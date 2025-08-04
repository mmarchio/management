package models

import (
	"time"

	"github.com/google/uuid"
)

type PromptTemplate struct {
	Model
	ID 			string 	`form:"id" json:"id"`
	Name		string  `form:"name" json:"name"`
	Template 	string 	`form:"template" json:"template"`
	Vars 		string 	`form:"vars" json:"vars"`
}

type ShallowPromptTemplate struct {
	Model
	ID 			string 	`form:"id" json:"id"`
	Name		string  `form:"name" json:"name"`
	Template 	string 	`form:"template" json:"template"`
	Vars 		string 	`form:"vars" json:"vars"`
}

func NewShallowPromptTemplate(id *string) ShallowPromptTemplate {
	c := ShallowPromptTemplate{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_prompttemplate"
	return c
}
