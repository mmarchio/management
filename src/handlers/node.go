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
	g.GET("/comfy/new", HandleComfyNew)
	g.GET("/comfy/edit/:id", HandleComfyEdit)
	g.GET("/comfy/delete/:id", HandleComfyDelete)
	g.GET("/comfy/list", HandleComfyList)
	g.POST("/comfy/save/:id", HandleComfySave)

	g.GET("/ollama/new", HandleOllamaNew)
	g.GET("/ollama/edit/:id", HandleOllamaEdit)
	g.GET("/ollama/delete/:id", HandleOllamaDelete)
	g.GET("/ollama/list", HandleOllamaList)
	g.POST("/ollama/save/:id", HandleOllamaSave)
	
	g.GET("/ssh/new", HandleSSHNew)
	g.GET("/ssh/edit/:id", HandleSSHEdit)
	g.GET("/ssh/delete/:id", HandleSSHDelete)
	g.GET("/ssh/list", HandleSSHList)
	g.POST("/ssh/save/:id", HandleSSHSave)
}

func HandleComfyNew(c echo.Context) error {
	dt := DisplayComfyNode{
		ComfyNode: types.NewComfyNode(nil),
		DisplayType: "new",
	}
	dt.Menu.Href = "/nodes/comfy"
	dt.Menu.Title = "Comfy Nodes"
	return c.Render(http.StatusOK, "node.comfy.tpl", dt)
}

func HandleComfyEdit(c echo.Context) error {
	ctx := GetEchoCtx(c)
	dt := DisplayComfyNode{
		DisplayType: "edit",
	}
	dt.Menu.Href = "/nodes/comfy"
	dt.Menu.Title = "Comfy Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewComfyNode(&id)
		if err := cn.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt.ComfyNode = cn
	}
	return c.Render(http.StatusOK, "node.comfy.tpl", dt)
}

func HandleComfySave(c echo.Context) error {
	ctx := GetEchoCtx(c)
	dt := DisplayComfyNode{
		DisplayType: "edit",
	}
	dt.Menu.Href = "/nodes/comfy"
	dt.Menu.Title = "Comfy Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewComfyNode(&id)
		if err := cn.Get(ctx); err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
			}
		}
		if err := c.Bind(&cn); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleComfySave"}.Wrap(err))
		}
		if err := cn.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt.ComfyNode = cn
		return c.Render(http.StatusCreated, "node.comfy.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleComfyDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	dt := DisplayComfyNode{
		DisplayType: "list",
	}
	dt.Menu.Href = "/nodes/comfy"
	dt.Menu.Title = "Comfy Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewComfyNode(&id)
		if err := cn.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleComfyList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleComfyList(c echo.Context) error {
	ctx := GetEchoCtx(c)
	dt := DisplayComfyNode{
		DisplayType: "list",
	}
	dt.Menu.Href = "/nodes/comfy"
	dt.Menu.Title = "Comfy Nodes"
	cn := types.NewComfyNode(nil)
	list, err := cn.List(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt.List = list
	return c.Render(http.StatusOK, "node.comfy.tpl", dt)
}

func HandleOllamaNew(c echo.Context) error {
	dt := DisplayOllamaNode{
		OllamaNode: types.NewOllamaNode(nil),
		DisplayType: "new",
	}
	dt.Menu.Href = "/nodes/ollama"
	dt.Menu.Title = "Ollama Nodes"
	return c.Render(http.StatusOK, "node.ollama.tpl", dt)
}

func HandleOllamaEdit(c echo.Context) error {
	ctx := GetEchoCtx(c)
	dt := DisplayOllamaNode{
		DisplayType: "edit",
	}
	dt.Menu.Href = "/nodes/ollama"
	dt.Menu.Title = "Ollama Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewOllamaNode(&id)
		if err := cn.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt.OllamaNode = cn
	}
	return c.Render(http.StatusOK, "node.ollama.tpl", dt)
}

func HandleOllamaSave(c echo.Context) error {
	ctx := GetEchoCtx(c)
	dt := DisplayOllamaNode{
		DisplayType: "edit",
	}
	dt.Menu.Href = "/nodes/ollama"
	dt.Menu.Title = "Ollama Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewOllamaNode(&id)
		if err := cn.Get(ctx); err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
			}
		}
		if err := c.Bind(&cn); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleOllamaSave"}.Wrap(err))
		}
		if err := cn.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt.OllamaNode = cn
		return c.Render(http.StatusCreated, "node.ollama.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleOllamaDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	dt := DisplayOllamaNode{
		DisplayType: "list",
	}
	dt.Menu.Href = "/nodes/ollama"
	dt.Menu.Title = "Ollama Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewOllamaNode(&id)
		if err := cn.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleOllamaList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleOllamaList(c echo.Context) error {
	ctx := GetEchoCtx(c)
	dt := DisplayOllamaNode{
		DisplayType: "list",
	}
	dt.Menu.Href = "/nodes/ollama"
	dt.Menu.Title = "Ollama Nodes"
	cn := types.NewOllamaNode(nil)
	list, err := cn.List(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt.List = list
	return c.Render(http.StatusOK, "node.ollama.tpl", dt)
}

func HandleSSHNew(c echo.Context) error {
	dt := DisplaySSHNode{
		SSHNode: types.NewSSHNode(nil),
		DisplayType: "new",
	}
	dt.Menu.Href = "/nodes/ssh"
	dt.Menu.Title = "SSH Nodes"
	return c.Render(http.StatusOK, "node.ssh.tpl", dt)
}

func HandleSSHEdit(c echo.Context) error {
	ctx := GetEchoCtx(c)
	dt := DisplaySSHNode{
		DisplayType: "edit",
	}
	dt.Menu.Href = "/nodes/ssh"
	dt.Menu.Title = "SSH Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewSSHNode(&id)
		if err := cn.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt.SSHNode = cn
	}
	return c.Render(http.StatusOK, "node.ssh.tpl", dt)
}

func HandleSSHSave(c echo.Context) error {
	ctx := GetEchoCtx(c)
	dt := DisplaySSHNode{
		DisplayType: "edit",
	}
	dt.Menu.Href = "/nodes/ssh"
	dt.Menu.Title = "SSH Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewSSHNode(&id)
		if err := cn.Get(ctx); err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
			}
		}
		if err := c.Bind(&cn); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleComfySave"}.Wrap(err))
		}
		if err := cn.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt.SSHNode = cn
		return c.Render(http.StatusCreated, "node.ssh.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleSSHDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	dt := DisplaySSHNode{
		DisplayType: "list",
	}
	dt.Menu.Href = "/nodes/ssh"
	dt.Menu.Title = "SSH Nodes"
	if id := c.Param("id"); id != "" {
		cn := types.NewSSHNode(&id)
		if err := cn.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleSSHList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleSSHList(c echo.Context) error {
	ctx := GetEchoCtx(c)
	dt := DisplaySSHNode{
		DisplayType: "list",
	}
	dt.Menu.Href = "/nodes/ssh"
	dt.Menu.Title = "SSH Nodes"
	cn := types.NewSSHNode(nil)
	list, err := cn.List(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt.List = list
	return c.Render(http.StatusOK, "node.ssh.tpl", dt)
}
