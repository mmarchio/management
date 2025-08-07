package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type ShallowContext struct {
	ShallowModel
	PromptModel							string `json:"prompt_model"`
	DispositionModel 					string `json:"disposition_model"`
	JobRunID 							RunID `json:"job_run_id"`
	SettingsModel 						string `json:"settings_model"`
	GetResearchPromptModel 				string `json:"get_research_prompt_model"`
	GetResearchOutputModel 				string `json:"get_research_output_model"`
	ScreenwritingPromptModel 			string `json:"screenwriting_prompt_model"`
	ScreenWritingOutputModel 			string `json:"screenwriting_output_model"`
	VideoPromptModel 					string `json:"video_prompt_model"`
	AudioPromptModel 					string `json:"audio_prompt_model"`
	AudioOutputModel 					string `json:"audio_output_model"`
	VideoLipsyncOutputModel 			string `json:"video_lipsync_output_model"`
	VideoTransparencyOutputModel 		string `json:"video_transparency_output_model"`
	VideoBackgroundOutputModel 			string `json:"video_background_output_model"`
	VideoLayerMergeModel 				string `json:"video_layer_merge_model"`
	VideoJoinArrayModel 				[]string `json:"video_join_array_model"`
	ImageThumbnailPromptModel 			string `json:"image_thumbnail_prompt_model"`
	ImageThumbnailOutputModel 			string `json:"image_thumbnail_output_model"`
	ImageBackgroundContextOutputModel 	string `json:"image_background_context_output_model"`
	PublishVideoYoutubeModel 			string `json:"publish_video_youtube_model"`
	PublishVideoTiktokModel 			string `json:"publish_video_tiktok_model"`
	PublishVideoRumbleModel 			string `json:"publish_video_rumble_model"`
	PublishVideoFacebookModel 			string `json:"publish_video_facebook_model"`
	PublishSocialFacebookModel 			string `json:"publish_social_facebook_model"`
	PublishSocialXModel 				string `json:"publish_social_x_model"`
	PublishSocialYoutubeModel 			string `json:"publish_social_youtube_model"`
	PublishSocialTruthModel 			string `json:"publish_social_truth_model"`
}

func NewShallowContext(prompt Prompt, jobRunID RunID, disposition Disposition) ShallowContext {
	c := ShallowContext{}
	c.ShallowModel.ID = uuid.NewString()
	c.ID = c.ShallowModel.ID
	c.PromptModel = prompt.Model.ID
	c.DispositionModel = disposition.Model.ID
	c.JobRunID = jobRunID
	c.SettingsModel = prompt.SettingsModel.ID
	c.GetResearchPromptModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.GetResearchOutputModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.ScreenwritingPromptModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.ScreenWritingOutputModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.VideoPromptModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.AudioPromptModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.VideoLipsyncOutputModel = ShallowVideoLipsyncOutput{}.New().ShallowModel.ID
	c.VideoTransparencyOutputModel = ShallowVideoTransparencyOutput{}.New().ShallowModel.ID
	c.VideoBackgroundOutputModel = ShallowVideoBackgroundOutput{}.New().ShallowModel.ID
	c.VideoLayerMergeModel = ShallowVideoLayerMergeOutput{}.New().ShallowModel.ID
	c.ImageThumbnailPromptModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.ImageThumbnailOutputModel = ShallowImageThumbnailOutput{}.New().ShallowModel.ID
	c.ImageBackgroundContextOutputModel = ShallowImageBackgroundContextOutput{}.New().ShallowModel.ID
	c.PublishVideoYoutubeModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.PublishVideoTiktokModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.PublishVideoRumbleModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.PublishVideoFacebookModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.PublishSocialFacebookModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.PublishSocialXModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.PublishSocialYoutubeModel = ShallowStats{}.New(nil).ShallowModel.ID
	c.PublishSocialTruthModel = ShallowStats{}.New(nil).ShallowModel.ID
	return c
}

func (c ShallowContext) ToModel() (*models.ShallowContext, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{Package:"types", Struct:"ShallowContext", Function: "ToModel"}.Wrap(err)
	}
	r := models.ShallowContext{}
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, merrors.JSONUnmarshallingError{Package:"types", Struct:"ShallowContext", Function: "FromModel"}.Wrap(err) 
	}
	return &r, nil
}

func (c *ShallowContext) FromModel(ptr *models.ShallowContext) error {
	if ptr != nil {
		ctx := *ptr
		b, err := json.Marshal(ctx)
		if err != nil {
			return merrors.JSONMarshallingError{Package:"types", Struct:"ShallowContext", Function: "FromModel"}.Wrap(err)
		}
		d := ShallowContext{}
		if err := json.Unmarshal(b, &d); err != nil {
			return merrors.JSONUnmarshallingError{Package:"types", Struct:"ShallowContext", Function: "FromModel"}.Wrap(err)
		}
		c = &d
	}
	return nil
}

func (c *ShallowContext) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowContext) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func GetShallowSystemPrompts() ([]string, error) {
	ct := "shallowsystemprompt"
	ctx := context.Background()
	systemPrompt := NewShallowSystemPrompt(nil, &ct)
	systemPrompts, err := systemPrompt.List(ctx)
	if err != nil {
		return nil, merrors.ContentListError{Package: "types", Function: "GetSystemPrompts"}.Wrap(err)
	}
	list := make([]string, 0)
	for _, sp := range systemPrompts {
		list = append(list, sp.ShallowModel.ID)
	}
	return list, nil
}
