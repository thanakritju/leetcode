package perfectsquares

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSquares(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{12, 3},
		{13, 2},
	}

	for _, test := range tests {
		actual := numSquares(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}
