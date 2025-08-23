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
	GetLogger(4).Flogger("HandleComfyNew called")
	if workflowid := c.Param("workflowid"); workflowid != "" {
		dt := DisplayComfyNode{}
		if err := dt.Init(c, "new"); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return c.Render(http.StatusOK, "node.comfy.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleComfyEdit(c echo.Context) error {
	GetLogger(4).Flogger("HandleComfyEdit called")
	dt := DisplayComfyNode{}
	if err := dt.Init(c, "edit"); err != nil {
		c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
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
	GetLogger(4).Flogger("HandleComfySave called")
	var cn types.ComfyNode
	var update bool
	if id := c.Param("id"); id != "" {
		update = true
		ctype := "comfynode"
		valid, err := types.CheckType(c, &id, &ctype)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", merrors.ContentCheckError{CalledBy: "handlers.HandleComfySave"}.Wrap(err).Log().Error())
		}
		if valid {
			cn = types.NewComfyNode(&id)
			if err := cn.Get(c); err != nil {
				if !strings.Contains(err.Error(), "not found") {
					return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
				}
			}
		} else {
			cn = types.NewComfyNode(nil)
		}
	} else {
		cn = types.NewComfyNode(nil)
	}
	if err := c.Bind(&cn); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{CalledBy:"handlers.HandleComfySave"}.Wrap(err).Log().Error())
	}
	if tv := c.FormValue("template_values"); tv != "" {
		cn.TemplateValues = tv
	}

	if wfid := c.Param("workflowid"); wfid != "" {
		ctype := "workflow"
		valid, err := types.CheckType(c, &wfid, &ctype)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if !valid {
			return c.Render(http.StatusInternalServerError, "error.tpl", merrors.ContentCheckError{CalledBy: "handlers.HandleComfySave"}.New("workflowid %s not valid", wfid).Log().Error())
		}
		cn.WorkflowID = types.WorkflowID(wfid)
	} else if wfid := c.Param("id"); wfid != "" {
		ctype := "workflow"
		valid, err := types.CheckType(c, &wfid, &ctype)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if !valid {
			return c.Render(http.StatusInternalServerError, "error.tpl", merrors.ContentCheckError{CalledBy: "handlers.HandleComfySave"}.New("workflowid %s not valid", wfid).Log().Error())
		}
		cn.WorkflowID = types.WorkflowID(wfid)
	}
	cn.Model.Slug = strings.ReplaceAll(cn.Name, " ", "-")
	cn.Enabled.Value = false
	if enabled := c.FormValue("comfynode_enabled"); enabled == "on" {
		cn.Enabled.Value = true
	}
	cn.Bypass.Value = false
	if bypass := c.FormValue("comfynode_bypass"); bypass == "on" {
		cn.Bypass.Value = true
	}
	GetLogger(3).Flogger("comfynode pre-save: %#v", cn)
	if err := cn.Set(c, update); err != nil {
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
		if err := wf.Set(c, true); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	return HandleWorkflowList(c)
}

func HandleComfyDelete(c echo.Context) error {
	GetLogger(4).Flogger("HandleComfyDelete called")
	if id := c.Param("id"); id != "" {
		cn := types.NewComfyNode(&id)
		if err := cn.Delete(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleComfyList(c)
	}
	return HandleWorkflowList(c)
}

func HandleComfyList(c echo.Context) error {
	GetLogger(4).Flogger("HandleComfyList called")
	dt := DisplayComfyNode{}
	if err := dt.Init(c, "list"); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	return c.Render(http.StatusOK, "node.comfy.tpl", dt)
}

func HandleOllamaNew(c echo.Context) error {
	GetLogger(4).Flogger("HandleOllamaNew called")
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
	GetLogger(4).Flogger("HandleOllamaEdit called")
	dt := DisplayOllamaNode{}
	dt.Init(c, "edit")
	return c.Render(http.StatusOK, "node.ollama.tpl", dt)
}

func HandleOllamaSave(c echo.Context) error {
	GetLogger(4).Flogger("HandleOllamaSave called")
	var update bool
	cn := types.NewOllamaNode(nil)
	if id := c.Param("id"); id != "" {
		update = true
		cn = types.NewOllamaNode(&id)
		if err := cn.Get(c); err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
			}
		}
	}
	if err := c.Bind(&cn); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleOllamaSave"}.Wrap(err))
	}
	cn.Model.Slug = strings.ReplaceAll(cn.Name, " ", "-")
	if err := cn.Set(c, update); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if !cn.WorkflowID.IsNil() {
		wfid := cn.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		if err := wf.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wf.OllamaNodeAppend(cn)
		if err := wf.Set(c, true); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	return HandleWorkflowList(c)
}

func HandleOllamaDelete(c echo.Context) error {
	GetLogger(4).Flogger("HandleOllamaDelete called")
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
	GetLogger(4).Flogger("HandleOllamaList called")
	dt := DisplayOllamaNode{}
	dt.Init(c, "list")
	return c.Render(http.StatusOK, "node.ollama.tpl", dt)
}

func HandleSSHNew(c echo.Context) error {
	GetLogger(4).Flogger("HandleSSHNew called")
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
	GetLogger(4).Flogger("HandleSSHEdit called")
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
	GetLogger(4).Flogger("HandleSSHSave called")
	dt := DisplaySSHNode{
		DisplayType: "edit",
	}
	dt.Menu.Href = "nodes/ssh"
	dt.Menu.Title = "SSH Nodes"
	var cn types.SSHNode
	var update bool
	if id := c.Param("id"); id != "" {
		update = true
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
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleComfySave"}.Wrap(err))
	}
	cn.Model.Slug = strings.ReplaceAll(cn.Name, " ", "-")
	if err := cn.Set(c, update); err != nil {
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
		if err := wf.Set(c, true); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	dt.SSHNode = cn
	return c.Render(http.StatusCreated, "node.ssh.tpl", dt)
}

func HandleSSHDelete(c echo.Context) error {
	GetLogger(4).Flogger("HandleSSHDelete called")
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
	GetLogger(4).Flogger("HandleSSHList called")
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
