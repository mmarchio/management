package handlers

import (
	"net/http"

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
	
	GetLogger().Flogger("HandleAPIGetStep called")
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
	GetLogger().Flogger("HandleAPISetStep called")
	entity := types.NewStep(nil)
	if err = c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	entity, err = entity.SetID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err = entity.Set(c); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, entity)
}

func HandleAPIListStep(c echo.Context) error {
	GetLogger().Flogger("HandleAPIListStep called")
	prompt := types.NewStep(nil)
	prompts, err := prompt.List(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, prompts)
}

func HandleStep(c echo.Context) error {
	GetLogger().Flogger("HandleSteps called")
	dt := DisplayStep{
		Step: types.Step{},
		DisplayType: "new",
		Menu: Menu{
			Href: "step",
			Title: "Step",
		},
	}
	return c.Render(http.StatusOK, "step.tpl", dt)
}

func HandleStepNew(c echo.Context) error {
	GetLogger().Flogger("HandleStepsNew called")
	dt := DisplayStep{
		Step: types.Step{},
		DisplayType: "new",
		Menu: Menu{
			Href: "step",
			Title: "step",
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

func HandleStepEdit(c echo.Context) error {
	if id := c.Param("id"); id != "" {
		entity := types.NewStep(&id)
		if err := entity.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt := DisplayStep{
			Step: entity,
			DisplayType: "edit",
			Menu: Menu{
				Href: "step",
				Title: "Step",
			},
			Enabled: types.Toggle{
				ID: entity.Enabled.ID,
				NamePrefix: entity.Enabled.NamePrefix,
				IdPrefix: entity.Enabled.IdPrefix,
				Suffix: entity.Enabled.Suffix,
				Title: entity.Enabled.Title,
			},
			Bypass: types.Toggle{
				ID: entity.Bypass.ID,
				NamePrefix: entity.Bypass.NamePrefix,
				IdPrefix: entity.Bypass.IdPrefix,
				Suffix: entity.Bypass.Suffix,
				Title: entity.Bypass.Title,
			},
		}
		return c.Render(http.StatusOK, "step.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleStepSave(c echo.Context) error {
	GetLogger().Flogger("HandleStepSave called")
	var err error
	var entity types.Step
	if id := c.Param("id"); id != "" {
		entity = types.NewStep(&id)
	} else {
		entity = types.NewStep(nil)
	}
	if err = c.Bind(&entity); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleStepSave"}.Wrap(err))
	}
	entity, err = entity.SetID()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if err = entity.Set(c); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}

	return HandleStepList(c)
}

func HandleStepList(c echo.Context) error {
	var err error
	GetLogger().Flogger("HandleStepList called")
	entity := types.NewStep(nil)
	entities, err := entity.List(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}

	dt := DisplayStep{
		Step: types.Step{},
		List: entities,
		DisplayType: "list",
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
	return c.Render(http.StatusOK, "prompttemplates.tpl", dt)
}

func HandleStepGet(c echo.Context) error {
	GetLogger().Flogger("HandleStepsGet called")
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
	GetLogger().Flogger("HandleStepsDelete called")
	if id := c.Param("id"); id != "" {
		entity := types.NewStep(&id)
		if err := entity.Delete(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleStepList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

