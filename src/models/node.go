package models

type params interface {
	GetType() string
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

