package models

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

