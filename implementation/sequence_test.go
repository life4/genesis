package implementation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequenceCount(t *testing.T) {
	f := func(start T, step T, count int, expected []T) {
		ctx, cancel := context.WithCancel(context.Background())
		s := Sequence{ctx: ctx}
		seq := s.Count(start, step)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []T{1, 3, 5, 7})
}

func TestSequenceExponential(t *testing.T) {
	f := func(start T, factor T, count int, expected []T) {
		ctx, cancel := context.WithCancel(context.Background())
		s := Sequence{ctx: ctx}
		seq := s.Exponential(start, factor)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 1, 4, []T{1, 1, 1, 1})
	f(1, 2, 4, []T{1, 2, 4, 8})
}

func TestSequenceIterate(t *testing.T) {
	f := func(start T, count int, expected []T) {
		ctx, cancel := context.WithCancel(context.Background())
		s := Sequence{ctx: ctx}
		double := func(val T) T { return val * 2 }
		seq := s.Iterate(start, double)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, []T{1, 2, 4, 8})
}

func TestSequenceRange(t *testing.T) {
	f := func(start T, stop T, step T, expected []T) {
		ctx, cancel := context.WithCancel(context.Background())
		s := Sequence{ctx: ctx}
		seq := s.Range(start, stop, step)
		actual := Channel{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []T{1, 2, 3})
	f(3, 0, -1, []T{3, 2, 1})
	f(1, 1, 1, []T{})
	f(1, 2, 1, []T{1})
}

func TestSequenceRepeat(t *testing.T) {
	f := func(count int, given T, expected []T) {
		ctx, cancel := context.WithCancel(context.Background())
		s := Sequence{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []T{1, 1})
}
