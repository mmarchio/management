package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
	"golang.org/x/net/websocket"
)

func RegisterJobRunRoutes(e *echo.Echo) {
	g := e.Group("/jobruns")
	g.GET("", HandleJobRuns)
	g.GET("/new", HandleJobRuns)
	g.GET("/list", HandleJobRunsList)
	g.GET("/delete/:id", HandleJobRunsDelete)
	g.GET("/context/:id", HandleJobRunsContextGet)
	g.GET("/run/:id", HandleJobRunRun)
}

var wg sync.WaitGroup

func HandleAPIGetJobRun(c echo.Context) error {
	GetLogger().Flogger("HandleAPIGetJobRun called")
	if id := c.Param("id"); id != "" {
		jobRun := types.NewJobRun(&id)
		if err := jobRun.Get(c); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, jobRun)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPIListJobRun(c echo.Context) error {
	GetLogger().Flogger("HandleAPIListJobRun called")
	jobRun := types.NewJobRun(nil)
	jobRuns, err := jobRun.List(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, jobRuns)
}

func HandleAPIListJobRunBy(c echo.Context) error {
	GetLogger().Flogger("HandleAPIListJobRunBy called")
	if id := c.Param("id"); id != "" {
		jobRun := types.NewJobRun(nil)
		jobruns, err := jobRun.ListBy(c, "job_id", id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, jobruns)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPISaveJobRun(c echo.Context) error {
	GetLogger().Flogger("HandleAPISaveJobRun called")
	job := types.NewJobRun(nil)
	if err := c.Bind(&job); err != nil {
		return c.JSON(http.StatusInternalServerError, merrors.EchoBindError{Package: "handlers", Function: "HandleAPISaveJobRun"}.Wrap(nil, err))
	}
	if err := job.Set(c); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, job)
}

func HandleJobRuns(c echo.Context) error {
	GetLogger().Flogger("HandleJobRuns called")
	dt := DisplayJobRun{
		JobRun: types.JobRun{},
		DisplayType: "none",
		Menu: Menu{
			Href: "jobruns",
			Title: "Job Run",
		},
	}
	return c.Render(http.StatusOK, "jobruns.tpl", dt)
}

func HandleJobRunsList(c echo.Context) error {
	GetLogger().Flogger("HandleJobRunsList called")
	jobRun := types.NewJobRun(nil)
	jobRuns, err := jobRun.List(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt := DisplayJobRun{
		JobRun: types.JobRun{},
		List: jobRuns,
		DisplayType: "list",
		Menu: Menu{
			Href: "jobruns",
			Title: "Job Run",
		},
	}
	return c.Render(http.StatusOK, "jobruns.tpl", dt)
}

func HandleJobRunsDelete(c echo.Context) error {
	GetLogger().Flogger("HandleJobRunsDelete called")
	if id := c.Param("id"); id != "" {
		entity := types.NewJobRun(&id)
		if err := entity.Delete(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleJobRunsList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleJobRunsContextGet(c echo.Context) error {
	GetLogger().Flogger("HandleJobRunsContextGet called")
	if id := c.Param("id"); id != "" {
		entity := types.NewJobRun(&id)
		if err := entity.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		b, err := json.MarshalIndent(entity.TruncatedContextModel, "", "  ")
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", merrors.JSONMarshallingError{}.Wrap(nil, err))
		}
		return c.Render(http.StatusOK, "jobrun.context.tpl", string(b))
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleAPIJobRunsContextGet(c echo.Context) error {
	GetLogger().Flogger("HandleAPIJobRunsContextGet called")
	if id := c.Param("id"); id != "" {
		entity := types.NewJobRun(&id)
		if err := entity.Get(c); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if entity.TruncatedContextModel.ID == "" {
			tc, err := entity.ContextModel.Truncate(c)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			if tc != nil {
				entity.TruncatedContextModel = *tc
			}
		}
		return c.JSON(http.StatusOK, entity.TruncatedContextModel)
	}
	return c.JSON(http.StatusBadRequest, "bad request: missing id")
}

func HandleAPIJobRunsContextSet(c echo.Context) error {
	GetLogger().Flogger("HandleAPIJobRunsContextSet called")
	if id := c.Param("id"); id != "" {
		entity := types.NewJobRun(&id)
		if err := entity.Get(c); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if entity.TruncatedContextModel.ID == "" {
			tc, err := entity.ContextModel.Truncate(c)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			if tc != nil {
				entity.TruncatedContextModel = *tc
			}
		}
		tc := types.TruncatedContext{}
		if err := json.NewDecoder(c.Request().Body).Decode(&tc); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if err := entity.Set(c); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, entity.ID)
	}
	return c.JSON(http.StatusBadRequest, "bad request: missing id")
}

func HandleJobRunsContextSet(c echo.Context) error {
	GetLogger().Flogger("HandleJobRunsContextSet called")
	if id := c.Param("id"); id != "" {
		jobRun := types.NewJobRun(&id)
		if err := jobRun.Get(c); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		Ctx, err := jobRun.ContextModel.Truncate(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, Ctx)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPINextJobRun(c echo.Context) error {
	GetLogger().Flogger("HandleAPINextJobRun called")
	entity := types.NewJobRun(nil)
	q := fmt.Sprintf("SELECT id, created_at, updated_at, content_type, content FROM content WHERE content_type = 'jobrun' AND content @> '{\"latest_status_type\":\"start\"}' AND content @> '{\"latest_status_value\":\"queued\"}' ORDER BY updated_at ASC LIMIT 1")
	res, err := entity.CustomQuery(c, false, q)
	if err != nil {
		if e, ok := err.(merrors.WrappedError); ok {
			if e.GetCode() == merrors.ErrorCode(404) || len(res) == 0 {
				return c.JSON(http.StatusNotFound, "no jobruns found")
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res[0])
}

func HandleWorkflowRun(c echo.Context) error {
	GetLogger().Flogger("HandleWorkflowRun called")
	if id := c.Param("id"); id != "" {
		entity := types.NewJobRun(&id)
		if err := entity.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wfid := entity.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		if err := wf.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleJobRunRun(c echo.Context) error {
	GetLogger().Flogger("HandleJobRunRun called")
	if entityID := c.Param("id"); entityID != "" {
		entity := types.NewJobRun(&entityID)
		if err := entity.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		jobid := entity.JobID.String()
		job := types.NewJob(&jobid)
		if err := job.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wfid := job.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		wf.ID = job.WorkflowID
		if err := wf.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wf.NodeCleanup(c)
		for _, node := range wf.NodeOrder {
				GetLogger().Flogger("wf.NodeOrder:type %s", node.NodeType)
				switch node.NodeType {
				case "comfynode":
				case "ollamanode":
					for _, v := range wf.OllamaNodesArrayModel {
						if v.ID == node.NodeID {
							if err := v.Exec(c, &entity); err != nil {
								return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
							}
						}
					}
				case "sshnode":
				}
		}
	}
		return HandleJobRunsList(c)
}

func run(c echo.Context, resp types.OllamaResponse, Response chan types.OllamaResponse, Error chan error, sleep int, semaphore chan struct{}) {
	GetLogger().Flogger("run called")
	GetLogger().Flogger("websocket worker starting\n")
	defer wg.Done()
	defer func() { <-semaphore }()

	semaphore <- struct{}{}
	websocket.Handler(func(ws *websocket.Conn){
		defer ws.Close()
		for {
			//write
			if len(Response) > 0 {
				resp = <-Response
				if err := websocket.Message.Send(ws, resp.Response); err != nil {
					c.Logger().Error(err)
				}
			}
			if len(Error) > 0 {
				if err := <-Error; err != nil {
					c.Logger().Error(err)
				}
			}
			//read
			msg := ""
			if err := websocket.Message.Receive(ws, &msg); err != nil {
				c.Logger().Error(err)
			}
			time.Sleep(time.Duration(time.Duration(sleep)*time.Second))
		}
	}).ServeHTTP(c.Response(), c.Request())
	GetLogger().Flogger("websocket worker exiting\n")
}