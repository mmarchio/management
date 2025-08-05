package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type Steps struct {
	EmbedModel
	ID 									StepsID `json:"id"`
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
	GenerateBackgroundCountextModel 	Toggle `json:"generate_background_context_model"`
	GenerateBackgroundModel 			Toggle `json:"generate_background_model"`
	FFMPEGLipsyncPostModel 				Toggle `json:"ffmpeg_lipsync_post_model"`
	FFMPEGMergeModel 					Toggle `json:"ffmpeg_merge_model"`
	PublishVideoModel 					Toggle `json:"publish_video_model"`
	PublishThumbnailModel 				Toggle `json:"publish_thumbnail_model"`
	PublishMetadataModel 				Toggle `json:"publish_metadata_model"`
}

func (c *Steps) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Steps) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Steps) New(parent ITable, contentType string) (string, error) {
	c.EmbedModel.New(contentType)
	c.EmbedModel.ID = parent.GetID()
	c.GetResearchOutputModel.init()
	c.GetResearchOutputModel.New(*c)
	c.GetResearchPromptModel.init()
	c.GetResearchPromptModel.New(*c)
	c.ScreenwritingStartModel.init()
	c.ScreenwritingStartModel.New(*c)
	c.ScreenwritingGetPromptInputModel.init()
	c.ScreenwritingGetPromptInputModel.New(*c)
	c.ScreenwritingGetPromptOutputModel.init()
	c.ScreenwritingGetPromptOutputModel.New(*c)
	c.ScreenwritingOutputModel.init()
	c.ScreenwritingOutputModel.New(*c)
	c.ContainerSwapModel.init()
	c.ContainerSwapModel.New(*c)
	c.GenerateAudioModel.init()
	c.GenerateAudioModel.New(*c)
	c.GenerateLipsyncModel.init()
	c.GenerateLipsyncModel.New(*c)
	c.GenerateThumbnailsModel.init()
	c.GenerateThumbnailsModel.New(*c)
	c.GenerateBackgroundCountextModel.init()
	c.GenerateBackgroundCountextModel.New(*c)
	c.GenerateBackgroundModel.init()
	c.GenerateBackgroundModel.New(*c)
	c.FFMPEGLipsyncPostModel.init()
	c.FFMPEGLipsyncPostModel.New(*c)
	c.FFMPEGMergeModel.init()
	c.FFMPEGMergeModel.New(*c)
	c.PublishVideoModel.init()
	c.PublishVideoModel.New(*c)
	c.PublishThumbnailModel.init()
	c.PublishThumbnailModel.New(*c)
	c.PublishMetadataModel.init()
	c.PublishMetadataModel.New(*c)
	embedBytes, err := json.Marshal(c)
	if err != nil {
		return "", merrors.JSONMarshallingError{}.Wrap(err)
	}
	return string(embedBytes), nil
}

func (c Steps) GetContentType() string {
	return c.EmbedModel.ContentType
}

func (c Steps) GetID() string {
	return c.EmbedModel.ID
}

func (c *Steps) SetID() error {
	var err error
	c.ID = StepsID(c.EmbedModel.ID)
	if err != nil {
		return merrors.IDSetError{Info: "steps"}.Wrap(err)
	}
	return nil
}

func (c Steps) FromModel(model models.Steps) Steps {
	toggle := Toggle{}
	c.EmbedModel.FromModel(model.Model)
	c.SetID()
	c.GetResearchOutputModel = toggle.FromModel(model.GetResearchOutputModel)
	c.GetResearchPromptModel = toggle.FromModel(model.GetResearchPromptModel)
	c.ScreenwritingStartModel = toggle.FromModel(model.ScreenwritingStartModel)
	c.ScreenwritingGetPromptInputModel = toggle.FromModel(model.ScreenwritingGetPromptInputModel)
	c.ScreenwritingGetPromptOutputModel = toggle.FromModel(model.ScreenwritingGetPromptOutputModel)
	c.ScreenwritingOutputModel = toggle.FromModel(model.ScreenwritingOutputModel)
	c.ContainerSwapModel = toggle.FromModel(model.ContainerSwapModel)
	c.GenerateAudioModel = toggle.FromModel(model.GenerateAudioModel)
	c.GenerateLipsyncModel = toggle.FromModel(model.GenerateLipsyncModel)
	c.GenerateThumbnailsModel = toggle.FromModel(model.GenerateThumbnailsModel)
	c.GenerateBackgroundCountextModel = toggle.FromModel(model.GenerateBackgroundModelCountextModel)
	c.GenerateBackgroundModel = toggle.FromModel(model.GenerateBackgroundModel)
	c.FFMPEGLipsyncPostModel = toggle.FromModel(model.FFMPEGLipsyncPostModel)
	c.FFMPEGMergeModel = toggle.FromModel(model.FFMPEGMergeModel)
	c.PublishVideoModel = toggle.FromModel(model.PublishVideoModel)
	c.PublishThumbnailModel = toggle.FromModel(model.PublishThumbnailModel)
	c.PublishMetadataModel = toggle.FromModel(model.PublishMetadataModel)
	return c
}

func (c Steps) Bind(e echo.Context) (Steps, error) {
	if e.FormValue(c.GetResearchOutputModel.NamePrefix+c.GetResearchOutputModel.Suffix) == "on" {
		c.GetResearchOutputModel.Value = true
	}
	if e.FormValue(c.GetResearchPromptModel.NamePrefix+c.GetResearchPromptModel.Suffix) == "on" {
		c.GetResearchPromptModel.Value = true
	}
	if e.FormValue(c.ScreenwritingStartModel.NamePrefix+c.ScreenwritingStartModel.Suffix) == "on" {
		c.ScreenwritingStartModel.Value = true
	}
	if e.FormValue(c.ScreenwritingGetPromptInputModel.NamePrefix+c.ScreenwritingGetPromptInputModel.Suffix) == "on" {
		c.ScreenwritingGetPromptInputModel.Value = true
	}
	if e.FormValue(c.ScreenwritingGetPromptOutputModel.NamePrefix+c.ScreenwritingGetPromptOutputModel.Suffix) == "on" {
		c.ScreenwritingGetPromptOutputModel.Value = true
	}
	if e.FormValue(c.ScreenwritingOutputModel.NamePrefix+c.ScreenwritingOutputModel.Suffix) == "on" {
		c.ScreenwritingOutputModel.Value = true
	}
	if e.FormValue(c.ContainerSwapModel.NamePrefix+c.ContainerSwapModel.Suffix) == "on" {
		c.ContainerSwapModel.Value = true
	}
	if e.FormValue(c.GenerateAudioModel.NamePrefix+c.GenerateAudioModel.Suffix) == "on" {
		c.GenerateAudioModel.Value = true
	}
	if e.FormValue(c.GenerateLipsyncModel.NamePrefix+c.GenerateLipsyncModel.Suffix) == "on" {
		c.GenerateLipsyncModel.Value = true
	}
	if e.FormValue(c.GenerateThumbnailsModel.NamePrefix+c.GenerateThumbnailsModel.Suffix) == "on" {
		c.GenerateThumbnailsModel.Value = true
	}
	if e.FormValue(c.GenerateBackgroundCountextModel.NamePrefix+c.GenerateBackgroundCountextModel.Suffix) == "on" {
		c.GenerateBackgroundCountextModel.Value = true
	}
	if e.FormValue(c.GenerateBackgroundModel.NamePrefix+c.GenerateBackgroundModel.Suffix) == "on" {
		c.GenerateBackgroundModel.Value = true
	}
	if e.FormValue(c.FFMPEGLipsyncPostModel.NamePrefix+c.FFMPEGLipsyncPostModel.Suffix) == "on" {
		c.FFMPEGLipsyncPostModel.Value = true
	}
	if e.FormValue(c.FFMPEGMergeModel.NamePrefix+c.FFMPEGMergeModel.Suffix) == "on" {
		c.FFMPEGMergeModel.Value = true
	}
	if e.FormValue(c.PublishVideoModel.NamePrefix+c.PublishVideoModel.Suffix) == "on" {
		c.PublishVideoModel.Value = true
	}
	if e.FormValue(c.PublishThumbnailModel.NamePrefix+c.PublishThumbnailModel.Suffix) == "on" {
		c.PublishThumbnailModel.Value = true
	}
	if e.FormValue(c.PublishMetadataModel.NamePrefix+c.PublishMetadataModel.Suffix) == "on" {
		c.PublishMetadataModel.Value = true
	}
	return c, nil
}

func (c Steps) Truncate() map[string]interface{} {
	msi := make(map[string]interface{})
	msi["get_research_output"] = c.GetResearchOutputModel.Value
	msi["get_research_prompt"] = c.GetResearchPromptModel.Value
	msi["screenwriting_start"] = c.ScreenwritingStartModel.Value
	msi["screenwriting_get_prompt_input"] = c.ScreenwritingGetPromptInputModel.Value
	msi["screenwriting_get_prompt_output"] = c.ScreenwritingGetPromptOutputModel.Value
	msi["scrrenwriting_output"] = c.ScreenwritingOutputModel.Value
	msi["container_swap"] = c.ContainerSwapModel.Value
	msi["generate_audio"] = c.GenerateAudioModel.Value
	msi["generate_lipsync"] = c.GenerateLipsyncModel.Value
	msi["generate_thumbnails"] = c.GenerateThumbnailsModel.Value
	msi["generate_background_context"] = c.GenerateBackgroundCountextModel.Value
	msi["generate_background"] = c.GenerateBackgroundModel.Value
	msi["ffmpeg_lipsync_post"] = c.FFMPEGLipsyncPostModel.Value
	msi["ffmpeg_merge"] = c.FFMPEGMergeModel.Value
	msi["publish_video"] = c.PublishVideoModel.Value
	msi["publish_thumbnail"] = c.PublishThumbnailModel.Value
	msi["publish_metadata"] = c.PublishMetadataModel.Value
	return msi
}

func ValidateSteps(s Steps, prefix string) (Steps, error) {
	s.GetResearchOutputModel = ValidateToggle(s.GetResearchOutputModel, uuid.NewString(), prefix, "get_research_output", "get_research_output")
	s.GetResearchPromptModel = ValidateToggle(s.GetResearchPromptModel, uuid.NewString(), prefix, "get_research_prompt", "get_research_prompt")
	s.ScreenwritingStartModel = ValidateToggle(s.ScreenwritingStartModel, uuid.NewString(), prefix, "screenwriting_start", "screenwriting start")
	s.ScreenwritingGetPromptInputModel = ValidateToggle(s.ScreenwritingGetPromptInputModel, uuid.NewString(), prefix, "screenwriting_get_prompt_input", "screenwriting get prompt input")
	s.ScreenwritingGetPromptOutputModel = ValidateToggle(s.ScreenwritingGetPromptOutputModel, uuid.NewString(), prefix, "screenwriting_get_prompt_output", "screenwriting get prompt output")
	s.ScreenwritingOutputModel = ValidateToggle(s.ScreenwritingOutputModel, uuid.NewString(), prefix, "screenwriting_output", "screenwriting_output")
	s.ContainerSwapModel = ValidateToggle(s.ContainerSwapModel, uuid.NewString(), prefix, "container_swap", "container swap")
	s.GenerateAudioModel = ValidateToggle(s.GenerateAudioModel, uuid.NewString(), prefix, "generate_audio", "generate audio")
	s.GenerateLipsyncModel = ValidateToggle(s.GenerateLipsyncModel, uuid.NewString(), prefix, "generate_lipsync", "generate audio")
	s.GenerateThumbnailsModel = ValidateToggle(s.GenerateThumbnailsModel, uuid.NewString(), prefix, "generate_thumbnails", "generate thumbnails")
	s.GenerateBackgroundCountextModel = ValidateToggle(s.GenerateBackgroundCountextModel, uuid.NewString(), prefix, "generate_background_context_images", "generate background context images")
	s.GenerateBackgroundModel = ValidateToggle(s.GenerateBackgroundModel, uuid.NewString(), prefix, "generate_background", "generate background")
	s.FFMPEGLipsyncPostModel = ValidateToggle(s.FFMPEGLipsyncPostModel, uuid.NewString(), prefix, "ffmpeg_lipsync_post", "ffmpeg lipsync post")
	s.FFMPEGMergeModel = ValidateToggle(s.FFMPEGMergeModel, uuid.NewString(), prefix, "ffmpeg_merge", "ffmpeg merge")
	s.PublishVideoModel = ValidateToggle(s.PublishVideoModel, uuid.NewString(), prefix, "publish_video", "publish video")
	s.PublishThumbnailModel = ValidateToggle(s.PublishThumbnailModel, uuid.NewString(), prefix, "publish_thumbnail", "publish thumbnail")
	s.PublishMetadataModel = ValidateToggle(s.PublishMetadataModel, uuid.NewString(), prefix, "publish_metadata", "publish metadata")
	return s, nil
}


