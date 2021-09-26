package complexnumbermultiplication

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplexNumberMultiply(t *testing.T) {
	var tests = []struct {
		input    string
		input2   string
		expected string
	}{
		{"1+1i", "1+1i", "0+2i"},
		{"1+-1i", "1+-1i", "0+-2i"},
		{"78+-76i", "-86+72i", "-1236+12152i"},
	}

	for _, test := range tests {
		actual := complexNumberMultiply(test.input, test.input2)
		assert.Equal(t, test.expected, actual, test)
	}
}
