package maxevents

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxEvents(t *testing.T) {
	var tests = []struct {
		events   [][]int
		expected int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}}, 3},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}, 4},
	}

	for _, test := range tests {
		actual := maxEvents(test.events)
		assert.Equal(t, test.expected, actual, test)
	}
}
