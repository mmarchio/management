package types

import (
	"encoding/json"
	"fmt"
	"strings"

	merrors "github.com/mmarchio/management/errors"
)

type PromptGenerationResponse struct {
	Name string 							`json:"name"`
	Think string 							`json:"think"`
	Topics []string 						`json:"topics"`
	Prompt string 							`json:"prompt"`
	Output interface{} 						`json:"output"`
	ComfyResponse ComfyResponse 			`json:"comfy_response"`
	ComfyResponseData []ComfyResponseData 	`json:"comfy_response_data"`
	Responses []string 						`json:"responses"`
	Segments []ComfyScriptSegment			`json:"segments"`			
}

func (c PromptGenerationResponse) IsNil() bool {
	if c.Think != "" {
		return false
	}
	if len(c.Topics) > 0 {
		return false
	}
	if c.Prompt != "" {
		return false
	}
	if c.Output != "" {
		return false
	}
	return true
}

func (c *PromptGenerationResponse) Unmarshal(s string, step *Step) error {
	GetLogger(3).Flogger("step %d: input: %s", step.Order, s)
	s = strings.ReplaceAll(s, "'", "&#39")
	s = strings.Replace(s, "{", "", 1)
	s = strings.Replace(s, "<think>", "{\"think\":\"", 1)
	s = strings.Replace(s, "\u003cthink\u003e", "{\"think\":\"", 1)
	s = strings.Replace(s, "</think>", "\",", 1)
	s = strings.Replace(s, "\u003c/think\u003e", "\",", 1)
	// s = strings.ReplaceAll(s, "'", "\\'")
	s = fmt.Sprintf("{%s", s)
	GetLogger(3).Flogger("step %d: output: %s", step.Order, s)
	if err := json.Unmarshal([]byte(s), c); err != nil {
		return merrors.JSONUnmarshallingError{Info: s}.Wrap(err).Log()
	}
	c.Think = ""
	GetLogger(3).Flogger("promptgenerationresponse.unmarshal: %#v", c)
	return nil
}

func (c PromptGenerationResponse) Marshal() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	return string(b), nil
}

