package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/mmarchio/management/database"
	merrors "github.com/mmarchio/management/errors"
)

type ContextKeyT int64
var contextKey ContextKeyT = 3


type Context struct {
	Model
	Prompt Prompt `json:"prompt"`
	Disposition Disposition `json:"disposition"`
	JobRunID string `json:"job_run_id"`
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

type Stats struct {
	Model
	ID string `json:"stats_id"`
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
	Input string `json:"input"`
	Output string `json:"output"`
	Duration time.Duration `json:"duration"`
	Files []File `json:"files"`
	Status string `json:"status"`
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

type Output struct {
	Model
	ID string `json:"id"`
	Stats Stats `json:"stats"`
	Files []File `json:"files"`
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


type AudioOutput Output

type VideoLipsyncOutput Output

type VideoTransparancyOutput Output

type VideoBackgroundOutput Output

type VideoLayerMergeOutput Output

type ImageThumbnailOutput Output

type ImageBackgroundContextOutput Output

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
