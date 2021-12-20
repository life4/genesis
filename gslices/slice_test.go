package gslices_test

import (
	"testing"

	"github.com/life4/genesis/gchannels"
	"github.com/life4/genesis/gerrors"
	"github.com/life4/genesis/gslices"
	"github.com/stretchr/testify/require"
)

func TestSliceAny(t *testing.T) {
	f := func(given []int, expected bool) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := gslices.Any(given, even)
		require.Equal(t, expected, actual)
	}
	f([]int{}, false)
	f([]int{1, 3}, false)
	f([]int{2}, true)
	f([]int{1, 2}, true)
}

func TestSliceAll(t *testing.T) {
	f := func(given []int, expected bool) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := gslices.All(given, even)
		require.Equal(t, expected, actual)
	}
	f([]int{}, true)
	f([]int{2}, true)
	f([]int{1}, false)
	f([]int{2, 4}, true)
	f([]int{2, 4, 1}, false)
	f([]int{1, 2, 4}, false)
}

func TestSliceChoice(t *testing.T) {
	f := func(given []int, seed int64, expected int) {
		actual, _ := gslices.Choice(given, seed)
		require.Equal(t, expected, actual)
	}
	f([]int{1}, 0, 1)
	f([]int{1, 2, 3}, 1, 3)
	f([]int{1, 2, 3}, 2, 2)
}

func TestSliceChunkBy(t *testing.T) {
	f := func(given []int, expected [][]int) {
		reminder := func(t int) int { return (t % 2) }
		actual := gslices.ChunkBy(given, reminder)
		require.Equal(t, expected, actual)
	}
	f([]int{}, [][]int{})
	f([]int{1}, [][]int{{1}})
	f([]int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f([]int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEvery(t *testing.T) {
	f := func(count int, given []int, expected [][]int) {
		actual, _ := gslices.ChunkEvery(given, count)
		require.Equal(t, expected, actual)
	}
	f(2, []int{}, [][]int{})
	f(2, []int{1}, [][]int{{1}})
	f(2, []int{1, 2, 3, 4}, [][]int{{1, 2}, {3, 4}})
	f(2, []int{1, 2, 3, 4, 5}, [][]int{{1, 2}, {3, 4}, {5}})
}

func TestSliceContains(t *testing.T) {
	f := func(el int, given []int, expected bool) {
		actual := gslices.Contains(given, el)
		require.Equal(t, expected, actual)
	}
	f(1, []int{}, false)
	f(1, []int{1}, true)
	f(1, []int{2}, false)
	f(1, []int{2, 3, 4, 5}, false)
	f(1, []int{2, 3, 1, 4, 5}, true)
	f(1, []int{2, 3, 1, 1, 4, 5}, true)
}

func TestSliceCount(t *testing.T) {
	f := func(el int, given []int, expected int) {
		actual := gslices.Count(given, el)
		require.Equal(t, expected, actual)
	}
	f(1, []int{}, 0)
	f(1, []int{1}, 1)
	f(1, []int{2}, 0)
	f(1, []int{2, 3, 4, 5}, 0)
	f(1, []int{2, 3, 1, 4, 5}, 1)
	f(1, []int{2, 3, 1, 1, 4, 5}, 2)
	f(1, []int{1, 1, 1, 1, 1}, 5)
}

func TestSliceCountBy(t *testing.T) {
	f := func(given []int, expected int) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := gslices.CountBy(given, even)
		require.Equal(t, expected, actual)
	}
	f([]int{}, 0)
	f([]int{1}, 0)
	f([]int{2}, 1)
	f([]int{1, 2, 3, 4, 5}, 2)
	f([]int{1, 2, 3, 4, 5, 6}, 3)
}

func TestSliceCycle(t *testing.T) {
	f := func(count int, given []int, expected []int) {
		c := gslices.Cycle(given)
		seq := gchannels.Take(c, count)
		actual := gchannels.ToSlice(seq)
		require.Equal(t, expected, actual)
	}
	f(5, []int{}, []int{})
	f(5, []int{1}, []int{1, 1, 1, 1, 1})
	f(5, []int{1, 2}, []int{1, 2, 1, 2, 1})
}

func TestSliceDedup(t *testing.T) {
	f := func(given []int, expected []int) {
		actual := gslices.Dedup(given)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3, 3, 3, 2, 1, 1}, []int{1, 2, 3, 2, 1})
}

func TestSliceDedupBy(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(el int) int { return el % 2 }
		actual := gslices.DedupBy(given, even)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 4, 3, 5, 7, 10}, []int{1, 2, 3, 10})
}

func TestSliceDelete(t *testing.T) {
	f := func(given []int, el int, expected []int) {
		actual := gslices.Delete(given, el)
		require.Equal(t, expected, actual)
	}
	f([]int{}, 1, []int{})
	f([]int{1}, 1, []int{})
	f([]int{2}, 1, []int{2})
	f([]int{1, 2}, 1, []int{2})
	f([]int{1, 2, 3}, 2, []int{1, 3})
	f([]int{1, 2, 2, 3, 2}, 2, []int{1, 2, 3, 2})
}

func TestSliceDeleteAll(t *testing.T) {
	f := func(given []int, el int, expected []int) {
		actual := gslices.DeleteAll(given, el)
		require.Equal(t, expected, actual)
	}
	f([]int{}, 1, []int{})
	f([]int{1}, 1, []int{})
	f([]int{2}, 1, []int{2})
	f([]int{1, 2}, 1, []int{2})
	f([]int{1, 2, 3}, 2, []int{1, 3})
	f([]int{1, 2, 2, 3, 2}, 2, []int{1, 3})
}

func TestSliceDeleteAt(t *testing.T) {
	f := func(given []int, indices []int, expected []int) {
		actual, _ := gslices.DeleteAt(given, indices...)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{}, []int{})
	f([]int{1}, []int{0}, []int{})
	f([]int{1, 2}, []int{0}, []int{2})

	f([]int{1, 2, 3}, []int{0}, []int{2, 3})
	f([]int{1, 2, 3}, []int{1}, []int{1, 3})
	f([]int{1, 2, 3}, []int{2}, []int{1, 2})

	f([]int{1, 2, 3}, []int{0, 1}, []int{3})
	f([]int{1, 2, 3}, []int{0, 2}, []int{2})
	f([]int{1, 2, 3}, []int{1, 2}, []int{1})
}

func TestSliceDropEvery(t *testing.T) {
	f := func(given []int, nth int, from int, expected []int) {
		actual, _ := gslices.DropEvery(given, nth, from)
		require.Equal(t, expected, actual)
	}
	f([]int{}, 1, 1, []int{})
	f([]int{1, 2, 3}, 1, 1, []int{})

	f([]int{1, 2, 3, 4}, 2, 1, []int{1, 3})
	f([]int{1, 2, 3, 4, 5}, 2, 1, []int{1, 3, 5})

	f([]int{1, 2, 3, 4}, 2, 0, []int{2, 4})
	f([]int{1, 2, 3, 4, 5}, 2, 0, []int{2, 4})

	f([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, []int{1, 3, 5, 7, 9})
	f([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 1, []int{1, 2, 4, 5, 7, 8, 10})
}

func TestSliceDropWhile(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(el int) bool { return el%2 == 0 }
		actual := gslices.DropWhile(given, even)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{})
	f([]int{2}, []int{})
	f([]int{1}, []int{1})
	f([]int{2, 1}, []int{1})
	f([]int{2, 1, 2}, []int{1, 2})
	f([]int{1, 2}, []int{1, 2})
	f([]int{2, 4, 6, 1, 8}, []int{1, 8})
}

func TestSliceEndsWith(t *testing.T) {
	f := func(given []int, suffix []int, expected bool) {
		actual := gslices.EndsWith(given, suffix)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{}, true)
	f([]int{1}, []int{1}, true)
	f([]int{1}, []int{2}, false)
	f([]int{2, 3}, []int{1, 2, 3}, false)

	f([]int{1, 2, 3}, []int{3}, true)
	f([]int{1, 2, 3}, []int{2, 3}, true)
	f([]int{1, 2, 3}, []int{1, 2, 3}, true)

	f([]int{1, 2, 3}, []int{1}, false)
	f([]int{1, 2, 3}, []int{2}, false)
	f([]int{1, 2, 3}, []int{1, 2}, false)
	f([]int{1, 2, 3}, []int{3, 2}, false)
}

func TestSliceEqual(t *testing.T) {
	f := func(left []int, right []int, expected bool) {
		actual := gslices.Equal(left, right)
		require.Equal(t, expected, actual)

		actual = gslices.Equal(right, left)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{}, true)
	f([]int{1}, []int{1}, true)
	f([]int{1}, []int{2}, false)
	f([]int{1, 2, 3, 3}, []int{1, 2, 3, 3}, true)
	f([]int{1, 2, 3, 3}, []int{1, 2, 2, 3}, false)
	f([]int{1, 2, 3, 3}, []int{1, 2, 4, 3}, false)

	// different len
	f([]int{1, 2, 3}, []int{1, 2}, false)
	f([]int{1, 2}, []int{1, 2, 3}, false)
	f([]int{}, []int{1, 2, 3}, false)
	f([]int{1, 2, 3}, []int{}, false)
}

func TestSliceFilter(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := gslices.Filter(given, even)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1, 2, 3, 4}, []int{2, 4})
	f([]int{1, 3}, []int{})
	f([]int{2, 4}, []int{2, 4})
}

func TestSliceFind(t *testing.T) {
	f := func(given []int, expectedEl int, expectedErr error) {
		even := func(t int) bool { return (t % 2) == 0 }
		el, err := gslices.Find(given, even)
		require.Equal(t, expectedEl, el)
		require.Equal(t, expectedErr, err)
	}
	f([]int{}, 0, gerrors.ErrNotFound)
	f([]int{1}, 0, gerrors.ErrNotFound)
	f([]int{1}, 0, gerrors.ErrNotFound)
	f([]int{2}, 2, nil)
	f([]int{1, 2}, 2, nil)
	f([]int{1, 2, 3}, 2, nil)
	f([]int{1, 3, 5}, 0, gerrors.ErrNotFound)
}

func TestSliceFindIndex(t *testing.T) {
	f := func(given []int, expectedInd int) {
		even := func(t int) bool { return (t % 2) == 0 }
		index := gslices.FindIndex(given, even)
		require.Equal(t, expectedInd, index)
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

func TestSliceJoin(t *testing.T) {
	is := require.New(t)
	is.Equal(gslices.Join([]int{}, ""), "")
	is.Equal(gslices.Join([]int{}, "|"), "")

	is.Equal(gslices.Join([]int{1}, ""), "1")
	is.Equal(gslices.Join([]int{1}, "|"), "1")

	is.Equal(gslices.Join([]int{1, 2, 3}, ""), "123")
	is.Equal(gslices.Join([]int{1, 2, 3}, "|"), "1|2|3")
	is.Equal(gslices.Join([]int{1, 2, 3}, "<T>"), "1<T>2<T>3")
	is.Equal(gslices.Join([]string{"hello", "world"}, " "), "hello world")
}

func TestSliceGroupBy(t *testing.T) {
	f := func(given []int, expected map[int][]int) {
		reminder := func(t int) int { return (t % 2) }
		actual := gslices.GroupBy(given, reminder)
		require.Equal(t, expected, actual)
	}
	f([]int{}, map[int][]int{})
	f([]int{1}, map[int][]int{1: {1}})
	f([]int{1, 3, 2, 4, 5}, map[int][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceInsertAt(t *testing.T) {
	f := func(given []int, index int, expected []int, expectedErr error) {
		actual, err := gslices.InsertAt(given, index, 10)
		require.Equal(t, expected, actual)
		require.Equal(t, expectedErr, err)
	}
	f([]int{}, -1, []int{}, gerrors.ErrNegativeValue)
	f([]int{}, 0, []int{10}, nil)
	f([]int{}, 1, []int{}, gerrors.ErrOutOfRange)

	f([]int{1, 2, 3}, -1, []int{1, 2, 3}, gerrors.ErrNegativeValue)
	f([]int{1, 2, 3}, 0, []int{10, 1, 2, 3}, nil)
	f([]int{1, 2, 3}, 1, []int{1, 10, 2, 3}, nil)
	f([]int{1, 2, 3}, 3, []int{1, 2, 3, 10}, nil)
	f([]int{1, 2, 3}, 4, []int{1, 2, 3}, gerrors.ErrOutOfRange)
}

func TestSliceIntersperse(t *testing.T) {
	f := func(el int, given []int, expected []int) {
		actual := gslices.Intersperse(given, el)
		require.Equal(t, expected, actual)
	}
	f(0, []int{}, []int{})
	f(0, []int{1}, []int{1})
	f(0, []int{1, 2}, []int{1, 0, 2})
	f(0, []int{1, 2, 3}, []int{1, 0, 2, 0, 3})
}

func TestSliceLast(t *testing.T) {
	f := func(given []int, expectedEl int, expectedErr error) {
		el, err := gslices.Last(given)
		require.Equal(t, expectedEl, el)
		require.Equal(t, expectedErr, err)
	}
	f([]int{}, 0, gerrors.ErrEmpty)
	f([]int{1}, 1, nil)
	f([]int{1, 2, 3}, 3, nil)
}

func TestSliceMap(t *testing.T) {
	f := func(given []int, expected []int) {
		double := func(t int) int { return (t * 2) }
		actual := gslices.Map(given, double)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{2})
	f([]int{1, 2, 3}, []int{2, 4, 6})
}

func TestSliceMax(t *testing.T) {
	f := func(given []int, expectedEl int, expectedErr error) {
		el, err := gslices.Max(given)
		require.Equal(t, expectedEl, el)
		require.Equal(t, expectedErr, err)
	}
	f([]int{}, 0, gerrors.ErrEmpty)
	f([]int{1}, 1, nil)
	f([]int{1, 2, 3}, 3, nil)
	f([]int{1, 3, 2}, 3, nil)
	f([]int{3, 2, 1}, 3, nil)
}

func TestSliceMin(t *testing.T) {
	f := func(given []int, expectedEl int, expectedErr error) {
		el, err := gslices.Min(given)
		require.Equal(t, expectedEl, el)
		require.Equal(t, expectedErr, err)
	}
	f([]int{}, 0, gerrors.ErrEmpty)
	f([]int{1}, 1, nil)
	f([]int{1, 2, 3}, 1, nil)
	f([]int{2, 1, 3}, 1, nil)
	f([]int{3, 2, 1}, 1, nil)
}

func TestSlicesPermutations(t *testing.T) {
	f := func(size int, given []int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		for el := range gslices.Permutations(given, size) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		require.Equal(t, expected, actual)
	}
	f(2, []int{}, [][]int{})
	f(2, []int{1}, [][]int{{1}})
	f(2, []int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}})
}

func TestSliceProduct(t *testing.T) {
	f := func(given []int, repeat int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		for el := range gslices.Product(given, repeat) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		require.Equal(t, expected, actual)
	}

	f([]int{1, 2}, 0, [][]int{})
	f([]int{}, 2, [][]int{})
	f([]int{1}, 2, [][]int{{1, 1}})

	f([]int{1, 2}, 1, [][]int{{1}, {2}})
	f([]int{1, 2}, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}})
	f([]int{1, 2}, 3, [][]int{
		{1, 1, 1}, {1, 1, 2}, {1, 2, 1}, {1, 2, 2},
		{2, 1, 1}, {2, 1, 2}, {2, 2, 1}, {2, 2, 2},
	})
}

func TestSliceReduce(t *testing.T) {
	f := func(given []int, expected int) {
		sum := func(el int, acc int) int { return (el) + acc }
		actual := gslices.Reduce(given, 0, sum)
		require.Equal(t, expected, actual)
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
}

func TestSliceReduceWhile(t *testing.T) {
	f := func(given []int, expected int) {
		sum := func(el int, acc int) (int, error) {
			if el == 0 {
				return acc, gerrors.ErrEmpty
			}
			return (el) + acc, nil
		}
		actual, _ := gslices.ReduceWhile(given, 0, sum)
		require.Equal(t, expected, actual)
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
	f([]int{1, 2, 0, 3}, 3)
}

func TestSliceReverse(t *testing.T) {
	f := func(given []int, expected []int) {
		actual := gslices.Reverse(given)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 2}, []int{2, 1})
	f([]int{1, 2, 3}, []int{3, 2, 1})
	f([]int{1, 2, 2, 3, 3}, []int{3, 3, 2, 2, 1})
}

func TestSliceSame(t *testing.T) {
	f := func(given []int, expected bool) {
		actual := gslices.Same(given)
		require.Equal(t, expected, actual)
	}
	f([]int{}, true)
	f([]int{1}, true)
	f([]int{1, 1}, true)
	f([]int{1, 1, 1}, true)

	f([]int{1, 2, 1}, false)
	f([]int{1, 2, 2}, false)
	f([]int{1, 1, 2}, false)
}

func TestSliceScan(t *testing.T) {
	f := func(given []int, expected []int) {
		sum := func(el int, acc int) int { return (el) + acc }
		actual := gslices.Scan(given, 0, sum)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 2}, []int{1, 3})
	f([]int{1, 2, 3}, []int{1, 3, 6})
	f([]int{1, 2, 3, 4}, []int{1, 3, 6, 10})
}

func TestSliceShuffle(t *testing.T) {
	f := func(given []int, seed int64, expected []int) {
		actual := gslices.Shuffle(given, seed)
		require.Equal(t, expected, actual)
	}
	f([]int{}, 0, []int{})
	f([]int{1}, 0, []int{1})
	f([]int{1, 2, 3, 4, 5, 6}, 2, []int{3, 5, 4, 1, 6, 2})
	f([]int{1, 2, 2, 3, 3}, 2, []int{3, 2, 3, 2, 1})
}

func TestSliceSorted(t *testing.T) {
	f := func(given []int, expected bool) {
		actual := gslices.Sorted(given)
		require.Equal(t, expected, actual)
	}
	f([]int{}, true)
	f([]int{1}, true)
	f([]int{1, 1}, true)
	f([]int{1, 2, 2}, true)
	f([]int{1, 2, 3}, true)

	f([]int{2, 1}, false)
	f([]int{1, 2, 1}, false)
}

func TestSliceSplit(t *testing.T) {
	f := func(given []int, sep int, expected [][]int) {
		actual := gslices.Split(given, sep)
		require.Equal(t, expected, actual)
	}
	f([]int{}, 1, [][]int{{}})
	f([]int{2}, 1, [][]int{{2}})
	f([]int{2, 1, 3}, 1, [][]int{{2}, {3}})
	f([]int{1, 3}, 1, [][]int{{}, {3}})
	f([]int{2, 1}, 1, [][]int{{2}, {}})
	f([]int{2, 1, 3, 4, 1, 5, 6, 7}, 1, [][]int{{2}, {3, 4}, {5, 6, 7}})
}

func TestSliceStartsWith(t *testing.T) {
	f := func(given []int, suffix []int, expected bool) {
		actual := gslices.StartsWith(given, suffix)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{}, true)
	f([]int{1}, []int{1}, true)
	f([]int{1}, []int{2}, false)
	f([]int{1, 2}, []int{1, 2, 3}, false)

	f([]int{1, 2, 3}, []int{1}, true)
	f([]int{1, 2, 3}, []int{1, 2}, true)
	f([]int{1, 2, 3}, []int{1, 2, 3}, true)

	f([]int{1, 2, 3}, []int{2}, false)
	f([]int{1, 2, 3}, []int{3}, false)
	f([]int{1, 2, 3}, []int{2, 3}, false)
	f([]int{1, 2, 3}, []int{3, 2}, false)
	f([]int{1, 2, 3}, []int{2, 1}, false)
}

func TestSliceSum(t *testing.T) {
	f := func(given []int, expected int) {
		actual := gslices.Sum(given)
		require.Equal(t, expected, actual)
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
}

func TestSliceTakeEvery(t *testing.T) {
	f := func(given []int, nth int, from int, expected []int) {
		actual, _ := gslices.TakeEvery(given, nth, from)
		require.Equal(t, expected, actual)
	}

	// step 1
	f([]int{}, 1, 1, []int{})
	f([]int{1, 2, 3}, 1, 0, []int{1, 2, 3})

	// step 2 from 0
	f([]int{1, 2, 3, 4, 5}, 2, 0, []int{1, 3, 5})
	f([]int{1, 2, 3, 4, 5, 6}, 2, 0, []int{1, 3, 5})

	// step 2 from 1
	f([]int{1, 2, 3, 4}, 2, 1, []int{2, 4})
	f([]int{1, 2, 3, 4, 5}, 2, 1, []int{2, 4})
}

func TestSliceTakeRandom(t *testing.T) {
	f := func(given []int, count int, seed int64, expected []int) {
		actual, _ := gslices.TakeRandom(given, count, seed)
		require.Equal(t, expected, actual)
	}
	f([]int{1}, 1, 0, []int{1})
	f([]int{1, 2, 3, 4, 5}, 3, 1, []int{3, 1, 2})
	f([]int{1, 2, 3, 4, 5}, 5, 1, []int{3, 1, 2, 5, 4})
}

func TestSliceTakeWhile(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(el int) bool { return el%2 == 0 }
		actual := gslices.TakeWhile(given, even)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{})
	f([]int{2}, []int{2})
	f([]int{2, 4, 6, 1, 8}, []int{2, 4, 6})
	f([]int{1, 2, 3}, []int{})
}

func TestSliceUniq(t *testing.T) {
	f := func(given []int, expected []int) {
		actual := gslices.Uniq(given)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 1}, []int{1, 2})
	f([]int{1, 2, 1, 2}, []int{1, 2})
	f([]int{1, 2, 1, 2, 3, 2, 1, 1}, []int{1, 2, 3})
}

func TestSliceWindow(t *testing.T) {
	f := func(given []int, size int, expected [][]int) {
		actual, _ := gslices.Window(given, size)
		require.Equal(t, expected, actual)
	}
	f([]int{}, 1, [][]int{})
	f([]int{1, 2, 3, 4}, 1, [][]int{{1}, {2}, {3}, {4}})
	f([]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {2, 3}, {3, 4}})
	f([]int{1, 2, 3, 4}, 3, [][]int{{1, 2, 3}, {2, 3, 4}})
	f([]int{1, 2, 3, 4}, 4, [][]int{{1, 2, 3, 4}})
}

func TestSliceWithout(t *testing.T) {
	f := func(given []int, items []int, expected []int) {
		actual := gslices.Without(given, items...)
		require.Equal(t, expected, actual)
	}
	f([]int{}, []int{}, []int{})
	f([]int{}, []int{1, 2}, []int{})

	f([]int{1}, []int{1}, []int{})
	f([]int{1}, []int{1, 2}, []int{})
	f([]int{1}, []int{2}, []int{1})

	f([]int{1, 2, 3, 4}, []int{1}, []int{2, 3, 4})
	f([]int{1, 2, 3, 4}, []int{2}, []int{1, 3, 4})
	f([]int{1, 2, 3, 4}, []int{4}, []int{1, 2, 3})

	f([]int{1, 2, 3, 4}, []int{1, 2}, []int{3, 4})
	f([]int{1, 2, 3, 4}, []int{1, 2, 4}, []int{3})
	f([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, []int{})
	f([]int{1, 2, 3, 4}, []int{2, 4}, []int{1, 3})

	f([]int{1, 1, 2, 3, 1, 4, 1}, []int{1}, []int{2, 3, 4})
}
