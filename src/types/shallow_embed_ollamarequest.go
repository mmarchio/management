package types

import "context"

type ShallowOllamaRequest struct {
	Model 		string `json:"model"`
	System 		string `json:"system"`
	Prompt 		string `json:"prompt"`
	Stream 		bool `json:"stream"`
	Format 		string `json:"format"`
	KeepAlive 	string `json:"keep_alive"`
}

func (c ShallowOllamaRequest) Expand(ctx context.Context) (*OllamaRequest, error) {
	r := OllamaRequest{}
	r.Model = c.Model
	r.System = c.System
	r.Prompt = c.Prompt
	r.Stream = c.Stream
	r.Format = c.Format
	r.KeepAlive = c.KeepAlive
	return &r, nil
}

func (c ShallowOllamaRequest) IsShallowModel() bool {
	return true
}
