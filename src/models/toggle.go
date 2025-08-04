package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
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

func (c Toggle) Pack(ctx context.Context) (string, error) {
	err = c.Marshal(ctx)
	if err != nil {
		return "", merrors.JSONMarshallingError{}.Wrap(err)
	}
	c.Base64Value, err = ToBase64(c.JsonValue)
	if err != nil {
		return  "", merrors.Base64EncodingError{Info: string(c.JsonValue)}.Wrap(err)
	}
	return c.Base64Value, nil
}

func (c *Toggle) Unpack(ctx context.Context) error {
	j, err := FromBase64(c.Base64Value)
	if err != nil {
		return merrors.Base64DecodingError{Info: string(c.Base64Value)}.Wrap(err)
	}
	err = c.Unmarshal(ctx, j)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: j, Package: "models", Struct: "toggle", Function: "Unpack"}.Wrap(err)
	}
	return nil
}