package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mmarchio/management/config"
	"github.com/mmarchio/management/handlers"
	"github.com/mmarchio/management/logger"
	"github.com/mmarchio/management/types"
)

func GetLogger(severity int) logger.LoggingContext {
	r := logger.LoggingContext{Severity: severity}
	r.Init()
	return r
}

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.Templates.ExecuteTemplate(w, name, data)
	if err != nil {
		terr := fmt.Errorf("Template rendering error: %w", err) // Log the error
		GetLogger(0).Flogger("%s, %#v, %#v", err.Error(), terr, data)
		// You can choose to send a generic error page or a plain string
		return c.String(http.StatusInternalServerError, "Error rendering template.")
	}
	return nil
}

func main() {
	e := echo.New()

	e.GET("/", handleIndex)
	e.GET("/js", handleJSIndex)

	handlers.RegisterAPIPageRoutes(e)

	// e.GET("/swagger", echoSwagger.WrapHandler)

	// e.GET("/debug/content/view/:id", handlers.HandleDebugContentView)

	// e.GET("/api/prompt/:id", handlers.HandleAPIGetPrompt)
	// e.POST("/api/prompt", handlers.HandleAPISetPrompt)
	// e.GET("/api/prompts", handlers.HandleAPIListPrompt)

	// e.GET("/api/comfy/:id", handlers.HandleAPIGetComfyUITemplate)
	// e.POST("/api/comfy", handlers.HandleAPISetComfyUITemplate)
	// e.GET("/api/comfys", handlers.HandleAPIListComfyUITemplate)

	// e.GET("/api/systemprompt/:id", handlers.HandleAPIGetSystemPrompt)
	// e.POST("/api/systemprompt", handlers.HandleAPISetSystemPrompt)
	// e.GET("/api/systemprompts", handlers.HandleAPIListSystemPrompt)

	// e.GET("/api/job/:id", handlers.HandleAPIGetJob)
	// e.GET("/api/jobs", handlers.HandleAPIListJob)
	// e.POST("/api/jobs/set", handlers.HandleAPISaveJob)

	// e.GET("/api/jobrun/:id", handlers.HandleAPIGetJobRun)
	// e.GET("/api/jobruns", handlers.HandleAPIListJobRun)
	// e.GET("/api/jobruns/:id", handlers.HandleAPIListJobRunBy)
	// e.POST("/api/jobruns/set", handlers.HandleAPISaveJobRun)
	// e.GET("/api/jobruns/next", handlers.HandleAPINextJobRun)
	// e.GET("/api/jobruns/context/:id", handlers.HandleAPIJobRunsContextGet)
	// e.POST("/api/jobruns/context/:id", handlers.HandleAPIJobRunsContextSet)

	// e.GET("/api/disposition/:id", handlers.HandleAPIGetDisposition)
	// e.POST("/api/disposition", handlers.HandleAPISetDisposition)
	// e.GET("/api/dispositions", handlers.HandleAPIListDisposition)

	// e.GET("/:ContentType", handlers.ContentType)
	// e.GET("/:ContentType/new", handlers.ContentTypeNew)
	// e.POST("/:ContentType/save", handlers.ContentTypeSave)
	// e.POST("/:ContentType/save/:id", handlers.ContentTypeSave)
	// e.GET("/:ContentType/list", handlers.ContentTypeList)
	// e.GET("/:ContentType/edit/:id", handlers.ContentTypeEdit)
	// e.GET("/:ContentType/delete/:id", handlers.ContentTypeDelete)


	handlers.RegisterWorkflowRoutes(e)
	handlers.RegisterNodesRoutes(e)
	handlers.RegisterPromptsRoutes(e)
	handlers.RegisterComfyUITemplatesRoutes(e)
	handlers.RegisterSystemPromptsRoutes(e)
	handlers.RegisterJobRoutes(e)
	handlers.RegisterJobRunRoutes(e)
	handlers.RegisterDispositionRoutes(e)
	handlers.RegisterPromptTemplateRoutes(e)
	handlers.RegisterStepRoutes(e)
	handlers.RegisterSeedFinderRoutes(e)


	//	e.Use(middleware.Static("/public/static"))
    e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
        Root:       "public/static",
        Browse:     false,
        IgnoreBase: false,
    }))

	t := &Template{
		Templates: template.Must(template.New("").Funcs(template.FuncMap{
			"prettyJSON": PrettyJSON,
			"contains": Contains,
		}).ParseGlob("public/views/*.tpl")),
	}

	e.Renderer = t

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)))	
}

func handleIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index.tpl", nil)
}

func handleJSIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index.js.tpl", nil)
}


func PrettyJSON(v interface{}) (template.HTML, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return template.HTML(fmt.Sprintf("<pre>%s</pre>", strings.Replace(string(b), "\\\"", "\"", -1))), nil
}

func Contains(needle string, haystack []types.Disposition) bool {
	for _, h := range haystack {
		if needle == h.Model.ID {
			return true
		}
	}
	return false
}