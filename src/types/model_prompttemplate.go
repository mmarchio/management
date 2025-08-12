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
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]PromptTemplate, 0)
	for _, model := range contents {
		cut := NewPromptTemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "PromptTemplate", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c PromptTemplate) ListBy(e echo.Context, key string, value interface{}) ([]PromptTemplate, error) {
	content := NewPromptTemplateModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]PromptTemplate, 0)
	for _, model := range contents {
		cut := NewPromptTemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "PromptTemplate", Function: "ListBy"}.Wrap(err)
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
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "PromptTemplate", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c PromptTemplate) Set(e echo.Context) error {
	content := NewPromptTemplateTypeContent()
	content.FromType(c)
	err := content.Set(e)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c PromptTemplate) Delete(e echo.Context) error {
	content := NewComfyUITypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(err)
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
		return c, merrors.JSONUnmarshallingError{Info: j, Package: "types", Struct: "PromptTemplate", Function: "Unmarshal"}.Wrap(err)
	}
	c.Model.FromModel(model.Model)

	d, err := c.SetID()
	if err != nil {
		return c, merrors.IDSetError{Info: "PromptTemplate"}.Wrap(err)
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
		return c, merrors.IDSetError{Info: "PromptTemplate"}.Wrap(err)
	}
	return c, nil
}

func (c *PromptTemplate) Prepare(e echo.Context, jobrun *JobRun) error {
	vars := make(map[string]interface{})
	if err := json.Unmarshal([]byte(c.Vars), &vars); err != nil {
		GetLogger().Flogger("error unmarshalling vars: %s", err.Error())
		return merrors.JSONUnmarshallingError{Info: c.Vars}.Wrap(err)
	}
	ctxb, err := json.Marshal(jobrun.ContextModel)
	if err != nil {
		GetLogger().Flogger("error marshalling context: %s", err.Error())
		return merrors.JSONMarshallingError{}.Wrap(err)
	}
	for k, v := range vars {
		if val, ok := v.(string); ok {
			if strings.Contains(val, "context.") {
				val = strings.Replace(val, "context.", "$.", 1)
				vars[k], err = jpath(string(ctxb), val)
				if err != nil {
					GetLogger().Flogger("error navigating with jsonpath err: %s", err.Error())
					return merrors.JSONUnmarshallingError{}.Wrap(err)
				}
			}
		}
	}
	b, err := json.Marshal(vars)
	if err != nil {
		GetLogger().Flogger("error marshalling converted vars err: %s", err.Error())
		return merrors.JSONMarshallingError{}.Wrap(err)
	}
	c.Vars = string(b)
	return nil
}
