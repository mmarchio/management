package models

type ComfyUITemplate struct {
	Model
	Name 		string `form:"name" json:"name"`
	Endpoint 	string `form:"endpoint" json:"endpoint"`
	Base 		string `form:"base" json:"base"`
	Template 	string `form:"template" json:"template"`
}

func (c *ComfyUITemplate) New() {
}