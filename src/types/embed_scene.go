package types

import (
	"context"
	"encoding/json"
	"time"

)

type Scene struct {
	EmbedModel
	ID 					SceneID `json:"id"`
	Start 				time.Time `json:"start"`
	End 				time.Time `json:"end"`
	SceneNumber 		int64 `json:"scene_number"`
	Path 				string `json:"path"`
	FilesArrayModel 	[]File `json:"files_array_model"`
	SceneFileModel 		File `json:"scene_file_model"`
}

func (c *Scene) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Scene) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

