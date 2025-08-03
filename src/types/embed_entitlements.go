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
	YouTube 	Toggle `form:"youtube" json:"youtube"`
	TikTok 		Toggle `form:"tiktok" json:"tiktok"`
	Rumble 		Toggle `form:"rumble" json:"rumble"`
	Patreon 	Toggle `form:"patreon" json:"patreon"`
	Facebook 	Toggle `form:"facebook" json:"facebook"`
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
	c.YouTube.init()
	c.YouTube.New(*c)
	c.TikTok.init()
	c.TikTok.New(*c)
	c.Rumble.init()
	c.Rumble.New(*c)
	c.Patreon.init()
	c.Patreon.New(*c)
	c.Facebook.init()
	c.Facebook.New(*c)
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
	c.YouTube = toggle.FromModel(model.YouTube)
	c.TikTok = toggle.FromModel(model.TikTok)
	c.Rumble = toggle.FromModel(model.Rumble)
	c.Patreon = toggle.FromModel(model.Patreon)
	c.Facebook = toggle.FromModel(model.Facebook)
	return c
}

func (c Entitlements) Bind(e echo.Context) (Entitlements, error) {
	if e.FormValue(c.YouTube.NamePrefix+c.YouTube.Suffix) == "on" {
		c.YouTube.Value = true
	}
	if e.FormValue(c.TikTok.NamePrefix+c.TikTok.Suffix) == "on" {
		c.TikTok.Value = true
	}
	if e.FormValue(c.Rumble.NamePrefix+c.Rumble.Suffix) == "on" {
		c.Rumble.Value = true
	}
	if e.FormValue(c.Patreon.NamePrefix+c.Patreon.Suffix) == "on" {
		c.Patreon.Value = true
	}
	if e.FormValue(c.Facebook.NamePrefix+c.Facebook.Suffix) == "on" {
		c.Facebook.Value = true
	}
	return c, nil
}

func (c Entitlements) Truncate() map[string]interface{} {
	entitlements := make(map[string]interface{})
	entitlements["youtube"] = c.YouTube.Value
	entitlements["tiktok"] = c.TikTok.Value
	entitlements["rumble"] = c.Rumble.Value
	entitlements["patreon"] = c.Patreon.Value
	entitlements["facebook"] = c.Facebook.Value
	return entitlements
}

func ValidateEntitlements(c Entitlements, prefix string) (Entitlements, error) {
	c.YouTube = ValidateToggle(c.YouTube, uuid.NewString(), prefix, "youtube", "youtube")
	c.TikTok = ValidateToggle(c.TikTok, uuid.NewString(), prefix, "tiktok", "tiktok")
	c.Rumble = ValidateToggle(c.Rumble, uuid.NewString(), prefix, "rumble", "rumble")
	c.Patreon = ValidateToggle(c.Patreon, uuid.NewString(), prefix, "patreon", "patreon")
	c.Facebook = ValidateToggle(c.Facebook, uuid.NewString(), prefix, "facebook", "facebook")
	return c, nil
}