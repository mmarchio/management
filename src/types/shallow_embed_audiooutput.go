package types

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
)

type ShallowAudioOutput struct {
	ShallowModel
	ID AudioOutputID `json:"id"`
	StatsModel string `json:"stats_model"`
	FilesArrayModel []string `json:"files_array_model"`
}

func (c ShallowAudioOutput) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(nil, err)
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowAudioOutput) Expand(e echo.Context) (*AudioOutput, error) {
	r := AudioOutput{}
	if c.ShallowModel.CreatedAt.IsZero() && c.ShallowModel.ID != "" {
		sc, err := c.ShallowModel.Get(e)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(nil, err)
		}
		if err := json.Unmarshal([]byte(sc.Content), &r); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(nil, err)
		}
		return &r, nil
	}
	r.EmbedModel = r.EmbedModel.FromShallowModel(c.ShallowModel)
	ss := ShallowStats{}
	ss.ShallowModel.ID = c.StatsModel
	stats, err := ss.Expand(e)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(nil, err)
	}
	r.StatsModel = *stats
	r.FilesArrayModel = make([]File, 0)
	for _, id := range c.FilesArrayModel {
		sf := ShallowFile{}
		sf.ShallowModel.ID = id
		f, err := sf.Expand(e)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(nil, err)
		}
		r.FilesArrayModel = append(r.FilesArrayModel, *f)
	}
	return &r, nil
}

func (c *ShallowAudioOutput) Unmarshal(e echo.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowAudioOutput) Marshal(e echo.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowAudioOutput) IsShallowModel() bool {
	return true
}
