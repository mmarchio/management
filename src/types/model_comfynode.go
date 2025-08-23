package types

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	// "os"
	// "os/exec"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
	"github.com/mmarchio/management/strrep"
)

const COMFYRESTBASE string = "http://172.17.0.1:8188"
const COMFYWSBASE string = "ws://172.17.0.1:8188"

func NewComfyModelContent(idPtr *string) models.Content {
	var id string
	if idPtr == nil {
		id = uuid.NewString()
	} else {
		id = *idPtr
	}
	r := models.Content{
		ContentType: "comfynode",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	r.Model.ID = id
	r.Model.ContentType = "comfynode"
	return r
}

type datareq interface{
	Unmarshal(string) (datareq, error)
}

type ComfyNode struct {
	Model
	Name 			string 			`form:"name" json:"name"`
	Prompt          string          `form:"prompt" json:"prompt"`
	APIBase 		string 			`form:"api_base" json:"api_base"`
	APITemplate 	string 			`form:"api_template" json:"api_template"`
	TemplateValues  string			`json:"template_values"`
	WorkflowID  	WorkflowID 		`form:"workflow_id" json:"workflow_id"`
	Type 			string 			`form:"type" json:"type"`
	Enabled 		Toggle  		`json:"enabled"`
	Bypass 			Toggle			`json:"bypass"`
	Output 			string 			`form:"output" json:"output"`
	Dependencies	DependencyMap
}

func (c ComfyNode) Pack() []shallowmodel {
	sms := make([]shallowmodel, 0)
	sm := ShallowComfyNode{}
	sm.ShallowModel = sm.ShallowModel.FromTypeModel(c.Model)
	sm.ID = c.ID
	sm.Name = c.Name
	sm.Prompt = c.Prompt
	sm.APIBase = c.APIBase
	sm.APITemplate = c.APITemplate
	sm.TemplateValues = c.TemplateValues
	sm.WorkflowID = c.WorkflowID
	sm.Type = c.Type
	sm.Enabled = c.Enabled.Value
	sm.Bypass = c.Bypass.Value
	sm.Output = c.Output
	sms = append(sms, sm)
	return sms
}

func (c *ComfyNode) Validate() ComfyNode {
	valid := true
	if !c.Model.Validate() {
		GetLogger(3).Flogger("model validation failed")
		valid = false
	}
	if c.Model.ContentType != "comfynode" {
		GetLogger(3).Flogger("wrong content type: %s", c.Model.ContentType)
		valid = false
	}
	if c.Name == "" || c.APIBase == "" || c.APITemplate == "" {
		GetLogger(3).Flogger("nil content")
		valid = false
	}
	if c.WorkflowID == "" {
		GetLogger(3).Flogger("workflow is nil")
		valid = false
	}
	c.Model.Validated = valid
	GetLogger(3).Flogger("c.model.validated: %t", c.Model.Validated)
	return *c
}

func (c ComfyNode) GetValidated() bool {
	return c.Model.Validated
}

func (c ComfyNode) GetName() string {
	return c.Name
}

func (c ComfyNode) GetCommand() string {
	return ""
}

func (c ComfyNode) GetUser() string {
	return ""
}

func (c ComfyNode) GetHost() string {
	return ""
}

func (c ComfyNode) GetModel() string {
	return ""
}

func (c ComfyNode) GetSystemPrompt() string {
	return ""
}

func (c ComfyNode) GetPrompt() string {
	return ""
}

func (c ComfyNode) GetPromptTemplate() string {
	return ""
}

func (c ComfyNode) GetApiBase() string {
	return c.APIBase
}

func (c ComfyNode) GetApiTemplate() string {
	return c.APITemplate
}
func (c ComfyNode) GetType() string {
	return "comfy_node"
}

func (c *ComfyNode) FromMSI(msi map[string]interface{}) error {
	var err error
	if id, ok := msi["id"].(string); ok {
		c.Model.ID = id
	}
	if createdAt, ok := msi["CreatedAt"].(string); ok {
		c.Model.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "createdAt", Package: "types", Struct:"ComfyNode", Function: "FromMSI"}.Wrap(err).Log()
		}
	}
	if updatedAt, ok := msi["UpdatedAt"].(string); ok {
		c.Model.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			return merrors.MSIConversionError{Info: "updatedAt", Package: "types", Struct:"ComfyNode", Function: "FromMSI"}.Wrap(err).Log()
		}
	}
	if ct, ok := msi["ContentType"].(string); ok {
		c.Model.ContentType = ct
	}
	if name, ok := msi["name"].(string); ok {
		c.Name = name
	}
	if ab, ok := msi["api_base"].(string); ok {
		c.APIBase = ab
	}
	if at, ok := msi["api_template"].(string); ok {
		c.APITemplate = at
	}
	return nil
}

func (c *ComfyNode) Get(e echo.Context) error {
	content := NewComfyNodeTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "comfynode"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	if err := json.Unmarshal([]byte(content.Content), c); err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Model.ID}.Wrap(err).Log()
	}
	return nil
}

func (c *ComfyNode) GetShallow(e echo.Context) error {
	content := NewComfyNodeTypeContent()
	content.Model.ID = c.Model.ID
	content.Model.ContentType = "comfynode"
	content, err := content.Get(e)
	if err != nil {
		return merrors.ContentGetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "node", Function: "Get"}.Wrap(err).Log()
	}
	return nil
}

func NewComfyNodeTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "comfynode"
	return c
}

func (c ComfyNode) Delete(e echo.Context) error {
	content := NewComfyNodeTypeContent()
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	content.ID = c.ID
	if err := c.Get(e); err != nil {
		return merrors.ContentGetError{CalledBy: "types.ComfyNode.Delete"}.Wrap(err).Log()
	}
	wf := NewWorkflow(nil)
	wf.Model.ID = c.WorkflowID.String()
	if err := wf.Get(e); err != nil {
		return merrors.ContentGetError{}.Wrap(err).Log()
	}
	wf.CutNode(c.Model.ID)
	wf.CutNodeOrder(c.Model.ID)
	if err := wf.Set(e, false); err != nil {
		return merrors.ContentSetError{}.Wrap(err).Log()
	}
	if err := content.Delete(e); err != nil {
		return merrors.ContentDeleteError{Info: c.Model.ID, CalledBy: "types.ComfyNode.Delete"}.Wrap(err).Log()
	}
	return nil
}

func (c ComfyNode) GetContentType() string {
	return c.Model.ContentType
}

func (c ComfyNode) GetID() string {
	return c.Model.ID
}

func NewComfyNode(id *string) ComfyNode {
	c := ComfyNode{}
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.ID = c.Model.ID
	c.Model.ContentType = "comfynode"
	if c.Model.CreatedAt.IsZero() {
		c.Model.CreatedAt = time.Now()
		c.Model.UpdatedAt = c.Model.CreatedAt
	} else {
		c.Model.UpdatedAt = time.Now()
	}
	return c
}

func (c ComfyNode) Set(e echo.Context, update bool) error {
	d := c.Validate()
	if !d.Model.Validated {
		return merrors.ContentValidationError{CalledBy: "ComfyNode.Set"}.New("validation failed").Log()
	}
	content := NewComfyNodeTypeContent()
	content.FromType(c, c.Model)
	content.Model.ID = c.Model.ID
	content.ID = c.Model.ID
	err := content.Set(e, update)
	if err != nil {
		return merrors.ContentSetError{Info: c.Model.ID}.Wrap(err).Log()
	}
	return nil
}

func (c ComfyNode) List(e echo.Context) ([]ComfyNode, error) {
	content := NewComfyModelContent(nil)
	content.Model.ContentType = "comfynode"
	contents, err := content.List(e)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.Model.ContentType}.Wrap(err).Log()
	}
	cuts := make([]ComfyNode, 0)
	for _, model := range contents {
		cut := ComfyNode{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ComfyNode", Function: "List"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ComfyNode) ListBy(e echo.Context, key string, value interface{}) ([]ComfyNode, error) {
	content := NewComfyModelContent(nil)
	content.Model.ContentType = "comfynode"
	list, err := content.ListBy(e, key, value)
	if err != nil {
		return nil, merrors.ContentListByError{Info: fmt.Sprintf("{\"%s\":\"%s\"}", key, value), Package: "types", Struct: "ComfyNode", Function: "ListBy"}.Wrap(err).Log()
	}
	cuts := make([]ComfyNode, 0)
	for _, model := range list {
		cut := ComfyNode{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ComfyNode", Function: "ListBy"}.Wrap(err).Log()
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ComfyNode) Exec(e echo.Context, jobrun *JobRun, step *Step) error {
	var err error
	clientID := uuid.NewString()
	GetLogger(3).Flogger("clientID: %s", clientID)
	c.Dependencies = DependencyMap{}
	c.Dependencies.Unmarshal(c.TemplateValues)
	var crd *ComfyResponseData
	for _, dep := range c.Dependencies.Dependencies {
		if dep.Type == "ComfyPrompt" {
			if err := c.Dependencies.ComfyPrompt.Unmarshal(dep.Source, jobrun); err != nil {
				return err
			}
			crd, err = c.Call(e, jobrun, step, clientID, dep.Key, c.Dependencies.ComfyPrompt)
			if err != nil {
				return merrors.NodeExecError{Info: dep.Type, CalledBy: "types.ComfyNode.Exec"}.Wrap(err).Log()
			}
		}
		if dep.Type == "ComfyScript" {
			GetLogger(3).Flogger("dependencies: %#v", c.Dependencies)
			GetLogger(3).Flogger("comfyscript: %#v", c.Dependencies.ComfyScript)
			cs, err := c.Dependencies.ComfyScript.Unmarshal(dep.Source, jobrun)
			if err != nil {
				return err
			}
			if c.Dependencies.ComfyScript != nil {
				for _, segment := range c.Dependencies.ComfyScript.Segments {
					// prompt, err := c.ParseApiTemplate(step, dep.Key, data)
					// if err != nil {
					// 	return err
					// }
					// GetLogger(3).Flogger("prompt: %s", prompt)
					crd, err = c.Call(e, jobrun, step, clientID, dep.Key, segment)
					if err != nil {
						return merrors.NodeExecError{Info: dep.Type, CalledBy: "types.ComfyNode.Exec"}.Wrap(err).Log()
					}
				}
			} else {
				for _, segment := range cs.Segments {
					crd, err = c.Call(e, jobrun, step, clientID, dep.Key, segment)
					if err != nil {
						return merrors.NodeExecError{Info: dep.Type, CalledBy: "types.ComfyNode.Exec"}.Wrap(err).Log()
					}
				}
				GetLogger(3).Flogger("comfyscript is nil for step: %s", step.Name)
			}
		}
	}
	crds := make([]ComfyResponseData, 0)
	if crd != nil {
		crds = append(crds, *crd)
		jobrun.ValueCache[fmt.Sprintf("steps[%d].output.prompt", step.Order-1)] = ""
		jobrun.ValueCache[fmt.Sprintf("steps[%d].output.topics", step.Order-1)] = nil
		jobrun.ValueCache[fmt.Sprintf("steps[%d].output.output", step.Order-1)] = crds
		step.Stats.Output.Output = crds
	}

	return nil
}

func (c ComfyNode) Call(e echo.Context, jobrun *JobRun, step *Step, clientID string, key string, data interface{}) (*ComfyResponseData, error) {
	node := NewComfyNode(&step.Node)
	if err := node.Get(e); err != nil {
		return nil, merrors.ContentGetError{CalledBy: "ComfyNode.Call"}.Wrap(err).Log()
	}
	if data == nil {
		return nil, merrors.NilContentError{}.New("data supplied for interpolation is nil").Log()
	}

	prompt, err := c.ParseApiTemplate(step, key, data)
	if err != nil {
		return nil, err
	}

	cp := ComfyMessage{
		Prompt: prompt,
		ClientID: clientID,
	}

	b, err := json.Marshal(cp)
	if err != nil {
		return nil, merrors.JSONMarshallingError{CalledBy: "types.ComfyNode.Call"}.Wrap(err).Log()
	}
	GetLogger(3).Flogger("payload: %s", string(b))
	crd, err := c.QueuePrompt(b)
	if err != nil {
		return nil, merrors.HTTPRequestError{CalledBy: "types.ComfyNode.Call"}.Wrap(err).Log()
	}
	return crd, nil
}

func (c ComfyNode) ApiConn(ctx context.Context, uri string) (*websocket.Conn, *http.Response, error) {
	dialer := websocket.DefaultDialer
	dialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn, resp, err := dialer.Dial(uri, http.Header{})
	if err != nil {
		return nil, nil, merrors.WebsocketDialError{CalledBy: "types.ComfyNode.ApiConn"}.Wrap(err).Log()
	}
	return conn, resp, err
}

func (c ComfyNode) QueuePrompt(payload []byte) (*ComfyResponseData, error) {
	comfyprompturi := fmt.Sprintf("%s/oneapi/v1/execute", COMFYRESTBASE)
	cp := ComfyPrompt{}
	if err := json.Unmarshal(payload, &cp); err != nil {
		return nil, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
	}
	
	oneapi := OneAPI{
		Workflow: cp.Prompt,
		Timeout: 300,
		WaitForResult: true,
	}
	pd, err := oneapi.Serialize()
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer([]byte(pd))
	req, err := http.NewRequest("POST", comfyprompturi, buf)
	if err != nil {
		return nil, merrors.HTTPRequestError{CalledBy: "types.ComfyNode.QueuePrompt"}.Wrap(err).Log()
	}
	req.Header.Add("content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, merrors.HTTPRequestError{CalledBy: "types.ComfyNode.QueuePrompt"}.Wrap(err).Log()
	}
	crd := ComfyResponseData{}
	if err := json.NewDecoder(resp.Body).Decode(&crd); err != nil {
		return nil, merrors.JSONUnmarshallingError{CalledBy: "types.ComfyNode.QueuePrompt"}.Wrap(err).Log()		
	}
	return &crd, nil
}

func (c ComfyNode) ParseApiTemplate(step *Step, key string, templateValues interface{}) (string, error) {
	msi := make(map[string]interface{})
	if css, ok := templateValues.(ComfyScriptSegment); ok {
		templateValues = css.Text
	}
	msi[key] = templateValues
	GetLogger(3).Flogger("strrep data: %#v", msi)
	return strrep.Strrep(c.APITemplate, msi)	
} 

type OneAPI struct {
	Workflow string `json:"workflow"`
	Params struct{
		Prompt string `json:"prompt"`
	} `json:"params"`
	WaitForResult bool `json:"wait_for_result"`
	Timeout int64 `json:"timeout"`
}

func (c OneAPI) ToMSI() (map[string]interface{}, error) {
	msi := make(map[string]interface{})
	b, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &msi); err != nil {
		return nil, err
	}
	return msi, nil
}

func (c OneAPI) Serialize() (string, error) {
	msi, err := c.ToMSI()
	if err != nil {
		return "", err
	}

	workflow := make(map[string]interface{})
	if err := json.Unmarshal([]byte(c.Workflow), &workflow); err != nil {
		return "", merrors.JSONUnmarshallingError{Info: c.Workflow}.Wrap(err).Log()
	}
	if p, ok := workflow["prompt"].(string); ok {
		prompt := make(map[string]interface{})
		if err := json.Unmarshal([]byte(p), &prompt); err != nil {
			return "", merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
		workflow["prompt"] = prompt
	}
	msi["workflow"] = workflow
	b, err := json.Marshal(msi)
	if err != nil {
		return "", merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	return string(b), err
}

type ComfyMessage struct {
	Prompt string `json:"prompt"`
	ClientID string `json:"client_id"`
}

func (c ComfyMessage) Serialize() ([]byte, error) {
	msi := make(map[string]interface{})
	p := make(map[string]interface{})
	msi["prompt"] = p
	msi["client_id"] = c.ClientID
	pd, err := json.Marshal(msi)
	if err != nil {
		return nil, err
	}
	return pd, nil
}

type ComfyPrompt struct {
	Prompt string
}

func (c *ComfyPrompt) Unmarshal(in interface{}, jobrun *JobRun) error {
	data := jobrun.GetValueCache(in)
	GetLogger(3).Flogger("in: %s, data: %#v", in, data)
	if s, ok := data.(string); ok {
		if err := json.Unmarshal([]byte(s), c); err != nil {
			return merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
	}
	if b, ok := data.([]byte); ok {
		if err := json.Unmarshal(b, c); err != nil {
			return merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
	}
	if m, ok := data.(map[string]interface{}); ok {
		b, err := json.Marshal(m)
		if err != nil {
			return merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
		if b == nil {
			GetLogger(3).Flogger("nil bytes: source: %s, data: %#v, map: %#v", in, data, m)
		}
		t := ComfyPrompt{}
		if err := json.Unmarshal(b, &t); err != nil {
			return merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
		c = &t
	}
	return nil
}

type ComfyResonse struct {
	Type string `json:"type"`
	Data ComfyResponseData `json:"data"`
}
type ComfyResponseData struct {
	Images []string `json:"images"`
	ImagesByVar map[string]interface{} `json:"images_by_var"`
	Outputs map[string]interface{} `json:"outputs"`
	NodeErrors interface{} `json:"node_errors"`
	PromptID string `json:"prompt"`
	Status string `json:"status"`
}

type ComfyScriptSegment struct {
	Text string `json:"text"`
	Time int64 `json:"time"`
}

type ComfyScript struct {
	Segments []ComfyScriptSegment `json:"segments"`
}

func (c *ComfyScript) Unmarshal(in interface{}, jobrun *JobRun) (*ComfyScript, error) {
	data := jobrun.GetValueCache(in)
	GetLogger(3).Flogger("data: %#v", data)
	if s, ok := data.(string); ok {
		if err := json.Unmarshal([]byte(s), c); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
	}
	if b, ok := data.([]byte); ok {
		if err := json.Unmarshal(b, c); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
	}
	if m, ok := data.(map[string]interface{}); ok {
		b, err := json.Marshal(m)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
		t := ComfyScript{}
		if err := json.Unmarshal(b, &t); err != nil {
			return nil, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
		c = &t
		GetLogger(3).Flogger("comfyscript unmarshaled: %#v", c)
		return &t, nil
	}
	return nil, nil
}

type DependencyMap struct {
	Dependencies []Dependency
	ComfyScript *ComfyScript
	ComfyPrompt *ComfyPrompt
}

type Dependency struct {
	Key string `json:"key"`
	Type string `json:"type"`
	Source string `json:"source"`
}

func (c *DependencyMap) Unmarshal(vars string) error {
	if err := json.Unmarshal([]byte(vars), c); err != nil {
		return merrors.JSONUnmarshallingError{CalledBy: "types.DependencyMap.Unmarshal"}.Wrap(err).Log()
	}
	return nil
}

func (c DependencyMap) Hydrate(cn *ComfyNode, jobrun *JobRun, step *Step) error {
	for _, dep := range c.Dependencies {
		switch dep.Type {
		case "ComfyScript":
			cacheVal := jobrun.GetValueCache(dep.Source)
			if cv, ok := cacheVal.(string); ok {
				if err := json.Unmarshal([]byte(cv), c.ComfyScript); err != nil {
					return merrors.JSONUnmarshallingError{Info: cv, CalledBy: "types.DependencyMap.Hydrate"}.Wrap(err).Log()
				}
			}
		case "ComfyPrompt":
			cacheVal := jobrun.GetValueCache(dep.Source)
			if cv, ok := cacheVal.(string); ok {
				if err := json.Unmarshal([]byte(cv), c.ComfyPrompt); err != nil {
					return merrors.JSONUnmarshallingError{Info: cv, CalledBy: "types.DependencyMap.Hydrate"}.Wrap(err).Log()
				}
			}
		default:
		}
	}
	return nil
}