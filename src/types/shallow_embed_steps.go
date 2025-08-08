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

func (c ShallowSteps) Expand(ctx context.Context) (*Steps, error) {
	f := func(ctx context.Context, c string, r *Toggle) error {
		t := ShallowToggle{}
		t.ShallowModel.ID = c
		tog, err := t.Expand(ctx)
		if err != nil {
			return merrors.ContentGetError{}.Wrap(err)
		}
		if tog != nil {
			r = tog
		}
		return nil
	}
	r := Steps{}
	r.EmbedModel = r.EmbedModel.FromShallowModel(c.ShallowModel)
	if err := f(ctx, c.GetResearchOutputModel, &r.GetResearchOutputModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.GetResearchPromptModel, &r.GetResearchPromptModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.ScreenwritingStartModel, &r.ScreenwritingStartModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.ScreenwritingGetPromptInputModel, &r.ScreenwritingGetPromptInputModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.ScreenwritingGetPromptOutputModel, &r.ScreenwritingGetPromptOutputModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.ScreenwritingOutputModel, &r.ScreenwritingOutputModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.ContainerSwapModel, &r.ContainerSwapModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.GenerateAudioModel, &r.GenerateAudioModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.GenerateLipsyncModel, &r.GenerateLipsyncModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.GenerateThumbnailsModel, &r.GenerateThumbnailsModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.GenerateBackgroundContextModel, &r.GenerateBackgroundContextModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.GenerateBackgroundModel, &r.GenerateBackgroundModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.FFMPEGLipsyncPostModel, &r.FFMPEGLipsyncPostModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.FFMPEGMergeModel, &r.FFMPEGMergeModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.PublishVideoModel, &r.PublishVideoModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.PublishThumbnailModel, &r.PublishThumbnailModel); err != nil {
		return nil, err
	}
	if err := f(ctx, c.PublishMetadataModel, &r.PublishMetadataModel); err != nil {
		return nil, err
	}
	return &r, nil
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
	c.GetResearchPromptModel = model.GetResearchPromptModel
	c.ScreenwritingStartModel = model.ScreenwritingStartModel
	c.ScreenwritingGetPromptInputModel = model.ScreenwritingGetPromptInputModel
	c.ScreenwritingGetPromptOutputModel = model.ScreenwritingGetPromptOutputModel
	c.ScreenwritingOutputModel = model.ScreenwritingOutputModel
	c.ContainerSwapModel = model.ContainerSwapModel
	c.GenerateAudioModel = model.GenerateAudioModel
	c.GenerateLipsyncModel = model.GenerateLipsyncModel
	c.GenerateThumbnailsModel = model.GenerateThumbnailsModel
	c.GenerateBackgroundContextModel = model.GenerateBackgroundContextModel
	c.GenerateBackgroundModel = model.GenerateBackgroundModel
	c.FFMPEGLipsyncPostModel = model.FFMPEGLipsyncPostModel
	c.FFMPEGMergeModel = model.FFMPEGMergeModel
	c.PublishVideoModel = model.PublishVideoModel
	c.PublishThumbnailModel = model.PublishThumbnailModel
	c.PublishMetadataModel = model.PublishMetadataModel
	return c
}

func (c ShallowSteps) IsShallowModel() bool {
	return true
}
