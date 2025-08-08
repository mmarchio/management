package models

type params interface {
	GetType() string
	Validate() params
	GetValidated() bool
}

type Node struct {
	Model
	ID 		string `json:"id"`
	Name 	string `json:"name"`
	Type 	string `json:"type"`
	Params 	params `json:"params"`
	Enabled bool `json:"enabled"`
	Bypass 	bool `json:"bypass"`
	Output 	string `json:"output"`
}

func (c *Node) Validate() {
	valid := true
	if c.Model.ID == "" {
		valid = false
	}
	if c.Model.CreatedAt.IsZero() || c.Model.UpdatedAt.IsZero() {
		valid = false
	}
	if c.Model.ContentType != "node" {
		valid = false
	}
	if c.ID == "" {
		valid = false
	}
	if c.Name == "" {
		valid = false
	}
	if c.Type == "" {
		valid = false
	}
	if c.Params != nil {
		c.Params = c.Params.Validate()
		if !c.Params.GetValidated() {
			valid = false
		}
	}
	c.Model.Validated = valid
}