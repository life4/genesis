package slices_test

import (
	"testing"

	"github.com/life4/genesis/channels"
	"github.com/life4/genesis/slices"
	"github.com/stretchr/testify/assert"
)

func TestAnyAsync(t *testing.T) {
	f := func(check func(t int) bool, given []int, expected bool) {
		actual := slices.AnyAsync(given, 2, check)
		assert.Equal(t, expected, actual)
	}
	isEven := func(t int) bool { return (t % 2) == 0 }

	f(isEven, []int{}, false)
	f(isEven, []int{1}, false)
	f(isEven, []int{1, 3}, false)
	f(isEven, []int{2}, true)
	f(isEven, []int{1, 2}, true)
	f(isEven, []int{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []int{1, 3, 5, 7, 9, 12}, true)
}

func TestAllAsync(t *testing.T) {
	f := func(check func(t int) bool, given []int, expected bool) {
		actual := slices.AllAsync(given, 2, check)
		assert.Equal(t, expected, actual)
	}
	isEven := func(t int) bool { return (t % 2) == 0 }

	f(isEven, []int{}, true)
	f(isEven, []int{1}, false)
	f(isEven, []int{1, 3}, false)
	f(isEven, []int{2}, true)
	f(isEven, []int{2, 4}, true)
	f(isEven, []int{2, 3}, false)
	f(isEven, []int{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []int{2, 4, 6, 8, 10, 11}, false)
}

func TestEachAsync(t *testing.T) {
	f := func(given []int) {
		result := make(chan int, len(given))
		mapper := func(t int) { result <- t }
		slices.EachAsync(given, 2, mapper)
		close(result)
		actual := channels.ToSlice(result)
		sorted := slices.Sort(actual)
		assert.Equal(t, given, sorted)
	}

	f([]int{})
	f([]int{1})
	f([]int{1, 2, 3})
	f([]int{1, 2, 3, 4, 5, 6, 7})
}

func TestFilterAsync(t *testing.T) {
	f := func(given []int, expected []int) {
		filter := func(t int) bool { return t > 10 }
		actual := slices.FilterAsync(given, 2, filter)
		assert.Equal(t, expected, actual)
	}

	f([]int{}, []int{})
	f([]int{5}, []int{})
	f([]int{15}, []int{15})
	f([]int{9, 11, 12, 13, 6}, []int{11, 12, 13})
}

func TestMapAsync(t *testing.T) {
	f := func(mapper func(t int) int, given []int, expected []int) {
		actual := slices.MapAsync(given, 2, mapper)
		assert.Equal(t, expected, actual)
	}
	double := func(t int) int { return (t * 2) }

	f(double, []int{}, []int{})
	f(double, []int{1}, []int{2})
	f(double, []int{1, 2, 3}, []int{2, 4, 6})
}
func TestReduceAsync(t *testing.T) {
	f := func(reducer func(a int, b int) int, given []int, expected int) {
		actual := slices.ReduceAsync(given, 4, reducer)
		assert.Equal(t, expected, actual)
	}
	sum := func(a int, b int) int { return a + b }

	f(sum, []int{}, 0)
	f(sum, []int{1}, 1)
	f(sum, []int{1, 2}, 3)
	f(sum, []int{1, 2, 3}, 6)
	f(sum, []int{1, 2, 3, 4}, 10)
	f(sum, []int{1, 2, 3, 4, 5}, 15)
}
