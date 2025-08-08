package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Settings struct {
	Model
	ID 				string `json:"id"`
	Template 		Template `json:"template_id"`
	GlobalBypass 	Steps `json:"global_bypass"`
	Recurring 		Toggle `json:"recurring"`
	Interval 		int64 `json:"interval"`
	Base64Value		string
	JsonValue		[]byte
}

type ShallowSettings struct {
	Model
	ID 				string `json:"id"`
	Template 		string `json:"template_id"`
	GlobalBypass 	string `json:"global_bypass"`
	Recurring 		string `json:"recurring"`
	Interval 		int64  `json:"interval"`
	Base64Value		string
	JsonValue		[]byte
}

func NewShallowSettings(id *string) ShallowSettings {
	c := ShallowSettings{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_settings"
	return c
}

func (c *Settings) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Settings) Marshal(ctx context.Context) error {
	c.JsonValue, err = json.Marshal(c)
	return err
}

