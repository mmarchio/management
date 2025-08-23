package types

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
)

type ShallowImageBackgroundContextOutput struct {
	ShallowModel
	StatsModel 			string `json:"stats_model"`
	FilesArrayModel 	[]string `json:"files_array_model"`
}

func (c ShallowImageBackgroundContextOutput) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowImageBackgroundContextOutput) Expand(e echo.Context) (*ImageBackgroundContextOutput, error) {
	r := ImageBackgroundContextOutput{}
	if c.ShallowModel.CreatedAt.IsZero() && c.ShallowModel.ID != "" {
		sc, err := c.ShallowModel.Get(e)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err).Log()
		}
		if err := json.Unmarshal([]byte(sc.Content), &r); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
		return &r, nil
	}
	r.EmbedModel = r.EmbedModel.FromShallowModel(c.ShallowModel)
	ss := ShallowStats{}
	ss.ShallowModel.ID = c.StatsModel
	stats, err := ss.Expand(e)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err).Log()
	}
	r.StatsModel = *stats
	r.FilesArrayModel = make([]File, 0)
	for _, id := range c.FilesArrayModel {
		sf := ShallowFile{}
		sf.ShallowModel.ID = id
		f, err := sf.Expand(e)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err).Log()
		}
		r.FilesArrayModel = append(r.FilesArrayModel, *f)
	}
	return &r, nil
}

func (c *ShallowImageBackgroundContextOutput) Unmarshal(e echo.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowImageBackgroundContextOutput) Marshal(e echo.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowImageBackgroundContextOutput) New() ShallowImageBackgroundContextOutput {
	c.ShallowModel.ID = uuid.NewString()
	c.StatsModel = ShallowStats{}.New(nil).ShallowModel.ID
	return c
}

func (c ShallowImageBackgroundContextOutput) IsShallowModel() bool {
	return true
}
