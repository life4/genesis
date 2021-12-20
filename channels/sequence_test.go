package channels_test

import (
	"context"
	"testing"

	"github.com/life4/genesis/channels"
	"github.com/stretchr/testify/assert"
)

func TestSequenceCount(t *testing.T) {
	f := func(start int, step int, count int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		seq := channels.Counter(ctx, start, step)
		seq2 := channels.Take(seq, count)
		actual := channels.ToSlice(seq2)
		cancel()
		assert.Equal(t, expected, actual)
	}
	f(1, 2, 4, []int{1, 3, 5, 7})
}

func TestSequenceExponential(t *testing.T) {
	f := func(start int, factor int, count int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		seq := channels.Exponential(ctx, start, factor)
		seq2 := channels.Take(seq, count)
		actual := channels.ToSlice(seq2)
		cancel()
		assert.Equal(t, expected, actual)
	}
	f(1, 1, 4, []int{1, 1, 1, 1})
	f(1, 2, 4, []int{1, 2, 4, 8})
}

func TestSequenceIterate(t *testing.T) {
	f := func(start int, count int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		double := func(val int) int { return val * 2 }
		seq := channels.Iterate(ctx, start, double)
		seq2 := channels.Take(seq, count)
		actual := channels.ToSlice(seq2)
		cancel()
		assert.Equal(t, expected, actual)
	}
	f(1, 4, []int{1, 2, 4, 8})
}

func TestSequenceRange(t *testing.T) {
	f := func(start int, stop int, step int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		seq := channels.Range(ctx, start, stop, step)
		actual := channels.ToSlice(seq)
		cancel()
		assert.Equal(t, expected, actual)
	}
	f(1, 4, 1, []int{1, 2, 3})
	f(3, 0, -1, []int{3, 2, 1})
	f(1, 1, 1, []int{})
	f(1, 2, 1, []int{1})
}

func TestSequenceRepeat(t *testing.T) {
	f := func(count int, given int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		seq := channels.Repeat(ctx, given)
		seq2 := channels.Take(seq, count)
		actual := channels.ToSlice(seq2)
		cancel()
		assert.Equal(t, expected, actual)
	}
	f(2, 1, []int{1, 1})
}

func TestSequenceReplicate(t *testing.T) {
	f := func(count int, given int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		seq := channels.Replicate(ctx, given, count)
		actual := channels.ToSlice(seq)
		cancel()
		assert.Equal(t, expected, actual)
	}
	f(0, 1, []int{})
	f(1, 1, []int{1})
	f(5, 1, []int{1, 1, 1, 1, 1})
}
