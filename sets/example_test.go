package sets_test

import (
	"fmt"

	"github.com/life4/genesis/sets"
)

func ExampleContains() {
	s := sets.New(3, 4, 5)
	result := sets.Contains(s, 4)
	fmt.Println(result)
	// Output: true
}

func ExampleDisjoint() {
	a := sets.New(3, 4, 5)
	b := sets.New(5, 6, 7)
	result := sets.Disjoint(a, b)
	fmt.Println(result) // a and b both share 5
	// Output: false
}

func ExampleDisjointMany() {
	a := sets.New(3, 4)
	b := sets.New(5, 6)
	c := sets.New(6, 7)
	result := sets.DisjointMany(a, b, c)
	fmt.Println(result) // b and c both share 6
	// Output: false
}

func ExampleEmpty() {
	s := sets.New[int]()
	result := sets.Empty(s)
	fmt.Println(result)
	// Output: true
}

func ExampleEqual() {
	a := sets.New(3, 4, 5)
	b := sets.New(3, 4, 6)
	result := sets.Equal(a, b)
	fmt.Println(result)
	// Output: false
}

func ExampleEqualMany() {
	a := sets.New(3, 4, 5)
	b := sets.New(3, 4, 5)
	c := sets.New(3, 4, 6)
	result := sets.EqualMany(a, b, c)
	fmt.Println(result)
	// Output: false
}

func ExampleIntersect() {
	a := sets.New(3, 4, 5)
	b := sets.New(5, 6, 7)
	result := sets.Intersect(a, b)
	fmt.Println(result)
	// Output: true
}

func ExampleSubset() {
	a := sets.New(4, 5)
	b := sets.New(3, 4, 5, 6)
	result := sets.Subset(a, b)
	fmt.Println(result)
	// Output: true
}

func ExampleSuperset() {
	a := sets.New(3, 4, 5, 6)
	b := sets.New(4, 5)
	result := sets.Superset(a, b)
	fmt.Println(result)
	// Output: true
}
