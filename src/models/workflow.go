package models

type Workflow struct {
	Model
	ID string `json:"id"`
	Name string `form:"name" json:"name"`
	Nodes []Node `json:"nodes"`
}
