package types

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type Workflow struct {
	Model
	ID 						WorkflowID 		`form: "id" json:"id"`
	Name 					string 			`form: "name" json:"name"`
	ComfyNodesArrayModel 	[]ComfyNode 	`form: "comfy_nodes" json: "comfy_nodes_array_model"`
	OllamaNodesArrayModel 	[]OllamaNode 	`form: "ollama_nodes" json: "ollama_nodes_array_model"`
	SSHNodesArrayModel 		[]SSHNode 		`form: "ssh_nodes" json: "ssh_nodes_array_model"`
	NodeOrder 				map[string]int 	`form: "node_order" "json: "node_order"`
}

func (c Workflow) Pack() []shallowmodel {
	sw := ShallowWorkflow{}
	sw.ShallowModel = sw.ShallowModel.FromTypeModel(c.Model)
	sms := make([]shallowmodel, 0)
	for _, sm := range c.ComfyNodesArrayModel {
		sw.ComfyNodesArrayModel = append(sw.ComfyNodesArrayModel, sm.ID)
		sms = append(sms, sm.Pack()...)
	}
	for _, sm := range c.OllamaNodesArrayModel {
		sw.OllamaNodesArrayModel = append(sw.OllamaNodesArrayModel, sm.ID)
		sms = append(sms, sm.Pack()...)
	}
	for _, sm := range c.SSHNodesArrayModel {
		sw.SSHNodesArrayModel = append(sw.SSHNodesArrayModel, sm.ID)
		sms = append(sms, sm.Pack()...)
	}
	sms = append(sms, sw)
	return sms
}

func (c *Workflow) Validate() {
	valid := true
	if !c.Model.Validate() {
		fmt.Printf("types.workflow.model is not valid\n")
		valid = false
	}
	if c.ID.IsNil() || c.ID.String() != c.Model.ID {
		fmt.Printf("types.workflow.id does not match model")
		valid = false
	}
	if c.Name == "" {
		fmt.Printf("types.workflow.name is nil")
		valid = false
	}
	for _, node := range c.ComfyNodesArrayModel {
		node.Validate()
		if !node.Model.Validated {
			fmt.Printf("types.workflow.node[%s] failed validation\n", node.Model.ID)
			valid = false
		}
	}
	for _, node := range c.OllamaNodesArrayModel {
		node.Validate()
		if !node.Model.Validated {
			fmt.Printf("types.workflow.node[%s] failed validation\n", node.Model.ID)
			valid = false
		}
	}
	for _, node := range c.SSHNodesArrayModel {
		node.Validate()
		if !node.Model.Validated {
			fmt.Printf("types.workflow.node[%s] failed validation\n", node.Model.ID)
			valid = false
		}
	}
	c.Model.Validated = valid
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
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]Workflow, 0)
	for _, model := range contents {
		cut := NewWorkflow(&model.Model.ID)
		if err := json.Unmarshal([]byte(model.Content), &cut); err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Workflow", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c Workflow) ListBy(ctx context.Context, key string, value interface{}) ([]Workflow, error) {
	content := NewWorkflowModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]Workflow, 0)
	for _, model := range contents {
		cut := Workflow{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Workflow", Function: "ListBy"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *Workflow) Get(ctx context.Context) error {
	fmt.Printf("types:workflow:get model.id: %s id: %s\n", c.Model.ID, c.ID.String())
	content := NewWorkflowTypeContent()
	content.Model.ID = c.Model.ID
	content.ID = c.Model.ID
	content.Model.ContentType = "workflow"
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err)
	}
	if err := json.Unmarshal([]byte(content.Content), c); err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Workflow", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c Workflow) Set(ctx context.Context) error {
	c.Validate()
	if !c.Model.Validated {
		return merrors.ContentValidationError{Package: "types", Struct: "workflow", Function: "set"}.Wrap(fmt.Errorf("validation failed"))
	}
	content := NewWorkflowTypeContent()
	content.FromType(c)
	content.Model.ID = c.ID.String()
	err := content.Set(ctx)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c Workflow) Delete(ctx context.Context) error {
	content := NewWorkflowTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(err)
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

func (c *Workflow) CutNodeOrder(index string) {
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

func (c *Workflow) CutNode(id string) {
	found := false
	if !found {
		comfynodes := make([]ComfyNode, 0)
		for _, v := range c.ComfyNodesArrayModel {
			if v.Model.ID == id {
				found = true
				continue
			}
			comfynodes = append(comfynodes, v)
		}
		c.ComfyNodesArrayModel = comfynodes
	}
	if !found {
		ollamanodes := make([]OllamaNode, 0)
		for _, v := range c.OllamaNodesArrayModel {
			if v.Model.ID == id {
				found = true
				continue
			}
			ollamanodes = append(ollamanodes, v)
		}
		c.OllamaNodesArrayModel = ollamanodes
	}
	if !found {
		sshnodes := make([]SSHNode, 0)
		for _, v := range c.SSHNodesArrayModel {
			if v.Model.ID == id {
				found = true
				continue
			}
			sshnodes = append(sshnodes, v)
		}
		c.SSHNodesArrayModel = sshnodes
	}
}