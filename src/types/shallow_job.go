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

func NewShallowJob(id *string) ShallowJob {
	c := ShallowJob{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
	}
	c.ID = JobID(c.ShallowModel.ID)
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	c.ShallowModel.ContentType = "shallowjob"
	
	return c
} 

func NewShallowJobModelContent() models.ShallowContent {
	c := models.ShallowContent{}
	c.ShallowModel.ContentType = "shallowjob"
	return c
}

func NewShallowJobTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "shallowjob"
	return c
}

type ShallowJob struct {
	ShallowModel
	ID 				JobID 		`json:"id"`
	PromptID 		PromptID 	`json:"prompt_id"`
	WorkflowID		WorkflowID	`form:"workflow_id" json:"workflow_id"`
	Recurring   	bool        `json:"recurring"`
	Interval    	int64       `form:"interval" json:"interval"`
	LastCompleted 	time.Time 	`json:"last_completed"`
}

func (c *ShallowJob) New() {
	c.ID = c.ID.New()
	c.ShallowModel.ID = c.ID.String()
	c.ShallowModel.CreatedAt = time.Now()
	c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
	c.ShallowModel.Table = "jobs"
	c.ShallowModel.Columns = "id, created_at, updated_at, prompt_id"
	c.ShallowModel.Values = "$1, $2, $3, $4"
	c.ShallowModel.Conflict = "DO NOTHING"
}

func (c ShallowJob) List(ctx context.Context) ([]ShallowJob, error) {
	content := NewJobModelContent()
	content.Model.ContentType = "job"
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	cuts := make([]ShallowJob, 0)
	for _, model := range contents {
		cut := ShallowJob{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Job", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowJob) ListBy(ctx context.Context, key string, value interface{}) ([]ShallowJob, error) {
	content := NewShallowJobModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	cuts := make([]ShallowJob, 0)
	for _, model := range contents {
		cut := NewShallowJob(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Job", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *ShallowJob) Get(ctx context.Context) error {
	var err error
	content := NewShallowJobTypeContent()
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ID = content.ShallowModel.ID
	content.ShallowModel.ContentType = "job"
	content, err = content.Get(ctx)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	if err := json.Unmarshal([]byte(content.Content), c); err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Job", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c *ShallowJob) FindBy(ctx context.Context, key, value string) (ShallowJob, error) {
	var err error
	job := ShallowJob{}
	content := NewShallowJobTypeContent()
	content.ShallowModel.ID = c.ShallowModel.ID
	content, err = content.FindBy(ctx, key, value) 
	if err != nil {
		return job, merrors.ContentFindByError{Info: fmt.Sprintf("key: %s, value: %s", key, value)}.Wrap(err)
	}
	if err = json.Unmarshal([]byte(content.Content), &job); err != nil {
		if _, ok := err.(merrors.WrappedError); ok {
			return job, merrors.NilContentError{Info: content.Content, Package: "types", Struct: "ShallowJob", Function: "FindBy"}.Wrap(err).BubbleCode()
		}
		return job, merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowJob", Function: "FindBy"}.Wrap(err)
	}
	c = &job
	return job, nil
}

func (c ShallowJob) Set(ctx context.Context) error {
	content := NewShallowJobTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ID = c.ShallowModel.ID
	err := content.Set(ctx)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowJob) Delete(ctx context.Context) error {
	content := NewJobTypeContent()
	content.FromType(c)
	content.Model.ID = c.ShallowModel.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.ContentDeleteError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowJob) GetID() string {
	return c.ShallowModel.ID
}

func (c ShallowJob) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowJob) GetTable() string {
	return c.ShallowModel.Table
}