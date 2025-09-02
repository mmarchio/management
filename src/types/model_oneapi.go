package types

import (
	"encoding/json"

	merrors "github.com/mmarchio/management/errors"
)

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

