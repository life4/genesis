package genesis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterInt32(t *testing.T) {
	f := func(filter func(t int32) bool, given []int32, expected []int32) {
		actual := FilterInt32(given, filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t int32) bool { return t > 0 }

	f(filterPositive, []int32{1, -1, 2, -2, 3, -3}, []int32{1, 2, 3})
	f(filterPositive, []int32{1, 2, 3}, []int32{1, 2, 3})
	f(filterPositive, []int32{-1, -2, -3}, []int32{})
}

func TestMapInt32(t *testing.T) {
	f := func(mapper func(t int32) int32, given []int32, expected []int32) {
		actual := MapInt32(given, mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int32 { return t * 2 }

	f(double, []int32{1, 2, 3}, []int32{2, 4, 6})
}

func TestChunkEveryInt32(t *testing.T) {
	f := func(count int, given []int32, expected [][]int32) {
		actual := ChunkEveryInt32(given, count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int32{1, 2, 3, 4}, [][]int32{[]int32{1, 2}, []int32{3, 4}})
	f(2, []int32{1, 2, 3, 4, 5}, [][]int32{[]int32{1, 2}, []int32{3, 4}, []int32{5}})
}

func TestIntersperseInt32(t *testing.T) {
	f := func(el int32, given []int32, expected []int32) {
		actual := IntersperseInt32(given, el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int32{1, 2, 3}, []int32{1, 0, 2, 0, 3})
}
