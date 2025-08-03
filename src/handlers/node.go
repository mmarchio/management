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
	p := e.Group("/params")
	g.GET("", HandleNode)
	g.GET("/new", HandleNodeNew)
	g.GET("/new/:id", HandleNodeNew)
	g.POST("/save", HandleNodeSave)
	g.POST("/save/:id", HandleNodeSave)
	g.GET("/list", HandleNodeList)
	g.GET("/edit/:id", HandleNodeEdit)
	g.GET("/delete/:id", HandleNodeDelete)
	p.GET("/edit/:id", HandleParamsEdit)
	p.POST("/save/:id", HandleParamsSave)
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
		nodes := make([]types.Node, len(wf.Nodes)-1)
		for _, n := range wf.Nodes {
			if n.ID == entity.ID {
				continue
			}
			nodes = append(nodes, n)
		}
		wf.Nodes = nodes
		if err := wf.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleWorkflowList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleNodeSave(c echo.Context) error {
	ctx := GetEchoCtx(c)
	entity := types.NewNode(nil)
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
	if err := entity.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if entity.WorkflowID != "" {
		wfid := entity.WorkflowID.String()
		workflow := types.NewWorkflow(&wfid)
		if err := workflow.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if workflow.Nodes == nil {
			nodes := make([]types.Node, 0)
			nodes = append(nodes, entity)
			workflow.Nodes = nodes
		} else {
			workflow.Nodes = append(workflow.Nodes, entity)
		}
		if err := workflow.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
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

func HandleParamsEdit(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewNode(&id)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		jr := types.JobRun{}
		jr.ContentType = "jobrun"
		jr.WorkflowID = entity.WorkflowID
		if err := jr.FindBy(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		node := DisplayNode{}
		node.New(entity)
		node.Prompt = jr.Context.Prompt
		node.Disposition = jr.Context.Disposition
		if err := node.GetSystemPrompts(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if err := node.GetPromptTemplates(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if entity.Type == "ollama_node" {
			onode := types.OllamaNode{
				Name: entity.Params.GetName(),
				OllamaModel: entity.Params.GetModel(),
				SystemPrompt: entity.Params.GetSystemPrompt(),
				Prompt: entity.Params.GetPrompt(),
				PromptTemplate: entity.Params.GetPromptTemplate(),
			}
			if err := onode.ParsePromptTemplate(ctx); err != nil {
				return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
			}
			node.Params = onode
		}
		if entity.Type == "comfy_node" {
			node.Params = types.ComfyNode{}
		}
		if entity.Type == "ssh_node" {
			node.Params = types.SSHNode{}
		}
		return c.Render(http.StatusOK, "params.edit.tpl", node)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleParamsSave(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.FormValue("node_id"); id != "" {
		entity := types.NewNode(&id)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if entity.Type == "ollama_node" {
			param := types.OllamaNode{}
			if err := c.Bind(&param); err != nil {
				return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
			}
			entity.Params = param
		}
		if entity.Type == "comfy_node" {
			param := types.ComfyNode{}
			if err := c.Bind(&param); err != nil {
				return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
			}
			entity.Params = param
		}
		if entity.Type == "ssh_node" {
			param := types.SSHNode{}
			if err := c.Bind(&param); err != nil {
				return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
			}
			entity.Params = param
		}
		if err := entity.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wfid := entity.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		if err := wf.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if wf.Nodes == nil {
			nodes := make([]types.Node, 0)
			nodes = append(nodes, entity)
			wf.Nodes = nodes
		} else {
			exists := false
			var index int
			for i, wn := range wf.Nodes {
				if wn.ID == entity.ID {
					exists = true
					index = i
					break
				}
			}
			if !exists {
				wf.Nodes = append(wf.Nodes, entity)
			} else {
				wf.Nodes[index] = entity
			}
		}
		if err := wf.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleWorkflowList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}