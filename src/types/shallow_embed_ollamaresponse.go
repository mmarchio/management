package types

import "context"

type ShallowOllamaResponse struct {
	ShallowModel
	Model 		string `json:"model"`
	CreatedAt 	string `json:"created_at"`
	Done 		bool `json:"done"`
	Response 	string `json:"response"`
}

func (c ShallowOllamaResponse) GetResponse() string {
	return c.Response
}

func (c ShallowOllamaResponse) Expand(ctx context.Context) (*OllamaResponse, error) {
	r := OllamaResponse{}
	r.Model = c.Model
	r.CreatedAt = c.CreatedAt
	r.Done = c.Done
	r.Response = c.Response
	return &r, nil
}

func (c ShallowOllamaResponse) IsShallowModel() bool {
	return true
}
