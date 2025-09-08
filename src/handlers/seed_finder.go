package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/strrep"
	"github.com/mmarchio/management/types"
)

func RegisterSeedFinderRoutes(e *echo.Echo) {
	g := e.Group("/seedfinder")
	g.GET("/configure", HandleSeedFinderConfigure)
	// g.GET("/new", HandleWorkflowNew)
	// g.POST("/save", HandleWorkflowSave)
	// g.POST("/save/:id", HandleWorkflowSave)
	// g.GET("/list", HandleWorkflowList)
	// g.GET("/edit/:id", HandleWorkflowEdit)
	// g.GET("/delete/:id", HandleWorkflowDelete)
	g.POST("/run", HandleSeedFinderRun)
	g.GET("/set/:jobid/:stepid/:seed", HandleSeedFinderSet)
}

func HandleSeedFinderConfigure(c echo.Context) error {
	dt := DisplaySeedFinder{}
	dt.Init(c, "configure")
	data := make(map[string]interface{})
	data["dt"] = dt
	return c.Render(http.StatusOK, "seedfinder.tpl", data)
}

func HandleSeedFinderRun(c echo.Context) error {
	var err error
	dt := DisplaySeedFinder{}
	dt.Init(c, "run")
	data := make(map[string]interface{})
	data["dt"] = dt
	if wfid := c.FormValue("workflow_id"); wfid != "" {
		wf := types.NewWorkflow(&wfid)
		if err := wf.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", merrors.ContentGetError{}.Wrap(err).Log())
		}
		if len(wf.ComfyNodesArrayModel) == 1 {
			var limit int = 8
			if num := c.FormValue("number"); num != "" {
				limit, err = strconv.Atoi(num)
				if err != nil {
					return c.Render(http.StatusInternalServerError, "error.tpl", merrors.GeneralError{}.Wrap(err).Log().Error())
				}
			}
			node := wf.ComfyNodesArrayModel[0]
			for i := 0; i < limit; i++ {
				clientID := uuid.NewString()
				seed := rand.Int32()
				templateValues := make(map[string]interface{})
				templateValues["seed"] = seed
				templateValues["segment"] = "/ComfyUI/input/seedfinder/ComfyUI_00001_.mp3"
				prompt, err := strrep.Strrep(node.APITemplate, templateValues)
				if err != nil {
					return c.Render(http.StatusInternalServerError, "error.tpl", merrors.GeneralError{}.Wrap(err).Log())
				}
				cp := types.ComfyMessage {
					ClientID: clientID,
					Prompt: prompt,
				}
				cpb, err := json.Marshal(cp)
				if err != nil {
					return c.Render(http.StatusInternalServerError, "error.tpl", merrors.JSONMarshallingError{}.Wrap(err).Log())
				}
				crd, err := node.QueuePrompt(cpb)
				if err != nil {
					return c.Render(http.StatusInternalServerError, "error.tpl", merrors.GeneralError{}.Wrap(err).Log())
				}
				if crd != nil {
					crdp := *crd
					patternString := "\"filename\".+" +regexp.QuoteMeta(".") + "mp4"
					pattern := regexp.MustCompile(patternString)
					file := pattern.Find([]byte(crdp.Raw))
					fullpath := strings.Replace(string(file), "\"filename\": \"", "", 1)

					fullpath = strings.Replace(fullpath, "-audio.mp4\"", ".mp4", 1)
					data[fmt.Sprintf("%d", seed)] = fullpath
				}
				GetLogger(3).Flogger("%d of %d", i, limit+1)
			}
		} else {
			GetLogger(3).Flogger("wf.comfynode len: %d", len(wf.ComfyNodesArrayModel))
		}
	}
	GetLogger(3).Flogger("returning: %#v", data)
	err = c.Render(http.StatusOK, "seedfinder.tpl", data)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	return err
}

func HandleSeedFinderSet(c echo.Context) error {
	if jobid := c.Param("jobid"); jobid != "" {
		job := types.NewJob(&jobid)
		if err := job.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", merrors.ContentGetError{}.Wrap(err).Log().Error())
		}
		if stepid := c.Param("stepid"); stepid != "" {
			step := types.NewStep(&stepid)
			if err := step.Get(c); err != nil {
				return c.Render(http.StatusInternalServerError, "error.tpl", merrors.ContentGetError{}.Wrap(err).Log().Error())
			}
			if s := c.Param("seed"); s != "" {
				seed, err := strconv.Atoi(s)
				if err != nil {
					return c.Render(http.StatusInternalServerError, "error.tpl", merrors.GeneralError{}.Wrap(err).Log().Error())
				}
				if job.SeedMap == nil {
					job.SeedMap = make(map[string]int32)
				}
				job.SeedMap[step.Node] = int32(seed)
				if err := job.Set(c, true); err != nil {
					return c.Render(http.StatusInternalServerError, "error.tpl", merrors.ContentSetError{}.Wrap(err).Log().Error())
				}
			}
		}
	}
	return HandleSeedFinderConfigure(c)
}

