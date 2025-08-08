package handlers

import (
	"encoding/json"
	_ "fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

func HandleDebugContentView(c echo.Context) error {
	ctx := GetEchoCtx(c)
	content := types.Content{}
	if c.Param("id") != "" {
		content.Model.ID = c.Param("id")
		content.ID = content.Model.ID
	} else {
		return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
	}
	entity, err := content.Get(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err)
	}
	msi := make(map[string]interface{})
	if err := json.Unmarshal([]byte(entity.Content), &msi); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err)
	}
	switch entity.ContentType {
	case "jobrun":
		jobrun := types.NewJobRun(nil)
		if err := json.Unmarshal([]byte(entity.Content), &jobrun); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err)
		}
		jobrun.Model.ID = entity.Model.ID
		jobrun.ID = types.RunID(jobrun.Model.ID)
		jobrun.ContentType = "jobrun"
		
	}
	b, err := json.MarshalIndent(msi, "", "  ")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err)
	}
	return c.Render(http.StatusOK, "debug.content.tpl", string(b))


}