package gslices_test

import (
	"fmt"

	"github.com/life4/genesis/gchannels"
	"github.com/life4/genesis/gslices"
)

func ExampleAny() {
	even := func(item int) bool { return item%2 == 0 }
	result := gslices.Any([]int{1, 2, 3}, even)
	fmt.Println(result)
	result = gslices.Any([]int{1, 3, 5}, even)
	fmt.Println(result)
	// Output:
	// true
	// false
}
func ExampleFindIndex() {
	type UserId int
	index := gslices.FindIndex(
		[]UserId{1, 2, 3, 4, 5},
		func(el UserId) bool { return el == 3 },
	)
	fmt.Println(index)
	// Output: 2
}
func ExampleAll() {
	even := func(item int) bool { return item%2 == 0 }
	result := gslices.All([]int{2, 4, 6}, even)
	fmt.Println(result)
	result = gslices.All([]int{2, 4, 5}, even)
	fmt.Println(result)

	// Output:
	// true
	// false
}

func ExampleChunkBy() {
	s := []int{1, 3, 4, 6, 8, 9}
	remainder := func(item int) int { return item % 2 }
	result := gslices.ChunkBy(s, remainder)
	fmt.Println(result)
	// Output: [[1 3] [4 6 8] [9]]
}

func ExampleChunkEvery() {
	s := []int{1, 1, 2, 3, 5, 8, 13}
	result, _ := gslices.ChunkEvery(s, 3)
	fmt.Println(result)
	// Output: [[1 1 2] [3 5 8] [13]]
}

func ExampleContains() {
	s := []int{2, 4, 6, 8}
	result := gslices.Contains(s, 4)
	fmt.Println(result)
	result = gslices.Contains(s, 3)
	fmt.Println(result)

	// Output:
	// true
	// false
}

func ExampleCount() {
	s := []int{1, 0, 1, 0, 0, 1, 1, 0, 1, 0}
	result := gslices.Count(s, 1)
	fmt.Println(result)
	// Output: 5
}

func ExampleCountBy() {
	s := []int{1, 2, 3, 4, 5, 6}
	even := func(item int) bool { return item%2 == 0 }
	result := gslices.CountBy(s, even)
	fmt.Println(result)
	// Output: 3
}

func ExampleCycle() {
	s := []int{1, 2, 3}
	c := gslices.Cycle(s)
	c = gchannels.Take(c, 5)
	result := gchannels.ToSlice(c)
	fmt.Println(result)
	// Output: [1 2 3 1 2]
}

func ExampleDedup() {
	s := []int{1, 2, 2, 3, 3, 3, 2, 3, 1, 1}
	result := gslices.Dedup(s)
	fmt.Println(result)
	// Output: [1 2 3 2 3 1]
}

func ExampleMin() {
	s := []int{42, 7, 13}
	min, _ := gslices.Min(s)
	fmt.Println(min)
	// Output: 7
}

func ExampleMax() {
	s := []int{7, 42, 13}
	max, _ := gslices.Max(s)
	fmt.Println(max)
	// Output: 42
}

func ExampleMap() {
	s := []int{4, 8, 15, 16, 23, 42}
	double := func(el int) int { return el * 2 }
	doubled := gslices.Map(s, double)
	fmt.Println(doubled)
	// Output: [8 16 30 32 46 84]
}
