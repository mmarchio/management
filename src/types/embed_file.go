package types

import (
	"context"
	"encoding/json"
	"time"

)

type File struct {
	EmbedModel
	ID 			FileID `json:"file_id"`
	Type 		string `json:"type"`
	Path 		string `json:"path"`
	Duration 	time.Duration `json:"duration"`
	Scene 		SceneID `json:"scene"`
	Joined 		bool `json:"joined"`
}

func (c File) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowFile{}
	sm.ShallowModel = sm.ShallowModel.FromEmbedModel(c.EmbedModel)
	sm.ID = c.ID
	sm.Type = c.Type
	sm.Path = c.Path
	sm.Duration = c.Duration
	sm.Scene = c.Scene
	sm.Joined = c.Joined
	sms = append(sms, sm)
	return sms
}

func (c *File) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c File) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

