package types

import (
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
)

type ShallowFile struct {
	ShallowModel
	Type 		string `json:"type"`
	Path 		string `json:"path"`
	Duration 	time.Duration `json:"duration"`
	Scene 		SceneID `json:"scene"`
	Joined 		bool `json:"joined"`
}

func (c ShallowFile) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowFile) Expand(e echo.Context) (*File, error) {
	r := File{}
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
	r.ID = c.ID
	r.Type = c.Type
	r.Path = c.Path
	r.Duration = c.Duration
	r.Scene = c.Scene
	r.Joined = c.Joined
	return &r, nil
}

func (c *ShallowFile) Unmarshal(e echo.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowFile) Marshal(e echo.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowFile) IsShallowModel() bool {
	return true
}
