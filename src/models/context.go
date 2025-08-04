package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mmarchio/management/database"
	merrors "github.com/mmarchio/management/errors"
)

type ContextKeyT int64
var contextKey ContextKeyT = 3


type Context struct {
	Model
	Prompt 							Prompt `json:"prompt"`
	Disposition 					Disposition `json:"disposition"`
	JobRunID 						string `json:"job_run_id"`
	Settings 						Settings `json:"settings"`
	GetResearchPrompt 				Stats `json:"get_research_prompt"`
	GetResearchOutput 				Stats `json:"get_research_output"`
	ScreenwritingPrompt 			Stats `json:"screenwriting_prompt"`
	ScreenWritingOutput 			Stats `json:"screenwriting_output"`
	VideoPrompt 					Stats `json:"video_prompt"`
	AudioPrompt 					Stats `json:"audio_prompt"`
	AudioOutput 					AudioOutput `json:"audio_output"`
	VideoLipsyncOutput 				VideoLipsyncOutput `json:"video_lipsync_output"`
	VideoTransparencyOutput 		VideoTransparancyOutput `json:"video_transparency_output"`
	VideoBackgroundOutput 			VideoBackgroundOutput `json:"video_background_output"`
	VideoLayerMerge 				VideoLayerMergeOutput `json:"video_layer_merge"`
	VideoJoin 						[]Scene `json:"video_join"`
	ImageThumbnailPrompt 			Stats `json:"image_thumbnail_prompt"`
	ImageThumbnailOutput 			ImageThumbnailOutput `json:"image_thumbnail_output"`
	ImageBackgroundContextOutput 	ImageBackgroundContextOutput `json:"image_background_context_output"`
	PublishVideoYoutube 			Stats `json:"publish_video_youtube"`
	PublishVideoTiktok 				Stats `json:"publish_video_tiktok"`
	PublishVideoRumble 				Stats `json:"publish_video_rumble"`
	PublishVideoFacebook 			Stats `json:"publish_video_facebook"`
	PublishSocialFacebook 			Stats `json:"publish_social_facebook"`
	PublishSocialX 					Stats `json:"publish_social_x"`
	PublishSocialYoutube 			Stats `json:"publish_social_youtube"`
	PublishSocialTruth 				Stats `json:"publish_social_truth"`
}

type ShallowContext struct {
	ShallowModel
	Prompt 							string `json:"prompt"`
	Disposition 					string `json:"disposition"`
	JobRunID 						string `json:"job_run_id"`
	Settings 						string `json:"settings"`
	GetResearchPrompt 				string `json:"get_research_prompt"`
	GetResearchOutput 				string `json:"get_research_output"`
	ScreenwritingPrompt 			string `json:"screenwriting_prompt"`
	ScreenWritingOutput 			string `json:"screenwriting_output"`
	VideoPrompt 					string `json:"video_prompt"`
	AudioPrompt 					string `json:"audio_prompt"`
	AudioOutput 					string `json:"audio_output"`
	VideoLipsyncOutput 				string `json:"video_lipsync_output"`
	VideoTransparencyOutput 		string `json:"video_transparency_output"`
	VideoBackgroundOutput 			string `json:"video_background_output"`
	VideoLayerMerge 				string `json:"video_layer_merge"`
	VideoJoin 						[]string `json:"video_join"`
	ImageThumbnailPrompt 			string `json:"image_thumbnail_prompt"`
	ImageThumbnailOutput 			string `json:"image_thumbnail_output"`
	ImageBackgroundContextOutput 	string `json:"image_background_context_output"`
	PublishVideoYoutube 			string `json:"publish_video_youtube"`
	PublishVideoTiktok 				string `json:"publish_video_tiktok"`
	PublishVideoRumble 				string `json:"publish_video_rumble"`
	PublishVideoFacebook 			string `json:"publish_video_facebook"`
	PublishSocialFacebook 			string `json:"publish_social_facebook"`
	PublishSocialX 					string `json:"publish_social_x"`
	PublishSocialYoutube 			string `json:"publish_social_youtube"`
	PublishSocialTruth 				string `json:"publish_social_truth"`
}

func (c ShallowContext) Set(ctx context.Context) error {
	return nil
}

func (c ShallowContext) SetPrompt(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.Prompt = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowContext) SetDisposition(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.Disposition = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetSettings(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.Settings = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetGetResearchPrompt(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.GetResearchPrompt = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetGetResearchOutput(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.GetResearchOutput = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetJobRunID(ctx context.Context, id string) error {
	c.JobRunID = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowContext) SetScreenwritingOutput(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.ScreenWritingOutput = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetScreenwritingPrompt(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.ScreenwritingPrompt = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetVideoPrompt(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.VideoPrompt = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetAudioPrompt(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.AudioPrompt = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetAudioOutput(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.AudioOutput = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetVideoLipsyncOutput(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.VideoLipsyncOutput = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetVideoTransparencyOutput(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.VideoTransparencyOutput = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetVideoBackgroundOutput(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.VideoBackgroundOutput = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetVideoLayerMerge(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.VideoLayerMerge = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetVideoJoin(ctx context.Context, ids []string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, ids...)
	c.VideoJoin = ids
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) AppendVideoJoin(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.VideoJoin = append(c.VideoJoin, id)
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetImageThumbnailPrompt(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.ImageThumbnailPrompt = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetImageThumbnailOutput(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.ImageThumbnailOutput = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetImageBackgroundContextOutput(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.ImageBackgroundContextOutput = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishVideoYoutube(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishVideoYoutube = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishVideoTiktok(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishVideoTiktok = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishVideoRumble(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishVideoRumble = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishVideoFacebook(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishVideoFacebook = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishSocialFacebook(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishSocialFacebook = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishSocialTruth(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishSocialTruth = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishSocialX(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishSocialX = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowContext) SetPublishSocialYoutube(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.PublishSocialYoutube = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}



func (c ShallowContext) Get(ctx context.Context, mode string) (*Context, *ShallowContext, error) {
	content := Content{ID: c.Model.ID}
	if err := content.Get(ctx); err != nil {
		return nil, nil, merrors.ShallowWorkflowGetError{Info: c.Model.ID, Package: "models", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(err)
	}
	if err := json.Unmarshal([]byte(content.Content), &c); err != nil {
		return nil, nil, merrors.JSONUnmarshallingError{Info: content.Content, Package: "models", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(err)
	}
	if mode == "shallow" {
		return nil, &c, nil
	}
	if mode == "full" {
		m := Context{}
		m.Model = c.Model
		promptptr, _, err := NewShallowPrompt(&c.Prompt).Get(ctx, "full")
		if err != nil {
			return nil, nil, err
		}
		if promptptr != nil {
			m.Prompt = *promptptr
		}
		dispositionptr, _, err := NewShallowDisposition(&c.Disposition).Get(ctx, "full")
		if err != nil {
			return nil, nil, err
		}
		if dispositionptr != nil {
			m.Disposition = *dispositionptr
		}

	}
	return nil, nil, merrors.ShallowOllamaNodeGetError{Package: "models", Struct: "ShallowWorkflow", Function: "Get"}.Wrap(fmt.Errorf("unknown mode: %s", mode))
}

type Stats struct {
	Model
	ID 			string `json:"stats_id"`
	Start 		time.Time `json:"start"`
	End 		time.Time `json:"end"`
	Input 		string `json:"input"`
	Output 		string `json:"output"`
	Duration 	time.Duration `json:"duration"`
	Files 		[]File `json:"files"`
	Status 		string `json:"status"`
}

type ShallowStats struct {
	ShallowModel
	ID 			string `json:"stats_id"`
	Start 		time.Time `json:"start"`
	End 		time.Time `json:"end"`
	Input 		string `json:"input"`
	Output 		string `json:"output"`
	Duration 	time.Duration `json:"duration"`
	Files 		[]string `json:"files"`
	Status 		string `json:"status"`
}

func (c ShallowStats) Set(ctx context.Context) error {
	return nil
}

func (c ShallowStats) SetStart(ctx context.Context, id time.Time) error {
	c.Start = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) SetEnd(ctx context.Context, id time.Time) error {
	c.End = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) SetInput(ctx context.Context, id string) error {
	c.Input = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) SetOutput(ctx context.Context, id string) error {
	c.Output = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) SetDuration(ctx context.Context, id time.Duration) error {
	c.Duration = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) SetFiles(ctx context.Context, ids []string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, ids...)
	c.Files = ids
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) AppendFiles(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.Files = append(c.Files, id)
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil	
}

func (c ShallowStats) SetStatus(ctx context.Context, id string) error {
	c.Status = id
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
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
	Stats Stats `json:"stats"`
	Files []File `json:"files"`
}

type ShallowOutput struct {
	ShallowModel
	ID string `json:"id"`
	Stats string `json:"stats"`
	Files []string `json:"files"`
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
	Files []File `json:"files"`
	SceneFile File `json:"scene_file"`
}

type ShallowScene struct {
	ShallowModel
	ID string `json:"id"`
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
	SceneNumber int64 `json:"scene_number"`
	Path string `json:"path"`
	Files []string `json:"files"`
	SceneFile string `json:"scene_file"`
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

type AudioOutput Output
type ShallowAudioOutput ShallowOutput

func NewShallowAudioOutput(id *string) ShallowAudioOutput {
	c := ShallowAudioOutput{}
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


type VideoLipsyncOutput Output
type ShallowVideoLipsyncOutput ShallowOutput

func NewShallowVideoLipsyncOutput(id *string) ShallowVideoLipsyncOutput {
	c := ShallowVideoLipsyncOutput{}
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

func NewShallowVideoTransparencyOutput(id *string) ShallowVideoTransparancyOutput {
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

type VideoBackgroundOutput Output
type ShallowVideoBackgroundOutput ShallowOutput

func NewShallowVideoBackgroundOutput(id *string) ShallowVideoBackgroundOutput {
	c := ShallowVideoBackgroundOutput{}
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

type VideoLayerMergeOutput Output
type ShallowVideoLayerMergeOutput ShallowOutput

func NewShallowVideoLayerMergeOutput(id *string) ShallowVideoLayerMergeOutput {
	c := ShallowVideoLayerMergeOutput{}
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

type ImageThumbnailOutput Output
type ShallowImageThumbnailOutput ShallowOutput

func NewShallowImageThumbnailOutput(id *string) ShallowImageThumbnailOutput {
	c := ShallowImageThumbnailOutput{}
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

type ImageBackgroundContextOutput Output
type ShallowImageBackgroundContextOutput ShallowOutput

func NewShallowImageBackgroundContextOutput(id *string) ShallowImageBackgroundContextOutput {
	c := ShallowImageBackgroundContextOutput{}
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

func (c *Context) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Context) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Context) Get(ctx context.Context) (context.Context, error) {
	ctx, tx := database.GetDBTransaction(ctx)
	var j string
	err := tx.QueryRow(ctx, "SELECT status_context FROM job_status WHERE id = $1", c.JobRunID).Scan(&j)
	if err != nil {
		e := merrors.ContextGetError{}.Wrap(err)
		return ctx, &e
	}
	ctx = c.SetCtx(ctx)
	return ctx, nil
}

func (c Context) Set(ctx context.Context) (context.Context, error) {
	ctx, tx := database.GetDBTransaction(ctx)
	j, err := c.Marshal(ctx)
	if err != nil {
		tx.Rollback(ctx)
		e := merrors.ContextSetError{}.Wrap(err)
		return ctx, &e
	}
	_, err = tx.Exec(ctx, "UPDATE job_status SET status_context = $1 WHERE id = $2", j, c.JobRunID)
	if err != nil {
		tx.Rollback(ctx)
		e := merrors.ContextSetError{}.Wrap(err)
		return ctx, &e
	}
	ctx = c.SetCtx(ctx)
	return ctx, nil
}

func (c Context) GetCtx(ctx context.Context) (*Context, error) {
	ctxInterface := ctx.Value(contextKey)
	if innerContext, ok := ctxInterface.(Context); ok {
		return &innerContext, nil
	} else {
		ctx = c.SetCtx(ctx)
	}
	return &c, nil
}

func (c Context) SetCtx(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, contextKey, c)
	return ctx
}
