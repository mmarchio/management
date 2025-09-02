package handlers

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mmarchio/management/config"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

// type DisplayNode struct {
// 	types.Node
// 	Menu
// 	SystemPrompts []types.SystemPrompt
// 	Prompt types.Prompt
// 	Disposition types.Disposition
// 	PromptTemplates []types.PromptTemplate
// 	DisplayType string
// 	List []types.Node
// }

// func (c *DisplayNode) GetSystemPrompts(e echo.Context) error {
// 	entity := types.NewSystemPrompt(nil)
// 	list, err := entity.List(e)
// 	if err != nil {
// 		return err
// 	}
// 	c.SystemPrompts = list
// 	return nil
// }

// func (c *DisplayNode) GetPromptTemplates(e echo.Context) error {
// 	entity := types.NewPromptTemplate(nil)
// 	list, err := entity.List(e)
// 	if err != nil {
// 		return err
// 	}
// 	c.PromptTemplates = list
// 	return nil
// }

// func (c *DisplayNode) New(node types.Node) {
// 	c.Node = node
// }

type DisplayJob struct {
	types.Job
	Menu
	Workflows []types.Workflow
	DisplayType string
	List []types.Job
	Prompt types.Prompt
	Workflow types.Workflow
}

type DisplayPrompt struct {
	types.Prompt
	Menu
	Workflows []types.Workflow
	DisplayType string
	List []types.Prompt
	Debug interface{}
	MSI map[string]interface{}
}

func (c *DisplayPrompt) Init(e echo.Context, mode string) error {
	var err error
	c.Prompt = types.NewPrompt(nil)
	c.Menu = Menu{
		Href: "prompts",
		Title: "Prompt",
	}
	c.DisplayType = mode
	switch c.DisplayType {
	case "edit":
		if id := e.Param("id"); id != "" {
			c.Prompt = types.NewPrompt(&id)
			if err := c.Prompt.Get(e); err != nil {
				return merrors.ContentGetError{}.Wrap(err).Log()
			}
			wf := types.NewWorkflow(nil)
			c.Workflows, err = wf.List(e)
			if err != nil {
				return merrors.ContentListError{}.Wrap(err).Log()
			}
		}
	case "list":
		c.List, err = c.Prompt.List(e)
		if err != nil {
			return merrors.ContentListError{}.Wrap(err).Log()
		}
	default:
	}
	c.ToMSI()
	return nil
}

func (c *DisplayPrompt) ToMSI() error {
	b, err := json.Marshal(c)
	if err != nil {
		return merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	if err := json.Unmarshal(b, c.MSI); err != nil {
		return merrors.JSONUnmarshallingError{}.Wrap(err).Log()
	}
	return nil
}

type DisplayComfyUITemplate struct {
	types.ComfyUITemplate
	Menu
	DisplayType string
	List []types.ComfyUITemplate
}

type DisplayDisposition struct {
	types.Disposition
	Menu
	DisplayType string
	List []types.Disposition
}

type DisplayJobRun struct {
	types.JobRun
	Menu
	DisplayType string
	List []types.JobRun
	Workflow types.Workflow
	Steps map[string]string 
	Nodes map[string]string
}

type DisplaySystemPrompt struct {
	types.SystemPrompt
	Menu
	DisplayType string
	List []types.SystemPrompt
	MSI map[string]interface{}
}

func (c *DisplaySystemPrompt) Init(e echo.Context, mode string) error {
	var err error
	c.Menu = Menu{
		Href: "systemprompts",
		Title: "System Prompt",
	}
	c.SystemPrompt = types.NewSystemPrompt(nil)
	c.DisplayType = mode
	if mode == "list" {
		c.List, err = c.SystemPrompt.List(e)
		if err != nil {
			return merrors.ContentListError{CalledBy: "handlers.DisplaySystemPrompt.Init"}.Wrap(err).Log()
		}
	}
	if mode == "edit" {
		if id := e.Param("id"); id != "" {
			c.SystemPrompt = types.NewSystemPrompt(&id)
			if err := c.SystemPrompt.Get(e); err != nil {
				return merrors.ContentGetError{CalledBy: "handlers.DisplaySystemPrompt.Init"}.Wrap(err).Log()
			}
		}
	}
	if err := c.ToMSI(); err != nil {
		return merrors.MSIConversionError{CalledBy: "handlers.DisplaySystemPrompt.Init"}.Wrap(err).Log()
	}
	return nil
}

func (c *DisplaySystemPrompt) ToMSI() error {
	b, err := json.Marshal(c)
	if err != nil {
		return merrors.JSONMarshallingError{CalledBy: "handlers.DisplaySystemPrompt.ToMSI"}.Wrap(err).Log()
	}
	msi := make(map[string]interface{})
	if err := json.Unmarshal(b, &msi); err != nil {
		return merrors.JSONUnmarshallingError{CalledBy: "handlers.DisplaySystemPrompt.ToMSI"}.Wrap(err).Log()
	}
	c.MSI = msi
	return nil
}

type DisplayWorkflow struct {
	types.Workflow
	Menu
	DisplayType string
	List []types.Workflow
	ComfyNodes []types.ComfyNode
	OllamaNodes []types.OllamaNode
	SSHNodes []types.SSHNode
	MSI map[string]interface{}
}

func (c *DisplayWorkflow) Init(e echo.Context, mode string) error {
	var err error
	c.Menu = Menu{
		Href: "workflow",
		Title: "Workflow",
	}
	c.Workflow = types.NewWorkflow(nil)
	c.DisplayType = mode
	switch c.DisplayType {
	case "list":
		c.List, err = c.Workflow.List(e);

		if err != nil {
			return merrors.ContentListError{CalledBy: "handlers.DisplayWorkflow.Init"}.Wrap(err).Log()
		}
	case "edit":
		if id := e.Param("id"); id != "" {
			c.Workflow = types.NewWorkflow(&id)
			if err := c.Workflow.Get(e); err != nil {
				return merrors.ContentGetError{CalledBy: "handlers.DisplayWorkflow.Init"}.Wrap(err).Log()
			}
			cn := types.NewComfyNode(nil)
			c.ComfyNodes, err = cn.ListBy(e, "workflow_id", c.Model.ID)
			if err != nil {
				return merrors.ContentListError{CalledBy: "handlers.DisplayWorkflow.Init"}.Wrap(err).Log()
			}
			on := types.NewOllamaNode(nil)
			c.OllamaNodes, err = on.ListBy(e, "workflow_id", c.Model.ID)
			if err != nil {
				return merrors.ContentListError{CalledBy: "handlers.DisplayWorkflow.Init"}.Wrap(err).Log()
			}
			sn := types.NewSSHNode(nil)
			c.SSHNodes, err = sn.ListBy(e, "workflow_id", c.Model.ID)
			if err != nil {
				return merrors.ContentListError{CalledBy: "handlers.DisplayWorkflow.Init"}.Wrap(err).Log()
			}
			c.Workflow.NodeCleanup(e)
		}
	default:
	}
	if err := c.ToMSI(); err != nil {
		return merrors.MSIConversionError{CalledBy: "handlers.DisplayWorkflow.Init"}.Wrap(err).Log()
	}
	return nil
}

func (c *DisplayWorkflow) ToMSI() error {
	b, err := json.Marshal(c)
	if err != nil {
		return merrors.JSONMarshallingError{CalledBy: "handlers.DisplayWorkflow.ToMSI"}.Wrap(err).Log()
	}
	if err := json.Unmarshal(b, &c.MSI); err != nil {
		return merrors.JSONUnmarshallingError{CalledBy: "handlers.DisplayWorkflow.ToMSI"}.Wrap(err).Log()
	}
	return nil
}

type DisplayPromptTemplate struct {
	types.PromptTemplate
	Menu
	DisplayType string
	List []types.PromptTemplate
	Context string
}

type DisplayOllamaNode struct {
	types.OllamaNode
	Menu
	DisplayType string
	List []types.OllamaNode
	Enabled types.Toggle
	Bypass types.Toggle
	SystemPrompts []types.SystemPrompt
	PromptTemplates []types.PromptTemplate
	MSI map[string]interface{}
}

func (c *DisplayOllamaNode) Init(e echo.Context, mode string) error {
	var err error
	c.OllamaNode = types.NewOllamaNode(nil)
	c.Menu = Menu{
		Href: "node",
		Title: "Node",
	}
	c.DisplayType = mode
	c.Enabled = types.Toggle{
		NamePrefix: "ollamanode_",
		IdPrefix: "ollamanode_",
		Suffix: "enabled",
		Title: "Enabled",
	}
	c.Bypass = types.Toggle{
		NamePrefix: "ollamanode_",
		IdPrefix: "ollamanode_",
		Suffix: "bypass",
		Title: "bypass",
	}
	sp := types.NewSystemPrompt(nil)
	c.SystemPrompts, err = sp.List(e)
	if err != nil {
		return merrors.ContentListError{}.Wrap(err).Log()
	}
	pt := types.NewPromptTemplate(nil)
	c.PromptTemplates, err = pt.List(e)
	if err != nil {
		return merrors.ContentListError{}.Wrap(err).Log()
	}
	if mode == "list" {
		c.List, err = c.OllamaNode.List(e)
		if err != nil {
			return merrors.ContentListError{}.Wrap(err).Log()
		}
	}
	if mode == "edit" {
		if id := e.Param("id"); id != "" {
			c.OllamaNode = types.NewOllamaNode(&id)
			if err := c.OllamaNode.Get(e); err != nil {
				return merrors.ContentGetError{}.Wrap(err).Log()
			}
		}
	}
	if err := c.ToMSI(); err != nil {
		return err
	}
	return nil
}

func (c *DisplayOllamaNode) ToMSI() error {
	b, err := json.Marshal(c)
	if err != nil {
		return merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	msi := make(map[string]interface{})
	if err := json.Unmarshal(b, &msi); err != nil {
		return merrors.JSONUnmarshallingError{}.Wrap(err).Log()
	}
	c.MSI = msi
	return nil
}

type Menu struct {
	Href string
	Title string
}

type DisplayComfyNode struct {
	types.ComfyNode
	Menu
	DisplayType string
	List []types.ComfyNode
	Enabled types.Toggle
	Bypass types.Toggle
	WorkflowID string
	Services []ComfyService
	MSI map[string]interface{}
}

func (c *DisplayComfyNode) ToMSI() error {
	b, err := json.Marshal(c)
	if err != nil {
		return merrors.JSONMarshallingError{CalledBy: "DisplayComfyNode.ToMSI"}.Wrap(err).Log()
	}
	msi := make(map[string]interface{})
	if err := json.Unmarshal(b, &msi); err != nil {
		return merrors.JSONUnmarshallingError{CalledBy: "DisplayComfyNode.ToMSI"}.Wrap(err).Log()
	}
	c.MSI = msi
	return nil
}

type ComfyService struct {
	Port int64
	Title string
}

func (c *DisplayComfyNode) Init(e echo.Context, mode string) error {
	var err error
	c.ComfyNode = types.NewComfyNode(nil)
	c.Menu = Menu{
		Href: "node/comfy",
		Title: "Comfy Node",
	}
	c.DisplayType = mode
	c.Enabled = types.Toggle{
		NamePrefix: "comfynode",
		IdPrefix: "comfynode",
		Suffix: "_enabled",
		Title: "Enabled",
	}
	c.Bypass = types.Toggle{
		NamePrefix: "comfynode",
		IdPrefix: "comfynode",
		Suffix: "_enabled",
		Title: "Bypass",
	}
	c.Services = make([]ComfyService, 0)
	c.Services = append(c.Services, ComfyService{
		Port: int64(config.ComfyUIClassifyImagePort),
		Title: "classify image",
	})
	c.Services = append(c.Services, ComfyService{
		Port: int64(config.ComfyUIGenerateAudioPort),
		Title: "generate audio",
	})
	c.Services = append(c.Services, ComfyService{
		Port: int64(config.ComfyUIGenerateImagePort),
		Title: "generate image",
	})
	c.Services = append(c.Services, ComfyService{
		Port: int64(config.ComfyUIGenerateLipsyncPort),
		Title: "generate lipsync",
	})
	if wfid := e.Param("workflowid"); wfid != "" {
		c.WorkflowID = wfid
	}
	GetLogger(3).Flogger("displaycomfynode.workflowid: %s", c.WorkflowID)
	switch c.DisplayType {
	case "list":
		c.List, err = c.ComfyNode.List(e)
		if err != nil {
			return merrors.ContentListError{CalledBy: "DisplayComfyNode.Init"}.Wrap(err).Log()
		}
	case "edit":
		if id := e.Param("id"); id != "" {
			c.ComfyNode = types.NewComfyNode(&id)
			if err = c.ComfyNode.Get(e); err != nil {
				return merrors.ContentGetError{CalledBy: "DisplayComfyNode.Init"}.Wrap(err).Log()
			}
		}
	default:
	}
	c.ToMSI()
	return nil
}

type DisplaySSHNode struct {
	types.SSHNode
	Menu
	DisplayType string
	List []types.SSHNode
	Enabled types.Toggle
	Bypass types.Toggle
}

func (c *DisplaySSHNode) Init(e echo.Context, mode string) error {
	var err error
	c.Menu = Menu{
		Href: "ssh",
		Title: "SSH Node",
	}
	c.SSHNode = types.NewSSHNode(nil)
	c.DisplayType = mode
	c.Enabled = types.Toggle{
		NamePrefix: "sshnode",
		IdPrefix: "sshnode",
		Suffix: "_enabled",
		Title: "Enabled",
	}
	c.Bypass = types.Toggle{
		NamePrefix: "sshnode",
		IdPrefix: "sshnode",
		Suffix: "_enabled",
		Title: "Bypass",
	}
	switch mode {
	case "none":
	case "new":
	case "edit":
		if id := e.Param("id"); id != "" {
			c.SSHNode = types.NewSSHNode(&id)
			if err := c.SSHNode.Get(e); err != nil {
				return merrors.ContentGetError{}.Wrap(err).Log()
			}
			c.Enabled.Value = c.SSHNode.Enabled
			c.Bypass.Value = c.SSHNode.Bypass
		}
	case "list":
		c.List, err = c.SSHNode.List(e);
		if err != nil {
			return merrors.ContentListError{}.Wrap(err).Log()
		}
	default:
	}
	return nil
}

type DisplayStep struct {
	types.Step
	Menu
	DisplayType string
	List []StepListItem
	Enabled types.Toggle
	Bypass types.Toggle
	Dispositions []types.Disposition
	Workflows []types.Workflow
	SystemPrompts []types.SystemPrompt
	PromptTemplates []types.PromptTemplate
	Nodes []Node
	Dependencies []types.Step
}

type Node struct {
	ID string
	Name string
	Type string
	SystemPrompt string
	PromptTemplate string
}

func (c *DisplayStep) Init(e echo.Context, mode string) error {
	var err error
	c.Menu = Menu{
		Href: "step",
		Title: "Step",
	}
	c.DisplayType = mode
	c.Enabled = types.Toggle{
		NamePrefix: "step_",
		IdPrefix: "step_",
		Suffix: "enabled",
		Title: "Enabled",
	}
	c.Bypass = types.Toggle{
		NamePrefix: "step_",
		IdPrefix: "step_",
		Suffix: "bypass",
		Title: "bypass",
	}
	t := types.Step{}
	t.Model.ContentType = "step"
	c.Dependencies, err = t.List(e, "order desc")
	if err != nil {
		return merrors.ContentListError{}.Wrap(err)
	}
	
	d := types.NewDisposition(nil)
	c.Dispositions, err = d.List(e)
	if err != nil {
		return merrors.ContentListError{}.Wrap(err).Log()
	}
	w := types.NewWorkflow(nil)
	c.Workflows, err = w.List(e)
	if err != nil {
		return merrors.ContentListError{}.Wrap(err).Log()
	}
	sp := types.NewSystemPrompt(nil)
	c.SystemPrompts, err = sp.List(e)
	if err != nil {
		return merrors.ContentListError{}.Wrap(err).Log()
	}
	pt := types.NewPromptTemplate(nil)
	c.PromptTemplates, err = pt.List(e)
	if err != nil {
		return merrors.ContentListError{}.Wrap(err).Log()
	}
	nodes := make([]Node, 0)
	cn := types.NewComfyNode(nil)
	cns, err := cn.List(e)
	if err != nil {
		return merrors.ContentListError{}.Wrap(err).Log()
	}
	for _, node := range cns {
		nodes = append(nodes, Node{
			ID: node.ID,
			Name: node.Name,
			Type: node.Model.ContentType,
		})
	}
	on := types.NewOllamaNode(nil)
	ons, err := on.List(e)
	if err != nil {
		return merrors.ContentListError{}.Wrap(err).Log()
	}
	for _, node := range ons {
		nodes = append(nodes, Node{
			ID: node.ID,
			Name: node.Name,
			Type: node.Model.ContentType,
		})
	}
	sn := types.NewSSHNode(nil)
	sns, err := sn.List(e)
	if err != nil {
		return merrors.ContentListError{}.Wrap(err).Log()
	}
	for _, node := range sns {
		nodes = append(nodes, Node{
			ID: node.ID,
			Name: node.Name,
			Type: node.Model.ContentType,
		})
	}
	c.Nodes = nodes
	dep := types.Step{}
	dep.Model.ID = "nil"
	c.Dependency = &dep
	switch mode {
	case "list":
		items := make([]StepListItem, 0)
		for _, item := range c.Dependencies {
			sli := StepListItem{}
			sli.Step = item
			if item.WorkflowID != "" {
				wfid := item.WorkflowID.String()
				wf := types.NewWorkflow(&wfid)
				if err := wf.Get(e); err != nil {
					return merrors.ContentGetError{}.Wrap(err).Log()
				}
				sli.WorkflowModel = wf
			}
			if !item.DispositionID.IsNil() {
				dpid := item.DispositionID.String()
				dp := types.NewDisposition(&dpid)
				if err := dp.Get(e); err != nil {
					return merrors.ContentGetError{}.Wrap(err).Log()
				}
				sli.DispositionModel = dp
			}
			if item.Node != "" {
				content := types.Content{}
				content.Model.ID = item.Node
				content, err := content.Get(e)
				if err != nil {
					return merrors.ContentGetError{}.Wrap(err).Log()
				}
				switch content.ContentType {
				case "comfynode":
					n := types.NewComfyNode(&item.Node)
					if err := n.Get(e); err != nil {
						return merrors.ContentGetError{}.Wrap(err).Log()
					}
					sli.NodeModel = Node{
						ID: n.Model.ID,
						Name: n.Name,
						Type: n.Model.ContentType,
					}
				case "ollamanode":
					n := types.NewOllamaNode(&item.Node)
					if err := n.Get(e); err != nil {
						return merrors.ContentGetError{}.Wrap(err).Log()
					}
					sli.NodeModel = Node{
						ID: n.Model.ID,
						Name: n.Name,
						Type: n.Model.ContentType,
						SystemPrompt: n.SystemPrompt,
						PromptTemplate: n.PromptTemplate,
					}
				case "sshnode":
					n := types.NewSSHNode(&item.Node)
					if err := n.Get(e); err != nil {
						return merrors.ContentGetError{}.Wrap(err).Log()
					}
					sli.NodeModel = Node{
						ID: n.Model.ID,
						Name: n.Name,
						Type: n.Model.ContentType,
					}
				default:
				}
				if sli.NodeModel.SystemPrompt != "" {
					_, err = uuid.Parse(sli.NodeModel.SystemPrompt)
					if err == nil {
						spid := sli.NodeModel.SystemPrompt
						sp := types.NewSystemPrompt(&spid)
						if err := sp.Get(e); err != nil {
							return merrors.ContentGetError{}.Wrap(err).Log()
						}
						sli.NodeModel.SystemPrompt = sp.Name
					}
				}
				if sli.NodeModel.PromptTemplate != "" {
					_, err = uuid.Parse(sli.NodeModel.PromptTemplate)
					if err == nil {
						ptid := sli.NodeModel.PromptTemplate
						pt := types.NewPromptTemplate(&ptid)
						if err := pt.Get(e); err != nil {
							return merrors.ContentGetError{}.Wrap(err).Log()
						}
						sli.NodeModel.PromptTemplate = pt.Name
					}
				}
			}
			items = append(items, sli)
		}
		c.List = items
		sorted := make([]StepListItem, len(c.List))
		for i := 0; i<len(c.List); i++ {
			sorted[c.List[i].Order-1] = c.List[i]
		}
		c.List = sorted
	case "new":
		t := types.NewStep(nil)
		t.Dependency = &dep
		c.Step = t
	case "edit":
		if id := e.Param("id"); id != "" {
			t := types.NewStep(&id)
			if err := t.Get(e); err != nil {
				return merrors.ContentGetError{}.Wrap(err)
			}
			if t.Dependency == nil {
				t.Dependency = &dep
			}
			c.Step = t
			c.Enabled.Value = c.Step.Enabled.Value
			c.Bypass.Value = c.Step.Bypass.Value
		}
	}
	return nil
}

type StepListItem struct {
	types.Step
	DispositionModel types.Disposition
	WorkflowModel types.Workflow
	SystemPromptModel types.SystemPrompt
	PromptTemplateModel types.PromptTemplate
	NodeModel Node
}