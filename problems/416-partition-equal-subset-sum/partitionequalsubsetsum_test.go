package partitionequalsubsetsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartition(t *testing.T) {
	var tests = []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
	}

	for _, test := range tests {
		actual := canPartition(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}
