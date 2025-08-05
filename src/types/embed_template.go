package types

import (
	"context"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type Template struct {
	EmbedModel
	ID 						TemplateID `form:"id" json:"id"`
	Name 					string `form:"name" json:"name"`
	DispositionsArrayModel 	[]Disposition `form:"dispositions" json:"dispositions_array_model"`
	CurrentDisposition 		int64
	AvailableDispositions 	[]Disposition
}

func (c *Template) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Template) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c Template) Bind(e echo.Context) Template {
	if e.FormValue("template_name") != "" {
		c.Name = e.FormValue(("template_name"))
	}
	for _, ad := range c.AvailableDispositions {
		if e.FormValue(ad.Model.ID) == "on" {
			c.Dispositions = append(c.Dispositions, ad)
		}
	}
	return c
}

