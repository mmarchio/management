package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

func NewJob(id *string) Job {
	c := Job{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	c.Model.ContentType = "job"

	return c
}

func NewJobModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "job"
	return c
}

func NewJobTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "job"
	return c
}

type Job struct {
	Model
	PromptID      PromptID   `json:"prompt_id"`
	WorkflowID    WorkflowID `form:"workflow_id" json:"workflow_id"`
	Recurring     bool       `json:"recurring"`
	Interval      int64      `form:"interval" json:"interval"`
	LastCompleted time.Time  `json:"last_completed"`
}

func (c Job) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowJob{}
	sm.ShallowModel = sm.ShallowModel.FromTypeModel(c.Model)
	sm.ID = c.ID
	sm.PromptID = c.PromptID
	sm.WorkflowID = c.WorkflowID
	sm.Recurring = c.Recurring
	sm.Interval = c.Interval
	sm.LastCompleted = c.LastCompleted
	sms = append(sms, sm)
	return sms
}

func (c *Job) New() {
	c.Model.ID = uuid.NewString()
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
	c.Model.Table = "jobs"
	c.Model.Columns = "id, created_at, updated_at, prompt_id"
	c.Model.Values = "$1, $2, $3, $4"
	c.Model.Conflict = "DO NOTHING"
}

func (c Job) List(e echo.Context) ([]Job, error) {
	content := NewJobModelContent()
	content.Model.ContentType = "job"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]Job, 0)
	for _, model := range contents {
		cut := Job{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Job", Function: "List"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c Job) ListBy(e echo.Context, key string, value interface{}) ([]Job, error) {
	content := NewJobModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]Job, 0)
	for _, model := range contents {
		cut := NewJob(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Job", Function: "ListBy"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *Job) Get(e echo.Context) error {
	var err error
	content := NewJobTypeContent()
	content.Model.ID = c.Model.ID
	content.ID = content.Model.ID
	content.Model.ContentType = "job"
	content, err = content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	if err := json.Unmarshal([]byte(content.Content), c); err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Job", Function: "Get"}.Wrap(err).Log()
	}
	return nil
}

func (c *Job) FindBy(e echo.Context, key, value string) (Job, error) {
	var err error
	job := Job{}
	content := NewJobTypeContent()
	content.Model.ID = c.Model.ID
	content, err = content.FindBy(e, key, value)
	if err != nil {
		return job, merrors.ContentFindByError{Info: fmt.Sprintf("key: %s, value: %s", key, value)}.Wrap(err).Log()
	}
	if err = json.Unmarshal([]byte(content.Content), &job); err != nil {
		if _, ok := err.(merrors.WrappedError); ok {
			return job, merrors.NilContentError{Info: content.Content, Package: "types", Struct: "Job", Function: "FindBy"}.Wrap(err).Log().BubbleCode()
		}
		return job, merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Job", Function: "FindBy"}.Wrap(err).Log()
	}
	c = &job
	return job, nil
}

func (c Job) Set(e echo.Context, update bool) error {
	content := NewJobTypeContent()
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	content.ID = c.Model.ID
	err := content.Set(e, update)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	return nil
}

func (c Job) Delete(e echo.Context) error {
	content := NewJobTypeContent()
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(err).Log()
	}
	return nil
}

func (c Job) GetID() string {
	return c.Model.ID
}

func (c Job) GetContentType() string {
	return c.Model.ContentType
}

func (c Job) GetTable() string {
	return c.Model.Table
}
