package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

func RegisterStepRoutes(e *echo.Echo) {
	g := e.Group("/step")
	g.GET("", HandleStep)
	g.GET("/:id", HandleStepGet)
	g.GET("/list", HandleStepList)
	g.GET("/new", HandleStepNew)
	g.GET("/edit/:id", HandleStepEdit)
	g.POST("/save", HandleStepSave)
	g.POST("/save/:id", HandleStepSave)
	g.GET("/delete/:id", HandleStepDelete)
}


func HandleAPIGetStep(c echo.Context) error {
	
	GetLogger(4).Flogger("HandleAPIGetStep called")
	if id := c.Param("id"); id != "" {
		entity := types.NewStep(&id)
		if err := entity.Get(c); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, entity)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPISetStep(c echo.Context) error {
	var err error
	var update bool
	GetLogger(4).Flogger("HandleAPISetStep called")
	if id := c.Param("id"); id != "" {
		update = true
	}
	entity := types.NewStep(nil)
	if err = c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err = entity.Set(c, update); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, entity)
}

func HandleAPIListStep(c echo.Context) error {
	GetLogger(4).Flogger("HandleAPIListStep called")
	prompt := types.NewStep(nil)
	prompts, err := prompt.List(c, "order desc")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, prompts)
}

func HandleStep(c echo.Context) error {
	GetLogger(4).Flogger("HandleSteps called")
	dt := DisplayStep{
		Step: types.Step{},
		DisplayType: "none",
		Menu: Menu{
			Href: "step",
			Title: "Step",
		},
	}
	return c.Render(http.StatusOK, "step.tpl", dt)
}

func HandleStepNew(c echo.Context) error {
	GetLogger(4).Flogger("HandleStepsNew called")
	dt := DisplayStep{}
	dt.Init(c, "new")
	return c.Render(http.StatusOK, "step.tpl", dt)
}

func HandleStepEdit(c echo.Context) error {
	if id := c.Param("id"); id != "" {
		entity := types.NewStep(&id)
		if err := entity.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt := DisplayStep{}
		dt.Init(c, "edit")
		return c.Render(http.StatusOK, "step.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleStepSave(c echo.Context) error {
	GetLogger(4).Flogger("HandleStepSave called")
	var err error
	var entity types.Step
	var update bool
	if id := c.Param("id"); id != "" {
		update = true
		entity = types.NewStep(&id)
		if err := entity.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		} 
	} else {
		entity = types.NewStep(nil)
	}
	if err = c.Bind(&entity); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleStepSave"}.Wrap(err).Log())
	}
	entity.Model.Slug = strings.ReplaceAll(entity.Name, " ", "-")
	entity.Enabled.Value = false
	if enabled := c.FormValue("step_enabled"); enabled == "on" {
		entity.Enabled.Value = true
	}
	entity.Bypass.Value = false
	if bypass := c.FormValue("step_bypass"); bypass == "on" {
		entity.Bypass.Value = true
	}
	if err = entity.Set(c, update); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}

	return HandleStepList(c)
}

func HandleStepList(c echo.Context) error {
	GetLogger(4).Flogger("HandleStepList called")
	dt := DisplayStep{}
	dt.Init(c, "list")
	return c.Render(http.StatusOK, "step.tpl", dt)
}

func HandleStepGet(c echo.Context) error {
	GetLogger(4).Flogger("HandleStepsGet called")
	if cutid := c.Param("id"); cutid != "" {
		entity := types.NewStep(&cutid)
		if err := entity.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt := DisplayStep{
			Step: types.Step{},
			DisplayType: "new",
			Menu: Menu{
				Href: "step",
				Title: "Step",
			},
			Enabled: types.Toggle{
				NamePrefix: "",
				IdPrefix: "",
				Suffix: "",
				Title: "",
			},
			Bypass: types.Toggle{
				NamePrefix: "",
				IdPrefix: "",
				Suffix: "",
				Title: "",
			},
		}
		return c.Render(http.StatusOK, "step.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleStepDelete(c echo.Context) error {
	GetLogger(4).Flogger("HandleStepsDelete called")
	if id := c.Param("id"); id != "" {
		entity := types.NewStep(&id)
		if err := entity.Delete(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleStepList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

