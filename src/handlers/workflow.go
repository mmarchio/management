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
	ctx := GetEchoCtx(c)
	if wfid := c.Param("id"); wfid != "" {
		entity := types.NewWorkflow(&wfid)
		if err := entity.Get(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, entity)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleAPIListWorkflow(c echo.Context) error {
	ctx := GetEchoCtx(c)
	entity := types.NewWorkflow(nil)
	entities, err := entity.List(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, entities)
}

func HandleAPISaveWorkflow(c echo.Context) error {
	ctx := GetEchoCtx(c)
	entity := types.NewWorkflow(nil)
	if err := c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, merrors.EchoBindError{Package: "handlers", Function: "HandleAPISaveWorkflow"}.Wrap(err))
	}
	if entity.Name == "" && c.FormValue("name") != "" {
		entity.Name = c.FormValue("name")
	}
	if err := entity.Set(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, entity)
}

func HandleWorkflow(c echo.Context) error {
	dt := DisplayWorkflow{
		Workflow: types.Workflow{},
		DisplayType: "none",
		Menu: Menu{
			Href: "workflow",
			Title: "Workflow",
		},
	}
	return c.Render(http.StatusOK, "workflow.tpl", dt)
}

func HandleWorkflowList(c echo.Context) error {
	ctx := GetEchoCtx(c)
	entity := types.NewWorkflow(nil)
	list, err := entity.List(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt := DisplayWorkflow{
		Workflow: types.Workflow{},
		List: list,
		DisplayType: "list",
		Menu: Menu{
			Href: "workflow",
			Title: "Workflow",
		},
	}
	return c.Render(http.StatusOK, "workflow.tpl", dt)
}

func HandleWorkflowDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if wfid := c.Param("id"); wfid != "" {
		entity := types.NewWorkflow(&wfid)
		if err := entity.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleWorkflowList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleWorkflowSave(c echo.Context) error {
	ctx := GetEchoCtx(c)
	entity := types.NewWorkflow(nil)
	if wfid := c.Param("id"); wfid != "" {
		entity := types.NewWorkflow(&wfid)
		if err := entity.Get(ctx); err != nil {
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
	if err := entity.Set(ctx); err != nil {
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
	entity := types.NewWorkflow(nil)
	entity.Model.ID = ""
	entity.ID = types.WorkflowID("")
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

func HandleWorkflowEdit(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if wfid := c.Param("id"); wfid != "" {
		entity := types.NewWorkflow(&wfid)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt := DisplayWorkflow{
			Workflow: entity,
			DisplayType: "new",
			Menu: Menu{
				Href: "workflow",
				Title: "Workflow",
			},
		}
		return c.Render(http.StatusOK, "workflow.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

// func get(id string) (*types.Workflow, error) {
// 	ctx := GetEchoCtx(c)

// 	entity := types.NewWorkflow()
// 	entity.Model.ID = id
// 	entity.ID = types.WorkflowID(entity.Model.ID)
// 	if err := entity.Get(ctx); err != nil {
// 		return nil, err
// 	}
// 	return &entity, nil
// }