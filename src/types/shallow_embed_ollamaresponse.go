package types

type ShallowOllamaResponse struct {
	Model 		string `json:"model"`
	CreatedAt 	string `json:"created_at"`
	Done 		bool `json:"done"`
	Response 	string `json:"response"`
}

func (c ShallowOllamaResponse) GetResponse() string {
	return c.Response
}