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
	ShallowModel
	Name 		string `form:"name" json:"name"`
	Endpoint 	string `form:"endpoint" json:"endpoint"`
	Base 		string `form:"base" json:"base"`
	Template 	string `form:"template" json:"template"`
}

func NewShallowComfyUITemplate(id *string) ShallowComfyUITemplate {
	c := ShallowComfyUITemplate{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_comfyuitemplate"
	return c
}
