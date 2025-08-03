package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mmarchio/management/database"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

func RegisterDispositionRoutes(e *echo.Echo) {
	g := e.Group("/dispositions")
	g.GET("", HandleDispositions)
	g.GET("/:id", HandleDispositionGet)
	g.GET("/list", HandleDispositionsList)
	g.GET("/new", HandleDispositionsNew)
	g.POST("/save", HandleDispositionsSave)
	g.GET("/delete/:id", HandleDispositionsDelete)
}

func HandleAPIGetDisposition(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewDisposition(&id)
		if err := entity.Get(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, entity)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPISetDisposition(c echo.Context) error {
	var err error
	ctx := GetEchoCtx(c)
	entity := types.NewDisposition(nil)
	if err := c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, merrors.EchoBindError{Package: "handlers", Function: "HandleAPISetDisposition"}.Wrap(err))
	}
	entity.Model.ID = c.FormValue("id")
	entity, err = entity.SetID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = entity.Set(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, entity)
}

func HandleAPIListDisposition(c echo.Context) error {
	ctx := GetEchoCtx(c)
	disposition := types.NewDisposition(nil)
	dispositions, err := disposition.List(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dispositions)
}

func HandleDispositions(c echo.Context) error {
	dt := DisplayDisposition{
		DisplayType: "none",
		Menu: Menu{
			Href: "dispositions",
			Title: "Disposition",
		},
	}
	return c.Render(http.StatusOK, "dispositions.tpl", dt)
}

func HandleDispositionsNew(c echo.Context) error {
	dt := DisplayDisposition{
		Disposition: types.Disposition{},
		DisplayType: "new",
		Menu: Menu{
			Href: "dispositions",
			Title: "Disposition",
		},
	}
	return c.Render(http.StatusOK, "dispositions.tpl", dt)
}

func HandleDispositionGet(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		disposition := types.NewDisposition(&id)
		if err := disposition.Get(ctx); err != nil {
			return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
		}
		dt := DisplayDisposition{
			Disposition: disposition,
			DisplayType: "new",
			Menu: Menu{
				Href: "dispositions",
				Title: "Disposition",
			},
		}
		return c.Render(http.StatusOK, "disposition.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "missing id")
}

func HandleDispositionsList(c echo.Context) error {
	ctx := GetEchoCtx(c)
	disposition := types.NewDisposition(nil)
	dispositions, err := disposition.List(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt := DisplayDisposition{
		Disposition: types.Disposition{},
		List: dispositions,
		DisplayType: "list",
		Menu: Menu{
			Href: "dispositions",
			Title: "Disposition",
		},
	}
	return c.Render(http.StatusOK, "dispositions.tpl", dt)

}

func HandleDispositionsSave(c echo.Context) error {
	var err error
	ctx := database.GetDatabaseCtx()
	entity := types.NewDisposition(nil)
	if err := c.Bind(&entity); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleDispositionsSave"}.Wrap(err))
	}
	entity, err = entity.Bind(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error)
	}
	if err = entity.Set(ctx); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err)
	}
	
	dt := DisplayDisposition{
		Disposition: entity,
		DisplayType: "new",
		Menu: Menu{
			Href: "dispositions",
			Title: "Disposition",
		},
	}
	return c.Render(http.StatusCreated, "dispositions.tpl", dt)
}

func HandleDispositionsDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewDisposition(&id)
		entity.Model.ID = c.Param("id")
		if err := entity.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleDispositionsList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}
