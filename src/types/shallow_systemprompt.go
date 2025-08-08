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

func NewShallowSystemPrompt(id, ct *string) ShallowSystemPrompt {
	c := ShallowSystemPrompt{}
	c.ShallowModel.New(id, ct)
	c.ID = SystemPromptID(c.ShallowModel.ID)
	c.ShallowModel.ContentType = "systemprompt"
	return c
} 

func NewShallowSystemPromptModelContent() models.ShallowContent {
	c := models.ShallowContent{}
	c.ShallowModel.ContentType = "systemprompt"
	return c
}

func NewShallowSystemPromptTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "systemprompt"
	return c
}

type ShallowSystemPrompt struct {
	ShallowModel
	ID 		SystemPromptID `form:"id" json:"id"`
	Name 	string `form:"name" json:"name"`
	Domain 	string `form:"domain" json:"domain"`
	Prompt 	string `form:"prompt" json:"prompt"`
}

func (c ShallowSystemPrompt) Expand(ctx context.Context) (*SystemPrompt, error) {
	r := SystemPrompt{}
	if c.ShallowModel.CreatedAt.IsZero() && c.ShallowModel.ID != "" {
		sc, err := c.ShallowModel.Get(ctx)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err)
		}
		if err := json.Unmarshal([]byte(sc.Content), &r); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err)
		}
		return &r, nil
	}
	r.Model = r.Model.FromShallowModel(c.ShallowModel)
	r.ID = c.ID
	r.Name = c.Name
	r.Domain = c.Domain
	r.Prompt = c.Prompt
	return &r, nil
}

func (c *ShallowSystemPrompt) New(id *string) {
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
	}
	c.ID = SystemPromptID(c.ShallowModel.ID)
	c.ShallowModel.CreatedAt = time.Now()
	c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
	c.ShallowModel.ContentType = "systemprompt"
}


func (c ShallowSystemPrompt) List(ctx context.Context) ([]ShallowSystemPrompt, error) {
	content := NewShallowSystemPromptModelContent()
	content.ShallowModel.ContentType = "systemprompt"
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	ct := "shallowsystemprompt"
	cuts := make([]ShallowSystemPrompt, 0)
	for _, model := range contents {
		cut := NewShallowSystemPrompt(nil, &ct)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowSystemPrompt", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowSystemPrompt) ListBy(ctx context.Context, key string, value interface{}) ([]ShallowSystemPrompt, error) {
	content := NewShallowSystemPromptModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	ct := "shallowsystemprompt"
	cuts := make([]ShallowSystemPrompt, 0)
	for _, model := range contents {
		cut := NewShallowSystemPrompt(nil, &ct)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowSystemPrompt", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *ShallowSystemPrompt) FromContent(content *ShallowContent) error {
	if content == nil {
		return merrors.ContentToTypeError{}.Wrap(fmt.Errorf("content is nil"))
	}
	if err := json.Unmarshal([]byte(content.Content), c); err != nil {
		return merrors.JSONUnmarshallingError{}.Wrap(err)
	}
	return nil	
}

func (c *ShallowSystemPrompt) Get(ctx context.Context) error {
	content, err := c.ShallowModel.Get(ctx)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowSystemPrompt", Function: "Get"}.Wrap(err)
	}
	if err := c.FromContent(content); err != nil {
		return merrors.ContentToTypeError{}.Wrap(err)
	}
	return nil
}

func (c ShallowSystemPrompt) Set(ctx context.Context) error {
	content := NewShallowSystemPromptTypeContent()
	content.FromType(c)
	err := content.Set(ctx)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowSystemPrompt) Delete(ctx context.Context) error {
	content := NewShallowSystemPromptTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.ContentDeleteError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowSystemPrompt) GetID() string {
	return c.ShallowModel.ID
}

func (c ShallowSystemPrompt) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowSystemPrompt) GetTable() string {
	return c.ShallowModel.Table
}

func (c ShallowSystemPrompt) SetID() (ShallowSystemPrompt, error) {
	var err error
	c.ID = SystemPromptID(c.ShallowModel.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "systemprompt"}.Wrap(err)
	}
	return c, nil
}