package slices_test

import (
	"testing"

	"github.com/life4/genesis/channels"
	"github.com/life4/genesis/lambdas"
	"github.com/life4/genesis/slices"
	"github.com/matryer/is"
)

func TestChoice(t *testing.T) {
	is := is.New(t)
	f := func(given []int, seed int64, expected int) {
		actual := lambdas.Must(slices.Choice(given, seed))
		is.Equal(expected, actual)
	}
	f([]int{1}, 0, 1)
	f([]int{1, 2, 3}, 1, 3)
	f([]int{1, 2, 3}, 2, 2)
}

func TestChunkEvery(t *testing.T) {
	is := is.New(t)
	f := func(count int, given []int, expected [][]int) {
		actual, _ := slices.ChunkEvery(given, count)
		is.Equal(expected, actual)
	}
	f(2, []int{}, [][]int{})
	f(2, []int{1}, [][]int{{1}})
	f(2, []int{1, 2, 3, 4}, [][]int{{1, 2}, {3, 4}})
	f(2, []int{1, 2, 3, 4, 5}, [][]int{{1, 2}, {3, 4}, {5}})
}

func TestContains(t *testing.T) {
	is := is.New(t)
	f := func(el int, given []int, expected bool) {
		actual := slices.Contains(given, el)
		is.Equal(expected, actual)
	}
	f(1, []int{}, false)
	f(1, []int{1}, true)
	f(1, []int{2}, false)
	f(1, []int{2, 3, 4, 5}, false)
	f(1, []int{2, 3, 1, 4, 5}, true)
	f(1, []int{2, 3, 1, 1, 4, 5}, true)
}

func TestCount(t *testing.T) {
	is := is.New(t)
	f := func(el int, given []int, expected int) {
		actual := slices.Count(given, el)
		is.Equal(expected, actual)
	}
	f(1, []int{}, 0)
	f(1, []int{1}, 1)
	f(1, []int{2}, 0)
	f(1, []int{2, 3, 4, 5}, 0)
	f(1, []int{2, 3, 1, 4, 5}, 1)
	f(1, []int{2, 3, 1, 1, 4, 5}, 2)
	f(1, []int{1, 1, 1, 1, 1}, 5)
}

func TestCycle(t *testing.T) {
	is := is.New(t)
	f := func(count int, given []int, expected []int) {
		c := slices.Cycle(given)
		seq := channels.Take(c, count)
		actual := channels.ToSlice(seq)
		is.Equal(expected, actual)
	}
	f(5, []int{}, []int{})
	f(5, []int{1}, []int{1, 1, 1, 1, 1})
	f(5, []int{1, 2}, []int{1, 2, 1, 2, 1})
}

func TestDedup(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		actual := slices.Dedup(given)
		is.Equal(expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3, 3, 3, 2, 1, 1}, []int{1, 2, 3, 2, 1})
}

func TestDelete(t *testing.T) {
	is := is.New(t)
	f := func(given []int, el int, expected []int) {
		actual := slices.Delete(given, el)
		is.Equal(expected, actual)
	}
	f([]int{}, 1, []int{})
	f([]int{1}, 1, []int{})
	f([]int{2}, 1, []int{2})
	f([]int{1, 2}, 1, []int{2})
	f([]int{1, 2, 3}, 2, []int{1, 3})
	f([]int{1, 2, 2, 3, 2}, 2, []int{1, 2, 3, 2})
}

func TestDeleteAll(t *testing.T) {
	is := is.New(t)
	f := func(given []int, el int, expected []int) {
		actual := slices.DeleteAll(given, el)
		is.Equal(expected, actual)
	}
	f([]int{}, 1, []int{})
	f([]int{1}, 1, []int{})
	f([]int{2}, 1, []int{2})
	f([]int{1, 2}, 1, []int{2})
	f([]int{1, 2, 3}, 2, []int{1, 3})
	f([]int{1, 2, 2, 3, 2}, 2, []int{1, 3})
}

func TestDeleteAt(t *testing.T) {
	is := is.New(t)
	f := func(given []int, indices []int, expected []int) {
		actual, _ := slices.DeleteAt(given, indices...)
		is.Equal(expected, actual)
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

func TestDropEvery(t *testing.T) {
	is := is.New(t)
	f := func(given []int, nth int, from int, expected []int) {
		actual, _ := slices.DropEvery(given, nth, from)
		is.Equal(expected, actual)
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

func TestEndsWith(t *testing.T) {
	is := is.New(t)
	f := func(given []int, suffix []int, expected bool) {
		actual := slices.EndsWith(given, suffix)
		is.Equal(expected, actual)
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

func TestEqual(t *testing.T) {
	is := is.New(t)
	f := func(left []int, right []int, expected bool) {
		actual := slices.Equal(left, right)
		is.Equal(expected, actual)

		actual = slices.Equal(right, left)
		is.Equal(expected, actual)
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

func TestJoin(t *testing.T) {
	is := is.New(t)
	is.Equal(slices.Join([]int{}, ""), "")
	is.Equal(slices.Join([]int{}, "|"), "")

	is.Equal(slices.Join([]int{1}, ""), "1")
	is.Equal(slices.Join([]int{1}, "|"), "1")

	is.Equal(slices.Join([]int{1, 2, 3}, ""), "123")
	is.Equal(slices.Join([]int{1, 2, 3}, "|"), "1|2|3")
	is.Equal(slices.Join([]int{1, 2, 3}, "<T>"), "1<T>2<T>3")
	is.Equal(slices.Join([]string{"hello", "world"}, " "), "hello world")
}

func TestInsertAt(t *testing.T) {
	is := is.New(t)
	f := func(given []int, index int, expected []int, expectedErr error) {
		actual, err := slices.InsertAt(given, index, 10)
		is.Equal(expected, actual)
		is.Equal(expectedErr, err)
	}
	f([]int{}, -1, []int{}, slices.ErrNegativeValue)
	f([]int{}, 0, []int{10}, nil)
	f([]int{}, 1, []int{}, slices.ErrOutOfRange)

	f([]int{1, 2, 3}, -1, []int{1, 2, 3}, slices.ErrNegativeValue)
	f([]int{1, 2, 3}, 0, []int{10, 1, 2, 3}, nil)
	f([]int{1, 2, 3}, 1, []int{1, 10, 2, 3}, nil)
	f([]int{1, 2, 3}, 3, []int{1, 2, 3, 10}, nil)
	f([]int{1, 2, 3}, 4, []int{1, 2, 3}, slices.ErrOutOfRange)
}

func TestIntersperse(t *testing.T) {
	is := is.New(t)
	f := func(el int, given []int, expected []int) {
		actual := slices.Intersperse(given, el)
		is.Equal(expected, actual)
	}
	f(0, []int{}, []int{})
	f(0, []int{1}, []int{1})
	f(0, []int{1, 2}, []int{1, 0, 2})
	f(0, []int{1, 2, 3}, []int{1, 0, 2, 0, 3})
}

func TestLast(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expectedEl int, expectedErr error) {
		el, err := slices.Last(given)
		is.Equal(expectedEl, el)
		is.Equal(expectedErr, err)
	}
	f([]int{}, 0, slices.ErrEmpty)
	f([]int{1}, 1, nil)
	f([]int{1, 2, 3}, 3, nil)
}

func TestMax(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expectedEl int, expectedErr error) {
		el, err := slices.Max(given)
		is.Equal(expectedEl, el)
		is.Equal(expectedErr, err)
	}
	f([]int{}, 0, slices.ErrEmpty)
	f([]int{1}, 1, nil)
	f([]int{1, 2, 3}, 3, nil)
	f([]int{1, 3, 2}, 3, nil)
	f([]int{3, 2, 1}, 3, nil)
}

func TestMin(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expectedEl int, expectedErr error) {
		el, err := slices.Min(given)
		is.Equal(expectedEl, el)
		is.Equal(expectedErr, err)
	}
	f([]int{}, 0, slices.ErrEmpty)
	f([]int{1}, 1, nil)
	f([]int{1, 2, 3}, 1, nil)
	f([]int{2, 1, 3}, 1, nil)
	f([]int{3, 2, 1}, 1, nil)
}

func TestPermutations(t *testing.T) {
	is := is.New(t)
	f := func(size int, given []int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		for el := range slices.Permutations(given, size) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		is.Equal(expected, actual)
	}
	f(2, []int{}, [][]int{})
	f(2, []int{1}, [][]int{{1}})
	f(2, []int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}})
}

func TestProduct(t *testing.T) {
	is := is.New(t)
	f := func(given []int, repeat int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		for el := range slices.Product(given, repeat) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		is.Equal(expected, actual)
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

func TestReverse(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		actual := slices.Reverse(given)
		is.Equal(expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 2}, []int{2, 1})
	f([]int{1, 2, 3}, []int{3, 2, 1})
	f([]int{1, 2, 2, 3, 3}, []int{3, 3, 2, 2, 1})
}

func TestRepeat(t *testing.T) {
	is := is.New(t)
	is.Equal(slices.Repeat([]int{}, 0), []int{})
	is.Equal(slices.Repeat([]int{1}, 0), []int{})
	is.Equal(slices.Repeat([]int{1, 2}, 0), []int{})

	is.Equal(slices.Repeat([]int{}, 1), []int{})
	is.Equal(slices.Repeat([]int{1}, 1), []int{1})
	is.Equal(slices.Repeat([]int{1, 2}, 1), []int{1, 2})

	is.Equal(slices.Repeat([]int{}, 3), []int{})
	is.Equal(slices.Repeat([]int{1}, 3), []int{1, 1, 1})
	is.Equal(slices.Repeat([]int{1, 2}, 3), []int{1, 2, 1, 2, 1, 2})
}

func TestSame(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected bool) {
		actual := slices.Same(given)
		is.Equal(expected, actual)
	}
	f([]int{}, true)
	f([]int{1}, true)
	f([]int{1, 1}, true)
	f([]int{1, 1, 1}, true)

	f([]int{1, 2, 1}, false)
	f([]int{1, 2, 2}, false)
	f([]int{1, 1, 2}, false)
}

func TestShuffle(t *testing.T) {
	is := is.New(t)
	f := func(given []int, seed int64, expected []int) {
		actual := slices.Shuffle(given, seed)
		is.Equal(expected, actual)
	}
	f([]int{}, 0, []int{})
	f([]int{1}, 0, []int{1})
	f([]int{1, 2, 3, 4, 5, 6}, 2, []int{3, 5, 4, 1, 6, 2})
	f([]int{1, 2, 2, 3, 3}, 2, []int{3, 2, 3, 2, 1})
}

func TestSorted(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected bool) {
		actual := slices.Sorted(given)
		is.Equal(expected, actual)
	}
	f([]int{}, true)
	f([]int{1}, true)
	f([]int{1, 1}, true)
	f([]int{1, 2, 2}, true)
	f([]int{1, 2, 3}, true)

	f([]int{2, 1}, false)
	f([]int{1, 2, 1}, false)
}

func TestSplit(t *testing.T) {
	is := is.New(t)
	f := func(given []int, sep int, expected [][]int) {
		actual := slices.Split(given, sep)
		is.Equal(expected, actual)
	}
	f([]int{}, 1, [][]int{{}})
	f([]int{2}, 1, [][]int{{2}})
	f([]int{2, 1, 3}, 1, [][]int{{2}, {3}})
	f([]int{1, 3}, 1, [][]int{{}, {3}})
	f([]int{2, 1}, 1, [][]int{{2}, {}})
	f([]int{2, 1, 3, 4, 1, 5, 6, 7}, 1, [][]int{{2}, {3, 4}, {5, 6, 7}})
}

func TestStartsWith(t *testing.T) {
	is := is.New(t)
	f := func(given []int, suffix []int, expected bool) {
		actual := slices.StartsWith(given, suffix)
		is.Equal(expected, actual)
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

func TestSum(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected int) {
		actual := slices.Sum(given)
		is.Equal(expected, actual)
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
}

func TestTakeEvery(t *testing.T) {
	is := is.New(t)
	f := func(given []int, nth int, from int, expected []int) {
		actual, _ := slices.TakeEvery(given, nth, from)
		is.Equal(expected, actual)
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

func TestTakeRandom(t *testing.T) {
	is := is.New(t)
	f := func(given []int, count int, seed int64, expected []int) {
		actual, _ := slices.TakeRandom(given, count, seed)
		is.Equal(expected, actual)
	}
	f([]int{1}, 1, 0, []int{1})
	f([]int{1, 2, 3, 4, 5}, 3, 1, []int{3, 1, 2})
	f([]int{1, 2, 3, 4, 5}, 5, 1, []int{3, 1, 2, 5, 4})
}

func TestUniq(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		actual := slices.Uniq(given)
		is.Equal(expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 1}, []int{1, 2})
	f([]int{1, 2, 1, 2}, []int{1, 2})
	f([]int{1, 2, 1, 2, 3, 2, 1, 1}, []int{1, 2, 3})
}

func TestWindow(t *testing.T) {
	is := is.New(t)
	f := func(given []int, size int, expected [][]int) {
		actual, _ := slices.Window(given, size)
		is.Equal(expected, actual)
	}
	f([]int{}, 1, [][]int{})
	f([]int{1, 2, 3, 4}, 1, [][]int{{1}, {2}, {3}, {4}})
	f([]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {2, 3}, {3, 4}})
	f([]int{1, 2, 3, 4}, 3, [][]int{{1, 2, 3}, {2, 3, 4}})
	f([]int{1, 2, 3, 4}, 4, [][]int{{1, 2, 3, 4}})
}

func TestWithout(t *testing.T) {
	is := is.New(t)
	f := func(given []int, items []int, expected []int) {
		actual := slices.Without(given, items...)
		is.Equal(expected, actual)
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
