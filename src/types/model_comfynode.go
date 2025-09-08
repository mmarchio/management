package types

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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

const COMFYRESTBASE string = "http://172.17.0.1"
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
	ServicePort		int32 			`form:"service_port" json:"service_port"`
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
	c.Dependencies = DependencyMap{}
	if err := c.Dependencies.Unmarshal(c.TemplateValues); err != nil {
		return merrors.JSONUnmarshallingError{}.Wrap(err).Log()
	}

	if step.Stats.Output.ComfyResponseData == nil {
		step.Stats.Output.ComfyResponseData = make([]ComfyResponseData, 0)
	}
	for _, dep := range c.Dependencies.Dependencies {
		if dep.Type == "ComfyPrompt" {
			if err := c.Dependencies.ComfyPrompt.Unmarshal(dep.Source, jobrun); err != nil {
				return err
			}
			_, err = c.Call(e, jobrun, step, clientID, dep.Key, c.Dependencies.ComfyPrompt)
			if err != nil {
				return merrors.NodeExecError{Info: dep.Type, CalledBy: "types.ComfyNode.Exec"}.Wrap(err).Log()
			}
		}
		if dep.Type == "ComfyScriptSegment" {
			cs := ComfyScript{}
			msi := jobrun.GetValueCache(dep.Source)
			if pgr, ok := msi.(map[string]interface{}); ok {
				cs.Segments = make([]ComfyScriptSegment, 0)
				// if output, ok := segments.(map[string]interface{}); ok {
				if segmentsSlice, ok := pgr["segments"].([]interface{}); ok {
					for _, segmentSlice := range segmentsSlice {
						css := ComfyScriptSegment{}
						if segmentMSI, ok := segmentSlice.(map[string]interface{}); ok {
							if text, ok := segmentMSI["text"].(string); ok {
								css.Text = text
							}
							if tm, ok := segmentMSI["time"].(int64); ok {
								css.Time = tm
							}
							cs.Segments = append(cs.Segments, css)
						}
					}
				}
				// }
				if len(cs.Segments) > 0 {
					if err := handleSegments(e, c, cs.Segments, jobrun, step, dep, clientID); err != nil {
						return err
					}
				}
				// GetLogger(3).Flogger("segments: %#v", segments)
				// if len(cs.Segments) > 0 {
				// 	GetLogger(3).Flogger("processing promptgenerationresponse")
				// 	if err := handleSegments(e, c, cs.Segments, jobrun, step, dep, clientID); err != nil {
				// 		return err
				// 	}
				// }
			} else {
				ncs, err := cs.Unmarshal(dep.Source, jobrun)
				if err != nil {
					return merrors.ContentGetError{}.Wrap(err).Log()
				}
				if &ncs == nil {
					GetLogger(3).Flogger("comfyscript is nil for step: %s: %#v", step.Name, cs)
					continue
				}
				c.Dependencies.ComfyScript = &ncs
				if ncs.Segments != nil {
					if len(ncs.Segments) > 0 {
						if err := handleSegments(e, c, ncs.Segments, jobrun, step, dep, clientID); err != nil {
							return err
						}
					} else {
						GetLogger(3).Flogger("segments len: %d", len(ncs.Segments))
					}
				} else {
					GetLogger(3).Flogger("segments nil")
				} 
			}
		}
		if dep.Type == "ComfyScript" {
			cs := ComfyScript{}
			ncs, err := cs.Unmarshal(dep.Source, jobrun)
			if err != nil {
				return merrors.ContentGetError{}.Wrap(err).Log()
			}
			c.Dependencies.ComfyScript = &ncs
			if &ncs == nil {
				GetLogger(3).Flogger("comfyscript is nil for step: %s: %#v", step.Name, cs)
				continue
			}
			if ncs.Segments != nil {
				if len(ncs.Segments) > 0 {
					if err := handleSegments(e, c, ncs.Segments, jobrun, step, dep, clientID); err != nil {
						return err
					}
				} else {
					GetLogger(3).Flogger("segments len: %d", len(ncs.Segments))
				}
			} else {
				GetLogger(3).Flogger("segments nil")
			} 
			if ncs.ComfyResponseData != nil && len(ncs.ComfyResponseData) > 0 {
				if err := handleComfyResponseData(e, c, ncs.ComfyResponseData, jobrun, step, dep, clientID); err != nil {
					return err
				}
			} else {
				GetLogger(3).Flogger("comfyresponsedata nil: %#v", ncs.ComfyResponseData)
			}
		}
	}

	return nil
}

func handleSegments(e echo.Context, c ComfyNode, segments []ComfyScriptSegment, jobrun *JobRun, step *Step, dep Dependency, clientID string) error {
	var err error
	for _, segment := range segments {
		_, err = c.Call(e, jobrun, step, clientID, dep.Key, segment)
		if err != nil {
			return merrors.NodeExecError{Info: dep.Type, CalledBy: "types.ComfyNode.Exec"}.Wrap(err).Log()
		}
	}
	return nil
}

func handleComfyResponseData(e echo.Context, c ComfyNode, cs []ComfyResponseData, jobrun *JobRun, step *Step, dep Dependency, clientID string) error {
	var err error
	for _, segment := range cs {
		filename := fmt.Sprintf("/ComfyUI/output/%s/%s", segment.Outputs.Audio.Subfolder, segment.Outputs.Audio.Filename)
		_, err = c.Call(e, jobrun, step, clientID, dep.Key, filename)
		if err != nil {
			return merrors.NodeExecError{Info: dep.Type, CalledBy: "types.ComfyNode.Exec"}.Wrap(err).Log()
		}
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
	jid := step.JobID.String()
	job := NewJob(&jid)
	if err := job.Get(e); err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err).Log()
	}
	if seed, ok := job.SeedMap[step.Node]; ok {
		if d, ok := data.(map[string]interface{}); ok {
			d["seed"] = seed
			data = d
		}

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

	crd, err := c.QueuePrompt(b)
	if err != nil {
		return nil, merrors.HTTPRequestError{CalledBy: "types.ComfyNode.Call"}.Wrap(err).Log()
	}
	if crd != nil {
		step.Stats.Output.ComfyResponseData = append(step.Stats.Output.ComfyResponseData, *crd)
		jobrun.ValueCache[fmt.Sprintf("steps[%d].output.comfy_response_data", step.Order-1)] = step.Stats.Output.ComfyResponseData
	} else {
		GetLogger(3).Flogger("crd is nil for step: %d", step.Order)
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
	comfyprompturi := fmt.Sprintf("%s:%d/oneapi/v1/execute", COMFYRESTBASE, c.ServicePort)
	cp := ComfyPrompt{}
	if err := json.Unmarshal(payload, &cp); err != nil {
		return nil, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
	}
	
	oneapi := OneAPI{
		Workflow: cp.Prompt,
		Timeout: 1200,
		WaitForResult: true,
	}
	pd, err := oneapi.Serialize()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer([]byte(pd))
	req, err := http.NewRequest("POST", comfyprompturi, buf)
	if err != nil {
		return nil, merrors.HTTPRequestError{Info: comfyprompturi, CalledBy: "types.ComfyNode.QueuePrompt"}.Wrap(err).Log()
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, merrors.HTTPRequestError{Info: comfyprompturi, CalledBy: "types.ComfyNode.QueuePrompt"}.Wrap(err).Log()
	}
	defer resp.Body.Close()
	crdm := ComfyResponseDataMap{}
	crd := ComfyResponseData{}
	bf := new(strings.Builder)
	_, err = io.Copy(bf, resp.Body)
	if err != nil {
		return nil, merrors.HTTPRequestError{Info: comfyprompturi}.Wrap(err).Log()
	}
	reader := strings.NewReader(bf.String())
	if err := json.NewDecoder(reader).Decode(&crdm); err != nil {
		return nil, merrors.JSONUnmarshallingError{Info: bf.String(), CalledBy: "types.ComfyNode.QueuePrompt"}.Wrap(err).Log()		
	}
	
	err = crd.Hydrate(crdm)
	if err != nil {
		return nil, merrors.JSONUnmarshallingError{CalledBy: "types.ComfyNode.ComfyResponseData.Hydrate"}.Wrap(err).Log()
	}
	crd.Raw = bf.String()
	return &crd, nil
}

func (c ComfyNode) ParseApiTemplate(step *Step, key string, templateValues interface{}) (string, error) {
	msi := make(map[string]interface{})
	if css, ok := templateValues.(ComfyScriptSegment); ok {
		templateValues = css.Text
	} else if strings.Contains(key, "dependency") {
		parts := strings.Split(key, ".")
		if len(parts) == 2 {
			switch parts[1] {
			case "prompt":
				templateValues = step.Stats.Output.Prompt
			case "topics":
				templateValues = step.Stats.Output.Topics
			case "output":
				templateValues = step.Stats.Output.Output
			case "think":
				templateValues = step.Stats.Output.Think
			default:
			}
		}
	}
	templateMSI := make(map[string]interface{})
	if err := json.Unmarshal([]byte(c.APITemplate), &templateMSI); err != nil {
		return "", merrors.JSONUnmarshallingError{Info: c.APITemplate}.Wrap(err).Log()
	}
	msi[key] = templateValues
	return strrep.Strrep(c.APITemplate, msi)	
} 

