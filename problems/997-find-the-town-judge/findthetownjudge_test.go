package findthetownjudge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindJudge(t *testing.T) {
	var tests = []struct {
		input      int
		inputTrust [][]int
		expected   int
	}{
		{2, [][]int{{1, 2}}, 2},
		{3, [][]int{{1, 3}, {2, 3}}, 3},
		{3, [][]int{{1, 2}, {2, 3}, {3, 1}}, -1},
		{3, [][]int{{1, 2}, {2, 3}}, -1},
		{4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}, 3},
	}

	for _, test := range tests {
		actual := findJudge(test.input, test.inputTrust)
		assert.Equal(t, test.expected, actual, test)
	}
}
