package types

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type ShallowComfyNode struct {
	ShallowModel
	Name           string       `form:"name" json:"name"`
	Prompt         string       `form:"prompt" json:"prompt"`
	APIBase        string       `form:"api_base" json:"api_base"`
	APITemplate    string       `form:"api_template" json:"api_template"`
	TemplateValues string		`json:"template_values"`
	WorkflowID     WorkflowID	`form:"workflow_id" json:"workflow_id"`
	Type           string       `form:"type" json:"type"`
	Enabled        bool	        `json:"enabled"`
	Bypass         bool         `json:"bypass"`
	Output         string       `form:"output" json:"output"`
}

func (c ShallowComfyNode) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowComfyNode) Expand(e echo.Context) (*ComfyNode, error) {
	r := ComfyNode{}
	if c.ShallowModel.CreatedAt.IsZero() && c.ShallowModel.ID != "" {
		sc, err := c.ShallowModel.Get(e)
		if err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err).Log()
		}
		if err := json.Unmarshal([]byte(sc.Content), &r); err != nil {
			return nil, merrors.ContentGetError{}.Wrap(err).Log()
		}
		return &r, nil
	}
	r.Model = r.Model.FromShallowModel(c.ShallowModel)
	r.ID = c.ID
	r.Name = c.Name
	r.Prompt = c.Prompt
	r.APIBase = c.APIBase
	r.APITemplate = c.APITemplate
	r.TemplateValues = c.TemplateValues
	r.WorkflowID = c.WorkflowID
	r.Type = c.Type
	r.Enabled.Value = c.Enabled
	r.Bypass.Value = c.Bypass
	r.Output = c.Output
	return &r, nil
}

func (c ShallowComfyNode) Validate() ShallowComfyNode {
	valid := true
	if !c.ShallowModel.Validate() {
		valid = false
	}
	if c.ShallowModel.ContentType != "comfynode" {
		valid = false
	}
	if c.ID == "" || c.ID != c.ShallowModel.ID {
		valid = false
	}
	if c.Name == "" || c.APIBase == "" || c.APITemplate == "" {
		valid = false
	}
	c.ShallowModel.Validated = valid
	return c
}

func (c ShallowComfyNode) GetValidated() bool {
	return c.ShallowModel.Validated
}

func (c ShallowComfyNode) GetName() string {
	return c.Name
}

func (c ShallowComfyNode) GetCommand() string {
	return ""
}

func (c ShallowComfyNode) GetUser() string {
	return ""
}

func (c ShallowComfyNode) GetHost() string {
	return ""
}

func (c ShallowComfyNode) GetModel() string {
	return ""
}

func (c ShallowComfyNode) GetSystemPrompt() string {
	return ""
}

func (c ShallowComfyNode) GetPrompt() string {
	return ""
}

func (c ShallowComfyNode) GetPromptTemplate() string {
	return ""
}

func (c ShallowComfyNode) GetApiBase() string {
	return c.APIBase
}

func (c ShallowComfyNode) GetApiTemplate() string {
	return c.APITemplate
}
func (c ShallowComfyNode) GetType() string {
	return "comfy_node"
}

func (c *ShallowComfyNode) FromMSI(msi map[string]interface{}) error {
	var err error
	if id, ok := msi["id"].(string); ok {
		c.ShallowModel.ID = id
	}
	if createdAt, ok := msi["CreatedAt"].(string); ok {
		c.ShallowModel.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct: "ComfyNode", Function: "FromMSI"}.Wrap(err).Log()
		}
	}
	if updatedAt, ok := msi["UpdatedAt"].(string); ok {
		c.ShallowModel.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct: "ComfyNode", Function: "FromMSI"}.Wrap(err).Log()
		}
	}
	if ct, ok := msi["ContentType"].(string); ok {
		c.ShallowModel.ContentType = ct
	}
	if name, ok := msi["name"].(string); ok {
		c.Name = name
	}
	if ab, ok := msi["api_base"].(string); ok {
		c.APIBase = ab
	}
	if at, ok := msi["api_template"].(string); ok {
		c.APITemplate = at
	}
	return nil
}

func (c *ShallowComfyNode) Get(e echo.Context) error {
	content := NewComfyNodeTypeContent()
	content.Model.ID = c.ShallowModel.ID
	content.Model.ContentType = "shallowcomfynode"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err).Log()
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowComfyNode", Function: "Get"}.Wrap(err).Log()
	}
	return nil
}

func (c *ShallowComfyNode) GetShallow(e echo.Context) error {
	content := NewComfyNodeTypeContent()
	content.Model.ID = c.ShallowModel.ID
	content.Model.ContentType = "shallowcomfynode"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err).Log()
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowComfyNode", Function: "Get"}.Wrap(err).Log()
	}
	return nil
}

func NewShallowComfyNodeTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "shallowcomfynode"
	return c
}

func NewShallowComfyNodeModelContent() models.ShallowContent {
	c := models.ShallowContent{}
	c.ShallowModel.ContentType = "shallowcomfynode"
	return c
}

func (c ShallowComfyNode) Delete(e echo.Context) error {
	content := NewShallowComfyNodeTypeContent()
	content.FromType(c, c.ShallowModel)
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ID = c.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.ShallowModel.ID, Package: "types", Struct: "ShallowComfyNode", Function: "delete"}.Wrap(err).Log()
	}
	return nil
}

func (c ShallowComfyNode) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowComfyNode) GetID() string {
	return c.ShallowModel.ID
}

func NewShallowComfyNode(id *string) ShallowComfyNode {
	c := ShallowComfyNode{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "comfynode"
	if c.ShallowModel.CreatedAt.IsZero() {
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
	} else {
		c.ShallowModel.UpdatedAt = time.Now()
	}
	return c
}

func (c ShallowComfyNode) Set(e echo.Context, update bool) error {
	c.Validate()
	if !c.ShallowModel.Validated {
		return merrors.ContentValidationError{Package: "types", Struct: "node", Function: "set"}.New("validation failed")
	}
	content := NewComfyNodeTypeContent()
	content.FromType(c, content.Model)
	content.Model.ID = c.ShallowModel.ID
	content.Model.CreatedAt = c.ShallowModel.CreatedAt
	content.Model.UpdatedAt = c.ShallowModel.UpdatedAt
	content.Model.ContentType = c.ShallowModel.ContentType
	err := content.Set(e, update)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(err).Log()
	}
	return nil
}

func (c ShallowComfyNode) IsShallowModel() bool {
	return true
}
