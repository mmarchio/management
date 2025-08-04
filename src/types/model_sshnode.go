package types

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
)

type SSHNode struct {
	Model
	ID 			string `json:"id"`
	Name 		string `form:"name" json:"name"`
	Command 	string `form:"command" json:"command"`
	User 		string `form:"user" json:"user"`
	Host 		string `form:"host" json:"host"`
	WorkflowID  WorkflowID `form:"workflow_id" json:"workflow_id"`
	Type 		string `form:"type" json:"type"`
	Enabled 	bool   `json:"enabled"`
	Bypass 		bool   `json:"bypass"`
	Output 		string `form:"output" json:"output"`
}

func (c SSHNode) Validate() params {
	valid := true
	if !c.Model.Validate() {
		valid = false
	}
	if c.Model.ContentType != "sshnode" {
		valid = false
	}
	if c.ID == "" || c.Name == "" || c.Command == "" || c.User == "" || c.Host == "" {
		valid = false
	}
	if c.ID != c.Model.ID {
		valid = false
	}
	c.Model.Validated = valid
	return c
}

func (c SSHNode) GetValidated() bool {
	return c.Model.Validated
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

func (c *SSHNode) Get(ctx context.Context) error {
	content := NewSSHNodeTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "sshnode"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.NodeGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "node", Function: "Get"}.Wrap(err)
	}
	return nil
}

func NewSSHNodeTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "sshnode"
	return c
}

func (c SSHNode) Delete(ctx context.Context) error {
	content := NewSSHNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	content.ID = c.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.NodeDeleteError{Info: c.Model.ID, Package: "types", Struct: "sshnode", Function: "delete"}.Wrap(err)
	}
	return nil
}

func (c SSHNode) GetContentType() string {
	return c.Model.ContentType
}

func (c SSHNode) GetID() string {
	return c.Model.ID
}

func NewSSHNode(id *string) SSHNode {
	c := SSHNode{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "sshnode"
	if c.Model.CreatedAt.IsZero() {
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	} else {
		c.Model.UpdatedAt = time.Now()
	}
	return c
}

func (c SSHNode) Set(ctx context.Context) error {
	c.Validate()
	if !c.Model.Validated {
		return merrors.NodeValidationError{Package: "types", Struct: "node", Function: "set"}.Wrap(fmt.Errorf("validation failed"))
	}
	content := NewSSHNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	content.ID = c.Model.ID
	err := content.Set(ctx)
	if err != nil {
		return merrors.NodeSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

