package types

import (
	"context"
	"encoding/json"
	"time"

	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

func NewComfyUITemplate(id *string) ComfyUITemplate {
	c := ComfyUITemplate{}
	c.Model.New(id)
	c.Model.ContentType = "comfyuitemplate"
	return c
} 

func NewComfyUIModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "comfyuitemplate"
	return c
}

func NewComfyUITypeContent() Content {
	c := Content{}
	c.Model.ContentType = "comfyuitemplate"
	return c
}

type ComfyUITemplate struct {
	Model
	ID 			ComfyUITemplateID 	`form:"id" json:"id"`
	Name 		string 			`form:"name" json:"name"`
	Endpoint 	string 		`form:"enpoint" json:"endpoint"`
	Base 		string 			`form:"base"json:"base"`
	Template 	string 		`form:"template" json:"template"`
}

func (c *ComfyUITemplate) New() {
	c.ID = c.ID.New()
	c.Model.ID = c.ID.String()
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
}

func (c ComfyUITemplate) List(ctx context.Context) ([]ComfyUITemplate, error) {
	content := NewComfyUIModelContent()
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.ComfyUITemplateListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]ComfyUITemplate, 0)
	for _, model := range contents {
		cut := NewComfyUITemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ComfyUITemplate", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ComfyUITemplate) ListBy(ctx context.Context, key string, value interface{}) ([]ComfyUITemplate, error) {
	content := NewComfyUIModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ComfyUITemplateListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]ComfyUITemplate, 0)
	for _, model := range contents {
		cut := NewComfyUITemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ComfyUITemplate", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *ComfyUITemplate) Get(ctx context.Context) error {
	content := NewComfyUITypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "comfyuitemplate"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.ComfyUITemplateGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ComfyUITemplate", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c ComfyUITemplate) Set(ctx context.Context) error {
	content := NewComfyUITypeContent()
	content.FromType(c)
	err := content.Set(ctx)
	if err != nil {
		return merrors.ComfyUITemplateSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c ComfyUITemplate) Delete(ctx context.Context) error {
	content := NewComfyUITypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.ComfyUITemplateDeleteError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c ComfyUITemplate) GetID() string {
	return c.Model.ID
}

func (c ComfyUITemplate) GetContentType() string {
	return c.Model.ContentType
}

func (c ComfyUITemplate) GetTable() string {
	return c.Model.Table
}

func (c ComfyUITemplate) Unmarshal(j string) (ComfyUITemplate, error) {
	model := models.ComfyUITemplate{}
	if err := json.Unmarshal([]byte(j), &model); err != nil {
		return c, merrors.JSONUnmarshallingError{Info: j, Package: "types", Struct: "ComfyUITemplate", Function: "Unmarshal"}.Wrap(err)
	}
	c.Model.FromModel(model.Model)

	d, err := c.SetID()
	if err != nil {
		return c, merrors.IDSetError{Info: "disposition"}.Wrap(err)
	}
	c = d
	c.Name = model.Name
	c.Endpoint = model.Endpoint
	c.Base = model.Base
	c.Template = model.Template
	return c, nil
} 

func (c ComfyUITemplate) SetID() (ComfyUITemplate, error) {
	var err error
	c.ID = ComfyUITemplateID(c.Model.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "disposition"}.Wrap(err)
	}
	return c, nil
}