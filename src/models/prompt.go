package models

import (
	"context"
	"encoding/base64"
	// "encoding/json"
	merrors "github.com/mmarchio/management/errors"
)

type Prompt struct {
	Model
	Name 		string 		`form:"name" json:"name"`
	Prompt 		string 		`form:"prompt" json:"prompt"`
	Domain 		string 		`form:"domain" json:"domain"`
	Category 	string 		`form:"category" json:"category"`
	Settings 	string		`form:"settings" json:"settings"`
}

type nestedPrompt struct {
	Model
	Name string
	Prompt string
	Domain string
	Category string
	Settings Settings
	JsonValue []byte
	Base64Value string
}

func (c *Prompt) ToBase64() {
	c.Settings = base64.StdEncoding.EncodeToString([]byte(c.Settings))
}

func (c Prompt) Scan(ctx context.Context, rows Scannable) (ITable, error) {
	for rows.Next() {
		err := rows.Scan(&c.Model.ID, &c.Model.CreatedAt, &c.Model.UpdatedAt, &c.Name, &c.Domain, &c.Category, &c.Settings)
		if err != nil {
			return nil, merrors.PromptModelScanError{}.Wrap(err)
		}
	}
	return c, nil
}

func (c Prompt) Values(ctx context.Context) ([]any, error) {
	r := make([]any, 0)
	r = append(r, c.ID)
	r = append(r, c.CreatedAt)
	r = append(r, c.UpdatedAt)
	r = append(r, c.Name)
	r = append(r, c.Domain)
	r = append(r, c.Category)
	r = append(r, c.Settings)
	return r, nil
}

func (c *Prompt) New() {
	c.Model.Columns = "id, created_at, updated_at, name, prompt, domain, category, settings"
	c.Model.Values = "$1, $2, $3, $4, $5, $6, $7, $8"
	c.Model.Conflict = "DO UPDATE SET updated_at = $3, name = $4, prompt = $5, domain = $6, category = $7, settings = $8"
}


func (c *Prompt) Get(ctx context.Context) error {
	var err error
	itable, err := c.Model.Get(ctx, c)
	if err != nil {
		return merrors.PromptGetError{}.Wrap(err)
	}
	if entity, ok := itable.(Prompt); ok {
		c = &entity
		err = c.Unpack()
		if err != nil {
			return merrors.PromptUnpackError{}.Wrap(err)
		}
	}
	return nil
}

func (c Prompt) Set(ctx context.Context) error {
	err := c.Model.Set(ctx, c)
	if err != nil {
		return merrors.PromptSetError{}.Wrap(err)
	}
	return nil
}

func (c Prompt) GetSlice(ctx context.Context) []Prompt {
	return make([]Prompt, 0)
}

// func (c Prompt) List(ctx context.Context) ([]Prompt, error) {
// 	contents, err := c.Model.List(ctx, c)
// 	if err != nil {
// 		return nil, merrors.PromptListError{}.Wrap(err)
// 	}
// 	r := make([]Prompt, 0)
// 	var prompt Prompt
// 	for _, i := range contents {
// 		err = json.Unmarshal([]byte(i.Content), &prompt)
// 		if err != nil {
// 			return nil, merrors.JSONUnmarshallingError{Info: i.Content}.Wrap(err)
// 		}
// 		err = prompt.Unpack()
// 		if err != nil {
// 			return nil, merrors.PromptUnpackError{Info: c.Settings}.Wrap(err)
// 		}
// 		r = append(r, prompt)
// 	}
// 	return r, nil
// }

func (c *Prompt) Unpack() error {
	var err error
	c.Settings, err = FromBase64(c.Settings)
	if err != nil {
		return merrors.Base64DecodingError{Info: c.Settings}.Wrap(err)
	}
	return nil
}

func (c Prompt) GetID() string {
	return c.ID
}

func (c Prompt) GetContentType() string {
	return c.Model.ContentType
}
