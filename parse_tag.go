package stool

import (
	"go/ast"
	"strings"
)

func ParseTag(tag *ast.BasicLit) map[string][]string {
	if tag == nil {
		return nil
	}

	out := make(map[string][]string)
	tags := strings.Split(strings.Trim(tag.Value, "`"), " ")
	for _, t := range tags {
		kv := strings.Split(t, ":")
		if len(kv) != 2 {
			continue
		}

		v := strings.Trim(kv[1], `"`)
		out[kv[0]] = strings.Split(v, ",")
	}
	return out
}
