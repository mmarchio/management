package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/mmarchio/management/models"
)

type IID interface {
	String() string
	IsNil() bool
}

type Model struct {
	ID 			string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	ContentType string
	Table 		string
	Columns 	string
	Values 		string
	Conflict 	string
	Validated   bool
}

func (c Model) Validate() bool {
	valid := true
	if c.ID == "" {
		valid = false
	}
	if c.CreatedAt.IsZero() || c.UpdatedAt.IsZero() {
		valid = false
	}
	return valid
}

func (c *Model) New(id *string) {
	if id != nil {
		c.ID = *id
	} else {
		c.ID = uuid.NewString()
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
}

type EmbedModel struct {
	ID string
	CreatedAt time.Time
	UpdatedAt time.Time
	TokenCount int64
	ContentType string
}

func (c *EmbedModel) New(contentType string) {
	c.ID = uuid.New().String()
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	c.ContentType = contentType
}

func (c EmbedModel) GetID() string {
	return c.ID
}

func (c *Model) FromModel(m models.Model) {
	c.ID = m.ID
	c.CreatedAt = m.CreatedAt
	c.UpdatedAt = m.UpdatedAt
	c.ContentType = m.ContentType
}

func (c Model) ToModel() models.Model {
	m := models.Model{}
	m.ID = c.ID
	m.CreatedAt = c.CreatedAt
	m.UpdatedAt = c.UpdatedAt
	m.ContentType = c.ContentType
	return m
}

func (c *EmbedModel) FromModel(m models.Model) {
	c.ID = m.ID
	c.CreatedAt = m.CreatedAt
	c.UpdatedAt = m.UpdatedAt
	c.ContentType = m.ContentType
}