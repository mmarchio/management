package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mmarchio/management/database"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

func RegisterPromptTemplateRoutes(e *echo.Echo) {
	g := e.Group("/prompttemplates")
	g.GET("", HandlePromptTemplates)
	g.GET("/:id", HandlePromptTemplatesGet)
	g.GET("/list", HandlePromptTemplateList)
	g.GET("/new", HandlePromptTemplatesNew)
	g.POST("/save", HandlePromptTemplateSave)
	g.GET("/delete/:id", HandlePromptTemplatesDelete)
}

func HandleAPIGetPromptTemplate(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewPromptTemplate(&id)
		if err := entity.Get(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, entity)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPISetPromptTemplate(c echo.Context) error {
	var err error
	ctx := GetEchoCtx(c)
	entity := types.NewPromptTemplate(nil)
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

func HandleAPIListPromptTemplate(c echo.Context) error {
	ctx := GetEchoCtx(c)
	prompt := types.NewPromptTemplate(nil)
	prompts, err := prompt.List(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, prompts)
}

func HandlePromptTemplates(c echo.Context) error {
	dt := DisplayPromptTemplate{
		PromptTemplate: types.PromptTemplate{},
		DisplayType: "new",
		Menu: Menu{
			Href: "prompttemplates",
			Title: "Prompt Template",
		},
	}
	return c.Render(http.StatusOK, "prompttemplates.tpl", dt)
}

func HandlePromptTemplatesNew(c echo.Context) error {
	ctx := types.Context{}
	b, err := json.MarshalIndent(ctx, "", "  ")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt := DisplayPromptTemplate{
		PromptTemplate: types.PromptTemplate{},
		DisplayType: "new",
		Menu: Menu{
			Href: "prompttemplates",
			Title: "Prompt Template",
		},
		Context: string(b),
	}
	return c.Render(http.StatusOK, "prompttemplates.tpl", dt)
}

func HandlePromptTemplateSave(c echo.Context) error {
	var err error
	ctx := database.GetDatabaseCtx()
	entity := types.NewPromptTemplate(nil)
	if err = c.Bind(&entity); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandlePromptTemplateSave"}.Wrap(err))
	}
	entity, err = entity.SetID()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if c.FormValue("vars") != "" {
		entity.Vars = c.FormValue("vars")
	}
	if err = entity.Set(ctx); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}

	dt := DisplayPromptTemplate{
		PromptTemplate: entity,
		DisplayType: "new",
		Menu: Menu{
			Href: "prompttemplates",
			Title: "Prompt Template",
		},
	}
	return c.Render(http.StatusCreated, "prompttemplates.tpl", dt)
}

func HandlePromptTemplateList(c echo.Context) error {
	var err error
	ctx := GetEchoCtx(c)
	entity := types.NewPromptTemplate(nil)
	entities, err := entity.List(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}

	dt := DisplayPromptTemplate{
		PromptTemplate: types.PromptTemplate{},
		List: entities,
		DisplayType: "list",
		Menu: Menu{
			Href: "prompttemplates",
			Title: "Prompt Template",
		},
	}
	return c.Render(http.StatusOK, "prompttemplates.tpl", dt)
}

func HandlePromptTemplatesGet(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if cutid := c.Param("id"); cutid != "" {
		entity := types.NewPromptTemplate(&cutid)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt := DisplayPromptTemplate{
			PromptTemplate: types.PromptTemplate{},
			DisplayType: "new",
			Menu: Menu{
				Href: "prompttemplates",
				Title: "Prompt Template",
			},
		}
		return c.Render(http.StatusOK, "prompttemplates.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandlePromptTemplatesDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewPromptTemplate(&id)
		if err := entity.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandlePromptTemplateList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}
