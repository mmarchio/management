package models

import (
	"time"

	"github.com/google/uuid"
)

type Disposition struct {
	Model
	Name 					string			`json:"name"`
	MinDuration 			int64			`json:"min_duration"`
	MaxDuration 			int64			`json:"max_duration"`
	AdvertisementDuration 	int64			`json:"advertisement_duration"`
	Entitlements 			Entitlements	`json:"entitlements"`
	Verification 			Steps			`json:"verification"`
	Bypass 					Steps			`json:"bypass"`
	JsonValue				[]byte
	Base64Value				string
}

type ShallowDisposition struct {
	Model
	Name 					string			`json:"name"`
	MinDuration 			int64			`json:"min_duration"`
	MaxDuration 			int64			`json:"max_duration"`
	AdvertisementDuration 	int64			`json:"advertisement_duration"`
	Entitlements 			string			`json:"entitlements"`
	Verification 			string			`json:"verification"`
	Bypass 					string			`json:"bypass"`
	JsonValue				[]byte
	Base64Value				string
}

func NewShallowDisposition(id *string) ShallowDisposition {
	c := ShallowDisposition{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_disposition"
	return c
}
