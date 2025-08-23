package models

import (
	"time"

	"github.com/google/uuid"
)

type PromptTemplate struct {
	Model
	Name		string  `form:"name" json:"name"`
	Template 	string 	`form:"template" json:"template"`
	Vars 		string 	`form:"vars" json:"vars"`
}

type ShallowPromptTemplate struct {
	ShallowModel
	Name		string  `form:"name" json:"name"`
	Template 	string 	`form:"template" json:"template"`
	Vars 		string 	`form:"vars" json:"vars"`
}

func NewShallowPromptTemplate(id *string) ShallowPromptTemplate {
	c := ShallowPromptTemplate{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_prompttemplate"
	return c
}
