package strrep

import (
	"fmt"
	"strings"
)

func Strrep(str string, vars map[string]interface{}) string {
	for k, v := range vars {
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
	return str
}