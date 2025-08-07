package types

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

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

func (c ShallowComfyNode) Validate() params {
	valid := true
	if !c.ShallowModel.Validate() {
		valid = false
	}
	if c.ShallowModel.ContentType != "comfynode" {
		valid = false
	}
	if c.ID == "" || c.ID != c.ShallowModel.ID {
		valid = false
	}
	if c.Name == "" || c.APIBase == "" || c.APITemplate == "" {
		valid = false
	}
	c.ShallowModel.Validated = valid
	return c
}

func (c ShallowComfyNode) GetValidated() bool {
	return c.ShallowModel.Validated
}

func (c ShallowComfyNode) GetName() string {
	return c.Name
}

func (c ShallowComfyNode) GetCommand() string {
	return ""
}

func (c ShallowComfyNode) GetUser() string {
	return ""
}

func (c ShallowComfyNode) GetHost() string {
	return ""
}

func (c ShallowComfyNode) GetModel() string {
	return ""
}

func (c ShallowComfyNode) GetSystemPrompt() string {
	return ""
}

func (c ShallowComfyNode) GetPrompt() string {
	return ""
}

func (c ShallowComfyNode) GetPromptTemplate() string {
	return ""
}

func (c ShallowComfyNode) GetApiBase() string {
	return c.APIBase
}

func (c ShallowComfyNode) GetApiTemplate() string {
	return c.APITemplate
}
func (c ShallowComfyNode) GetType() string {
	return "comfy_node"
}

func (c *ShallowComfyNode) FromMSI(msi map[string]interface{}) error {
	var err error
	if id, ok := msi["id"].(string); ok {
		c.ShallowModel.ID = id
	}
	if createdAt, ok := msi["CreatedAt"].(string); ok {
		c.ShallowModel.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct:"ComfyNode", Function: "FromMSI"}.Wrap(err)
		}
	}
	if updatedAt, ok := msi["UpdatedAt"].(string); ok {
		c.ShallowModel.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct:"ComfyNode", Function: "FromMSI"}.Wrap(err)
		}
	}
	if ct, ok := msi["ContentType"].(string); ok {
		c.ShallowModel.ContentType = ct
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

func (c *ShallowComfyNode) Get(ctx context.Context) error {
	content := NewComfyNodeTypeContent()
	content.Model.ID = c.ShallowModel.ID
	content.Model.ContentType = "shallowcomfynode"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.NodeGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowComfyNode", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c *ShallowComfyNode) GetShallow(ctx context.Context) error {
	content := NewComfyNodeTypeContent()
	content.Model.ID = c.ShallowModel.ID
	content.Model.ContentType = "shallowcomfynode"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.NodeGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowComfyNode", Function: "Get"}.Wrap(err)
	}
	return nil
}

func NewShallowComfyNodeTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "shallowcomfynode"
	return c
}

func NewShallowComfyNodeModelContent() models.ShallowContent {
	c := models.ShallowContent{}
	c.ShallowModel.ContentType = "shallowcomfynode"
	return c
}

func (c ShallowComfyNode) Delete(ctx context.Context) error {
	content := NewShallowComfyNodeTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ID = c.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.NodeDeleteError{Info: c.ShallowModel.ID, Package: "types", Struct: "ShallowComfyNode", Function: "delete"}.Wrap(err)
	}
	return nil
}

func (c ShallowComfyNode) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowComfyNode) GetID() string {
	return c.ShallowModel.ID
}

func NewShallowComfyNode(id *string) ShallowComfyNode {
	c := ShallowComfyNode{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "comfynode"
	if c.ShallowModel.CreatedAt.IsZero() {
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
	} else {
		c.ShallowModel.UpdatedAt = time.Now()
	}
	return c
}

func (c ShallowComfyNode) Set(ctx context.Context) error {
	c.Validate()
	if !c.ShallowModel.Validated {
		return merrors.NodeValidationError{Package: "types", Struct: "node", Function: "set"}.Wrap(fmt.Errorf("validation failed"))
	}
	content := NewComfyNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.ShallowModel.ID
	content.ID = c.ShallowModel.ID
	err := content.Set(ctx)
	if err != nil {
		return merrors.NodeSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

