package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mmarchio/management/database"
	merrors "github.com/mmarchio/management/errors"
)

type ContextKeyT int64
var contextKey ContextKeyT = 3


type Context struct {
	Model
	PromptModel 						Prompt `json:"prompt_model"`
	DispositionModel 					Disposition `json:"disposition_model"`
	JobRunID 							string `json:"job_run_id"`
	SettingsModel 						Settings `json:"settings_model"`
	GetResearchPromptModel 				Stats `json:"get_research_prompt_model"`
	GetResearchOutputModel 				Stats `json:"get_research_output_model"`
	ScreenwritingPromptModel 			Stats `json:"screenwriting_prompt_model"`
	ScreenWritingOutputModel 			Stats `json:"screenwriting_output_model"`
	VideoPromptModel 					Stats `json:"video_prompt_model"`
	AudioPromptModel 					Stats `json:"audio_prompt_model"`
	AudioOutputModel 					AudioOutputModel `json:"audio_output_model"`
	VideoLipsyncOutputModel 			VideoLipsyncOutputModel `json:"video_lipsync_output_model"`
	VideoTransparencyOutputModel 		VideoTransparancyOutput `json:"video_transparency_output_model"`
	VideoBackgroundOutputModel 			VideoBackgroundOutputModel `json:"video_background_output_model"`
	VideoLayerMergeModel 				VideoLayerMergeModelOutput `json:"video_layer_merge_model"`
	VideoJoinArrayModel 				[]Scene `json:"video_join_array_model"`
	ImageThumbnailPromptModel 			Stats `json:"image_thumbnail_prompt_model"`
	ImageThumbnailOutputModel 			ImageThumbnailOutputModel `json:"image_thumbnail_output_model"`
	ImageBackgroundContextOutputModel 	ImageBackgroundContextOutputModel `json:"image_background_context_output_model"`
	PublishVideoYoutubeModel 			Stats `json:"publish_video_youtube_model"`
	PublishVideoTiktokModel 			Stats `json:"publish_video_tiktok_model"`
	PublishVideoRumbleModel 			Stats `json:"publish_video_rumble_model"`
	PublishVideoFacebookModel 			Stats `json:"publish_video_facebook_model"`
	PublishSocialFacebookModel 			Stats `json:"publish_social_facebook_model"`
	PublishSocialXModel 				Stats `json:"publish_social_x_model"`
	PublishSocialYoutubeModel 			Stats `json:"publish_social_youtube_model"`
	PublishSocialTruthModel 			Stats `json:"publish_social_truth_model"`
}

type ShallowContext struct {
	ShallowModel
	PromptModel 						string `json:"prompt_model"`
	DispositionModel 					string `json:"disposition_model"`
	JobRunID 							string `json:"job_run_id"`
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

func (c ShallowContext) Set(e echo.Context) error {
	return nil
}

func (c ShallowContext) SetPromptModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PromptModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowContext) SetDispositionModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.DispositionModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetSettingsModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.SettingsModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetGetResearchPromptModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.GetResearchPromptModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetGetResearchOutputModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.GetResearchOutputModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetJobRunID(e echo.Context, id string) error {
	c.JobRunID = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowContext) SetScreenwritingOutput(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.ScreenWritingOutputModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetScreenwritingPromptModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.ScreenwritingPromptModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetVideoPromptModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.VideoPromptModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetAudioPromptModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.AudioPromptModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetAudioOutputModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.AudioOutputModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetVideoLipsyncOutputModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.VideoLipsyncOutputModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetVideoTransparencyOutputModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.VideoTransparencyOutputModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetVideoBackgroundOutputModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.VideoBackgroundOutputModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetVideoLayerMergeModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.VideoLayerMergeModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetVideoJoinArrayModel(e echo.Context, ids []string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, ids...)
	c.VideoJoinArrayModel = ids
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) AppendVideoJoinArrayModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.VideoJoinArrayModel = append(c.VideoJoinArrayModel, id)
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetImageThumbnailPromptModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.ImageThumbnailPromptModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetImageThumbnailOutputModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.ImageThumbnailOutputModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetImageBackgroundContextOutputModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.ImageBackgroundContextOutputModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishVideoYoutubeModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishVideoYoutubeModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishVideoTiktokModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishVideoTiktokModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishVideoRumbleModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishVideoRumbleModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishVideoFacebookModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishVideoFacebookModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishSocialFacebookModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishSocialFacebookModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishSocialTruthModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishSocialTruthModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishSocialXModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishSocialXModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishSocialYoutubeModel(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishSocialYoutubeModel = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) Get(e echo.Context, mode string) (*Context, *ShallowContext, error) {
	content := Content{ID: c.ShallowModel.ID}
	if err := content.Get(e); err != nil {
		return nil, nil, merrors.ContentGetError{Info: c.ShallowModel.ID, Package: "models", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(nil, err)
	}
	if err := json.Unmarshal([]byte(content.Content), &c); err != nil {
		return nil, nil, merrors.JSONUnmarshallingError{Info: content.Content, Package: "models", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(nil, err)
	}
	if mode == "shallow" {
		return nil, &c, nil
	}
	if mode == "full" {
		m := Context{}
		m.Model.ID = c.ShallowModel.ID
		m.Model.CreatedAt = c.ShallowModel.CreatedAt
		m.Model.UpdatedAt = c.ShallowModel.UpdatedAt
		m.Model.ContentType = c.ShallowModel.ContentType

		promptptr, _, err := NewShallowPrompt(&c.PromptModel).Get(e, "full")
		if err != nil {
			return nil, nil, err
		}
		if promptptr != nil {
			m.PromptModel = *promptptr
		}
		dispositionptr, _, err := NewShallowDisposition(&c.DispositionModel).Get(e, "full")
		if err != nil {
			return nil, nil, err
		}
		if dispositionptr != nil {
			m.DispositionModel = *dispositionptr
		}

	}
	return nil, nil, merrors.ContentGetError{Package: "models", Struct: "ShallowWorkflow", Function: "Get"}.New(nil, "unknown mode: %s", mode)
}

type Stats struct {
	Model
	ID 			string `json:"stats_id"`
	Start 		time.Time `json:"start"`
	End 		time.Time `json:"end"`
	Input 		string `json:"input"`
	Output 		string `json:"output"`
	Duration 	time.Duration `json:"duration"`
	FilesArrayModel	[]File `json:"files_array_model"`
	Status		string `json:"status"`
}

type ShallowStats struct {
	ShallowModel
	ID 				string `json:"stats_id"`
	Start 			time.Time `json:"start"`
	End 			time.Time `json:"end"`
	Input 			string `json:"input"`
	Output 			string `json:"output"`
	Duration 		time.Duration `json:"duration"`
	FilesArrayModel	[]string `json:"files_array_model"`
	Status			string `json:"status"`
}

func (c ShallowStats) Set(e echo.Context) error {
	return nil
}

func (c ShallowStats) SetStart(e echo.Context, id time.Time) error {
	c.Start = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) SetEnd(e echo.Context, id time.Time) error {
	c.End = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) SetInput(e echo.Context, id string) error {
	c.Input = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) SetOutput(e echo.Context, id string) error {
	c.Output = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) SetDuration(e echo.Context, id time.Duration) error {
	c.Duration = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) SetFiles(e echo.Context, ids []string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, ids...)
	c.FilesArrayModel = ids
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) AppendFiles(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.FilesArrayModel = append(c.FilesArrayModel, id)
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) SetStatus(e echo.Context, id string) error {
	c.Status = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(e); err != nil {
			return err
		}
	}
	return nil	
}

func NewShallowStats(id *string) ShallowStats {
	c := ShallowStats{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_stats"
	return c
}

type File struct {
	Model
	ID string `json:"file_id"`
	Type string `json:"type"`
	Path string `json:"path"`
	Duration time.Duration `json:"duration"`
	Scene string `json:"scene"`
	Joined bool `json:"joined"`
}

type ShallowFile struct {
	Model
	ID string `json:"file_id"`
	Type string `json:"type"`
	Path string `json:"path"`
	Duration time.Duration `json:"duration"`
	Scene string `json:"scene"`
	Joined bool `json:"joined"`
}

func NewShallowFile(id *string) ShallowFile {
	c := ShallowFile{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_file"
	return c
}

type Output struct {
	Model
	ID string `json:"id"`
	FilesModel	[]File `json:"files_array_model"`
	StatusModel	string `json:"status_model"`
}

type ShallowOutput struct {
	ShallowModel
	ID string `json:"id"`
	FilesModel	[]File `json:"files_array_model"`
	StatusModel	string `json:"status_model"`
}

func NewShallowOutput(id *string) ShallowOutput {
	c := ShallowOutput{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_output"
	return c
}

type Scene struct {
	Model
	ID string `json:"id"`
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
	SceneNumber int64 `json:"scene_number"`
	Path string `json:"path"`
	FilesArrayModel []File `json:"files_array_model"`
	SceneFileModel File `json:"scene_file_model"`
}

type ShallowScene struct {
	ShallowModel
	ID string `json:"id"`
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
	SceneNumber int64 `json:"scene_number"`
	Path string `json:"path"`
	FilesArrayModel []string `json:"files_array_model"`
	SceneFile string `json:"scene_file_model"`
}

func NewShallowScene(id *string) ShallowScene {
	c := ShallowScene{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_scent"
	return c
}

func NewShallowContext(id *string) ShallowContext {
	c := ShallowContext{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_context"
	return c
}

func NewContext(id *string) Context {
	c := Context{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "context"
	return c
}

type AudioOutputModel Output
type ShallowAudioOutputModel ShallowOutput

func NewShallowAudioOutputModel(id *string) ShallowAudioOutputModel {
	c := ShallowAudioOutputModel{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_audiooutput"
	return c
}


type VideoLipsyncOutputModel Output
type ShallowVideoLipsyncOutputModel ShallowOutput

func NewShallowVideoLipsyncOutputModel(id *string) ShallowVideoLipsyncOutputModel {
	c := ShallowVideoLipsyncOutputModel{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_videolipsyncoutput"
	return c
}

type VideoTransparancyOutput Output
type ShallowVideoTransparancyOutput ShallowOutput

func NewShallowVideoTransparencyOutputModel(id *string) ShallowVideoTransparancyOutput {
	c := ShallowVideoTransparancyOutput{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_transparencyoutput"
	return c
}

type VideoBackgroundOutputModel Output
type ShallowVideoBackgroundOutputModel ShallowOutput

func NewShallowVideoBackgroundOutputModel(id *string) ShallowVideoBackgroundOutputModel {
	c := ShallowVideoBackgroundOutputModel{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_videobackgroundoutput"
	return c
}

type VideoLayerMergeModelOutput Output
type ShallowVideoLayerMergeModelOutput ShallowOutput

func NewShallowVideoLayerMergeModelOutput(id *string) ShallowVideoLayerMergeModelOutput {
	c := ShallowVideoLayerMergeModelOutput{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_videolayermergeoutput"
	return c
}

type ImageThumbnailOutputModel Output
type ShallowImageThumbnailOutputModel ShallowOutput

func NewShallowImageThumbnailOutputModel(id *string) ShallowImageThumbnailOutputModel {
	c := ShallowImageThumbnailOutputModel{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_imagethumbnailoutput"
	return c
}

type ImageBackgroundContextOutputModel Output
type ShallowImageBackgroundContextOutputModel ShallowOutput

func NewShallowImageBackgroundContextOutputModel(id *string) ShallowImageBackgroundContextOutputModel {
	c := ShallowImageBackgroundContextOutputModel{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_imagebackgroundcontextoutput"
	return c
}

func (c *Context) Unmarshal(e echo.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Context) Marshal(e echo.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Context) Get(e echo.Context) (*Context, error) {
	ctx := GetLogger().Flogger("Get called").Ctx
	db := database.GetPQDatabase(ctx)
	defer db.Close()
	tx := database.GetPQTx(ctx)
	var j string
	err := tx.QueryRow("SELECT status_context FROM job_status WHERE id = $1", c.JobRunID).Scan(&j)
	if err != nil {
		e := merrors.ContextGetError{}.Wrap(db, err)
		return nil, &e
	}
	if err := tx.Commit(); err != nil {
		return nil, merrors.TransactionCommitError{}.Wrap(db, err)
	}
	ctx = c.SetCtx(e)
	if err != nil {
		return nil, merrors.ContextSetError{}.Wrap(db, err)
	}
	return c, nil
}

func (c Context) Set(e echo.Context) (*Context, error) {
	ctx := GetLogger().Flogger("Set called").Ctx
	db := database.GetPQDatabase(ctx)
	defer db.Close()
	tx := database.GetPQTx(ctx)
	j, err := c.Marshal(e)
	if err != nil {
		tx.Rollback()
		e := merrors.ContextSetError{}.Wrap(db, err)
		return nil, &e
	}
	_, err = tx.Exec("UPDATE job_status SET status_context = $1 WHERE id = $2", j, c.JobRunID)
	if err != nil {
		tx.Rollback()
		e := merrors.ContextSetError{}.Wrap(db, err)
		return nil, &e
	}
	if err := tx.Commit(); err != nil {
		return nil, merrors.TransactionCommitError{}.Wrap(db, err)
	}
	ctx = c.SetCtx(e)
	return &c, nil
}

func (c Context) GetCtx(e echo.Context) (*Context, error) {
	ctx := GetLogger().Flogger("GetCtx called").Ctx
	eInterface := ctx.Value(contextKey)
	if innerContext, ok := eInterface.(Context); ok {
		return &innerContext, nil
	} else {
		ctx = c.SetCtx(e)
	}
	return &c, nil
}

func (c Context) SetCtx(e echo.Context) context.Context {
	ctx := GetLogger().Flogger("SetCtx called").Ctx
	ctx = context.WithValue(ctx, contextKey, c)
	return ctx
}
