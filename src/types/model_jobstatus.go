package types

import (
	"context"
	"encoding/json"
	"time"

	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

func NewJobStatus() JobStatus {
	c := JobStatus{}
	c.Model.ContentType = "comfyuitemplate"
	return c
} 

func NewJobStatusModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "comfyuitemplate"
	return c
}

func NewJobStatusTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "comfyuitemplate"
	return c
}

type JobStatus struct {
	Model
	ID 				JobStatusID `json:"id"`
	JobID 			JobID 		`json:"job_id"`
	StatusType 		string 		`json:"status_type"`
	StatusContext 	Context 	`json:"status_context"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
}

func (c *JobStatus) New() {
	c.ID = c.ID.New()
	c.Model.ID = c.ID.String()
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
}

func (c JobStatus) List(ctx context.Context) ([]JobStatus, error) {
	content := NewJobStatusModelContent()
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]JobStatus, 0)
	for _, model := range contents {
		cut := NewJobStatus()
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "JobStatus", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c JobStatus) ListBy(ctx context.Context, key string, value interface{}) ([]JobStatus, error) {
	content := NewJobStatusModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]JobStatus, 0)
	for _, model := range contents {
		cut := JobStatus{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "JobStatus", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *JobStatus) Get(ctx context.Context) error {
	content := NewJobStatusTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "jobstatus"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "JobStatus", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c JobStatus) Delete(ctx context.Context) error {
	content := NewJobStatusTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c JobStatus) GetID() string {
	return c.Model.ID
}

func (c JobStatus) GetContentType() string {
	return c.Model.ContentType
}

func (c JobStatus) GetTable() string {
	return c.Model.Table
}