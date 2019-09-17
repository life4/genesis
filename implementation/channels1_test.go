package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExponential(t *testing.T) {
	f := func(start T, factor T, count int, expected []T) {
		actual := Take(Exponential(start, factor), count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []T{1, 2, 4, 8})
}
func TestRange(t *testing.T) {
	f := func(start T, stop T, step T, expected []T) {
		actual := TakeAll(Range(start, stop, step))
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []T{1, 2, 3})
	f(3, 0, -1, []T{3, 2, 1})
}

func TestRepeat(t *testing.T) {
	f := func(count int, given T, expected []T) {
		actual := Take(Repeat(given), count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []T{1, 1})
}

func TestTake(t *testing.T) {
	f := func(count int, given T, expected []T) {
		actual := Take(Repeat(given), count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []T{})
	f(1, 1, []T{1})
	f(2, 1, []T{1, 1})
}

func TestTakeAll(t *testing.T) {
	f := func(start T, stop T, step T, expected []T) {
		actual := TakeAll(Range(start, stop, step))
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []T{1, 2, 3})
}
