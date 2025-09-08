package types

import (
	"encoding/json"
	"strings"

	merrors "github.com/mmarchio/management/errors"
)

type ComfyResponse struct {
	Type 	string 					`json:"type"`
	DataMap ComfyResponseDataMap 	`json:"data"`
}

type ComfyResponseDataMap struct {
	Images []string `json:"images"`
	ImagesByVar map[string]interface{} `json:"images_by_var"`
	Outputs map[string]interface{} `json:"outputs"`
	NodeErrors interface{} `json:"node_errors"`
	PromptID string `json:"prompt"`
	Status string `json:"status"`
}

type ComfyResponseData struct {
	Images []string `json:"images"`
	ImagesByVar map[string]interface{} `json:"images_by_var"`
	Outputs ComfyResponseDataOutput `json:"outputs"`
	NodeErrors interface{} `json:"node_errors"`
	PromptID string `json:"prompt"`
	Status string `json:"status"`
	Videos []string `json:"videos"`
	Raw string `json:"raw"`
}

type ComfyResponseDataOutput struct {
	Audio ComfyResponseDataOutputAudio `json:"audio"`
}

type ComfyResponseDataOutputAudio struct {
	Filename string `json:"filename"`
	Subfolder string `json:"subfolder"`
	Type string `json:"type"`
}

func (c *ComfyResponseData) Hydrate(crdm ComfyResponseDataMap) error {
	c.Images = crdm.Images
	c.ImagesByVar = crdm.ImagesByVar
	c.NodeErrors = crdm.NodeErrors
	c.PromptID = crdm.PromptID
	c.Status = crdm.Status
	crdo := ComfyResponseDataOutput{}
	b, err := json.Marshal(crdm)
	if err != nil {
		return merrors.JSONMarshallingError{}.Wrap(err).Log()
	}
	bs := string(b)
	bs = strings.Replace(bs, "{\"5\":{\"audio\":[", "{\"audio\":", 1)
	bs = strings.Replace(bs, "}]}}", "}}", 1)
	bsmsi := make(map[string]interface{})
	if err := json.Unmarshal([]byte(bs), &bsmsi); err != nil {
		if err := jsonFix(bs, &bsmsi); err != nil {
			return merrors.JSONUnmarshallingError{Info: bs}.Wrap(err).Log()
		}
	}
	if outputs, ok := bsmsi["outputs"].(map[string]interface{}); ok {
		if audio, ok := outputs["audio"].(map[string]interface{}); ok {
			if filename, ok := audio["filename"].(string); ok {
				crdo.Audio.Filename = filename
			}
			if subfolder, ok := audio["subfolder"].(string); ok {
				crdo.Audio.Subfolder = subfolder
			}
			if t, ok := audio["type"].(string); ok {
				crdo.Audio.Type = t
			}
		}
	}
	c.Outputs = crdo
	return nil
}

func jsonFix(s string, dest interface{}) error {
	s = strings.ReplaceAll(s, "}},\"node_errors", "}]},\"node_errors")
	s += "}"
	if err := json.Unmarshal([]byte(s), dest); err != nil {
		return merrors.JSONMarshallingError{Info: s}.Wrap(err).Log()
	}
	return nil
}
