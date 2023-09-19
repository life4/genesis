package sets_test

import (
	"testing"

	"github.com/life4/genesis/sets"
	"github.com/matryer/is"
)

func TestCopy(t *testing.T) {
	is := is.New(t)
	a := sets.New(4, 5)
	b := sets.Copy(a)
	a[6] = struct{}{}
	b[7] = struct{}{}
	is.Equal(a, sets.New(4, 5, 6))
	is.Equal(b, sets.New(4, 5, 7))

	is.Equal(sets.Copy(nilSet), nil)
}

func TestDifference(t *testing.T) {
	is := is.NewRelaxed(t)
	is.Equal(sets.Difference(sets.New(3, 4, 5), sets.New(3, 4, 5)), sets.New[int]())
	is.Equal(sets.Difference(sets.New(3, 4, 5), sets.New(3)), sets.New[int](4, 5))
	is.Equal(sets.Difference(sets.New(3, 4, 5), sets.New(3, 6, 7)), sets.New[int](4, 5))
	is.Equal(sets.Difference(sets.New(3), sets.New(3, 6, 7)), sets.New[int]())
	is.Equal(sets.Difference(sets.New[int](), sets.New(3, 6, 7)), sets.New[int]())
	is.Equal(sets.Difference(nilSet, sets.New(3, 6, 7)), sets.New[int]())
}

func TestFilter(t *testing.T) {
	is := is.NewRelaxed(t)
	even := func(x int) bool { return x%2 == 0 }
	is.Equal(sets.Filter(sets.New(3, 5, 7), even), sets.New[int]())
	is.Equal(sets.Filter(sets.New(3, 4, 5), even), sets.New(4))
	is.Equal(sets.Filter(sets.New(3, 4, 5, 6), even), sets.New(4, 6))
	is.Equal(sets.Filter(sets.New(3, 4, 5, 6, 4, 4), even), sets.New(4, 6, 4, 4))
	is.Equal(sets.Filter(nilSet, even), nilSet)
}

func TestFromSlice(t *testing.T) {
	is := is.NewRelaxed(t)
	var nilSlice []int = nil
	is.Equal(sets.FromSlice(nilSlice), sets.New[int]())
	is.Equal(sets.FromSlice([]int{}), sets.New[int]())
	is.Equal(sets.FromSlice([]int{4, 5, 6}), sets.New(4, 5, 6))
	is.Equal(sets.FromSlice([]int{6, 4, 5, 6, 6}), sets.New(4, 5, 6))
}

func TestIntersection(t *testing.T) {
	is := is.NewRelaxed(t)
	is.Equal(sets.Intersection(sets.New(3, 4, 5), sets.New(6, 7, 8)), sets.New[int]())
	is.Equal(sets.Intersection(sets.New(3, 4, 5), sets.New(3, 4, 5)), sets.New[int](3, 4, 5))
	is.Equal(sets.Intersection(sets.New(2, 3, 4), sets.New(3, 4, 5)), sets.New[int](3, 4))
	is.Equal(sets.Intersection(sets.New[int](), sets.New(3, 6, 7)), sets.New[int]())
	is.Equal(sets.Intersection(nilSet, sets.New(3, 6, 7)), sets.New[int]())
}

func TestMap(t *testing.T) {
	is := is.NewRelaxed(t)
	double := func(x int) int { return x * 2 }
	is.Equal(sets.Map(sets.New(3, 4, 5), double), sets.New(6, 8, 10))
	is.Equal(sets.Map(sets.New[int](), double), sets.New[int]())
	is.Equal(sets.Map(nilSet, double), sets.New[int]())
}

func TestNew(t *testing.T) {
	is := is.New(t)
	s := sets.New(3, 4, 5)
	is.Equal(s, map[int]struct{}{3: {}, 4: {}, 5: {}})
}

func TestSymmetricDifference(t *testing.T) {
	is := is.NewRelaxed(t)
	is.Equal(sets.SymmetricDifference(sets.New(3, 4, 5), sets.New(3, 4, 5)), sets.New[int]())
	is.Equal(sets.SymmetricDifference(sets.New(3, 4, 5), sets.New(3)), sets.New[int](4, 5))
	is.Equal(sets.SymmetricDifference(sets.New(3, 4, 5), sets.New(3, 6, 7)), sets.New[int](4, 5, 6, 7))
	is.Equal(sets.SymmetricDifference(sets.New(3), sets.New(3, 6, 7)), sets.New[int](6, 7))
	is.Equal(sets.SymmetricDifference(sets.New[int](), sets.New(3, 6, 7)), sets.New(3, 6, 7))
	is.Equal(sets.SymmetricDifference(nilSet, sets.New(3, 6, 7)), sets.New(3, 6, 7))
}

func TestUnion(t *testing.T) {
	is := is.NewRelaxed(t)
	is.Equal(sets.Union(sets.New(3, 4, 5), sets.New(3, 4, 5)), sets.New(3, 4, 5))
	is.Equal(sets.Union(sets.New(3, 4, 5), sets.New(3)), sets.New[int](3, 4, 5))
	is.Equal(sets.Union(sets.New(3, 4, 5), sets.New(3, 6, 7)), sets.New[int](3, 4, 5, 6, 7))
	is.Equal(sets.Union(sets.New(3), sets.New(3, 6, 7)), sets.New[int](3, 6, 7))
	is.Equal(sets.Union(sets.New(3, 4), sets.New(6, 7)), sets.New[int](3, 4, 6, 7))
	is.Equal(sets.Union(sets.New[int](), sets.New(3, 6, 7)), sets.New(3, 6, 7))
	is.Equal(sets.Union(sets.New[int](), sets.New[int]()), sets.New[int]())
	is.Equal(sets.Union(nilSet, sets.New(3, 6, 7)), sets.New[int](3, 6, 7))
}

func TestUnionMany(t *testing.T) {
	is := is.NewRelaxed(t)

	is.Equal(sets.UnionMany[map[int]struct{}](), sets.New[int]())
	is.Equal(sets.UnionMany(sets.New(3, 4, 5)), sets.New(3, 4, 5))

	is.Equal(sets.UnionMany(sets.New(3, 4, 5), sets.New(3, 4, 5)), sets.New(3, 4, 5))
	is.Equal(sets.UnionMany(sets.New(3, 4, 5), sets.New(3)), sets.New[int](3, 4, 5))
	is.Equal(sets.UnionMany(sets.New(3, 4, 5), sets.New(3, 6, 7)), sets.New[int](3, 4, 5, 6, 7))
	is.Equal(sets.UnionMany(sets.New(3), sets.New(3, 6, 7)), sets.New[int](3, 6, 7))
	is.Equal(sets.UnionMany(sets.New(3, 4), sets.New(6, 7)), sets.New[int](3, 4, 6, 7))
	is.Equal(sets.UnionMany(sets.New[int](), sets.New(3, 6, 7)), sets.New(3, 6, 7))
	is.Equal(sets.UnionMany(sets.New[int](), sets.New[int]()), sets.New[int]())
	is.Equal(sets.UnionMany(nilSet, sets.New(3, 6, 7)), sets.New[int](3, 6, 7))
}
