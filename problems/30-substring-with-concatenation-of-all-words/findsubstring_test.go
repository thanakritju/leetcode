package findsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubstring(t *testing.T) {
	var tests = []struct {
		inputS     string
		inputWords []string
		expected   []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
	}

	for _, test := range tests {
		actual := findSubstring(test.inputS, test.inputWords)
		assert.Equal(t, test.expected, actual, test)
	}
}
