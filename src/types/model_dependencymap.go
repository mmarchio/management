package types

import (
	"encoding/json"

	merrors "github.com/mmarchio/management/errors"
)

type DependencyMap struct {
	Dependencies []Dependency
	ComfyScript *ComfyScript
	ComfyPrompt *ComfyPrompt
}

type Dependency struct {
	Key string `json:"key"`
	Type string `json:"type"`
	Source string `json:"source"`
}

func (c *DependencyMap) Unmarshal(vars string) error {
	if err := json.Unmarshal([]byte(vars), c); err != nil {
		return merrors.JSONUnmarshallingError{CalledBy: "types.DependencyMap.Unmarshal"}.Wrap(err).Log()
	}
	return nil
}

func (c DependencyMap) Hydrate(cn *ComfyNode, jobrun *JobRun, step *Step) error {
	for _, dep := range c.Dependencies {
		switch dep.Type {
		case "ComfyScript":
			cacheVal := jobrun.GetValueCache(dep.Source)
			if cv, ok := cacheVal.(string); ok {
				if err := json.Unmarshal([]byte(cv), c.ComfyScript); err != nil {
					return merrors.JSONUnmarshallingError{Info: cv, CalledBy: "types.DependencyMap.Hydrate"}.Wrap(err).Log()
				}
			}
		case "ComfyPrompt":
			cacheVal := jobrun.GetValueCache(dep.Source)
			if cv, ok := cacheVal.(string); ok {
				if err := json.Unmarshal([]byte(cv), c.ComfyPrompt); err != nil {
					return merrors.JSONUnmarshallingError{Info: cv, CalledBy: "types.DependencyMap.Hydrate"}.Wrap(err).Log()
				}
			}
		default:
		}
	}
	return nil
}