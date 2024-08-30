package ringsandrods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPoints(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"B0B6G0R6R0R6G9", 1},
		{"B0R0G0R9R0B0G0", 1},
		{"G4", 0},
	}

	for _, test := range tests {
		actual := countPoints(test.input)
		assert.Equal(t, test.expected, actual, test)
	}
}
