package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mmarchio/management/database"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

func RegisterSystemPromptsRoutes(e *echo.Echo) {
	g := e.Group("/systemprompts")
	g.GET("", HandleSystemPrompts)
	g.GET("/new", HandleSystemPromptsNew)
	g.POST("/save", HandleSystemPromptSave)
	g.POST("/save/:id", HandleSystemPromptSave)
	g.GET("/list", HandleSystemPromptsList)
	g.GET("/:id", HandleSystemPromptsGet)
	g.GET("/delete/:id", HandleSystemPromptDelete)
}

func HandleAPIGetSystemPrompt(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewSystemPrompt(&id)
		if err := entity.Get(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal server error: %w", err))
		}
		return c.JSON(http.StatusOK, entity)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPISetSystemPrompt(c echo.Context) error {
	var err error
	ctx := GetEchoCtx(c)
	entity := types.NewSystemPrompt(nil)
	if err = c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err = entity.Set(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, entity)
}

func HandleAPIListSystemPrompt(c echo.Context) error {
	ctx := GetEchoCtx(c)
	prompt := types.NewSystemPrompt(nil)
	prompts, err := prompt.List(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, prompts)
}

func HandleSystemPrompts(c echo.Context) error {
	dt := DisplaySystemPrompt{
		SystemPrompt: types.SystemPrompt{},
		DisplayType: "none",
		Menu: Menu{
			Href: "systemprompts",
			Title: "System Prompt",
		},
	}
	return c.Render(http.StatusOK, "systemprompts.tpl", dt)
}

func HandleSystemPromptsNew(c echo.Context) error {
	dt := DisplaySystemPrompt{
		SystemPrompt: types.SystemPrompt{},
		DisplayType: "new",
		Menu: Menu{
			Href: "systemprompts",
			Title: "System Prompt",
		},
	}
	return c.Render(http.StatusOK, "systemprompts.tpl", dt)
}

func HandleSystemPromptSave(c echo.Context) error {
	var err error
	var prompt types.SystemPrompt
	ctx := database.GetDatabaseCtx()
	if id := c.Param("id"); id != "" {
		prompt = types.NewSystemPrompt(&id)
		if err := prompt.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	} else {
		prompt = types.NewSystemPrompt(nil)
	}
	if err = c.Bind(&prompt); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandleSystemPromptSave"}.Wrap(err))
	}
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if err = prompt.Set(ctx); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}

	dt := DisplaySystemPrompt{
		SystemPrompt: types.SystemPrompt{},
		DisplayType: "new",
		Menu: Menu{
			Href: "systemprompts",
			Title: "System Prompt",
		},
	}
	return c.Render(http.StatusCreated, "systemprompts.tpl", dt)
}

func HandleSystemPromptsList(c echo.Context) error {
	var err error
	ctx := GetEchoCtx(c)
	prompt := types.NewSystemPrompt(nil)
	prompts, err := prompt.List(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}

	dt := DisplaySystemPrompt{
		SystemPrompt: types.SystemPrompt{},
		List: prompts,
		DisplayType: "list",
		Menu: Menu{
			Href: "systemprompts",
			Title: "System Prompt",
		},
	}
	return c.Render(http.StatusOK, "systemprompts.tpl", dt)
}

func HandleSystemPromptsGet(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		prompt := types.NewSystemPrompt(&id)
		if err := prompt.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt := DisplaySystemPrompt{
			SystemPrompt: types.SystemPrompt{},
			DisplayType: "new",
			Menu: Menu{
				Href: "systemprompts",
				Title: "System Prompt",
			},
		}
		return c.Render(http.StatusOK, "systemprompts.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleSystemPromptDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewSystemPrompt(&id)
		if err := entity.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleSystemPromptsList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}
