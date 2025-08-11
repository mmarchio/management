package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/logger"
	"github.com/mmarchio/management/strrep"
)

var wg sync.WaitGroup

type OllamaNode struct {
	Model
	ID 				string `json:"id"`
	Name 			string `form:"name" json:"name"`
	OllamaModel 	string `form:"ollama_model" json:"model"`
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

func (c OllamaNode) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowOllamaNode{}
	sm.ShallowModel = sm.ShallowModel.FromTypeModel(c.Model)
	sm.ID = c.ID
	sm.OllamaModel = c.OllamaModel
	sm.SystemPrompt = c.SystemPrompt
	sm.Prompt = c.Prompt
	sm.PromptTemplate = c.PromptTemplate
	sm.WorkflowID = c.WorkflowID
	sm.Enabled = c.Enabled
	sm.Output = c.Output
	sms = append(sms, sm)
	return sms
}

func (c OllamaNode) Validate() params {
	msg := ""
	valid := true
	if !c.Model.Validate() {
		msg = "failed model validation"
		valid = false
	}
	if c.Model.ContentType != "ollamanode" {
		msg = "failed content type validation: "+c.Model.ContentType
		valid = false
	}
	if c.ID == "" || c.ID != c.Model.ID {
		msg = "failed id match"
		valid = false
	}
	if c.Name == "" || c.OllamaModel == "" {
		msg = fmt.Sprintf("failed name: %s ollama_model: %s", c.Name, c.OllamaModel)
		valid = false
	}
	if c.Prompt == "" && c.PromptTemplate == "" {
		msg = fmt.Sprintf("failed prompt: %s prompt_template: %s", c.Prompt, c.PromptTemplate)
		valid = false
	}
	c.Model.Validated = valid
	if !c.Model.Validated {
		fmt.Printf("validation failure: %s\n", msg)
	}
	return c
}

func (c OllamaNode) ValidateV2() bool {
	msg := ""
	valid := true
	if !c.Model.Validate() {
		msg = "failed model validation"
		valid = false
	}
	if c.Model.ContentType != "ollamanode" {
		msg = "failed content type validation: "+c.Model.ContentType
		valid = false
	}
	if c.ID == "" || c.ID != c.Model.ID {
		msg = "failed id match"
		valid = false
	}
	if c.Name == "" || c.OllamaModel == "" {
		msg = fmt.Sprintf("failed name: %s ollama_model: %s", c.Name, c.OllamaModel)
		valid = false
	}
	if c.Prompt == "" && c.PromptTemplate == "" {
		msg = fmt.Sprintf("failed prompt: %s prompt_template: %s", c.Prompt, c.PromptTemplate)
		valid = false
	}
	fmt.Printf("%s", msg)
	return valid
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

func (c *OllamaNode) ParsePromptTemplate(e echo.Context) error {
	if c.PromptTemplate != "" {
		if len(c.PromptTemplate) == 36 {
			id := c.PromptTemplate
			pt := NewPromptTemplate(&id)
			if err := pt.Get(e); err != nil {
				return merrors.ContentGetError{Package: "types", Struct:"OllamaNode", Function: "ParsePromptTemplate"}.Wrap(err)
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

func (c *OllamaNode) Call(e echo.Context, respchan chan OllamaResponse, errchan chan error) error {
	log, ok := e.Get("logger").(logger.LoggingContext)
	if !ok {
		return fmt.Errorf("logger is nil")
	}
	log.Flogger("Get called")
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
	log.Flogger(fmt.Sprintf("req data %#v", oreq))
	c.ResponseModel = OllamaResponse{}
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
		c.ResponseModel.Response = workerResp.Response
		time.Sleep(5*time.Second)
		ctr++
	}
	wg.Wait()
	end := time.Now()

	log.Flogger(fmt.Sprintf("start time: %s\n", start.Format(time.RFC3339)))
	log.Flogger(fmt.Sprintf("end time: %s\n", end.Format(time.RFC3339)))
	log.Flogger(fmt.Sprintf("time elapsed: %f\n", time.Since(start).Seconds()))
	return nil
}

func worker(c *OllamaNode, req *http.Request, respchan chan OllamaResponse, errchan chan error, semaphore chan struct{}) {
	defer wg.Done()
	defer func() { <-semaphore }()

	semaphore <- struct{}{}
	var err error
	cresp := OllamaResponse{}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("request err: %s\n", err.Error())
	}
	if resp != nil {
		err = json.NewDecoder(resp.Body).Decode(&cresp)
		if err != nil {
			errchan <- err
		}
		c.ResponseModel.Response = cresp.Response
		c.ResponseModel.Done = cresp.Done
		respchan <- c.ResponseModel
	}
}

func (c *OllamaNode) Get(e echo.Context) error {
	content := NewComfyNodeTypeContent()
	content.Model.ID = c.Model.ID
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err)
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

func (c OllamaNode) Delete(e echo.Context) error {
	content := NewSSHNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	content.ID = c.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID, Package: "types", Struct: "ollamanode", Function: "delete"}.Wrap(err)
	}
	return nil
}

func (c OllamaNode) GetContentType() string {
	return c.Model.ContentType
}

func (c OllamaNode) GetID() string {
	return c.Model.ID
}

func (c OllamaNode) Set(e echo.Context) error {
	if !c.ValidateV2() {
		fmt.Printf("validated: %t\nid: %s\ncreated_at: %s\nupdated_at: %s\ncontent_type: %s\nname: %s\nollama_model: %s\nsystem_prompt: %s\nprompt: %s\nprompt_template: %s", c.Model.Validated, c.Model.ID, c.Model.CreatedAt, c.Model.UpdatedAt, c.Model.ContentType, c.Name, c.OllamaModel, c.SystemPrompt, c.Prompt, c.PromptTemplate)
		return merrors.ContentValidationError{Info: fmt.Sprintf("ollamanode: %#v", c), Package: "types", Struct: "OllamaNode", Function: "Set"}.Wrap(fmt.Errorf("validation failed"))
	}
	content := NewOllamaNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	content.ID = c.Model.ID
	err := content.Set(e)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c *OllamaNode) GetNodeFromWorkflow(id string, wf Workflow) {
	for _, v := range wf.OllamaNodesArrayModel {
		if id == v.Model.ID {
			c = &v
			break
		}
	}
}

func (c OllamaNode) Exec(e echo.Context) error {
	log, ok := e.Get("logger").(logger.LoggingContext)
	if !ok {
		return fmt.Errorf("logger is nil")
	}
	log.Flogger("Exec called")
	start := time.Now()
	if c.SystemPrompt != "" && len(c.SystemPrompt) == 36 {
		spid := c.SystemPrompt
		sp := NewSystemPrompt(&spid)
		if err := sp.Get(e); err != nil {
			return err
		}
		c.SystemPrompt = sp.Prompt
	}
	Response := make(chan OllamaResponse, 1000)
	Error := make(chan error, 1)
	go c.Call(e, Response, Error)
	go func(){
		if len(Error) > 0 {
			err := <- Error
			log.Flogger("err from errchan: %s\n", err.Error())
		}
	}()
	wg.Add(1)
	// var resp types.OllamaResponseModel
	// go run(c, resp, ResponseModel, Error, 1, semaphore)
	// wg.Add(1)
	wg.Wait()
	resp := <-Response
	c.Output = resp.Response
	end := time.Now()
	c.Context.GetResearchPromptModel.Start = start
	c.Context.GetResearchPromptModel.End = end 
	c.Context.GetResearchPromptModel.Status = "done"
	c.Context.GetResearchPromptModel.Output = c.ResponseModel.Response
	if err := c.Set(e); err != nil {
		return err
	}
	return nil
}

func (c OllamaNode) List(e echo.Context) ([]OllamaNode, error) {
	content := NewOllamaNodeTypeContent()
	content.Model.ContentType = "ollamanode"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]OllamaNode, 0)
	for _, model := range contents {
		cut := OllamaNode{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "OllamaNode", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c OllamaNode) ListBy(e echo.Context, key string, value interface{}) ([]OllamaNode, error) {
	content := NewOllamaNodeTypeContent()
	content.Model.ContentType = "ollamanode"
	list, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{Info: fmt.Sprintf("{\"%s\":\"%s\"}", key, value), Package: "types", Struct: "OllamaNode", Function: "ListBy"}.Wrap(err)
	}
	cuts := make([]OllamaNode, 0)
	for _, model := range list {
		cut := OllamaNode{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "OllamaNode", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}