package handlers

import (

	"github.com/labstack/echo/v4"
	"github.com/mmarchio/management/types"
)

type DisplayNode struct {
	types.Node
	Menu
	SystemPrompts []types.SystemPrompt
	Prompt types.Prompt
	Disposition types.Disposition
	PromptTemplates []types.PromptTemplate
	DisplayType string
	List []types.Node
}

func (c *DisplayNode) GetSystemPrompts(e echo.Context) error {
	entity := types.NewSystemPrompt(nil)
	list, err := entity.List(e)
	if err != nil {
		return err
	}
	c.SystemPrompts = list
	return nil
}

func (c *DisplayNode) GetPromptTemplates(e echo.Context) error {
	entity := types.NewPromptTemplate(nil)
	list, err := entity.List(e)
	if err != nil {
		return err
	}
	c.PromptTemplates = list
	return nil
}

func (c *DisplayNode) New(node types.Node) {
	c.Node = node
}

type DisplayJob struct {
	types.Job
	Menu
	Workflows []types.Workflow
	DisplayType string
	List []types.Job
}

type DisplayPrompt struct {
	types.Prompt
	Menu
	Workflows []types.Workflow
	DisplayType string
	List []types.Prompt
	Debug interface{}
}

type DisplayComfyUITemplate struct {
	types.ComfyUITemplate
	Menu
	DisplayType string
	List []types.ComfyUITemplate
}

type DisplayDisposition struct {
	types.Disposition
	Menu
	DisplayType string
	List []types.Disposition
}

type DisplayJobRun struct {
	types.JobRun
	Menu
	DisplayType string
	List []types.JobRun
}

type DisplaySystemPrompt struct {
	types.SystemPrompt
	Menu
	DisplayType string
	List []types.SystemPrompt
}

type DisplayWorkflow struct {
	types.Workflow
	Menu
	DisplayType string
	List []types.Workflow
	ComfyNodes []types.ComfyNode
	OllamaNodes []types.OllamaNode
	SSHNodes []types.SSHNode
}

type DisplayPromptTemplate struct {
	types.PromptTemplate
	Menu
	DisplayType string
	List []types.PromptTemplate
	Context string
}

type DisplayOllamaNode struct {
	types.OllamaNode
	Menu
	DisplayType string
	List []types.OllamaNode
	Enabled types.Toggle
	Bypass types.Toggle
	SystemPrompts []types.SystemPrompt
	PromptTemplates []types.PromptTemplate
}

type Menu struct {
	Href string
	Title string
}

type DisplayComfyNode struct {
	types.ComfyNode
	Menu
	DisplayType string
	List []types.ComfyNode
	Enabled types.Toggle
	Bypass types.Toggle
}

type DisplaySSHNode struct {
	types.SSHNode
	Menu
	DisplayType string
	List []types.SSHNode
	Enabled types.Toggle
	Bypass types.Toggle
}