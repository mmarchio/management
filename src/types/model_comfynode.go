package types

import (
	"time"

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