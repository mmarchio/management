package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
)

type Workflow struct {
	Model
	Name 		string 			`form:"name" json:"name"`
	ComfyNodes 	[]ComfyNode 	`form:"comfy_nodes" json:"comfy_nodes"`
	OllamaNodes []OllamaNode 	`form:"ollama_nodes" json:"ollama_nodes"`
	SSHNodes 	[]SSHNode 		`form:"ssh_nodes" json:"ssh_nodes"`
	NodeOrder 	map[string]int 	`form:"node_order" json:"node_order"`
}

type ShallowWorkflow struct {
	ShallowModel
	Name 		string 			`form:"name" json:"name"`
	ComfyNodes 	[]string 		`form:"comfy_nodes" json:"comfy_nodes"`
	OllamaNodes []string 		`form:"ollama_nodes" json:"ollama_nodes"`
	SSHNodes 	[]string 		`form:"ssh_nodes" json:"ssh_nodes"`
	NodeOrder 	map[string]int 	`form:"node_order" json:"node_order"`
}

func (c ShallowWorkflow) Set(e echo.Context, update bool) error {
	content := Content{}
	content.ID = c.ShallowModel.ID
	content.Model.ID = c.ShallowModel.ID
	content.ContentType = c.ShallowModel.ContentType
	content.Model.ContentType = c.ShallowModel.ContentType
	content.Model.CreatedAt = c.ShallowModel.CreatedAt
	content.Model.UpdatedAt = c.ShallowModel.UpdatedAt
	b, err := json.Marshal(c)
	if err != nil {
		return merrors.JSONMarshallingError{Info: c.ShallowModel.ID, Package: "models", Struct: "ShallowWorkflow", Function: "Set"}.Wrap(err).Log()
	}
	content.Content = string(b)
	if err := content.Set(e, update); err != nil {
		return err
	}
	return nil
}

func (c ShallowWorkflow) Get(e echo.Context, mode string) (*Workflow, *ShallowWorkflow, error) {
	content := Content{}
	content.Model.ID = c.ShallowModel.ID
	if err := content.Get(e); err != nil {
		return nil, nil, merrors.ContentGetError{Info: c.ShallowModel.ID, Package: "models", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(err).Log()
	}
	if mode == "shallow" {
		if err := json.Unmarshal([]byte(content.Content), &c); err != nil {
			return nil, nil, merrors.JSONUnmarshallingError{Info: content.Content, Package: "models", Struct: "ShallowOllamaNode", Function: "Get"}.Wrap(err).Log()
		}
		return nil, &c, nil
	}
	if mode == "full" {
		list, err := content.GetIn(e)
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
	return nil, nil, merrors.ContentGetError{Package: "models", Struct: "ShallowWorkflow", Function: "Get"}.New("unknown mode: %s", mode)
} 

func (c ShallowWorkflow) SetName(e echo.Context, id string) error {
	c.Name = id
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(e, true); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) SetComfyNodes(e echo.Context, ids []string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, ids...)
	c.ComfyNodes = append(c.ComfyNodes, ids...)
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(e, true); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) AppendComfyNodes(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.ComfyNodes = append(c.ComfyNodes, id)
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(e, true); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) SetOllamaNodes(e echo.Context, ids []string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, ids...)
	c.OllamaNodes = append(c.OllamaNodes, ids...)
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(e, true); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) AppendOllamaNodes(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.OllamaNodes = append(c.OllamaNodes, id)
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(e, true); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) SetSSHNodes(e echo.Context, ids []string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, ids...)
	c.SSHNodes = append(c.SSHNodes, ids...)
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(e, true); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) AppendSSHNodes(e echo.Context, id string) error {
	c.ShallowModel.Manifest = append(c.ShallowModel.Manifest, id)
	c.SSHNodes = append(c.SSHNodes, id)
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(e, true); err != nil {
			return err
		}
	}
	return nil
}

func (c ShallowWorkflow) SetNodeOrder(e echo.Context, no map[string]int) error {
	c.NodeOrder = no
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(e, true); err != nil {
			return err
		}
	}
	return nil
} 

func (c ShallowWorkflow) SetNodeOrderIndex(e echo.Context, id string, v int) error {
	c.NodeOrder[id] = v
	c.ShallowModel.UpdatedAt = time.Now()
	if c.ShallowModel.ID != "" {
		if err := c.Set(e, true); err != nil {
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
		GetLogger(4).Flogger("workflow model is not valid")
		valid = false
	}
	if c.ID != c.Model.ID {
		GetLogger(4).Flogger("id does not match model")
	}
	if c.Model.ContentType != "workflow" {
		GetLogger(4).Flogger("content type is wrong")
		valid = false
	}
	if c.ID == "" {
		GetLogger(4).Flogger("id is nil")
		valid = false
	}
	if c.Name == "" {
		GetLogger(4).Flogger("name is nil")
		valid = false
	}
	for _, node := range c.ComfyNodes {
		node.Validate()
		if !node.Model.Validated {
			GetLogger(4).Flogger("node: %s is not valid", node.Model.ID)
			valid = false
		}
	}
	for _, node := range c.OllamaNodes {
		node.Validate()
		if !node.Model.Validated {
			GetLogger(4).Flogger("node: %s is not valid\n", node.Model.ID)
			valid = false
		}
	}
	for _, node := range c.SSHNodes {
		node.Validate()
		if !node.Model.Validated {
			GetLogger(4).Flogger("node: %s is not valid\n", node.Model.ID)
			valid = false
		}
	}
	c.Model.Validated = valid
}
