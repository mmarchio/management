package types

import (
	"context"
	"encoding/json"

)

type AudioOutput struct {
	EmbedModel
	ID AudioOutputID `json:"id"`
	StatsModel Stats `json:"stats_model"`
	FilesArrayModel []File `json:"files_array_model"`
}

func (c AudioOutput) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowAudioOutput{}
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

func (c *AudioOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c AudioOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
