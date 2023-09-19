package slices_test

import (
	"testing"

	"github.com/life4/genesis/slices"
	"github.com/matryer/is"
)

func TestConcat(t *testing.T) {
	is := is.New(t)
	f := func(given [][]int, expected []int) {
		is.Equal(slices.Concat(given...), expected)
	}
	f([][]int{}, []int{})
	f([][]int{{}}, []int{})
	f([][]int{{1}}, []int{1})
	f([][]int{{1}, {}}, []int{1})
	f([][]int{{}, {1}}, []int{1})
	f([][]int{{1, 2}, {3, 4, 5}}, []int{1, 2, 3, 4, 5})
}

func TestDifference(t *testing.T) {
	is := is.New(t)
	f := func(left, right []int, exp []int) {
		act := slices.Difference(left, right)
		is.Equal(act, exp)
	}
	f([]int{}, []int{}, []int{})
	f([]int{}, []int{4}, []int{})
	f([]int{4}, []int{}, []int{4})
	f([]int{4}, []int{5}, []int{4})
	f([]int{4}, []int{4}, []int{})
	f([]int{4}, []int{4, 4}, []int{})
	f([]int{4, 4}, []int{4}, []int{})
	f([]int{4, 4}, []int{4, 4}, []int{})
	f([]int{4, 4}, []int{3}, []int{4})
	f([]int{4, 5, 3, 1, 2, 1}, []int{1, 5, 5}, []int{4, 3, 2})
}

func TestIntersect(t *testing.T) {
	is := is.New(t)
	f := func(left, right []int, exp []int) {
		act := slices.Intersect(left, right)
		is.Equal(act, exp)
	}
	f([]int{}, []int{}, []int{})
	f([]int{}, []int{4}, []int{})
	f([]int{4}, []int{}, []int{})
	f([]int{4}, []int{5}, []int{})
	f([]int{4}, []int{4}, []int{4})
	f([]int{4}, []int{4, 4}, []int{4})
	f([]int{4, 4}, []int{4}, []int{4})
	f([]int{4, 4}, []int{4, 4}, []int{4})
	f([]int{1, 2, 2, 3, 1}, []int{2, 3, 4, 2, 5}, []int{2, 3})
}

func TestProduct2(t *testing.T) {
	is := is.New(t)
	f := func(given [][]int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		for el := range slices.Product2(given...) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		is.Equal(expected, actual)
	}
	f([][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]int{{1, 2}, {3}, {4, 5}}, [][]int{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}

func TestUnion(t *testing.T) {
	is := is.New(t)
	f := func(left, right []int, exp []int) {
		act := slices.Union(left, right)
		is.Equal(act, exp)
	}
	f([]int{}, []int{}, []int{})
	f([]int{}, []int{4}, []int{4})
	f([]int{4}, []int{}, []int{4})
	f([]int{4}, []int{5}, []int{4, 5})
	f([]int{4}, []int{4}, []int{4})
	f([]int{5}, []int{4}, []int{5, 4})
	f([]int{1, 2, 2, 3, 1}, []int{2, 3, 4, 2, 5}, []int{1, 2, 3, 4, 5})
}

func TestZip(t *testing.T) {
	is := is.New(t)
	f := func(given [][]int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		for el := range slices.Zip(given...) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		is.Equal(expected, actual)
	}
	f([][]int{}, [][]int{})
	f([][]int{{1}, {2}}, [][]int{{1, 2}})
	f([][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {2, 4}})
	f([][]int{{1, 2, 3}, {4, 5}}, [][]int{{1, 4}, {2, 5}})
}
