package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	_ "sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/strrep"
)

//var wg sync.WaitGroup

type OllamaNode struct {
	Model
	Name           string         			`form:"name" json:"name"`
	OllamaModel    string         			`form:"ollama_model" json:"model"`
	SystemPrompt   string         			`form:"system_prompt" json:"system_prompt"`
	Prompt         string         			`form:"prompt" json:"prompt"`
	PromptTemplate string         			`form:"prompt_template" json:"prompt_template"`
	ResponseModel  OllamaResponse 			`json:"response_model"`
	WorkflowID     WorkflowID     			`form:"workflow_id" json:"workflow_id"`
	Enabled        bool           			`json:"enabled"`
	Bypass         bool           			`json:"bypass"`
	Output         PromptGenerationResponse	`form:"output" json:"output"`
	ContextObject  *Stats
	Context        *Context
	Dependencies   []Dependency				`form:"dependencies" json:"dependencies"`
}

func (c OllamaNode) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowOllamaNode{}
	sm.ShallowModel = sm.ShallowModel.FromTypeModel(c.Model)
	sm.ID = c.ID
	sm.OllamaModel = c.OllamaModel
	sm.SystemPrompt = c.SystemPrompt
	sm.Prompt = c.Prompt
	sm.PromptTemplate = c.PromptTemplate
	sm.WorkflowID = c.WorkflowID
	sm.Enabled = c.Enabled
	sms = append(sms, sm)
	return sms
}

func (c OllamaNode) Validate() OllamaNode {
	msg := ""
	valid := true
	if !c.Model.Validate() {
		msg = "failed model validation"
		valid = false
	}
	if c.Model.ContentType != "ollamanode" {
		msg = "failed content type validation: " + c.Model.ContentType
		valid = false
	}
	if c.ID == "" || c.ID != c.Model.ID {
		msg = "failed id match"
		valid = false
	}
	if c.Name == "" || c.OllamaModel == "" {
		msg = fmt.Sprintf("failed name: %s ollama_model: %s", c.Name, c.OllamaModel)
		valid = false
	}
	if c.Prompt == "" && c.PromptTemplate == "" {
		msg = fmt.Sprintf("failed prompt: %s prompt_template: %s", c.Prompt, c.PromptTemplate)
		valid = false
	}
	c.Model.Validated = valid
	if !c.Model.Validated {
		fmt.Printf("validation failure: %s\n", msg)
	}
	return c
}

func (c OllamaNode) ValidateV2() bool {
	msg := ""
	valid := true
	if !c.Model.Validate() {
		msg = "failed model validation"
		valid = false
	}
	if c.Model.ContentType != "ollamanode" {
		msg = "failed content type validation: " + c.Model.ContentType
		valid = false
	}
	if c.ID == "" || c.ID != c.Model.ID {
		msg = "failed id match"
		valid = false
	}
	if c.Name == "" || c.OllamaModel == "" {
		msg = fmt.Sprintf("failed name: %s ollama_model: %s", c.Name, c.OllamaModel)
		valid = false
	}
	if c.Prompt == "" && c.PromptTemplate == "" {
		msg = fmt.Sprintf("failed prompt: %s prompt_template: %s", c.Prompt, c.PromptTemplate)
		valid = false
	}
	fmt.Printf("%s", msg)
	return valid
}

func (c OllamaNode) GetValidated() bool {
	return c.Model.Validated
}

func NewOllamaNode(id *string) OllamaNode {
	c := OllamaNode{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "ollamanode"
	if c.Model.CreatedAt.IsZero() {
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	} else {
		c.Model.UpdatedAt = time.Now()
	}
	return c
}

func (c OllamaNode) GetName() string {
	return c.Name
}

func (c OllamaNode) GetCommand() string {
	return ""
}

func (c OllamaNode) GetUser() string {
	return ""
}

func (c OllamaNode) GetHost() string {
	return ""
}

func (c OllamaNode) GetModel() string {
	return c.OllamaModel
}

func (c OllamaNode) GetSystemPrompt() string {
	return c.SystemPrompt
}

func (c OllamaNode) GetPrompt() string {
	return c.Prompt
}

func (c OllamaNode) GetPromptTemplate() string {
	return c.PromptTemplate
}

func (c OllamaNode) GetApiBase() string {
	return ""
}

func (c OllamaNode) GetApiTemplate() string {
	return ""
}

func (c OllamaNode) GetType() string {
	return "ollama_node"
}

func (c *OllamaNode) ParsePromptTemplate(e echo.Context) error {
	var err error
	if c.PromptTemplate != "" {
		GetLogger(3).Flogger("prompt template value: %s", c.PromptTemplate)
		if len(c.PromptTemplate) == 36 {
			id := c.PromptTemplate
			pt := NewPromptTemplate(&id)
			if err := pt.Get(e); err != nil {
				return merrors.ContentGetError{Package: "types", Struct: "OllamaNode", Function: "ParsePromptTemplate"}.Wrap(err).Log()
			}
			msi := make(map[string]interface{})
			if err := json.Unmarshal([]byte(pt.Vars), &msi); err != nil {
				return merrors.JSONUnmarshallingError{Info: pt.Vars, Package: "types", Struct: "OllamaNode", Function: "ParsePromptTemplate"}.Wrap(err).Log()
			}
			c.Prompt, err = strrep.Strrep(pt.Template, msi)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *OllamaNode) FromMSI(msi map[string]interface{}) error {
	var err error
	if p, ok := msi["params"].(map[string]interface{}); ok {
		if id, ok := p["id"].(string); ok {
			c.Model.ID = id
		}
		if createdAt, ok := p["CreatedAt"].(string); ok {
			c.Model.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
			if err != nil {
				return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct: "OllamaNode", Function: "FromMSI"}.Wrap(err).Log()
			}
		}
		if updatedAt, ok := p["UpdatedAt"].(string); ok {
			c.Model.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
			if err != nil {
				return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct: "OllamaNode", Function: "FromMSI"}.Wrap(err).Log()
			}
		}
		if ct, ok := p["ContentType"].(string); ok {
			c.Model.ContentType = ct
		}
		if id, ok := p["ID"].(string); ok {
			c.ID = id
		}
		if name, ok := p["name"].(string); ok {
			c.Name = name
		}
		if prompt, ok := p["prompt"].(string); ok {
			c.Prompt = prompt
		}
		if pt, ok := p["prompt_template"].(string); ok {
			c.PromptTemplate = pt
		}
		if sp, ok := p["system_prompt"].(string); ok {
			c.SystemPrompt = sp
		}
	}
	return nil
}

func (c *OllamaNode) Call(e echo.Context, jobrun *JobRun, step *Step, node OllamaNode) error {
	GetLogger(4).Flogger("step %d: Call called", step.Order)
	start := time.Now()
	if c.OllamaModel == "" {
		return merrors.NilContentError{}.New("ollamamodel is nil").Log()
	}
	// onode := c
	// GetLogger(3).Flogger("step %d: Getting OllamaNode instance", step.Order)
	// if err := onode.Get(e); err != nil {
	// 	return merrors.ContentGetError{}.Wrap(err).Log().Log()
	// }
	// GetLogger(3).Flogger("step %d: Executing preparePrompt", step.Order)
	// prompt, err := onode.preparePrompt(e, jobrun, step)
	// if err != nil {
	// 	GetLogger(1).Flogger("step %d: error preparing prompt: %s", step.Order, err.Error())
	// 	return err
	// }
	GetLogger(4).Flogger("step %d: onode retrieved", step.Order)
	oreq := OllamaRequest{
		Model:  node.OllamaModel,
		Prompt: node.Prompt,
		System: node.SystemPrompt,
		Format: "json",
		Stream:    false,
		KeepAlive: "5m",
		Think: false,
	}
	data, err := json.Marshal(oreq)
	if err != nil {
		return merrors.JSONMarshallingError{CalledBy: "types.OllamaNode.Call"}.Wrap(err).Log()
	}
	c.ResponseModel = OllamaResponse{}
	reader := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "http://172.17.0.1:7869/api/generate", reader)
	if err != nil {
		return merrors.HTTPRequestError{Package: "types", Struct: "OllamaNode", Function: "Call"}.Wrap(err).Log().Log()
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := worker(c, req, step)
	if err != nil {
		return err
	}
	pgr := PromptGenerationResponse{}
	pgr.Unmarshal(resp.Response, step)
	pgr.Think = ""
	GetLogger(3).Flogger("pgr: %#v", pgr)
	pgrb, err := json.Marshal(pgr)
	if err != nil {
		return merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	jobrun.ValueCache[fmt.Sprintf("steps[%d].output", step.Order -1)] = pgr
	jobrun.ValueCache[fmt.Sprintf("steps[%d].output.think", step.Order -1)] = pgr.Think
	jobrun.ValueCache[fmt.Sprintf("steps[%d].output.topics", step.Order -1)] = pgr.Topics
	jobrun.ValueCache[fmt.Sprintf("steps[%d].output.prompt", step.Order -1)] = pgr.Prompt
	jobrun.ValueCache[fmt.Sprintf("steps[%d].output.output", step.Order -1)] = pgr.Output
	jobrun.ValueCache[fmt.Sprintf("steps[%d].output.segments", step.Order -1)] = pgr.Segments

	c.Output = pgr
	c.ResponseModel.Done = resp.Done
	c.ResponseModel.Response = string(pgrb)
	end := time.Now()

	GetLogger(4).Flogger("step %d: start time: %s\n", step.Order, start.Format(time.RFC3339))
	GetLogger(4).Flogger("step %d: end time: %s\n", step.Order, end.Format(time.RFC3339))
	GetLogger(4).Flogger("step %d: time elapsed: %f\n", step.Order, time.Since(start).Seconds())
	return nil
}

func (c OllamaNode) PreparePrompt(e echo.Context, jobrun *JobRun, step *Step) (string, error) {
	prompt := c.Prompt
	if c.PromptTemplate != "" {
		pt := NewPromptTemplate(&c.PromptTemplate)
		if err := pt.Get(e); err != nil {
			return "", merrors.ContentGetError{}.Wrap(err).Log()
		}
		if err := pt.PrepareV2(e, jobrun, step); err != nil {
			return "", merrors.ContentGetError{}.Wrap(err).Log()
		}
		prompt = pt.Template
	}
	return prompt, nil
}

func worker(c *OllamaNode, req *http.Request, step *Step) (*OllamaResponse, error) {
	GetLogger(4).Flogger("step %d: worker called", step.Order)
	var err error
	cresp := OllamaResponse{}
	resp, err := http.DefaultClient.Do(req)
	GetLogger(4).Flogger("step %d: request made", step.Order)
	if err != nil {
		return nil, merrors.HTTPRequestError{}.Wrap(err).Log()
	}
	if resp != nil {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, merrors.HTTPRequestError{}.Wrap(err).Log()	
		}
		if err := json.Unmarshal(body, &cresp); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
		GetLogger(3).Flogger("ollama response status code: %d", resp.StatusCode)
		c.ResponseModel.Response = cresp.Response
		c.ResponseModel.Done = cresp.Done
	} else {
		GetLogger(1).Flogger("step %d: response is nil", step.Order)
	}
	return &cresp, nil
}

func (c *OllamaNode) Get(e echo.Context) error {
	content := NewOllamaNodeTypeContent(&c.Model.ID)
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	if err = json.Unmarshal([]byte(content.Content), c); err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Model.ID, Package: "types", Struct: "ollamanode", Function: "Get"}.Wrap(err).Log()
	}
	return nil
}

func NewOllamaNodeTypeContent(id *string) Content {
	c := Content{}
	if id != nil {
		c.Model.ID = *id
	}
	c.Model.ContentType = "ollamanode"
	return c
}

func (c OllamaNode) Delete(e echo.Context) error {
	content := NewOllamaNodeTypeContent(&c.Model.ID)
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	content.ID = c.ID
	onode := c
	if err := onode.Get(e); err != nil {
		return merrors.ContentGetError{}.Wrap(err).Log()
	}
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID, Package: "types", Struct: "ollamanode", Function: "delete"}.Wrap(err).Log()
	}
	wf := NewWorkflow(nil)
	wf.Model.ID = onode.WorkflowID.String()
	if err := wf.Get(e); err != nil {
		return merrors.ContentGetError{}.Wrap(err).Log()
	}
	wf.CutNode(c.Model.ID)
	wf.CutNodeOrder(c.Model.ID)
	if err := wf.Set(e, false); err != nil {
		return merrors.ContentSetError{}.Wrap(err).Log()
	}
	return nil
}

func (c OllamaNode) GetContentType() string {
	return c.Model.ContentType
}

func (c OllamaNode) GetID() string {
	return c.Model.ID
}

func (c OllamaNode) Set(e echo.Context, update bool) error {
	if !c.ValidateV2() {
		return merrors.ContentValidationError{Info: fmt.Sprintf("ollamanode: %#v", c)}.New("validation failed").Log()
	}
	c.Context = nil
	content := NewOllamaNodeTypeContent(&c.Model.ID)
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	err := content.Set(e, update)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	return nil
}

func (c *OllamaNode) GetNodeFromWorkflow(id string, wf Workflow) {
	for _, v := range wf.OllamaNodesArrayModel {
		if id == v.Model.ID {
			c = &v
			break
		}
	}
}

func (c OllamaNode) Exec(e echo.Context, jobrun *JobRun, step *Step) error {
	if step.Bypass.Value {
		return nil
	}
	pt := NewPromptTemplate(&c.PromptTemplate)
	if err := pt.Get(e); err != nil {
		return merrors.ContentGetError{}.Wrap(err).Log()
	}
	dep := DependencyMap{}
	dep.Unmarshal(pt.Values)
	GetLogger(4).Flogger("step %d: Exec called", step.Order)
	start := time.Now()
	if c.SystemPrompt != "" && len(c.SystemPrompt) == 36 {
		spid := c.SystemPrompt
		sp := NewSystemPrompt(&spid)
		if err := sp.Get(e); err != nil {
			return err
		}
		c.SystemPrompt = sp.Prompt
	}
	GetLogger(3).Flogger("step %d: Executing Call", step.Order)
	if err := c.Call(e, jobrun, step, c); err != nil {
		return err
	}

	end := time.Now()
	step.Stats.Start = start
	step.Stats.End = end
	step.Stats.Status = "done"
	step.Stats.Output = c.Output
	if err := step.Set(e, true); err != nil {
		return merrors.ContentSetError{}.Wrap(err).Log()
	}
	if err := c.Set(e, true); err != nil {
		return merrors.ContentSetError{}.Wrap(err).Log()
	}
	return nil
}

func (c OllamaNode) List(e echo.Context) ([]OllamaNode, error) {
	GetLogger(3).Flogger("OllamaNode List called")
	content := NewOllamaNodeTypeContent(nil)
	content.Model.ContentType = "ollamanode"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]OllamaNode, 0)
	for _, model := range contents {
		cut := OllamaNode{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "OllamaNode", Function: "List"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c OllamaNode) ListBy(e echo.Context, key string, value interface{}) ([]OllamaNode, error) {
	content := NewOllamaNodeTypeContent(nil)
	content.Model.ContentType = "ollamanode"
	list, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{Info: fmt.Sprintf("{\"%s\":\"%s\"}", key, value), Package: "types", Struct: "OllamaNode", Function: "ListBy"}.Wrap(err).Log()
	}
	cuts := make([]OllamaNode, 0)
	for _, model := range list {
		cut := OllamaNode{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "OllamaNode", Function: "ListBy"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}



