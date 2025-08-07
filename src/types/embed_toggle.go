package types

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/mmarchio/management/models"
)

type Toggle struct {
	Model
	ID 			string `json:"id"`
	NamePrefix 	string `json:"name_prefix"`
	IdPrefix 	string `json:"id_suffix"`
	Suffix 		string `json:"suffix"`
	Value 		bool `json:"value"`
	Title 		string `json:"title"`
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