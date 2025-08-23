package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type PromptTemplate struct {
	Model
	Name     string `form:"name" json:"name"`
	Template string `form:"template" json:"template"`
	Vars     string `form:"vars" json:"vars"`
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
	if id != nil {
		c.Model.ID = *id		
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
}

func (c PromptTemplate) List(e echo.Context) ([]PromptTemplate, error) {
	content := NewPromptTemplateModelContent()
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]PromptTemplate, 0)
	for _, model := range contents {
		cut := NewPromptTemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "PromptTemplate", Function: "List"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c PromptTemplate) ListBy(e echo.Context, key string, value interface{}) ([]PromptTemplate, error) {
	content := NewPromptTemplateModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]PromptTemplate, 0)
	for _, model := range contents {
		cut := NewPromptTemplate(nil)
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "PromptTemplate", Function: "ListBy"}.Wrap(err).Log()
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
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "PromptTemplate", Function: "Get"}.Wrap(err).Log()
	}
	return nil
}

func (c PromptTemplate) Set(e echo.Context, update bool) error {
	content := NewPromptTemplateTypeContent()
	content.FromType(c, c.Model)
	err := content.Set(e, update)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	return nil
}

func (c PromptTemplate) Delete(e echo.Context) error {
	content := NewComfyUITypeContent()
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(err).Log()
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
		return c, merrors.JSONUnmarshallingError{Info: j}.Wrap(err).Log().Log()
	}
	c.Model.FromModel(model.Model)

	c.Template = model.Template
	c.Vars = model.Vars
	return c, nil
}

func (c *PromptTemplate) Prepare(e echo.Context, jobrun *JobRun) error {
	vars := make(map[string]interface{})
	if err := json.Unmarshal([]byte(c.Vars), &vars); err != nil {
		return merrors.JSONUnmarshallingError{Info: c.Vars}.Wrap(err).Log().Log()
	}
	ctxb, err := json.Marshal(jobrun.ContextModel)
	if err != nil {
		return merrors.JSONMarshallingError{}.Wrap(err).Log().Log()
	}
	for k, v := range vars {
		if cacheVal := jobrun.GetValueCache(v); cacheVal != nil {
			vars[k] = cacheVal
			continue
		}

		if val, ok := v.(string); ok {
			if strings.Contains(val, "context.") {
				val = strings.Replace(val, "context.", "$.", 1)
				vars[k], err = jpath(string(ctxb), val)
				if err != nil {
					return merrors.JSONUnmarshallingError{}.Wrap(err).Log().Log()
				}
			}
			if strings.Contains(val, "steps[") {
				step := NewStep(nil)
				q := fmt.Sprintf("SELECT %s FROM content WHERE content_type = 'step' AND content @> '{\"workflow_id\":\"%s\"}'", step.Model.Columns, jobrun.WorkflowID.String())
				contents, err := Content{}.CustomQuery(e, false, q)
				if err != nil {
					return merrors.ContentCustomQueryError{Info: q}.Wrap(err).Log()
				}
				for _, content := range contents {
					s := Step{}
					if err := json.Unmarshal([]byte(content.Content), &s); err != nil {
						return merrors.JSONUnmarshallingError{}.Wrap(err).Log()
					}
					stepVal, err := c.ParseStep(val, s)
					if err != nil {
						return err
					}

					if stepVal != nil {
						if stepString, ok := stepVal.(string); ok {
							vars[k] = stepString
						}
						if stepTime, ok := stepVal.(time.Time); ok {
							vars[k] = stepTime.Format(time.RFC3339)
						}
						if stepDur, ok := stepVal.(time.Duration); ok {
							vars[k] = stepDur
						}
					}
				}
			}
		}
	}
	b, err := json.Marshal(vars)
	if err != nil {
		return merrors.JSONMarshallingError{}.Wrap(err).Log().Log()
	}
	c.Vars = string(b)
	return nil
}

func (c PromptTemplate) ParseStep(s string, step Step) (interface{}, error) {
	parts := strings.Split(s, ".")
	index := strings.Replace(parts[0], "steps[", "", 1)
	index = strings.Replace(index, "]", "", 1)
	indexInt, err := strconv.Atoi(index)
	if err != nil {
		return nil, merrors.ParseContent{}.Wrap(err).Log()
	}
	if indexInt == step.Order-1 {
		if len(parts) > 1 {
			GetLogger(3).Flogger("parts %#v", parts)
			switch strings.ToLower(parts[1]) {
			case "output":
				if len(parts) > 2 {
					GetLogger(3).Flogger("pgr %#v", step.Stats.Output)
					switch parts[2] {
					case "think":
						return step.Stats.Output.Think, nil
					case "topics":
						return step.Stats.Output.Topics, nil
					case "prompt":
						return step.Stats.Output.Prompt, nil
					default:
						return nil, nil
					}
				}
				return step.Stats.Output, nil
			case "input":
				return step.Stats.Input, nil
			case "start":
				return step.Stats.Start, nil
			case "end":
				return step.Stats.End, nil
			case "duration":
				return step.Stats.Duration, nil
			default:
				return nil, merrors.ParseContent{}.New("unknown field %s", parts[1])
			}
		}
	}
	return nil, nil
}
