package types

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

func NewShallowPrompt(id *string) ShallowPrompt {
	c := ShallowPrompt{}
	c.New(id)
	c.ShallowModel.ContentType = "shallowprompt"
	c, _ = ValidateShallowPrompt(c)
	return c
} 

func NewShallowPromptModelContent() models.ShallowContent {
	c := models.ShallowContent{}
	c.ShallowModel.ContentType = "shallowprompt"
	return c
}

func NewShallowPromptTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "shallowprompt"
	return c
}

type ShallowPrompt struct {
	ShallowModel
	ID 				PromptID 	`json:"id"`
	Name 			string 		`form:"name" json:"name"`
	Prompt 			string 		`form:"prompt" json:"prompt"`
	Domain 			string 		`form:"domain" json:"domain"`
	Category 		string 		`form:"category" json:"category"`
	Characters 		[]string 	`form:"characters" json:"characters"`
	SettingsModel 	string 		`form:"settings" json:"settings_model"`
}

func (c ShallowPrompt) Expand(ctx context.Context) (*Prompt, error) {
	r := Prompt{}
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
	r.Prompt = c.Prompt
	r.Domain = c.Domain
	r.Category = c.Category
	ss := ShallowSettings{}
	ss.ShallowModel.ID = c.SettingsModel
	settings, err := ss.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.SettingsModel = *settings
	return &r, nil
}

func (c *ShallowPrompt) New(id *string) {
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
	}
	c.ID = PromptID(c.ShallowModel.ID)
	c.ShallowModel.CreatedAt = time.Now()
	c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
	c.SettingsModel = ""
}


func (c ShallowPrompt) List(ctx context.Context) ([]ShallowPrompt, error) {
	content := NewPromptModelContent()
	content.Model.ContentType = "shallowprompt"
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	cuts := make([]ShallowPrompt, 0)
	for _, model := range contents {
		cut := NewShallowPrompt(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowPrompt", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowPrompt) ListBy(ctx context.Context, key string, value interface{}) ([]Prompt, error) {
	content := NewPromptModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	cuts := make([]Prompt, 0)
	for _, model := range contents {
		cut := NewPrompt(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Prompt", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *ShallowPrompt) Get(ctx context.Context) error {
	content := NewShallowPromptTypeContent()
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ShallowModel.ContentType = "shallowprompt"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Prompt", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c ShallowPrompt) Set(ctx context.Context) error {
	content := NewShallowPromptTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ID.String()
	err := content.Set(ctx)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowPrompt) Delete(ctx context.Context) error {
	content := NewShallowPromptTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.ContentDeleteError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowPrompt) GetDispositions(ctx context.Context) (ShallowPrompt, error) {
	var err error
	//disposition := NewShallowDisposition(nil)
//	c.SettingsModel.TemplateModel.AvailableDispositions, err = disposition.List(ctx)
	if err != nil {
		return c, merrors.ContentListError{Package: "types", Struct: "ShallowPrompt", Function: "GetDispositions"}.Wrap(err)
	}
	return c, nil
}

func (c ShallowPrompt) GetID() string {
	return c.ShallowModel.ID
}

func (c ShallowPrompt) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowPrompt) GetTable() string {
	return c.ShallowModel.Table
}

func (c ShallowPrompt) SetID() (ShallowPrompt, error) {
	var err error
	c.ID = PromptID(c.ShallowModel.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "prompt"}.Wrap(err)
	}
	return c, nil
}

func ValidateShallowPrompt(p ShallowPrompt) (ShallowPrompt, error) {
	if p.SettingsModel != "" {
		_, err := uuid.Parse(p.SettingsModel)
		if err != nil {
			return p, err
		}
	}
	return p, nil
}

func (c ShallowPrompt) Bind(e echo.Context) (ShallowPrompt, error) {
	var err error
	c.SettingsModel = e.FormValue("settings")
	return c, err
}

func (c ShallowPrompt) IsShallowModel() bool {
	return true
}
