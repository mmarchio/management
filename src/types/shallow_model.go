package types

import (
	"context"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type shallowmodel interface{
	IsShallowModel() bool
}

type ShallowModel struct {
	ID 			string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	ContentType string
	TokenCount 	int64
	Table 		string
	Columns 	string
	Values 		string
	Conflict 	string
	Validated   bool
}

func (c ShallowModel) FromTypeModel(m Model) ShallowModel {
	c.ID = m.ID
	c.CreatedAt = m.CreatedAt
	c.UpdatedAt = m.UpdatedAt
	c.ContentType = m.ContentType
	return c
}

func (c ShallowModel) FromEmbedModel(m EmbedModel) ShallowModel {
	c.ID = m.ID
	c.CreatedAt = m.CreatedAt
	c.UpdatedAt = m.UpdatedAt
	c.ContentType = m.ContentType
	return c
}


func (c ShallowModel) Validate() bool {
	valid := true
	if c.ID == "" {
		valid = false
	}
	if c.CreatedAt.IsZero() || c.UpdatedAt.IsZero() {
		valid = false
	}
	return valid
}

func (c ShallowModel) Expand() Model {
	r := Model{}
	r.ID = c.ID
	r.CreatedAt = c.CreatedAt
	r.UpdatedAt = c.UpdatedAt
	r.ContentType = c.ContentType
	r.Validated = c.Validated
	return r
}

func (c *ShallowModel) New(id *string, contentType *string) {
	if id != nil {
		c.ID = *id
	} else {
		c.ID = uuid.NewString()
	}
	if contentType != nil {
		c.ContentType = *contentType
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
}

func (c *ShallowModel) FromModel(m models.ShallowModel) {
	c.ID = m.ID
	c.CreatedAt = m.CreatedAt
	c.UpdatedAt = m.UpdatedAt
	c.ContentType = m.ContentType
}

func (c ShallowModel) ToModel() models.Model {
	m := models.Model{}
	m.ID = c.ID
	m.CreatedAt = c.CreatedAt
	m.UpdatedAt = c.UpdatedAt
	m.ContentType = c.ContentType
	return m
}

func (c ShallowModel) Get(ctx context.Context) (*ShallowContent, error) {
	sc := ShallowContent{}
	sc.ShallowModel = c
	rc, err := sc.Get(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	return &rc, nil
}

func (c ShallowModel) Set(ctx context.Context, content ShallowContent) error {
	if err := content.Set(ctx); err != nil {
		return merrors.ContentSetError{}.Wrap(err)
	}
	return nil
}

func (c ShallowModel) List(ctx context.Context) ([]ShallowContent, error) {
	sc := ShallowContent{}
	list, err := sc.List(ctx)
	if err != nil {
		return nil, merrors.ContentListError{}.Wrap(err)
	}
	return list, nil
}

func (c ShallowModel) ListBy(ctx context.Context, key, value interface{}) ([]*ShallowContent, error) {
	sc := ShallowContent{}
	list, err := sc.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{}.Wrap(err)
	}
	return list, nil
}

func (c ShallowModel) FindBy(ctx context.Context, key, value string) (*ShallowContent, error) {
	sc := ShallowContent{}
	rc, err := sc.FindBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ContentFindByError{}.Wrap(err)
	}
	return &rc, nil
}

func (c ShallowModel) CustomQuery(ctx context.Context, write bool, q string, vars []interface{}) ([]ShallowContent, error) {
	sc := ShallowContent{}
	list, err := sc.CustomQuery(ctx, write, q, vars...)
	if err != nil {
		return nil, merrors.ContentCustomQueryError{}.Wrap(err)
	}
	return list, nil
}

func (c ShallowModel) IsShallowModel() bool {
	return true
}
