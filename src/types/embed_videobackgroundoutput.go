package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type VideoBackgroundOutput struct {
	EmbedModel
	ID 				VideoBackgroundOutputID `json:"id"`
	StatsModel 		Stats `json:"stats_model"`
	FilesArrayModel []File `json:"files_array_model"`
}

func (c VideoBackgroundOutput) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowVideoBackgroundOutput{}
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

func (c *VideoBackgroundOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c VideoBackgroundOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c VideoBackgroundOutput) New() VideoBackgroundOutput {
	c.ID = VideoBackgroundOutputID(uuid.NewString())
	c.StatsModel = c.StatsModel.New(nil)
	return c
}
