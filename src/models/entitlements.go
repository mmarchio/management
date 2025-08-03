package models

var err error

type Entitlements struct {
	YouTube 	Toggle `form:"youtube" json:"youtube"`
	TikTok 		Toggle `form:"tiktok" json:"tiktok"`
	Rumble 		Toggle `form:"rumble" json:"rumble"`
	Patreon 	Toggle `form:"patreon" json:"patreon"`
	Facebook 	Toggle `form:"facebook" json:"facebook"`
	Base64Value string
	JsonValue 	[]byte
}
