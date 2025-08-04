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
	Model
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
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_entitlements"
	return c
}
