package types

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/strrep"
)

type ShallowOllamaNode struct {
	ShallowModel
	ID 				string `json:"id"`
	Name 			string `form:"name" json:"name"`
	OllamaModel 	string `form:"model" json:"model"`
	SystemPrompt 	string `form:"system_prompt" json:"system_prompt"`
	Prompt 			string `form:"prompt" json:"prompt"`
	PromptTemplate  string `form:"prompt_template" json:"prompt_template"`
	ResponseModel   string `json:"response_model"`
	WorkflowID  	WorkflowID `form:"workflow_id" json:"workflow_id"`
	Enabled 		bool   `json:"enabled"`
	Bypass 			bool   `json:"bypass"`
	Output 			string `form:"output" json:"output"`
	Context			Context
}

func (c ShallowOllamaNode) Validate() params {
	valid := true
	if !c.ShallowModel.Validate() {
		valid = false
	}
	if c.ShallowModel.ContentType != "ollamanode" {
		valid = false
	}
	if c.ID == "" || c.ID != c.ShallowModel.ID {
		valid = false
	}
	if c.Name == "" || c.OllamaModel == "" {
		valid = false
	}
	if c.Prompt == "" && c.PromptTemplate == "" {
		valid = false
	}
	c.ShallowModel.Validated = valid
	return c
}

func (c ShallowOllamaNode) GetValidated() bool {
	return c.ShallowModel.Validated
}

func NewShallowOllamaNode(id *string) ShallowOllamaNode {
	c := ShallowOllamaNode{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "ollamanode"
	if c.ShallowModel.CreatedAt.IsZero() {
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
	} else {
		c.ShallowModel.UpdatedAt = time.Now()
	}
	return c
}

func (c ShallowOllamaNode) GetName() string {
	return c.Name
}

func (c ShallowOllamaNode) GetCommand() string {
	return ""
}

func (c ShallowOllamaNode) GetUser() string {
	return ""
}

func (c ShallowOllamaNode) GetHost() string {
	return ""
}

func (c ShallowOllamaNode) GetModel() string {
	return c.OllamaModel
}

func (c ShallowOllamaNode) GetSystemPrompt() string {
	return c.SystemPrompt
}

func (c ShallowOllamaNode) GetPrompt() string {
	return c.Prompt
}

func (c ShallowOllamaNode) GetPromptTemplate() string {
	return c.PromptTemplate
}

func (c ShallowOllamaNode) GetApiBase() string {
	return ""
}

func (c ShallowOllamaNode) GetApiTemplate() string {
	return ""
}
func (c ShallowOllamaNode) GetType() string {
	return "ollama_node"
}

func (c *ShallowOllamaNode) ParsePromptTemplate(ctx context.Context) error {
	if c.PromptTemplate != "" {
		if len(c.PromptTemplate) == 36 {
			id := c.PromptTemplate
			pt := NewPromptTemplate(&id)
			if err := pt.Get(ctx); err != nil {
				return merrors.GetPromptTemplateError{Package: "types", Struct:"ShallowOllamaNode", Function: "ParsePromptTemplate"}.Wrap(err)
			}
			msi := make(map[string]interface{})
			if err := json.Unmarshal([]byte(pt.Vars), &msi); err != nil {
				return merrors.JSONUnmarshallingError{Info: pt.Vars, Package: "types", Struct:"ShallowOllamaNode", Function: "ParsePromptTemplate"}.Wrap(err)
			}
			c.Prompt = strrep.Strrep(pt.Template, msi)
		}
	}
	return nil
}

func (c *ShallowOllamaNode) FromMSI(msi map[string]interface{}) error {
	var err error
	if p, ok := msi["params"].(map[string]interface{}); ok {
		if id, ok := p["id"].(string); ok {
			c.ShallowModel.ID = id
		}
		if createdAt, ok := p["CreatedAt"].(string); ok {
			c.ShallowModel.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
			if err != nil {
				return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct:"ShallowOllamaNode", Function: "FromMSI"}.Wrap(err)
			}
		}
		if updatedAt, ok := p["UpdatedAt"].(string); ok {
			c.ShallowModel.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
			if err != nil {
				return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct:"ShallowOllamaNode", Function: "FromMSI"}.Wrap(err)
			}
		}
		if ct, ok := p["ContentType"].(string); ok {
			c.ShallowModel.ContentType = ct
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

func (c *ShallowOllamaNode) Get(ctx context.Context) error {
	content := NewComfyNodeTypeContent()
	content.Model.ID = c.ShallowModel.ID
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.NodeGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(err)
	}
	return nil
}

func NewShallowOllamaNodeTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "ollamanode"
	return c
}

func (c ShallowOllamaNode) Delete(ctx context.Context) error {
	content := NewShallowSSHNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.ShallowModel.ID
	content.ID = c.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.NodeDeleteError{Info: c.ShallowModel.ID, Package: "types", Struct: "ollamanode", Function: "delete"}.Wrap(err)
	}
	return nil
}

func (c ShallowOllamaNode) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowOllamaNode) GetID() string {
	return c.ShallowModel.ID
}

func (c ShallowOllamaNode) Set(ctx context.Context) error {
	c.Validate()
	if !c.ShallowModel.Validated {
		return merrors.NodeValidationError{Package: "types", Struct: "node", Function: "set"}.Wrap(fmt.Errorf("validation failed"))
	}
	content := NewShallowOllamaNodeTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ID = c.ShallowModel.ID
	err := content.Set(ctx)
	if err != nil {
		return merrors.NodeSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

// func (c *ShallowOllamaNode) GetNodeFromWorkflow(id string, wf Workflow) {
// 	for _, v := range wf.OllamaNodesArrayModel {
// 		if id == v {
// 			c = &v
// 			break
// 		}
// 	}
// }

