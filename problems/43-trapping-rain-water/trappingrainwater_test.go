package trappingrainwater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{0, 1, 0, 2, 2, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{4, 2, 3}, 1},
		{[]int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}, 23},
		{[]int{1000, 999, 998, 997, 996, 995, 994, 993, 992, 991, 990, 989, 988, 987, 986, 985, 984, 983, 982, 981, 980, 979, 978, 977, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966}, 0},
	}

	for _, test := range tests {
		actual := trap(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestIsLocalMaxima(t *testing.T) {
	var tests = []struct {
		inputIndex int
		inputArray []int
		expected   bool
	}{
		{0, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, false},
		{1, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, true},
		{3, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, true},
		{7, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, true},
		{1, []int{0, 1, 1, 1, 0}, true},
		{3, []int{0, 1, 1, 1, 0}, true},
		{3, []int{0, 1, 1, 1}, true},
		{1, []int{0, 1}, true},
		{2, []int{4, 2, 3}, true},
		{0, []int{4, 2, 3}, true},
	}

	for _, test := range tests {
		actual := isLocalMaxima(test.inputIndex, test.inputArray)
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestIsLocalMinima(t *testing.T) {
	var tests = []struct {
		inputIndex int
		inputArray []int
		expected   bool
	}{
		{2, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, true},
		{5, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, true},
		{7, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, false},
		{1, []int{2, 1}, true},
		{0, []int{2, 1}, false},
		{1, []int{3, 1, 1, 1, 3}, true},
		{3, []int{3, 1, 1, 1, 3}, true},
		{3, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, false},
		{4, []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}, true},
		{5, []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}, true},
		{9, []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}, true},
		{0, []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}, false},
		{1, []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}, false},
	}

	for _, test := range tests {
		actual := isLocalMinima(test.inputIndex, test.inputArray)
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestGetWaterBetweenIndexes(t *testing.T) {
	var tests = []struct {
		inputIndex1 int
		inputIndex2 int
		inputArray  []int
		expected    int
	}{
		{3, 7, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 4},
		{1, 3, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 1},
		{8, 10, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 1},
		{7, 10, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 1},
	}

	for _, test := range tests {
		actual := getWaterBetweenIndexes(test.inputIndex1, test.inputIndex2, test.inputArray)
		assert.Equal(t, test.expected, actual, test)
	}
}
