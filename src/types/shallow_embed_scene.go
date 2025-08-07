package types

import (
	"context"
	"encoding/json"
	"time"

)

type ShallowScene struct {
	ShallowModel
	ID 					SceneID `json:"id"`
	Start 				time.Time `json:"start"`
	End 				time.Time `json:"end"`
	SceneNumber 		int64 `json:"scene_number"`
	Path 				string `json:"path"`
	FilesArrayModel 	[]string `json:"files_array_model"`
	SceneFileModel 		string `json:"scene_file_model"`
}

func (c *ShallowScene) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowScene) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

