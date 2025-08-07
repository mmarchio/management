package types

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type ShallowContent struct {
	ShallowModel
	Content string `json:"content"`
}


func (c ShallowContent) New(ct string) ShallowContent {
	c.ShallowModel.ID = uuid.New().String()
	c.ShallowModel.CreatedAt = time.Now()
	c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
	c.ShallowModel.ContentType = ct
	c.ShallowModel.Table = "content"
	c.ShallowModel.Columns = "id, created_at, updated_at, content_type, content"
	c.ShallowModel.Values = "$1, $2, $3, $4, $5::jsonb"
	c.ShallowModel.Conflict = "DO UPDATE SET updated_at = $3, content = $5::jsonb"
	return c
} 

func (c *ShallowContent) Get(ctx context.Context) (ShallowContent, error) {
	contentModel :=  models.ShallowContent{}
	contentModel.ShallowModel.ID = c.ShallowModel.ID
	contentModel.ID = c.ID
	err := contentModel.Get(ctx)
	if err != nil {
		return *c, merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	d := c.FromModel(contentModel)
	return d, nil
}

func (c ShallowContent) CustomQuery(ctx context.Context, write bool, q string, vars ...any) ([]ShallowContent, error) {
	if write {
		contentModel := c.ToModel()
		contentModel.ShallowModel.ID = c.ShallowModel.ID
		contentModel.ID = contentModel.ShallowModel.ID
		contentModel.ContentType = c.ContentType
		_, err := contentModel.CustomQuery(ctx, write, q, vars...)
		if err != nil {
			return nil, merrors.ContentCustomQueryError{Info: c.ShallowModel.ID, Package: "types", Struct: "Content", Function: "CustomQuery"}.Wrap(err)
		}
		return nil, nil
	}
	contentModel := models.ShallowContent{}
	contentModel.ShallowModel.ID = c.ShallowModel.ID
	contentModel.ID = contentModel.ShallowModel.ID
	contentModel.ContentType = c.ContentType
	res, err := contentModel.CustomQuery(ctx, write, q, vars...)
	if err != nil {
		return nil, merrors.ContentCustomQueryError{Info: c.ShallowModel.ID, Package: "types", Struct: "Content", Function: "CustomQuery"}.Wrap(err).BubbleCode()
	}
	r := make([]ShallowContent, 0)
	for _, t := range res {
		d := c.FromModel(t)
		r = append(r, d)
	}
	return r, nil
}

func (c ShallowContent) Set(ctx context.Context) error {
	contentModel := c.ToModel()
	contentModel.ShallowModel.ID = c.ShallowModel.ID
	contentModel.ID = c.ShallowModel.ID
	err := contentModel.Set(ctx)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c *ShallowContent) FindBy(ctx context.Context, key, value string) (ShallowContent, error) {
	contentModel := models.ShallowContent{}
	contentModel.ShallowModel.ID = c.ShallowModel.ID
	if err := contentModel.FindBy(ctx, key, value); err != nil {
		return *c, merrors.ContentFindByError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	if contentModel.Content == "" {
		return *c, merrors.NilContentError{Package: "types", Struct: "ShallowContent", Function: "FindBy"}.Wrap(fmt.Errorf("nil content error")).BubbleCode()
	}
	d := c.FromModel(contentModel)
	return d, nil
}

func (c ShallowContent) Delete(ctx context.Context) error {
	contentModel := models.Content{}
	contentModel.Model.ID = c.ShallowModel.ID
	contentModel.ID = c.ID
	if err := contentModel.Delete(ctx); err != nil {
		return merrors.ContentModelDeleteError{}.Wrap(err)
	}
	return nil
}

func (c *ShallowContent) FromModel(m models.ShallowContent) ShallowContent {
	c.New(m.ShallowModel.ContentType)
	c.ShallowModel.FromModel(m.ShallowModel)
	c.Content = m.Content
	return *c
}

func (c ShallowContent) ToModel() models.ShallowContent {
	m := models.ShallowContent{}
	m.ShallowModel = c.ShallowModel.ToModel()
	m.Content = c.Content
	return m
}

func (c *ShallowContent) FromType(m ITable) error {
	b, err := json.Marshal(m)
	if err != nil {
		return merrors.JSONMarshallingError{Info: m.GetContentType()}.Wrap(err)
	}
	c.Content = string(b)
	return nil
}
