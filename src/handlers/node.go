package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

func RegisterNodesRoutes(e *echo.Echo) {
	g := e.Group("/node")
	g.GET("/comfy/new/:workflowid", HandleComfyNew)
	g.GET("/comfy/edit/:id", HandleComfyEdit)
	g.GET("/comfy/delete/:id", HandleComfyDelete)
	g.GET("/comfy/list", HandleComfyList)
	g.POST("/comfy/save", HandleComfySave)
	g.POST("/comfy/save/:id", HandleComfySave)

	g.GET("/ollama/new/:workflowid", HandleOllamaNew)
	g.GET("/ollama/edit/:id", HandleOllamaEdit)
	g.GET("/ollama/delete/:id", HandleOllamaDelete)
	g.GET("/ollama/list", HandleOllamaList)
	g.POST("/ollama/save", HandleOllamaSave)
	g.POST("/ollama/save/:id", HandleOllamaSave)
	
	g.GET("/ssh/new/:workflowid", HandleSSHNew)
	g.GET("/ssh/edit/:id", HandleSSHEdit)
	g.GET("/ssh/delete/:id", HandleSSHDelete)
	g.GET("/ssh/list", HandleSSHList)
	g.POST("/ssh/save", HandleSSHSave)
	g.POST("/ssh/save/:id", HandleSSHSave)
}

func HandleComfyNew(c echo.Context) error {
	GetLogger().Flogger("HandleComfyNew called")
	if workflowid := c.Param("workflowid"); workflowid != "" {
		dt := DisplayComfyNode{
			ComfyNode: types.NewComfyNode(nil),
			DisplayType: "new",
			Enabled: types.Toggle{
				NamePrefix: "comfynode_",
				IdPrefix: "comfynode_",
				Suffix: "enabled",
				Title: "Enabled",
			},
			Bypass: types.Toggle{
				NamePrefix: "comfynode_",
				IdPrefix: "comfynode_",
				Suffix: "bypass",
				Title: "Bypass",
			},
		}
		dt.ComfyNode.WorkflowID = types.WorkflowID(workflowid)
		dt.Menu.Href = "nodes/comfy"
		dt.Menu.Title = "Comfy Nodes"

		return c.Render(http.StatusOK, "node.comfy.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleComfyEdit(c echo.Context) error {
	GetLogger().Flogger("HandleComfyEdit called")
	dt := DisplayComfyNode{
		DisplayType: "edit",
	}
	dt.Menu.Href = "nodes/comfy"
	dt.Menu.Title = "Comfy Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewComfyNode(&id)
		if err := cn.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt.ComfyNode = cn
	}
	return c.Render(http.StatusOK, "node.comfy.tpl", dt)
}

func HandleComfySave(c echo.Context) error {
	GetLogger().Flogger("HandleComfySave called")
	dt := DisplayComfyNode{
		DisplayType: "edit",
	}
	dt.Menu.Href = "nodes/comfy"
	dt.Menu.Title = "Comfy Nodes"
	var cn types.ComfyNode
	if id := c.Param("id"); id != "" {
		cn = types.NewComfyNode(&id)
		if err := cn.Get(c); err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
			}
		}
	} else {
		cn = types.NewComfyNode(nil)
	}
	if err := c.Bind(&cn); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleComfySave"}.Wrap(nil, err))
	}
	if err := cn.Set(c); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if !cn.WorkflowID.IsNil() {
		wfid := cn.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		if err := wf.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if wf.ComfyNodesArrayModel == nil {
			wf.ComfyNodesArrayModel = make([]types.ComfyNode, 0)
		}
		wf.ComfyNodesArrayModel = append(wf.ComfyNodesArrayModel, cn)
		if err := wf.Set(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	dt.ComfyNode = cn
	return c.Render(http.StatusCreated, "node.comfy.tpl", dt)
}

func HandleComfyDelete(c echo.Context) error {
	GetLogger().Flogger("HandleComfyDelete called")
	dt := DisplayComfyNode{
		DisplayType: "list",
	}
	dt.Menu.Href = "nodes/comfy"
	dt.Menu.Title = "Comfy Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewComfyNode(&id)
		if err := cn.Delete(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleComfyList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleComfyList(c echo.Context) error {
	GetLogger().Flogger("HandleComfyList called")
	dt := DisplayComfyNode{
		DisplayType: "list",
	}
	dt.Menu.Href = "nodes/comfy"
	dt.Menu.Title = "Comfy Nodes"
	cn := types.NewComfyNode(nil)
	list, err := cn.List(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt.List = list
	return c.Render(http.StatusOK, "node.comfy.tpl", dt)
}

func HandleOllamaNew(c echo.Context) error {
	GetLogger().Flogger("HandleOllamaNew called")
	if workflowid := c.Param("workflowid"); workflowid != "" {
		wf := types.NewWorkflow(&workflowid)
		if err := wf.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		
		dt := DisplayOllamaNode{
			OllamaNode: types.NewOllamaNode(nil),
			DisplayType: "new",
			Enabled: types.Toggle{
				NamePrefix: "ollamanode_",
				IdPrefix: "ollamanode_",
				Suffix: "enabled",
				Title: "Enabled",
			},
			Bypass: types.Toggle{
				NamePrefix: "ollamanode_",
				IdPrefix: "ollamanode_",
				Suffix: "bypass",
				Title: "Bypass",
			},
		}
		dt.Enabled.Value = dt.OllamaNode.Enabled
		dt.Bypass.Value = dt.OllamaNode.Bypass
		sp := types.SystemPrompt{}
		sps, err := sp.List(c)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		pt := types.PromptTemplate{}
		pts, err := pt.List(c)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt.SystemPrompts = sps
		dt.PromptTemplates = pts
		dt.OllamaNode.WorkflowID = types.WorkflowID(workflowid)
		dt.Menu.Href = "nodes/ollama"
		dt.Menu.Title = "Ollama Nodes"
		return c.Render(http.StatusOK, "node.ollama.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleOllamaEdit(c echo.Context) error {
	GetLogger().Flogger("HandleOllamaEdit called")
	dt := DisplayOllamaNode{
		DisplayType: "edit",
		Enabled: types.Toggle{
			NamePrefix: "ollamanode_",
			IdPrefix: "ollamanode_",
			Suffix: "enabled",
			Title: "Enabled",
		},
		Bypass: types.Toggle{
			NamePrefix: "ollamanode_",
			IdPrefix: "ollamanode_",
			Suffix: "bypass",
			Title: "Bypass",
		},
	}
	sp := types.SystemPrompt{}
	sps, err := sp.List(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	pt := types.PromptTemplate{}
	pts, err := pt.List(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt.SystemPrompts = sps
	dt.PromptTemplates = pts
	dt.Menu.Href = "nodes/ollama"
	dt.Menu.Title = "Ollama Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewOllamaNode(&id)
		if err := cn.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt.OllamaNode = cn
		dt.Enabled.Value = cn.Enabled
		dt.Bypass.Value = cn.Bypass
	}
	return c.Render(http.StatusOK, "node.ollama.tpl", dt)
}

func HandleOllamaSave(c echo.Context) error {
	GetLogger().Flogger("HandleOllamaSave called")
	dt := DisplayOllamaNode{
		DisplayType: "edit",
	}
	dt.Menu.Href = "nodes/ollama"
	dt.Menu.Title = "Ollama Nodes"
	var cn types.OllamaNode
	if id := c.Param("id"); id != "" {
		cn = types.NewOllamaNode(&id)
		if err := cn.Get(c); err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
			}
		}
	} else {
		cn = types.NewOllamaNode(nil)
	}
	if err := c.Bind(&cn); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleOllamaSave"}.Wrap(nil, err))
	}
	if err := cn.Set(c); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if !cn.WorkflowID.IsNil() {
		wfid := cn.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		if err := wf.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wf.OllamaNodeAppend(cn)
		if err := wf.Set(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	dt.OllamaNode = cn
	return HandleWorkflowList(c)
}

func HandleOllamaDelete(c echo.Context) error {
	GetLogger().Flogger("HandleOllamaDelete called")
	dt := DisplayOllamaNode{
		DisplayType: "list",
	}
	dt.Menu.Href = "nodes/ollama"
	dt.Menu.Title = "Ollama Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewOllamaNode(&id)
		if err := cn.Delete(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleOllamaList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleOllamaList(c echo.Context) error {
	GetLogger().Flogger("HandleOllamaList called")
	dt := DisplayOllamaNode{
		DisplayType: "list",
	}
	dt.Menu.Href = "nodes/ollama"
	dt.Menu.Title = "Ollama Nodes"
	cn := types.NewOllamaNode(nil)
	list, err := cn.List(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt.List = list
	return c.Render(http.StatusOK, "node.ollama.tpl", dt)
}

func HandleSSHNew(c echo.Context) error {
	GetLogger().Flogger("HandleSSHNew called")
	if workflowid := c.Param("workflowid"); workflowid != "" {
		dt := DisplaySSHNode{
			SSHNode: types.NewSSHNode(nil),
			DisplayType: "new",
			Enabled: types.Toggle{
				NamePrefix: "sshnode_",
				IdPrefix: "sshnode_",
				Suffix: "enabled",
				Title: "Enabled",
			},
			Bypass: types.Toggle{
				NamePrefix: "sshnode_",
				IdPrefix: "sshnode_",
				Suffix: "bypass",
				Title: "Bypass",
			},
		}
		dt.SSHNode.WorkflowID = types.WorkflowID(workflowid)
		dt.Menu.Href = "nodes/ssh"
		dt.Menu.Title = "SSH Nodes"
		return c.Render(http.StatusOK, "node.ssh.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleSSHEdit(c echo.Context) error {
	GetLogger().Flogger("HandleSSHEdit called")
	dt := DisplaySSHNode{
		DisplayType: "edit",
	}
	dt.Menu.Href = "nodes/ssh"
	dt.Menu.Title = "SSH Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewSSHNode(&id)
		if err := cn.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt.SSHNode = cn
	}
	return c.Render(http.StatusOK, "node.ssh.tpl", dt)
}

func HandleSSHSave(c echo.Context) error {
	GetLogger().Flogger("HandleSSHSave called")
	dt := DisplaySSHNode{
		DisplayType: "edit",
	}
	dt.Menu.Href = "nodes/ssh"
	dt.Menu.Title = "SSH Nodes"
	var cn types.SSHNode
	if id := c.Param("id"); id != "" {
		cn = types.NewSSHNode(&id)
		if err := cn.Get(c); err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
			}
		}
	} else {
		cn = types.NewSSHNode(nil)
	}
	if err := c.Bind(&cn); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleComfySave"}.Wrap(nil, err))
	}
	if err := cn.Set(c); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if !cn.WorkflowID.IsNil() {
		wfid := cn.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		if err := wf.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if wf.SSHNodesArrayModel == nil {
			wf.SSHNodesArrayModel = make([]types.SSHNode, 0)
		}
		wf.SSHNodesArrayModel = append(wf.SSHNodesArrayModel, cn)
		if err := wf.Set(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	dt.SSHNode = cn
	return c.Render(http.StatusCreated, "node.ssh.tpl", dt)
}

func HandleSSHDelete(c echo.Context) error {
	GetLogger().Flogger("HandleSSHDelete called")
	dt := DisplaySSHNode{
		DisplayType: "list",
	}
	dt.Menu.Href = "nodes/ssh"
	dt.Menu.Title = "SSH Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewSSHNode(&id)
		if err := cn.Delete(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleSSHList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleSSHList(c echo.Context) error {
	GetLogger().Flogger("HandleSSHList called")
	dt := DisplaySSHNode{
		DisplayType: "list",
	}
	dt.Menu.Href = "nodes/ssh"
	dt.Menu.Title = "SSH Nodes"
	cn := types.NewSSHNode(nil)
	list, err := cn.List(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt.List = list
	return c.Render(http.StatusOK, "node.ssh.tpl", dt)
}
