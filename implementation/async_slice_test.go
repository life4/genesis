package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsyncSliceAny(t *testing.T) {
	f := func(check func(t T) bool, given []T, expected bool) {
		s := AsyncSlice{Data: given, workers: 2}
		actual := s.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t T) bool { return (t % 2) == 0 }

	f(isEven, []T{}, false)
	f(isEven, []T{1}, false)
	f(isEven, []T{1, 3}, false)
	f(isEven, []T{2}, true)
	f(isEven, []T{1, 2}, true)
	f(isEven, []T{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []T{1, 3, 5, 7, 9, 12}, true)
}

func TestAsyncSliceAll(t *testing.T) {
	f := func(check func(t T) bool, given []T, expected bool) {
		s := AsyncSlice{Data: given, workers: 2}
		actual := s.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t T) bool { return (t % 2) == 0 }

	f(isEven, []T{}, true)
	f(isEven, []T{1}, false)
	f(isEven, []T{1, 3}, false)
	f(isEven, []T{2}, true)
	f(isEven, []T{2, 4}, true)
	f(isEven, []T{2, 3}, false)
	f(isEven, []T{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []T{2, 4, 6, 8, 10, 11}, false)
}

func TestAsyncSliceEach(t *testing.T) {
	f := func(given []T) {
		s := AsyncSlice{Data: given, workers: 2}
		result := make(chan T, len(given))
		mapper := func(t T) { result <- t }
		s.Each(mapper)
		close(result)
		actual := Channel{result}.ToSlice()
		sorted := Slice{actual}.Sort()
		assert.Equal(t, given, sorted, "they should be equal")
	}

	f([]T{})
	f([]T{1})
	f([]T{1, 2, 3})
	f([]T{1, 2, 3, 4, 5, 6, 7})
}

func TestAsyncSliceMap(t *testing.T) {
	f := func(mapper func(t T) G, given []T, expected []G) {
		s := AsyncSlice{Data: given, workers: 2}
		actual := s.Map(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t T) G { return G((t * 2)) }

	f(double, []T{}, []G{})
	f(double, []T{1}, []G{2})
	f(double, []T{1, 2, 3}, []G{2, 4, 6})
}
func TestAsyncSliceReduce(t *testing.T) {
	f := func(reducer func(a T, b T) T, given []T, expected T) {
		s := AsyncSlice{Data: given, workers: 4}
		actual := s.Reduce(reducer)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	sum := func(a T, b T) T { return a + b }

	f(sum, []T{}, 0)
	f(sum, []T{1}, 1)
	f(sum, []T{1, 2}, 3)
	f(sum, []T{1, 2, 3}, 6)
	f(sum, []T{1, 2, 3, 4}, 10)
	f(sum, []T{1, 2, 3, 4, 5}, 15)
}
