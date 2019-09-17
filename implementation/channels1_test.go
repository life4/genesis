package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	f := func(count int, given T, expected []T) {
		actual := Take(Repeat(1), count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []T{1, 1})
}

func TestTake(t *testing.T) {
	f := func(count int, given T, expected []T) {
		actual := Take(Repeat(1), count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []T{})
	f(1, 1, []T{1})
	f(2, 1, []T{1, 1})
}
