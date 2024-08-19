package maximumprofitinjobscheduling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJobScheduling(t *testing.T) {
	var tests = []struct {
		startTime []int
		endTime   []int
		profit    []int
		expected  int
	}{
		{[]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}, 120},
		{[]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}, 150},
		{[]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}, 6},
	}

	t.Skip("skipping for now")

	for _, test := range tests {
		actual := jobScheduling(test.startTime, test.endTime, test.profit)
		assert.Equal(t, test.expected, actual, test)
	}
}
