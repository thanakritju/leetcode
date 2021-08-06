package trappingrainwater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	out := trap(input)

	assert.Equal(t, out, 6, "Example from leetcode")
}

func TestTrap2(t *testing.T) {
	input := []int{4, 2, 0, 3, 2, 5}

	out := trap(input)

	assert.Equal(t, out, 9, "Example from leetcode")
}
