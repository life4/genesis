package slices_test

import (
	"fmt"

	"github.com/life4/genesis/channels"
	"github.com/life4/genesis/slices"
)

func ExampleAny() {
	even := func(item int) bool { return item%2 == 0 }
	result := slices.Any([]int{1, 2, 3}, even)
	fmt.Println(result)
	result = slices.Any([]int{1, 3, 5}, even)
	fmt.Println(result)
	// Output:
	// true
	// false
}
func ExampleFindIndex() {
	type UserId int
	index := slices.FindIndex(
		[]UserId{1, 2, 3, 4, 5},
		func(el UserId) bool { return el == 3 },
	)
	fmt.Println(index)
	// Output: 2
}
func ExampleAll() {
	even := func(item int) bool { return item%2 == 0 }
	result := slices.All([]int{2, 4, 6}, even)
	fmt.Println(result)
	result = slices.All([]int{2, 4, 5}, even)
	fmt.Println(result)

	// Output:
	// true
	// false
}

func ExampleChunkBy() {
	s := []int{1, 3, 4, 6, 8, 9}
	remainder := func(item int) int { return item % 2 }
	result := slices.ChunkBy(s, remainder)
	fmt.Println(result)
	// Output: [[1 3] [4 6 8] [9]]
}

func ExampleChunkEvery() {
	s := []int{1, 1, 2, 3, 5, 8, 13}
	result, _ := slices.ChunkEvery(s, 3)
	fmt.Println(result)
	// Output: [[1 1 2] [3 5 8] [13]]
}

func ExampleConcat() {
	s1 := []int{3, 4, 5}
	s2 := []int{6, 7, 8}
	result := slices.Concat(s1, s2)
	fmt.Println(result)
	// Output: [3 4 5 6 7 8]
}

func ExampleContains() {
	s := []int{2, 4, 6, 8}
	result := slices.Contains(s, 4)
	fmt.Println(result)
	result = slices.Contains(s, 3)
	fmt.Println(result)

	// Output:
	// true
	// false
}

func ExampleCount() {
	s := []int{1, 0, 1, 0, 0, 1, 1, 0, 1, 0}
	result := slices.Count(s, 1)
	fmt.Println(result)
	// Output: 5
}

func ExampleCountBy() {
	s := []int{1, 2, 3, 4, 5, 6}
	even := func(item int) bool { return item%2 == 0 }
	result := slices.CountBy(s, even)
	fmt.Println(result)
	// Output: 3
}

func ExampleCycle() {
	s := []int{1, 2, 3}
	c := slices.Cycle(s)
	c = channels.Take(c, 5)
	result := channels.ToSlice(c)
	fmt.Println(result)
	// Output: [1 2 3 1 2]
}

func ExampleDedup() {
	s := []int{1, 2, 2, 3, 3, 3, 2, 3, 1, 1}
	result := slices.Dedup(s)
	fmt.Println(result)
	// Output: [1 2 3 2 3 1]
}

func ExampleDelete() {
	s := []int{3, 4, 5, 3, 4, 5}
	result := slices.Delete(s, 4)
	fmt.Println(result)
	// Output: [3 5 3 4 5]
}

func ExampleDeleteAt() {
	s := []int{3, 4, 5, 3, 4, 5}
	result, _ := slices.DeleteAt(s, 1, 3)
	fmt.Println(result)
	// Output: [3 5 4 5]
}

func ExampleDropEvery() {
	s := []int{3, 4, 5, 6, 7, 8}
	result, _ := slices.DropEvery(s, 2, 0)
	fmt.Println(result)
	// Output: [4 6 8]
}

func ExampleMin() {
	s := []int{42, 7, 13}
	min, _ := slices.Min(s)
	fmt.Println(min)
	// Output: 7
}

func ExampleMax() {
	s := []int{7, 42, 13}
	max, _ := slices.Max(s)
	fmt.Println(max)
	// Output: 42
}

func ExampleMap() {
	s := []int{4, 8, 15, 16, 23, 42}
	double := func(el int) int { return el * 2 }
	doubled := slices.Map(s, double)
	fmt.Println(doubled)
	// Output: [8 16 30 32 46 84]
}

func ExampleMapFilter() {
	s := []int{4, 8, 15, 16, 23, 42}
	isEven := func(t int) (string, bool) {
		if t%2 == 0 {
			s := fmt.Sprintf("%d", t)
			return s, true
		} else {
			return "", false
		}

	}
	doubled := slices.MapFilter(s, isEven)
	fmt.Println(doubled)
	// Output: [4 8 16 42]
}
func ExampleMapAsync() {
	pages := slices.MapAsync(
		[]string{"google.com", "go.dev", "golang.org"},
		0,
		func(url string) string {
			return fmt.Sprintf("<web page for %s>", url)
		},
	)
	fmt.Println(pages)
	// [<web page for google.com> <web page for go.dev> <web page for golang.org>]
}
