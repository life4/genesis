package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelToSlice(t *testing.T) {
	s := Sequence{}
	f := func(start T, stop T, step T, expected []T) {
		seq := s.Range(start, stop, step)
		actual := Channel{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []T{1, 2, 3})
}

func TestChannelAny(t *testing.T) {
	f := func(given []T, expected bool) {
		even := func(t T) bool { return t%2 == 0 }
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, false)
	f([]T{1}, false)
	f([]T{2}, true)
	f([]T{1, 2}, true)
	f([]T{1, 2, 3}, true)
	f([]T{1, 3, 5}, false)
	f([]T{1, 3, 5, 7, 9, 11}, false)
	f([]T{1, 3, 5, 7, 10, 11}, true)
}

func TestChannelAll(t *testing.T) {
	f := func(given []T, expected bool) {
		even := func(t T) bool { return t%2 == 0 }
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, true)
	f([]T{1}, false)
	f([]T{2}, true)
	f([]T{1, 2}, false)
	f([]T{2, 4}, true)
	f([]T{2, 4, 6, 8, 10, 12}, true)
	f([]T{2, 4, 6, 8, 11, 12}, false)
}

func TestChannelChunkEvery(t *testing.T) {
	f := func(size int, given []T, expected [][]T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := Channel{c}.ChunkEvery(size)
		actual := make([][]T, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []T{}, [][]T{})
	f(2, []T{1}, [][]T{{1}})
	f(2, []T{1, 2}, [][]T{{1, 2}})
	f(2, []T{1, 2, 3, 4}, [][]T{{1, 2}, {3, 4}})
	f(2, []T{1, 2, 3, 4, 5}, [][]T{{1, 2}, {3, 4}, {5}})
}

func TestChannelCount(t *testing.T) {
	f := func(element T, given []T, expected int) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.Count(element)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []T{}, 0)
	f(1, []T{1}, 1)
	f(1, []T{2}, 0)
	f(1, []T{1, 2, 3, 1, 4}, 2)
}
