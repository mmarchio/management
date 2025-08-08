package types

import (
	"context"
	"encoding/json"
	"time"

	merrors "github.com/mmarchio/management/errors"
)

type ShallowScene struct {
	ShallowModel
	ID 					SceneID `json:"id"`
	Start 				time.Time `json:"start"`
	End 				time.Time `json:"end"`
	SceneNumber 		int64 `json:"scene_number"`
	Path 				string `json:"path"`
	FilesArrayModel 	[]string `json:"files_array_model"`
	SceneFileModel 		string `json:"scene_file_model"`
}

func (c ShallowScene) Expand(ctx context.Context) (*Scene, error) {
	r := Scene{}
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
	r.Start = c.Start
	r.End = c.End
	r.SceneNumber = r.SceneNumber
	r.Path = c.Path
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
	sf := ShallowFile{}
	sf.ShallowModel.ID = c.SceneFileModel
	f, err := sf.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.SceneFileModel = *f
	return &r, nil	
}

func (c *ShallowScene) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowScene) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowScene) IsShallowModel() bool {
	return true
}
