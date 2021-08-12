package courseschedule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	var tests = []struct {
		inputNumCourses    int
		inputPrerequisites [][]int
		expected           bool
	}{
		{5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}, true},
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
		{20, [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}, false},
	}

	for _, test := range tests {
		actual := canFinish(test.inputNumCourses, test.inputPrerequisites)
		assert.Equal(t, test.expected, actual, test)
	}
}
