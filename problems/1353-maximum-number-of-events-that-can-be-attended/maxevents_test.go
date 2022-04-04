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
		{[][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}, 4},
		{[][]int{{1, 2}, {2, 2}, {3, 3}, {3, 4}, {3, 4}}, 4},
		{[][]int{{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5}}, 5},
		{[][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3}}, 5},
	}

	for _, test := range tests {
		actual := maxEvents(test.events)
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestSortEvents(t *testing.T) {
	var tests = []struct {
		events   [][]int
		expected [][]int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 2}, {2, 3}, {3, 4}}},
		{[][]int{{3, 4}, {1, 3}, {2, 3}, {1, 2}}, [][]int{{1, 2}, {1, 3}, {2, 3}, {3, 4}}},
	}

	for _, test := range tests {
		sortEvents(test.events)
		assert.Equal(t, test.expected, test.events, test)
	}
}
