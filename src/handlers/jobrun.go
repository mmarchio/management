package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "sync"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

func RegisterJobRunRoutes(e *echo.Echo) {
	g := e.Group("/jobruns")
	g.GET("", HandleJobRuns)
	g.GET("/new", HandleJobRuns)
	g.GET("/edit/:id", HandleJobRunEdit)
	g.POST("/save/:id", HandleJobRunsSave)
	g.GET("/list", HandleJobRunsList)
	g.GET("/delete/:id", HandleJobRunsDelete)
	g.GET("/context/:id", HandleJobRunsContextGet)
	g.GET("/run/:id", HandleJobRunRun)
}

//var wg sync.WaitGroup

func HandleAPIGetJobRun(c echo.Context) error {
	GetLogger(4).Flogger("HandleAPIGetJobRun called")
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
	GetLogger(4).Flogger("HandleAPIListJobRun called")
	jobRun := types.NewJobRun(nil)
	jobRuns, err := jobRun.List(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, jobRuns)
}

func HandleAPIListJobRunBy(c echo.Context) error {
	GetLogger(4).Flogger("HandleAPIListJobRunBy called")
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
	GetLogger(4).Flogger("HandleAPISaveJobRun called")
	var update bool
	if id := c.Param("id"); id != "" {
		update = true
	}
	job := types.NewJobRun(nil)
	if err := c.Bind(&job); err != nil {
		return c.JSON(http.StatusInternalServerError, merrors.EchoBindError{Package: "handlers", Function: "HandleAPISaveJobRun"}.Wrap(err))
	}
	if err := job.Set(c, update); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, job)
}

func HandleJobRuns(c echo.Context) error {
	GetLogger(4).Flogger("HandleJobRuns called")
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

func HandleJobRunEdit(c echo.Context) error {
	if id := c.Param("id"); id != "" {
		GetLogger(4).Flogger("##########Getting Job Run##########")
		entity := types.NewJobRun(&id)
		if err := entity.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		GetLogger(4).Flogger("##########Getting Workflow##########")
		wfid := entity.ContextModel.SettingsModel.Workflow.String()
		GetLogger(4).Flogger("wfid: %s", wfid)
		wf := types.NewWorkflow(&wfid)
		if err := wf.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		GetLogger(4).Flogger("##########Getting Steps##########")
		steps, err := entity.GetSteps(c)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		GetLogger(4).Flogger("##########Getting Nodes##########")
		nodes, err := entity.GetNodes(c)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt := DisplayJobRun{
			JobRun: entity,
			Menu: Menu{
				Href: "jobrun",
				Title: "Job Run",
			},
			Workflow: wf,
			DisplayType: "edit",
			Steps: steps,
			Nodes: nodes,
		}
		GetLogger(4).Flogger("steps: %#v", steps)
		GetLogger(4).Flogger("nodes: %#v", nodes)
		return c.Render(http.StatusOK, "jobruns.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandleJobRunsSave(c echo.Context) error {
	GetLogger(4).Flogger("HandleJobRunsSave called")
	var update bool
	jobrun := types.NewJobRun(nil)
	if id := c.Param("id"); id != "" {
		update = true
		jobrun = types.NewJobRun(&id)
		if err := jobrun.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	steps, err := jobrun.GetSteps(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if jobrun.NodeStepMap == nil {
		jobrun.NodeStepMap = make(map[string]string)
	}
	for stepid, _ := range steps {
		if nodeid := c.FormValue(stepid); nodeid != "" {
			jobrun.NodeStepMap[stepid] = nodeid
		}
	}
	if err := jobrun.Set(c, update); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	return HandleJobRunsList(c)
}

func HandleJobRunsList(c echo.Context) error {
	GetLogger(4).Flogger("HandleJobRunsList called")
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
	GetLogger(4).Flogger("HandleJobRunsDelete called")
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
	GetLogger(4).Flogger("HandleJobRunsContextGet called")
	if id := c.Param("id"); id != "" {
		entity := types.NewJobRun(&id)
		if err := entity.Get(c); err != nil {
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
	GetLogger(4).Flogger("HandleAPIJobRunsContextGet called")
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
	GetLogger(4).Flogger("HandleAPIJobRunsContextSet called")
	var update bool
	entity := types.NewJobRun(nil)
	if id := c.Param("id"); id != "" {
		update = true
		entity = types.NewJobRun(&id)
		if err := entity.Get(c); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
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
	if err := entity.Set(c, update); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, entity.ID)
}

func HandleJobRunsContextSet(c echo.Context) error {
	GetLogger(4).Flogger("HandleJobRunsContextSet called")
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
	GetLogger(4).Flogger("HandleAPINextJobRun called")
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
	GetLogger(4).Flogger("HandleWorkflowRun called")
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
	GetLogger(4).Flogger("HandleJobRunRun called")
	if entityID := c.Param("id"); entityID != "" {
		entity := types.NewJobRun(&entityID)
		GetLogger(4).Flogger("HandleJobRunRun Get JobRun")
		if err := entity.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		entity.ValueCache = make(map[string]interface{})
		jobid := entity.JobID.String()
		job := types.NewJob(&jobid)
		GetLogger(4).Flogger("HandleJobRunRun Get Job")
		if err := job.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wfid, err := GetWorkflowID(c, &entity)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wid := wfid.String()
		entity.WorkflowID = wfid
		wf := types.NewWorkflow(&wid)
		wf.Model.ID = job.WorkflowID.String()
		GetLogger(4).Flogger("HandleJobRunRun Get Workflow")
		if err := wf.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		wf.NodeCleanup(c)
		s := types.NewStep(nil)
		GetLogger(4).Flogger("HandleJobRunRun Get Step List")
		steps, err := s.ListBy(c, "", "workflow_id", wfid)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		orderedSteps := make([]types.Step, len(steps))
		for _, step := range steps {
			orderedSteps[step.Order-1] = step
		}
		for _, step := range orderedSteps {
			if step.Bypass.Value && !step.Stats.Output.IsNil() {
				entity.ValueCache[fmt.Sprintf("steps[%d].output.think", step.Order-1)] = step.Stats.Output.Think
				entity.ValueCache[fmt.Sprintf("steps[%d].output.prompt", step.Order-1)] = step.Stats.Output.Prompt
				entity.ValueCache[fmt.Sprintf("steps[%d].output.topics", step.Order-1)] = step.Stats.Output.Topics
				entity.ValueCache[fmt.Sprintf("steps[%d].output.output", step.Order-1)] = step.Stats.Output.Output
				GetLogger(3).Flogger("bypassing step %d", step.Order)
				continue
			}
			content := types.Content{}
			content.Model.ID = step.Node
			GetLogger(3).Flogger("HandleJobRunRun Get Node Content")
			GetLogger(3).Flogger("step.name: %s, step.node: %s", step.Name, step.Node)
			content, err := content.Get(c)
			if err != nil {
				return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
			}
			switch content.Model.ContentType {
			case "comfynode":
				GetLogger(4).Flogger("HandleJobRunRun Unwrapping comfynode")
				node := types.ComfyNode{}
				if err := json.Unmarshal([]byte(content.Content), &node); err != nil {
					return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
				}
				GetLogger(4).Flogger("HandleJobRunRun Executing comfynode")
				if err := node.Exec(c, &entity, &step); err != nil {
					return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
				}
			case "ollamanode":
				GetLogger(4).Flogger("HandleJobRunRun Unwrapping ollamanode")
				node := types.OllamaNode{}
				if err := json.Unmarshal([]byte(content.Content), &node); err != nil {
					return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
				}
				GetLogger(4).Flogger("HandleJobRunRun Executing ollamanode")
				entity.ValueCache["context.prompt_model.prompt"] = entity.ContextModel.PromptModel.Prompt
				if err := node.Exec(c, &entity, &step); err != nil {
					return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
				}
			case "sshnode":
				GetLogger(4).Flogger("HandleJobRunRun Unwrapping sshnode")
				node := types.SSHNode{}
				if err := json.Unmarshal([]byte(content.Content), &node); err != nil {
					return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
				}
				GetLogger(4).Flogger("HandleJobRunRun Executing sshnode")
				if err := node.Exec(c, &entity, &step); err != nil {
					return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
				}
			default:
			}
			GetLogger(3).Flogger("Done executing step: %d", step.Order)
			entity.ValueCache[fmt.Sprintf("steps[%d].output.think", step.Order-1)] = step.Stats.Output.Think
			entity.ValueCache[fmt.Sprintf("steps[%d].output.prompt", step.Order-1)] = step.Stats.Output.Prompt
			entity.ValueCache[fmt.Sprintf("steps[%d].output.topics", step.Order-1)] = step.Stats.Output.Topics
			entity.ValueCache[fmt.Sprintf("steps[%d].output.output", step.Order-1)] = step.Stats.Output.Output
			GetLogger(3).Flogger("entity.ValueCache[steps[%d].output.think]: %s", step.Order-1, step.Stats.Output.Think)
			GetLogger(3).Flogger("entity.ValueCache[steps[%d].output.prompt]: %s", step.Order-1, step.Stats.Output.Prompt)
			GetLogger(3).Flogger("entity.ValueCache[steps[%d].output.topics]: %s", step.Order-1, step.Stats.Output.Topics)
			GetLogger(3).Flogger("entity.ValueCache[steps[%d].output.output]: %s", step.Order-1, step.Stats.Output.Output)
		}
		GetLogger(3).Flogger("Done executing all nodes")
	}
	return HandleJobRunsList(c)
}

func GetWorkflowID(e echo.Context, jobrun *types.JobRun) (types.WorkflowID, error) {
	jid := jobrun.JobID.String()
	job := types.NewJob(&jid)
	if err := job.Get(e); err != nil {
		return "", merrors.ContentGetError{CalledBy: "handlers.GetWorkflowID"}.Wrap(err).Log()
	}
	pid := job.PromptID.String()
	prompt := types.NewPrompt(&pid)
	if err := prompt.Get(e); err != nil {
		return "", merrors.ContentGetError{CalledBy: "handlers.GetWorkflowID"}.Wrap(err).Log()
	}
	return prompt.SettingsModel.Workflow, nil
}

