package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

func HandleAPIGetComfyUITemplate(e echo.Context) error {
	GetLogger(4).Flogger("HandleAPIGetComfyUITemplate called")
	if id := e.Param("id"); id != "" {
		entity := types.NewComfyUITemplate(&id)
		if err := entity.Get(e); err != nil {
			return e.JSON(http.StatusInternalServerError, err.Error())
		}
		return e.JSON(http.StatusOK, entity)
	}
	return e.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPISetComfyUITemplate(c echo.Context) error {
	GetLogger(4).Flogger("HandleAPISetComfyUITemplate called")
	var err error
	var update bool
	if id := c.Param("id"); id != "" {
		update = true
	}
	entity := types.NewComfyUITemplate(nil)
	if err = c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err = entity.Set(c, update); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, entity)
}

func HandleAPIListComfyUITemplate(c echo.Context) error {
	GetLogger(4).Flogger("HandleAPIListComfyUITemplate called")
	prompt := types.NewComfyUITemplate(nil)
	prompts, err := prompt.List(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, prompts)
}

func HandleComfyUITemplates(c echo.Context) error {
	GetLogger(4).Flogger("HandleComfyUITemplates called")
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
	GetLogger(4).Flogger("HandleComfyUITemplatesNew called")
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
	var update bool
	GetLogger(4).Flogger("HandleComfyUITemplateSave called")
	if id := c.Param("id"); id != "" {
		update = true
	}
	entity := types.NewComfyUITemplate(nil)
	if err = c.Bind(&entity); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleComfyUITemplateSave"}.Wrap(err))
	}
	if err = entity.Set(c, update); err != nil {
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
	GetLogger(4).Flogger("HandleComfyUITemplateList called")
	entity := types.NewComfyUITemplate(nil)
	entities, err := entity.List(c)
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
	GetLogger(4).Flogger("HandleComfyUITemplatesGet called")
	if cutid := c.Param("id"); cutid != "" {
		entity := types.NewComfyUITemplate(&cutid)
		if err := entity.Get(c); err != nil {
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
	GetLogger(4).Flogger("HandleComfyUITemplatesDelete called")
	if id := c.Param("id"); id != "" {
		entity := types.NewComfyUITemplate(&id)
		if err := entity.Delete(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleComfyUITemplateList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}
