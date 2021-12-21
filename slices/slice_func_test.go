package slices_test

import (
	"testing"

	"github.com/life4/genesis/slices"
	"github.com/matryer/is"
)

func TestAny(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected bool) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := slices.Any(given, even)
		is.Equal(expected, actual)
	}
	f([]int{}, false)
	f([]int{1, 3}, false)
	f([]int{2}, true)
	f([]int{1, 2}, true)
}

func TestAll(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected bool) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := slices.All(given, even)
		is.Equal(expected, actual)
	}
	f([]int{}, true)
	f([]int{2}, true)
	f([]int{1}, false)
	f([]int{2, 4}, true)
	f([]int{2, 4, 1}, false)
	f([]int{1, 2, 4}, false)
}
func TestChunkBy(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected [][]int) {
		reminder := func(t int) int { return (t % 2) }
		actual := slices.ChunkBy(given, reminder)
		is.Equal(expected, actual)
	}
	f([]int{}, [][]int{})
	f([]int{1}, [][]int{{1}})
	f([]int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f([]int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestCountFunc(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected int) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := slices.CountBy(given, even)
		is.Equal(expected, actual)
	}
	f([]int{}, 0)
	f([]int{1}, 0)
	f([]int{2}, 1)
	f([]int{1, 2, 3, 4, 5}, 2)
	f([]int{1, 2, 3, 4, 5, 6}, 3)
}

func TestDedupBy(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		even := func(el int) int { return el % 2 }
		actual := slices.DedupBy(given, even)
		is.Equal(expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 4, 3, 5, 7, 10}, []int{1, 2, 3, 10})
}

func TestDropWhile(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		even := func(el int) bool { return el%2 == 0 }
		actual := slices.DropWhile(given, even)
		is.Equal(expected, actual)
	}
	f([]int{}, []int{})
	f([]int{2}, []int{})
	f([]int{1}, []int{1})
	f([]int{2, 1}, []int{1})
	f([]int{2, 1, 2}, []int{1, 2})
	f([]int{1, 2}, []int{1, 2})
	f([]int{2, 4, 6, 1, 8}, []int{1, 8})
}

func TestFilter(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := slices.Filter(given, even)
		is.Equal(expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1, 2, 3, 4}, []int{2, 4})
	f([]int{1, 3}, []int{})
	f([]int{2, 4}, []int{2, 4})
}

func TestFind(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expectedEl int, expectedErr error) {
		even := func(t int) bool { return (t % 2) == 0 }
		el, err := slices.Find(given, even)
		is.Equal(expectedEl, el)
		is.Equal(expectedErr, err)
	}
	f([]int{}, 0, slices.ErrNotFound)
	f([]int{1}, 0, slices.ErrNotFound)
	f([]int{1}, 0, slices.ErrNotFound)
	f([]int{2}, 2, nil)
	f([]int{1, 2}, 2, nil)
	f([]int{1, 2, 3}, 2, nil)
	f([]int{1, 3, 5}, 0, slices.ErrNotFound)
}

func TestFindIndex(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expectedInd int) {
		even := func(t int) bool { return (t % 2) == 0 }
		index := slices.FindIndex(given, even)
		is.Equal(expectedInd, index)
	}
	f([]int{}, -1)
	f([]int{1}, -1)
	f([]int{1}, -1)
	f([]int{2}, 0)
	f([]int{1, 2}, 1)
	f([]int{1, 2, 3}, 1)
	f([]int{1, 3, 5, 7, 9, 2}, 5)
	f([]int{1, 3, 5}, -1)
}

func TestGroupBy(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected map[int][]int) {
		reminder := func(t int) int { return (t % 2) }
		actual := slices.GroupBy(given, reminder)
		is.Equal(expected, actual)
	}
	f([]int{}, map[int][]int{})
	f([]int{1}, map[int][]int{1: {1}})
	f([]int{1, 3, 2, 4, 5}, map[int][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestMap(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		double := func(t int) int { return (t * 2) }
		actual := slices.Map(given, double)
		is.Equal(expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{2})
	f([]int{1, 2, 3}, []int{2, 4, 6})
}

func TestReduce(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected int) {
		sum := func(el int, acc int) int { return (el) + acc }
		actual := slices.Reduce(given, 0, sum)
		is.Equal(expected, actual)
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
}

func TestReduceWhile(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected int) {
		sum := func(el int, acc int) (int, error) {
			if el == 0 {
				return acc, slices.ErrEmpty
			}
			return (el) + acc, nil
		}
		actual, _ := slices.ReduceWhile(given, 0, sum)
		is.Equal(expected, actual)
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
	f([]int{1, 2, 0, 3}, 3)
}

func TestScan(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		sum := func(el int, acc int) int { return (el) + acc }
		actual := slices.Scan(given, 0, sum)
		is.Equal(expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 2}, []int{1, 3})
	f([]int{1, 2, 3}, []int{1, 3, 6})
	f([]int{1, 2, 3, 4}, []int{1, 3, 6, 10})
}

func TestTakeWhile(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		even := func(el int) bool { return el%2 == 0 }
		actual := slices.TakeWhile(given, even)
		is.Equal(expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{})
	f([]int{2}, []int{2})
	f([]int{2, 4, 6, 1, 8}, []int{2, 4, 6})
	f([]int{1, 2, 3}, []int{})
}
