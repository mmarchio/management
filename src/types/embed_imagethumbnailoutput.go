package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type ImageThumbnailOutput struct {
	EmbedModel
	ID 					ImageThumbnailOutputID `json:"id"`
	StatsModel 			Stats `json:"stats_model"`
	FilesArrayModel 	[]File `json:"files_array_model"`
}

func (c *ImageThumbnailOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ImageThumbnailOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ImageThumbnailOutput) New() ImageThumbnailOutput {
	c.ID = ImageThumbnailOutputID(uuid.NewString())
	c.StatsModel = c.StatsModel.New(c.ID.String())
	return c
}
