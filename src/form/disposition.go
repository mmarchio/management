package form

type Disposition struct {
	Name 					Text			`form:"name" json:"name" placeholder:"name"`
	MinDuration 			Number			`form:"min_duration" json:"min_duration" placeholder:"min duration"`
	MaxDuration 			Number			`form:"max_duration" json:"max_duration" placeholder:"max duration"`
	AdvertisementDuration 	Number			`form:"advertisement_duration" json:"advertisement_duration" placeholder:"advertisement_duration"`
}