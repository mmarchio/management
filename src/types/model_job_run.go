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

func NewJobRun(id *string) JobRun {
	c := JobRun{}
	c.New(id)
	c.Model.ContentType = "jobrun"
	c.Model.Columns = "id, created_at, updated_at, content_type, content"
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
	JobID                 JobID             					`json:"job_id"`
	WorkflowID            WorkflowID        					`json:"workflow_id"`
	ContextModel          Context           					`json:"context_model"`
	TruncatedContextModel TruncatedContext  					`json:"truncated_context_model"`
	SettingsModel         Settings          					`json:"settings_model"`
	DispositionModel      Disposition       					`json:"disposition_model"`
	NodeStepMap           map[string]string 					`json:"node_step_map"`
	Tokens                int64             					`json:"tokens"`
	LatestStatusType      string            					`json:"latest_status_type"`
	LatestStatusValue     string            					`json:"latest_status_value"`
	ValueCache            map[string]interface{}
}

func (c JobRun) GetValueCache(k interface{}) interface{} {
	if key, ok := k.(string); ok {
		if c.ValueCache[key] != nil {
			return c.ValueCache[key]
		}
	}
	return nil
}

func (c JobRun) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowJobRun{}
	sm.ShallowModel = sm.ShallowModel.FromTypeModel(c.Model)
	sm.ID = c.ID
	sm.JobID = c.JobID
	sm.WorkflowID = c.WorkflowID
	sm.SettingsModel = c.SettingsModel.ID
	sms = append(sms, c.SettingsModel.Pack()...)
	sm.DispositionModel = c.DispositionModel.Model.ID
	sms = append(sms, c.DispositionModel.Pack()...)
	sm.TokenCount = c.Tokens
	sm.LatestStatusType = c.LatestStatusType
	sm.LatestStatusValue = c.LatestStatusValue
	sms = append(sms, sm)
	return sms
}

func (c *JobRun) New(id *string) {
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
}

func (c JobRun) List(e echo.Context) ([]JobRun, error) {
	content := NewJobRunModelContent()
	content.Model.ContentType = "jobrun"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType, Package: "types", Struct: "JobRun", Function: "List"}.Wrap(err).Log()
	}
	cuts := make([]JobRun, 0)
	for _, model := range contents {
		cut := JobRun{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "JobRun", Function: "List"}.Wrap(err).Log()
		}
		cut.ContextModel.SetCtx(e)
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c JobRun) ListBy(e echo.Context, key string, value interface{}) ([]JobRun, error) {
	content := NewJobRunModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType, Package: "types", Struct: "JobRun", Function: "ListBy"}.Wrap(err).Log()
	}
	cuts := make([]JobRun, 0)
	for _, model := range contents {
		cut := NewJobRun(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "JobRun", Function: "ListBy"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *JobRun) Get(e echo.Context) error {
	content := NewJobRunTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "jobrun"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "JobRun", Function: "Get"}.Wrap(err).Log()
	}
	return nil
}

func (c *JobRun) FindBy(e echo.Context) error {
	var err error
	content := NewJobRunTypeContent()
	if c.Model.ID != "" {
		content, err = content.FindBy(e, "id", c.Model.ID)
	}
	if !c.JobID.IsNil() {
		content, err = content.FindBy(e, "job_id", c.JobID.String())
	}
	if !c.WorkflowID.IsNil() {
		content, err = content.FindBy(e, "workflow_id", c.WorkflowID.String())
	}
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "JobRun", Function: "Get"}.Wrap(err).Log()
	}
	return nil
}

func (c *JobRun) CustomQuery(e echo.Context, write bool, q string, vars ...any) ([]JobRun, error) {
	content := NewJobRunTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "jobrun"
	content.ID = content.Model.ID
	if write {
		content.FromType(c, c.Model)
		_, err := content.CustomQuery(e, write, q, vars)
		if err != nil {
			return nil, merrors.ContentCustomQueryError{Info: c.Model.ID}.Wrap(err).Log()
		}
		return nil, nil
	}
	res, err := content.CustomQuery(e, write, q, vars)
	if err != nil {
		if err != nil {
			return nil, merrors.ContentCustomQueryError{Info: c.Model.ID}.Wrap(err).Log().BubbleCode()
		}
	}
	r := make([]JobRun, 0)
	for _, t := range res {
		jr := NewJobRun(nil)
		if err = json.Unmarshal([]byte(t.Content), &jr); err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "JobRun", Function: "CustomQuery"}.Wrap(err).Log()
		}
		r = append(r, jr)
	}
	return r, nil
}

func (c JobRun) Set(e echo.Context, update bool) error {
	content := NewJobRunTypeContent()
	content.FromType(c, c.Model)
	content.ContentType = "jobrun"
	content.ID = c.Model.ID
	err := content.Set(e, update)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	return nil
}

func (c JobRun) Delete(e echo.Context) error {
	content := NewJobRunTypeContent()
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(err).Log()
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

func (c JobRun) GetNodes(e echo.Context) (map[string]string, error) {
	GetLogger(4).Flogger("GetNodes called")
	d := NewJobRun(nil)
	q := fmt.Sprintf("SELECT %s FROM content WHERE (content_type = 'comfynode' OR content_type = 'ollamanode' OR content_type = 'sshnode') AND content @> '{\"workflow_id\":\"%s\"}'", d.Model.Columns, c.ContextModel.SettingsModel.Workflow)
	content := Content{}
	nodes, err := content.UnrestrictedCustomQuery(e, false, q)
	if err != nil {
		return nil, err
	}
	r := make(map[string]string)
	for _, node := range nodes {
		switch node.ContentType {
		case "comfynode":
			n := ComfyNode{}
			if err := json.Unmarshal([]byte(node.Content), &n); err != nil {
				return nil, err
			}
			r[n.ID] = n.Name
		case "ollamanode":
			n := OllamaNode{}
			if err := json.Unmarshal([]byte(node.Content), &n); err != nil {
				return nil, err
			}
			r[n.ID] = n.Name
		case "sshnode":
			n := SSHNode{}
			if err := json.Unmarshal([]byte(node.Content), &n); err != nil {
				return nil, err
			}
			r[n.ID] = n.Name
		default:
		}
	}
	return r, nil
}

func (c JobRun) GetSteps(e echo.Context) (map[string]string, error) {
	GetLogger(4).Flogger("GetSteps called")
	d := NewJobRun(nil)
	q := fmt.Sprintf("SELECT %s FROM content WHERE content_type = 'step' AND content @> '{\"workflow_id\":\"%s\"}'", d.Model.Columns, c.ContextModel.SettingsModel.Workflow)
	content := Content{}
	steps, err := content.UnrestrictedCustomQuery(e, false, q)
	if err != nil {
		return nil, err
	}
	r := make(map[string]string)
	for _, step := range steps {
		s := Step{}
		if err := json.Unmarshal([]byte(step.Content), &s); err != nil {
			return nil, err
		}
		r[s.Model.ID] = s.Name
	}
	return r, nil
}
