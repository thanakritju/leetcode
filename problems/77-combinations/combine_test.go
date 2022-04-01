package combinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	var tests = []struct {
		inputN   int
		inputK   int
		expected [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{1, 1, [][]int{{1}}},
	}

	for _, test := range tests {
		actual := combine(test.inputN, test.inputK)
		assert.Equal(t, test.expected, actual, test)
	}
}
