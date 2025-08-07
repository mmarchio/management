package types

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/mmarchio/management/models"
)

type ShallowToggle struct {
	ShallowModel
	ID 			string `json:"id"`
	NamePrefix 	string `json:"name_prefix"`
	IdPrefix 	string `json:"id_suffix"`
	Suffix 		string `json:"suffix"`
	Value 		bool `json:"value"`
	Title 		string `json:"title"`
}

func (c *ShallowToggle) init() {
	c.ID = uuid.NewString()
}

func (c *ShallowToggle) New(parent Embeddable) {
	c.NamePrefix = parent.GetContentType()
	c.IdPrefix = parent.GetContentType()
	fieldName := FindFieldName(c.ID, parent)
	if fieldName != "" {
		c.Suffix = fieldName
		c.Title = strings.Replace(fieldName, "_", "", -1)
	}
	c.ID = parent.GetID()

}

func (c *ShallowToggle) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowToggle) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowToggle) FromModel(model models.ShallowToggle) ShallowToggle{
	c.ID = model.ID
	c.IdPrefix = model.IdPrefix
	c.NamePrefix = model.NamePrefix
	c.Suffix = model.Suffix
	c.Title = model.Title
	c.Value = model.Value
	return c
}
