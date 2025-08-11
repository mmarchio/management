package types

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

func NewShallowJobRun(id *string) ShallowJobRun {
	c := ShallowJobRun{}
	c.New(id)
	c.ShallowModel.ContentType = "shallowjobrun"
	return c
} 

func NewShallowJobRunModelContent() models.ShallowContent {
	c := models.ShallowContent{}
	c.ShallowModel.ContentType = "shallowjobrun"
	return c
}

func NewShallowJobRunTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "shallowjobrun"
	return c
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

func (c ShallowJobRun) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(nil, err)
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowJobRun) Expand(e echo.Context) (*JobRun, error) {
	r := JobRun{}
	if c.ShallowModel.CreatedAt.IsZero() && c.ShallowModel.ID != "" {
		sc, err := c.ShallowModel.Get(e)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(nil, err)
		}
		if err := json.Unmarshal([]byte(sc.Content), &r); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(nil, err)
		}
		return &r, nil
	}
	r.Model = r.Model.FromShallowModel(c.ShallowModel)
	r.ID = c.ID
	r.JobID = c.JobID
	r.WorkflowID = c.WorkflowID
	context, err := r.Model.GetCtx(e)
	if err != nil {
		return nil, merrors.ContextGetError{}.Wrap(nil, err)
	}
	r.ContextModel = *context
	truncated, err := r.ContextModel.Truncate(e)
	if err != nil {
		return nil, merrors.ContextGetError{}.Wrap(nil, err)
	}
	r.TruncatedContextModel = *truncated
	ss := ShallowSettings{}
	ss.ShallowModel.ID = c.SettingsModel
	settings, err := ss.Expand(e)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(nil, err)
	}
	r.SettingsModel = *settings
	sd := ShallowDisposition{}
	sd.ShallowModel.ID = c.DispositionModel
	disposition, err := sd.Expand(e)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(nil, err)
	}
	r.DispositionModel = *disposition
	r.Tokens = c.Tokens
	r.LatestStatusType = c.LatestStatusType
	r.LatestStatusValue = c.LatestStatusValue
	return &r, nil
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

func (c ShallowJobRun) List(e echo.Context) ([]ShallowJobRun, error) {
	content := NewJobRunModelContent()
	content.Model.ContentType = "jobrun"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType, Package: "types", Struct: "JobRun", Function: "List"}.Wrap(nil, err)
	}
	cuts := make([]ShallowJobRun, 0)
	for _, model := range contents {
		cut := ShallowJobRun{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "JobRun", Function: "List"}.Wrap(nil, err)
		}
		contextModel := Context{}
		if err := contextModel.GetCtx(e); err != nil {
			return nil, err
		}
		_, err = contextModel.SetCtx(e)
		if err != nil {
			return nil, err
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowJobRun) ListBy(e echo.Context, key string, value interface{}) ([]ShallowJobRun, error) {
	content := NewShallowJobRunModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{Info: c.ShallowModel.ContentType, Package: "types", Struct: "JobRun", Function: "ListBy"}.Wrap(nil, err)
	}
	cuts := make([]ShallowJobRun, 0)
	for _, model := range contents {
		cut := NewShallowJobRun(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "JobRun", Function: "ListBy"}.Wrap(nil, err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *ShallowJobRun) Get(e echo.Context) error {
	content := NewShallowJobRunTypeContent()
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ShallowModel.ContentType = "shallowjobrun"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(nil, err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "JobRun", Function: "Get"}.Wrap(nil, err)
	}
	return nil
}

func (c *ShallowJobRun) FindBy(e echo.Context) error {
	var err error
	content := NewShallowJobRunTypeContent()
	if !c.ID.IsNil() {
		content, err = content.FindBy(e, "id", c.ID.String())
	}
	if !c.JobID.IsNil() {
		content, err = content.FindBy(e, "job_id", c.JobID.String())
	}
	if !c.WorkflowID.IsNil() {
		content, err = content.FindBy(e, "workflow_id", c.WorkflowID.String())
	}
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(nil, err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "JobRun", Function: "Get"}.Wrap(nil, err)
	}
	return nil
}

func (c *ShallowJobRun) CustomQuery(e echo.Context, write bool, q string, vars ...any) ([]ShallowJobRun, error) {
	content := NewShallowJobRunTypeContent()
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ShallowModel.ContentType = "jobrun"
	content.ID = content.ShallowModel.ID
	if write {
		content.FromType(c)
		_, err := content.CustomQuery(e, write, q, vars)
		if err != nil {
			return nil, merrors.ContentCustomQueryError{Info: c.ShallowModel.ID}.Wrap(nil, err)
		}
		return nil, nil
	}
	res, err := content.CustomQuery(e, write, q, vars)
	if err != nil {
		if err != nil {
			return nil, merrors.ContentCustomQueryError{Info: c.ShallowModel.ID}.Wrap(nil, err).BubbleCode()
		}
	}
	r := make([]ShallowJobRun, 0)
	for _, t := range res {
		jr := NewShallowJobRun(nil)
		if err = json.Unmarshal([]byte(t.Content), &jr); err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "JobRun", Function: "CustomQuery"}.Wrap(nil, err)
		}
		r = append(r, jr)
	}
	return r, nil
}

func (c ShallowJobRun) Set(e echo.Context) error {
	content := NewShallowJobRunTypeContent()
	content.FromType(c)
	content.ContentType = "shallowjobrun"
	content.ID = c.ShallowModel.ID
	err := content.Set(e)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(nil, err)
	}
	return nil
}

func (c ShallowJobRun) Delete(e echo.Context) error {
	content := NewShallowJobRunTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.ShallowModel.ID}.Wrap(nil, err)
	}
	return nil
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

func (c ShallowJobRun) IsShallowModel() bool {
	return true
}
