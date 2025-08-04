package types

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

func NewJobRun(id *string) JobRun {
	c := JobRun{}
	c.New(id)
	c.Model.ContentType = "jobrun"
	return c
} 

func NewJobRunModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "jobrun"
	return c
}

func NewJobRunTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "jobrun"
	return c
}

type JobRun struct {
	Model
	ID 					RunID			 `json:"id"`
	JobID 				JobID			 `json:"job_id"`
	WorkflowID			WorkflowID  	 `json:"workflow_id"`
	Context 			Context			 `json:"context"`
	TruncatedContext    TruncatedContext `json:"truncated_context"`
	Settings 			Settings		 `json:"settings"`
	Disposition         Disposition 	 `json:"disposition"`
	Tokens 				int64			 `json:"tokens"`
	LatestStatusType 	string 			 `json:"latest_status_type"`
	LatestStatusValue 	string 			 `json:"latest_status_value"`
}

func (c *JobRun) New(id *string) {
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.ID = RunID(c.Model.ID)
	c.ID = c.ID.New()
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
}

func (c JobRun) List(ctx context.Context) ([]JobRun, error) {
	content := NewJobRunModelContent()
	content.Model.ContentType = "jobrun"
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.JobRunListError{Info: c.Model.ContentType, Package: "types", Struct: "JobRun", Function: "List"}.Wrap(err)
	}
	cuts := make([]JobRun, 0)
	for _, model := range contents {
		cut := JobRun{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "JobRun", Function: "List"}.Wrap(err)
		}
		cut.Context.SetCtx(ctx)
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c JobRun) ListBy(ctx context.Context, key string, value interface{}) ([]JobRun, error) {
	content := NewJobRunModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.JobRunListError{Info: c.Model.ContentType, Package: "types", Struct: "JobRun", Function: "ListBy"}.Wrap(err)
	}
	cuts := make([]JobRun, 0)
	for _, model := range contents {
		cut := NewJobRun(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "JobRun", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *JobRun) Get(ctx context.Context) error {
	content := NewJobRunTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "jobrun"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.JobRunGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "JobRun", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c *JobRun) FindBy(ctx context.Context) error {
	var err error
	content := NewJobRunTypeContent()
	if !c.ID.IsNil() {
		content, err = content.FindBy(ctx, "id", c.ID.String())
	}
	if !c.JobID.IsNil() {
		content, err = content.FindBy(ctx, "job_id", c.JobID.String())
	}
	if !c.WorkflowID.IsNil() {
		content, err = content.FindBy(ctx, "workflow_id", c.WorkflowID.String())
	}
	if err != nil {
		return merrors.JobRunGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "JobRun", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c *JobRun) CustomQuery(ctx context.Context, write bool, q string, vars ...any) ([]JobRun, error) {
	content := NewJobRunTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "jobrun"
	content.ID = content.Model.ID
	if write {
		content.FromType(c)
		_, err := content.CustomQuery(ctx, write, q, vars)
		if err != nil {
			return nil, merrors.JobRunCustomQueryError{Info: c.Model.ID}.Wrap(err)
		}
		return nil, nil
	}
	res, err := content.CustomQuery(ctx, write, q, vars)
	if err != nil {
		if err != nil {
			return nil, merrors.JobRunCustomQueryError{Info: c.Model.ID}.Wrap(err).BubbleCode()
		}
	}
	r := make([]JobRun, 0)
	for _, t := range res {
		jr := NewJobRun(nil)
		if err = json.Unmarshal([]byte(t.Content), &jr); err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "JobRun", Function: "CustomQuery"}.Wrap(err)
		}
		r = append(r, jr)
	}
	return r, nil
}

func (c JobRun) Set(ctx context.Context) error {
	content := NewJobRunTypeContent()
	content.FromType(c)
	content.ContentType = "jobrun"
	content.ID = c.Model.ID
	err := content.Set(ctx)
	if err != nil {
		return merrors.JobRunSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c JobRun) Delete(ctx context.Context) error {
	content := NewJobRunTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.JobRunDeleteError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c JobRun) GetID() string {
	return c.Model.ID
}

func (c JobRun) GetContentType() string {
	return c.Model.ContentType
}

func (c JobRun) GetTable() string {
	return c.Model.Table
}

