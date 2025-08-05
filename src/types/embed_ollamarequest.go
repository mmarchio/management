package types

type OllamaRequest struct {
	Model 		string `json:"model"`
	System 		string `json:"system"`
	Prompt 		string `json:"prompt"`
	Stream 		bool `json:"stream"`
	Format 		string `json:"format"`
	KeepAlive 	string `json:"keep_alive"`
}

