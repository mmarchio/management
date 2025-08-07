package types

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
)

type ComfyNode struct {
	Model
	ID 				string 					`json:"id"`
	Name 			string 					`form:"name" json:"name"`
	Prompt          string                  `form:"prompt" json:"prompt"`
	APIBase 		string 					`form:"api_base" json:"api_base"`
	APITemplate 	string 					`form:"api_template" json:"api_template"`
	TemplateValues  map[string]interface{} 	`json:"template_values"`
	WorkflowID  	WorkflowID 				`form:"workflow_id" json:"workflow_id"`
	Type 			string 					`form:"type" json:"type"`
	Enabled 		bool   					`json:"enabled"`
	Bypass 			bool   					`json:"bypass"`
	Output 			string 					`form:"output" json:"output"`
}

type ShallowComfyNode struct {
	ShallowModel
	ID 				string 					`json:"id"`
	Name 			string 					`form:"name" json:"name"`
	Prompt          string                  `form:"prompt" json:"prompt"`
	APIBase 		string 					`form:"api_base" json:"api_base"`
	APITemplate 	string 					`form:"api_template" json:"api_template"`
	TemplateValues  map[string]interface{} 	`json:"template_values"`
	WorkflowID  	WorkflowID 				`form:"workflow_id" json:"workflow_id"`
	Type 			string 					`form:"type" json:"type"`
	Enabled 		bool   					`json:"enabled"`
	Bypass 			bool   					`json:"bypass"`
	Output 			string 					`form:"output" json:"output"`
}

func (c ComfyNode) Validate() params {
	valid := true
	if !c.Model.Validate() {
		valid = false
	}
	if c.Model.ContentType != "comfynode" {
		valid = false
	}
	if c.ID == "" || c.ID != c.Model.ID {
		valid = false
	}
	if c.Name == "" || c.APIBase == "" || c.APITemplate == "" {
		valid = false
	}
	c.Model.Validated = valid
	return c
}

func (c ComfyNode) GetValidated() bool {
	return c.Model.Validated
}

func (c ComfyNode) GetName() string {
	return c.Name
}

func (c ComfyNode) GetCommand() string {
	return ""
}

func (c ComfyNode) GetUser() string {
	return ""
}

func (c ComfyNode) GetHost() string {
	return ""
}

func (c ComfyNode) GetModel() string {
	return ""
}

func (c ComfyNode) GetSystemPrompt() string {
	return ""
}

func (c ComfyNode) GetPrompt() string {
	return ""
}

func (c ComfyNode) GetPromptTemplate() string {
	return ""
}

func (c ComfyNode) GetApiBase() string {
	return c.APIBase
}

func (c ComfyNode) GetApiTemplate() string {
	return c.APITemplate
}
func (c ComfyNode) GetType() string {
	return "comfy_node"
}

func (c *ComfyNode) FromMSI(msi map[string]interface{}) error {
	var err error
	if id, ok := msi["id"].(string); ok {
		c.Model.ID = id
	}
	if createdAt, ok := msi["CreatedAt"].(string); ok {
		c.Model.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct:"ComfyNode", Function: "FromMSI"}.Wrap(err)
		}
	}
	if updatedAt, ok := msi["UpdatedAt"].(string); ok {
		c.Model.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct:"ComfyNode", Function: "FromMSI"}.Wrap(err)
		}
	}
	if ct, ok := msi["ContentType"].(string); ok {
		c.Model.ContentType = ct
	}
	if name, ok := msi["name"].(string); ok {
		c.Name = name
	}
	if ab, ok := msi["api_base"].(string); ok {
		c.APIBase = ab
	}
	if at, ok := msi["api_template"].(string); ok {
		c.APITemplate = at
	}
	return nil
}

func (c *ComfyNode) Get(ctx context.Context) error {
	content := NewComfyNodeTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "comfynode"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.NodeGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "node", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c *ComfyNode) GetShallow(ctx context.Context) error {
	content := NewComfyNodeTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "comfynode"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.NodeGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "node", Function: "Get"}.Wrap(err)
	}
	return nil
}

func NewComfyNodeTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "comfynode"
	return c
}

func NewShallowComfyNodeTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "comfynode"
	return c
}

func (c ComfyNode) Delete(ctx context.Context) error {
	content := NewSSHNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	content.ID = c.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.NodeDeleteError{Info: c.Model.ID, Package: "types", Struct: "comfynode", Function: "delete"}.Wrap(err)
	}
	return nil
}

func (c ComfyNode) GetContentType() string {
	return c.Model.ContentType
}

func (c ComfyNode) GetID() string {
	return c.Model.ID
}

func NewComfyNode(id *string) ComfyNode {
	c := ComfyNode{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "comfynode"
	if c.Model.CreatedAt.IsZero() {
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	} else {
		c.Model.UpdatedAt = time.Now()
	}
	return c
}

func (c ComfyNode) Set(ctx context.Context) error {
	c.Validate()
	if !c.Model.Validated {
		return merrors.NodeValidationError{Package: "types", Struct: "node", Function: "set"}.Wrap(fmt.Errorf("validation failed"))
	}
	content := NewComfyNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	content.ID = c.Model.ID
	err := content.Set(ctx)
	if err != nil {
		return merrors.NodeSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

