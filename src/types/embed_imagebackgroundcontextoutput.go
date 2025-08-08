package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type ImageBackgroundContextOutput struct {
	EmbedModel
	ID 					ImageBackgroundContextOutputID `json:"id"`
	StatsModel 			Stats `json:"stats_model"`
	FilesArrayModel 	[]File `json:"files_array_model"`
}

func (c ImageBackgroundContextOutput) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowImageBackgroundContextOutput{}
	sm.ShallowModel = sm.ShallowModel.FromEmbedModel(c.EmbedModel)
	sm.ID = c.ID
	sm.StatsModel = c.StatsModel.EmbedModel.ID
	sms = append(sms, c.StatsModel.Pack()...)
	sm.FilesArrayModel = make([]string, 0)
	for _, id := range c.FilesArrayModel {
		sm.FilesArrayModel = append(sm.FilesArrayModel, id.EmbedModel.ID)
		sms = append(sms, id.Pack()...)
	}
	sms = append(sms, sm)
	return sms
}

func (c *ImageBackgroundContextOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ImageBackgroundContextOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c ImageBackgroundContextOutput) New() ImageBackgroundContextOutput {
	c.ID = ImageBackgroundContextOutputID(uuid.NewString())
	c.StatsModel = c.StatsModel.New(nil)
	return c
}