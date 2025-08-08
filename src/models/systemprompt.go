package models

import (
	"time"

	"github.com/google/uuid"
)

type SystemPrompt struct {
	Model
	ID string
	Name string
	Domain string
	Prompt string
}

type ShallowSystemPrompt struct {
	Model
	ID string
	Name string
	Domain string
	Prompt string
}

func NewShallowSystemPrompt(id *string) ShallowSystemPrompt {
	c := ShallowSystemPrompt{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_systemprompt"
	return c
}

