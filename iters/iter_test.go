package iters_test

import (
	"iter"
	"testing"

	"github.com/life4/genesis/iters"
	"github.com/life4/genesis/lambdas"
	"github.com/matryer/is"
)

var ts = iters.ToSlice[int]

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
	is.Equal(len(ts(iters.StepBy(new[int](), 1))), 0)
	is.Equal(ts(iters.StepBy(new(4, 5, 6, 7, 8), 1)), []int{4, 5, 6, 7, 8})
	is.Equal(ts(iters.StepBy(new(4, 5, 6, 7, 8), 2)), []int{4, 6, 8})
	is.Equal(ts(iters.StepBy(new(4, 5, 6, 7, 8), 3)), []int{4, 7})
}

func TestIter_Chain(t *testing.T) {
	is := is.New(t)
	is.Equal(len(ts(iters.Chain(new[int](), new[int]()))), 0)
	is.Equal(ts(iters.Chain(new(4, 5), new(6, 7))), []int{4, 5, 6, 7})
}

func TestIter_Intersperse(t *testing.T) {
	is := is.New(t)
	is.Equal(len(ts(iters.Intersperse(new[int](), 4))), 0)
	is.Equal(ts(iters.Intersperse(new(4, 5, 6), 8)), []int{4, 8, 5, 8, 6})
}

func TestIter_Map(t *testing.T) {
	is := is.New(t)
	double := func(v int) int { return v * 2 }
	is.Equal(ts(iters.Map(new(4, 5, 6), double)), []int{8, 10, 12})
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
	is.Equal(ts(iters.Filter(new(4, 5, 6, 7, 8), even)), []int{4, 6, 8})
}

func TestAll(t *testing.T) {
	is := is.New(t)
	even := func(v int) bool { return v%2 == 0 }
	is.True(iters.All(new[int](), even))
	is.True(iters.All(new(2), even))
	is.True(iters.All(new(2, 4, 6), even))

	is.True(!iters.All(new(3), even))
	is.True(!iters.All(new(3, 5), even))
	is.True(!iters.All(new(2, 3), even))
	is.True(!iters.All(new(2, 3, 6), even))
	is.True(!iters.All(new(2, 4, 3, 6), even))
	is.True(!iters.All(new(2, 4, 3), even))
}

func TestAny(t *testing.T) {
	is := is.New(t)
	even := func(v int) bool { return v%2 == 0 }
	is.True(iters.Any(new(2), even))
	is.True(iters.Any(new(2, 4, 6), even))
	is.True(iters.Any(new(2, 3), even))
	is.True(iters.Any(new(2, 3), even))
	is.True(iters.Any(new(3, 5, 6), even))
	is.True(iters.Any(new(3, 4, 5), even))

	is.True(!iters.Any(new[int](), even))
	is.True(!iters.Any(new(3), even))
	is.True(!iters.Any(new(3, 5), even))
}

func TestContains(t *testing.T) {
	is := is.New(t)
	is.True(iters.Contains(new(2), 2))
	is.True(iters.Contains(new(2, 4, 6), 2))
	is.True(iters.Contains(new(2, 4, 6), 4))
	is.True(iters.Contains(new(2, 4, 6), 6))

	is.True(!iters.Contains(new[int](), 2))
	is.True(!iters.Contains(new(2, 4, 6), 3))
	is.True(!iters.Contains(new(2, 4, 6), 5))
}

func TestEq(t *testing.T) {
	is := is.New(t)
	is.True(iters.Eq(new[int](), new[int]()))
	is.True(iters.Eq(new(2), new(2)))
	is.True(iters.Eq(new(2, 4, 6), new(2, 4, 6)))

	is.True(!iters.Eq(new(1), new(2)))
	is.True(!iters.Eq(new(2, 4), new(2, 4, 6)))
	is.True(!iters.Eq(new(2, 4, 6), new(2, 4)))
	is.True(!iters.Eq(new(2, 4, 6), new(2, 6, 4)))
}

func TestSkip(t *testing.T) {
	is := is.NewRelaxed(t)
	is.Equal(ts(iters.Skip(new(3, 4, 5), 1)), []int{4, 5})
	is.Equal(ts(iters.Skip(new(3, 4, 5), 2)), []int{5})
	is.Equal(ts(iters.Skip(new(3, 4, 5), 0)), []int{3, 4, 5})
	is.Equal(ts(iters.Skip(new(3, 4, 5), -5)), []int{3, 4, 5})
	is.Equal(ts(iters.Skip(new(3, 4, 5), 3)), []int{})
	is.Equal(ts(iters.Skip(new(3, 4, 5), 6)), []int{})
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

func TestFromSlice(t *testing.T) {
	is := is.NewRelaxed(t)
	is.Equal(ts(iters.FromSlice([]int{3, 4, 5})), []int{3, 4, 5})
	is.Equal(ts(iters.FromSlice([]int{3})), []int{3})
	is.Equal(ts(iters.FromSlice([]int{})), []int{})
	is.Equal(ts(iters.FromSlice([]int(nil))), []int{})
}

func TestFold(t *testing.T) {
	is := is.NewRelaxed(t)
	add := func(x, a int) int { return x + a }
	is.Equal(iters.Fold(new(3, 4, 5), 0, add), 12)
	is.Equal(iters.Fold(new(3, 4, 5), 3, add), 15)
	is.Equal(iters.Fold(new[int](), 0, add), 0)
	is.Equal(iters.Fold(new[int](), 7, add), 7)
}

func TestReduce(t *testing.T) {
	is := is.NewRelaxed(t)
	add := func(x, a int) int { return x + a }
	is.Equal(lambdas.Must(iters.Reduce(new(3, 4, 5), add)), 12)
	_, err := iters.Reduce(new[int](), add)
	is.True(err != nil)
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

func TestSum(t *testing.T) {
	is := is.NewRelaxed(t)
	is.Equal(iters.Sum(new[int]()), 0)
	is.Equal(iters.Sum(new(3, 4, 5)), 12)
	is.Equal(iters.Sum(new("ab", "cd")), "abcd")
}

func TestSorted(t *testing.T) {
	is := is.New(t)
	is.True(iters.Sorted(new[int]()))
	is.True(iters.Sorted(new(4)))
	is.True(iters.Sorted(new(4, 5, 6)))
	is.True(iters.Sorted(new(4, 6)))

	is.True(!iters.Sorted(new(4, 6, 5)))
	is.True(!iters.Sorted(new(4, 3)))
}
