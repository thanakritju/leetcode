package minimumwindowsubstring

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	timeout := time.After(1 * time.Second)
	done := make(chan bool)
	go func() {
		var tests = []struct {
			inputS   string
			inputT   string
			expected string
		}{
			{"cabwefgewcwaefgcf", "cae", "cwae"},
			{"ab", "a", "a"},
			{"ab", "A", ""},
			{"a", "b", ""},
			{"ZZADOBECODEBANCZZ", "ABC", "BANC"},
			{"ADOBECODEBANC", "ABC", "BANC"},
			{"a", "a", "a"},
			{"a", "aa", ""},
			{"ab", "ba", "ab"},
			{"cccabcccc", "ba", "ab"},
		}

		for _, test := range tests {
			actual := minWindow(test.inputS, test.inputT)
			assert.Equal(t, test.expected, actual, test)
		}
		done <- true
	}()

	select {
	case <-timeout:
		t.Fatal("Test didn't finish in time")
	case <-done:
	}
}

func TestAreSubset(t *testing.T) {
	var tests = []struct {
		inputS   string
		inputT   string
		expected bool
	}{
		{"ADOBECODEBANC", "ABC", true},
		{"BANC", "ABC", true},
		{"a", "a", true},
		{"acc", "aa", false},
		{"ab", "ba", true},
	}

	for _, test := range tests {
		actual := areSubset(test.inputS, test.inputT)
		assert.Equal(t, test.expected, actual, test)
	}
}
