package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Disposition struct {
	Model
	Name 					string			`json:"name"`
	MinDuration 			int64			`json:"min_duration"`
	MaxDuration 			int64			`json:"max_duration"`
	AdvertisementDuration 	int64			`json:"advertisement_duration"`
	EntitlementsModel 			Entitlements	`json:"entitlements_model"`
	VerificationModel 			Steps			`json:"verification_model"`
	BypassModel 					Steps			`json:"bypass_model"`
	JsonValue				[]byte
	Base64Value				string
}

type ShallowDisposition struct {
	Model
	Name 					string			`json:"name"`
	MinDuration 			int64			`json:"min_duration"`
	MaxDuration 			int64			`json:"max_duration"`
	AdvertisementDuration 	int64			`json:"advertisement_duration"`
	EntitlementsModel 			string			`json:"entitlements_model"`
	VerificationModel 			string			`json:"verification_model"`
	BypassModel 					string			`json:"bypass_model"`
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

func (c ShallowDisposition) Get(ctx context.Context, mode string) (*Disposition, *ShallowDisposition, error) {
	return nil, nil, nil 
}
