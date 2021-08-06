package countingbits

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	timeout := time.After(1 * time.Second)
	var tests = []struct {
		input    int
		expected []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}

	done := make(chan bool)
	go func() {
		for _, test := range tests {
			actual := countBits(test.input)
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
