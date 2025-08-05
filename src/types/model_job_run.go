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

func NewShallowJobRun(id *string) ShallowJobRun {
	c := ShallowJobRun{}
	c.New(id)
	c.ShallowModel.ContentType = "shallowjobrun"
	return c
} 

func NewJobRunModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "jobrun"
	return c
}

func NewShallowJobRunModelContent() models.ShallowContent {
	c := models.ShallowContent{}
	c.ShallowModel.ContentType = "shallowjobrun"
	return c
}

func NewJobRunTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "jobrun"
	return c
}

func NewShallowJobRunTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "shallowjobrun"
	return c
}

type JobRun struct {
	Model
	ID 						RunID			 `json:"id"`
	JobID 					JobID			 `json:"job_id"`
	WorkflowID				WorkflowID  	 `json:"workflow_id"`
	ContextModel 			Context			 `json:"context_model"`
	TruncatedContextModel  	TruncatedContext `json:"truncated_context_model"`
	SettingsModel 			Settings		 `json:"settings_model"`
	DispositionModel        Disposition 	 `json:"disposition_model"`
	Tokens 					int64			 `json:"tokens"`
	LatestStatusType 		string 			 `json:"latest_status_type"`
	LatestStatusValue 		string 			 `json:"latest_status_value"`
}

type ShallowJobRun struct {
	ShallowModel
	ID 						RunID			 `json:"id"`
	JobID 					JobID			 `json:"job_id"`
	WorkflowID				WorkflowID  	 `json:"workflow_id"`
	ContextModel 			string			 `json:"context_model"`
	TruncatedContextModel  	string 			 `json:"truncated_context_model"`
	SettingsModel 			string			 `json:"settings_model"`
	DispositionModel        string		 	 `json:"disposition_model"`
	Tokens 					int64			 `json:"tokens"`
	LatestStatusType 		string 			 `json:"latest_status_type"`
	LatestStatusValue 		string 			 `json:"latest_status_value"`
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

func (c *ShallowJobRun) New(id *string) {
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
	}
	c.ID = RunID(c.ShallowModel.ID)
	c.ID = c.ID.New()
	c.ShallowModel.CreatedAt = time.Now()
	c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
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
		cut.ContextModel.SetCtx(ctx)
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowJobRun) List(ctx context.Context) ([]ShallowJobRun, error) {
	content := NewJobRunModelContent()
	content.Model.ContentType = "jobrun"
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.JobRunListError{Info: c.ShallowModel.ContentType, Package: "types", Struct: "JobRun", Function: "List"}.Wrap(err)
	}
	cuts := make([]ShallowJobRun, 0)
	for _, model := range contents {
		cut := ShallowJobRun{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "JobRun", Function: "List"}.Wrap(err)
		}
		contextModel := Context{}
		if err := contextModel.GetCtx(ctx); err != nil {
			return nil, err
		}
		ctx, err = contextModel.SetCtx(ctx)
		if err != nil {
			return nil, err
		}
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

func (c ShallowJobRun) ListBy(ctx context.Context, key string, value interface{}) ([]ShallowJobRun, error) {
	content := NewShallowJobRunModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.JobRunListError{Info: c.ShallowModel.ContentType, Package: "types", Struct: "JobRun", Function: "ListBy"}.Wrap(err)
	}
	cuts := make([]ShallowJobRun, 0)
	for _, model := range contents {
		cut := NewShallowJobRun(nil)
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

func (c *ShallowJobRun) Get(ctx context.Context) error {
	content := NewShallowJobRunTypeContent()
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ShallowModel.ContentType = "shallowjobrun"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.JobRunGetError{Info: c.ShallowModel.ID}.Wrap(err)
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

func (c *ShallowJobRun) FindBy(ctx context.Context) error {
	var err error
	content := NewShallowJobRunTypeContent()
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
		return merrors.JobRunGetError{Info: c.ShallowModel.ID}.Wrap(err)
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

func (c *ShallowJobRun) CustomQuery(ctx context.Context, write bool, q string, vars ...any) ([]ShallowJobRun, error) {
	content := NewShallowJobRunTypeContent()
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ShallowModel.ContentType = "jobrun"
	content.ID = content.ShallowModel.ID
	if write {
		content.FromType(c)
		_, err := content.CustomQuery(ctx, write, q, vars)
		if err != nil {
			return nil, merrors.JobRunCustomQueryError{Info: c.ShallowModel.ID}.Wrap(err)
		}
		return nil, nil
	}
	res, err := content.CustomQuery(ctx, write, q, vars)
	if err != nil {
		if err != nil {
			return nil, merrors.JobRunCustomQueryError{Info: c.ShallowModel.ID}.Wrap(err).BubbleCode()
		}
	}
	r := make([]ShallowJobRun, 0)
	for _, t := range res {
		jr := NewShallowJobRun(nil)
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

func (c ShallowJobRun) Set(ctx context.Context) error {
	content := NewShallowJobRunTypeContent()
	content.FromType(c)
	content.ContentType = "shallowjobrun"
	content.ID = c.ShallowModel.ID
	err := content.Set(ctx)
	if err != nil {
		return merrors.JobRunSetError{Info: c.ShallowModel.ID}.Wrap(err)
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

func (c ShallowJobRun) Delete(ctx context.Context) error {
	content := NewShallowJobRunTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.JobRunDeleteError{Info: c.ShallowModel.ID}.Wrap(err)
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

func (c ShallowJobRun) GetID() string {
	return c.ShallowModel.ID
}

func (c ShallowJobRun) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowJobRun) GetTable() string {
	return c.ShallowModel.Table
}

