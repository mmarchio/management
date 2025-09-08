package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

func RegisterWorkflowRoutes(e *echo.Echo) {
	g := e.Group("/workflow")
	g.GET("", HandleWorkflow)
	g.GET("/new", HandleWorkflowNew)
	g.POST("/save", HandleWorkflowSave)
	g.POST("/save/:id", HandleWorkflowSave)
	g.GET("/list", HandleWorkflowList)
	g.GET("/edit/:id", HandleWorkflowEdit)
	g.GET("/delete/:id", HandleWorkflowDelete)
}

func HandleAPIGetWorkflow(c echo.Context) error {
	GetLogger(4).Flogger("HandleAPIGetWorkflow called")
	if wfid := c.Param("id"); wfid != "" {
		entity := types.NewWorkflow(&wfid)
		if err := entity.Get(c); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, entity)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleAPIListWorkflow(c echo.Context) error {
	GetLogger(4).Flogger("HandleAPIListWorkflow called")
	entity := types.NewWorkflow(nil)
	entities, err := entity.List(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, entities)
}

func HandleAPISaveWorkflow(c echo.Context) error {
	GetLogger(4).Flogger("HandleAPISaveWorkflow called")
	var update bool
	if id := c.Param("id"); id != "" {
		update = true
	}
	entity := types.NewWorkflow(nil)
	if err := c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, merrors.EchoBindError{Package: "handlers", Function: "HandleAPISaveWorkflow"}.Wrap(err))
	}
	if entity.Name == "" && c.FormValue("name") != "" {
		entity.Name = c.FormValue("name")
	}
	if err := entity.Set(c, update); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, entity)
}

func HandleWorkflow(c echo.Context) error {
	GetLogger(4).Flogger("HandleWorkflow called")
	dt := DisplayWorkflow{}
	dt.Init(c, "none")
	return c.Render(http.StatusOK, "workflow.tpl", dt)
}

func HandleWorkflowList(c echo.Context) error {
	GetLogger(4).Flogger("HandleWorkflowList called")
	dt := DisplayWorkflow{}
	dt.Init(c, "list")
	return c.Render(http.StatusOK, "workflow.tpl", dt)
}

func HandleWorkflowDelete(c echo.Context) error {
	GetLogger(4).Flogger("HandleWorkflowDelete called")
	if wfid := c.Param("id"); wfid != "" {
		entity := types.NewWorkflow(&wfid)
		if err := entity.Delete(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		entity.NodeCleanup(c)
		return HandleWorkflowList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleWorkflowSave(c echo.Context) error {
	GetLogger(4).Flogger("HandleWorkflowSave called")
	var update bool
	id := "new"
	entity := types.NewWorkflow(&id)
	if wfid := c.Param("id"); wfid != "" {
		update = true
		entity := types.NewWorkflow(&wfid)
		if err := entity.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	if err := c.Bind(&entity); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if entity.Name == "" && c.FormValue("name") != "" {
		entity.Name = c.FormValue("name")
	}
	if entity.Name == "" {
		fmt.Printf(c.FormValue("name"))
	}
	if err := entity.Set(c, update); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt := DisplayWorkflow{
		Workflow: types.Workflow{},
		DisplayType: "new",
		Menu: Menu{
			Href: "workflow",
			Title: "Workflow",
		},
	}
	return c.Render(http.StatusOK, "workflow.tpl", dt)
}

func HandleWorkflowNew(c echo.Context) error {
	GetLogger(4).Flogger("HandleWorkflowNew called")
	dt := DisplayWorkflow{}
	dt.Init(c, "new")
	return c.Render(http.StatusOK, "workflow.tpl", dt)
}

func HandleWorkflowEdit(c echo.Context) error {
	GetLogger(4).Flogger("HandleWorkflowEdit called")
	if wfid := c.Param("id"); wfid != "" {
		dt := DisplayWorkflow{}
		dt.Init(c, "edit")
		return c.Render(http.StatusOK, "workflow.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}
