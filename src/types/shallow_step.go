package types

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

func NewShallowStep(id, ct *string) ShallowStep {
	c := ShallowStep{}
	c.ShallowModel.New(id, ct)
	c.ShallowModel.ContentType = "systemprompt"
	return c
}

func NewShallowStepModelContent() models.ShallowContent {
	c := models.ShallowContent{}
	c.ShallowModel.ContentType = "systemprompt"
	return c
}

func NewShallowStepTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "systemprompt"
	return c
}

type ShallowStep struct {
	ShallowModel
	Name          	string        `form:"name" json:"name"`
	Order         	int           `form:"order" json:"order"`
	DispositionID 	DispositionID `form:"disposition_id" json:"disposition_id"`
	Stats         	string        `form:"stats" json:"stats"`
	Enabled       	string        `form:"enabled" json:"enabled"`
	Bypass        	string        `form:"bypass" json:"bypass"`
	SystemPrompt  	string        `form:"system_prompt" json:"system_prompt"`
	PromptTemplate	string        `form:"prompt_template" json:"prompt_template"`
	Node            string        `form:"node" json:"node"`
	NodeType        string        `form:"node_type" json:"node_type"`
	WorkflowID      string        `form:"workflow_id" json:"workflow_id"`
}

func (c ShallowStep) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowStep) Expand(e echo.Context) (*Step, error) {
	r := Step{}
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
	statsInput := Stats{}
	statsInput.EmbedModel.ID = c.Stats
	statsOutput, err := statsInput.Get(e)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err).Log()
	}
	enabledInput := Toggle{}
	enabledInput.ID = c.Enabled
	enabledInput.Model.ID = c.Enabled
	enabledOutput, err := enabledInput.Get(e)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err).Log()
	}
	bypassInput := Toggle{}
	bypassInput.ID = c.Bypass
	bypassInput.Model.ID = c.Bypass
	bypassOutput, err := bypassInput.Get(e)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err).Log()
	}

	r.Model = r.Model.FromShallowModel(c.ShallowModel)
	r.ID = c.ID
	r.Name = c.Name
	r.Order = c.Order
	r.DispositionID = c.DispositionID
	r.Stats = *statsOutput
	r.Enabled = *enabledOutput
	r.Bypass = *bypassOutput
	return &r, nil
}

func (c *ShallowStep) New(id *string) {
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
	}
	c.ShallowModel.CreatedAt = time.Now()
	c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
	c.ShallowModel.ContentType = "systemprompt"
}

func (c ShallowStep) List(e echo.Context) ([]ShallowStep, error) {
	content := NewShallowStepModelContent()
	content.ShallowModel.ContentType = "systemprompt"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(err).Log()
	}
	ct := "shallowsystemprompt"
	cuts := make([]ShallowStep, 0)
	for _, model := range contents {
		cut := NewShallowStep(nil, &ct)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowStep", Function: "List"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowStep) ListBy(e echo.Context, key string, value interface{}) ([]ShallowStep, error) {
	content := NewShallowStepModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{Info: c.ShallowModel.ContentType}.Wrap(err).Log()
	}
	ct := "shallowsystemprompt"
	cuts := make([]ShallowStep, 0)
	for _, model := range contents {
		cut := NewShallowStep(nil, &ct)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowStep", Function: "ListBy"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *ShallowStep) FromContent(content *ShallowContent) error {
	if content == nil {
		return merrors.ContentToTypeError{}.New("content is nil")
	}
	if err := json.Unmarshal([]byte(content.Content), c); err != nil {
		return merrors.JSONUnmarshallingError{}.Wrap(err).Log()
	}
	return nil
}

func (c *ShallowStep) Get(e echo.Context) error {
	content, err := c.ShallowModel.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err).Log()
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowStep", Function: "Get"}.Wrap(err).Log()
	}
	if err := c.FromContent(content); err != nil {
		return merrors.ContentToTypeError{}.Wrap(err).Log()
	}
	return nil
}

func (c ShallowStep) Set(e echo.Context, update bool) error {
	content := NewShallowStepTypeContent()
	content.FromType(c, c.ShallowModel)
	err := content.Set(e, update)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(err).Log()
	}
	return nil
}

func (c ShallowStep) Delete(e echo.Context) error {
	content := NewShallowStepTypeContent()
	content.FromType(c, c.ShallowModel)
	content.ShallowModel.ID = c.ShallowModel.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.ShallowModel.ID}.Wrap(err).Log()
	}
	return nil
}

func (c ShallowStep) GetID() string {
	return c.ShallowModel.ID
}

func (c ShallowStep) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowStep) GetTable() string {
	return c.ShallowModel.Table
}

func (c ShallowStep) IsShallowModel() bool {
	return true
}
