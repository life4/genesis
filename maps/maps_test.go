package maps_test

import (
	"fmt"
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

	is.Equal(maps.Copy[map[int]int](nil), nil)
}

func TestEqual(t *testing.T) {
	is := is.New(t)
	is.True(maps.Equal(map[int]int{}, map[int]int{}))
	is.True(maps.Equal(map[int]int{1: 2}, map[int]int{1: 2}))
	is.True(maps.Equal(map[int]int{1: 2, 3: 4}, map[int]int{1: 2, 3: 4}))
	is.True(maps.Equal[map[int]int, map[int]int](nil, nil))

	is.True(!maps.Equal(map[int]int{1: 2, 3: 4}, map[int]int{1: 2, 3: 4, 5: 6}))
	is.True(!maps.Equal(map[int]int{1: 2, 3: 4}, map[int]int{}))
	is.True(!maps.Equal(map[int]int{1: 2, 3: 4}, map[int]int{1: 2, 3: 5}))
	is.True(!maps.Equal(map[int]int{1: 2, 3: 4}, map[int]int{6: 7}))
	is.True(!maps.Equal(map[int]int{1: 2}, map[int]int{3: 4}))
	is.True(!maps.Equal(map[int]int{1: 2}, map[int]int{1: 3}))
	is.True(!maps.Equal(map[int]int{1: 2}, map[int]int{2: 2}))
	is.True(!maps.Equal(map[int]int{1: 2}, map[int]int{2: 1}))
}

func TestHasKey(t *testing.T) {
	is := is.New(t)
	m := map[int]int{1: 2, 3: 4}
	is.True(maps.HasKey(m, 1))
	is.True(maps.HasKey(m, 3))

	is.True(!maps.HasKey(m, 2))
	is.True(!maps.HasKey(m, 4))
	is.True(!maps.HasKey(m, 5))
	is.True(!maps.HasKey[map[int]int](nil, 3))
}

func TestFromKeys(t *testing.T) {
	is := is.New(t)
	arr := []int{4, 5, 6}
	exp := map[int]int{4: 7, 5: 7, 6: 7}
	is.Equal(maps.FromKeys(arr, 7), exp)
}

func TestMap(t *testing.T) {
	is := is.New(t)
	m := map[int32]int64{1: 2, 3: 4, 5: 6}
	f := func(k int32, v int64) (int, string) {
		return int(k + 1), fmt.Sprintf("%d", v)
	}
	is.Equal(maps.Map(m, f), map[int]string{2: "2", 4: "4", 6: "6"})
}

func TestMapKeys(t *testing.T) {
	is := is.New(t)
	m := map[int32]int{1: 2, 3: 4, 5: 6}
	f := func(k int32) int64 {
		return int64(k * 2)
	}
	is.Equal(maps.MapKeys(m, f), map[int64]int{2: 2, 6: 4, 10: 6})
}

func TestMapValues(t *testing.T) {
	is := is.New(t)
	m := map[int]int32{1: 2, 3: 4, 5: 6}
	f := func(v int32) int64 {
		return int64(v * 2)
	}
	is.Equal(maps.MapValues(m, f), map[int]int64{1: 4, 3: 8, 5: 12})
}

func TestMerge(t *testing.T) {
	is := is.New(t)
	m1 := map[int]string{1: "one", 2: "two"}
	m2 := map[int]string{2: "new two", 3: "three"}
	exp := map[int]string{1: "one", 2: "new two", 3: "three"}
	is.Equal(maps.Merge(m1, m2), exp)
}

func TestMergeBy(t *testing.T) {
	is := is.New(t)
	m1 := map[int]string{1: "one", 2: "two"}
	m2 := map[int]string{2: "new two", 3: "three"}
	f := func(k int, a, b string) string {
		is.Equal(k, 2)
		return fmt.Sprintf("%s|%s", a, b)
	}
	exp := map[int]string{1: "one", 2: "two|new two", 3: "three"}
	is.Equal(maps.MergeBy(m1, m2, f), exp)
}

func TestKeys(t *testing.T) {
	is := is.New(t)
	m := map[int]int{1: 2, 3: 4, 5: 6}
	act := maps.Keys(m)
	is.Equal(maps.FromKeys(act, 0), map[int]int{1: 0, 3: 0, 5: 0})
}

func TestValues(t *testing.T) {
	is := is.New(t)
	m := map[int]int{1: 2, 3: 4, 5: 6}
	act := maps.Values(m)
	is.Equal(maps.FromKeys(act, 0), map[int]int{2: 0, 4: 0, 6: 0})
}

func TestTake(t *testing.T) {
	is := is.New(t)
	m := map[int]int{1: 2, 3: 4, 5: 6}
	is.Equal(maps.Take(m, 1, 5), map[int]int{1: 2, 5: 6})
}

func TestWithout(t *testing.T) {
	is := is.New(t)
	m := map[int]int{1: 2, 3: 4, 5: 6}
	is.Equal(maps.Without(m, 1, 5), map[int]int{3: 4})
}
