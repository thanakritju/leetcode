package firstbadversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstBadVersion(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{5, 4},
	}

	for _, test := range tests {
		actual := firstBadVersion(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}
