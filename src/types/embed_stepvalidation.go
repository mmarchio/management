package types

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	merrors "github.com/mmarchio/management/errors"
)

type StepValidation struct {
	EmbedModel
	Type string `form:"type" json:"type"`
	Field string `form:"field" json:"field"`
	Value string `form:"value" json:"value"`
}

func (c StepValidation) ValidateTemplate(s string) bool {
	pattern := regexp.QuoteMeta("{") + regexp.QuoteMeta("{") + "[a-zA-Z0-9]*" + regexp.QuoteMeta("(") + "." + regexp.QuoteMeta(")") + regexp.QuoteMeta("}") + regexp.QuoteMeta("}")
	r := regexp.MustCompile(pattern)
	match := r.Find([]byte(s))
	if match != nil {
		GetLogger(3).Flogger("pattern match found: %s", string(match))
		return false
	}
	return true	
}

func (c StepValidation) Validate(step *Step) (bool, error) {
	
	if err := json.Unmarshal([]byte(step.Validation), &c); err != nil {
		return false, merrors.JSONUnmarshallingError{}.Wrap(err).Log()
	}
	if c.Type == "PromptGenerationResponse" {
		target := step.Stats.Output
		switch c.Field {
		case "think":
			switch c.Value {
			case "empty":
				if target.Think == "" {
					return false, nil
				}
			}
		case "prompt":
			switch c.Value {
			case "emtpy":
				if target.Prompt == "" {
					return false, nil
				}
			}
		case "topics":
			parts := strings.Split(c.Value, ":")
			if len(parts) == 2 {
				if parts[0] == "len" {
					n, err := strconv.Atoi(parts[1])
					if err != nil {
						return false, err
					}
					if len(target.Topics) != n {
						return false, nil
					}
				}
			}
		case "output":
		case "ComfyResponse":
		case "ComfyResponseData":
		case "Responses":
		}
	}
	return true, nil
}