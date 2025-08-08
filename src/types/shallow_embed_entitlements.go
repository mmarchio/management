package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type ShallowEntitlements struct {
	ShallowModel
	YouTubeModel 	string `form:"youtube" json:"youtube"`
	TikTokModel 	string `form:"tiktok" json:"tiktok"`
	RumbleModel 	string `form:"rumble" json:"rumble"`
	PatreonModel 	string `form:"patreon" json:"patreon"`
	FacebookModel 	string `form:"facebook" json:"facebook"`
}

func (c ShallowEntitlements) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(err)
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowEntitlements) Expand(ctx context.Context) (*Entitlements, error) {
	r := Entitlements{}
	r.EmbedModel.FromShallowModel(c.ShallowModel)
	t := ShallowToggle{}
	t.ShallowModel.ID = c.YouTubeModel
	youTubeModel, err := t.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.YouTubeModel = *youTubeModel
	t.ShallowModel.ID = c.TikTokModel
	tikTokModel, err := t.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.TikTokModel = *tikTokModel
	t.ShallowModel.ID = c.RumbleModel
	rumbleModel, err := t.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.RumbleModel = *rumbleModel
	t.ShallowModel.ID = c.PatreonModel
	patreonModel, err := t.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.PatreonModel = *patreonModel
	t.ShallowModel.ID = c.FacebookModel
	facebookModel, err := t.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.FacebookModel = *facebookModel
	return &r, nil
}

func (c *ShallowEntitlements) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c ShallowEntitlements) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ShallowEntitlements) New(parent ITable) (string, error) {
	ct := "entitlement"
	c.ShallowModel.New(nil, &ct)
	c.ShallowModel.ID = parent.GetID()
	embedBytes, err := json.Marshal(c)
	if err != nil {
		return "", merrors.JSONMarshallingError{}.Wrap(err)
	}
	return string(embedBytes), nil
}

func (c ShallowEntitlements) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowEntitlements) GetID() string {
	return c.ShallowModel.ID
}

func (c ShallowEntitlements) FromModel(model models.ShallowEntitlements) ShallowEntitlements {
	c.ShallowModel.FromModel(model.ShallowModel)
	c.YouTubeModel = model.YouTube
	c.TikTokModel = model.TikTok
	c.RumbleModel = model.Rumble
	c.PatreonModel = model.Patreon
	c.FacebookModel = model.Facebook
	return c
}

func (c ShallowEntitlements) Truncate() map[string]interface{} {
	entitlements := make(map[string]interface{})
	entitlements["youtube"] = c.YouTubeModel
	entitlements["tiktok"] = c.TikTokModel
	entitlements["rumble"] = c.RumbleModel
	entitlements["patreon"] = c.PatreonModel
	entitlements["facebook"] = c.FacebookModel
	return entitlements
}

func ValidateShallowEntitlements(c ShallowEntitlements, prefix string) (ShallowEntitlements, error) {
	if c.YouTubeModel != "" {
		_, err := uuid.Parse(c.YouTubeModel)
		if err != nil {
			return c, err
		}
	}
	if c.TikTokModel != "" {
		_, err := uuid.Parse(c.TikTokModel)
		if err != nil {
			return c, err
		}
	}
	if c.RumbleModel != "" {
		_, err := uuid.Parse(c.RumbleModel)
		if err != nil {
			return c, err
		}
	}
	if c.PatreonModel != "" {
		_, err := uuid.Parse(c.PatreonModel)
		if err != nil {
			return c, err
		}
	}
	if c.FacebookModel != "" {
		_, err := uuid.Parse(c.FacebookModel)
		if err != nil {
			return c, err
		}
	}

	return c, nil
}

func (c ShallowEntitlements) IsShallowModel() bool {
	return true
}
