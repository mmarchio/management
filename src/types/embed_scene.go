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

func (c Scene) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowScene{}
	sm.ShallowModel = sm.ShallowModel.FromEmbedModel(c.EmbedModel)
	sm.ID = c.ID
	sm.Start = c.Start
	sm.End = c.End
	sm.SceneNumber = c.SceneNumber
	sm.Path = c.Path
	sm.FilesArrayModel = make([]string, 0)
	for _, id := range c.FilesArrayModel {
		sm.FilesArrayModel = append(sm.FilesArrayModel, id.EmbedModel.ID)
		sms = append(sms, id.Pack()...)
	}
	sm.SceneFileModel = c.SceneFileModel.EmbedModel.ID
	sms = append(sms, c.SceneFileModel.Pack()...)
	sms = append(sms, sm)
	return sms
}

func (c *Scene) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Scene) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

