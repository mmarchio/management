package strrep

import (
	"encoding/json"
	"fmt"
	"strings"

	merrors "github.com/mmarchio/management/errors"
)

func Strrep(str string, vars map[string]interface{}) (string, error) {
	for k, v := range vars {
		if sm, ok := v.(map[string]interface{}); ok {
			b, err := json.Marshal(sm)
			if err != nil {
				return "", merrors.JSONMarshallingError{}.Wrap(err).Log()
			}
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(s)}}", k), string(b))
		}
		if si, ok := v.([]interface{}); ok {
			var collector []string
			for _, siv := range si {
				if s, ok := siv.(string); ok {
					collector = append(collector, s)
				}
			}
			j := strings.Join(collector, ", ")
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(s)}}", k), j)
		}
		if s, ok := v.(string); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(s)}}", k), s)
		}
		if d, ok := v.(int); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(d)}}", k), string(d))
		}
		if d, ok := v.(int8); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(d)}}", k), string(d))
		}
		if d, ok := v.(int16); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(d)}}", k), string(d))
		}
		if d, ok := v.(int32); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(d)}}", k), string(d))
		}
		if d, ok := v.(int64); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(d)}}", k), string(d))
		}
		if d, ok := v.(uint); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(d)}}", k), string(d))
		}
		if d, ok := v.(uint8); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(d)}}", k), string(d))
		}
		if d, ok := v.(uint16); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(d)}}", k), string(d))
		}
		if d, ok := v.(uint32); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(d)}}", k), string(d))
		}
		if d, ok := v.(uint64); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(d)}}", k), string(d))
		}
		if f, ok := v.(float32); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(f)}}", k), fmt.Sprintf("%f", f))
		}
		if f, ok := v.(float64); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(f)}}", k), fmt.Sprintf("%f", f))
		}
		if b, ok := v.(bool); ok {
			str = strings.ReplaceAll(str, fmt.Sprintf("{{%s(b)}}", k), fmt.Sprintf("%t", b))
		}
	}
	return str, nil
}