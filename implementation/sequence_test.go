package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExponential(t *testing.T) {
	s := Sequence{}
	f := func(start T, factor T, count int, expected []T) {
		seq := s.Exponential(start, factor)
		actual := Channel{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []T{1, 2, 4, 8})
}

func TestRange(t *testing.T) {
	s := Sequence{}
	f := func(start T, stop T, step T, expected []T) {
		seq := s.Range(start, stop, step)
		actual := Channel{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []T{1, 2, 3})
	f(3, 0, -1, []T{3, 2, 1})
}

func TestRepeat(t *testing.T) {
	s := Sequence{}
	f := func(count int, given T, expected []T) {
		seq := s.Repeat(given)
		actual := Channel{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []T{1, 1})
}

func TestTake(t *testing.T) {
	s := Sequence{}
	f := func(count int, given T, expected []T) {
		seq := s.Repeat(given)
		actual := Channel{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []T{})
	f(1, 1, []T{1})
	f(2, 1, []T{1, 1})
}

func TestToSlice(t *testing.T) {
	s := Sequence{}
	f := func(start T, stop T, step T, expected []T) {
		seq := s.Range(start, stop, step)
		actual := Channel{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []T{1, 2, 3})
}
