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
	Prompt Prompt `json:"prompt"`
	Disposition Disposition `json:"disposition"`
	JobRunID RunID `json:"job_run_id"`
	Settings Settings `json:"settings"`
	GetResearchPrompt Stats `json:"get_research_prompt"`
	GetResearchOutput Stats `json:"get_research_output"`
	ScreenwritingPrompt Stats `json:"screenwriting_prompt"`
	ScreenWritingOutput Stats `json:"screenwriting_output"`
	VideoPrompt Stats `json:"video_prompt"`
	AudioPrompt Stats `json:"audio_prompt"`
	AudioOutput AudioOutput `json:"audio_output"`
	VideoLipsyncOutput VideoLipsyncOutput `json:"video_lipsync_output"`
	VideoTransparencyOutput VideoTransparancyOutput `json:"video_transparency_output"`
	VideoBackgroundOutput VideoBackgroundOutput `json:"video_background_output"`
	VideoLayerMerge VideoLayerMergeOutput `json:"video_layer_merge"`
	VideoJoin []Scene `json:"video_join"`
	ImageThumbnailPrompt Stats `json:"image_thumbnail_prompt"`
	ImageThumbnailOutput ImageThumbnailOutput `json:"image_thumbnail_output"`
	ImageBackgroundContextOutput ImageBackgroundContextOutput `json:"image_background_context_output"`
	PublishVideoYoutube Stats `json:"publish_video_youtube"`
	PublishVideoTiktok Stats `json:"publish_video_tiktok"`
	PublishVideoRumble Stats `json:"publish_video_rumble"`
	PublishVideoFacebook Stats `json:"publish_video_facebook"`
	PublishSocialFacebook Stats `json:"publish_social_facebook"`
	PublishSocialX Stats `json:"publish_social_x"`
	PublishSocialYoutube Stats `json:"publish_social_youtube"`
	PublishSocialTruth Stats `json:"publish_social_truth"`
}

func NewContext(prompt Prompt, jobRunID RunID, disposition Disposition) Context {
	c := Context{}
	c.EmbedModel.ID = uuid.NewString()
	c.ID = c.EmbedModel.ID
	c.Prompt = prompt
	c.Disposition = disposition
	c.JobRunID = jobRunID
	c.Settings = prompt.Settings
	c.GetResearchPrompt = c.GetResearchPrompt.New("")
	c.GetResearchOutput = c.GetResearchOutput.New("")
	c.ScreenwritingPrompt = c.ScreenwritingPrompt.New("")
	c.ScreenWritingOutput = c.ScreenWritingOutput.New("")
	c.VideoPrompt = c.VideoPrompt.New("")
	c.AudioPrompt = c.AudioPrompt.New("")
	c.VideoLipsyncOutput = c.VideoLipsyncOutput.New()
	c.VideoTransparencyOutput = c.VideoTransparencyOutput.New()
	c.VideoBackgroundOutput = c.VideoBackgroundOutput.New()
	c.VideoLayerMerge = c.VideoLayerMerge.New()
	c.ImageThumbnailPrompt = c.ImageThumbnailPrompt.New("")
	c.ImageThumbnailOutput = c.ImageThumbnailOutput.New()
	c.ImageBackgroundContextOutput = c.ImageBackgroundContextOutput.New()
	c.PublishVideoYoutube = c.PublishVideoYoutube.New("")
	c.PublishVideoTiktok = c.PublishVideoTiktok.New("")
	c.PublishVideoRumble = c.PublishVideoRumble.New("")
	c.PublishVideoFacebook = c.PublishVideoFacebook.New("")
	c.PublishSocialFacebook = c.PublishSocialFacebook.New("")
	c.PublishSocialX = c.PublishSocialX.New("")
	c.PublishSocialYoutube = c.PublishSocialYoutube.New("")
	c.PublishSocialTruth = c.PublishSocialTruth.New("")
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
		return nil, merrors.SystemPromptListError{Package: "types", Function: "GetSystemPrompts"}.Wrap(err)
	}
	return systemPrompts, nil
}

func (c Context) Truncate() (*TruncatedContext, error) {
	var err error
	truncated := TruncatedContext{}
	truncated.JobRunID = c.JobRunID.String()
	truncated.PromptID = c.Prompt.Model.ID
	truncated.Prompt = c.Prompt.Prompt
	truncated.Domain = c.Prompt.Domain
	truncated.Category = c.Prompt.Category
	gb := c.Prompt.Settings.GlobalBypass.Truncate()
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
	en := c.Disposition.Entitlements
	truncated.Entitlements.Youtube = en.YouTube.Value
	truncated.Entitlements.Tiktok = en.TikTok.Value
	truncated.Entitlements.Rumble = en.Rumble.Value
	truncated.Entitlements.Patreon = en.Patreon.Value
	truncated.Entitlements.Facebook = en.Facebook.Value

	truncated.Disposition.MinDuration = c.Disposition.MinDuration
	truncated.Disposition.MaxDuration = c.Disposition.MaxDuration
	truncated.Disposition.AdvertisementDuration = c.Disposition.AdvertisementDuration
	truncated.Settings.ID = c.Prompt.Settings.ID
	truncated.Settings.Name = c.Prompt.Settings.Name
	truncated.Settings.Recurring = c.Prompt.Settings.Recurring.Value
	truncated.Settings.Interval = c.Prompt.Settings.Interval

	stats := make(map[string]Stats)
	stats["get_research_prompt"] = c.GetResearchPrompt
	stats["get_research_output"] = c.GetResearchOutput
	stats["screenwriting_prompt"] = c.ScreenwritingPrompt
	stats["screenwriting_output"] = c.ScreenWritingOutput
	stats["video_prompt"] = c.VideoPrompt
	stats["audio_prompt"] = c.AudioPrompt
	stats["video_lipsync_output"] = c.VideoLipsyncOutput.Stats
	stats["video_transparency_output"] = c.VideoTransparencyOutput.Stats
	stats["video_background_output"] = c.VideoBackgroundOutput.Stats
	stats["video_merge"] = c.VideoLayerMerge.Stats
	stats["image_thumbnail_prompt"] = c.ImageThumbnailPrompt
	stats["image_thumbnail_output"] = c.ImageThumbnailOutput.Stats
	stats["image_background_context_output"] = c.ImageBackgroundContextOutput.Stats
	stats["publish_video_youtube"] = c.PublishVideoYoutube
	stats["publish_video_tiktok"] = c.PublishVideoTiktok
	stats["publish_video_rumble"] = c.PublishVideoRumble
	stats["publish_video_facebook"] = c.PublishVideoFacebook
	stats["publish_social_facebook"] = c.PublishSocialFacebook
	stats["publish_social_x"] = c.PublishSocialX
	stats["publish_social_youtube"] = c.PublishSocialYoutube
	stats["publish_social_truth"] = c.PublishSocialTruth
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