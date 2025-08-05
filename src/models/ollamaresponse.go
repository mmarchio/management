package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
)

type OllamaResponse struct {
	Model
	OllamaModel string `json:"model"`
	CreatedAt string `json:"created_at"`
	Done bool `json:"done"`
	Response string `json:"response"`
}

type ShallowOllamaResponse struct {
	Model
	OllamaModel string `json:"model"`
	CreatedAt string `json:"created_at"`
	Done bool `json:"done"`
	Response string `json:"response"`
}

func (c OllamaResponse) GetResponse() string {
	return c.Response
}

func NewShallowOllamaResponse(id *string) ShallowOllamaResponse {
	c := ShallowOllamaResponse{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_ollamaresponse"
	return c
}

func NewOllamaResponse(id *string) OllamaResponse {
	c := OllamaResponse{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "ollamaresponse"
	return c
}


func (c ShallowOllamaResponse) Get(ctx context.Context, mode string) (*OllamaResponse, *ShallowOllamaResponse, error) {
	content := Content{ID: c.Model.ID}
	if err := content.Get(ctx); err != nil {
		return nil, nil, merrors.ShallowWorkflowGetError{Info: c.Model.ID, Package: "models", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(err)
	}
	if err := json.Unmarshal([]byte(content.Content), &c); err != nil {
		return nil, nil, merrors.JSONUnmarshallingError{Info: content.Content, Package: "models", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(err)
	}
	if mode == "shallow" {
		return nil, &c, nil
	}
	if mode == "full" {
		m := OllamaResponse{}
		m.Model = c.Model
		m.OllamaModel = c.OllamaModel
		m.CreatedAt = c.CreatedAt
		m.Done = c.Done
		m.Response = c.Response
		return &m, nil, nil
	}
	return nil, nil, merrors.ShallowOllamaResponseGetError{Package: "models", Struct: "ShallowOllamaResponse", Function: "Get"}.Wrap(fmt.Errorf("unknown mode: %s", mode))
} 
