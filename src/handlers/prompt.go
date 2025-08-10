package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/logger"
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
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandleAPIGetPrompt called")
	}
	if id := c.Param("id"); id != "" {
		prompt := types.NewPrompt(&id)
		if err := prompt.Get(c); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}		
		return c.JSON(http.StatusOK, prompt)
	}
	return c.JSON(http.StatusBadRequest, "missing id")
}

func HandleAPISetPrompt(c echo.Context) error {
	var err error
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandleAPISetPrompt called")
	}
	entity := types.NewPrompt(nil)
	if err := c.Bind(&entity); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	entity, err = entity.SetID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err = entity.Set(c); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, entity)
}

func HandleAPIListPrompt(c echo.Context) error {
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandleAPIListPrompt called")
	}
	prompt := types.NewPrompt(nil)
	prompts, err := prompt.List(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, prompts)
}

func HandlePrompts(c echo.Context) error {
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandlePrompts called")
	}
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
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandlePromptsNew called")
	}
	prompt := types.NewPrompt(nil)
	prompt.ID = types.PromptID("")
	prompt.Model.ID = ""
	if id := c.Param("id"); id != "" {
		prompt = types.NewPrompt(&id)
	}
	prompt, err = prompt.GetDispositions(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	wf := types.NewWorkflow(nil)
	wfs, err := wf.List(c)
	b, _ := json.Marshal(prompt)
	msi := make(map[string]interface{})
	_ = json.Unmarshal(b, &msi)
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
		Debug: msi,
	}
	fmt.Printf("len available: %d\n", len(dt.Prompt.SettingsModel.TemplateModel.AvailableDispositions))
	for _, ad := range dt.Prompt.SettingsModel.TemplateModel.AvailableDispositions {
		fmt.Printf("ad.id: %s at.name: %s\n", ad.ID, ad.Name)
	}
	return c.Render(http.StatusOK, "prompts.tpl", dt)
}

func HandlePromptSave(c echo.Context) error {
	var err error
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandlePromptSave called")
	}
	prompt := types.NewPrompt(nil)
	if id := c.Param("id"); id != "" {
		//existing entity
		prompt = types.NewPrompt(&id)
		if err := prompt.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
	}
	if err = c.Bind(&prompt); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", merrors.EchoBindError{Package: "handlers", Function: "HandlePromptSave"}.Wrap(err))
	}
	prompt, err = prompt.GetDispositions(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	prompt, err = prompt.Bind(c)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if err = prompt.Set(c); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	prompt.SettingsModel.TemplateModel.AvailableDispositions = reconcileDispositions(prompt)

	existingJob := types.NewJob(nil)
	existingJob.PromptID = prompt.ID
	job, err := findExistingJob(c, existingJob.PromptID.String(), prompt)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	if err = createPromptJobRuns(c, job, prompt); err != nil {
		return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
	}
	return HandlePromptsList(c)
}

func reconcileDispositions(prompt types.Prompt) []types.Disposition {
	newAvailableDispositions := make([]types.Disposition, 0)
	for _, d := range prompt.SettingsModel.TemplateModel.DispositionsArrayModel {
		exists := false
		for _, ad := range prompt.SettingsModel.TemplateModel.AvailableDispositions {
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

func findExistingJob(e echo.Context, placeholderID string, prompt types.Prompt) (*types.Job, error){
	findbyJob := types.Job{}
	findbyJob.PromptID = prompt.ID
	foundJob, err := findbyJob.FindBy(e, "prompt_id", prompt.ID.String())
	if err != nil {
		if er, ok := err.(merrors.WrappedError); ok {
			if er.GetCode() != merrors.ErrorCode(404) {
				return nil, err
			}
			if er.GetCode() == merrors.ErrorCode(404) {
				newJob, err := createPromptJob(e, prompt)
				if err != nil {
					return nil, err
				}
				return newJob, nil
			}
		}
	}
	if foundJob.Model.ID == "" {
		newJob, err := createPromptJob(e, prompt)
		if err != nil {
			return nil, err
		}
		return newJob, nil
	}
	var job types.Job
	if foundJob.PromptID.String() != placeholderID {
		job = types.NewJob(nil)
		job.PromptID = types.PromptID(prompt.Model.ID)
		job.WorkflowID = prompt.SettingsModel.Workflow
		if err := job.Set(e); err != nil {
			return nil, err
		}
	} else {
		job = foundJob
	}
	return &job, nil
}

func updateJobRuns(e echo.Context, jobruns []types.JobRun, jobrun types.JobRun, prompt types.Prompt) error {
	for _, jobrun = range jobruns {
		if prompt.SettingsModel.UpdatedAt != jobrun.SettingsModel.UpdatedAt {
			jobrun.SettingsModel = prompt.SettingsModel
			jobrun.Model.UpdatedAt = time.Now()
		}
		for _, d := range prompt.SettingsModel.TemplateModel.DispositionsArrayModel {
			if jobrun.DispositionModel.Model.ID == d.Model.ID && jobrun.DispositionModel.UpdatedAt != d.Model.UpdatedAt {
				jobrun.DispositionModel = d
				jobrun.Model.UpdatedAt = time.Now()
			}
		}
		jobrun.Model.ID = jobrun.ID.String()
		if err := jobrun.Set(e); err != nil {
			return err
		}
	}
	return nil
}

func createJobRuns(e echo.Context, prompt types.Prompt, job *types.Job) error {
	for _, disposition := range prompt.SettingsModel.TemplateModel.DispositionsArrayModel {
		jobrun := types.NewJobRun(nil)
		jobrun.Model.ContentType = "jobrun"
		jobrun.JobID = types.JobID(job.Model.ID)
		jobrun.ContextModel = types.NewContext(prompt, jobrun.ID, disposition)
		jobrun.LatestStatusType = "start"
		jobrun.LatestStatusValue = "queued"
		jobrun.DispositionModel = disposition
		if err := jobrun.Set(e); err != nil {
			return err
		}
	}
	return nil
}

func createPromptJob(e echo.Context, prompt types.Prompt) (*types.Job, error) {
	newJob := types.NewJob(nil)
	newJob.PromptID = prompt.ID
	newJob.Recurring = prompt.SettingsModel.RecurringModel.Value
	newJob.Interval = prompt.SettingsModel.Interval
	newJob.WorkflowID = prompt.SettingsModel.Workflow
	if err := newJob.Set(e); err != nil {
		return nil, err
	}
	return &newJob, nil
}

func createPromptJobRuns(e echo.Context, job *types.Job, prompt types.Prompt) error {
	jobrun := types.NewJobRun(nil)
	jobruns, err := jobrun.ListBy(e, "job_id", job.Model.ID)
	if err != nil {
		return err
	}
	if len(jobruns) > 0 {
		//update existing job runs
		if err = updateJobRuns(e, jobruns, jobrun, prompt); err != nil {
			return err
		}
	} else {
		//create new job runs
		if err = createJobRuns(e, prompt, job); err != nil {
			return err
		}
	}
	return nil
}

func HandlePromptsList(c echo.Context) error {
	var err error
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandlePromptsList called")
	}
	prompt := types.NewPrompt(nil)
	prompts, err := prompt.List(c)
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
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandlePromptsGet called")
	}
	if id := c.Param("id"); id != "" {
		prompt := types.NewPrompt(&id)
		if err = prompt.Get(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		content := types.NewPromptTypeContent()
		if err = content.FromType(prompt); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		prompt, err = prompt.GetDispositions(c)
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
	if cc, ok := (c).(logger.LoggingContext); ok {
		cc.Flogger("HandlePromptDelete called")
	}
	if id := c.Param("id"); id != "" {
		entity := types.NewPrompt(&id)
		if err := entity.Delete(c); err != nil {
			return c.Render(http.StatusInternalServerError, "error.tpl", err.Error())
		}
		return HandlePromptsList(c)
	}
	return c.Render(http.StatusBadRequest, "error.tpl", "bad request: missing id")
}

func debug(msg string) {
	fmt.Printf(msg)
}
