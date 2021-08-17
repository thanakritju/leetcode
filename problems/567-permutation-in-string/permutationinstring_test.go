package permutationinstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOrder(t *testing.T) {
	var tests = []struct {
		inputS1  string
		inputS2  string
		expected bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"abc", "xxxxcabxxxx", true},
		{"abb", "xxxxabxxxx", false},
		{"abb", "xxxxbabxxxx", true},
		{"ab", "xbabxxxxab", true},
		{"abc", "ccccbbbbaaaa", false},
	}

	for _, test := range tests {
		actual := checkInclusion(test.inputS1, test.inputS2)
		assert.Equal(t, test.expected, actual, test)
	}
}
