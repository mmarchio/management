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

func (c Template) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowTemplate{}
	sm.ShallowModel = sm.ShallowModel.FromEmbedModel(c.EmbedModel)
	sm.ID = c.ID
	sm.Name = c.Name
	sm.DispositionsArrayModel = make([]string, 0)
	for _, id := range c.DispositionsArrayModel {
		sm.DispositionsArrayModel = append(sm.DispositionsArrayModel, id.Model.ID)
		sms = append(sms, c.Pack()...)
	}
	sm.CurrentDisposition = c.CurrentDisposition
	sm.AvailableDispositions = make([]string, 0)
	for _, id := range c.AvailableDispositions {
		sm.AvailableDispositions = append(sm.AvailableDispositions, id.Model.ID)
		sms = append(sms, c.Pack()...)
	}
	sms = append(sms, sm)
	return sms
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
			c.DispositionsArrayModel = append(c.DispositionsArrayModel, ad)
		}
	}
	return c
}

