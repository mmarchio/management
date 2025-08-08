package types

import (
	"context"
	"encoding/json"

)

type VideoOutput struct {
	EmbedModel
	ID 				VideoOutputID `json:"id"`
	StatsModel 		Stats `json:"stats_model"`
	FilesArrayModel []File `json:"files_model"`
}

func (c VideoOutput) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowVideoOutput{}
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

func (c *VideoOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c VideoOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	c.StatsModel = c.StatsModel.New(nil)
	return string(b), err
}


