package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceAny(t *testing.T) {
	f := func(check func(t T) bool, given []T, expected bool) {
		actual := Slice{given}.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t T) bool { return (t % 2) == 0 }

	f(isEven, []T{}, false)
	f(isEven, []T{1, 3}, false)
	f(isEven, []T{2}, true)
	f(isEven, []T{1, 2}, true)
}

func TestSliceAll(t *testing.T) {
	f := func(check func(t T) bool, given []T, expected bool) {
		actual := Slice{given}.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t T) bool { return (t % 2) == 0 }

	f(isEven, []T{}, true)
	f(isEven, []T{2}, true)
	f(isEven, []T{1}, false)
	f(isEven, []T{2, 4}, true)
	f(isEven, []T{2, 4, 1}, false)
	f(isEven, []T{1, 2, 4}, false)
}

func TestSliceChunkBy(t *testing.T) {
	f := func(mapper func(t T) G, given []T, expected [][]T) {
		actual := Slice{given}.ChunkBy(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t T) G { return G((t % 2)) }

	f(remainder, []T{1}, [][]T{{1}})
	f(remainder, []T{1, 2, 3}, [][]T{{1}, {2}, {3}})
	f(remainder, []T{1, 3, 2, 4, 5}, [][]T{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEvery(t *testing.T) {
	f := func(count int, given []T, expected [][]T) {
		actual := Slice{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []T{1, 2, 3, 4}, [][]T{[]T{1, 2}, []T{3, 4}})
	f(2, []T{1, 2, 3, 4, 5}, [][]T{[]T{1, 2}, []T{3, 4}, []T{5}})
}

func TestSliceFilter(t *testing.T) {
	f := func(filter func(t T) bool, given []T, expected []T) {
		actual := Slice{given}.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t T) bool { return t > 0 }

	f(filterPositive, []T{1, -1, 2, -2, 3, -3}, []T{1, 2, 3})
	f(filterPositive, []T{1, 2, 3}, []T{1, 2, 3})
	f(filterPositive, []T{-1, -2, -3}, []T{})
}

func TestSliceGroupBy(t *testing.T) {
	f := func(mapper func(t T) G, given []T, expected map[G][]T) {
		actual := Slice{given}.GroupBy(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t T) G { return G((t % 2)) }

	f(remainder, []T{}, map[G][]T{})
	f(remainder, []T{1}, map[G][]T{1: {1}})
	f(remainder, []T{1, 3, 2, 4, 5}, map[G][]T{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceIntersperse(t *testing.T) {
	f := func(el T, given []T, expected []T) {
		actual := Slice{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []T{1, 2, 3}, []T{1, 0, 2, 0, 3})
}

func TestSliceMap(t *testing.T) {
	f := func(mapper func(t T) G, given []T, expected []G) {
		actual := Slice{given}.Map(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t T) G { return G((t * 2)) }

	f(double, []T{1, 2, 3}, []G{2, 4, 6})
}
