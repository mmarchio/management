package types

import (
	"context"
	"encoding/json"
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

