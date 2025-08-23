package types

import (
	"encoding/json"
	"strings"

	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/logger"
)

func GetLogger(severity int) logger.LoggingContext {
	r := logger.LoggingContext{Severity: severity}
	r.Init()
	return r
}

func jpath(j, p string) (interface{}, error) {
	parts := strings.Split(p, ".")
	msi := make(map[string]interface{})
	if err := json.Unmarshal([]byte(j), &msi); err != nil {
		return "", merrors.JSONUnmarshallingError{}.Wrap(err).Log()
	}
	return jpathRecurse(parts, msi)
}

func jpathRecurse(parts []string, msi map[string]interface{}) (interface{}, error) {
	GetLogger(4).Flogger("parts: %#v", parts)
	if len(parts) == 1 {
		if sub, ok := msi[parts[0]].(string); ok {
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
			GetLogger(3).Flogger("unknown index %s type: %T", parts[0], msi[parts[0]])
		}
	}
	return nil, merrors.JPATHError{}.New("resource not found")
}

func CheckType(e echo.Context, id, ctype *string) (bool, error) {
	m := Content{}
	if id == nil {
		return false, merrors.ContentCheckError{CalledBy: "types.CheckType"}.New("id is nil").Log()
	}
	if ctype == nil {
		return false, merrors.ContentCheckError{CalledBy: "types.CheckType"}.New("content type is nil").Log()
	}
	m.Model.ID = *id
	m.Model.ContentType = *ctype
	exists, err := m.CheckType(e)
	if err != nil {
		return false, merrors.ContentCheckError{CalledBy: "types.CheckType"}.Wrap(err).Log()
	}
	return exists, nil
}

