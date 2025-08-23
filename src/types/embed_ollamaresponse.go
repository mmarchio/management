package types

type OllamaResponse struct {
	Model 		string `json:"model"`
	CreatedAt 	string `json:"created_at"`
	Done 		bool `json:"done"`
	Response 	string `json:"response"`
}

func (c OllamaResponse) GetResponse() string {
	return c.Response
}