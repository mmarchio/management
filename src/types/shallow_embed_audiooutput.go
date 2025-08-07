package types

import (
	"context"
	"encoding/json"

)

type ShallowAudioOutput struct {
	ShallowModel
	ID AudioOutputID `json:"id"`
	StatsModel string `json:"stats_model"`
	FilesArrayModel []string `json:"files_array_model"`
}

func (c *ShallowAudioOutput) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowAudioOutput) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
