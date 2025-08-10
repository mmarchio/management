package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/logger"
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
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandleAPIGetSystemPrompt called")
	}
	if id := c.Param("id"); id != "" {
		entity := types.NewSystemPrompt(&id)
		if err := entity.Get(c); err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal server error: %w", err))
		}
		return c.JSON(http.StatusOK, entity)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPISetSystemPrompt(c echo.Context) error {
	var err error
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandleAPISetSystemPrompt called")
	}
	entity := types.NewSystemPrompt(nil)
	if err = c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err = entity.Set(c); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, entity)
}

func HandleAPIListSystemPrompt(c echo.Context) error {
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandleAPIListSystemPrompt called")
	}
	prompt := types.NewSystemPrompt(nil)
	prompts, err := prompt.List(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, prompts)
}

func HandleSystemPrompts(c echo.Context) error {
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandleSystemPrompts called")
	}
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
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandleSystemPromptsNew called")
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

func HandleSystemPromptSave(c echo.Context) error {
	var err error
	var prompt types.SystemPrompt
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandleSystemPromptSave called")
	}
	if id := c.Param("id"); id != "" {
		prompt = types.NewSystemPrompt(&id)
		if err := prompt.Get(c); err != nil {
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
	if err = prompt.Set(c); err != nil {
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
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandleSystemPromptsList called")
	}
	prompt := types.NewSystemPrompt(nil)
	prompts, err := prompt.List(c)
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
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandleSystemPromptsGet called")
	}
	if id := c.Param("id"); id != "" {
		prompt := types.NewSystemPrompt(&id)
		if err := prompt.Get(c); err != nil {
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
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandleSystemPromptDelete called")
	}
	if id := c.Param("id"); id != "" {
		entity := types.NewSystemPrompt(&id)
		if err := entity.Delete(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleSystemPromptsList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}
