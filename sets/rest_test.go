package sets_test

import (
	"testing"

	"github.com/life4/genesis/sets"
	"github.com/matryer/is"
)

func TestMax(t *testing.T) {
	is := is.NewRelaxed(t)

	act, err := sets.Max(sets.New(3, 4, 5))
	is.NoErr(err)
	is.Equal(act, 5)

	act, err = sets.Max(sets.New(-3, -4, -5))
	is.NoErr(err)
	is.Equal(act, -3)

	act, err = sets.Max(sets.New[int]())
	is.Equal(err, sets.ErrEmpty)
	is.Equal(act, 0)

	act, err = sets.Max(nilSet)
	is.Equal(err, sets.ErrEmpty)
	is.Equal(act, 0)
}

func TestMin(t *testing.T) {
	is := is.NewRelaxed(t)

	act, err := sets.Min(sets.New(3, 4, 5))
	is.NoErr(err)
	is.Equal(act, 3)

	act, err = sets.Min(sets.New(-3, -4, -5))
	is.NoErr(err)
	is.Equal(act, -5)

	act, err = sets.Min(sets.New[int]())
	is.Equal(err, sets.ErrEmpty)
	is.Equal(act, 0)

	act, err = sets.Min(nilSet)
	is.Equal(err, sets.ErrEmpty)
	is.Equal(act, 0)
}

func TestReduce(t *testing.T) {
	is := is.NewRelaxed(t)
	add := func(x, y int) int { return x + y }
	is.Equal(sets.Reduce(sets.New(3, 4, 5), 0, add), 12)
	is.Equal(sets.Reduce(sets.New(3, 4, 5), -10, add), 2)
	is.Equal(sets.Reduce(sets.New(-3, -4, -5), 0, add), -12)
	is.Equal(sets.Reduce(sets.New[int](), 0, add), 0)
	is.Equal(sets.Reduce(nilSet, 0, add), 0)
}

func TestSum(t *testing.T) {
	is := is.NewRelaxed(t)
	is.Equal(sets.Sum(sets.New(3, 4, 5)), 12)
	is.Equal(sets.Sum(sets.New(-3, -4, -5)), -12)
	is.Equal(sets.Sum(sets.New[int]()), 0)
	is.Equal(sets.Sum(nilSet), 0)
}

func TestToSlice(t *testing.T) {
	is := is.NewRelaxed(t)
	s := sets.ToSlice(sets.New(3, 4, 5))
	is.Equal(sets.FromSlice(s), sets.New(3, 4, 5))
	is.Equal(sets.ToSlice(sets.New[int]()), []int{})
	is.Equal(sets.ToSlice(nilSet), nil)
}
