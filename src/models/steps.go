package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Steps struct {
	Model
	GetResearchOutputModel				Toggle `json:"get_research_output_model"`
	GetResearchPromptModel				Toggle `json:"get_research_prompt_model"`
	ScreenwritingStartModel 			Toggle `json:"screenwriting_start_model"`
	ScreenwritingGetPromptInputModel 	Toggle `json:"screenwriting_get_prompt_input_model"`
	ScreenwritingGetPromptOutputModel 	Toggle `json:"screenwriting_get_prompt_output_model"`
	ScreenwritingOutputModel 			Toggle `json:"screenwriting_output_model"`
	ContainerSwapModel 					Toggle `json:"container_swap_model"`
	GenerateAudioModel 					Toggle `json:"generate_audio_model"`
	GenerateLipsyncModel 				Toggle `json:"generate_lipsync_model"`
	GenerateThumbnailsModel 			Toggle `json:"generate_thumbnails_model"`
	GenerateBackgroundContextModel Toggle `json:"generate_background_context_model"`
	GenerateBackgroundModel 			Toggle `json:"generate_background_model"`
	FFMPEGLipsyncPostModel 				Toggle `json:"ffmpeg_lipsync_post_model"`
	FFMPEGMergeModel 					Toggle `json:"ffmpeg_merge_model"`
	PublishVideoModel 					Toggle `json:"publish_video_model"`
	PublishThumbnailModel 				Toggle `json:"publish_thumbnail_model"`
	PublishMetadataModel 				Toggle `json:"publish_metadata_model"`
	Base64Value						string
	JsonValue						[]byte
}

type ShallowSteps struct {
	ShallowModel
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
	GenerateBackgroundContextModel		string `json:"generate_background_context_model"`
	GenerateBackgroundModel 			string `json:"generate_background_model"`
	FFMPEGLipsyncPostModel 				string `json:"ffmpeg_lipsync_post_model"`
	FFMPEGMergeModel 					string `json:"ffmpeg_merge_model"`
	PublishVideoModel 					string `json:"publish_video_model"`
	PublishThumbnailModel 				string `json:"publish_thumbnail_model"`
	PublishMetadataModel 				string `json:"publish_metadata_model"`
	Base64Value						string
	JsonValue						[]byte
}

func NewShallowSteps(id *string) ShallowSteps {
	c := ShallowSteps{}
	ct := "shallowsteps"
	c.ShallowModel = c.ShallowModel.New(nil, &ct)
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallowsteps"
	return c
}

func (c *Steps) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Steps) Marshal(ctx context.Context) error {
	c.JsonValue, err = json.Marshal(c)
	return err
}
