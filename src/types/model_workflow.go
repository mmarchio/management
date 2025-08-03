package types

import (
	"context"
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type Workflow struct {
	Model
	ID WorkflowID `json:"id"`
	Name string `form:"name" json:"name"`
	Nodes []Node `json:"nodes"`
}

func (c *Workflow) FromMSI(msi map[string]interface{}) error {
	var err error
	if id, ok := msi["id"].(string); ok {
		c.Model.ID = id
	}
	if createdAt, ok := msi["CreatedAt"].(string); ok {
		c.Model.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct:"Workflow", Function: "FromMSI"}.Wrap(err)
		}
	}
	if updatedAt, ok := msi["UpdatedAt"].(string); ok {
		c.Model.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct:"Workflow", Function: "FromMSI"}.Wrap(err)
		}
	}
	if contentType, ok := msi["ContentType"].(string); ok {
		c.Model.ContentType = contentType
	}
	if id, ok := msi["ID"].(string); ok {
		c.ID = WorkflowID(id)
	}
	if name, ok := msi["Name"].(string); ok {
		c.Name = name
	}
	if name, ok := msi["name"].(string); ok {
		c.Name = name
	}
	if nodes, ok := msi["nodes"].([]interface{}); ok {
		NODES := make([]Node, 0)
		for _, nodeInterface := range nodes {
			NODE := Node{}
			if node, ok := nodeInterface.(map[string]interface{}); ok {
				NODE.FromMSI(node)
			}
			NODES = append(NODES, NODE)
		}
		c.Nodes = NODES
	}
	return nil
}

func NewWorkflow(id *string) Workflow {
	c := Workflow{}
	c.New(id)
	c.Model.ContentType = "workflow"
	c, _ = ValidateWorkflow(c)
	return c
} 

func NewWorkflowModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "workflow"
	return c
}

func NewWorkflowTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "workflow"
	return c
}

func (c *Workflow) New(id *string) {
	c.ID = c.ID.New(id)
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = c.ID.String()
	}
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
}


func (c Workflow) List(ctx context.Context) ([]Workflow, error) {
	content := NewWorkflowModelContent()
	content.Model.ContentType = "workflow"
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.WorkflowListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]Workflow, 0)
	for _, model := range contents {
		cut := NewWorkflow(&model.Model.ID)
		msi := make(map[string]interface{})
		if err := json.Unmarshal([]byte(model.Content), &msi); err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Workflow", Function: "List"}.Wrap(err)
		}
		if err := cut.FromMSI(msi); err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: "from msi", Package: "types", Struct: "Workflow", Function: "List"}.Wrap(err)
		}
		if err != nil {
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c Workflow) ListBy(ctx context.Context, key string, value interface{}) ([]Workflow, error) {
	content := NewWorkflowModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.WorkflowListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]Workflow, 0)
	for _, model := range contents {
		cut := Workflow{}
		msi := make(map[string]interface{})
		err = json.Unmarshal([]byte(model.Content), &msi)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Workflow", Function: "ListBy"}.Wrap(err)
		}
		if err := cut.FromMSI(msi); err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: "from msi", Package: "types", Struct: "Workflow", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *Workflow) Get(ctx context.Context) error {
	content := NewWorkflowTypeContent()
	content.Model.ID = c.Model.ID
	content.ID = c.Model.ID
	content.Model.ContentType = "workflow"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.WorkflowGetError{Info: c.Model.ID}.Wrap(err)
	}
	msi := make(map[string]interface{})
	if err := json.Unmarshal([]byte(content.Content), &msi); err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Workflow", Function: "Get"}.Wrap(err)
	}
	if err := c.FromMSI(msi); err != nil {
		return merrors.JSONUnmarshallingError{Info: "from msi", Package: "types", Struct: "Workflow", Function: "List"}.Wrap(err)
	}
	return nil
}

func (c Workflow) Set(ctx context.Context) error {
	content := NewWorkflowTypeContent()
	content.FromType(c)
	content.Model.ID = c.ID.String()
	err := content.Set(ctx)
	if err != nil {
		return merrors.WorkflowSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c Workflow) Delete(ctx context.Context) error {
	content := NewWorkflowTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.WorkflowDeleteError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c Workflow) GetID() string {
	return c.Model.ID
}

func (c Workflow) GetContentType() string {
	return c.Model.ContentType
}

func (c Workflow) GetTable() string {
	return c.Model.Table
}

func (c Workflow) SetID() (Workflow, error) {
	var err error
	c.ID = WorkflowID(c.Model.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "workflow"}.Wrap(err)
	}
	return c, nil
}

func ValidateWorkflow(p Workflow) (Workflow, error) {
	var err error
	return p, err
}

func (c Workflow) Bind(e echo.Context) (Workflow, error) {
	var err error
	return c, err
}

func (c Workflow) Next(e echo.Context, ctx context.Context) (*models.Context, error) {
	systemContext, err := models.Context{}.GetCtx(ctx)
	if err != nil {
		return nil, merrors.ContextGetError{Package: "types", Struct: "Workflow", Function: "Next"}.Wrap(err)
	}
	return systemContext, nil
}

func (c Workflow) Run(ctx context.Context) error {
	for _, node := range c.Nodes {
		switch node.Params.GetType() {
		case "ollama_node":
		case "comfy_node":
		case "ssh_node":
		default:
		}
	}
	return nil
} 

