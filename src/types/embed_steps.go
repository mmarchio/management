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
	ID 								StepsID `json:"id"`
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
	c.GetResearchOutput.init()
	c.GetResearchOutput.New(*c)
	c.GetResearchPrompt.init()
	c.GetResearchPrompt.New(*c)
	c.ScreenwritingStart.init()
	c.ScreenwritingStart.New(*c)
	c.ScreenwritingGetPromptInput.init()
	c.ScreenwritingGetPromptInput.New(*c)
	c.ScreenwritingGetPromptOutput.init()
	c.ScreenwritingGetPromptOutput.New(*c)
	c.ScreenwritingOutput.init()
	c.ScreenwritingOutput.New(*c)
	c.ContainerSwap.init()
	c.ContainerSwap.New(*c)
	c.GenerateAudio.init()
	c.GenerateAudio.New(*c)
	c.GenerateLipsync.init()
	c.GenerateLipsync.New(*c)
	c.GenerateThumbnails.init()
	c.GenerateThumbnails.New(*c)
	c.GenerateBackgroundCountext.init()
	c.GenerateBackgroundCountext.New(*c)
	c.GenerateBackground.init()
	c.GenerateBackground.New(*c)
	c.FFMPEGLipsyncPost.init()
	c.FFMPEGLipsyncPost.New(*c)
	c.FFMPEGMerge.init()
	c.FFMPEGMerge.New(*c)
	c.PublishVideo.init()
	c.PublishVideo.New(*c)
	c.PublishThumbnail.init()
	c.PublishThumbnail.New(*c)
	c.PublishMetadata.init()
	c.PublishMetadata.New(*c)
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
	c.GetResearchOutput = toggle.FromModel(model.GetResearchOutput)
	c.GetResearchPrompt = toggle.FromModel(model.GetResearchPrompt)
	c.ScreenwritingStart = toggle.FromModel(model.ScreenwritingStart)
	c.ScreenwritingGetPromptInput = toggle.FromModel(model.ScreenwritingGetPromptInput)
	c.ScreenwritingGetPromptOutput = toggle.FromModel(model.ScreenwritingGetPromptOutput)
	c.ScreenwritingOutput = toggle.FromModel(model.ScreenwritingOutput)
	c.ContainerSwap = toggle.FromModel(model.ContainerSwap)
	c.GenerateAudio = toggle.FromModel(model.GenerateAudio)
	c.GenerateLipsync = toggle.FromModel(model.GenerateLipsync)
	c.GenerateThumbnails = toggle.FromModel(model.GenerateThumbnails)
	c.GenerateBackgroundCountext = toggle.FromModel(model.GenerateBackgroundCountext)
	c.GenerateBackground = toggle.FromModel(model.GenerateBackground)
	c.FFMPEGLipsyncPost = toggle.FromModel(model.FFMPEGLipsyncPost)
	c.FFMPEGMerge = toggle.FromModel(model.FFMPEGMerge)
	c.PublishVideo = toggle.FromModel(model.PublishVideo)
	c.PublishThumbnail = toggle.FromModel(model.PublishThumbnail)
	c.PublishMetadata = toggle.FromModel(model.PublishMetadata)
	return c
}

func (c Steps) Bind(e echo.Context) (Steps, error) {
	if e.FormValue(c.GetResearchOutput.NamePrefix+c.GetResearchOutput.Suffix) == "on" {
		c.GetResearchOutput.Value = true
	}
	if e.FormValue(c.GetResearchPrompt.NamePrefix+c.GetResearchPrompt.Suffix) == "on" {
		c.GetResearchPrompt.Value = true
	}
	if e.FormValue(c.ScreenwritingStart.NamePrefix+c.ScreenwritingStart.Suffix) == "on" {
		c.ScreenwritingStart.Value = true
	}
	if e.FormValue(c.ScreenwritingGetPromptInput.NamePrefix+c.ScreenwritingGetPromptInput.Suffix) == "on" {
		c.ScreenwritingGetPromptInput.Value = true
	}
	if e.FormValue(c.ScreenwritingGetPromptOutput.NamePrefix+c.ScreenwritingGetPromptOutput.Suffix) == "on" {
		c.ScreenwritingGetPromptOutput.Value = true
	}
	if e.FormValue(c.ScreenwritingOutput.NamePrefix+c.ScreenwritingOutput.Suffix) == "on" {
		c.ScreenwritingOutput.Value = true
	}
	if e.FormValue(c.ContainerSwap.NamePrefix+c.ContainerSwap.Suffix) == "on" {
		c.ContainerSwap.Value = true
	}
	if e.FormValue(c.GenerateAudio.NamePrefix+c.GenerateAudio.Suffix) == "on" {
		c.GenerateAudio.Value = true
	}
	if e.FormValue(c.GenerateLipsync.NamePrefix+c.GenerateLipsync.Suffix) == "on" {
		c.GenerateLipsync.Value = true
	}
	if e.FormValue(c.GenerateThumbnails.NamePrefix+c.GenerateThumbnails.Suffix) == "on" {
		c.GenerateThumbnails.Value = true
	}
	if e.FormValue(c.GenerateBackgroundCountext.NamePrefix+c.GenerateBackgroundCountext.Suffix) == "on" {
		c.GenerateBackgroundCountext.Value = true
	}
	if e.FormValue(c.GenerateBackground.NamePrefix+c.GenerateBackground.Suffix) == "on" {
		c.GenerateBackground.Value = true
	}
	if e.FormValue(c.FFMPEGLipsyncPost.NamePrefix+c.FFMPEGLipsyncPost.Suffix) == "on" {
		c.FFMPEGLipsyncPost.Value = true
	}
	if e.FormValue(c.FFMPEGMerge.NamePrefix+c.FFMPEGMerge.Suffix) == "on" {
		c.FFMPEGMerge.Value = true
	}
	if e.FormValue(c.PublishVideo.NamePrefix+c.PublishVideo.Suffix) == "on" {
		c.PublishVideo.Value = true
	}
	if e.FormValue(c.PublishThumbnail.NamePrefix+c.PublishThumbnail.Suffix) == "on" {
		c.PublishThumbnail.Value = true
	}
	if e.FormValue(c.PublishMetadata.NamePrefix+c.PublishMetadata.Suffix) == "on" {
		c.PublishMetadata.Value = true
	}
	return c, nil
}

func (c Steps) Truncate() map[string]interface{} {
	msi := make(map[string]interface{})
	msi["get_research_output"] = c.GetResearchOutput.Value
	msi["get_research_prompt"] = c.GetResearchPrompt.Value
	msi["screenwriting_start"] = c.ScreenwritingStart.Value
	msi["screenwriting_get_prompt_input"] = c.ScreenwritingGetPromptInput.Value
	msi["screenwriting_get_prompt_output"] = c.ScreenwritingGetPromptOutput.Value
	msi["scrrenwriting_output"] = c.ScreenwritingOutput.Value
	msi["container_swap"] = c.ContainerSwap.Value
	msi["generate_audio"] = c.GenerateAudio.Value
	msi["generate_lipsync"] = c.GenerateLipsync.Value
	msi["generate_thumbnails"] = c.GenerateThumbnails.Value
	msi["generate_background_context"] = c.GenerateBackgroundCountext.Value
	msi["generate_background"] = c.GenerateBackground.Value
	msi["ffmpeg_lipsync_post"] = c.FFMPEGLipsyncPost.Value
	msi["ffmpeg_merge"] = c.FFMPEGMerge.Value
	msi["publish_video"] = c.PublishVideo.Value
	msi["publish_thumbnail"] = c.PublishThumbnail.Value
	msi["publish_metadata"] = c.PublishMetadata.Value
	return msi
}

func ValidateSteps(s Steps, prefix string) (Steps, error) {
	s.GetResearchOutput = ValidateToggle(s.GetResearchOutput, uuid.NewString(), prefix, "get_research_output", "get_research_output")
	s.GetResearchPrompt = ValidateToggle(s.GetResearchPrompt, uuid.NewString(), prefix, "get_research_prompt", "get_research_prompt")
	s.ScreenwritingStart = ValidateToggle(s.ScreenwritingStart, uuid.NewString(), prefix, "screenwriting_start", "screenwriting start")
	s.ScreenwritingGetPromptInput = ValidateToggle(s.ScreenwritingGetPromptInput, uuid.NewString(), prefix, "screenwriting_get_prompt_input", "screenwriting get prompt input")
	s.ScreenwritingGetPromptOutput = ValidateToggle(s.ScreenwritingGetPromptOutput, uuid.NewString(), prefix, "screenwriting_get_prompt_output", "screenwriting get prompt output")
	s.ScreenwritingOutput = ValidateToggle(s.ScreenwritingOutput, uuid.NewString(), prefix, "screenwriting_output", "screenwriting_output")
	s.ContainerSwap = ValidateToggle(s.ContainerSwap, uuid.NewString(), prefix, "container_swap", "container swap")
	s.GenerateAudio = ValidateToggle(s.GenerateAudio, uuid.NewString(), prefix, "generate_audio", "generate audio")
	s.GenerateLipsync = ValidateToggle(s.GenerateLipsync, uuid.NewString(), prefix, "generate_lipsync", "generate audio")
	s.GenerateThumbnails = ValidateToggle(s.GenerateThumbnails, uuid.NewString(), prefix, "generate_thumbnails", "generate thumbnails")
	s.GenerateBackgroundCountext = ValidateToggle(s.GenerateBackgroundCountext, uuid.NewString(), prefix, "generate_background_context_images", "generate background context images")
	s.GenerateBackground = ValidateToggle(s.GenerateBackground, uuid.NewString(), prefix, "generate_background", "generate background")
	s.FFMPEGLipsyncPost = ValidateToggle(s.FFMPEGLipsyncPost, uuid.NewString(), prefix, "ffmpeg_lipsync_post", "ffmpeg lipsync post")
	s.FFMPEGMerge = ValidateToggle(s.FFMPEGMerge, uuid.NewString(), prefix, "ffmpeg_merge", "ffmpeg merge")
	s.PublishVideo = ValidateToggle(s.PublishVideo, uuid.NewString(), prefix, "publish_video", "publish video")
	s.PublishThumbnail = ValidateToggle(s.PublishThumbnail, uuid.NewString(), prefix, "publish_thumbnail", "publish thumbnail")
	s.PublishMetadata = ValidateToggle(s.PublishMetadata, uuid.NewString(), prefix, "publish_metadata", "publish metadata")
	return s, nil
}


