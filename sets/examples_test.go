package sets_test

import (
	"fmt"

	"github.com/life4/genesis/sets"
)

func ExampleAdd() {
	s := sets.New(3, 4)
	sets.Add(s, 4)
	sets.Add(s, 5)
	fmt.Println(s)
	// Output: map[3:{} 4:{} 5:{}]
}

func ExampleClear() {
	s := sets.New(3, 4)
	sets.Clear(s)
	fmt.Println(s)
	// Output: map[]
}

func ExampleContains() {
	s := sets.New(3, 4, 5)
	result := sets.Contains(s, 4)
	fmt.Println(result)
	// Output: true
}

func ExampleCopy() {
	a := sets.New(3, 4)
	b := sets.Copy(a)
	sets.Add(a, 5)
	sets.Add(b, 6)
	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// map[3:{} 4:{} 5:{}]
	// map[3:{} 4:{} 6:{}]
}

func ExampleDiscard() {
	s := sets.New(3, 4)
	sets.Discard(s, 4)
	sets.Discard(s, 5)
	fmt.Println(s)
	// Output: map[3:{}]
}

func ExampleDifference() {
	a := sets.New(3, 4, 5)
	b := sets.New(5, 6, 7)
	result := sets.Difference(a, b)
	fmt.Println(result)
	// Output: map[3:{} 4:{}]
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

func ExampleFilter() {
	s := sets.New(3, 4, 5, 6, 7, 8)
	even := func(x int) bool { return x%2 == 0 }
	result := sets.Filter(s, even)
	fmt.Println(result)
	// Output: map[4:{} 6:{} 8:{}]
}

func ExampleFromSlice() {
	s := []int{3, 4, 5}
	result := sets.FromSlice(s)
	fmt.Println(result)
	// Output: map[3:{} 4:{} 5:{}]
}

func ExampleIntersect() {
	a := sets.New(3, 4, 5)
	b := sets.New(5, 6, 7)
	result := sets.Intersect(a, b)
	fmt.Println(result)
	// Output: true
}

func ExampleIntersection() {
	a := sets.New(3, 4, 5)
	b := sets.New(4, 5, 6)
	result := sets.Intersection(a, b)
	fmt.Println(result)
	// Output: map[4:{} 5:{}]
}

func ExampleMap() {
	s := sets.New(3, 4, 5)
	double := func(x int) int { return x * 2 }
	result := sets.Map(s, double)
	fmt.Println(result)
	// Output: map[6:{} 8:{} 10:{}]
}

func ExampleMax() {
	s := sets.New(3, 6, 4, 5)
	result, _ := sets.Max(s)
	fmt.Println(result)
	// Output: 6
}

func ExampleMin() {
	s := sets.New(4, 5, 3, 6)
	result, _ := sets.Min(s)
	fmt.Println(result)
	// Output: 3
}

func ExampleNew() {
	result := sets.New(3, 4, 4, 5, 3)
	fmt.Println(result)
	// Output: map[3:{} 4:{} 5:{}]
}

func ExamplePop() {
	s := sets.New(4)
	val, _ := sets.Pop(s)
	fmt.Println(val, s)
	// Output: 4 map[]
}

func ExampleReduce() {
	s := sets.New(3, 4, 5)
	add := func(x, acc int) int { return x + acc }
	result := sets.Reduce(s, 0, add)
	fmt.Println(result)
	// Output: 12
}

func ExampleSubset() {
	a := sets.New(4, 5)
	b := sets.New(3, 4, 5, 6)
	result := sets.Subset(a, b)
	fmt.Println(result)
	// Output: true
}

func ExampleSum() {
	s := sets.New(3, 4, 5)
	result := sets.Sum(s)
	fmt.Println(result)
	// Output: 12
}

func ExampleSymmetricDifference() {
	a := sets.New(3, 4, 5)
	b := sets.New(5, 6, 7)
	result := sets.SymmetricDifference(a, b)
	fmt.Println(result)
	// Output: map[3:{} 4:{} 6:{} 7:{}]
}

func ExampleToSlice() {
	s := sets.New(3)
	result := sets.ToSlice(s)
	fmt.Println(result)
	// Output: [3]
}

func ExampleSuperset() {
	a := sets.New(3, 4, 5, 6)
	b := sets.New(4, 5)
	result := sets.Superset(a, b)
	fmt.Println(result)
	// Output: true
}

func ExampleUnion() {
	a := sets.New(3, 4, 5)
	b := sets.New(5, 6, 7)
	result := sets.Union(a, b)
	fmt.Println(result)
	// Output: map[3:{} 4:{} 5:{} 6:{} 7:{}]
}

func ExampleUnionMany() {
	a := sets.New(3, 4)
	b := sets.New(5, 6)
	c := sets.New(6, 7)
	result := sets.UnionMany(a, b, c)
	fmt.Println(result)
	// Output: map[3:{} 4:{} 5:{} 6:{} 7:{}]
}

func ExampleUpdate() {
	a := sets.New(4, 5)
	b := sets.New(5, 6)
	sets.Update(a, b)
	fmt.Println(a)
	// Output: map[4:{} 5:{} 6:{}]
}
