package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
)

type ShallowImageThumbnailOutput struct {
	ShallowModel
	ID 					ImageThumbnailOutputID `json:"id"`
	StatsModel 			string `json:"stats_model"`
	FilesArrayModel 	[]string `json:"files_array_model"`
}

func (c ShallowImageThumbnailOutput) Expand(ctx context.Context) (*ImageThumbnailOutput, error) {
	r := ImageThumbnailOutput{}
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
	ss := ShallowStats{}
	ss.ShallowModel.ID = c.StatsModel
	stats, err := ss.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.StatsModel = *stats
	r.FilesArrayModel = make([]File, 0)
	for _, id := range c.FilesArrayModel {
		sf := ShallowFile{}
		sf.ShallowModel.ID = id
		f, err := sf.Expand(ctx)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err)
		}
		r.FilesArrayModel = append(r.FilesArrayModel, *f)
	}
	return &r, nil
}

func (c *ShallowImageThumbnailOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowImageThumbnailOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowImageThumbnailOutput) New() ShallowImageThumbnailOutput {
	c.ID = ImageThumbnailOutputID(uuid.NewString())
	c.StatsModel = ShallowStats{}.New(nil).ShallowModel.ID
	return c
}

func (c ShallowImageThumbnailOutput) IsShallowModel() bool {
	return true
}
