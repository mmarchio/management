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
	ShallowModel
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
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_disposition"
	return c
}

func (c ShallowDisposition) Get(ctx context.Context, mode string) (*Disposition, *ShallowDisposition, error) {
	return nil, nil, nil 
}
