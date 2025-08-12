package types

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type Step struct {
	Model
	ID StepID 					`form:"id" json:"id"`
	Name string 				`form:"name" json:"name"`
	Order int 					`form:"order" json:"order"`
	DispositionID DispositionID `form:"disposition_id" json:"disposition_id"`
	Stats Stats					`form:"stats" json:"stats"`
	Enabled Toggle 				`form:"enabled" json:"enabled"`
	Bypass Toggle				`form:"bypass" json:"bypass"`
}

func NewStep(id *string) Step {
	c := Step{}
	c.Model.New(id)
	c.ID = StepID(c.Model.ID)
	c.Model.ContentType = "step"
	return c
} 

func NewStepModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "systemprompt"
	return c
}

func NewStepTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "systemprompt"
	return c
}

func (c Step) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowStep{}
	sm.ShallowModel = sm.ShallowModel.FromTypeModel(c.Model)
	sm.ID = c.ID
	sm.Name = c.Name
	sm.Order = c.Order
	sm.DispositionID = c.DispositionID
	sm.Stats = c.Stats.EmbedModel.ID
	sm.Enabled = c.Enabled.Model.ID
	sm.Bypass = c.Bypass.Model.ID
	sms = append(sms, sm)
	return sms
}

func (c *Step) New(id *string) {
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.ID = StepID(c.Model.ID)
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
	c.Model.ContentType = "systemprompt"
}


func (c Step) List(e echo.Context) ([]Step, error) {
	content := NewStepModelContent()
	
	content.Model.ContentType = "systemprompt"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]Step, 0)
	for _, model := range contents {
		cut := NewStep(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Step", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c Step) ListBy(e echo.Context, key string, value interface{}) ([]Step, error) {
	content := NewStepModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]Step, 0)
	for _, model := range contents {
		cut := NewStep(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Step", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *Step) Get(e echo.Context) error {
	content := NewStepTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "systemprompt"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Step", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c Step) Set(e echo.Context) error {
	content := NewStepTypeContent()
	content.FromType(c)
	err := content.Set(e)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c Step) Delete(e echo.Context) error {
	content := NewStepTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c Step) GetID() string {
	return c.Model.ID
}

func (c Step) GetContentType() string {
	return c.Model.ContentType
}

func (c Step) GetTable() string {
	return c.Model.Table
}

func (c Step) SetID() (Step, error) {
	var err error
	c.ID = StepID(c.Model.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "systemprompt"}.Wrap(err)
	}
	return c, nil
}