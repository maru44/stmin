package stmin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypePrefixFromString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  TypePrefix
	}{
		{
			name:  "ok: ptr",
			input: "*",
			want:  TypePrefix{IsPtr: true},
		},
		{
			name:  "ok: slice",
			input: "[]",
			want:  TypePrefix{IsSlice: true},
		},
		{
			name:  "ok: array",
			input: "[5]",
			want:  TypePrefix{IsArray: true, Len: 5},
		},
		{
			name:  "not ok: invalid pattern",
			input: "bad",
			want:  TypePrefix{},
		},
		{
			name:  "not ok: invalid pattern like array",
			input: "[bad]",
			want:  TypePrefix{},
		},
		{
			name:  "not ok: invalid pattern 2",
			input: "[bad",
			want:  TypePrefix{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := TypePrefixFromString(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
