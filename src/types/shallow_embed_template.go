package types

import (
	"context"
	"encoding/json"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
)

type ShallowTemplate struct {
	ShallowModel
	ID 						TemplateID `form:"id" json:"id"`
	Name 					string `form:"name" json:"name"`
	DispositionsArrayModel 	[]string `form:"dispositions" json:"dispositions_array_model"`
	CurrentDisposition 		int64
	AvailableDispositions 	[]string
}

func (c ShallowTemplate) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(err)
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowTemplate) Expand(ctx context.Context) (*Template, error) {
	r := Template{}
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
	sd := ShallowDisposition{}
	r.DispositionsArrayModel = make([]Disposition, 0)
	for _, id := range c.DispositionsArrayModel {
		d := Disposition{}
		sd.ShallowModel.ID = id
		sc, err := sd.ShallowModel.Get(ctx)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err)
		}		
		if err := json.Unmarshal([]byte(sc.Content), &d); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err)
		}
		r.DispositionsArrayModel = append(r.DispositionsArrayModel, d)
	}
	r.AvailableDispositions = make([]Disposition, 0)
	for _, id := range c.AvailableDispositions {
		d := Disposition{}
		sd.ShallowModel.ID = id
		sc, err := sd.ShallowModel.Get(ctx)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err)
		}		
		if err := json.Unmarshal([]byte(sc.Content), &d); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err)
		}
		r.AvailableDispositions = append(r.AvailableDispositions, d)
	}
	return &r, nil
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

func (c ShallowTemplate) IsShallowModel() bool {
	return true
}
