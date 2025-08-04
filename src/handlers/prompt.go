package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mmarchio/management/database"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/types"
)

func RegisterPromptsRoutes(e *echo.Echo) {
	g := e.Group("/prompts")
	g.GET("", HandlePrompts)
	g.GET("/new", HandlePromptsNew)
	g.POST("/save", HandlePromptSave)
	g.POST("/save/:id", HandlePromptSave)
	g.GET("/list", HandlePromptsList)
	g.GET("/edit/:id", HandlePromptsGet)
	g.GET("/delete/:id", HandlePromptDelete)
}

func HandleAPIGetPrompt(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		prompt := types.NewPrompt(&id)
		if err := prompt.Get(ctx); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}		
		return c.JSON(http.StatusOK, prompt)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPISetPrompt(c echo.Context) error {
	var err error
	ctx := GetEchoCtx(c)
	entity := types.NewPrompt(nil)
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

func HandleAPIListPrompt(c echo.Context) error {
	ctx := GetEchoCtx(c)
	prompt := types.NewPrompt(nil)
	prompts, err := prompt.List(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, prompts)
}

func HandlePrompts(c echo.Context) error {
	dt := DisplayPrompt{
		Prompt: types.Prompt{},
		DisplayType: "none",
		Menu: Menu{
			Href: "prompts",
			Title: "Prompt",
		},
	}
	return c.Render(http.StatusOK, "prompts.tpl", dt)
}

func HandlePromptsNew(c echo.Context) error {
	var err error
	ctx := GetEchoCtx(c)
	prompt := types.NewPrompt(nil)
	prompt.ID = types.PromptID("")
	prompt.Model.ID = ""
	if id := c.Param("id"); err != nil {
		prompt = types.NewPrompt(&id)
	}
	prompt, err = prompt.GetDispositions(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	wf := types.NewWorkflow(nil)
	wfs, err := wf.List(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt := DisplayPrompt{
		Prompt: prompt,
		Workflows: wfs,
		DisplayType: "new",
		Menu: Menu{
			Href: "prompts",
			Title: "Prompt",
		},
	}
	return c.Render(http.StatusOK, "prompts.tpl", dt)
}

func HandlePromptSave(c echo.Context) error {
	var err error
	ctx := database.GetDatabaseCtx()
	prompt := types.NewPrompt(nil)
	if id := c.Param("id"); id != "" {
		//existing entity
		prompt = types.NewPrompt(&id)
		if err := prompt.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	if err = c.Bind(&prompt); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandlePromptSave"}.Wrap(err))
	}
	prompt, err = prompt.GetDispositions(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	prompt, err = prompt.Bind(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if err = prompt.Set(ctx); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	prompt.Settings.Template.AvailableDispositions = reconcileDispositions(prompt)

	existingJob := types.NewJob(nil)
	existingJob.PromptID = prompt.ID
	job, err := findExistingJob(ctx, existingJob.PromptID.String(), prompt)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if err = createPromptJobRuns(ctx, job, prompt); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	return HandlePromptsList(c)
}

func reconcileDispositions(prompt types.Prompt) []types.Disposition {
	newAvailableDispositions := make([]types.Disposition, 0)
	for _, d := range prompt.Settings.Template.Dispositions {
		exists := false
		for _, ad := range prompt.Settings.Template.AvailableDispositions {
			if d.Model.ID == ad.Model.ID {
				exists = true
			}
		}
		if !exists {
			newAvailableDispositions = append(newAvailableDispositions, d)
		}
	}
	return newAvailableDispositions
}

func findExistingJob(ctx context.Context, placeholderID string, prompt types.Prompt) (*types.Job, error){
	findbyJob := types.Job{}
	findbyJob.PromptID = prompt.ID
	foundJob, err := findbyJob.FindBy(ctx, "prompt_id", prompt.ID.String())
	if err != nil {
		if e, ok := err.(merrors.WrappedError); ok {
			if e.GetCode() != merrors.ErrorCode(404) {
				return nil, err
			}
			if e.GetCode() == merrors.ErrorCode(404) {
				newJob, err := createPromptJob(ctx, prompt)
				if err != nil {
					return nil, err
				}
				return newJob, nil
			}
		}
	}
	if foundJob.Model.ID == "" {
		newJob, err := createPromptJob(ctx, prompt)
		if err != nil {
			return nil, err
		}
		return newJob, nil
	}
	var job types.Job
	if foundJob.PromptID.String() != placeholderID {
		job = types.NewJob(nil)
		job.PromptID = types.PromptID(prompt.Model.ID)
		job.WorkflowID = prompt.Settings.Workflow
		if err := job.Set(ctx); err != nil {
			return nil, err
		}
	} else {
		job = foundJob
	}
	return &job, nil
}

func updateJobRuns(ctx context.Context, jobruns []types.JobRun, jobrun types.JobRun, prompt types.Prompt) error {
	for _, jobrun = range jobruns {
		if prompt.Settings.UpdatedAt != jobrun.Settings.UpdatedAt {
			jobrun.Settings = prompt.Settings
			jobrun.Model.UpdatedAt = time.Now()
		}
		for _, d := range prompt.Settings.Template.Dispositions {
			if jobrun.Disposition.Model.ID == d.Model.ID && jobrun.Disposition.UpdatedAt != d.Model.UpdatedAt {
				jobrun.Disposition = d
				jobrun.Model.UpdatedAt = time.Now()
			}
		}
		jobrun.Model.ID = jobrun.ID.String()
		if err := jobrun.Set(ctx); err != nil {
			return err
		}
	}
	return nil
}

func createJobRuns(ctx context.Context, prompt types.Prompt, job *types.Job) error {
	for _, disposition := range prompt.Settings.Template.Dispositions {
		jobrun := types.NewJobRun(nil)
		jobrun.Model.ContentType = "jobrun"
		jobrun.JobID = types.JobID(job.Model.ID)
		jobrun.Context = types.NewContext(prompt, jobrun.ID, disposition)
		jobrun.LatestStatusType = "start"
		jobrun.LatestStatusValue = "queued"
		jobrun.Disposition = disposition
		if err := jobrun.Set(ctx); err != nil {
			return err
		}
	}
	return nil
}

func createPromptJob(ctx context.Context, prompt types.Prompt) (*types.Job, error) {
	newJob := types.NewJob(nil)
	newJob.PromptID = prompt.ID
	newJob.Recurring = prompt.Settings.Recurring.Value
	newJob.Interval = prompt.Settings.Interval
	newJob.WorkflowID = prompt.Settings.Workflow
	if err := newJob.Set(ctx); err != nil {
		return nil, err
	}
	return &newJob, nil
}

func createPromptJobRuns(ctx context.Context, job *types.Job, prompt types.Prompt) error {
	jobrun := types.NewJobRun(nil)
	jobruns, err := jobrun.ListBy(ctx, "job_id", job.Model.ID)
	if err != nil {
		return err
	}
	if len(jobruns) > 0 {
		//update existing job runs
		if err = updateJobRuns(ctx, jobruns, jobrun, prompt); err != nil {
			return err
		}
	} else {
		//create new job runs
		if err = createJobRuns(ctx, prompt, job); err != nil {
			return err
		}
	}
	return nil
}

func HandlePromptsList(c echo.Context) error {
	var err error
	ctx := GetEchoCtx(c)
	prompt := types.NewPrompt(nil)
	prompts, err := prompt.List(ctx)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	dt := DisplayPrompt{
		Prompt: prompt,
		List: prompts,
		DisplayType: "list",
		Menu: Menu{
			Href: "prompts",
			Title: "Prompt",
		},
	}
	return c.Render(http.StatusOK, "prompts.tpl", dt)
}

func HandlePromptsGet(c echo.Context) error {
	var err error
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		prompt := types.NewPrompt(&id)
		if err = prompt.Get(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		content := types.NewPromptTypeContent()
		if err = content.FromType(prompt); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		prompt, err = prompt.GetDispositions(ctx)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		dt := DisplayPrompt{
			Prompt: prompt,
			DisplayType: "new",
			Menu: Menu{
				Href: "prompts",
				Title: "Prompt",
			},
		}
		return c.Render(http.StatusOK, "prompts.tpl", dt)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func HandlePromptDelete(c echo.Context) error {
	ctx := GetEchoCtx(c)
	if id := c.Param("id"); id != "" {
		entity := types.NewPrompt(&id)
		if err := entity.Delete(ctx); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandlePromptsList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func debug(msg string) {
	fmt.Printf(msg)
}
