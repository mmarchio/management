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
	Name                  string            `form:"name" json:"name"`
	ComfyNodesArrayModel  []ComfyNode       `form:"comfy_nodes" json:"comfy_nodes_array_model"`
	OllamaNodesArrayModel []OllamaNode      `form:"ollama_nodes" json:"ollama_nodes_array_model"`
	SSHNodesArrayModel    []SSHNode         `form:"ssh_nodes" json:"ssh_nodes_array_model"`
	NodeOrder             map[int]NodeOrder `form:"node_order" json:"node_order"`
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
	m := NewWorkflowTypeContent(nil)
	m.Model = c.Model
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	m.Content = string(b)
	return &m, nil
}

func (c *Workflow) Validate() error {
	if !c.Model.Validate() {
		GetLogger(3).Flogger("workflow: %#v", c)
		return merrors.ContentValidationError{CalledBy: "types.Workflow.Validate"}.New("model validation failed").Log()
	}
	if c.Model.IsNil() {
		return merrors.ContentValidationError{CalledBy: "types.Workflow.Validate"}.New("types.workflow.id does not match model").Log()
	}
	if c.Name == "" {
		return merrors.ContentValidationError{CalledBy: "types.Workflow.Validate"}.New("types.workflow.name is nil").Log()
	}
	for _, node := range c.ComfyNodesArrayModel {
		d := node.Validate()
		if !d.Model.Validated {
			return merrors.ContentValidationError{CalledBy: "types.Workflow.Validate"}.New("types.workflow.node[%s] failed validation", node.Model.ID).Log()
		}
	}
	for _, node := range c.OllamaNodesArrayModel {
		if !node.ValidateV2() {
			return merrors.ContentValidationError{CalledBy: "types.Workflow.Validate"}.New("types.workflow.node[%s] failed validation", node.Model.ID).Log()
		}
	}
	for _, node := range c.SSHNodesArrayModel {
		d := node.Validate()
		if !d.Model.Validated {
			return merrors.ContentValidationError{CalledBy: "types.Workflow.Validate"}.New("types.workflow.node[%s] failed validateion", node.Model.ID).Log()
		}
	}
	return nil
}

func NewWorkflow(id *string) Workflow {
	c := Workflow{}
	if id == nil {
	} else {
		idc := *id
		if idc == "new" {
			c.New(nil)		
		} else {
			c.New(id)
		}
	}
	c.Model.ContentType = "workflow"
	c, _ = ValidateWorkflow(c)
	return c
}

func NewWorkflowModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "workflow"
	return c
}

func NewWorkflowTypeContent(id *string) Content {
	c := Content{}
	if id != nil {
		if *id == "new" {
			c.Model.ID = uuid.NewString()
		} else {
			c.Model.ID = *id
		}
	}
	c.Model.ContentType = "workflow"
	return c
}

func (c *Workflow) New(id *string) {
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
}

func (c Workflow) List(e echo.Context) ([]Workflow, error) {
	content := NewWorkflowModelContent()
	content.Model.ContentType = "workflow"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]Workflow, 0)
	for _, model := range contents {
		cut := NewWorkflow(&model.Model.ID)
		if err := json.Unmarshal([]byte(model.Content), &cut); err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Workflow", Function: "List"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c Workflow) ListBy(e echo.Context, key string, value interface{}) ([]Workflow, error) {
	content := NewWorkflowModelContent()
	contents, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]Workflow, 0)
	for _, model := range contents {
		cut := Workflow{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Workflow", Function: "ListBy"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *Workflow) Get(e echo.Context) error {
	content := NewWorkflowTypeContent(&c.Model.ID)
	content.Model.ContentType = "workflow"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	if err := json.Unmarshal([]byte(content.Content), c); err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Workflow", Function: "Get"}.Wrap(err).Log()
	}
	c.NodeCleanup(e)
	return nil
}

func (c Workflow) Set(e echo.Context, update bool) error {
	err := c.Validate()
	if err != nil {
		return err
	}
	content := NewWorkflowTypeContent(nil)
	content.FromType(c, c.Model)
	err = content.Set(e, update)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	return nil
}

func (c Workflow) Delete(e echo.Context) error {
	content := NewWorkflowTypeContent(nil)
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID}.Wrap(err).Log()
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
		return nil, merrors.ContextGetError{Package: "types", Struct: "Workflow", Function: "Next"}.Wrap(err).Log()
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
	comfyTracker := make(map[string]interface{})
	newComfyNodes := make([]ComfyNode, 0)
	for _, node := range c.ComfyNodesArrayModel {
		if _, ok := comfyTracker[node.Model.ID].(string); !ok {
			if err := node.Get(e); err != nil {
				c.CutNodeOrder(node.Model.ID)
				GetLogger(4).Flogger("err: %s", err.Error())
				continue
			}
			newComfyNodes = append(newComfyNodes, node)
			comfyTracker[node.Model.ID] = "true"
		}
	}
	c.ComfyNodesArrayModel = newComfyNodes
	newOllamaNodes := make([]OllamaNode, 0)
	for _, node := range c.OllamaNodesArrayModel {
		if _, ok := comfyTracker[node.Model.ID].(string); !ok {
			if err := node.Get(e); err != nil {
				c.CutNodeOrder(node.ID)
				GetLogger(4).Flogger("err: %s", err.Error())
				continue
			}
			newOllamaNodes = append(newOllamaNodes, node)
			comfyTracker[node.Model.ID] = "true"
		}
	}
	c.OllamaNodesArrayModel = newOllamaNodes
	newSSHNodes := make([]SSHNode, 0)
	for _, node := range c.SSHNodesArrayModel {
		if _, ok := comfyTracker[node.Model.ID]; !ok {
			if err := node.Get(e); err != nil {
				c.CutNodeOrder(node.ID)
				GetLogger(4).Flogger("err: %s", err.Error())
			}
			newSSHNodes = append(newSSHNodes, node)
			comfyTracker[node.Model.ID] = "true"
		}
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
		WorkflowID: node.WorkflowID,
		NodeID:     node.ID,
		NodeType:   node.Model.ContentType,
		Order:      len(c.NodeOrder),
	}
	n.EmbedModel.ID = uuid.NewString()
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
		WorkflowID: node.WorkflowID,
		NodeID:     node.ID,
		NodeType:   node.Model.ContentType,
		Order:      len(c.NodeOrder),
	}
	n.EmbedModel.ID = uuid.NewString()
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
		WorkflowID: node.WorkflowID,
		NodeID:     node.ID,
		NodeType:   node.Model.ContentType,
		Order:      len(c.NodeOrder),
	}
	n.EmbedModel.ID = uuid.NewString()
	n.CreatedAt = time.Now()
	n.UpdatedAt = n.CreatedAt
	n.ContentType = "nodeorder"
	c.NodeOrder[len(c.NodeOrder)] = n
}
