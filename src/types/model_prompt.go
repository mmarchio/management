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

func NewPrompt(id *string) Prompt {
	c := Prompt{}
	c.New(id)
	c.Model.ContentType = "prompt"
	c, _ = ValidatePrompt(c)
	return c
} 

func NewPromptModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "prompt"
	return c
}

func NewPromptTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "prompt"
	return c
}

type Prompt struct {
	Model
	ID 			PromptID 	`json:"id"`
	Name 		string 		`form:"name" json:"name"`
	Prompt 		string 		`form:"prompt" json:"prompt"`
	Domain 		string 		`form:"domain" json:"domain"`
	Category 	string 		`form:"category" json:"category"`
	Characters 	[]Character `form:"characters" json:"characters"`
	SettingsModel 	Settings 	`form:"settings" json:"settings_model"`
}

func (c Prompt) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowPrompt{}
	sm.ShallowModel = sm.ShallowModel.FromTypeModel(c.Model)
	sm.ID = c.ID
	sm.Name = c.Name
	sm.Prompt = c.Prompt
	sm.Domain = c.Domain
	sm.Category = c.Category
	sm.SettingsModel = c.SettingsModel.ID
	sms = append(sms, c.SettingsModel.Pack()...)
	sms = append(sms, sm)
	return sms
}

func (c *Prompt) New(id *string) {
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.ID = PromptID(c.Model.ID)
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
	c.SettingsModel.New()
}


func (c Prompt) List(ctx context.Context) ([]Prompt, error) {
	content := NewPromptModelContent()
	content.Model.ContentType = "prompt"
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]Prompt, 0)
	for _, model := range contents {
		cut := NewPrompt(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Prompt", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c Prompt) ListBy(ctx context.Context, key string, value interface{}) ([]Prompt, error) {
	content := NewPromptModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err)
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

func (c *Prompt) Get(ctx context.Context) error {
	content := NewPromptTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "prompt"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Prompt", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c Prompt) Set(ctx context.Context) error {
	content := NewPromptTypeContent()
	content.FromType(c)
	content.Model.ID = c.ID.String()
	err := content.Set(ctx)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c Prompt) Delete(ctx context.Context) error {
	content := NewPromptTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c Prompt) GetDispositions(ctx context.Context) (Prompt, error) {
	var err error
	disposition := NewDisposition(nil)
	c.SettingsModel.TemplateModel.AvailableDispositions, err = disposition.List(ctx)
	if err != nil {
		return c, merrors.ContentListError{Package: "types", Struct: "Prompt", Function: "GetDispositions"}.Wrap(err)
	}
	return c, nil
}

func (c Prompt) GetID() string {
	return c.Model.ID
}

func (c Prompt) GetContentType() string {
	return c.Model.ContentType
}

func (c Prompt) GetTable() string {
	return c.Model.Table
}

func (c Prompt) SetID() (Prompt, error) {
	var err error
	c.ID = PromptID(c.Model.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "prompt"}.Wrap(err)
	}
	return c, nil
}

func ValidatePrompt(p Prompt) (Prompt, error) {
	var err error
	p.SettingsModel.GlobalBypassModel, err = ValidateSteps(p.SettingsModel.GlobalBypassModel, "global_bypass_")
	p.SettingsModel.RecurringModel = ValidateToggle(p.SettingsModel.RecurringModel, uuid.NewString(), "prompt_settings_", "recurring", "recurring")
	return p, err
}

func (c Prompt) Bind(e echo.Context) (Prompt, error) {
	var err error
	c.SettingsModel = c.SettingsModel.Bind(e)
	return c, err
}

func (c Prompt) Next(e echo.Context, ctx context.Context) (*models.Context, error) {
	systemContext, err := models.Context{}.GetCtx(ctx)
	if err != nil {
		return nil, merrors.ContextGetError{Package: "types", Struct: "Prompt", Function: "Next"}.Wrap(err)
	}
	return systemContext, nil
}