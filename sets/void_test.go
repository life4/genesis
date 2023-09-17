package sets_test

import (
	"testing"

	"github.com/life4/genesis/sets"
	"github.com/matryer/is"
)

func TestAdd(t *testing.T) {
	is := is.NewRelaxed(t)
	f := func(given map[int]sets.Z, val int, exp map[int]sets.Z) {
		sets.Add(given, val)
		is.Equal(given, exp)
	}
	f(sets.New(3, 4, 5), 3, sets.New(3, 4, 5))
	f(sets.New(3, 4, 5), 6, sets.New(3, 4, 5, 6))
	f(sets.New[int](), 6, sets.New(6))
	// f(nilSet, 6, sets.Set(6))
}

func TestClear(t *testing.T) {
	is := is.New(t)
	s := sets.New(3, 4, 5)
	sets.Clear(s)
	is.Equal(s, sets.New[int]())

	s = sets.New[int]()
	sets.Clear(s)
	is.Equal(s, sets.New[int]())

	s = nil
	sets.Clear(s)
	is.Equal(s, nil)
}

func TestDiscard(t *testing.T) {
	is := is.NewRelaxed(t)
	f := func(given map[int]sets.Z, val int, exp map[int]sets.Z) {
		sets.Discard(given, val)
		is.Equal(given, exp)
	}
	f(sets.New(3, 4, 5), 3, sets.New(4, 5))
	f(sets.New(3, 4, 5), 6, sets.New(3, 4, 5))
	f(sets.New[int](), 6, sets.New[int]())
	f(nilSet, 6, nilSet)
}

func TestPop(t *testing.T) {
	is := is.New(t)
	s := sets.New(3)
	val, err := sets.Pop(s)
	is.NoErr(err)
	is.Equal(val, 3)

	val, err = sets.Pop(s)
	is.Equal(err, sets.ErrEmpty)
	is.Equal(val, 0)

	s = nil
	val, err = sets.Pop(s)
	is.Equal(err, sets.ErrEmpty)
	is.Equal(val, 0)
}

func TestUpdate(t *testing.T) {
	is := is.NewRelaxed(t)
	f := func(target, values, exp map[int]sets.Z) {
		sets.Update(target, values)
		is.Equal(target, exp)
	}
	f(sets.New(3, 4, 5), sets.New(4, 5, 6), sets.New(3, 4, 5, 6))
	f(sets.New(3, 4, 5), sets.New(4), sets.New(3, 4, 5))
	f(sets.New[int](), sets.New(4), sets.New(4))
	f(sets.New[int](), sets.New[int](), sets.New[int]())
	f(sets.New[int](4), sets.New[int](), sets.New(4))
	f(nil, sets.New[int](), nil)
	f(nil, nil, nil)
	f(sets.New[int](4), nil, sets.New[int](4))
	// f(nil, sets.New[int](4), nil)
}
