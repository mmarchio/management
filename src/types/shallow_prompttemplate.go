package types

import (
	"context"
	"encoding/json"
	"time"

	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type ShallowPromptTemplate struct {
	ShallowModel
	ID 			PromptTemplateID `form:"id" json:"id"`
	Name 		string `form:"name" json:"name"`
	Template 	string `form:"template" json:"template"`
	Vars 		string `form:"vars" json:"vars"`
}


func NewShallowPromptTemplate(id *string) ShallowPromptTemplate {
	c := ShallowPromptTemplate{}
	c.New(id)
	c.ShallowModel.ContentType = "shallowprompttemplate"
	return c
} 

func NewShallowPromptTemplateModelContent() models.ShallowContent {
	c := models.ShallowContent{}
	c.ShallowModel.ContentType = "shallowprompttemplate"
	return c
}

func NewShallowPromptTemplateTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "shallowprompttemplate"
	return c
}


func (c *ShallowPromptTemplate) New(id *string) {
	c.ID = c.ID.New(id)
	c.ShallowModel.ID = c.ID.String()
	c.ShallowModel.CreatedAt = time.Now()
	c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
}

func (c ShallowPromptTemplate) List(ctx context.Context) ([]ShallowPromptTemplate, error) {
	content := NewShallowPromptTemplateModelContent()
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	cuts := make([]ShallowPromptTemplate, 0)
	for _, model := range contents {
		cut := NewShallowPromptTemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowPromptTemplate", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowPromptTemplate) ListBy(ctx context.Context, key string, value interface{}) ([]ShallowPromptTemplate, error) {
	content := NewShallowPromptTemplateModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	cuts := make([]ShallowPromptTemplate, 0)
	for _, model := range contents {
		cut := NewShallowPromptTemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowPromptTemplate", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *ShallowPromptTemplate) Get(ctx context.Context) error {
	content := NewShallowPromptTemplateTypeContent()
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ShallowModel.ContentType = "shallowprompttemplate"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowPromptTemplate", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c ShallowPromptTemplate) Set(ctx context.Context) error {
	content := NewShallowPromptTemplateTypeContent()
	content.FromType(c)
	err := content.Set(ctx)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowPromptTemplate) Delete(ctx context.Context) error {
	content := NewShallowComfyUITypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.ContentDeleteError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowPromptTemplate) GetID() string {
	return c.ShallowModel.ID
}

func (c ShallowPromptTemplate) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowPromptTemplate) GetTable() string {
	return c.ShallowModel.Table
}

func (c ShallowPromptTemplate) Unmarshal(j string) (ShallowPromptTemplate, error) {
	model := models.ShallowPromptTemplate{}
	if err := json.Unmarshal([]byte(j), &model); err != nil {
		return c, merrors.JSONUnmarshallingError{Info: j, Package: "types", Struct: "ShallowPromptTemplate", Function: "Unmarshal"}.Wrap(err)
	}
	c.ShallowModel.FromModel(model.ShallowModel)

	d, err := c.SetID()
	if err != nil {
		return c, merrors.IDSetError{Info: "ShallowPromptTemplate"}.Wrap(err)
	}
	c = d
	c.Template = model.Template
	c.Vars = model.Vars
	return c, nil
} 

func (c ShallowPromptTemplate) SetID() (ShallowPromptTemplate, error) {
	var err error
	c.ID = PromptTemplateID(c.ShallowModel.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "ShallowPromptTemplate"}.Wrap(err)
	}
	return c, nil
}