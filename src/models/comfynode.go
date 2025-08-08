package models

import (
	"time"

	"github.com/google/uuid"
)

type ComfyNode struct {
	Model
	ID 				string 					`json:"id"`
	Name 			string 					`form:"name" json:"name"`
	Prompt          string                  `form:"prompt" json:"prompt"`
	APIBase 		string 					`form:"api_base" json:"api_base"`
	APITemplate 	string 					`form:"api_template" json:"api_template"`
	TemplateValues  map[string]interface{} 	`json:"template_values"`
	WorkflowID  	string 					`form:"workflow_id" json:"workflow_id"`
	Type 			string 					`form:"type" json:"type"`
	Enabled 		bool   					`json:"enabled"`
	Bypass 			bool   					`json:"bypass"`
	Output 			string 					`form:"output" json:"output"`
}

type ShallowComfyNode struct {
	Model
	ID 				string 					`json:"id"`
	Name 			string 					`form:"name" json:"name"`
	Prompt          string                  `form:"prompt" json:"prompt"`
	APIBase 		string 					`form:"api_base" json:"api_base"`
	APITemplate 	string 					`form:"api_template" json:"api_template"`
	TemplateValues  map[string]interface{} 	`json:"template_values"`
	WorkflowID  	string 					`form:"workflow_id" json:"workflow_id"`
	Type 			string 					`form:"type" json:"type"`
	Enabled 		bool   					`json:"enabled"`
	Bypass 			bool   					`json:"bypass"`
	Output 			string 					`form:"output" json:"output"`
}

func NewShallowComfyNode(id *string) ShallowComfyNode {
	c := ShallowComfyNode{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_comfynode"
	return c
}

