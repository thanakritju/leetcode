package relativeranks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeRanks(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []string
	}{
		{[]int{5, 4, 3, 2, 1}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
		{[]int{10, 3, 8, 9, 4}, []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}},
	}

	for _, test := range tests {
		actual := findRelativeRanks(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}
