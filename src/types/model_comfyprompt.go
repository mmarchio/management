package types

import (
	"encoding/json"

	merrors "github.com/mmarchio/management/errors"
)

type ComfyPrompt struct {
	Prompt string
}

func (c *ComfyPrompt) Unmarshal(in interface{}, jobrun *JobRun) error {
	data := jobrun.GetValueCache(in)
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
