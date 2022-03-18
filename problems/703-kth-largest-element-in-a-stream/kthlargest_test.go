package kthlargest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargest1(t *testing.T) {
	obj := Constructor(3, []int{4, 5, 8, 2})
	assert.Equal(t, 4, obj.Add(3))
	assert.Equal(t, 5, obj.Add(5))
	assert.Equal(t, 5, obj.Add(10))
	assert.Equal(t, 8, obj.Add(9))
	assert.Equal(t, 8, obj.Add(4))
}

func TestKthLargest2(t *testing.T) {
	obj := Constructor(1, []int{-2})
	assert.Equal(t, -2, obj.Add(-3))
	assert.Equal(t, 0, obj.Add(0))
	assert.Equal(t, 2, obj.Add(2))
	assert.Equal(t, 2, obj.Add(-1))
	assert.Equal(t, 4, obj.Add(4))
}
