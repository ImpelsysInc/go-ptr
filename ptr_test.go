package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPtr(t *testing.T) {
	t.Run("StringPtr", func(t *testing.T) {
		emptyStr := ""
		nonEmptyStr := "non-empty"

		cases := []struct {
			name     string
			actual   *string
			expected string
		}{
			{
				name:     "empty string",
				actual:   StringPtr(emptyStr),
				expected: emptyStr,
			},
			{
				name:     "non-empty string",
				actual:   StringPtr(nonEmptyStr),
				expected: nonEmptyStr,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("StringValue", func(t *testing.T) {
		emptyStrPtr := StringPtr("")
		nonEmptyStrPtr := StringPtr("non-empty")

		cases := []struct {
			name     string
			actual   string
			expected string
		}{
			{
				name:     "empty string",
				actual:   StringValue(emptyStrPtr),
				expected: "",
			},
			{
				name:     "non-empty string",
				actual:   StringValue(nonEmptyStrPtr),
				expected: "non-empty",
			},
			{
				name:     "nil string",
				actual:   StringValue(nil),
				expected: "",
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})
}
