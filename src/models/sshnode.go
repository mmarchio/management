package models

import (
	"time"

	"github.com/google/uuid"
)

type SSHNode struct {
	Model
	ID 			string `json:"id"`
	Name 		string `form:"name" json:"name"`
	Command 	string `form:"command" json:"command"`
	User 		string `form:"user" json:"user"`
	Host 		string `form:"host" json:"host"`
	WorkflowID  string `form:"workflow_id" json:"workflow_id"`
	Type 		string `form:"type" json:"type"`
	Enabled 	bool   `json:"enabled"`
	Bypass 		bool   `json:"bypass"`
	Output 		string `form:"output" json:"output"`
}

type ShallowSSHNode struct {
	Model
	ID 			string `json:"id"`
	Name 		string `form:"name" json:"name"`
	Command 	string `form:"command" json:"command"`
	User 		string `form:"user" json:"user"`
	Host 		string `form:"host" json:"host"`
	WorkflowID  string `form:"workflow_id" json:"workflow_id"`
	Type 		string `form:"type" json:"type"`
	Enabled 	bool   `json:"enabled"`
	Bypass 		bool   `json:"bypass"`
	Output 		string `form:"output" json:"output"`
}

func NewShallowSSHNode(id *string) ShallowSSHNode {
	c := ShallowSSHNode{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.CreatedAt
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "shallow_sshnode"
	return c
}
