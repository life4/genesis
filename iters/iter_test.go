package iters_test

import (
	"iter"
	"testing"

	"github.com/life4/genesis/iters"
	"github.com/life4/genesis/lambdas"
	"github.com/matryer/is"
)

func new[T any](items ...T) iter.Seq[T] {
	return iters.FromSlice(items)
}

func TestIter_Count(t *testing.T) {
	is := is.New(t)
	is.Equal(iters.Count(new[int]()), 0)
	is.Equal(iters.Count(new(4, 5, 6)), 3)
	is.Equal(iters.Count(new(4)), 1)
}

func TestIter_First(t *testing.T) {
	is := is.New(t)
	_, err := iters.First(new[int]())
	is.True(err != nil)
	is.Equal(lambdas.Must(iters.First(new(9))), 9)
	is.Equal(lambdas.Must(iters.First(new(4, 5, 6))), 4)

	// Unlike the classic Next, First is re-entrant (more like Peek).
	it := new(9)
	is.Equal(lambdas.Must(iters.First(it)), 9)
	is.Equal(lambdas.Must(iters.First(it)), 9)
}

func TestIter_Last(t *testing.T) {
	is := is.New(t)
	_, err := iters.Last(new[int]())
	is.True(err != nil)
	is.Equal(lambdas.Must(iters.Last(new(9))), 9)
	is.Equal(lambdas.Must(iters.Last(new(4, 5, 6))), 6)
}

func TestIter_Nth(t *testing.T) {
	is := is.New(t)
	_, err := iters.Nth(new[int](), 0)
	is.True(err != nil)
	_, err = iters.Nth(new[int](), 1)
	is.True(err != nil)

	is.Equal(lambdas.Must(iters.Nth(new(9), 0)), 9)
	_, err = iters.Nth(new(9), 1)
	is.True(err != nil)

	is.Equal(lambdas.Must(iters.Nth(new(4, 5, 6), 0)), 4)
	is.Equal(lambdas.Must(iters.Nth(new(4, 5, 6), 2)), 6)
	is.Equal(lambdas.Must(iters.Nth(new(4, 5, 6), 1)), 5)
	_, err = iters.Nth(new(4, 5, 6), 3)
	is.True(err != nil)
}

func TestIter_StepBy(t *testing.T) {
	is := is.New(t)
	is.Equal(len(iters.ToSlice(iters.StepBy(new[int](), 1))), 0)
	is.Equal(iters.ToSlice(iters.StepBy(new(4, 5, 6, 7, 8), 1)), []int{4, 5, 6, 7, 8})
	is.Equal(iters.ToSlice(iters.StepBy(new(4, 5, 6, 7, 8), 2)), []int{4, 6, 8})
	is.Equal(iters.ToSlice(iters.StepBy(new(4, 5, 6, 7, 8), 3)), []int{4, 7})
}

func TestIter_Chain(t *testing.T) {
	is := is.New(t)
	is.Equal(len(iters.ToSlice(iters.Chain(new[int](), new[int]()))), 0)
	is.Equal(iters.ToSlice(iters.Chain(new(4, 5), new(6, 7))), []int{4, 5, 6, 7})
}

func TestIter_Intersperse(t *testing.T) {
	is := is.New(t)
	is.Equal(len(iters.ToSlice(iters.Intersperse(new[int](), 4))), 0)
	is.Equal(iters.ToSlice(iters.Intersperse(new(4, 5, 6), 8)), []int{4, 8, 5, 8, 6})
}

func TestIter_Map(t *testing.T) {
	is := is.New(t)
	double := func(v int) int { return v * 2 }
	is.Equal(iters.ToSlice(iters.Map(new(4, 5, 6), double)), []int{8, 10, 12})
}

func TestIter_Each(t *testing.T) {
	is := is.New(t)
	s := 0
	f := func(v int) { s += v }
	iters.Each(new(4, 5, 6), f)
	is.Equal(s, 15)
}

func TestIter_Filter(t *testing.T) {
	is := is.New(t)
	even := func(v int) bool { return v%2 == 0 }
	is.Equal(iters.ToSlice(iters.Filter(new(4, 5, 6, 7, 8), even)), []int{4, 6, 8})
}
