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

type ContentT string

func (c ContentT) Marshal() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

type Content struct {
	Model
	Content string `json:"content"`
}

func (c Content) Scan(ctx context.Context, rows Scannable) (Content, error) {
	err := rows.Scan(&c.Model.ID, &c.Model.CreatedAt, &c.Model.UpdatedAt, &c.Model.ContentType, &c.Content)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (c Content) Values() []any {
	r := make([]any, 0)
	r = append(r, c.Model.ID)
	r = append(r, c.Model.CreatedAt.Format(time.RFC3339))
	r = append(r, c.Model.UpdatedAt.Format(time.RFC3339))
	r = append(r, c.Model.ContentType)
	r = append(r, c.Content)
	return r
}

func (c Content) New(ct string) Content {
	c.Model.ID = uuid.New().String()
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
	c.Model.ContentType = ct
	c.Model.Table = "content"
	c.Model.Columns = "id, created_at, updated_at, content_type, content"
	c.Model.Values = "$1, $2, $3, $4, $5::jsonb"
	c.Model.Conflict = "DO UPDATE SET updated_at = $3, content = $5::jsonb"
	return c
} 

func (c *Content) Get(ctx context.Context) (Content, error) {
	contentModel :=  models.Content{}
	contentModel.Model.ID = c.Model.ID
	contentModel.ID = c.ID
	err := contentModel.Get(ctx)
	if err != nil {
		return *c, merrors.ContentGetError{Info: c.Model.ID}.Wrap(err)
	}
	d := c.FromModel(contentModel)
	return d, nil
}

func (c Content) CustomQuery(ctx context.Context, write bool, q string, vars ...any) ([]Content, error) {
	if write {
		contentModel := c.ToModel()
		contentModel.Model.ID = c.Model.ID
		contentModel.ID = contentModel.Model.ID
		contentModel.ContentType = c.ContentType
		_, err := contentModel.CustomQuery(ctx, write, q, vars...)
		if err != nil {
			return nil, merrors.ContentCustomQueryError{Info: c.Model.ID, Package: "types", Struct: "Content", Function: "CustomQuery"}.Wrap(err)
		}
		return nil, nil
	}
	contentModel := models.Content{}
	contentModel.Model.ID = c.Model.ID
	contentModel.ID = contentModel.Model.ID
	contentModel.ContentType = c.ContentType
	res, err := contentModel.CustomQuery(ctx, write, q, vars...)
	if err != nil {
		return nil, merrors.ContentCustomQueryError{Info: c.Model.ID, Package: "types", Struct: "Content", Function: "CustomQuery"}.Wrap(err).BubbleCode()
	}
	r := make([]Content, 0)
	for _, t := range res {
		d := c.FromModel(t)
		r = append(r, d)
	}
	return r, nil
}

func (c Content) Set(ctx context.Context) error {
	contentModel := c.ToModel()
	contentModel.Model.ID = c.Model.ID
	contentModel.ID = c.Model.ID
	err := contentModel.Set(ctx)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c *Content) FindBy(ctx context.Context, key, value string) (Content, error) {
	contentModel := models.Content{}
	contentModel.Model.ID = c.Model.ID
	if err := contentModel.FindBy(ctx, key, value); err != nil {
		return *c, merrors.ContentFindByError{Info: c.Model.ID}.Wrap(err)
	}
	if contentModel.Content == "" {
		return *c, merrors.NilContentError{Package: "types", Struct: "Content", Function: "FindBy"}.Wrap(fmt.Errorf("nil content error")).BubbleCode()
	}
	d := c.FromModel(contentModel)
	return d, nil
}

func (c Content) List(ctx context.Context) ([]Content, error) {
	contentModel := models.Content{}
	contentModels, err := contentModel.List(ctx)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err)
	}
	contents := make([]Content, 0)
	for _, model := range contentModels {
		content := Content{}
		content.FromModel(model)
		contents = append(contents, content)
	}
	return contents, nil
}

func (c Content) ListBy(ctx context.Context, key string, value interface{}) ([]Content, error) {
	contentModel := models.Content{}
	contentModels, err := contentModel.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{Info: fmt.Sprintf("content type: %s, filter: %s:%v", c.Model.ContentType, key, value)}.Wrap(err)
	}
	contents := make([]Content, 0)
	for _, model := range contentModels {
		content := Content{}
		content.FromModel(model)
		contents = append(contents, content)
	}
	return contents, nil
}

func (c Content) Delete(ctx context.Context) error {
	contentModel := models.Content{}
	contentModel.Model.ID = c.Model.ID
	contentModel.ID = c.ID
	if err := contentModel.Delete(ctx); err != nil {
		return merrors.ContentModelDeleteError{}.Wrap(err)
	}
	return nil
}

func (c *Content) FromModel(m models.Content) Content {
	c.New(m.Model.ContentType)
	c.Model.FromModel(m.Model)
	c.Content = m.Content
	return *c
}

func (c Content) ToModel() models.Content {
	m := models.Content{}
	m.Model = c.Model.ToModel()
	m.Content = c.Content
	return m
}

func (c *Content) FromType(m ITable) error {
	b, err := json.Marshal(m)
	if err != nil {
		return merrors.JSONMarshallingError{Info: m.GetContentType()}.Wrap(err)
	}
	c.Content = string(b)
	return nil
}
