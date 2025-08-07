package types

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
)

type ShallowSSHNode struct {
	ShallowModel
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

func (c ShallowSSHNode) Validate() params {
	valid := true
	if !c.ShallowModel.Validate() {
		valid = false
	}
	if c.ShallowModel.ContentType != "sshnode" {
		valid = false
	}
	if c.ID == "" || c.Name == "" || c.Command == "" || c.User == "" || c.Host == "" {
		valid = false
	}
	if c.ID != c.ShallowModel.ID {
		valid = false
	}
	c.ShallowModel.Validated = valid
	return c
}

func (c ShallowSSHNode) GetValidated() bool {
	return c.ShallowModel.Validated
}

func (c ShallowSSHNode) GetName() string {
	return c.Name
}

func (c ShallowSSHNode) GetCommand() string {
	return c.Command
}

func (c ShallowSSHNode) GetUser() string {
	return c.User
}

func (c ShallowSSHNode) GetHost() string {
	return c.Host
}

func (c ShallowSSHNode) GetModel() string {
	return ""
}

func (c ShallowSSHNode) GetSystemPrompt() string {
	return ""
}

func (c ShallowSSHNode) GetPrompt() string {
	return ""
}

func (c ShallowSSHNode) GetPromptTemplate() string {
	return ""
}

func (c ShallowSSHNode) GetApiBase() string {
	return ""
}

func (c ShallowSSHNode) GetApiTemplate() string {
	return ""
}
func (c ShallowSSHNode) GetType() string {
	return "ssh_node"
}

func (c *ShallowSSHNode) FromMSI(msi map[string]interface{}) error {
	var err error
	if id, ok := msi["id"].(string); ok {
		c.ShallowModel.ID = id
	}
	if createdAt, ok := msi["CreatedAt"].(string); ok {
		c.ShallowModel.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct:"SSHNode", Function: "FromMSI"}.Wrap(err)
		}
	}
	if updatedAt, ok := msi["UpdatedAt"].(string); ok {
		c.ShallowModel.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct:"SSHNode", Function: "FromMSI"}.Wrap(err)
		}
	}
	if ct, ok := msi["ContentType"].(string); ok {
		c.ShallowModel.ContentType = ct
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

func (c *ShallowSSHNode) Get(ctx context.Context) error {
	content := NewSSHNodeTypeContent()
	content.Model.ID = c.ShallowModel.ID
	content.Model.ContentType = "sshnode"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "node", Function: "Get"}.Wrap(err)
	}
	return nil
}

func NewShallowSSHNodeTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "sshnode"
	return c
}

func (c ShallowSSHNode) Delete(ctx context.Context) error {
	content := NewSSHNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.ShallowModel.ID
	content.ID = c.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.ContentDeleteError{Info: c.ShallowModel.ID, Package: "types", Struct: "sshnode", Function: "delete"}.Wrap(err)
	}
	return nil
}

func (c ShallowSSHNode) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowSSHNode) GetID() string {
	return c.ShallowModel.ID
}

func NewShallowSSHNode(id *string) SSHNode {
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

func (c ShallowSSHNode) Set(ctx context.Context) error {
	c.Validate()
	if !c.ShallowModel.Validated {
		return merrors.ContentValidationError{Package: "types", Struct: "node", Function: "set"}.Wrap(fmt.Errorf("validation failed"))
	}
	content := NewSSHNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.ShallowModel.ID
	content.ID = c.ShallowModel.ID
	err := content.Set(ctx)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

