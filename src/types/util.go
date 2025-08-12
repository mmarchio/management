package types

import (
	"encoding/json"
	"strings"

	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/logger"
)

func GetLogger() logger.LoggingContext {
	r := logger.LoggingContext{}
	r.Init()
	return r
}

func jpath(j, p string) (interface{}, error) {
	parts := strings.Split(p, ".")
	msi := make(map[string]interface{})
	if err := json.Unmarshal([]byte(j), &msi); err != nil {
		GetLogger().Flogger("jpath unmarshalling error: %s", err.Error())
		return "", merrors.JSONUnmarshallingError{}.Wrap(err)
	}
	return jpathRecurse(parts, msi)
}

func jpathRecurse(parts []string, msi map[string]interface{}) (interface{}, error) {
	GetLogger().Flogger("parts: %#v", parts)
	if len(parts) == 1 {
		if sub, ok := msi[parts[0]].(string); ok {
			GetLogger().Flogger("jpath result: %T:%v", sub, sub)
			return sub, nil
		}
	}
	if len(parts) > 0 {
		if parts[0] == "$" {
			return jpathRecurse(parts[1:], msi)
		}
		if sub, ok := msi[parts[0]].(map[string]interface{}); ok {
			return jpathRecurse(parts[1:], sub)
		} else if sub, ok := msi[parts[0]].([]interface{}); ok {
			if len(sub) > 0 {
				for _, sl := range sub {
					if s, ok := sl.(map[string]interface{}); ok {
						return jpathRecurse(parts[1:], s)
					}
				}
			}
		} else {
			GetLogger().Flogger("unknown index %s type: %T", parts[0], msi[parts[0]])
		}
	}
	return nil, merrors.JPATHError{}.New("resource not found")
}