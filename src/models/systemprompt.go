package models

type SystemPrompt struct {
	Model
	ID string
	Name string
	Domain string
	Prompt string
}