package types

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/strrep"
)

var wg sync.WaitGroup

type OllamaNode struct {
	Model
	ID 				string `json:"id"`
	Name 			string `form:"name" json:"name"`
	OllamaModel 	string `form:"model" json:"model"`
	SystemPrompt 	string `form:"system_prompt" json:"system_prompt"`
	Prompt 			string `form:"prompt" json:"prompt"`
	PromptTemplate  string `form:"prompt_template" json:"prompt_template"`
	ResponseModel   OllamaResponse `json:"response_model"`
	WorkflowID  	WorkflowID `form:"workflow_id" json:"workflow_id"`
	Enabled 		bool   `json:"enabled"`
	Bypass 			bool   `json:"bypass"`
	Output 			string `form:"output" json:"output"`
	Context			Context
}

func (c OllamaNode) Validate() params {
	valid := true
	if !c.Model.Validate() {
		valid = false
	}
	if c.Model.ContentType != "ollamanode" {
		valid = false
	}
	if c.ID == "" || c.ID != c.Model.ID {
		valid = false
	}
	if c.Name == "" || c.OllamaModel == "" {
		valid = false
	}
	if c.Prompt == "" && c.PromptTemplate == "" {
		valid = false
	}
	c.Model.Validated = valid
	return c
}

func (c OllamaNode) GetValidated() bool {
	return c.Model.Validated
}

func NewOllamaNode(id *string) OllamaNode {
	c := OllamaNode{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "ollamanode"
	if c.Model.CreatedAt.IsZero() {
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	} else {
		c.Model.UpdatedAt = time.Now()
	}
	return c
}

func (c OllamaNode) GetName() string {
	return c.Name
}

func (c OllamaNode) GetCommand() string {
	return ""
}

func (c OllamaNode) GetUser() string {
	return ""
}

func (c OllamaNode) GetHost() string {
	return ""
}

func (c OllamaNode) GetModel() string {
	return c.OllamaModel
}

func (c OllamaNode) GetSystemPrompt() string {
	return c.SystemPrompt
}

func (c OllamaNode) GetPrompt() string {
	return c.Prompt
}

func (c OllamaNode) GetPromptTemplate() string {
	return c.PromptTemplate
}

func (c OllamaNode) GetApiBase() string {
	return ""
}

func (c OllamaNode) GetApiTemplate() string {
	return ""
}
func (c OllamaNode) GetType() string {
	return "ollama_node"
}

func (c *OllamaNode) ParsePromptTemplate(ctx context.Context) error {
	if c.PromptTemplate != "" {
		if len(c.PromptTemplate) == 36 {
			id := c.PromptTemplate
			pt := NewPromptTemplate(&id)
			if err := pt.Get(ctx); err != nil {
				return merrors.GetPromptTemplateError{Package: "types", Struct:"OllamaNode", Function: "ParsePromptTemplate"}.Wrap(err)
			}
			msi := make(map[string]interface{})
			if err := json.Unmarshal([]byte(pt.Vars), &msi); err != nil {
				return merrors.JSONUnmarshallingError{Info: pt.Vars, Package: "types", Struct:"OllamaNode", Function: "ParsePromptTemplate"}.Wrap(err)
			}
			c.Prompt = strrep.Strrep(pt.Template, msi)
		}
	}
	return nil
}

func (c *OllamaNode) FromMSI(msi map[string]interface{}) error {
	var err error
	if p, ok := msi["params"].(map[string]interface{}); ok {
		if id, ok := p["id"].(string); ok {
			c.Model.ID = id
		}
		if createdAt, ok := p["CreatedAt"].(string); ok {
			c.Model.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
			if err != nil {
				return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct:"OllamaNode", Function: "FromMSI"}.Wrap(err)
			}
		}
		if updatedAt, ok := p["UpdatedAt"].(string); ok {
			c.Model.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
			if err != nil {
				return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct:"OllamaNode", Function: "FromMSI"}.Wrap(err)
			}
		}
		if ct, ok := p["ContentType"].(string); ok {
			c.Model.ContentType = ct
		}
		if id, ok := p["ID"].(string); ok {
			c.ID = id
		}
		if name, ok := p["name"].(string); ok {
			c.Name = name
		}
		if prompt, ok := p["prompt"].(string); ok {
			c.Prompt = prompt
		}
		if pt, ok := p["prompt_template"].(string); ok {
			c.PromptTemplate = pt
		}
		if sp, ok := p["system_prompt"].(string); ok {
			c.SystemPrompt = sp
		}
	}
	return nil
}

func (c *OllamaNode) Call(ctx context.Context, respchan chan OllamaResponseModel, errchan chan error) error {
	start := time.Now()
	if c.OllamaModel == "" {
		errchan <- fmt.Errorf("OllamaNode:OllamaModel is nil\n")
	}
	oreq := OllamaRequest{
		Model: c.OllamaModel,
		Prompt: c.Prompt,
		System: c.SystemPrompt,
//		Format: "json",
		Stream: false,
		KeepAlive: "1h",
	}
	data, err := json.Marshal(oreq)
	if err != nil {
		return merrors.JSONMarshallingError{Package: "types", Struct:"OllamaNode", Function: "Call"}.Wrap(err)
	}
	fmt.Printf("req data %#v\n", oreq)
	c.ResponseModel = OllamaResponseModel{}
	semaphore := make(chan struct{}, 1)
	ctr := 1
	for !c.ResponseModel.Done {
		reader := bytes.NewReader(data)
		req, err := http.NewRequest("POST", "http://172.17.0.1:7869/api/generate", reader)
		if err != nil {
			return merrors.HTTPRequestError{Package: "types", Struct:"OllamaNode", Function: "Call"}.Wrap(err)
		}
		wg.Add(1)
		go worker(c, req, respchan, errchan, semaphore)
		workerResp := <- respchan
		c.ResponseModel.Done = workerResp.Done
		c.ResponseModel.ResponseModel = workerResp.ResponseModel
		time.Sleep(5*time.Second)
		ctr++
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("start time: %s\n", start.Format(time.RFC3339))
	fmt.Printf("end time: %s\n", end.Format(time.RFC3339))
	fmt.Printf("time elapsed: %f\n", time.Since(start).Seconds())
	return nil
}

func worker(c *OllamaNode, req *http.Request, respchan chan OllamaResponseModel, errchan chan error, semaphore chan struct{}) {
	defer wg.Done()
	defer func() { <-semaphore }()

	semaphore <- struct{}{}
	var err error
	cresp := OllamaResponseModel{}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("request err: %s\n", err.Error())
	}
	if resp != nil {
		err = json.NewDecoder(resp.Body).Decode(&cresp)
		if err != nil {
			errchan <- err
		}
		c.ResponseModel.ResponseModel = cresp.ResponseModel
		c.ResponseModel.Done = cresp.Done
		respchan <- c.ResponseModel
	}
}

func (c *OllamaNode) Get(ctx context.Context) error {
	content := NewComfyNodeTypeContent()
	content.Model.ID = c.Model.ID
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.NodeGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ollamanode", Function: "Get"}.Wrap(err)
	}
	return nil
}

func NewOllamaNodeTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "ollamanode"
	return c
}

func (c OllamaNode) Delete(ctx context.Context) error {
	content := NewSSHNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	content.ID = c.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.NodeDeleteError{Info: c.Model.ID, Package: "types", Struct: "ollamanode", Function: "delete"}.Wrap(err)
	}
	return nil
}

func (c OllamaNode) GetContentType() string {
	return c.Model.ContentType
}

func (c OllamaNode) GetID() string {
	return c.Model.ID
}

func (c OllamaNode) Set(ctx context.Context) error {
	c.Validate()
	if !c.Model.Validated {
		return merrors.NodeValidationError{Package: "types", Struct: "node", Function: "set"}.Wrap(fmt.Errorf("validation failed"))
	}
	content := NewOllamaNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	content.ID = c.Model.ID
	err := content.Set(ctx)
	if err != nil {
		return merrors.NodeSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c *OllamaNode) GetNodeFromWorkflow(id string, wf Workflow) {
	for _, v := range wf.OllamaNodes {
		if id == v.Model.ID {
			c = &v
			break
		}
	}
}

func (c OllamaNode) Exec(ctx context.Context) error {
	fmt.Printf("ollama node found: %s\n", c.Name)
	start := time.Now()
	if c.SystemPrompt != "" && len(c.SystemPrompt) == 36 {
		spid := c.SystemPrompt
		sp := NewSystemPrompt(&spid)
		if err := sp.Get(ctx); err != nil {
			return err
		}
		c.SystemPrompt = sp.Prompt
	}
	ResponseModel := make(chan OllamaResponseModel, 1000)
	Error := make(chan error, 1)
	go c.Call(ctx, ResponseModel, Error)
	go func(){
		if len(Error) > 0 {
			err := <- Error
			fmt.Printf("err from errchan: %s\n", err.Error())
		}
	}()
	wg.Add(1)
	// var resp types.OllamaResponseModel
	// go run(c, resp, ResponseModel, Error, 1, semaphore)
	// wg.Add(1)
	wg.Wait()

	wf.Nodes[i].Params = p
	if err := wf.Set(ctx); err != nil {
		return err
	}
	end := time.Now()
	c.Context.GetResearchPromptModel.Start = start
	c.Context.GetResearchPromptModel.End = end 
	c.Context.GetResearchPromptModel.Status = "done"
	c.Context.GetResearchPromptModel.Output = c.ResponseModel.ResponseModel
	if err := c.Set(ctx); err != nil {
		return err
	}
}