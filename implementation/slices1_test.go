package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkEvery(t *testing.T) {
	f := func(count int, given []T, expected [][]T) {
		actual := ChunkEvery(given, count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []T{1, 2, 3, 4}, [][]T{[]T{1, 2}, []T{3, 4}})
	f(2, []T{1, 2, 3, 4, 5}, [][]T{[]T{1, 2}, []T{3, 4}, []T{5}})
}

func TestFilter(t *testing.T) {
	f := func(filter func(t T) bool, given []T, expected []T) {
		actual := Filter(given, filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t T) bool { return t > 0 }

	f(filterPositive, []T{1, -1, 2, -2, 3, -3}, []T{1, 2, 3})
	f(filterPositive, []T{1, 2, 3}, []T{1, 2, 3})
	f(filterPositive, []T{-1, -2, -3}, []T{})
}

func TestMap(t *testing.T) {
	f := func(mapper func(t T) G, given []T, expected []G) {
		actual := Map(given, mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t T) G { return G((t * 2)) }

	f(double, []T{1, 2, 3}, []G{2, 4, 6})
}

func TestIntersperse(t *testing.T) {
	f := func(el T, given []T, expected []T) {
		actual := Intersperse(given, el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []T{1, 2, 3}, []T{1, 0, 2, 0, 3})
}
