package types

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type params interface {
	GetName() string
	GetCommand() string 
	GetUser() string 
	GetHost() string 
	GetModel() string 
	GetSystemPrompt() string 
	GetPrompt() string 
	GetPromptTemplate() string 
	GetApiBase() string 
	GetApiTemplate() string 
	GetType() string
}

type Node struct {
	Model
	ID 			NodeID `json:"id"`
	WorkflowID  WorkflowID `form:"workflow_id" json:"workflow_id"`
	Name 		string `form:"name" json:"name"`
	Type 		string `form:"type" json:"type"`
	Params 		params `json:"params"`
	Enabled 	bool   `json:"enabled"`
	Bypass 		bool   `json:"bypass"`
	Output 		string `form:"output" json:"output"`
}

func NewNode(id *string) Node {
	c := Node{}
	c.New(id)
	c.Model.ContentType = "node"
	c, _ = ValidateNode(c)
	return c
} 

func NewNodeModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "node"
	return c
}

func NewNodeTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "node"
	return c
}

func (c *Node) New(id *string) {
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.ID = NodeID(c.Model.ID)
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
}


func (c Node) List(ctx context.Context) ([]Node, error) {
	content := NewNodeModelContent()
	content.Model.ContentType = "node"
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.NodeListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]Node, 0)
	for _, model := range contents {
		cut := NewNode(nil)
		msi := make(map[string]interface{})
		err = json.Unmarshal([]byte(model.Content), &msi)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "node", Function: "List"}.Wrap(err)
		}
		cut.FromMSI(msi)
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c Node) ListBy(ctx context.Context, key string, value interface{}) ([]Node, error) {
	content := NewNodeModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.NodeListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]Node, 0)
	for _, model := range contents {
		cut := Node{}
		msi := make(map[string]interface{})
		err = json.Unmarshal([]byte(model.Content), &msi)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "node", Function: "ListBy"}.Wrap(err)
		}
		cut.FromMSI(msi)
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *Node) Get(ctx context.Context) error {
	content := NewNodeTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "node"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.NodeGetError{Info: c.Model.ID}.Wrap(err)
	}
	msi := make(map[string]interface{})
	err = json.Unmarshal([]byte(content.Content), &msi)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "node", Function: "Get"}.Wrap(err)
	}
	if err := c.FromMSI(msi); err != nil {
		return merrors.JSONUnmarshallingError{Info: "from msi", Package: "types", Struct: "node", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c Node) Set(ctx context.Context) error {
	content := NewNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.ID.String()
	err := content.Set(ctx)
	if err != nil {
		return merrors.NodeSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c Node) Delete(ctx context.Context) error {
	content := NewNodeTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	content.ID = c.ID.String()
	if err := content.Delete(ctx); err != nil {
		return merrors.NodeDeleteError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c Node) GetID() string {
	return c.Model.ID
}

func (c Node) GetContentType() string {
	return c.Model.ContentType
}

func (c Node) GetTable() string {
	return c.Model.Table
}

func (c Node) SetID() (Node, error) {
	var err error
	c.ID = NodeID(c.Model.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "node"}.Wrap(err)
	}
	return c, nil
}

func ValidateNode(p Node) (Node, error) {
	var err error
	return p, err
}

func (c Node) Bind(e echo.Context) (Node, error) {
	var err error
	return c, err
}

func (c Node) Next(e echo.Context, ctx context.Context) (*models.Context, error) {
	systemContext, err := models.Context{}.GetCtx(ctx)
	if err != nil {
		return nil, merrors.ContextGetError{Package: "types", Struct: "Node", Function: "Next"}.Wrap(err)
	}
	return systemContext, nil
}

func (c *Node) FromMSI(msi map[string]interface{}) error {
	var err error
	if id, ok := msi["id"].(string); ok {
		c.Model.ID = id
	}
	if createdAt, ok := msi["CreatedAt"].(string); ok {
		c.Model.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct:"Node", Function: "FromMSI"}.Wrap(err)
		}
	}
	if updatedAt, ok := msi["UpdatedAt"].(string); ok {
		c.Model.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct:"ComfyNode", Function: "FromMSI"}.Wrap(err)
		}
	}
	if contentType, ok := msi["ContentType"].(string); ok {
		c.Model.ContentType = contentType
	}
	if id, ok := msi["ID"].(string); ok {
		c.ID = NodeID(id)
	}
	if name, ok := msi["Name"].(string); ok {
		c.Name = name
	}
	if name, ok := msi["name"].(string); ok {
		c.Name = name
	}
	if tp, ok := msi["type"].(string); ok {
		c.Type = tp
	}
	if bypass, ok := msi["bypass"].(bool); ok {
		c.Bypass = bypass
	}
	if wfid, ok := msi["workflow_id"].(string); ok {
		c.WorkflowID = WorkflowID(wfid)
	}
	if enabled, ok := msi["enabled"].(bool); ok {
		c.Enabled = enabled
	}
	if createdAt, ok := msi["CreatedAt"].(string); ok {
		c.Model.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct:"Node", Function: "FromMSI"}.Wrap(err)
		}
	}
	if updatedAt, ok := msi["UpdatedAt"].(string); ok {
		c.Model.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct:"Node", Function: "FromMSI"}.Wrap(err)
		}
	}
	if contentType, ok := msi["ContentType"].(string); ok {
		c.Model.ContentType = contentType
	}
	if c.Type == "ollama_node" {
		params := OllamaNode{}
		params.FromMSI(msi)
		c.Params = params
	}
	if c.Type == "comfy_node" {
		params := ComfyNode{}
		params.FromMSI(msi)
		c.Params = params
	}
	if c.Type == "ssh_node" {
		params := SSHNode{}
		params.FromMSI(msi)
		c.Params = params
	}
	return nil
}
