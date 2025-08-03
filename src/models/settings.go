package models

import (
	"context"
	"encoding/json"
)

type Settings struct {
	ID 				string `json:"id"`
	Template 		Template `json:"template_id"`
	GlobalBypass 	Steps `json:"global_bypass"`
	Recurring 		Toggle `json:"recurring"`
	Interval 		int64 `json:"interval"`
	Base64Value		string
	JsonValue		[]byte
}

func (c *Settings) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Settings) Marshal(ctx context.Context) error {
	c.JsonValue, err = json.Marshal(c)
	return err
}

