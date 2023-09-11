package sets_test

import (
	"testing"

	"github.com/life4/genesis/sets"
	"github.com/matryer/is"
)

var (
	nilSet   map[int]struct{} = nil
	emptySet map[int]struct{} = map[int]struct{}{}
)

func TestContains(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(sets.Contains(sets.Set(3, 4, 5), 3))
	is.True(sets.Contains(sets.Set(3, 4, 5), 4))
	is.True(sets.Contains(sets.Set(3, 4, 5), 5))
	is.True(!sets.Contains(sets.Set(3, 4, 5), 6))
	is.True(!sets.Contains(emptySet, 3))
	is.True(!sets.Contains(nilSet, 3))
}

func TestDisjoint(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(!sets.Disjoint(sets.Set(3, 4, 5), sets.Set(3, 4, 5)))
	is.True(!sets.Disjoint(sets.Set(3, 4, 5), sets.Set(5, 6, 7)))
	is.True(sets.Disjoint(sets.Set(3, 4, 5), sets.Set(6, 7)))
	is.True(sets.Disjoint(emptySet, emptySet))
	is.True(sets.Disjoint(emptySet, sets.Set(6)))
	is.True(sets.Disjoint(sets.Set(6), emptySet))
	is.True(sets.Disjoint(sets.Set(6), nilSet))
	is.True(sets.Disjoint(nilSet, sets.Set(6)))
	is.True(sets.Disjoint(nilSet, nilSet))
}

func TestDisjointMany(t *testing.T) {
	is := is.NewRelaxed(t)

	// two sets
	is.True(!sets.DisjointMany(sets.Set(3, 4, 5), sets.Set(3, 4, 5)))
	is.True(!sets.DisjointMany(sets.Set(3, 4, 5), sets.Set(5, 6, 7)))
	is.True(sets.DisjointMany(sets.Set(3, 4, 5), sets.Set(6, 7)))
	is.True(sets.DisjointMany(emptySet, emptySet))
	is.True(sets.DisjointMany(emptySet, sets.Set(6)))
	is.True(sets.DisjointMany(sets.Set(6), emptySet))
	is.True(sets.DisjointMany(sets.Set(6), nilSet))
	is.True(sets.DisjointMany(nilSet, sets.Set(6)))
	is.True(sets.DisjointMany(nilSet, nilSet))

	// three and more sets
	is.True(sets.DisjointMany(sets.Set(3), sets.Set(4), sets.Set(5)))
	is.True(!sets.DisjointMany(sets.Set(3), sets.Set(4), sets.Set(4, 5)))
	is.True(!sets.DisjointMany(sets.Set(3), sets.Set(4), sets.Set(3, 5)))

	// one and zero sets
	is.True(sets.DisjointMany(nilSet))
	is.True(sets.DisjointMany(emptySet))
	is.True(sets.DisjointMany[map[int]struct{}]())
}

func TestEmpty(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(sets.Empty(nilSet))
	is.True(sets.Empty(emptySet))
	is.True(!sets.Empty(sets.Set(3)))
	is.True(!sets.Empty(sets.Set(3, 4)))
}

func TestEqual(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(sets.Equal(nilSet, nilSet))
	is.True(sets.Equal(emptySet, emptySet))
	is.True(sets.Equal(nilSet, emptySet))
	is.True(sets.Equal(sets.Set(3, 4, 5), sets.Set(3, 4, 5)))
	is.True(!sets.Equal(sets.Set(3, 4, 5), sets.Set(3, 4)))
	is.True(!sets.Equal(sets.Set(3, 4), sets.Set(3, 4, 5)))
	is.True(!sets.Equal(sets.Set(3, 4, 6), sets.Set(3, 4, 5)))
	is.True(!sets.Equal(sets.Set(3), emptySet))
	is.True(!sets.Equal(emptySet, sets.Set(3)))
}

func TestEqualMany(t *testing.T) {
	is := is.NewRelaxed(t)

	// two sets
	is.True(sets.EqualMany(nilSet, nilSet))
	is.True(sets.EqualMany(emptySet, emptySet))
	is.True(sets.EqualMany(nilSet, emptySet))
	is.True(sets.EqualMany(sets.Set(3, 4, 5), sets.Set(3, 4, 5)))
	is.True(!sets.EqualMany(sets.Set(3, 4, 5), sets.Set(3, 4)))
	is.True(!sets.EqualMany(sets.Set(3, 4), sets.Set(3, 4, 5)))
	is.True(!sets.EqualMany(sets.Set(3, 4, 6), sets.Set(3, 4, 5)))
	is.True(!sets.EqualMany(sets.Set(3), emptySet))
	is.True(!sets.EqualMany(emptySet, sets.Set(3)))

	// one or zero sets
	is.True(sets.EqualMany(emptySet))
	is.True(sets.EqualMany(nilSet))
	is.True(sets.EqualMany[map[int]struct{}]())
}

func TestIntersect(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(sets.Intersect(sets.Set(3, 4, 5), sets.Set(3, 4, 5)))
	is.True(sets.Intersect(sets.Set(3, 4, 5), sets.Set(5, 6, 7)))
	is.True(!sets.Intersect(sets.Set(3, 4, 5), sets.Set(6, 7)))
	is.True(!sets.Intersect(emptySet, emptySet))
	is.True(!sets.Intersect(emptySet, sets.Set(6)))
	is.True(!sets.Intersect(sets.Set(6), emptySet))
	is.True(!sets.Intersect(sets.Set(6), nilSet))
	is.True(!sets.Intersect(nilSet, sets.Set(6)))
	is.True(!sets.Intersect(nilSet, nilSet))
}

func TestSubset(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(sets.Subset(sets.Set(3, 4, 5), sets.Set(3, 4, 5)))
	is.True(sets.Subset(sets.Set(3, 4), sets.Set(3, 4, 5)))
	is.True(sets.Subset(sets.Set(4), sets.Set(3, 4, 5)))
	is.True(sets.Subset(emptySet, sets.Set(3, 4, 5)))
	is.True(sets.Subset(emptySet, emptySet))
	is.True(sets.Subset(emptySet, nilSet))
	is.True(sets.Subset(nilSet, emptySet))

	is.True(!sets.Subset(sets.Set(4), emptySet))
	is.True(!sets.Subset(sets.Set(4), nilSet))
	is.True(!sets.Subset(sets.Set(4), sets.Set(5, 6)))
	is.True(!sets.Subset(sets.Set(4, 5), sets.Set(5, 6)))
	is.True(!sets.Subset(sets.Set(4, 5, 6), sets.Set(5, 6)))
}

func TestSuperset(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(sets.Superset(sets.Set(3, 4, 5), sets.Set(3, 4, 5)))
	is.True(sets.Superset(sets.Set(3, 4, 5), sets.Set(3, 4)))
	is.True(sets.Superset(sets.Set(3, 4, 5), sets.Set(4)))
	is.True(sets.Superset(sets.Set(3, 4, 5), emptySet))
	is.True(sets.Superset(emptySet, emptySet))
	is.True(sets.Superset(nilSet, emptySet))
	is.True(sets.Superset(emptySet, nilSet))

	is.True(!sets.Superset(emptySet, sets.Set(4)))
	is.True(!sets.Superset(nilSet, sets.Set(4)))
	is.True(!sets.Superset(sets.Set(5, 6), sets.Set(4)))
	is.True(!sets.Superset(sets.Set(5, 6), sets.Set(4, 5)))
	is.True(!sets.Superset(sets.Set(5, 6), sets.Set(4, 5, 6)))
}
