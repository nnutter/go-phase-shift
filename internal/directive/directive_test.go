package directive_test

import (
	"testing"

	"github.com/phasemerge/phase-shift-go/internal/directive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		comment string
		want    []directive.Check
	}{
		{
			name:    "bare tag",
			comment: "//phase:nonmutating",
			want: []directive.Check{
				{Name: "nonmutating"},
			},
		},
		{
			name:    "not a phase tag",
			comment: "//go:noinline",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := directive.Parse(test.comment)
			require.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParseErrors(t *testing.T) {
	tests := []struct {
		name    string
		comment string
	}{
		{name: "missing name", comment: "//phase:"},
		{name: "missing quote", comment: "//phase:nonmutating:\"unterminated"},
		{name: "unquoted value", comment: "//phase:nonmutating:strict"},
		{name: "invalid check tag", comment: "//phase:othercheck"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := directive.Parse(test.comment)
			assert.Error(t, err)
		})
	}
}

func TestHasNonmutating(t *testing.T) {
	assert.True(t, directive.HasNonmutating("//phase:nonmutating"))
	assert.False(t, directive.HasNonmutating("//go:noinline"))
}
