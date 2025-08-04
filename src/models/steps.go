package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Steps struct {
	Model
	GetResearchOutput				Toggle `json:"get_research_output"`
	GetResearchPrompt				Toggle `json:"get_research_prompt"`
	ScreenwritingStart 				Toggle `json:"screenwriting_start"`
	ScreenwritingGetPromptInput 	Toggle `json:"screenwriting_get_prompt_input"`
	ScreenwritingGetPromptOutput 	Toggle `json:"screenwriting_get_prompt_output"`
	ScreenwritingOutput 			Toggle `json:"screenwriting_output"`
	ContainerSwap 					Toggle `json:"container_swap"`
	GenerateAudio 					Toggle `json:"generate_audio"`
	GenerateLipsync 				Toggle `json:"generate_lipsync"`
	GenerateThumbnails 				Toggle `json:"generate_thumbnails"`
	GenerateBackgroundCountext 		Toggle `json:"generate_background_context"`
	GenerateBackground 				Toggle `json:"generate_background"`
	FFMPEGLipsyncPost 				Toggle `json:"ffmpeg_lipsync_post"`
	FFMPEGMerge 					Toggle `json:"ffmpeg_merge"`
	PublishVideo 					Toggle `json:"publish_video"`
	PublishThumbnail 				Toggle `json:"publish_thumbnail"`
	PublishMetadata 				Toggle `json:"publish_metadata"`
	Base64Value						string
	JsonValue						[]byte
}

type ShallowSteps struct {
	Model
	GetResearchOutput				string `json:"get_research_output"`
	GetResearchPrompt				string `json:"get_research_prompt"`
	ScreenwritingStart 				string `json:"screenwriting_start"`
	ScreenwritingGetPromptInput 	string `json:"screenwriting_get_prompt_input"`
	ScreenwritingGetPromptOutput 	string `json:"screenwriting_get_prompt_output"`
	ScreenwritingOutput 			string `json:"screenwriting_output"`
	ContainerSwap 					string `json:"container_swap"`
	GenerateAudio 					string `json:"generate_audio"`
	GenerateLipsync 				string `json:"generate_lipsync"`
	GenerateThumbnails 				string `json:"generate_thumbnails"`
	GenerateBackgroundCountext 		string `json:"generate_background_context"`
	GenerateBackground 				string `json:"generate_background"`
	FFMPEGLipsyncPost 				string `json:"ffmpeg_lipsync_post"`
	FFMPEGMerge 					string `json:"ffmpeg_merge"`
	PublishVideo 					string `json:"publish_video"`
	PublishThumbnail 				string `json:"publish_thumbnail"`
	PublishMetadata 				string `json:"publish_metadata"`
	Base64Value						string
	JsonValue						[]byte
}

func NewShallowSteps(id *string) ShallowSteps {
	c := ShallowSteps{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_steps"
	return c
}

func (c *Steps) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Steps) Marshal(ctx context.Context) error {
	c.JsonValue, err = json.Marshal(c)
	return err
}
