package models

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	merrors "github.com/mmarchio/management/errors"
)

type Workflow struct {
	Model
	ID 			string 			`form:"id" json:"id"`
	Name 		string 			`form:"name" json:"name"`
	ComfyNodes 	[]ComfyNode 	`form:"comfy_nodes" json:"comfy_nodes"`
	OllamaNodes []OllamaNode 	`form:"ollama_nodes" json:"ollama_nodes"`
	SSHNodes 	[]SSHNode 		`form:"ssh_nodes" json:"ssh_nodes"`
	NodeOrder 	map[string]int 	`form:"node_order" json:"node_order"`
}

type ShallowWorkflow struct {
	ShallowModel
	ID 			string 			`form:"id" json:"id"`
	Name 		string 			`form:"name" json:"name"`
	ComfyNodes 	[]string 		`form:"comfy_nodes" json:"comfy_nodes"`
	OllamaNodes []string 		`form:"ollama_nodes" json:"ollama_nodes"`
	SSHNodes 	[]string 		`form:"ssh_nodes" json:"ssh_nodes"`
	NodeOrder 	map[string]int 	`form:"node_order" json:"node_order"`
}

func (c ShallowWorkflow) Set(ctx context.Context) error {
	content := Content{}
	content.ID = c.ShallowModel.ID
	content.Model.ID = c.ShallowModel.ID
	content.ContentType = c.ShallowModel.ContentType
	content.Model.ContentType = c.ShallowModel.ContentType
	content.Model.CreatedAt = c.ShallowModel.CreatedAt
	content.Model.UpdatedAt = c.ShallowModel.UpdatedAt
	b, err := json.Marshal(c)
	if err != nil {
		return merrors.JSONMarshallingError{Info: c.ShallowModel.ID, Package: "models", Struct: "ShallowWorkflow", Function: "Set"}.Wrap(err)
	}
	content.Content = string(b)
	if err := content.Set(ctx); err != nil {
		return err
	}
	return nil
}

func (c ShallowWorkflow) Get(ctx context.Context, mode string) (*Workflow, *ShallowWorkflow, error) {
	content := Content{ID: c.ShallowModel.ID}
	if err := content.Get(ctx); err != nil {
		return nil, nil, merrors.ShallowWorkflowGetError{Info: c.ShallowModel.ID, Package: "models", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(err)
	}
	if mode == "shallow" {
		if err := json.Unmarshal([]byte(content.Content), &c); err != nil {
			return nil, nil, merrors.JSONUnmarshallingError{Info: content.Content, Package: "models", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(err)
		}
		return nil, &c, nil
	}
	if mode == "full" {
		list, err := content.GetIn(ctx, c.ShallowModel.Manifest)
		if err != nil {
			return nil, nil, err
		}
		mss := make(map[string]string)
		for _, v := range list {
			mss[v.ID] = v.Content
		}
		hydrated, err := HydrateShallowJson(content.Content, mss)
		if err != nil {
			return nil, nil, err
		}
		full := Workflow{}
		if err := json.Unmarshal([]byte(hydrated), &full); err != nil {
			return nil, nil, err
		}
		return &full, nil, nil
	}
	return nil, nil, merrors.ShallowWorkflowGetError{Package: "models", Struct: "ShallowWorkflow", Function: "Get"}.Wrap(fmt.Errorf("unknown mode: %s", mode))
} 

func (c ShallowWorkflow) SetName(ctx context.Context, id string) error {
	c.Name = id
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) SetComfyNodes(ctx context.Context, ids []string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, ids...)
	c.ComfyNodes = append(c.ComfyNodes, ids...)
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) AppendComfyNodes(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.ComfyNodes = append(c.ComfyNodes, id)
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) SetOllamaNodes(ctx context.Context, ids []string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, ids...)
	c.OllamaNodes = append(c.OllamaNodes, ids...)
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) AppendOllamaNodes(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.OllamaNodes = append(c.OllamaNodes, id)
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) SetSSHNodes(ctx context.Context, ids []string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, ids...)
	c.SSHNodes = append(c.SSHNodes, ids...)
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) AppendSSHNodes(ctx context.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.SSHNodes = append(c.SSHNodes, id)
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) SetNodeOrder(ctx context.Context, no map[string]int) error {
	c.NodeOrder = no
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil
} 

func (c ShallowWorkflow) SetNodeOrderIndex(ctx context.Context, id string, v int) error {
	c.NodeOrder[id] = v
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(ctx); err != nil {
			return err
		}
	}
	return nil
} 

func NewShallowWorkflow(id *string) ShallowWorkflow {
	c := ShallowWorkflow{}
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
		c.ShallowModel.CreatedAt = time.Now()
		c.ShallowModel.UpdatedAt = c.CreatedAt
	}
	c.ID = c.ShallowModel.ID
	c.ShallowModel.ContentType = "shallow_workflow"
	return c
}

func (c *Workflow) Validate() {
	valid := true
	if !c.Model.Validate() {
		fmt.Printf("workflow model is not valid\n")
		valid = false
	}
	if c.ID != c.Model.ID {
		fmt.Printf("id does not match model\n")
	}
	if c.Model.ContentType != "workflow" {
		fmt.Printf("content type is wrong\n")
		valid = false
	}
	if c.ID == "" {
		fmt.Printf("id is nil")
		valid = false
	}
	if c.Name == "" {
		fmt.Printf("name is nil")
		valid = false
	}
	for _, node := range c.ComfyNodes {
		node.Validate()
		if !node.Model.Validated {
			fmt.Printf("node: %s is not valid\n", node.Model.ID)
			valid = false
		}
	}
	for _, node := range c.OllamaNodes {
		node.Validate()
		if !node.Model.Validated {
			fmt.Printf("node: %s is not valid\n", node.Model.ID)
			valid = false
		}
	}
	for _, node := range c.SSHNodes {
		node.Validate()
		if !node.Model.Validated {
			fmt.Printf("node: %s is not valid\n", node.Model.ID)
			valid = false
		}
	}
	c.Model.Validated = valid
}
