package types

import "encoding/json"

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

