package models

import (
	"time"

	"github.com/google/uuid"
)

type ComfyUITemplate struct {
	Model
	Name 		string `form:"name" json:"name"`
	Endpoint 	string `form:"endpoint" json:"endpoint"`
	Base 		string `form:"base" json:"base"`
	Template 	string `form:"template" json:"template"`
}

func (c *ComfyUITemplate) New() {
}

type ShallowComfyUITemplate struct {
	Model
	Name 		string `form:"name" json:"name"`
	Endpoint 	string `form:"endpoint" json:"endpoint"`
	Base 		string `form:"base" json:"base"`
	Template 	string `form:"template" json:"template"`
}

func NewShallowComfyUITemplate(id *string) ShallowComfyUITemplate {
	c := ShallowComfyUITemplate{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_comfyuitemplate"
	return c
}
