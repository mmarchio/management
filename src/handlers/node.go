package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

func RegisterNodesRoutes(e *echo.Echo) {
	g := e.Group("/node")
	g.GET("", HandleNode)
	g.GET("/new", HandleNodeNew)
	g.GET("/new/:id", HandleNodeNew)
	// g.POST("/save", HandleNodeSave)
	// g.POST("/save/:id", HandleNodeSave)
	g.GET("/list", HandleNodeList)
	g.GET("/edit/:id", HandleNodeEdit)
	g.GET("/delete/:id", HandleNodeDelete)
	g.GET("/comfynode/delete/:id", HandleComfyNodeDelete)
	g.GET("/ollamanode/delete/:id", HandleOllamaNodeDelete)
	g.GET("/sshnode/delete/:id", HandleSSHNodeDelete)
	g.GET("/comfynode/save", HandleComfyNodeSave)
	g.GET("/ollamanode/save", HandleOllamaNodeSave)
	g.GET("/sshnode/save", HandleSSHNodeSave)
	g.GET("/comfynode/save/:id", HandleComfyNodeSave)
	g.GET("/ollamanode/save/:id", HandleOllamaNodeSave)
	g.GET("/sshnode/save/:id", HandleSSHNodeSave)
}

func HandleAPIGetNode(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewNode(&id)
		if err := entity.Get(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal server error: %w", err))
		}
		return c.JSON(http.StatusOK, entity)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPIListNode(c echo.Context) error {
	ctx := GetEchoCtx(c)
	entity := types.NewNode(nil)
	entities, err := entity.List(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, entities)
}

func HandleAPISaveNode(c echo.Context) error {
	ctx := GetEchoCtx(c)
	entity := types.NewNode(nil)
	if err := c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, merrors.EchoBindError{Package: "handlers", Function: "HandleAPISaveNode"}.Wrap(err))
	}
	if err := entity.Set(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, entity)
}

func HandleNode(c echo.Context) error {
	dt := DisplayNode{
		DisplayType: "none",
		Menu: Menu{
			Href: "node",
			Title: "Node",
		},
	}
	return c.Render(http.StatusOK, "node.tpl", dt)
}

func HandleNodeList(c echo.Context) error {
	ctx := GetEchoCtx(c)
	entity := types.NewNode(nil)
	list, err := entity.List(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt := DisplayNode{
		List: list,
		DisplayType: "list",
		Menu: Menu{
			Href: "node",
			Title: "Node",
		},
	}
	return c.Render(http.StatusOK, "node.tpl", dt)
}

func HandleNodeDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewNode(&id)
		if err := entity.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wfid := entity.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		fmt.Printf("handlers:node:HandleNodeDelete node: %#v\n", entity)
		fmt.Printf("handlers:node:HandleNodeDelete wf: %#v\n", wf)
		if err := wf.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wf.CutNode(entity.Model.ID)
		wf.CutNodeOrder(fmt.Sprintf("%s:%s", entity.Model.ID, entity.Model.ContentType))
		if err := wf.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleWorkflowList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleComfyNodeDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewComfyNode(&id)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wfid := entity.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		wf.CutNode(entity.Model.ID)
		wf.CutNodeOrder(fmt.Sprintf("%s:%s", entity.Model.ID, entity.Model.ContentType))
		if err := wf.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if err := entity.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleWorkflowList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleOllamaNodeDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewOllamaNode(&id)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wfid := entity.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		wf.CutNode(entity.Model.ID)
		wf.CutNodeOrder(fmt.Sprintf("%s:%s", entity.Model.ID, entity.Model.ContentType))
		if err := wf.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if err := entity.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleWorkflowList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleSSHNodeDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewSSHNode(&id)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wfid := entity.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		wf.CutNode(entity.Model.ID)
		wf.CutNodeOrder(fmt.Sprintf("%s:%s", entity.Model.ID, entity.Model.ContentType))
		if err := wf.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if err := entity.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleWorkflowList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

// func HandleNodeSave(c echo.Context) error {
// 	ctx := GetEchoCtx(c)
// 	entity := types.NewNode(nil)
// 	if err := c.Bind(&entity); err != nil {
// 			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
// 	}
// 	fmt.Printf("handlers:node:HandleNodeSave node: %#v\n", entity)
// 	if c.FormValue("enabled") == "on" {
// 		entity.Enabled = true
// 	}
// 	if c.FormValue("bypass") == "on" {
// 		entity.Bypass = true
// 	}
// 	if c.FormValue("workflow_id") != "" {
// 		entity.WorkflowID = types.WorkflowID(c.FormValue("workflow_id"))
// 	}
// 	if err := entity.Set(ctx); err != nil {
// 			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
// 	}
// 	if entity.WorkflowID != "" {
// 		wfid := entity.WorkflowID.String()
// 		workflow := types.NewWorkflow(&wfid)
// 		if err := workflow.Get(ctx); err != nil {
// 			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
// 		}
// 		if workflow.Nodes == nil {
// 			nodes := make([]types.Node, 0)
// 			nodes = append(nodes, entity)
// 			workflow.Nodes = nodes
// 		} else {
// 			workflow.Nodes = append(workflow.Nodes, entity)
// 		}
// 		if err := workflow.Set(ctx); err != nil {
// 			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
// 		}
// 	}
// 	return HandleNodeList(c)
// }

func HandleComfyNodeSave(c echo.Context) error {
	ctx := GetEchoCtx(c)
	entity := types.NewComfyNode(nil)
	if err := c.Bind(&entity); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if c.FormValue("enabled") == "on" {
		entity.Enabled = true
	}
	if c.FormValue("bypass") == "on" {
		entity.Bypass = true
	}
	if c.FormValue("workflow_id") != "" {
		entity.WorkflowID = types.WorkflowID(c.FormValue("workflow_id"))
	}
	if entity.WorkflowID != "" {
		wfid := entity.WorkflowID.String()
		workflow := types.NewWorkflow(&wfid)
		if err := workflow.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if workflow.ComfyNodes == nil {
			nodes := make([]types.ComfyNode, 0)
			nodes = append(nodes, entity)
			workflow.ComfyNodes = nodes
		} else {
			workflow.ComfyNodes = append(workflow.ComfyNodes, entity)
		}
		workflow.NodeOrder[fmt.Sprintf("%s:%s", entity.Model.ID, entity.Model.ContentType)] = len(workflow.NodeOrder)
		if err := workflow.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if err := entity.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	return HandleWorkflowList(c)
}

func HandleOllamaNodeSave(c echo.Context) error {
	ctx := GetEchoCtx(c)
	entity := types.NewOllamaNode(nil)
	if err := c.Bind(&entity); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if c.FormValue("enabled") == "on" {
		entity.Enabled = true
	}
	if c.FormValue("bypass") == "on" {
		entity.Bypass = true
	}
	if c.FormValue("workflow_id") != "" {
		entity.WorkflowID = types.WorkflowID(c.FormValue("workflow_id"))
	}
	if entity.WorkflowID != "" {
		wfid := entity.WorkflowID.String()
		workflow := types.NewWorkflow(&wfid)
		if err := workflow.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if workflow.OllamaNodes == nil {
			nodes := make([]types.OllamaNode, 0)
			nodes = append(nodes, entity)
			workflow.OllamaNodes = nodes
		} else {
			workflow.OllamaNodes = append(workflow.OllamaNodes, entity)
		}
		workflow.NodeOrder[entity.Model.ID] = len(workflow.NodeOrder)
		if err := workflow.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if err := entity.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	return HandleWorkflowList(c)
}

func HandleSSHNodeSave(c echo.Context) error {
	ctx := GetEchoCtx(c)
	entity := types.NewSSHNode(nil)
	if err := c.Bind(&entity); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if c.FormValue("enabled") == "on" {
		entity.Enabled = true
	}
	if c.FormValue("bypass") == "on" {
		entity.Bypass = true
	}
	if c.FormValue("workflow_id") != "" {
		entity.WorkflowID = types.WorkflowID(c.FormValue("workflow_id"))
	}
	if entity.WorkflowID != "" {
		wfid := entity.WorkflowID.String()
		workflow := types.NewWorkflow(&wfid)
		if err := workflow.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if workflow.SSHNodes == nil {
			nodes := make([]types.SSHNode, 0)
			nodes = append(nodes, entity)
			workflow.SSHNodes = nodes
		} else {
			workflow.SSHNodes = append(workflow.SSHNodes, entity)
		}
		workflow.NodeOrder[entity.Model.ID] = len(workflow.NodeOrder)
		if err := workflow.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if err := entity.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	return HandleWorkflowList(c)
}

func HandleNodeNew(c echo.Context) error {
	entity := types.NewNode(nil)
	if id := c.Param("id"); id != "" {
		entity.WorkflowID = types.WorkflowID(id)
	}
	dt := DisplayNode{
		Node: entity,
		DisplayType: "new",
		Menu: Menu{
			Href: "node",
			Title: "Node",
		},
	}
	return c.Render(http.StatusOK, "node.tpl", dt)
}

func HandleNodeEdit(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewNode(&id)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt := DisplayNode{
			Node: entity,
			DisplayType: "new",
			Menu: Menu{
				Href: "node",
				Title: "Node",
			},
		}
		return c.Render(http.StatusOK, "node.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

