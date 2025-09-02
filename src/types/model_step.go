package types

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type Step struct {
	Model
	Name          	string        	`form:"name" json:"name"`
	Order         	int           	`form:"order" json:"order"`
	DispositionID 	DispositionID 	`form:"disposition_id" json:"disposition_id"`
	WorkflowID    	WorkflowID    	`form:"workflow_id" json:"workflow_id"`
	JobID			JobID		  	`form:"job_id" json:"job_id"`
	Stats         	Stats         	`form:"stats" json:"stats"`
	Enabled       	Toggle        	`form:"enabled" json:"enabled"`
	Bypass        	Toggle        	`form:"bypass" json:"bypass"`
	SystemPrompt  	string        	`form:"system_prompt" json:"system_prompt"`
	PromptTemplate	string        	`form:"prompt_template" json:"prompt_template"`
	Node            string        	`form:"node" json:"node"`
	NodeType        string        	`form:"node_type" json:"node_type"`
	Dependency    	*Step		  	`form:"depencencies" json:"dependencies"`
	Validation      string			`form:"validation" json:"validation"`
}

func NewStep(id *string) Step {
	c := Step{}
	c.Model.New(id)
	c.Model.ContentType = "step"
	return c
}

func NewStepModelContent(id *string) models.Content {
	c := models.Content{}
	c.ID = uuid.NewString()
	c.Model.ID = c.ID
	if id != nil {
		c.ID = *id
		c.Model.ID = *id
	}
	c.Model.ContentType = "step"
	return c
}

func NewStepTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "step"
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
	sm.SystemPrompt = c.SystemPrompt
	sm.PromptTemplate = c.PromptTemplate
	sm.Node = c.Node
	sm.NodeType = c.NodeType
	sms = append(sms, sm)
	return sms
}

func (c *Step) New(id *string) {
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
	c.Model.ContentType = "step"
}

func (c Step) List(e echo.Context, orderby string) ([]Step, error) {
	content := NewStepModelContent(nil)
	content.Model.ContentType = "step"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]Step, 0)
	for _, model := range contents {
		cut := NewStep(nil)
		if err := json.Unmarshal([]byte(model.Content), &cut); err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Step", Function: "List"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c Step) ListBy(e echo.Context, orderby, key string, value interface{}) ([]Step, error) {
	content := NewStepModelContent(nil)
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]Step, 0)
	for _, model := range contents {
		cut := NewStep(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Step", Function: "ListBy"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *Step) Get(e echo.Context) error {
	content := NewStepTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "step"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Step", Function: "Get"}.Wrap(err).Log()
	}
	return nil
}

func (c Step) Set(e echo.Context, update bool) error {
	content := NewStepTypeContent()
	content.FromType(c, c.Model)
	content.ID = c.Model.ID
	list, err := c.ListBy(e, "order desc", "job_id", c.WorkflowID.String())
	if err != nil {
		return merrors.ContentListByError{}.Wrap(err).Log()
	}
	if order := e.FormValue("order"); order != "" {
		orderInt, err := strconv.Atoi(order)
		if err != nil {
			return err
		}
		if orderInt != 0 {
			c.Order = orderInt
		}
	}
	if len(list) == 0 {
		//first entry, does not need to be reordered
		if c.Order == 0 {
			c.Order = 1
		}
	}
	if c.Order == 0 {
		c.Order = len(list)
	}
	if err := content.Set(e, update); err != nil {
		return merrors.ContentSetError{}.Wrap(err).Log().Log()
	}
	return nil
}

func (c Step) Reorder(e echo.Context) error {
	list, err := c.List(e, "order desc")
	if err != nil {
		return merrors.ContentListError{}.Wrap(err).Log()
	}
	for _, step := range list {
		if step.Order < c.Order || step.ID == c.ID {
			continue
		} else {
			step.Order = step.Order + 1
		}
		if err := step.Set(e, true); err != nil {
			return merrors.ContentSetError{}.Wrap(err).Log()
		}
	}
	return nil
}

func (c Step) Delete(e echo.Context) error {
	content := NewStepTypeContent()
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(err).Log()
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
