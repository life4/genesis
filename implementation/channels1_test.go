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
func TestProduct(t *testing.T) {
	f := func(arrs [][]T, expected [][]T) {
		actual := make([][]T, 0)
		i := 0
		for el := range Product(arrs...) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]T{{1, 2}, {3, 4}}, [][]T{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]T{{1, 2}, {3}, {4, 5}}, [][]T{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
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
