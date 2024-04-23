package slices_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/life4/genesis/slices"
	"github.com/matryer/is"
)

func TestAll(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected bool) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := slices.All(given, even)
		is.Equal(actual, expected)
	}
	f([]int{}, true)
	f([]int{2}, true)
	f([]int{1}, false)
	f([]int{2, 4}, true)
	f([]int{2, 4, 1}, false)
	f([]int{1, 2, 4}, false)
}

func TestAny(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected bool) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := slices.Any(given, even)
		is.Equal(actual, expected)
	}
	f([]int{}, false)
	f([]int{1, 3}, false)
	f([]int{2}, true)
	f([]int{1, 2}, true)
}

func TestChunkBy(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected [][]int) {
		reminder := func(t int) int { return (t % 2) }
		actual := slices.ChunkBy(given, reminder)
		is.Equal(actual, expected)
	}
	f([]int{}, [][]int{})
	f([]int{1}, [][]int{{1}})
	f([]int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f([]int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestCountBy(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected int) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := slices.CountBy(given, even)
		is.Equal(actual, expected)
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
		is.Equal(actual, expected)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 4, 3, 5, 7, 10}, []int{1, 2, 3, 10})

	// DedupBy supports non-comparable types.
	g := func(given [][]int, expected [][]int) {
		first := func(s []int) int { return s[0] }
		actual := slices.DedupBy(given, first)
		is.Equal(actual, expected)
	}
	g([][]int{{1}, {1}}, [][]int{{1}})
	g([][]int{{1}, {1, 2}}, [][]int{{1}})
}

func TestDropWhile(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		even := func(el int) bool { return el%2 == 0 }
		actual := slices.DropWhile(given, even)
		is.Equal(actual, expected)
	}
	f([]int{}, []int{})
	f([]int{2}, []int{})
	f([]int{1}, []int{1})
	f([]int{2, 1}, []int{1})
	f([]int{2, 1, 2}, []int{1, 2})
	f([]int{1, 2}, []int{1, 2})
	f([]int{2, 4, 6, 1, 8}, []int{1, 8})
}

func TestEach(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		actual := make([]int, 0)
		double := func(t int) {
			actual = append(actual, t*2)
		}
		slices.Each(given, double)
		is.Equal(actual, expected)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{2})
	f([]int{1, 2, 3}, []int{2, 4, 6})
	f([]int{2, 5, 3}, []int{4, 10, 6})
}

func TestEachErr(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		actual := make([]int, 0)
		double := func(t int) error {
			if t == 13 {
				return errors.New("oh no")
			}
			actual = append(actual, t*2)
			return nil
		}
		_ = slices.EachErr(given, double)
		is.Equal(actual, expected)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{2})
	f([]int{1, 2, 3}, []int{2, 4, 6})
	f([]int{2, 5, 3}, []int{4, 10, 6})
	f([]int{2, 5, 3, 13, 4}, []int{4, 10, 6})
}

func TestEqualBy(t *testing.T) {
	is := is.New(t)
	f := func(left []int, right []int, expected bool) {
		f := func(a, b int) bool { return a == -b }
		actual := slices.EqualBy(left, right, f)
		is.Equal(actual, expected)

		actual = slices.EqualBy(right, left, f)
		is.Equal(actual, expected)
	}
	f([]int{}, []int{}, true)
	f([]int{1}, []int{-1}, true)
	f([]int{1}, []int{1}, false)
	f([]int{1}, []int{2}, false)
	f([]int{1, 2, 3, 3}, []int{-1, -2, -3, -3}, true)
	f([]int{1, 2, 3, 3}, []int{-1, -2, -3, 3}, false)
	f([]int{1, 2, 3, 3}, []int{-1, -2, -2, -3}, false)
	f([]int{1, 2, 3, 3}, []int{-1, -2, -4, -3}, false)

	// different len
	f([]int{1, 2, 3}, []int{-1, -2}, false)
	f([]int{1, 2}, []int{-1, -2, 3}, false)
	f([]int{}, []int{-1, -2, -3}, false)
	f([]int{1, 2, 3}, []int{}, false)
}

func TestFilter(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := slices.Filter(given, even)
		is.Equal(actual, expected)
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
		is.Equal(actual, expected)
	}
	f([]int{}, map[int][]int{})
	f([]int{1}, map[int][]int{1: {1}})
	f([]int{1, 3, 2, 4, 5}, map[int][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupBy_2(t *testing.T) {
	is := is.New(t)

	type employee struct {
		Name       string
		Department string
	}

	f := func(given []employee, expected map[string][]employee) {
		actual := slices.GroupBy(given, func(emp employee) string {
			return emp.Department
		})
		is.Equal(actual, expected)
	}

	employeeJohn := employee{Name: "John", Department: "Engineering"}
	employeeEva := employee{Name: "Eva", Department: "HR"}
	employeeCarlos := employee{Name: "Carlos", Department: "Engineering"}

	f([]employee{employeeJohn, employeeEva, employeeCarlos}, map[string][]employee{
		"Engineering": {employeeJohn, employeeCarlos},
		"HR":          {employeeEva},
	})

	f([]employee{}, map[string][]employee{})

	f(nil, map[string][]employee{})
}

func TestIndexBy(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expectedInd int) {
		even := func(t int) bool { return (t % 2) == 0 }
		index, _ := slices.IndexBy(given, even)
		is.Equal(expectedInd, index)
	}
	f([]int{}, 0)
	f([]int{6}, 0)
	f([]int{3, 6}, 1)
	f([]int{3, 6, 8}, 1)
	f([]int{3, 5, 8}, 2)
}

func TestMap(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		double := func(t int) int { return (t * 2) }
		actual := slices.Map(given, double)
		is.Equal(actual, expected)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{2})
	f([]int{1, 2, 3}, []int{2, 4, 6})
}

func TestMapFilter(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []string) {
		isEven := func(t int) (string, bool) {
			if t%2 == 0 {
				s := strconv.Itoa(t)
				return s, true
			} else {
				return "", false
			}

		}
		actual := slices.MapFilter(given, isEven)
		is.Equal(actual, expected)
	}
	f([]int{}, []string{})
	f([]int{1}, []string{})
	f([]int{1, 2, 3, 4}, []string{"2", "4"})
}

func TestPartition(t *testing.T) {
	is := is.New(t)
	f := func(given, exp1, exp2 []int) {
		even := func(x int) bool { return x%2 == 0 }
		act1, act2 := slices.Partition(given, even)
		is.Equal(act1, exp1)
		is.Equal(act2, exp2)
		is.Equal(len(act1)+len(act2), len(given))
	}
	f([]int{}, []int{}, []int{})
	f([]int{1}, []int{}, []int{1})
	f([]int{2}, []int{2}, []int{})
	f([]int{1, 2, 3, 4}, []int{2, 4}, []int{1, 3})
	f([]int{2, 1, 4, 3}, []int{2, 4}, []int{1, 3})
	f([]int{4, 3, 2, 1}, []int{4, 2}, []int{3, 1})
}

func TestReduce(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected int) {
		sum := func(el int, acc int) int { return (el) + acc }
		actual := slices.Reduce(given, 0, sum)
		is.Equal(actual, expected)
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
		is.Equal(actual, expected)
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
	f([]int{1, 2, 0, 3}, 3)
}

func TestReject(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		odd := func(t int) bool { return (t % 2) == 1 }
		actual := slices.Reject(given, odd)
		is.Equal(actual, expected)
	}
	f([]int{}, []int{})
	f([]int{1, 2, 3, 4}, []int{2, 4})
	f([]int{1, 3}, []int{})
	f([]int{2, 4}, []int{2, 4})
}

func TestScan(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		sum := func(el int, acc int) int { return (el) + acc }
		actual := slices.Scan(given, 0, sum)
		is.Equal(actual, expected)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 2}, []int{1, 3})
	f([]int{1, 2, 3}, []int{1, 3, 6})
	f([]int{1, 2, 3, 4}, []int{1, 3, 6, 10})
}

func TestSortBy(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		abs := func(x int) int {
			if x < 0 {
				return -x
			}
			return x
		}
		actual := slices.SortBy(given, abs)
		is.Equal(actual, expected)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, -2}, []int{1, -2})
	f([]int{-1, 2}, []int{-1, 2})
	f([]int{-1, -2}, []int{-1, -2})
	f([]int{2, -1}, []int{-1, 2})
	f([]int{-2, 1}, []int{1, -2})
	f([]int{-2, -1}, []int{-1, -2})
	f([]int{2, 3, 1}, []int{1, 2, 3})
	f([]int{2, -3, 1}, []int{1, 2, -3})
	f([]int{2, -3, 3, 1}, []int{1, 2, -3, 3})
	f([]int{2, 3, -3, 1}, []int{1, 2, 3, -3})
	f([]int{2, 3, -1, -2}, []int{-1, 2, -2, 3})
	f([]int{-2, -3, -1, -2}, []int{-1, -2, -2, -3})
}

func TestTakeWhile(t *testing.T) {
	is := is.New(t)
	f := func(given []int, expected []int) {
		even := func(el int) bool { return el%2 == 0 }
		actual := slices.TakeWhile(given, even)
		is.Equal(actual, expected)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{})
	f([]int{2}, []int{2})
	f([]int{2, 4, 6, 1, 8}, []int{2, 4, 6})
	f([]int{1, 2, 3}, []int{})
}
