package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequenceCount(t *testing.T) {
	s := Sequence{}
	f := func(start T, step T, count int, expected []T) {
		seq := s.Count(start, step)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []T{1, 3, 5, 7})
}

func TestSequenceExponential(t *testing.T) {
	s := Sequence{}
	f := func(start T, factor T, count int, expected []T) {
		seq := s.Exponential(start, factor)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 1, 4, []T{1, 1, 1, 1})
	f(1, 2, 4, []T{1, 2, 4, 8})
}

func TestSequenceRange(t *testing.T) {
	s := Sequence{}
	f := func(start T, stop T, step T, expected []T) {
		seq := s.Range(start, stop, step)
		actual := Channel{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []T{1, 2, 3})
	f(3, 0, -1, []T{3, 2, 1})
	f(1, 1, 1, []T{})
	f(1, 2, 1, []T{1})
}

func TestSequenceRepeat(t *testing.T) {
	s := Sequence{}
	f := func(count int, given T, expected []T) {
		seq := s.Repeat(given)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []T{1, 1})
}
