package minimumoperationstomakebinaryarrayelementsequaltoonei

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperation(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{0, 1, 1, 1, 0, 0}, 3},
		{[]int{0, 1, 1, 1}, -1},
	}

	for _, test := range tests {
		actual := minOperations(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}
