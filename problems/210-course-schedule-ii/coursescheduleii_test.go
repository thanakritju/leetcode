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
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 2, 1, 3}},
		{2, [][]int{{1, 0}}, []int{0, 1}},
		{2, [][]int{{1, 0}, {0, 1}}, []int{}},
		{1, [][]int{}, []int{0}},
	}

	for _, test := range tests {
		actual := findOrder(test.inputNumCourses, test.inputPrerequisites)
		assert.Equal(t, test.expected, actual, test)
	}
}
