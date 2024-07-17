package converttotitle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	var tests = []struct {
		input    int
		expected string
	}{
		{1, "A"},
		{26, "Z"},
		{27, "AA"},
		{28, "AB"},
		{701, "ZY"},
	}

	for _, test := range tests {
		actual := convertToTitle(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}
