package types

import (
	"context"
	"encoding/json"

	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type ShallowSteps struct {
	ShallowModel
	ID 									StepsID `json:"id"`
	GetResearchOutputModel				string `json:"get_research_output_model"`
	GetResearchPromptModel				string `json:"get_research_prompt_model"`
	ScreenwritingStartModel 			string `json:"screenwriting_start_model"`
	ScreenwritingGetPromptInputModel 	string `json:"screenwriting_get_prompt_input_model"`
	ScreenwritingGetPromptOutputModel 	string `json:"screenwriting_get_prompt_output_model"`
	ScreenwritingOutputModel 			string `json:"screenwriting_output_model"`
	ContainerSwapModel 					string `json:"container_swap_model"`
	GenerateAudioModel 					string `json:"generate_audio_model"`
	GenerateLipsyncModel 				string `json:"generate_lipsync_model"`
	GenerateThumbnailsModel 			string `json:"generate_thumbnails_model"`
	GenerateBackgroundContextModel 		string `json:"generate_background_context_model"`
	GenerateBackgroundModel 			string `json:"generate_background_model"`
	FFMPEGLipsyncPostModel 				string `json:"ffmpeg_lipsync_post_model"`
	FFMPEGMergeModel 					string `json:"ffmpeg_merge_model"`
	PublishVideoModel 					string `json:"publish_video_model"`
	PublishThumbnailModel 				string `json:"publish_thumbnail_model"`
	PublishMetadataModel 				string `json:"publish_metadata_model"`
}

func (c *ShallowSteps) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowSteps) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ShallowSteps) New(parent ITable, contentType string) (string, error) {
	embedBytes, err := json.Marshal(c)
	if err != nil {
		return "", merrors.JSONMarshallingError{}.Wrap(err)
	}
	return string(embedBytes), nil
}

func (c ShallowSteps) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowSteps) GetID() string {
	return c.ShallowModel.ID
}

func (c *ShallowSteps) SetID() error {
	var err error
	c.ID = StepsID(c.ShallowModel.ID)
	if err != nil {
		return merrors.IDSetError{Info: "steps"}.Wrap(err)
	}
	return nil
}

func (c ShallowSteps) FromModel(model models.ShallowSteps) ShallowSteps {
	c.ShallowModel.FromModel(model.ShallowModel)
	c.SetID()
	c.GetResearchOutputModel = model.GetResearchOutputModel
	c.GetResearchPromptModel = c.GetResearchPromptModel
	c.ScreenwritingStartModel = c.ScreenwritingStartModel
	c.ScreenwritingGetPromptInputModel = c.ScreenwritingGetPromptInputModel
	c.ScreenwritingGetPromptOutputModel = c.ScreenwritingGetPromptOutputModel
	c.ScreenwritingOutputModel = c.ScreenwritingOutputModel
	c.ContainerSwapModel = c.ContainerSwapModel
	c.GenerateAudioModel = c.GenerateAudioModel
	c.GenerateLipsyncModel = c.GenerateLipsyncModel
	c.GenerateThumbnailsModel = c.GenerateThumbnailsModel
	c.GenerateBackgroundContextModel = c.GenerateBackgroundContextModel
	c.GenerateBackgroundModel = c.GenerateBackgroundModel
	c.FFMPEGLipsyncPostModel = c.FFMPEGLipsyncPostModel
	c.FFMPEGMergeModel = c.FFMPEGMergeModel
	c.PublishVideoModel = c.PublishVideoModel
	c.PublishThumbnailModel = c.PublishThumbnailModel
	c.PublishMetadataModel = c.PublishMetadataModel
	return c
}
