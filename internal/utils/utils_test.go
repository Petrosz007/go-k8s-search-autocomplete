package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniques(t *testing.T) {
	// Arrange
	tests := []struct {
		input, expected []string
	}{
		{[]string{}, []string{}},
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "a"}, []string{"a", "b"}},
		{[]string{"a", "a", "b", "a"}, []string{"a", "b"}},
		{[]string{"a", "b", "a", "b", "a"}, []string{"a", "b"}},
		{[]string{"a", "b", "a", "b", "a", "b", "b"}, []string{"a", "b"}},
	}

	for _, tt := range tests {
		t.Run("unique should work", func(t *testing.T) {
			// Act
			result := Uniques(tt.input)

			// Assert
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}
