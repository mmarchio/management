package types

import (
	"context"
	"encoding/json"

	merrors "github.com/mmarchio/management/errors"
)

type ShallowSettings struct {
	ShallowModel
	ID 					string 		`json:"id"`
	Name 				string 		`json:"name"`
	TemplateModel 		string 		`json:"template_model"`
	GlobalBypassModel 	string 		`json:"global_bypass_model"`
	RecurringModel 		string 		`json:"recurring_model"`
	Interval 			int64 		`json:"interval"`
	Workflow 			WorkflowID 	`json:"workflow_id"`
}

func (c ShallowSettings) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(err)
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowSettings) Expand(ctx context.Context) (*Settings, error) {
	r := Settings{}
	if c.ShallowModel.CreatedAt.IsZero() && c.ShallowModel.ID != "" {
		sc, err := c.ShallowModel.Get(ctx)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err)
		}
		if err := json.Unmarshal([]byte(sc.Content), &r); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err)
		}
		return &r, nil
	}
	r.EmbedModel = r.EmbedModel.FromShallowModel(c.ShallowModel)
	r.ID = c.ID
	r.Name = c.Name
	st := ShallowTemplate{}
	st.ShallowModel.ID = c.TemplateModel
	template, err := st.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.TemplateModel = *template
	sgp := ShallowSteps{}
	sgp.ShallowModel.ID = c.GlobalBypassModel
	globalbypass, err := sgp.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.GlobalBypassModel = *globalbypass
	stg := ShallowToggle{}
	stg.ShallowModel.ID = c.RecurringModel
	recurring, err := stg.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.RecurringModel = *recurring
	r.Interval = c.Interval
	r.Workflow = c.Workflow
	return &r, nil
}

func (c ShallowSettings) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ShallowSettings) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c *ShallowSettings) New(globalBypass, template *string) {
	ct := "shallowsettings"
	c.ShallowModel.New(nil, &ct)
	if globalBypass != nil {
		c.GlobalBypassModel = *globalBypass
	}
	if template != nil {
		c.TemplateModel = *template
	}
}

func (c ShallowSettings) IsShallowModel() bool {
	return true
}
