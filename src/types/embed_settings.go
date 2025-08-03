package types

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Settings struct {
	EmbedModel
	ID 				string 		`json:"id"`
	Name 			string 		`json:"name"`
	Template 		Template 	`json:"template_id"`
	GlobalBypass 	Steps 		`json:"global_bypass"`
	Recurring 		Toggle 		`json:"recurring"`
	Interval 		int64 		`json:"interval"`
	Workflow 		WorkflowID 	`json:"workflow_id"`
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
	c.GlobalBypass.EmbedModel.New("global_bypass")
	c.Template.EmbedModel.New("template")
}

func (c Settings) Bind(e echo.Context) Settings {
	c.GlobalBypass, _ = c.GlobalBypass.Bind(e)
	c.Template = c.Template.Bind(e)
	if e.FormValue("template_name") != "" {
		c.Name = e.FormValue("template_name")
	}
	if e.FormValue("recurring") == "on" {
		c.Recurring.Value = true
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
