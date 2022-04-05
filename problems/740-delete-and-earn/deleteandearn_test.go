package deleteandearn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteAndEarn(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 3, 3, 3, 4}, 9},
		{[]int{3, 4, 2}, 6},
	}

	for _, test := range tests {
		actual := deleteAndEarn(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}
