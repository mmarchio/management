package types

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
)


type ShallowStats struct {
	ShallowModel
	ID 				StatsID `json:"stats_id"`
	Start 			time.Time `json:"start"`
	End 			time.Time `json:"end"`
	Input 			string `json:"input"`
	Output 			string `json:"output"`
	Duration 		time.Duration `json:"duration"`
	FilesArrayModel []string `json:"files_array_model"`
	Status 			string `json:"status"`
}

func (c ShallowStats) Expand(ctx context.Context) (*Stats, error) {
	r := Stats{}
	r.EmbedModel.ID = c.ShallowModel.ID
	r.EmbedModel.CreatedAt = c.ShallowModel.CreatedAt
	r.EmbedModel.UpdatedAt = c.ShallowModel.UpdatedAt
	r.EmbedModel.ContentType = c.ShallowModel.ContentType
	r.ID = c.ID
	r.Start = c.Start
	r.End = c.End
	r.Input = c.Input
	r.Output = c.Output
	r.Duration = c.Duration
	f := File{}
	fs := make([]File, 0)
	for _, id := range c.FilesArrayModel {
		sf := ShallowFile{}
		sf.ShallowModel.ID = id
		sc, err := sf.ShallowModel.Get(ctx)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err)
		}
		if err := json.Unmarshal([]byte(sc.Content), &f); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err)
		}
		fs = append(fs, f)
	}
	r.FilesArrayModel = fs
	r.Status = c.Status
	return &r, nil
}

func (c *ShallowStats) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowStats) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ShallowStats) New(id *string) ShallowStats {
	if id != nil {
		c.ID = StatsID(*id)
	} else {
		c.ID = StatsID(uuid.NewString())
	}
	c.Status = "queued"
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	c.ContentType = "contextitem"
	return c
}

func (c ShallowStats) IsShallowModel() bool {
	return true
}

