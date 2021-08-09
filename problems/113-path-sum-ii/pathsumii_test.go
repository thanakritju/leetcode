package pathsumii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum(t *testing.T) {
	var tests = []struct {
		inputArray     []int
		inputTargetSum int
		expected       [][]int
	}{
		{[]int{1, 2}, 0, [][]int{}},
		{[]int{1, 2, 3}, 5, [][]int{}},
	}

	for _, test := range tests {
		actual := pathSum(&TreeNode{}, test.inputTargetSum)
		assert.Equal(t, test.expected, actual, test)
	}
}
