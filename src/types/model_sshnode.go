package types

import (
	"time"

	merrors "github.com/mmarchio/management/errors"
)

type SSHNode struct {
	Model
	ID 			string `json:"id"`
	Name 		string `form:"name" json:"name"`
	Command 	string `form:"command" json:"command"`
	User 		string `form:"user" json:"user"`
	Host 		string `form:"host" json:"host"`
}

func (c SSHNode) GetName() string {
	return c.Name
}

func (c SSHNode) GetCommand() string {
	return c.Command
}

func (c SSHNode) GetUser() string {
	return c.User
}

func (c SSHNode) GetHost() string {
	return c.Host
}

func (c SSHNode) GetModel() string {
	return ""
}

func (c SSHNode) GetSystemPrompt() string {
	return ""
}

func (c SSHNode) GetPrompt() string {
	return ""
}

func (c SSHNode) GetPromptTemplate() string {
	return ""
}

func (c SSHNode) GetApiBase() string {
	return ""
}

func (c SSHNode) GetApiTemplate() string {
	return ""
}
func (c SSHNode) GetType() string {
	return "ssh_node"
}

func (c *SSHNode) FromMSI(msi map[string]interface{}) error {
	var err error
	if id, ok := msi["id"].(string); ok {
		c.Model.ID = id
	}
	if createdAt, ok := msi["CreatedAt"].(string); ok {
		c.Model.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct:"SSHNode", Function: "FromMSI"}.Wrap(err)
		}
	}
	if updatedAt, ok := msi["UpdatedAt"].(string); ok {
		c.Model.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct:"SSHNode", Function: "FromMSI"}.Wrap(err)
		}
	}
	if ct, ok := msi["ContentType"].(string); ok {
		c.Model.ContentType = ct
	}
	if cmd, ok := msi["command"].(string); ok {
		c.Command = cmd
	}
	if user, ok := msi["user"].(string); ok {
		c.User = user
	}
	if host, ok := msi["host"].(string); ok {
		c.Host = host
	}
	return nil
}