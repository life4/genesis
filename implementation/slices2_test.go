package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkBy(t *testing.T) {
	f := func(mapper func(t T) G, given []T, expected [][]T) {
		actual := ChunkBy(given, mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t T) G { return G((t % 2)) }

	f(remainder, []T{1}, [][]T{{1}})
	f(remainder, []T{1, 2, 3}, [][]T{{1}, {2}, {3}})
	f(remainder, []T{1, 3, 2, 4, 5}, [][]T{{1, 3}, {2, 4}, {5}})
}

func TestMap(t *testing.T) {
	f := func(mapper func(t T) G, given []T, expected []G) {
		actual := Map(given, mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t T) G { return G((t * 2)) }

	f(double, []T{1, 2, 3}, []G{2, 4, 6})
}
