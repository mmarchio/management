package models

import (
	"context"
	"encoding/json"

	merrors "github.com/mmarchio/management/errors"
)

type Toggle struct {
	ID 			string 	`json:"id"`
	NamePrefix 	string 	`json:"name_prefix"`
	IdPrefix 	string 	`json:"id_suffix"`
	Suffix 		string 	`json:"suffix"`
	Value 		bool 	`json:"value"`
	Title 		string 	`json:"title"`
	Base64Value string
	JsonValue	[]byte
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