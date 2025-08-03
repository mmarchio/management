package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mmarchio/management/database"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

func RegisterComfyUITemplatesRoutes(e *echo.Echo) {
	g := e.Group("/comfy")
	g.GET("", HandleComfyUITemplates)
	g.GET("/new", HandleComfyUITemplatesNew)
	g.POST("/save", HandleComfyUITemplateSave)
	g.GET("/list", HandleComfyUITemplateList)
	g.GET("/:id", HandleComfyUITemplatesGet)
	g.GET("/delete/:id", HandleComfyUITemplatesDelete)
}

func HandleAPIGetComfyUITemplate(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewComfyUITemplate(&id)
		if err := entity.Get(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, entity)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPISetComfyUITemplate(c echo.Context) error {
	var err error
	ctx := GetEchoCtx(c)
	entity := types.NewComfyUITemplate(nil)
	if err = c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	entity, err = entity.SetID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err = entity.Set(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, entity)
}

func HandleAPIListComfyUITemplate(c echo.Context) error {
	ctx := GetEchoCtx(c)
	prompt := types.NewComfyUITemplate(nil)
	prompts, err := prompt.List(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, prompts)
}

func HandleComfyUITemplates(c echo.Context) error {
	dt := DisplayComfyUITemplate{
		ComfyUITemplate: types.ComfyUITemplate{},
		DisplayType: "new",
		Menu: Menu{
			Href: "comfy",
			Title: "ComfyUI Template",
		},
	}
	return c.Render(http.StatusOK, "comfy.tpl", dt)
}

func HandleComfyUITemplatesNew(c echo.Context) error {
	dt := DisplayComfyUITemplate{
		ComfyUITemplate: types.ComfyUITemplate{},
		DisplayType: "new",
		Menu: Menu{
			Href: "comfy",
			Title: "ComfyUI Template",
		},
	}
	return c.Render(http.StatusOK, "comfy.tpl", dt)
}

func HandleComfyUITemplateSave(c echo.Context) error {
	var err error
	ctx := database.GetDatabaseCtx()
	entity := types.NewComfyUITemplate(nil)
	if err = c.Bind(&entity); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleComfyUITemplateSave"}.Wrap(err))
	}
	entity, err = entity.SetID()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if err = entity.Set(ctx); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}

	dt := DisplayComfyUITemplate{
		ComfyUITemplate: entity,
		DisplayType: "new",
		Menu: Menu{
			Href: "comfy",
			Title: "ComfyUI Template",
		},
	}
	return c.Render(http.StatusCreated, "comfy.tpl", dt)
}

func HandleComfyUITemplateList(c echo.Context) error {
	var err error
	ctx := GetEchoCtx(c)
	entity := types.NewComfyUITemplate(nil)
	entities, err := entity.List(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}

	dt := DisplayComfyUITemplate{
		ComfyUITemplate: types.ComfyUITemplate{},
		List: entities,
		DisplayType: "list",
		Menu: Menu{
			Href: "comfy",
			Title: "ComfyUI Template",
		},
	}
	return c.Render(http.StatusOK, "comfy.tpl", dt)
}

func HandleComfyUITemplatesGet(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if cutid := c.Param("id"); cutid != "" {
		entity := types.NewComfyUITemplate(&cutid)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt := DisplayComfyUITemplate{
			ComfyUITemplate: types.ComfyUITemplate{},
			DisplayType: "new",
			Menu: Menu{
				Href: "comfy",
				Title: "ComfyUI Template",
			},
		}
		return c.Render(http.StatusOK, "comfy.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleComfyUITemplatesDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewComfyUITemplate(&id)
		if err := entity.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleComfyUITemplateList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}
