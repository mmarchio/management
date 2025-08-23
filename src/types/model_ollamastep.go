package types

type OllamaStep struct {
	Model
	StepName string `json:"step_name"`
	OllamaModel string `json:"ollama_model"`
	SytemPrompt string `json:"system_prompt"`
	ContextPrompt string `json:"context_prompt"`
	PromptTemplate string `json:"prompt_template"`
	Uri string `json:"uri"`
	Port int32 `json:"port"`
}

func (c OllamaStep) GetType() string {
	return "ollama"
}

