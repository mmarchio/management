package types

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type ShallowContent struct {
	ShallowModel
	Content string `json:"content"`
}

func (c ShallowContent) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	m.Content = c.Content
	return &m, nil
}

func (c ShallowContent) Expand(e echo.Context) (*Content, error) {
	r := Content{}
	if c.ShallowModel.CreatedAt.IsZero() && c.ShallowModel.ID != "" {
		sc, err := c.ShallowModel.Get(e)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err)
		}
		if err := json.Unmarshal([]byte(sc.Content), &r); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err)
		}
		return &r, nil
	}
	r.Content = c.Content
	return &r, nil
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

func (c *ShallowContent) Get(e echo.Context) (ShallowContent, error) {
	contentModel :=  models.ShallowContent{}
	contentModel.ShallowModel.ID = c.ShallowModel.ID
	contentModel.ID = c.ID
	err := contentModel.Get(e)
	if err != nil {
		return *c, merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	d := c.FromModel(contentModel)
	return d, nil
}

func (c ShallowContent) CustomQuery(e echo.Context, write bool, q string, vars ...any) ([]ShallowContent, error) {
	if write {
		contentModel := c.ToModel()
		contentModel.ShallowModel.ID = c.ShallowModel.ID
		contentModel.ID = contentModel.ShallowModel.ID
		contentModel.ContentType = c.ContentType
		_, err := contentModel.CustomQuery(e, write, q, vars...)
		if err != nil {
			return nil, merrors.ContentCustomQueryError{Info: c.ShallowModel.ID, Package: "types", Struct: "Content", Function: "CustomQuery"}.Wrap(err)
		}
		return nil, nil
	}
	contentModel := models.ShallowContent{}
	contentModel.ShallowModel.ID = c.ShallowModel.ID
	contentModel.ID = contentModel.ShallowModel.ID
	contentModel.ContentType = c.ContentType
	res, err := contentModel.CustomQuery(e, write, q, vars...)
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

func (c ShallowContent) Set(e echo.Context) error {
	contentModel := c.ToModel()
	contentModel.ShallowModel.ID = c.ShallowModel.ID
	contentModel.ID = c.ShallowModel.ID
	err := contentModel.Set(e)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c *ShallowContent) FindBy(e echo.Context, key, value string) (ShallowContent, error) {
	contentModel := models.ShallowContent{}
	contentModel.ShallowModel.ID = c.ShallowModel.ID
	if err := contentModel.FindBy(e, key, value); err != nil {
		return *c, merrors.ContentFindByError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	if contentModel.Content == "" {
		return *c, merrors.NilContentError{Package: "types", Struct: "ShallowContent", Function: "FindBy"}.New("nil content error").BubbleCode()
	}
	d := c.FromModel(contentModel)
	return d, nil
}

func (c ShallowContent) Delete(e echo.Context) error {
	contentModel := models.Content{}
	contentModel.Model.ID = c.ShallowModel.ID
	contentModel.ID = c.ID
	if err := contentModel.Delete(e); err != nil {
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
	m.ID = c.ID
	m.CreatedAt = c.CreatedAt
	m.UpdatedAt = c.UpdatedAt
	m.ContentType = c.ContentType
	m.TokenCount = c.TokenCount
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

func (c ShallowContent) IsShallowModel() bool {
	return true
}
