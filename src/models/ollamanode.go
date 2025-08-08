package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
)

type OllamaNode struct {
	Model
	ID 				string `json:"id"`
	Name 			string `form:"name" json:"name"`
	OllamaModel 	string `form:"model" json:"model"`
	SystemPrompt 	string `form:"system_prompt" json:"system_prompt"`
	Prompt 			string `form:"prompt" json:"prompt"`
	PromptTemplate  string `form:"prompt_template" json:"prompt_template"`
	Response        OllamaResponse `json:"response"`
	WorkflowID  	string `form:"workflow_id" json:"workflow_id"`
	Enabled 		bool   `json:"enabled"`
	Bypass 			bool   `json:"bypass"`
	Output 			string `form:"output" json:"output"`
	Context			Context
}

type ShallowOllamaNode struct {
	Model
	ID 				string `json:"id"`
	Name 			string `form:"name" json:"name"`
	OllamaModel 	string `form:"model" json:"model"`
	SystemPrompt 	string `form:"system_prompt" json:"system_prompt"`
	Prompt 			string `form:"prompt" json:"prompt"`
	PromptTemplate  string `form:"prompt_template" json:"prompt_template"`
	Response        string `json:"response"`
	WorkflowID  	string `form:"workflow_id" json:"workflow_id"`
	Enabled 		bool   `json:"enabled"`
	Bypass 			bool   `json:"bypass"`
	Output 			string `form:"output" json:"output"`
	Context			string `json:"context"`
}

func NewShallowOllamaNode(id *string) ShallowOllamaNode {
	c := ShallowOllamaNode{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_ollamanode"
	return c
}

func (c ShallowOllamaNode) Get(ctx context.Context, mode string) (*OllamaNode, *ShallowOllamaNode, error) {
	content := Content{ID: c.Model.ID}
	if err := content.Get(ctx); err != nil {
		return nil, nil, merrors.ContentGetError{Info: c.Model.ID, Package: "models", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(err)
	}
	if err := json.Unmarshal([]byte(content.Content), &c); err != nil {
		return nil, nil, merrors.JSONUnmarshallingError{Info: content.Content, Package: "models", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(err)
	}
	if mode == "shallow" {
		return nil, &c, nil
	}
	if mode == "full" {
		m := OllamaNode{}
		m.Model.ID = c.Model.ID
		m.Model.CreatedAt = c.Model.CreatedAt
		m.Model.UpdatedAt = c.Model.UpdatedAt
		m.Model.ContentType = c.Model.ContentType
		m.ID = c.Model.ID
		m.Name = c.Name
		m.OllamaModel = c.OllamaModel
		m.SystemPrompt = c.SystemPrompt
		m.Prompt = c.Prompt
		m.PromptTemplate = c.PromptTemplate
		ollamaresponseptr, _, err := NewShallowOllamaResponse(&c.Response).Get(ctx, "full")
		if err != nil {
			return nil, nil, err
		}
		if ollamaresponseptr != nil {
			m.Response = *ollamaresponseptr
		}
		m.WorkflowID = c.WorkflowID
		m.Enabled = c.Enabled
		m.Bypass = c.Bypass
		m.Output = c.Output
		contextptr, _, err := NewShallowContext(&c.Context).Get(ctx, "full")
		if err != nil {
			return nil, nil, err
		}
		if contextptr != nil {
			m.Context = *contextptr
		}
		return &m, nil, nil
	}
	return nil, nil, merrors.ContentGetError{Package: "models", Struct: "ShallowWorkflow", Function: "Get"}.Wrap(fmt.Errorf("unknown mode: %s", mode))
}
