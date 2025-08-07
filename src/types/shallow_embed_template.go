package types

import (
	"context"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ShallowTemplate struct {
	ShallowModel
	ID 						TemplateID `form:"id" json:"id"`
	Name 					string `form:"name" json:"name"`
	DispositionsArrayModel 	[]string `form:"dispositions" json:"dispositions_array_model"`
	CurrentDisposition 		int64
	AvailableDispositions 	[]string
}

func (c *ShallowTemplate) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowTemplate) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowTemplate) Bind(e echo.Context) ShallowTemplate {
	if e.FormValue("template_name") != "" {
		c.Name = e.FormValue(("template_name"))
	}
	for _, ad := range c.AvailableDispositions {
		if e.FormValue(ad) == "on" {
			c.DispositionsArrayModel = append(c.DispositionsArrayModel, ad)
		}
	}
	return c
}

