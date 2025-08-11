package types

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

type Workflow struct {
	Model
	ID 						WorkflowID 			`form:"id" json:"id"`
	Name 					string 				`form:"name" json:"name"`
	ComfyNodesArrayModel 	[]ComfyNode 		`form:"comfy_nodes" json:"comfy_nodes_array_model"`
	OllamaNodesArrayModel 	[]OllamaNode 		`form:"ollama_nodes" json:"ollama_nodes_array_model"`
	SSHNodesArrayModel 		[]SSHNode 			`form:"ssh_nodes" json:"ssh_nodes_array_model"`
	NodeOrder 				map[int]NodeOrder 	`form:"node_order" json:"node_order"`
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

func (c Workflow) ToContent() (*Content, error) {
	m := NewWorkflowTypeContent()
	m.Model = c.Model
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(nil, err)
	}
	m.Content = string(b)
	return &m, nil
}

func (c *Workflow) Validate() {
	valid := true
	if !c.Model.Validate() {
		GetLogger().Flogger("types.workflow.model is not valid")
		valid = false
	}
	if c.ID.IsNil() || c.ID.String() != c.Model.ID {
		GetLogger().Flogger("types.workflow.id does not match model")
		valid = false
	}
	if c.Name == "" {
		GetLogger().Flogger("types.workflow.name is nil")
		valid = false
	}
	for _, node := range c.ComfyNodesArrayModel {
		node.Validate()
		if !node.Model.Validated {
			GetLogger().Flogger("types.workflow.node[%s] failed validation", node.Model.ID)
			valid = false
		}
	}
	for _, node := range c.OllamaNodesArrayModel {
		if !node.ValidateV2() {
			GetLogger().Flogger("types.workflow.node[%s] failed validation", node.Model.ID)
			valid = false
		}
	}
	for _, node := range c.SSHNodesArrayModel {
		node.Validate()
		if !node.Model.Validated {
			GetLogger().Flogger("types.workflow.node[%s] failed validation", node.Model.ID)
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

func (c Workflow) List(e echo.Context) ([]Workflow, error) {
	content := NewWorkflowModelContent()
	content.Model.ContentType = "workflow"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(nil, err)
	}
	cuts := make([]Workflow, 0)
	for _, model := range contents {
		cut := NewWorkflow(&model.Model.ID)
		if err := json.Unmarshal([]byte(model.Content), &cut); err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Workflow", Function: "List"}.Wrap(nil, err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c Workflow) ListBy(e echo.Context, key string, value interface{}) ([]Workflow, error) {
	content := NewWorkflowModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(nil, err)
	}
	cuts := make([]Workflow, 0)
	for _, model := range contents {
		cut := Workflow{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Workflow", Function: "ListBy"}.Wrap(nil, err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *Workflow) Get(e echo.Context) error {
	content := NewWorkflowTypeContent()
	content.Model.ID = c.Model.ID
	content.ID = c.Model.ID
	content.Model.ContentType = "workflow"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(nil, err)
	}
	if err := json.Unmarshal([]byte(content.Content), c); err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Workflow", Function: "Get"}.Wrap(nil, err)
	}
	c.NodeCleanup(e)
	return nil
}

func (c Workflow) Set(e echo.Context) error {
	c.Validate()
	if !c.Model.Validated {
		return merrors.ContentValidationError{Package: "types", Struct: "workflow", Function: "set"}.New(nil, "validation failed")
	}
	content := NewWorkflowTypeContent()
	content.FromType(c)
	content.Model.ID = c.ID.String()
	err := content.Set(e)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(nil, err)
	}
	return nil
}

func (c Workflow) Delete(e echo.Context) error {
	content := NewWorkflowTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(nil, err)
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
		return c, merrors.IDSetError{Info: "workflow"}.Wrap(nil, err)
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

func (c Workflow) Next(e echo.Context) (*models.Context, error) {
	systemContext, err := models.Context{}.GetCtx(e)
	if err != nil {
		return nil, merrors.ContextGetError{Package: "types", Struct: "Workflow", Function: "Next"}.Wrap(nil, err)
	}
	return systemContext, nil
}

func (c *Workflow) CutNodeOrder(nodeid string) {
	nodeSlice := make([]NodeOrder, 0)
	newOrder := make(map[int]NodeOrder)
	for _, n := range c.NodeOrder {
		if n.NodeID == nodeid {
			continue
		}
		n.Order = len(nodeSlice)
		nodeSlice = append(nodeSlice, n)
		newOrder[n.Order] = n
	}
	c.NodeOrder = newOrder
}

func (c *Workflow) CutNode(id string) {
	for _, n := range c.NodeOrder {
		if n.NodeID == id {
			switch n.NodeType {
			case "comfynode":
				for i, cn := range c.ComfyNodesArrayModel {
					if cn.ID == id {
						c.ComfyNodesArrayModel = append(c.ComfyNodesArrayModel[:i], c.ComfyNodesArrayModel[i:]...)
						break
					}
				}
			case "ollamanode":
				for i, cn := range c.OllamaNodesArrayModel {
					if cn.ID == id {
						c.OllamaNodesArrayModel = append(c.OllamaNodesArrayModel[:i], c.OllamaNodesArrayModel[i:]...)
						break
					}
				}
			case "sshnode":
				for i, cn := range c.SSHNodesArrayModel {
					if cn.ID == id {
						c.SSHNodesArrayModel = append(c.SSHNodesArrayModel[:i], c.SSHNodesArrayModel[i:]...)
						break
					}
				}
			default:
			}
		}
	}
}

func (c *Workflow) NodeCleanup(e echo.Context) {
	newComfyNodes := make([]ComfyNode, 0)
	for _, node := range c.ComfyNodesArrayModel {
		if err := node.Get(e); err != nil {
			c.CutNodeOrder(node.ID)
			GetLogger().Flogger("err: %s", err.Error())
			continue
		}
		newComfyNodes = append(newComfyNodes, node)
	}
	c.ComfyNodesArrayModel = newComfyNodes
	newOllamaNodes := make([]OllamaNode, 0)
	for _, node := range c.OllamaNodesArrayModel {
		if err := node.Get(e); err != nil {
			c.CutNodeOrder(node.ID)
			GetLogger().Flogger("err: %s", err.Error())
			continue
		}
		newOllamaNodes = append(newOllamaNodes, node)
	}
	c.OllamaNodesArrayModel = newOllamaNodes
	newSSHNodes := make([]SSHNode, 0)
	for _, node := range c.SSHNodesArrayModel {
		if err := node.Get(e); err != nil {
			c.CutNodeOrder(node.ID)
			GetLogger().Flogger("err: %s", err.Error())
		}
		newSSHNodes = append(newSSHNodes, node)
	}
	c.SSHNodesArrayModel = newSSHNodes
}

func (c *Workflow) ComfyNodeAppend(node ComfyNode) {
	if c.ComfyNodesArrayModel == nil {
		c.ComfyNodesArrayModel = make([]ComfyNode, 0)
	}
	c.ComfyNodesArrayModel = append(c.ComfyNodesArrayModel, node)
	if c.NodeOrder == nil {
		c.NodeOrder = make(map[int]NodeOrder)
	}
	n := NodeOrder{
		ID: NodeOrderID(uuid.NewString()),
		WorkflowID: node.WorkflowID,
		NodeID: node.ID,
		NodeType: node.Model.ContentType,
		Order: len(c.NodeOrder),
	}
	n.EmbedModel.ID = n.ID.String()
	n.CreatedAt = time.Now()
	n.UpdatedAt = n.CreatedAt
	n.ContentType = "nodeorder"
	c.NodeOrder[len(c.NodeOrder)] = n
}

func (c *Workflow) OllamaNodeAppend(node OllamaNode) {
	if c.OllamaNodesArrayModel == nil {
		c.OllamaNodesArrayModel = make([]OllamaNode, 0)
	}
	c.OllamaNodesArrayModel = append(c.OllamaNodesArrayModel, node)
	if c.NodeOrder == nil {
		c.NodeOrder = make(map[int]NodeOrder)
	}
	n := NodeOrder{
		ID: NodeOrderID(uuid.NewString()),
		WorkflowID: node.WorkflowID,
		NodeID: node.ID,
		NodeType: node.Model.ContentType,
		Order: len(c.NodeOrder),
	}
	n.EmbedModel.ID = n.ID.String()
	n.CreatedAt = time.Now()
	n.UpdatedAt = n.CreatedAt
	n.ContentType = "nodeorder"
	c.NodeOrder[len(c.NodeOrder)] = n
}

func (c *Workflow) SSHNodeAppend(node SSHNode) {
	if c.SSHNodesArrayModel == nil {
		c.SSHNodesArrayModel = make([]SSHNode, 0)
	}
	c.SSHNodesArrayModel = append(c.SSHNodesArrayModel, node)
	if c.NodeOrder == nil {
		c.NodeOrder = make(map[int]NodeOrder)
	}
	n := NodeOrder{
		ID: NodeOrderID(uuid.NewString()),
		WorkflowID: node.WorkflowID,
		NodeID: node.ID,
		NodeType: node.Model.ContentType,
		Order: len(c.NodeOrder),
	}
	n.EmbedModel.ID = n.ID.String()
	n.CreatedAt = time.Now()
	n.UpdatedAt = n.CreatedAt
	n.ContentType = "nodeorder"
	c.NodeOrder[len(c.NodeOrder)] = n
}