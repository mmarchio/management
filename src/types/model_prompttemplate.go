package types

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type PromptTemplate struct {
	Model
	ID 			PromptTemplateID `form:"id" json:"id"`
	Name 		string `form:"name" json:"name"`
	Template 	string `form:"template" json:"template"`
	Vars 		string `form:"vars" json:"vars"`
}

func (c PromptTemplate) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowPromptTemplate{}
	sm.ShallowModel = sm.ShallowModel.FromTypeModel(c.Model)
	sm.ID = c.ID
	sm.Name = c.Name
	sm.Template = c.Template
	sm.Vars = c.Vars
	sms = append(sms, sm)
	return sms
}

func NewPromptTemplate(id *string) PromptTemplate {
	c := PromptTemplate{}
	c.New(id)
	c.Model.ContentType = "prompttemplate"
	return c
} 

func NewPromptTemplateModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "prompttemplate"
	return c
}

func NewPromptTemplateTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "prompttemplate"
	return c
}


func (c *PromptTemplate) New(id *string) {
	c.ID = c.ID.New(id)
	c.Model.ID = c.ID.String()
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
}

func (c PromptTemplate) List(e echo.Context) ([]PromptTemplate, error) {
	content := NewPromptTemplateModelContent()
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(nil, err)
	}
	cuts := make([]PromptTemplate, 0)
	for _, model := range contents {
		cut := NewPromptTemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "PromptTemplate", Function: "List"}.Wrap(nil, err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c PromptTemplate) ListBy(e echo.Context, key string, value interface{}) ([]PromptTemplate, error) {
	content := NewPromptTemplateModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(nil, err)
	}
	cuts := make([]PromptTemplate, 0)
	for _, model := range contents {
		cut := NewPromptTemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "PromptTemplate", Function: "ListBy"}.Wrap(nil, err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *PromptTemplate) Get(e echo.Context) error {
	content := NewPromptTemplateTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "prompttemplate"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(nil, err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "PromptTemplate", Function: "Get"}.Wrap(nil, err)
	}
	return nil
}

func (c PromptTemplate) Set(e echo.Context) error {
	content := NewPromptTemplateTypeContent()
	content.FromType(c)
	err := content.Set(e)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(nil, err)
	}
	return nil
}

func (c PromptTemplate) Delete(e echo.Context) error {
	content := NewComfyUITypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(nil, err)
	}
	return nil
}

func (c PromptTemplate) GetID() string {
	return c.Model.ID
}

func (c PromptTemplate) GetContentType() string {
	return c.Model.ContentType
}

func (c PromptTemplate) GetTable() string {
	return c.Model.Table
}

func (c PromptTemplate) Unmarshal(j string) (PromptTemplate, error) {
	model := models.PromptTemplate{}
	if err := json.Unmarshal([]byte(j), &model); err != nil {
		return c, merrors.JSONUnmarshallingError{Info: j, Package: "types", Struct: "PromptTemplate", Function: "Unmarshal"}.Wrap(nil, err)
	}
	c.Model.FromModel(model.Model)

	d, err := c.SetID()
	if err != nil {
		return c, merrors.IDSetError{Info: "PromptTemplate"}.Wrap(nil, err)
	}
	c = d
	c.Template = model.Template
	c.Vars = model.Vars
	return c, nil
} 

func (c PromptTemplate) SetID() (PromptTemplate, error) {
	var err error
	c.ID = PromptTemplateID(c.Model.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "PromptTemplate"}.Wrap(nil, err)
	}
	return c, nil
}

func (c *PromptTemplate) Prepare(e echo.Context) error {
	var ctx Context
	ctxPtr, err := c.Model.GetCtx(e)
	if err != nil {
		GetLogger().Flogger("error retrieving context err: %s", err.Error())
		return merrors.ContextGetError{}.Wrap(nil, err)
	}
	if ctxPtr == nil {
		return merrors.ContextGetError{}.New(nil, "context is nil")
	}
	ctx = *ctxPtr
	vars := make(map[string]interface{})
	if err := json.Unmarshal([]byte(c.Vars), &vars); err != nil {
		GetLogger().Flogger("error unmarshalling vars: %s", err.Error())
		return merrors.JSONUnmarshallingError{}.Wrap(nil, err)
	}
	ctxb, err := json.Marshal(ctx)
	if err != nil {
		GetLogger().Flogger("error marshalling context: %s", err.Error())
		return merrors.JSONMarshallingError{}.Wrap(nil, err)
	}
	for k, v := range vars {
		if val, ok := v.(string); ok {
			if strings.Contains(val, "context.") {
				val = strings.Replace(val, "context.", "$.", 1)
				vars[k], err = jpath(string(ctxb), val)
				if err != nil {
					GetLogger().Flogger("error navigating with jsonpath err: %s", err.Error())
					return merrors.JSONUnmarshallingError{}.Wrap(nil, err)
				}
			}
		}
	}
	b, err := json.Marshal(vars)
	if err != nil {
		GetLogger().Flogger("error marshalling converted vars err: %s", err.Error())
		return merrors.JSONMarshallingError{}.Wrap(nil, err)
	}
	c.Vars = string(b)
	return nil
}

func jpath(j, p string) (interface{}, error) {
	parts := strings.Split(p, ".")
	msi := make(map[string]interface{})
	if err := json.Unmarshal([]byte(j), &msi); err != nil {
		GetLogger().Flogger("jpath unmarshalling error: %s", err.Error())
		return "", merrors.JSONUnmarshallingError{}.Wrap(nil, err)
	}
	return jpathRecurse(parts, msi)
}

func jpathRecurse(parts []string, msi map[string]interface{}) (interface{}, error) {
	GetLogger().Flogger("parts: %#v", parts)
	if len(parts) == 1 {
		if sub, ok := msi[parts[0]].(map[string]interface{}); !ok {
			return sub, nil
		}
	}
	for _, p := range parts {
		if p == "$" {
			return jpathRecurse(parts[1:], msi)
		}
		if sub, ok := msi[p].(map[string]interface{}); ok {
			return jpathRecurse(parts[1:], sub)
		}
		if sub, ok := msi[p].([]interface{}); ok {
			for _, sl := range sub {
				if s, ok := sl.(map[string]interface{}); ok {
					return jpathRecurse(parts[1:], s)
				}
			}
			//TODO: supprt slice lookups
		}
	}
	return nil, merrors.JPATHError{}.New(nil, "resource not found")
}