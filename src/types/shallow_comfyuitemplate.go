package types

import (
	"context"
	"encoding/json"
	"time"

	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

func NewShallowComfyUITemplate(id *string) ShallowComfyUITemplate {
	c := ShallowComfyUITemplate{}
	ct := "shallowcomfyuitemplate"
	c.ShallowModel.New(id, &ct)
	c.ShallowModel.ContentType = "shallowcomfyuitemplate"
	return c
} 

func NewShallowComfyUIModelContent() models.ShallowContent {
	c := models.ShallowContent{}
	c.ShallowModel.ContentType = "shallowcomfyuitemplate"
	return c
}

func NewShallowComfyUITypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "shallowcomfyuitemplate"
	return c
}

type ShallowComfyUITemplate struct {
	ShallowModel
	ID 			ComfyUITemplateID 	`form:"id" json:"id"`
	Name 		string 			`form:"name" json:"name"`
	Endpoint 	string 		`form:"enpoint" json:"endpoint"`
	Base 		string 			`form:"base"json:"base"`
	Template 	string 		`form:"template" json:"template"`
}

func (c *ShallowComfyUITemplate) New() {
	c.ID = c.ID.New()
	c.ShallowModel.ID = c.ID.String()
	c.ShallowModel.CreatedAt = time.Now()
	c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
}

func (c ShallowComfyUITemplate) List(ctx context.Context) ([]ShallowComfyUITemplate, error) {
	content := NewShallowComfyUIModelContent()
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	cuts := make([]ShallowComfyUITemplate, 0)
	for _, model := range contents {
		cut := NewShallowComfyUITemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowComfyUITemplate", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowComfyUITemplate) ListBy(ctx context.Context, key string, value interface{}) ([]ShallowComfyUITemplate, error) {
	content := NewShallowComfyUIModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	cuts := make([]ShallowComfyUITemplate, 0)
	for _, model := range contents {
		cut := NewShallowComfyUITemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowComfyUITemplate", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *ShallowComfyUITemplate) Get(ctx context.Context) error {
	content := NewShallowComfyUITypeContent()
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ShallowModel.ContentType = "shallowcomfyuitemplate"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowComfyUITemplate", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c ShallowComfyUITemplate) Set(ctx context.Context) error {
	content := NewShallowComfyUITypeContent()
	content.FromType(c)
	err := content.Set(ctx)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowComfyUITemplate) Delete(ctx context.Context) error {
	content := NewShallowComfyUITypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.ContentDeleteError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowComfyUITemplate) GetID() string {
	return c.ShallowModel.ID
}

func (c ShallowComfyUITemplate) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowComfyUITemplate) GetTable() string {
	return c.ShallowModel.Table
}

func (c ShallowComfyUITemplate) Unmarshal(j string) (ShallowComfyUITemplate, error) {
	model := models.ShallowComfyUITemplate{}
	if err := json.Unmarshal([]byte(j), &model); err != nil {
		return c, merrors.JSONUnmarshallingError{Info: j, Package: "types", Struct: "ShallowComfyUITemplate", Function: "Unmarshal"}.Wrap(err)
	}
	c.ShallowModel.FromModel(model.ShallowModel)

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

func (c ShallowComfyUITemplate) SetID() (ShallowComfyUITemplate, error) {
	var err error
	c.ID = ComfyUITemplateID(c.ShallowModel.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "disposition"}.Wrap(err)
	}
	return c, nil
}