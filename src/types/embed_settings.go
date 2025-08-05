package types

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Settings struct {
	EmbedModel
	ID 					string 		`json:"id"`
	Name 				string 		`json:"name"`
	TemplateModel 		Template 	`json:"template_model"`
	GlobalBypassModel 	Steps 		`json:"global_bypass_model"`
	RecurringModel 		Toggle 		`json:"recurring_model"`
	Interval 			int64 		`json:"interval"`
	Workflow 			WorkflowID 	`json:"workflow_id"`
}

func (c Settings) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Settings) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c *Settings) New() {
	c.EmbedModel.New("settings")
	c.GlobalBypassModel.EmbedModel.New("global_bypass")
	c.TemplateModel.EmbedModel.New("template")
}

func (c Settings) Bind(e echo.Context) Settings {
	c.GlobalBypassModel, _ = c.GlobalBypassModel.Bind(e)
	c.TemplateModel = c.TemplateModel.Bind(e)
	if e.FormValue("template_name") != "" {
		c.Name = e.FormValue("template_name")
	}
	if e.FormValue("recurring") == "on" {
		c.RecurringModel.Value = true
	}
	if e.FormValue("interval") != "" {
		i, _ := strconv.Atoi(e.FormValue("interval"))
		c.Interval = int64(i)
	}
	if e.FormValue("workflow") != "" {
		c.Workflow = WorkflowID(e.FormValue("workflow"))
	}
	return c
}
