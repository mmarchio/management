package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		jobRun := types.NewJobRun(&id)
		if err := jobRun.Get(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal server error: %w", err))
		}
		return c.JSON(http.StatusOK, jobRun)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPIListJobRun(c echo.Context) error {
	ctx := GetEchoCtx(c)
	jobRun := types.NewJobRun(nil)
	jobRuns, err := jobRun.List(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal server error: %w", err))
	}
	return c.JSON(http.StatusOK, jobRuns)
}

func HandleAPIListJobRunBy(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		jobRun := types.NewJobRun(nil)
		jobruns, err := jobRun.ListBy(ctx, "job_id", id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal server error: %w", err))
		}
		return c.JSON(http.StatusOK, jobruns)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPISaveJobRun(c echo.Context) error {
	ctx := GetEchoCtx(c)
	job := types.NewJobRun(nil)
	if err := c.Bind(&job); err != nil {
		return c.JSON(http.StatusInternalServerError, merrors.EchoBindError{Package: "handlers", Function: "HandleAPISaveJobRun"}.Wrap(err))
	}
	if err := job.Set(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, job)
}

func HandleJobRuns(c echo.Context) error {
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
	ctx := GetEchoCtx(c)
	jobRun := types.NewJobRun(nil)
	jobRuns, err := jobRun.List(ctx)
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
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewJobRun(&id)
		if err := entity.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandleJobRunsList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleJobRunsContextGet(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewJobRun(&id)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		b, err := json.MarshalIndent(entity.TruncatedContextModel, "", "  ")
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", merrors.JSONMarshallingError{}.Wrap(err))
		}
		return c.Render(http.StatusOK, "jobrun.context.tpl", string(b))
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleAPIJobRunsContextGet(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewJobRun(&id)
		if err := entity.Get(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if entity.TruncatedContextModel.ID == "" {
			tc, err := entity.ContextModel.Truncate()
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
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewJobRun(&id)
		if err := entity.Get(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if entity.TruncatedContextModel.ID == "" {
			tc, err := entity.ContextModel.Truncate()
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
		if err := entity.Set(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, entity.ID)
	}
	return c.JSON(http.StatusBadRequest, "bad request: missing id")
}

func HandleJobRunsContextSet(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		jobRun := types.NewJobRun(&id)
		if err := jobRun.Get(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		Ctx, err := jobRun.ContextModel.Truncate()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, Ctx)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPINextJobRun(c echo.Context) error {
	ctx := GetEchoCtx(c)
	entity := types.NewJobRun(nil)
	q := fmt.Sprintf("SELECT id, created_at, updated_at, content_type, content FROM content WHERE content_type = 'jobrun' AND content @> '{\"latest_status_type\":\"start\"}' AND content @> '{\"latest_status_value\":\"queued\"}' ORDER BY updated_at ASC LIMIT 1")
	res, err := entity.CustomQuery(ctx, false, q)
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
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewJobRun(&id)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wfid := entity.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		if err := wf.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleJobRunRun(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if entityID := c.Param("id"); entityID != "" {
		entity := types.NewJobRun(&entityID)
		if err := entity.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		jobid := entity.JobID.String()
		job := types.NewJob(&jobid)
		if err := job.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wfid := job.WorkflowID.String()
		wf := types.NewWorkflow(&wfid)
		wf.ID = job.WorkflowID
		if err := wf.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		for k := range wf.NodeOrder {
			parts := strings.Split(k, ":")
			if len(parts) == 2 {
				switch parts[1] {
				case "comfynode":
				case "ollamanode":
					onode := types.NewOllamaNode(nil)
					onode.GetNodeFromWorkflow(parts[0], wf)
					if err := onode.Exec(ctx); err != nil {
						return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
					}
				case "sshnode":
				}
			}
		}

		return HandleJobRunsList(c)
	}
	return c.Render(http.StatusInternalServerError, "error.tpl", "unknown node type")
}

func run(c echo.Context, resp types.OllamaResponse, Response chan types.OllamaResponse, Error chan error, sleep int, semaphore chan struct{}) {
	fmt.Printf("websocket worker starting\n")
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
	fmt.Printf("websocket worker exiting\n")
}