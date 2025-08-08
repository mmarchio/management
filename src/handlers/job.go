package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

func RegisterJobRoutes(e *echo.Echo) {
	g := e.Group("/jobs")
	g.GET("", HandleJobs)
	g.GET("/new", HandleJobs)
	g.GET("/list", HandleJobsList)
	g.GET("/workflow/add/:id", HandlerJobWorkflowAdd)
	g.POST("/save/:id", HandleJobSave)
}

func HandleAPIGetJob(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		job := types.NewJob(&id)
		if err := job.Get(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, job)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPIListJob(c echo.Context) error {
	ctx := GetEchoCtx(c)
	job := types.NewJob(nil)
	jobs, err := job.List(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, jobs)
}

func HandleAPISaveJob(c echo.Context) error {
	ctx := GetEchoCtx(c)
	job := types.NewJob(nil)
	if err := c.Bind(&job); err != nil {
		return c.JSON(http.StatusInternalServerError, merrors.EchoBindError{Package: "handlers", Function: "HandleAPISaveJob"}.Wrap(err))
	}
	if err := job.Set(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, job)
}

func HandleJobs(c echo.Context) error {
	dt := DisplayJob{
		Job: types.Job{},
		DisplayType: "none",
		Menu: Menu{
			Href: "jobs",
			Title: "Job",
		},
	}
	return c.Render(http.StatusOK, "jobs.tpl", dt)
}

func HandleJobsList(c echo.Context) error {
	ctx := GetEchoCtx(c)
	job := types.NewJob(nil)
	list, err := job.List(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt := DisplayJob{
		Job: types.Job{},
		DisplayType: "list",
		List: list,
		Menu: Menu{
			Href: "jobs",
			Title: "Job",
		},
	}
	return c.Render(http.StatusOK, "jobs.tpl", dt)
}

func HandleJobsDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewJob(&id)
		if err := entity.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleJobsList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandlerJobWorkflowAdd(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewJob(&id)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wf := types.NewWorkflow(nil)
		wfs, err := wf.List(ctx)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt := DisplayJob{
			Job: entity,
			Workflows: wfs,
			Menu: Menu{
				Href: "jobs",
				Title: "Job",
			},
		}
		return c.Render(http.StatusOK, "job.workflow.add.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleJobSave(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewJob(&id)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if err := c.Bind(&entity); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		if entity.WorkflowID.IsNil() && c.FormValue("workflow_id") != "" {
			entity.WorkflowID = types.WorkflowID(c.FormValue("workflow_id"))
		}
		if err := entity.Set(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleJobsList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}