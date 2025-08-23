package types

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type Toggle struct {
	Model
	NamePrefix 	string `json:"name_prefix"`
	IdPrefix 	string `json:"id_suffix"`
	Suffix 		string `json:"suffix"`
	Value 		bool `json:"value"`
	Title 		string `json:"title"`
}

func (c Toggle) IsNil() bool {
	if c.Model.IsNil() && c.ID == "" && c.NamePrefix == "" && c.IdPrefix == "" && c.Suffix == "" && c.Title == "" && !c.Value {
		return true
	}
	return false
}

func (c Toggle) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowToggle{}
	sm.ShallowModel = sm.ShallowModel.FromTypeModel(c.Model)
	sm.ID = c.ID
	sm.NamePrefix = c.NamePrefix
	sm.IdPrefix = c.IdPrefix
	sm.Suffix = c.Suffix
	sm.Value = c.Value
	sm.Title = c.Title
	sms = append(sms, sm)
	return sms
}

func (c *Toggle) init() {
	c.ID = uuid.NewString()
}

func (c *Toggle) New(parent Embeddable) {
	c.NamePrefix = parent.GetContentType()
	c.IdPrefix = parent.GetContentType()
	fieldName := FindFieldName(c.ID, parent)
	if fieldName != "" {
		c.Suffix = fieldName
		c.Title = strings.Replace(fieldName, "_", "", -1)
	}
	c.ID = parent.GetID()

}

type IsStruct interface{}

func FindFieldName(needle string, haystack IsStruct) string {
	t := reflect.TypeOf(haystack)
	v := reflect.ValueOf(haystack)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		toggleValues := fmt.Sprintf("%#v", value)
		if strings.Contains(toggleValues, needle) {
			return field.Tag.Get("json")
		}
	}
	return ""
}

func (c *Toggle) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Toggle) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c Toggle) FromModel(model models.Toggle) Toggle{
	c.ID = model.ID
	c.IdPrefix = model.IdPrefix
	c.NamePrefix = model.NamePrefix
	c.Suffix = model.Suffix
	c.Title = model.Title
	c.Value = model.Value
	return c
}

func ValidateToggle(p Toggle, id, prefix, suffix, title string) Toggle {
	if p.IdPrefix == "" {
		p.IdPrefix = prefix
	}
	if p.NamePrefix == "" {
		p.NamePrefix = prefix
	}
	if p.Suffix == "" {
		p.Suffix = suffix
	}
	if p.Title == "" {
		p.Title = title
	}
	if p.ID == "" {
		p.ID = id
	}
	return p
}

func (c Toggle) Get(e echo.Context) (*Toggle, error) {
	input := Content{}
	input.Model.ID = c.Model.ID
	input.ID = c.Model.ID
	output, err := input.Get(e)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err).Log()
	}
	toggle := c
	if err := json.Unmarshal([]byte(output.Content), &toggle); err != nil {
		return nil, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
	}
	return &toggle, nil
}