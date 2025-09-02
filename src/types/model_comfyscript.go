package types

import (
	"encoding/json"

	merrors "github.com/mmarchio/management/errors"
)

type ComfyScriptSegment struct {
	Text string `json:"text"`
	Time int64 `json:"time"`
}

type ComfyScript struct {
	Segments []ComfyScriptSegment `json:"segments"`
	ComfyResponseData []ComfyResponseData `json:"comfy_response_data"`
}

func (c *ComfyScript) Unmarshal(in interface{}, jobrun *JobRun) (ComfyScript, error) {
	GetLogger(3).Flogger("comfyscript.unmarshal in: %#v", in)
	data := jobrun.GetValueCache(in)
	d := *c
	GetLogger(3).Flogger("data: %#v", data)
	if s, ok := data.([]ComfyResponseData); ok {
		c.ComfyResponseData = make([]ComfyResponseData, 0)
		if s != nil {
			c.ComfyResponseData = s
			return *c, nil
		}
	}
	if s, ok := data.(string); ok {
		if err := json.Unmarshal([]byte(s), c); err != nil {
			return d, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
	}
	if b, ok := data.([]byte); ok {
		if err := json.Unmarshal(b, c); err != nil {
			return d, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
	}
	if m, ok := data.(map[string]interface{}); ok {
		b, err := json.Marshal(m)
		if err != nil {
			return d, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
		t := ComfyScript{}
		if err := json.Unmarshal(b, &t); err != nil {
			return d, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
		}
		GetLogger(3).Flogger("comfyscript unmarshaled: %#v", c)
		return t, nil
	}
	return d, nil
}

