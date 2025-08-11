package models

import (
	"time"

	// "encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
)

type ShallowPrompt struct {
	Model
	Name 		string 		`form:"name" json:"name"`
	Prompt 		string 		`form:"prompt" json:"prompt"`
	Domain 		string 		`form:"domain" json:"domain"`
	Category 	string 		`form:"category" json:"category"`
	Settings 	string		`form:"settings" json:"settings"`
}

func NewShallowPrompt(id *string) ShallowPrompt {
	c := ShallowPrompt{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_prompt"
	return c
}

func (c ShallowPrompt) Get(e echo.Context, mode string) (*Prompt, *ShallowPrompt, error) {
	return nil, nil, nil
}

type Prompt struct {
	Model
	Name string
	Prompt string
	Domain string
	Category string
	Settings Settings
	JsonValue []byte
	Base64Value string
}

func (c Prompt) Scan(e echo.Context, rows Scannable) (ITable, error) {
	for rows.Next() {
		err := rows.Scan(&c.Model.ID, &c.Model.CreatedAt, &c.Model.UpdatedAt, &c.Name, &c.Domain, &c.Category, &c.Settings)
		if err != nil {
			return nil, merrors.DBContentScanError{}.Wrap(nil, err)
		}
	}
	return c, nil
}

func (c Prompt) Values(e echo.Context) ([]any, error) {
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


func (c *Prompt) Get(e echo.Context) error {
	var err error
	if err != nil {
		return merrors.ContentGetError{}.Wrap(nil, err)
	}
	return nil
}

func (c Prompt) Set(e echo.Context) error {
	err := c.Model.Set(e, c)
	if err != nil {
		return merrors.ContentSetError{}.Wrap(nil, err)
	}
	return nil
}

func (c Prompt) GetSlice(e echo.Context) []Prompt {
	return make([]Prompt, 0)
}

// func (c Prompt) List(e echo.Context) ([]Prompt, error) {
// 	contents, err := c.Model.List(e, c)
// 	if err != nil {
// 		return nil, merrors.PromptListError{}.Wrap(db, err)
// 	}
// 	r := make([]Prompt, 0)
// 	var prompt Prompt
// 	for _, i := range contents {
// 		err = json.Unmarshal([]byte(i.Content), &prompt)
// 		if err != nil {
// 			return nil, merrors.JSONUnmarshallingError{Info: i.Content}.Wrap(db, err)
// 		}
// 		err = prompt.Unpack()
// 		if err != nil {
// 			return nil, merrors.PromptUnpackError{Info: c.Settings}.Wrap(db, err)
// 		}
// 		r = append(r, prompt)
// 	}
// 	return r, nil
// }

func (c Prompt) GetID() string {
	return c.ID
}

func (c Prompt) GetContentType() string {
	return c.Model.ContentType
}
