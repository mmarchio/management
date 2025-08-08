package types

import (
	"context"
	"encoding/json"
	"time"

	merrors "github.com/mmarchio/management/errors"
)

type ShallowFile struct {
	ShallowModel
	ID 			FileID `json:"file_id"`
	Type 		string `json:"type"`
	Path 		string `json:"path"`
	Duration 	time.Duration `json:"duration"`
	Scene 		SceneID `json:"scene"`
	Joined 		bool `json:"joined"`
}

func (c ShallowFile) Expand(ctx context.Context) (*File, error) {
	r := File{}
	if c.ShallowModel.CreatedAt.IsZero() && c.ShallowModel.ID != "" {
		sc, err := c.ShallowModel.Get(ctx)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err)
		}
		if err := json.Unmarshal([]byte(sc.Content), &r); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err)
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

func (c *ShallowFile) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowFile) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowFile) IsShallowModel() bool {
	return true
}
