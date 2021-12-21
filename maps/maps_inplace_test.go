package maps_test

import (
	"testing"

	"github.com/life4/genesis/maps"
	"github.com/matryer/is"
)

func TestClear(t *testing.T) {
	is := is.New(t)
	m := map[int]int{1: 2, 3: 4}
	maps.Clear(m)
	is.Equal(m, map[int]int{})
}

func TestDrop(t *testing.T) {
	is := is.New(t)
	m := map[int]int{1: 2, 3: 4, 5: 6}
	maps.Drop(m, 1, 5)
	is.Equal(m, map[int]int{3: 4})
}

func TestLeaveOnly(t *testing.T) {
	is := is.New(t)
	m := map[int]int{1: 2, 3: 4, 5: 6}
	maps.LeaveOnly(m, 1, 5)
	is.Equal(m, map[int]int{1: 2, 5: 6})
}

func TestPop(t *testing.T) {
	is := is.New(t)
	m := map[int]int{1: 2, 3: 4, 5: 6}
	val, err := maps.Pop(m, 1)
	is.Equal(err, nil)
	is.Equal(val, 2)
	is.Equal(m, map[int]int{3: 4, 5: 6})
}

func TestReplace(t *testing.T) {
	is := is.New(t)
	m := map[int]int{1: 2, 5: 6}
	maps.Replace(m, 5, 7)
	maps.Replace(m, 8, 9)
	is.Equal(m, map[int]int{1: 2, 5: 7})
}

func TestUpdate(t *testing.T) {
	is := is.New(t)
	m1 := map[int]int{1: 2, 3: 4}
	m2 := map[int]int{3: 5, 6: 7}
	maps.Update(m1, m2)
	is.Equal(m1, map[int]int{1: 2, 3: 5, 6: 7})
}
