package types

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

func NewShallowSystemPrompt(id, ct *string) ShallowSystemPrompt {
	c := ShallowSystemPrompt{}
	c.ShallowModel.New(id, ct)
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
	Name        string `form:"name" json:"name"`
	Domain      string `form:"domain" json:"domain"`
	SelfModel   string `form:"self_model" json:"self_model"`
	TargetModel string `form:"target_model" json:"target_model"`
	Prompt      string `form:"prompt" json:"prompt"`
}

func (c ShallowSystemPrompt) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowSystemPrompt) Expand(e echo.Context) (*SystemPrompt, error) {
	r := SystemPrompt{}
	if c.ShallowModel.CreatedAt.IsZero() && c.ShallowModel.ID != "" {
		sc, err := c.ShallowModel.Get(e)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err).Log()
		}
		if err := json.Unmarshal([]byte(sc.Content), &r); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
		return &r, nil
	}
	r.Model = r.Model.FromShallowModel(c.ShallowModel)
	r.ID = c.ID
	r.Name = c.Name
	r.Domain = c.Domain
	r.SelfModel = c.SelfModel
	r.TargetModel = c.TargetModel
	r.Prompt = c.Prompt
	return &r, nil
}

func (c *ShallowSystemPrompt) New(id *string) {
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
	}
	c.ShallowModel.CreatedAt = time.Now()
	c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
	c.ShallowModel.ContentType = "systemprompt"
}

func (c ShallowSystemPrompt) List(e echo.Context) ([]ShallowSystemPrompt, error) {
	content := NewShallowSystemPromptModelContent()
	content.ShallowModel.ContentType = "systemprompt"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(err).Log()
	}
	ct := "shallowsystemprompt"
	cuts := make([]ShallowSystemPrompt, 0)
	for _, model := range contents {
		cut := NewShallowSystemPrompt(nil, &ct)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowSystemPrompt", Function: "List"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowSystemPrompt) ListBy(e echo.Context, key string, value interface{}) ([]ShallowSystemPrompt, error) {
	content := NewShallowSystemPromptModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{Info: c.ShallowModel.ContentType}.Wrap(err).Log()
	}
	ct := "shallowsystemprompt"
	cuts := make([]ShallowSystemPrompt, 0)
	for _, model := range contents {
		cut := NewShallowSystemPrompt(nil, &ct)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowSystemPrompt", Function: "ListBy"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *ShallowSystemPrompt) FromContent(content *ShallowContent) error {
	if content == nil {
		return merrors.ContentToTypeError{}.New("content is nil")
	}
	if err := json.Unmarshal([]byte(content.Content), c); err != nil {
		return merrors.JSONUnmarshallingError{}.Wrap(err).Log()
	}
	return nil
}

func (c *ShallowSystemPrompt) Get(e echo.Context) error {
	content, err := c.ShallowModel.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err).Log()
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowSystemPrompt", Function: "Get"}.Wrap(err).Log()
	}
	if err := c.FromContent(content); err != nil {
		return merrors.ContentToTypeError{}.Wrap(err).Log()
	}
	return nil
}

func (c ShallowSystemPrompt) Set(e echo.Context, update bool) error {
	content := NewShallowSystemPromptTypeContent()
	content.FromType(c, c.ShallowModel)
	err := content.Set(e, update)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(err).Log()
	}
	return nil
}

func (c ShallowSystemPrompt) Delete(e echo.Context) error {
	content := NewShallowSystemPromptTypeContent()
	content.FromType(c, c.ShallowModel)
	content.ShallowModel.ID = c.ShallowModel.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.ShallowModel.ID}.Wrap(err).Log()
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

func (c ShallowSystemPrompt) IsShallowModel() bool {
	return true
}
