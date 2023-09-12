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
	is.True(sets.Contains(sets.New(3, 4, 5), 3))
	is.True(sets.Contains(sets.New(3, 4, 5), 4))
	is.True(sets.Contains(sets.New(3, 4, 5), 5))
	is.True(!sets.Contains(sets.New(3, 4, 5), 6))
	is.True(!sets.Contains(emptySet, 3))
	is.True(!sets.Contains(nilSet, 3))
}

func TestDisjoint(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(!sets.Disjoint(sets.New(3, 4, 5), sets.New(3, 4, 5)))
	is.True(!sets.Disjoint(sets.New(3, 4, 5), sets.New(5, 6, 7)))
	is.True(sets.Disjoint(sets.New(3, 4, 5), sets.New(6, 7)))
	is.True(sets.Disjoint(emptySet, emptySet))
	is.True(sets.Disjoint(emptySet, sets.New(6)))
	is.True(sets.Disjoint(sets.New(6), emptySet))
	is.True(sets.Disjoint(sets.New(6), nilSet))
	is.True(sets.Disjoint(nilSet, sets.New(6)))
	is.True(sets.Disjoint(nilSet, nilSet))
}

func TestDisjointMany(t *testing.T) {
	is := is.NewRelaxed(t)

	// two sets
	is.True(!sets.DisjointMany(sets.New(3, 4, 5), sets.New(3, 4, 5)))
	is.True(!sets.DisjointMany(sets.New(3, 4, 5), sets.New(5, 6, 7)))
	is.True(sets.DisjointMany(sets.New(3, 4, 5), sets.New(6, 7)))
	is.True(sets.DisjointMany(emptySet, emptySet))
	is.True(sets.DisjointMany(emptySet, sets.New(6)))
	is.True(sets.DisjointMany(sets.New(6), emptySet))
	is.True(sets.DisjointMany(sets.New(6), nilSet))
	is.True(sets.DisjointMany(nilSet, sets.New(6)))
	is.True(sets.DisjointMany(nilSet, nilSet))

	// three and more sets
	is.True(sets.DisjointMany(sets.New(3), sets.New(4), sets.New(5)))
	is.True(!sets.DisjointMany(sets.New(3), sets.New(4), sets.New(4, 5)))
	is.True(!sets.DisjointMany(sets.New(3), sets.New(4), sets.New(3, 5)))

	// one and zero sets
	is.True(sets.DisjointMany(nilSet))
	is.True(sets.DisjointMany(emptySet))
	is.True(sets.DisjointMany[map[int]struct{}]())
}

func TestEmpty(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(sets.Empty(nilSet))
	is.True(sets.Empty(emptySet))
	is.True(!sets.Empty(sets.New(3)))
	is.True(!sets.Empty(sets.New(3, 4)))
}

func TestEqual(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(sets.Equal(nilSet, nilSet))
	is.True(sets.Equal(emptySet, emptySet))
	is.True(sets.Equal(nilSet, emptySet))
	is.True(sets.Equal(sets.New(3, 4, 5), sets.New(3, 4, 5)))
	is.True(!sets.Equal(sets.New(3, 4, 5), sets.New(3, 4)))
	is.True(!sets.Equal(sets.New(3, 4), sets.New(3, 4, 5)))
	is.True(!sets.Equal(sets.New(3, 4, 6), sets.New(3, 4, 5)))
	is.True(!sets.Equal(sets.New(3), emptySet))
	is.True(!sets.Equal(emptySet, sets.New(3)))
}

func TestEqualMany(t *testing.T) {
	is := is.NewRelaxed(t)

	// two sets
	is.True(sets.EqualMany(nilSet, nilSet))
	is.True(sets.EqualMany(emptySet, emptySet))
	is.True(sets.EqualMany(nilSet, emptySet))
	is.True(sets.EqualMany(sets.New(3, 4, 5), sets.New(3, 4, 5)))
	is.True(!sets.EqualMany(sets.New(3, 4, 5), sets.New(3, 4)))
	is.True(!sets.EqualMany(sets.New(3, 4), sets.New(3, 4, 5)))
	is.True(!sets.EqualMany(sets.New(3, 4, 6), sets.New(3, 4, 5)))
	is.True(!sets.EqualMany(sets.New(3), emptySet))
	is.True(!sets.EqualMany(emptySet, sets.New(3)))

	// one or zero sets
	is.True(sets.EqualMany(emptySet))
	is.True(sets.EqualMany(nilSet))
	is.True(sets.EqualMany[map[int]struct{}]())
}

func TestIntersect(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(sets.Intersect(sets.New(3, 4, 5), sets.New(3, 4, 5)))
	is.True(sets.Intersect(sets.New(3, 4, 5), sets.New(5, 6, 7)))
	is.True(!sets.Intersect(sets.New(3, 4, 5), sets.New(6, 7)))
	is.True(!sets.Intersect(emptySet, emptySet))
	is.True(!sets.Intersect(emptySet, sets.New(6)))
	is.True(!sets.Intersect(sets.New(6), emptySet))
	is.True(!sets.Intersect(sets.New(6), nilSet))
	is.True(!sets.Intersect(nilSet, sets.New(6)))
	is.True(!sets.Intersect(nilSet, nilSet))
}

func TestSubset(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(sets.Subset(sets.New(3, 4, 5), sets.New(3, 4, 5)))
	is.True(sets.Subset(sets.New(3, 4), sets.New(3, 4, 5)))
	is.True(sets.Subset(sets.New(4), sets.New(3, 4, 5)))
	is.True(sets.Subset(emptySet, sets.New(3, 4, 5)))
	is.True(sets.Subset(emptySet, emptySet))
	is.True(sets.Subset(emptySet, nilSet))
	is.True(sets.Subset(nilSet, emptySet))

	is.True(!sets.Subset(sets.New(4), emptySet))
	is.True(!sets.Subset(sets.New(4), nilSet))
	is.True(!sets.Subset(sets.New(4), sets.New(5, 6)))
	is.True(!sets.Subset(sets.New(4, 5), sets.New(5, 6)))
	is.True(!sets.Subset(sets.New(4, 5, 6), sets.New(5, 6)))
}

func TestSuperset(t *testing.T) {
	is := is.NewRelaxed(t)
	is.True(sets.Superset(sets.New(3, 4, 5), sets.New(3, 4, 5)))
	is.True(sets.Superset(sets.New(3, 4, 5), sets.New(3, 4)))
	is.True(sets.Superset(sets.New(3, 4, 5), sets.New(4)))
	is.True(sets.Superset(sets.New(3, 4, 5), emptySet))
	is.True(sets.Superset(emptySet, emptySet))
	is.True(sets.Superset(nilSet, emptySet))
	is.True(sets.Superset(emptySet, nilSet))

	is.True(!sets.Superset(emptySet, sets.New(4)))
	is.True(!sets.Superset(nilSet, sets.New(4)))
	is.True(!sets.Superset(sets.New(5, 6), sets.New(4)))
	is.True(!sets.Superset(sets.New(5, 6), sets.New(4, 5)))
	is.True(!sets.Superset(sets.New(5, 6), sets.New(4, 5, 6)))
}
