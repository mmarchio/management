package models

import (
	"context"
	"encoding/json"
)

type Template struct {
	ID 					string 				`json:"id"`
	Name 				string 				`json:"name"`
	Dispositions 		[]Disposition `json:"dispositions"`
	CurrentDisposition 	int64 				`json:"current_disposition"`
	Base64Value 		string
	JsonValue			[]byte
}

func (c *Template) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Template) Marshal(ctx context.Context) error {
	c.JsonValue, err = json.Marshal(c)
	return err
}