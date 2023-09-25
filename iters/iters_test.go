package iters_test

import (
	"testing"

	"github.com/life4/genesis/iters"
	"github.com/matryer/is"
)

var ts = iters.ToSlice[int]

func new[T any](vs ...T) iters.Iter[T] {
	return iters.FromSlice(vs)
}

func TestFilter(t *testing.T) {
	is := is.NewRelaxed(t)
	even := func(x int) bool { return x%2 == 0 }
	is.Equal(ts(iters.Filter(new(3, 4, 5, 6), even)), []int{4, 6})
	is.Equal(ts(iters.Filter(new(4, 6), even)), []int{4, 6})
	is.Equal(ts(iters.Filter(new(3, 5), even)), []int{})
	is.Equal(ts(iters.Filter(new[int](), even)), []int{})
}

func TestFromChannel(t *testing.T) {
	is := is.New(t)
	ch := make(chan int)
	go func() {
		ch <- 3
		ch <- 4
		ch <- 5
		close(ch)
	}()
	is.Equal(ts(iters.FromChannel(ch)), []int{3, 4, 5})
}

func TestFromFunc(t *testing.T) {
	is := is.New(t)
	c := 0
	f := func() (int, bool) {
		c += 1
		if c <= 3 {
			return c + 2, true
		}
		return 0, false
	}
	is.Equal(ts(iters.FromFunc(f)), []int{3, 4, 5})
}

func TestFromSlice(t *testing.T) {
	is := is.NewRelaxed(t)
	is.Equal(ts(iters.FromSlice([]int{3, 4, 5})), []int{3, 4, 5})
	is.Equal(ts(iters.FromSlice([]int{3})), []int{3})
	is.Equal(ts(iters.FromSlice([]int{})), []int{})
	is.Equal(ts(iters.FromSlice([]int(nil))), []int{})
}

func TestMap(t *testing.T) {
	is := is.NewRelaxed(t)
	double := func(x int) int { return x * 2 }
	is.Equal(ts(iters.Map(new(3, 4, 5), double)), []int{6, 8, 10})
	is.Equal(ts(iters.Map(new[int](), double)), []int{})
}

func TestReduce(t *testing.T) {
	is := is.NewRelaxed(t)
	add := func(x, a int) int { return x + a }
	is.Equal(iters.Reduce(new(3, 4, 5), 0, add), 12)
	is.Equal(iters.Reduce(new(3, 4, 5), 3, add), 15)
	is.Equal(iters.Reduce(new[int](), 0, add), 0)
	is.Equal(iters.Reduce(new[int](), 7, add), 7)
}

func TestTake(t *testing.T) {
	is := is.NewRelaxed(t)
	is.Equal(ts(iters.Take(new(3, 4, 5), 2)), []int{3, 4})
	is.Equal(ts(iters.Take(new(3, 4, 5), 1)), []int{3})
	is.Equal(ts(iters.Take(new(3, 4, 5), 10)), []int{3, 4, 5})
	is.Equal(ts(iters.Take(new(3, 4, 5), 0)), []int{})
	is.Equal(ts(iters.Take(new(3, 4, 5), -1)), []int{})
	is.Equal(ts(iters.Take(new(3, 4, 5), -10)), []int{})
}

func TestToSlice(t *testing.T) {
	is := is.NewRelaxed(t)
	is.Equal(iters.ToSlice(new(3, 4, 5)), []int{3, 4, 5})
	is.Equal(iters.ToSlice(new[int]()), []int{})
}
