package maps_test

import (
	"testing"

	"github.com/life4/genesis/maps"
	"github.com/matryer/is"
)

func TestCopy(t *testing.T) {
	is := is.New(t)
	m1 := map[int]int{1: 2}
	m2 := maps.Copy(m1)
	m2[3] = 4
	is.Equal(m1, map[int]int{1: 2})
}

func TestEqual(t *testing.T) {
	is := is.New(t)
	is.True(maps.Equal(map[int]int{}, map[int]int{}))
	is.True(maps.Equal(map[int]int{1: 2}, map[int]int{1: 2}))
	is.True(maps.Equal(map[int]int{1: 2, 3: 4}, map[int]int{1: 2, 3: 4}))

	is.True(!maps.Equal(map[int]int{1: 2, 3: 4}, map[int]int{1: 2, 3: 4, 5: 6}))
	is.True(!maps.Equal(map[int]int{1: 2, 3: 4}, map[int]int{}))
	is.True(!maps.Equal(map[int]int{1: 2, 3: 4}, map[int]int{1: 2, 3: 5}))
	is.True(!maps.Equal(map[int]int{1: 2, 3: 4}, map[int]int{6: 7}))
}

func TestHasKey(t *testing.T) {
	is := is.New(t)
	m := map[int]int{1: 2, 3: 4}
	is.True(maps.HasKey(m, 1))
	is.True(maps.HasKey(m, 3))

	is.True(!maps.HasKey(m, 2))
	is.True(!maps.HasKey(m, 4))
	is.True(!maps.HasKey(m, 5))
}

func TestWithout(t *testing.T) {
	is := is.New(t)
	m := map[int]int{1: 2, 3: 4, 5: 6}
	is.Equal(maps.Without(m, 1, 5), map[int]int{3: 4})
}
