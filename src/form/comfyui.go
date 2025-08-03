package form

type ComfyUITemplate struct {
	Name 		Text     `form:"name" json:"name" placeholder:"name"`
	Endpoint 	Text 	 `form:"endpoint" json:"endpoint" placeholder:"endpoint"`
	Base 		Textarea `form:"base" json:"base" placeholder:"base"`
	Prompt      Textarea `form:"prompt" json:"prompt" placeholder:"prompt"`
	Template 	Textarea `form:"template" json:"template" placeholder:"template"`
}

func (c ComfyUITemplate) IsFormDefinition() bool {
	return true
}