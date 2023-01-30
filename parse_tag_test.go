package stool

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaeseTag(t *testing.T) {
	tests := []struct {
		name string
		tag  *ast.BasicLit
		want map[string][]string
	}{
		{
			name: "ok single",
			tag: &ast.BasicLit{
				Value: "`tag01:ok`",
			},
			want: map[string][]string{"tag01": {"ok"}},
		},
		{
			name: "ok multiple keys",
			tag: &ast.BasicLit{
				Value: "`tag01:\"ok\" tag02:\"ok02\"`",
			},
			want: map[string][]string{
				"tag01": {"ok"},
				"tag02": {"ok02"},
			},
		},
		{
			name: "ok multiple values",
			tag: &ast.BasicLit{
				Value: "`tag01:\"ok,ok02\"`",
			},
			want: map[string][]string{
				"tag01": {"ok", "ok02"},
			},
		},
		{
			name: "not ok: invalid",
			tag: &ast.BasicLit{
				Value: "`tag`",
			},
			want: make(map[string][]string),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := ParseTag(tt.tag)
			assert.Equal(t, tt.want, got)
		})
	}
}
