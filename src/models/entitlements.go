package models

import (
	"time"

	"github.com/google/uuid"
)

var err error

type Entitlements struct {
	Model
	YouTube 	Toggle `form:"youtube" json:"youtube"`
	TikTok 		Toggle `form:"tiktok" json:"tiktok"`
	Rumble 		Toggle `form:"rumble" json:"rumble"`
	Patreon 	Toggle `form:"patreon" json:"patreon"`
	Facebook 	Toggle `form:"facebook" json:"facebook"`
	Base64Value string
	JsonValue 	[]byte
}

type ShallowEntitlements struct {
	ShallowModel
	YouTube 	string `form:"youtube" json:"youtube"`
	TikTok 		string `form:"tiktok" json:"tiktok"`
	Rumble 		string `form:"rumble" json:"rumble"`
	Patreon 	string `form:"patreon" json:"patreon"`
	Facebook 	string `form:"facebook" json:"facebook"`
	Base64Value string
	JsonValue 	[]byte
}

func NewShallowEntitlements(id *string) ShallowEntitlements {
	c := ShallowEntitlements{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallowentitlements"
	return c
}
