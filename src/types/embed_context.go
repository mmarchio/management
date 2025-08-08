package types

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type Context struct {
	EmbedModel
	PromptModel							Prompt `json:"prompt_model"`
	DispositionModel 					Disposition `json:"disposition_model"`
	JobRunID 							RunID `json:"job_run_id"`
	SettingsModel 						Settings `json:"settings_model"`
	GetResearchPromptModel 				Stats `json:"get_research_prompt_model"`
	GetResearchOutputModel 				Stats `json:"get_research_output_model"`
	ScreenwritingPromptModel 			Stats `json:"screenwriting_prompt_model"`
	ScreenWritingOutputModel 			Stats `json:"screenwriting_output_model"`
	VideoPromptModel 					Stats `json:"video_prompt_model"`
	AudioPromptModel 					Stats `json:"audio_prompt_model"`
	AudioOutputModel 					AudioOutput `json:"audio_output_model"`
	VideoLipsyncOutputModel 			VideoLipsyncOutput `json:"video_lipsync_output_model"`
	VideoTransparencyOutputModel 		VideoTransparencyOutput `json:"video_transparency_output_model"`
	VideoBackgroundOutputModel 			VideoBackgroundOutput `json:"video_background_output_model"`
	VideoLayerMergeModel 				VideoLayerMergeOutput `json:"video_layer_merge_model"`
	VideoJoinArrayModel 				[]Scene `json:"video_join_array_model"`
	ImageThumbnailPromptModel 			Stats `json:"image_thumbnail_prompt_model"`
	ImageThumbnailOutputModel 			ImageThumbnailOutput `json:"image_thumbnail_output_model"`
	ImageBackgroundContextOutputModel 	ImageBackgroundContextOutput `json:"image_background_context_output_model"`
	PublishVideoYoutubeModel 			Stats `json:"publish_video_youtube_model"`
	PublishVideoTiktokModel 			Stats `json:"publish_video_tiktok_model"`
	PublishVideoRumbleModel 			Stats `json:"publish_video_rumble_model"`
	PublishVideoFacebookModel 			Stats `json:"publish_video_facebook_model"`
	PublishSocialFacebookModel 			Stats `json:"publish_social_facebook_model"`
	PublishSocialXModel 				Stats `json:"publish_social_x_model"`
	PublishSocialYoutubeModel 			Stats `json:"publish_social_youtube_model"`
	PublishSocialTruthModel 			Stats `json:"publish_social_truth_model"`
}

func NewContext(prompt Prompt, jobRunID RunID, disposition Disposition) Context {
	c := Context{}
	c.EmbedModel.ID = uuid.NewString()
	c.ID = c.EmbedModel.ID
	c.PromptModel = prompt
	c.DispositionModel = disposition
	c.JobRunID = jobRunID
	c.SettingsModel = prompt.SettingsModel
	c.GetResearchPromptModel = c.GetResearchPromptModel.New(nil)
	c.GetResearchOutputModel = c.GetResearchOutputModel.New(nil)
	c.ScreenwritingPromptModel = c.ScreenwritingPromptModel.New(nil)
	c.ScreenWritingOutputModel = c.ScreenWritingOutputModel.New(nil)
	c.VideoPromptModel = c.VideoPromptModel.New(nil)
	c.AudioPromptModel = c.AudioPromptModel.New(nil)
	c.VideoLipsyncOutputModel = c.VideoLipsyncOutputModel.New()
	c.VideoTransparencyOutputModel = c.VideoTransparencyOutputModel.New()
	c.VideoBackgroundOutputModel = c.VideoBackgroundOutputModel.New()
	c.VideoLayerMergeModel = c.VideoLayerMergeModel.New()
	c.ImageThumbnailPromptModel = c.ImageThumbnailPromptModel.New(nil)
	c.ImageThumbnailOutputModel = c.ImageThumbnailOutputModel.New()
	c.ImageBackgroundContextOutputModel = c.ImageBackgroundContextOutputModel.New()
	c.PublishVideoYoutubeModel = c.PublishVideoYoutubeModel.New(nil)
	c.PublishVideoTiktokModel = c.PublishVideoTiktokModel.New(nil)
	c.PublishVideoRumbleModel = c.PublishVideoRumbleModel.New(nil)
	c.PublishVideoFacebookModel = c.PublishVideoFacebookModel.New(nil)
	c.PublishSocialFacebookModel = c.PublishSocialFacebookModel.New(nil)
	c.PublishSocialXModel = c.PublishSocialXModel.New(nil)
	c.PublishSocialYoutubeModel = c.PublishSocialYoutubeModel.New(nil)
	c.PublishSocialTruthModel = c.PublishSocialTruthModel.New(nil)
	return c
}

func (c *Context) GetCtx(ctx context.Context) error {
	systemContext, err := models.Context{}.GetCtx(ctx)
	if err != nil {
		return merrors.ContextGetError{Package: "types", Struct: "Context", Function: "GetCtx"}.Wrap(err)
	}
	c.FromModel(systemContext)
	return nil
}

func (c Context) SetCtx(ctx context.Context) (context.Context, error) {
	s, err := c.ToModel()
	if err != nil {
		return ctx, merrors.SetContextError{Package:"types", Struct:"Context", Function: "SetCtx"}.Wrap(err)
	}
	ctx = s.SetCtx(ctx)
	return ctx, nil
}

func (c Context) ToModel() (*models.Context, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{Package:"types", Struct:"Context", Function: "ToModel"}.Wrap(err)
	}
	r := models.Context{}
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, merrors.JSONUnmarshallingError{Package:"types", Struct:"Context", Function: "FromModel"}.Wrap(err) 
	}
	return &r, nil
}

func (c *Context) FromModel(ptr *models.Context) error {
	if ptr != nil {
		ctx := *ptr
		b, err := json.Marshal(ctx)
		if err != nil {
			return merrors.JSONMarshallingError{Package:"types", Struct:"Context", Function: "FromModel"}.Wrap(err)
		}
		d := Context{}
		if err := json.Unmarshal(b, &d); err != nil {
			return merrors.JSONUnmarshallingError{Package:"types", Struct:"Context", Function: "FromModel"}.Wrap(err)
		}
		c = &d
	}
	return nil
}

func (c *Context) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Context) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func GetSystemPrompts() ([]SystemPrompt, error) {
	ctx := context.Background()
	systemPrompt := NewSystemPrompt(nil)
	systemPrompts, err := systemPrompt.List(ctx)
	if err != nil {
		return nil, merrors.ContentListError{Package: "types", Function: "GetSystemPrompts"}.Wrap(err)
	}
	return systemPrompts, nil
}

func (c Context) Truncate() (*TruncatedContext, error) {
	var err error
	truncated := TruncatedContext{}
	truncated.JobRunID = c.JobRunID.String()
	truncated.PromptID = c.PromptModel.Model.ID
	truncated.Prompt = c.PromptModel.Prompt
	truncated.Domain = c.PromptModel.Domain
	truncated.Category = c.PromptModel.Category
	gb := c.PromptModel.SettingsModel.GlobalBypassModel.Truncate()
	if v, ok := gb["get_research_prompt"].(bool); ok {
		truncated.GlobalBypass.GetResearchPrompt = v
	}
	if v, ok := gb["get_research_output"].(bool); ok {
		truncated.GlobalBypass.GetResearchOutput = v
	}
	if v, ok := gb["container_swap"].(bool); ok {
		truncated.GlobalBypass.ContainerSwap = v
	}
	if v, ok := gb["ffmpeg_lipsync_post"].(bool); ok {
		truncated.GlobalBypass.FfmpegLipsyncPost = v
	}
	if v, ok := gb["ffmpeg_merge"].(bool); ok {
		truncated.GlobalBypass.FfmpegMerge = v
	}
	if v, ok := gb["generate_audio"].(bool); ok {
		truncated.GlobalBypass.GenerateAudio = v
	}
	if v, ok := gb["generate_background"].(bool); ok {
		truncated.GlobalBypass.GenerateBackground = v
	}
	if v, ok := gb["generate_background_context"].(bool); ok {
		truncated.GlobalBypass.GenerateBackgroundContext = v
	}
	if v, ok := gb["generate_lipsync"].(bool); ok {
		truncated.GlobalBypass.GenerateLipsync = v
	}
	if v, ok := gb["generate_thumbnails"].(bool); ok {
		truncated.GlobalBypass.GenerateThumbnails = v
	}
	if v, ok := gb["publish_metadata"].(bool); ok {
		truncated.GlobalBypass.PublishMetadata = v
	}
	if v, ok := gb["publish_thumbnail"].(bool); ok {
		truncated.GlobalBypass.PublishThumbnail = v
	}
	if v, ok := gb["publish_video"].(bool); ok {
		truncated.GlobalBypass.PublishVideo = v
	}
	if v, ok := gb["screenwriting_get_prompt_input"].(bool); ok {
		truncated.GlobalBypass.ScreenwritingGetPromptInput = v
	}
	if v, ok := gb["screenwriting_get_prompt_output"].(bool); ok {
		truncated.GlobalBypass.ScreenwritingGetPromptOutput = v
	}
	if v, ok := gb["screenwriting_output"].(bool); ok {
		truncated.GlobalBypass.ScreenwritingOutput = v
	}
	if v, ok := gb["screenwriting_start"].(bool); ok {
		truncated.GlobalBypass.ScreenwritingStart = v
	}
	en := c.DispositionModel.EntitlementsModel
	truncated.Entitlements.Youtube = en.YouTubeModel.Value
	truncated.Entitlements.Tiktok = en.TikTokModel.Value
	truncated.Entitlements.Rumble = en.RumbleModel.Value
	truncated.Entitlements.Patreon = en.PatreonModel.Value
	truncated.Entitlements.Facebook = en.FacebookModel.Value

	truncated.Disposition.MinDuration = c.DispositionModel.MinDuration
	truncated.Disposition.MaxDuration = c.DispositionModel.MaxDuration
	truncated.Disposition.AdvertisementDuration = c.DispositionModel.AdvertisementDuration
	truncated.Settings.ID = c.PromptModel.SettingsModel.ID
	truncated.Settings.Name = c.PromptModel.SettingsModel.Name
	truncated.Settings.Recurring = c.PromptModel.SettingsModel.RecurringModel.Value
	truncated.Settings.Interval = c.PromptModel.SettingsModel.Interval

	stats := make(map[string]Stats)
	stats["get_research_prompt"] = c.GetResearchPromptModel
	stats["get_research_output"] = c.GetResearchOutputModel
	stats["screenwriting_prompt"] = c.ScreenwritingPromptModel
	stats["screenwriting_output"] = c.ScreenWritingOutputModel
	stats["video_prompt"] = c.VideoPromptModel
	stats["audio_prompt"] = c.AudioPromptModel
	stats["video_lipsync_output"] = c.VideoLipsyncOutputModel.StatsModel
	stats["video_transparency_output"] = c.VideoTransparencyOutputModel.StatsModel
	stats["video_background_output"] = c.VideoBackgroundOutputModel.StatsModel
	stats["video_merge"] = c.VideoLayerMergeModel.StatsModel
	stats["image_thumbnail_prompt"] = c.ImageThumbnailPromptModel
	stats["image_thumbnail_output"] = c.ImageThumbnailOutputModel.StatsModel
	stats["image_background_context_output"] = c.ImageBackgroundContextOutputModel.StatsModel
	stats["publish_video_youtube"] = c.PublishVideoYoutubeModel
	stats["publish_video_tiktok"] = c.PublishVideoTiktokModel
	stats["publish_video_rumble"] = c.PublishVideoRumbleModel
	stats["publish_video_facebook"] = c.PublishVideoFacebookModel
	stats["publish_social_facebook"] = c.PublishSocialFacebookModel
	stats["publish_social_x"] = c.PublishSocialXModel
	stats["publish_social_youtube"] = c.PublishSocialYoutubeModel
	stats["publish_social_truth"] = c.PublishSocialTruthModel
	sysPrompts, err := GetSystemPrompts()
	if err != nil {
		return nil, err
	}
	systemPrompts := make(map[string]string)
	for _, sp := range sysPrompts {
		systemPrompts[fmt.Sprintf("%s:%s", sp.Name, sp.Domain)] = sp.Prompt
	}
	truncated.SystemPrompts = systemPrompts
	truncated.Stats = stats
		
	return &truncated, nil
}

type TruncatedContext struct {
	ID 				string `json:"id"`
	Prompt 			string `json:"prompt"`
	Domain 			string `json:"domain"`
	Category 		string `json:"category"`
	PromptID 		string `json:"prompt_id"`
	JobID 			string `json:"job_id"`
	JobRunID 		string `json:"job_run_id"`
	GlobalBypass 	truncatedSteps `json:"global_bypass"`
	Settings 		truncatedSettings `json:"settings"`
	Disposition 	truncatedDisposition `json:"disposition"`
	Verification 	truncatedSteps `json:"verification"`
	Bypass 			truncatedSteps `json:"bypass"`
	Entitlements 	truncatedEntitlements `json:"entitlements"`
	SystemPrompts 	map[string]string `json:"system_prompts"`
	Stats 			map[string]Stats `json:"stats"`
}

type truncatedSteps struct {
	GetResearchOutput				bool `json:"get_research_output"`
	GetResearchPrompt				bool `json:"get_research_prompt"`
	ScreenwritingStart 				bool `json:"screenwriting_start"`
	ScreenwritingGetPromptInput 	bool `json:"screenwriting_get_prompt_input"`
	ScreenwritingGetPromptOutput 	bool `json:"screenwriting_get_prompt_output"`
	ScreenwritingOutput 			bool `json:"screenwriting_output"`
	ContainerSwap 					bool `json:"container_swap"`
	GenerateAudio 					bool `json:"generate_audio"`
	GenerateLipsync 				bool `json:"generate_lipsync"`
	GenerateThumbnails 				bool `json:"generate_thumbnails"`
	GenerateBackgroundContext 		bool `json:"generate_background_context"`
	GenerateBackground 				bool `json:"generate_background"`
	FfmpegLipsyncPost 				bool `json:"ffmpeg_lipsync_post"`
	FfmpegMerge 					bool `json:"ffmpeg_merge"`
	PublishVideo 					bool `json:"publish_video"`
	PublishThumbnail 				bool `json:"publish_thumbnail"`
	PublishMetadata 				bool `json:"publish_metadata"`
}

type truncatedEntitlements struct {
	Youtube 	bool `json:"youtube"`
	Tiktok		bool `json:"tiktok"`
	Rumble 		bool `json:"rumble"`
	Patreon 	bool `json:"patreon"`
	Facebook 	bool `json:"facebook"`
}

type truncatedSettings struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Recurring bool `json:"recurring"`
	Interval int64 `json:"interval"`
}

type truncatedDisposition struct {
	MinDuration 			int64 `json:"min_duration"`
	MaxDuration 			int64 `json:"max_duration"`
	AdvertisementDuration 	int64 `json:"advertisement_duration"`
}