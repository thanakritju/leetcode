package getpermutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPermutation(t *testing.T) {
	var tests = []struct {
		inputN   int
		inputK   int
		expected string
	}{
		{3, 3, "213"},
		{4, 9, "2314"},
		{3, 1, "123"},
	}

	for _, test := range tests {
		actual := getPermutation(test.inputN, test.inputK)
		assert.Equal(t, test.expected, actual, test)
	}
}
