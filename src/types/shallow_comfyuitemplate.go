package types

import (
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
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

func (c ShallowComfyUITemplate) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(nil, err)
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowComfyUITemplate) Expand(e echo.Context) (*ComfyUITemplate, error) {
	r := ComfyUITemplate{}
	if c.ShallowModel.CreatedAt.IsZero() && c.ShallowModel.ID != "" {
		sc, err := c.ShallowModel.Get(e)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(nil, err)
		}
		if err := json.Unmarshal([]byte(sc.Content), &r); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(nil, err)
		}
		return &r, nil
	}
	r.Model = r.Model.FromShallowModel(c.ShallowModel)
	r.ID = c.ID
	r.Name = c.Name
	r.Endpoint = c.Endpoint
	r.Base = c.Base
	r.Template = c.Template
	return &r, nil
}

func (c *ShallowComfyUITemplate) New() {
	c.ID = c.ID.New()
	c.ShallowModel.ID = c.ID.String()
	c.ShallowModel.CreatedAt = time.Now()
	c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
}

func (c ShallowComfyUITemplate) List(e echo.Context) ([]ShallowComfyUITemplate, error) {
	content := NewShallowComfyUIModelContent()
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(nil, err)
	}
	cuts := make([]ShallowComfyUITemplate, 0)
	for _, model := range contents {
		cut := NewShallowComfyUITemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowComfyUITemplate", Function: "List"}.Wrap(nil, err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowComfyUITemplate) ListBy(e echo.Context, key string, value interface{}) ([]ShallowComfyUITemplate, error) {
	content := NewShallowComfyUIModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(nil, err)
	}
	cuts := make([]ShallowComfyUITemplate, 0)
	for _, model := range contents {
		cut := NewShallowComfyUITemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowComfyUITemplate", Function: "ListBy"}.Wrap(nil, err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *ShallowComfyUITemplate) Get(e echo.Context) error {
	content := NewShallowComfyUITypeContent()
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ShallowModel.ContentType = "shallowcomfyuitemplate"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(nil, err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowComfyUITemplate", Function: "Get"}.Wrap(nil, err)
	}
	return nil
}

func (c ShallowComfyUITemplate) Set(e echo.Context) error {
	content := NewShallowComfyUITypeContent()
	content.FromType(c)
	err := content.Set(e)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(nil, err)
	}
	return nil
}

func (c ShallowComfyUITemplate) Delete(e echo.Context) error {
	content := NewShallowComfyUITypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.ShallowModel.ID}.Wrap(nil, err)
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
		return c, merrors.JSONUnmarshallingError{Info: j, Package: "types", Struct: "ShallowComfyUITemplate", Function: "Unmarshal"}.Wrap(nil, err)
	}
	c.ShallowModel.FromModel(model.ShallowModel)

	d, err := c.SetID()
	if err != nil {
		return c, merrors.IDSetError{Info: "disposition"}.Wrap(nil, err)
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
		return c, merrors.IDSetError{Info: "disposition"}.Wrap(nil, err)
	}
	return c, nil
}