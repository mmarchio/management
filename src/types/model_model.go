package types

import (
	"context"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type IID interface {
	String() string
	IsNil() bool
}

type ShallowModel struct {
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

func (c *ShallowModel) New(id *string) {
	if id != nil {
		c.ID = *id
	} else {
		c.ID = uuid.NewString()
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
}

func (c Model) GetCtx(ctx context.Context) (*Context, error) {
	typesContext := Context{}
	systemContext, err := models.Context{}.GetCtx(ctx)
	if err != nil {
		return nil, merrors.ContextGetError{Package: "types", Struct: "Context", Function: "GetCtx"}.Wrap(err)
	}
	typesContext.FromModel(systemContext)
	return &typesContext, nil
}

func (c Model) SetCtx(ctx context.Context) (context.Context, error) {
	systemContext := Context{}
	s, err := systemContext.ToModel()
	if err != nil {
		return ctx, merrors.SetContextError{Package:"types", Struct:"Context", Function: "SetCtx"}.Wrap(err)
	}
	ctx = s.SetCtx(ctx)
	return ctx, nil
}

func (c ShallowModel) GetCtx(ctx context.Context) (*Context, error) {
	typesContext := Context{}
	systemContext, err := models.Context{}.GetCtx(ctx)
	if err != nil {
		return nil, merrors.ContextGetError{Package: "types", Struct: "Context", Function: "GetCtx"}.Wrap(err)
	}
	typesContext.FromModel(systemContext)
	return &typesContext, nil
}

func (c ShallowModel) SetCtx(ctx context.Context) (context.Context, error) {
	systemContext := Context{}
	s, err := systemContext.ToModel()
	if err != nil {
		return ctx, merrors.SetContextError{Package:"types", Struct:"Context", Function: "SetCtx"}.Wrap(err)
	}
	ctx = s.SetCtx(ctx)
	return ctx, nil
}

type EmbedModel struct {
	ID string
	CreatedAt time.Time
	UpdatedAt time.Time
	TokenCount int64
	ContentType string
}

func (c EmbedModel) GetCtx(ctx context.Context) (*Context, error) {
	typesContext := Context{}
	systemContext, err := models.Context{}.GetCtx(ctx)
	if err != nil {
		return nil, merrors.ContextGetError{Package: "types", Struct: "Context", Function: "GetCtx"}.Wrap(err)
	}
	typesContext.FromModel(systemContext)
	return &typesContext, nil
}

func (c EmbedModel) SetCtx(ctx context.Context) (context.Context, error) {
	systemContext := Context{}
	s, err := systemContext.ToModel()
	if err != nil {
		return ctx, merrors.SetContextError{Package:"types", Struct:"Context", Function: "SetCtx"}.Wrap(err)
	}
	ctx = s.SetCtx(ctx)
	return ctx, nil
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

func (c *ShallowModel) FromModel(m models.ShallowModel) {
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

func (c ShallowModel) ToModel() models.ShallowModel {
	m := models.ShallowModel{}
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