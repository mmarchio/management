package types

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type ShallowWorkflow struct {
	ShallowModel
	ID 						WorkflowID 		`form: "id" json:"id"`
	Name 					string 			`form: "name" json:"name"`
	ComfyNodesArrayModel 	[]string 	`form: "comfy_nodes" json: "comfy_nodes_array_model"`
	OllamaNodesArrayModel 	[]string 	`form: "ollama_nodes" json: "ollama_nodes_array_model"`
	SSHNodesArrayModel 		[]string 		`form: "ssh_nodes" json: "ssh_nodes_array_model"`
	NodeOrder 				map[string]int 	`form: "node_order" "json: "node_order"`
}

func (c *ShallowWorkflow) Validate() {
	var err error
	valid := true
	if !c.ShallowModel.Validate() {
		fmt.Printf("types.shallowworkflow.model is not valid\n")
		valid = false
	}
	if c.ID.IsNil() || c.ID.String() != c.ShallowModel.ID {
		fmt.Printf("types.shallowworkflow.id does not match model")
		valid = false
	}
	if c.Name == "" {
		fmt.Printf("types.shallowworkflow.name is nil")
		valid = false
	}
	for _, cn := range c.ComfyNodesArrayModel {
		_, err = uuid.Parse(cn)
		if err != nil {
			valid = false
		}
	}
	for _, on := range c.OllamaNodesArrayModel {
		_, err = uuid.Parse(on)
		if err != nil {
			valid = false
		}
	}
	for _, sn := range c.SSHNodesArrayModel {
		_, err = uuid.Parse(sn)
		if err != nil {
			valid = false
		}
	}
	c.ShallowModel.Validated = valid
}

func NewShallowWorkflow(id *string) ShallowWorkflow {
	c := ShallowWorkflow{}
	c.New(id)
	c.ShallowModel.ContentType = "shallowworkflow"
	c, _ = ValidateShallowWorkflow(c)
	return c
} 

func NewShallowWorkflowModelContent() models.ShallowContent {
	c := models.ShallowContent{}
	c.ShallowModel.ContentType = "shallowworkflow"
	return c
}

func NewShallowWorkflowTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "shallowworkflow"
	return c
}

func (c *ShallowWorkflow) New(id *string) {
	c.ID = c.ID.New(id)
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = c.ID.String()
	}
	c.ShallowModel.CreatedAt = time.Now()
	c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
}


func (c ShallowWorkflow) List(ctx context.Context) ([]ShallowWorkflow, error) {
	content := NewShallowWorkflowModelContent()
	content.ShallowModel.ContentType = "shallowworkflow"
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.WorkflowListError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	cuts := make([]ShallowWorkflow, 0)
	for _, model := range contents {
		cut := NewShallowWorkflow(&model.ShallowModel.ID)
		if err := json.Unmarshal([]byte(model.Content), &cut); err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowWorkflow", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowWorkflow) ListBy(ctx context.Context, key string, value interface{}) ([]ShallowWorkflow, error) {
	content := NewShallowWorkflowModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.WorkflowListError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	cuts := make([]ShallowWorkflow, 0)
	for _, model := range contents {
		cut := ShallowWorkflow{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowWorkflow", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *ShallowWorkflow) Get(ctx context.Context) error {
	fmt.Printf("types:shallowworkflow:get model.id: %s id: %s\n", c.ShallowModel.ID, c.ID.String())
	content := NewShallowWorkflowTypeContent()
	content.ShallowModel.ID = c.ShallowModel.ID
	content.ID = c.ShallowModel.ID
	content.ShallowModel.ContentType = "shallowworkflow"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.WorkflowGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	if err := json.Unmarshal([]byte(content.Content), c); err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "ShallowWorkflow", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c ShallowWorkflow) Set(ctx context.Context) error {
	c.Validate()
	if !c.ShallowModel.Validated {
		return merrors.WorkflowValidationError{Package: "types", Struct: "shallowworkflow", Function: "set"}.Wrap(fmt.Errorf("validation failed"))
	}
	content := NewShallowWorkflowTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ID.String()
	err := content.Set(ctx)
	if err != nil {
		return merrors.WorkflowSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowWorkflow) Delete(ctx context.Context) error {
	content := NewShallowWorkflowTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.WorkflowDeleteError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowWorkflow) GetID() string {
	return c.ShallowModel.ID
}

func (c ShallowWorkflow) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowWorkflow) GetTable() string {
	return c.ShallowModel.Table
}

func (c ShallowWorkflow) SetID() (ShallowWorkflow, error) {
	var err error
	c.ID = WorkflowID(c.ShallowModel.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "shallowworkflow"}.Wrap(err)
	}
	return c, nil
}

func ValidateShallowWorkflow(p ShallowWorkflow) (ShallowWorkflow, error) {
	var err error
	return p, err
}

func (c ShallowWorkflow) Bind(e echo.Context) (ShallowWorkflow, error) {
	var err error
	return c, err
}

func (c ShallowWorkflow) Next(e echo.Context, ctx context.Context) (*models.Context, error) {
	systemContext, err := models.Context{}.GetCtx(ctx)
	if err != nil {
		return nil, merrors.ContextGetError{Package: "types", Struct: "ShallowWorkflow", Function: "Next"}.Wrap(err)
	}
	return systemContext, nil
}

func (c *ShallowWorkflow) CutNodeOrder(index string) {
	iv := c.NodeOrder[index]
	toslice := make([]string, len(c.NodeOrder))
	for k, v := range c.NodeOrder {
		toslice[v] = k
	}
	newslice := append(toslice[:iv], toslice[iv:]...)
	tomap := make(map[string]int)
	for k, v := range newslice {
		tomap[v] = k 
	}
	c.NodeOrder = tomap
}

func (c *ShallowWorkflow) CutNode(id string) {
	found := false
	if !found {
		comfynodes := make([]string, 0)
		for _, v := range c.ComfyNodesArrayModel {
			if v == id {
				found = true
				continue
			}
			comfynodes = append(comfynodes, v)
		}
		c.ComfyNodesArrayModel = comfynodes
	}
	if !found {
		ollamanodes := make([]string, 0)
		for _, v := range c.OllamaNodesArrayModel {
			if v == id {
				found = true
				continue
			}
			ollamanodes = append(ollamanodes, v)
		}
		c.OllamaNodesArrayModel = ollamanodes
	}
	if !found {
		sshnodes := make([]string, 0)
		for _, v := range c.SSHNodesArrayModel {
			if v == id {
				found = true
				continue
			}
			sshnodes = append(sshnodes, v)
		}
		c.SSHNodesArrayModel = sshnodes
	}
}