package models

import (
	"context"
	"encoding/json"
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

func (c *Steps) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Steps) Marshal(ctx context.Context) error {
	c.JsonValue, err = json.Marshal(c)
	return err
}
