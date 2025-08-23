package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
)

type SSHNode struct {
	Model
	Name       string     `form:"name" json:"name"`
	Command    string     `form:"command" json:"command"`
	User       string     `form:"user" json:"user"`
	Host       string     `form:"host" json:"host"`
	WorkflowID WorkflowID `form:"workflow_id" json:"workflow_id"`
	Type       string     `form:"type" json:"type"`
	Enabled    bool       `json:"enabled"`
	Bypass     bool       `json:"bypass"`
	Output     string     `form:"output" json:"output"`
}

func (c SSHNode) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowSSHNode{}
	sm.ShallowModel = sm.ShallowModel.FromTypeModel(c.Model)
	sm.ID = c.ID
	sm.Name = c.Name
	sm.Command = c.Command
	sm.User = c.User
	sm.Host = c.Host
	sm.WorkflowID = c.WorkflowID
	sm.Type = c.Type
	sm.Enabled = c.Enabled
	sm.Bypass = c.Bypass
	sm.Output = c.Output
	sms = append(sms, sm)
	return sms
}

func (c SSHNode) Validate() SSHNode {
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
			return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct: "SSHNode", Function: "FromMSI"}.Wrap(err).Log()
		}
	}
	if updatedAt, ok := msi["UpdatedAt"].(string); ok {
		c.Model.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct: "SSHNode", Function: "FromMSI"}.Wrap(err).Log()
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

func (c *SSHNode) Get(e echo.Context) error {
	content := NewSSHNodeTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "sshnode"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "node", Function: "Get"}.Wrap(err).Log()
	}
	return nil
}

func NewSSHNodeTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "sshnode"
	return c
}

func (c SSHNode) Delete(e echo.Context) error {
	content := NewSSHNodeTypeContent()
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	content.ID = c.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID, Package: "types", Struct: "sshnode", Function: "delete"}.Wrap(err).Log()
	}
	wf := NewWorkflow(nil)
	wf.Model.ID = c.WorkflowID.String()
	if err := wf.Get(e); err != nil {
		return merrors.ContentGetError{}.Wrap(err).Log()
	}
	wf.CutNode(c.Model.ID)
	wf.CutNodeOrder(c.Model.ID)
	if err := wf.Set(e, false); err != nil {
		return merrors.ContentSetError{}.Wrap(err).Log()
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

func (c SSHNode) Set(e echo.Context, update bool) error {
	c.Validate()
	if !c.Model.Validated {
		return merrors.ContentValidationError{Package: "types", Struct: "node", Function: "set"}.New("validation failed")
	}
	content := NewSSHNodeTypeContent()
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	content.ID = c.Model.ID
	err := content.Set(e, update)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	return nil
}

func (c SSHNode) List(e echo.Context) ([]SSHNode, error) {
	content := NewSSHNodeTypeContent()
	content.Model.ContentType = "sshnode"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]SSHNode, 0)
	for _, model := range contents {
		cut := SSHNode{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "SSHNode", Function: "List"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c SSHNode) ListBy(e echo.Context, key string, value interface{}) ([]SSHNode, error) {
	content := NewSSHNodeTypeContent()
	content.Model.ContentType = "sshnode"
	list, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{Info: fmt.Sprintf("{\"%s\":\"%s\"}", key, value), Package: "types", Struct: "SSHNode", Function: "ListBy"}.Wrap(err).Log()
	}
	cuts := make([]SSHNode, 0)
	for _, model := range list {
		cut := SSHNode{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "SSHNode", Function: "ListBy"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c SSHNode) Exec(e echo.Context, jobrun *JobRun, step *Step) error {
	return nil
}
