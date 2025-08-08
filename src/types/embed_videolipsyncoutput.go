package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type VideoLipsyncOutput struct {
	EmbedModel
	ID 				VideoLipsyncOutputID `json:"id"`
	StatsModel 		Stats `json:"stats_model"`
	FilesArrayModel []File `json:"files_array_model"`
}

func (c VideoLipsyncOutput) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowVideoLipsyncOutput{}
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

func (c *VideoLipsyncOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c VideoLipsyncOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c VideoLipsyncOutput) New() VideoLipsyncOutput {
	c.ID = VideoLipsyncOutputID(uuid.NewString())
	c.StatsModel = c.StatsModel.New(nil)
	return c
}

