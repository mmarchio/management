package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Template struct {
	Model
	ID 					string 				`json:"id"`
	Name 				string 				`json:"name"`
	Dispositions 		[]Disposition 		`json:"dispositions"`
	CurrentDisposition 	int64 				`json:"current_disposition"`
	Base64Value 		string
	JsonValue			[]byte
}

type ShallowTemplate struct {
	Model
	ID 					string 				`json:"id"`
	Name 				string 				`json:"name"`
	Dispositions 		[]string			`json:"dispositions"`
	CurrentDisposition 	int64 				`json:"current_disposition"`
	Base64Value 		string
	JsonValue			[]byte
}

func NewShallowSTemplate(id *string) ShallowTemplate {
	c := ShallowTemplate{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_template"
	return c
}


func (c *Template) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Template) Marshal(ctx context.Context) error {
	c.JsonValue, err = json.Marshal(c)
	return err
}