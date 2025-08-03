package types

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/strrep"
)

var wg sync.WaitGroup

type OllamaNode struct {
	Model
	ID 				string `json:"id"`
	Name 			string `form:"name" json:"name"`
	OllamaModel 	string `form:"ollama_model" json:"ollama_model"`
	SystemPrompt 	string `form:"system_prompt" json:"system_prompt"`
	Prompt 			string `form:"prompt" json:"prompt"`
	PromptTemplate  string `form:"prompt_template" json:"prompt_template"`
	Response        OllamaResponse `json:"response"`
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

func (c *OllamaNode) Call(ctx context.Context, respchan chan OllamaResponse, errchan chan error) error {
	start := time.Now()
	if c.OllamaModel == "" {
		return fmt.Errorf("OllamaNode:OllamaModel is nil\n")
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
	c.Response = OllamaResponse{}
	semaphore := make(chan struct{}, 1)
	ctr := 1
	for !c.Response.Done {
		reader := bytes.NewReader(data)
		req, err := http.NewRequest("POST", "http://172.17.0.1:7869/api/generate", reader)
		if err != nil {
			return merrors.HTTPRequestError{Package: "types", Struct:"OllamaNode", Function: "Call"}.Wrap(err)
		}
		wg.Add(1)
		go worker(c, req, respchan, errchan, semaphore)
		workerResp := <- respchan
		c.Response.Done = workerResp.Done
		c.Response.Response = workerResp.Response
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
		c.Response.Response = cresp.Response
		c.Response.Done = cresp.Done
		respchan <- c.Response
	}
}

