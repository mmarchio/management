package types

import (
	"context"
	"encoding/json"
	"time"

)

type ShallowFile struct {
	EmbedModel
	ID 			FileID `json:"file_id"`
	Type 		string `json:"type"`
	Path 		string `json:"path"`
	Duration 	time.Duration `json:"duration"`
	Scene 		SceneID `json:"scene"`
	Joined 		bool `json:"joined"`
}

func (c *ShallowFile) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowFile) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

