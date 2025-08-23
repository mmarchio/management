package types

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

func (c Content) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowContent{}
	sm.ShallowModel = sm.ShallowModel.FromTypeModel(c.Model)
	sm.Content = c.Content
	sms = append(sms, sm)
	return sms
}

func (c Content) CheckType(e echo.Context) (bool, error) {
	m := models.Content{}
	m.Model.ID = c.Model.ID
	m.Model.CreatedAt = c.Model.CreatedAt
	m.Model.UpdatedAt = c.Model.UpdatedAt
	m.Model.ContentType = c.Model.ContentType
	count, err := m.Check(e, m.Model.ID)
	if err != nil {
		return false, merrors.ContentCheckError{CalledBy: "models.Content.Check"}.Wrap(err).Log()
	}
	if count == 0 {
		return false, nil
	}
	if count > int64(0) {
		if err := m.Get(e); err != nil {
			return false, merrors.ContentGetError{CalledBy: "models.Content.Check"}.Wrap(err).Log()
		}
		if m.Model.ContentType == c.Model.ContentType {
			return true, nil
		}
	}

	return false, nil
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

func (c *Content) Get(e echo.Context) (Content, error) {
	contentModel :=  models.Content{}
	contentModel.Model.ID = c.Model.ID
	if err := contentModel.Get(e); err != nil {
		return *c, merrors.ContentGetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	d := c.FromModel(contentModel)
	return d, nil
}

func (c Content) CustomQuery(e echo.Context, write bool, q string, vars ...any) ([]Content, error) {
	if write {
		contentModel := c.ToModel()
		contentModel.Model.ID = c.Model.ID
		contentModel.ContentType = c.ContentType
		_, err := contentModel.CustomQuery(e, write, q, vars...)
		if err != nil {
			return nil, merrors.ContentCustomQueryError{Info: c.Model.ID, Package: "types", Struct: "Content", Function: "CustomQuery"}.Wrap(err).Log()
		}
		return nil, nil
	}
	contentModel := models.Content{}
	contentModel.Model.ID = c.Model.ID
	contentModel.ID = contentModel.Model.ID
	contentModel.ContentType = c.ContentType
	res, err := contentModel.CustomQuery(e, write, q, vars...)
	if err != nil {
		return nil, merrors.ContentCustomQueryError{Info: c.Model.ID, Package: "types", Struct: "Content", Function: "CustomQuery"}.Wrap(err).Log().BubbleCode()
	}
	r := make([]Content, 0)
	for _, t := range res {
		d := c.FromModel(t)
		r = append(r, d)
	}
	return r, nil
}

func (c Content) Set(e echo.Context, update bool) error {
	GetLogger(4).Flogger("#######Content Set called#######")
	contentModel := c.ToModel()
	err := contentModel.Set(e, update)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	return nil
}

func (c *Content) FindBy(e echo.Context, key, value string) (Content, error) {
	contentModel := models.Content{}
	contentModel.Model.ID = c.Model.ID
	if err := contentModel.FindBy(e, key, value); err != nil {
		return *c, merrors.ContentFindByError{Info: c.Model.ID}.Wrap(err).Log()
	}
	if contentModel.Content == "" {
		return *c, merrors.NilContentError{Package: "types", Struct: "Content", Function: "FindBy"}.New("nil content error").BubbleCode()
	}
	d := c.FromModel(contentModel)
	return d, nil
}

func (c Content) List(e echo.Context) ([]Content, error) {
	GetLogger(3).Flogger("types.Content.List Called")
	contentModel := models.Content{}
	contentModel.Model.ContentType = c.Model.ContentType
	contentModels, err := contentModel.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	contents := make([]Content, 0)
	for _, model := range contentModels {
		content := Content{}
		content.FromModel(model)
		contents = append(contents, content)
	}
	return contents, nil
}

func (c Content) ListBy(e echo.Context, key string, value interface{}) ([]Content, error) {
	contentModel := models.Content{}
	contentModels, err := contentModel.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{Info: fmt.Sprintf("content type: %s, filter: %s:%v", c.Model.ContentType, key, value)}.Wrap(err).Log()
	}
	contents := make([]Content, 0)
	for _, model := range contentModels {
		content := Content{}
		content.FromModel(model)
		contents = append(contents, content)
	}
	return contents, nil
}

func (c Content) Delete(e echo.Context) error {
	contentModel := models.Content{}
	contentModel.Model.ID = c.Model.ID
	if err := contentModel.Delete(e); err != nil {
		return merrors.ContentModelDeleteError{}.Wrap(err).Log()
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

func (c *Content) FromType(m ITable, model Model) error {
	b, err := json.Marshal(m)
	if err != nil {
		return merrors.JSONMarshallingError{Info: m.GetContentType()}.Wrap(err).Log()
	}
	c.Model = model
	c.Content = string(b)
	return nil
}

func (c Content) UnrestrictedCustomQuery(e echo.Context, write bool, q string, vars ...any) ([]Content, error) {
	return c.CustomQuery(e, write, q, vars...)
}