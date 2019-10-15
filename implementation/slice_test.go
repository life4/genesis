package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceAny(t *testing.T) {
	f := func(given []T, expected bool) {
		even := func(t T) bool { return (t % 2) == 0 }
		actual := Slice{given}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, false)
	f([]T{1, 3}, false)
	f([]T{2}, true)
	f([]T{1, 2}, true)
}

func TestSliceAll(t *testing.T) {
	f := func(given []T, expected bool) {
		even := func(t T) bool { return (t % 2) == 0 }
		actual := Slice{given}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, true)
	f([]T{2}, true)
	f([]T{1}, false)
	f([]T{2, 4}, true)
	f([]T{2, 4, 1}, false)
	f([]T{1, 2, 4}, false)
}

func TestSliceChunkBy(t *testing.T) {
	f := func(given []T, expected [][]T) {
		reminder := func(t T) G { return G((t % 2)) }
		actual := Slice{given}.ChunkBy(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, [][]T{})
	f([]T{1}, [][]T{{1}})
	f([]T{1, 2, 3}, [][]T{{1}, {2}, {3}})
	f([]T{1, 3, 2, 4, 5}, [][]T{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEvery(t *testing.T) {
	f := func(count int, given []T, expected [][]T) {
		actual, _ := Slice{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []T{}, [][]T{})
	f(2, []T{1}, [][]T{{1}})
	f(2, []T{1, 2, 3, 4}, [][]T{{1, 2}, {3, 4}})
	f(2, []T{1, 2, 3, 4, 5}, [][]T{{1, 2}, {3, 4}, {5}})
}

func TestSliceContains(t *testing.T) {
	f := func(el T, given []T, expected bool) {
		actual := Slice{given}.Contains(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []T{}, false)
	f(1, []T{1}, true)
	f(1, []T{2}, false)
	f(1, []T{2, 3, 4, 5}, false)
	f(1, []T{2, 3, 1, 4, 5}, true)
	f(1, []T{2, 3, 1, 1, 4, 5}, true)
}

func TestSliceCount(t *testing.T) {
	f := func(el T, given []T, expected int) {
		actual := Slice{given}.Count(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []T{}, 0)
	f(1, []T{1}, 1)
	f(1, []T{2}, 0)
	f(1, []T{2, 3, 4, 5}, 0)
	f(1, []T{2, 3, 1, 4, 5}, 1)
	f(1, []T{2, 3, 1, 1, 4, 5}, 2)
	f(1, []T{1, 1, 1, 1, 1}, 5)
}

func TestSliceCountBy(t *testing.T) {
	f := func(given []T, expected int) {
		even := func(t T) bool { return (t % 2) == 0 }
		actual := Slice{given}.CountBy(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 0)
	f([]T{1}, 0)
	f([]T{2}, 1)
	f([]T{1, 2, 3, 4, 5}, 2)
	f([]T{1, 2, 3, 4, 5, 6}, 3)
}

func TestSliceCycle(t *testing.T) {
	f := func(count int, given []T, expected []T) {
		c := Slice{given}.Cycle()
		seq := Channel{c}.Take(count)
		actual := Channel{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(5, []T{}, []T{})
	f(5, []T{1}, []T{1, 1, 1, 1, 1})
	f(5, []T{1, 2}, []T{1, 2, 1, 2, 1})
}

func TestSliceDedup(t *testing.T) {
	f := func(given []T, expected []T) {
		actual := Slice{given}.Dedup()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{1})
	f([]T{1, 1}, []T{1})
	f([]T{1, 2}, []T{1, 2})
	f([]T{1, 2, 3}, []T{1, 2, 3})
	f([]T{1, 2, 2, 3}, []T{1, 2, 3})
	f([]T{1, 2, 2, 3, 3, 3, 2, 1, 1}, []T{1, 2, 3, 2, 1})
}

func TestSliceDedupBy(t *testing.T) {
	f := func(given []T, expected []T) {
		even := func(el T) G { return G(el % 2) }
		actual := Slice{given}.DedupBy(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{1})
	f([]T{1, 1}, []T{1})
	f([]T{1, 2}, []T{1, 2})
	f([]T{1, 2, 3}, []T{1, 2, 3})
	f([]T{1, 2, 2, 3}, []T{1, 2, 3})
	f([]T{1, 2, 4, 3, 5, 7, 10}, []T{1, 2, 3, 10})
}

func TestSliceDelete(t *testing.T) {
	f := func(given []T, el T, expected []T) {
		actual := Slice{given}.Delete(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 1, []T{})
	f([]T{1}, 1, []T{})
	f([]T{2}, 1, []T{2})
	f([]T{1, 2}, 1, []T{2})
	f([]T{1, 2, 3}, 2, []T{1, 3})
	f([]T{1, 2, 2, 3, 2}, 2, []T{1, 2, 3, 2})
}

func TestSliceDeleteAll(t *testing.T) {
	f := func(given []T, el T, expected []T) {
		actual := Slice{given}.DeleteAll(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 1, []T{})
	f([]T{1}, 1, []T{})
	f([]T{2}, 1, []T{2})
	f([]T{1, 2}, 1, []T{2})
	f([]T{1, 2, 3}, 2, []T{1, 3})
	f([]T{1, 2, 2, 3, 2}, 2, []T{1, 3})
}

func TestSliceDeleteAt(t *testing.T) {
	f := func(given []T, indices []int, expected []T) {
		actual, _ := Slice{given}.DeleteAt(indices...)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []int{}, []T{})
	f([]T{1}, []int{0}, []T{})
	f([]T{1, 2}, []int{0}, []T{2})

	f([]T{1, 2, 3}, []int{0}, []T{2, 3})
	f([]T{1, 2, 3}, []int{1}, []T{1, 3})
	f([]T{1, 2, 3}, []int{2}, []T{1, 2})

	f([]T{1, 2, 3}, []int{0, 1}, []T{3})
	f([]T{1, 2, 3}, []int{0, 2}, []T{2})
	f([]T{1, 2, 3}, []int{1, 2}, []T{1})
}

func TestSliceDropEvery(t *testing.T) {
	f := func(given []T, nth int, from int, expected []T) {
		actual, _ := Slice{given}.DropEvery(nth, from)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 1, 1, []T{})
	f([]T{1, 2, 3}, 1, 1, []T{})

	f([]T{1, 2, 3, 4}, 2, 1, []T{1, 3})
	f([]T{1, 2, 3, 4, 5}, 2, 1, []T{1, 3, 5})

	f([]T{1, 2, 3, 4}, 2, 0, []T{2, 4})
	f([]T{1, 2, 3, 4, 5}, 2, 0, []T{2, 4})

	f([]T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, []T{1, 3, 5, 7, 9})
	f([]T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 1, []T{1, 2, 4, 5, 7, 8, 10})
}

func TestSliceDropWhile(t *testing.T) {
	f := func(given []T, expected []T) {
		even := func(el T) bool { return el%2 == 0 }
		actual := Slice{given}.DropWhile(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{2}, []T{})
	f([]T{1}, []T{1})
	f([]T{2, 1}, []T{1})
	f([]T{2, 1, 2}, []T{1, 2})
	f([]T{1, 2}, []T{1, 2})
	f([]T{2, 4, 6, 1, 8}, []T{1, 8})
}

func TestSliceEndsWith(t *testing.T) {
	f := func(given []T, suffix []T, expected bool) {
		actual := Slice{given}.EndsWith(suffix)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{}, true)
	f([]T{1}, []T{1}, true)
	f([]T{1}, []T{2}, false)
	f([]T{2, 3}, []T{1, 2, 3}, false)

	f([]T{1, 2, 3}, []T{3}, true)
	f([]T{1, 2, 3}, []T{2, 3}, true)
	f([]T{1, 2, 3}, []T{1, 2, 3}, true)

	f([]T{1, 2, 3}, []T{1}, false)
	f([]T{1, 2, 3}, []T{2}, false)
	f([]T{1, 2, 3}, []T{1, 2}, false)
	f([]T{1, 2, 3}, []T{3, 2}, false)
}

func TestSliceEqual(t *testing.T) {
	f := func(left []T, right []T, expected bool) {
		actual := Slice{left}.Equal(right)
		assert.Equal(t, expected, actual, "they should be equal")

		actual = Slice{right}.Equal(left)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{}, true)
	f([]T{1}, []T{1}, true)
	f([]T{1}, []T{2}, false)
	f([]T{1, 2, 3, 3}, []T{1, 2, 3, 3}, true)
	f([]T{1, 2, 3, 3}, []T{1, 2, 2, 3}, false)
	f([]T{1, 2, 3, 3}, []T{1, 2, 4, 3}, false)

	// different len
	f([]T{1, 2, 3}, []T{1, 2}, false)
	f([]T{1, 2}, []T{1, 2, 3}, false)
	f([]T{}, []T{1, 2, 3}, false)
	f([]T{1, 2, 3}, []T{}, false)
}

func TestSliceFilter(t *testing.T) {
	f := func(given []T, expected []T) {
		even := func(t T) bool { return (t % 2) == 0 }
		actual := Slice{given}.Filter(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1, 2, 3, 4}, []T{2, 4})
	f([]T{1, 3}, []T{})
	f([]T{2, 4}, []T{2, 4})
}

func TestSliceFind(t *testing.T) {
	f := func(given []T, expectedEl T, expectedErr error) {
		even := func(t T) bool { return (t % 2) == 0 }
		el, err := Slice{given}.Find(even)
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]T{}, 0, ErrNotFound)
	f([]T{1}, 0, ErrNotFound)
	f([]T{1}, 0, ErrNotFound)
	f([]T{2}, 2, nil)
	f([]T{1, 2}, 2, nil)
	f([]T{1, 2, 3}, 2, nil)
	f([]T{1, 3, 5}, 0, ErrNotFound)
}

func TestSliceFindIndex(t *testing.T) {
	f := func(given []T, expectedInd int, expectedErr error) {
		even := func(t T) bool { return (t % 2) == 0 }
		index, err := Slice{given}.FindIndex(even)
		assert.Equal(t, expectedInd, index, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]T{}, 0, ErrNotFound)
	f([]T{1}, 0, ErrNotFound)
	f([]T{1}, 0, ErrNotFound)
	f([]T{2}, 0, nil)
	f([]T{1, 2}, 1, nil)
	f([]T{1, 2, 3}, 1, nil)
	f([]T{1, 3, 5, 7, 9, 2}, 5, nil)
	f([]T{1, 3, 5}, 0, ErrNotFound)
}

func TestSliceJoin(t *testing.T) {
	f := func(given []T, sep string, expected string) {
		actual := Slice{given}.Join(sep)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, "", "")
	f([]T{}, "|", "")

	f([]T{1}, "", "1")
	f([]T{1}, "|", "1")

	f([]T{1, 2, 3}, "", "123")
	f([]T{1, 2, 3}, "|", "1|2|3")
	f([]T{1, 2, 3}, "<T>", "1<T>2<T>3")
}

func TestSliceGroupBy(t *testing.T) {
	f := func(given []T, expected map[G][]T) {
		reminder := func(t T) G { return G((t % 2)) }
		actual := Slice{given}.GroupBy(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, map[G][]T{})
	f([]T{1}, map[G][]T{1: {1}})
	f([]T{1, 3, 2, 4, 5}, map[G][]T{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceInsertAt(t *testing.T) {
	f := func(given []T, index int, expected []T, expectedErr error) {
		actual, err := Slice{given}.InsertAt(index, 10)
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]T{}, -1, []T{}, ErrNegativeValue)
	f([]T{}, 0, []T{10}, nil)
	f([]T{}, 1, []T{}, ErrOutOfRange)

	f([]T{1, 2, 3}, -1, []T{1, 2, 3}, ErrNegativeValue)
	f([]T{1, 2, 3}, 0, []T{10, 1, 2, 3}, nil)
	f([]T{1, 2, 3}, 1, []T{1, 10, 2, 3}, nil)
	f([]T{1, 2, 3}, 3, []T{1, 2, 3, 10}, nil)
	f([]T{1, 2, 3}, 4, []T{1, 2, 3}, ErrOutOfRange)
}

func TestSliceIntersperse(t *testing.T) {
	f := func(el T, given []T, expected []T) {
		actual := Slice{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []T{}, []T{})
	f(0, []T{1}, []T{1})
	f(0, []T{1, 2}, []T{1, 0, 2})
	f(0, []T{1, 2, 3}, []T{1, 0, 2, 0, 3})
}

func TestSliceLast(t *testing.T) {
	f := func(given []T, expectedEl T, expectedErr error) {
		el, err := Slice{given}.Last()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]T{}, 0, ErrEmpty)
	f([]T{1}, 1, nil)
	f([]T{1, 2, 3}, 3, nil)
}

func TestSliceMap(t *testing.T) {
	f := func(given []T, expected []G) {
		double := func(t T) G { return G((t * 2)) }
		actual := Slice{given}.Map(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []G{})
	f([]T{1}, []G{2})
	f([]T{1, 2, 3}, []G{2, 4, 6})
}

func TestSliceMax(t *testing.T) {
	f := func(given []T, expectedEl T, expectedErr error) {
		el, err := Slice{given}.Max()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]T{}, 0, ErrEmpty)
	f([]T{1}, 1, nil)
	f([]T{1, 2, 3}, 3, nil)
	f([]T{1, 3, 2}, 3, nil)
	f([]T{3, 2, 1}, 3, nil)
}

func TestSliceMin(t *testing.T) {
	f := func(given []T, expectedEl T, expectedErr error) {
		el, err := Slice{given}.Min()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]T{}, 0, ErrEmpty)
	f([]T{1}, 1, nil)
	f([]T{1, 2, 3}, 1, nil)
	f([]T{2, 1, 3}, 1, nil)
	f([]T{3, 2, 1}, 1, nil)
}

func TestSlicesPermutations(t *testing.T) {
	f := func(size int, given []T, expected [][]T) {
		actual := make([][]T, 0)
		i := 0
		s := Slice{given}
		for el := range s.Permutations(size) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []T{}, [][]T{})
	f(2, []T{1}, [][]T{{1}})
	f(2, []T{1, 2, 3}, [][]T{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}})
}

func TestSliceProduct(t *testing.T) {
	f := func(given []T, repeat int, expected [][]T) {
		actual := make([][]T, 0)
		i := 0
		s := Slice{given}
		for el := range s.Product(repeat) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]T{1, 2}, 0, [][]T{})
	f([]T{}, 2, [][]T{})
	f([]T{1}, 2, [][]T{{1, 1}})

	f([]T{1, 2}, 1, [][]T{{1}, {2}})
	f([]T{1, 2}, 2, [][]T{{1, 1}, {1, 2}, {2, 1}, {2, 2}})
	f([]T{1, 2}, 3, [][]T{
		{1, 1, 1}, {1, 1, 2}, {1, 2, 1}, {1, 2, 2},
		{2, 1, 1}, {2, 1, 2}, {2, 2, 1}, {2, 2, 2},
	})
}

func TestSliceReduce(t *testing.T) {
	f := func(given []T, expected G) {
		sum := func(el T, acc G) G { return G(el) + acc }
		actual := Slice{given}.Reduce(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 0)
	f([]T{1}, 1)
	f([]T{1, 2}, 3)
	f([]T{1, 2, 3}, 6)
}

func TestSliceReduceWhile(t *testing.T) {
	f := func(given []T, expected G) {
		sum := func(el T, acc G) (G, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return G(el) + acc, nil
		}
		actual, _ := Slice{given}.ReduceWhile(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 0)
	f([]T{1}, 1)
	f([]T{1, 2}, 3)
	f([]T{1, 2, 3}, 6)
	f([]T{1, 2, 0, 3}, 3)
}

func TestSliceReverse(t *testing.T) {
	f := func(given []T, expected []T) {
		actual := Slice{given}.Reverse()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{1})
	f([]T{1, 2}, []T{2, 1})
	f([]T{1, 2, 3}, []T{3, 2, 1})
	f([]T{1, 2, 2, 3, 3}, []T{3, 3, 2, 2, 1})
}

func TestSliceSame(t *testing.T) {
	f := func(given []T, expected bool) {
		actual := Slice{given}.Same()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, true)
	f([]T{1}, true)
	f([]T{1, 1}, true)
	f([]T{1, 1, 1}, true)

	f([]T{1, 2, 1}, false)
	f([]T{1, 2, 2}, false)
	f([]T{1, 1, 2}, false)
}

func TestSliceScan(t *testing.T) {
	f := func(given []T, expected []G) {
		sum := func(el T, acc G) G { return G(el) + acc }
		actual := Slice{given}.Scan(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []G{})
	f([]T{1}, []G{1})
	f([]T{1, 2}, []G{1, 3})
	f([]T{1, 2, 3}, []G{1, 3, 6})
	f([]T{1, 2, 3, 4}, []G{1, 3, 6, 10})
}
