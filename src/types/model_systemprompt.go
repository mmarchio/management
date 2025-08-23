package types

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

func NewSystemPrompt(id *string) SystemPrompt {
	c := SystemPrompt{}
	c.Model.New(id)
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
	Name        string `form:"name" json:"name"`
	SelfModel   string `form:"self_model" json:"self_model"`
	TargetModel string `form:"target_model" json:"target_model"`
	Domain      string `form:"domain" json:"domain"`
	Prompt      string `form:"prompt" json:"prompt"`
}

func (c SystemPrompt) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowSystemPrompt{}
	sm.ShallowModel = sm.ShallowModel.FromTypeModel(c.Model)
	sm.ID = c.ID
	sm.Name = c.Name
	sm.Domain = c.Domain
	sm.SelfModel = c.SelfModel
	sm.TargetModel = c.TargetModel
	sm.Prompt = c.Prompt
	sms = append(sms, sm)
	return sms
}

func (c *SystemPrompt) New(id *string) {
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
	c.Model.ContentType = "systemprompt"
}

func (c SystemPrompt) List(e echo.Context) ([]SystemPrompt, error) {
	content := NewSystemPromptModelContent()
	content.Model.ContentType = "systemprompt"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]SystemPrompt, 0)
	for _, model := range contents {
		cut := NewSystemPrompt(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "SystemPrompt", Function: "List"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c SystemPrompt) ListBy(e echo.Context, key string, value interface{}) ([]SystemPrompt, error) {
	content := NewSystemPromptModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]SystemPrompt, 0)
	for _, model := range contents {
		cut := NewSystemPrompt(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "SystemPrompt", Function: "ListBy"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *SystemPrompt) Get(e echo.Context) error {
	content := NewSystemPromptTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "systemprompt"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "SystemPrompt", Function: "Get"}.Wrap(err).Log()
	}
	return nil
}

func (c SystemPrompt) Set(e echo.Context, update bool) error {
	content := NewSystemPromptTypeContent()
	content.FromType(c, c.Model)
	err := content.Set(e, update)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	return nil
}

func (c SystemPrompt) Delete(e echo.Context) error {
	content := NewSystemPromptTypeContent()
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(err).Log()
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
