package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Toggle struct {
	Model
	ID 			string 	`json:"id"`
	NamePrefix 	string 	`json:"name_prefix"`
	IdPrefix 	string 	`json:"id_suffix"`
	Suffix 		string 	`json:"suffix"`
	Value 		bool 	`json:"value"`
	Title 		string 	`json:"title"`
	Base64Value string
	JsonValue	[]byte
}

type ShallowToggle struct {
	Model
	ID 			string 	`json:"id"`
	NamePrefix 	string 	`json:"name_prefix"`
	IdPrefix 	string 	`json:"id_suffix"`
	Suffix 		string 	`json:"suffix"`
	Value 		bool 	`json:"value"`
	Title 		string 	`json:"title"`
	Base64Value string
	JsonValue	[]byte
}

func NewShallowToggle(id *string) ShallowToggle {
	c := ShallowToggle{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_steps"
	return c
}


func (c *Toggle) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Toggle) Marshal(ctx context.Context) error {
	c.JsonValue, err = json.Marshal(c)
	return err
}
