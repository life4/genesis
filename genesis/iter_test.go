package genesis_test

import (
	"testing"

	"github.com/life4/genesis/genesis"
	"github.com/matryer/is"
)

func iter[T any](items ...T) genesis.Iter[T] {
	return genesis.NewSlice(items).Iter()
}

func TestIter_Count(t *testing.T) {
	is := is.New(t)
	is.Equal(iter[int]().Count(), 0)
	is.Equal(iter(4, 5, 6).Count(), 3)
	is.Equal(iter(4).Count(), 1)
}

func TestIter_First(t *testing.T) {
	is := is.New(t)
	is.True(iter[int]().First().IsNone())
	is.Equal(iter(9).First().Unwrap(), 9)
	is.Equal(iter(4, 5, 6).First().Unwrap(), 4)

	// Unlike the classic Next, First is re-entrant (more like Peek).
	it := iter(9)
	is.Equal(it.First().Unwrap(), 9)
	is.Equal(it.First().Unwrap(), 9)
}

func TestIter_Last(t *testing.T) {
	is := is.New(t)
	is.True(iter[int]().Last().IsNone())
	is.Equal(iter(9).Last().Unwrap(), 9)
	is.Equal(iter(4, 5, 6).Last().Unwrap(), 6)
}

func TestIter_Nth(t *testing.T) {
	is := is.New(t)
	is.True(iter[int]().Nth(0).IsNone())
	is.True(iter[int]().Nth(1).IsNone())

	is.Equal(iter(9).Nth(0).Unwrap(), 9)
	is.True(iter(9).Nth(1).IsNone())

	is.Equal(iter(4, 5, 6).Nth(0).Unwrap(), 4)
	is.Equal(iter(4, 5, 6).Nth(2).Unwrap(), 6)
	is.Equal(iter(4, 5, 6).Nth(1).Unwrap(), 5)
	is.True(iter(4, 5, 6).Nth(3).IsNone())
}

func TestIter_StepBy(t *testing.T) {
	is := is.New(t)
	is.Equal(len(iter[int]().StepBy(1).Slice()), 0)
	is.Equal(iter(4, 5, 6, 7, 8).StepBy(1).Slice(), genesis.NewSlice([]int{4, 5, 6, 7, 8}))
	is.Equal(iter(4, 5, 6, 7, 8).StepBy(2).Slice(), genesis.NewSlice([]int{4, 6, 8}))
	is.Equal(iter(4, 5, 6, 7, 8).StepBy(3).Slice(), genesis.NewSlice([]int{4, 7}))
}
