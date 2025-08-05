package types

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type Entitlements struct {
	EmbedModel
	YouTubeModel 	Toggle `form:"youtube" json:"youtube"`
	TikTokModel 		Toggle `form:"tiktok" json:"tiktok"`
	RumbleModel 		Toggle `form:"rumble" json:"rumble"`
	PatreonModel 	Toggle `form:"patreon" json:"patreon"`
	FacebookModel 	Toggle `form:"facebook" json:"facebook"`
}

type ShallowEntitlements struct {
	ShallowModel
	YouTubeModel 	string `form:"youtube" json:"youtube"`
	TikTokModel 	string `form:"tiktok" json:"tiktok"`
	RumbleModel 	string `form:"rumble" json:"rumble"`
	PatreonModel 	string `form:"patreon" json:"patreon"`
	FacebookModel 	string `form:"facebook" json:"facebook"`
}

func (c *Entitlements) Unmarshal(ctx context.Context, j string) error {
	return json.Unmarshal([]byte(j), c)
}

func (c Entitlements) Marshal(ctx context.Context) (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Entitlements) New(parent ITable) (string, error) {
	c.EmbedModel.New("entitlement")
	c.EmbedModel.ID = parent.GetID()
	c.YouTubeModel.init()
	c.YouTubeModel.New(*c)
	c.TikTokModel.init()
	c.TikTokModel.New(*c)
	c.RumbleModel.init()
	c.RumbleModel.New(*c)
	c.PatreonModel.init()
	c.PatreonModel.New(*c)
	c.FacebookModel.init()
	c.FacebookModel.New(*c)
	embedBytes, err := json.Marshal(c)
	if err != nil {
		return "", merrors.JSONMarshallingError{}.Wrap(err)
	}
	return string(embedBytes), nil
}

func (c Entitlements) GetContentType() string {
	return c.EmbedModel.ContentType
}

func (c Entitlements) GetID() string {
	return c.EmbedModel.ID
}

func (c Entitlements) FromModel(model models.Entitlements) Entitlements {
	toggle := Toggle{}
	c.YouTubeModel = toggle.FromModel(model.YouTube)
	c.TikTokModel = toggle.FromModel(model.TikTok)
	c.RumbleModel = toggle.FromModel(model.Rumble)
	c.PatreonModel = toggle.FromModel(model.Patreon)
	c.FacebookModel = toggle.FromModel(model.Facebook)
	return c
}

func (c Entitlements) Bind(e echo.Context) (Entitlements, error) {
	if e.FormValue(c.YouTubeModel.NamePrefix+c.YouTubeModel.Suffix) == "on" {
		c.YouTubeModel.Value = true
	}
	if e.FormValue(c.TikTokModel.NamePrefix+c.TikTokModel.Suffix) == "on" {
		c.TikTokModel.Value = true
	}
	if e.FormValue(c.RumbleModel.NamePrefix+c.RumbleModel.Suffix) == "on" {
		c.RumbleModel.Value = true
	}
	if e.FormValue(c.PatreonModel.NamePrefix+c.PatreonModel.Suffix) == "on" {
		c.PatreonModel.Value = true
	}
	if e.FormValue(c.FacebookModel.NamePrefix+c.FacebookModel.Suffix) == "on" {
		c.FacebookModel.Value = true
	}
	return c, nil
}

func (c Entitlements) Truncate() map[string]interface{} {
	entitlements := make(map[string]interface{})
	entitlements["youtube"] = c.YouTubeModel.Value
	entitlements["tiktok"] = c.TikTokModel.Value
	entitlements["rumble"] = c.RumbleModel.Value
	entitlements["patreon"] = c.PatreonModel.Value
	entitlements["facebook"] = c.FacebookModel.Value
	return entitlements
}

func ValidateEntitlements(c Entitlements, prefix string) (Entitlements, error) {
	c.YouTubeModel = ValidateToggle(c.YouTubeModel, uuid.NewString(), prefix, "youtube", "youtube")
	c.TikTokModel = ValidateToggle(c.TikTokModel, uuid.NewString(), prefix, "tiktok", "tiktok")
	c.RumbleModel = ValidateToggle(c.RumbleModel, uuid.NewString(), prefix, "rumble", "rumble")
	c.PatreonModel = ValidateToggle(c.PatreonModel, uuid.NewString(), prefix, "patreon", "patreon")
	c.FacebookModel = ValidateToggle(c.FacebookModel, uuid.NewString(), prefix, "facebook", "facebook")
	return c, nil
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