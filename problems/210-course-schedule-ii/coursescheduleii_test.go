package coursescheduleii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOrder(t *testing.T) {
	var tests = []struct {
		inputNumCourses    int
		inputPrerequisites [][]int
		expected           []int
	}{
		{5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}, []int{0, 1, 2, 3}},
		{2, [][]int{{1, 0}}, []int{0, 1}},
	}

	for _, test := range tests {
		actual := findOrder(test.inputNumCourses, test.inputPrerequisites)
		assert.Equal(t, test.expected, actual, test)
	}
}
