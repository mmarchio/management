package types

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

func NewSystemPrompt(id *string) SystemPrompt {
	c := SystemPrompt{}
	c.Model.New(id)
	c.ID = SystemPromptID(c.Model.ID)
	c.Model.ContentType = "systemprompt"
	return c
} 

func NewSystemPromptModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "systemprompt"
	return c
}

func NewSystemPromptTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "systemprompt"
	return c
}

type SystemPrompt struct {
	Model
	ID 		SystemPromptID `form:"id" json:"id"`
	Name 	string `form:"name" json:"name"`
	Domain 	string `form:"domain" json:"domain"`
	Prompt 	string `form:"prompt" json:"prompt"`
}

func (c *SystemPrompt) New(id *string) {
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.ID = SystemPromptID(c.Model.ID)
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
	c.Model.ContentType = "systemprompt"
}


func (c SystemPrompt) List(ctx context.Context) ([]SystemPrompt, error) {
	content := NewSystemPromptModelContent()
	content.Model.ContentType = "systemprompt"
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.SystemPromptListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]SystemPrompt, 0)
	for _, model := range contents {
		cut := NewSystemPrompt(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "SystemPrompt", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c SystemPrompt) ListBy(ctx context.Context, key string, value interface{}) ([]SystemPrompt, error) {
	content := NewSystemPromptModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.SystemPromptListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]SystemPrompt, 0)
	for _, model := range contents {
		cut := NewSystemPrompt(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "SystemPrompt", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *SystemPrompt) Get(ctx context.Context) error {
	content := NewSystemPromptTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "systemprompt"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.SystemPromptGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "SystemPrompt", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c SystemPrompt) Set(ctx context.Context) error {
	content := NewSystemPromptTypeContent()
	content.FromType(c)
	err := content.Set(ctx)
	if err != nil {
		return merrors.SystemPromptSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c SystemPrompt) Delete(ctx context.Context) error {
	content := NewSystemPromptTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.SystemPromptDeleteError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c SystemPrompt) GetID() string {
	return c.Model.ID
}

func (c SystemPrompt) GetContentType() string {
	return c.Model.ContentType
}

func (c SystemPrompt) GetTable() string {
	return c.Model.Table
}

func (c SystemPrompt) SetID() (SystemPrompt, error) {
	var err error
	c.ID = SystemPromptID(c.Model.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "systemprompt"}.Wrap(err)
	}
	return c, nil
}