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

func ExampleChoice() {
	result, _ := slices.Choice([]int{3, 4, 5, 6}, 13)
	fmt.Println(result)
	// Output: 3
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

func ExampleDedupBy() {
	s := []int{1, 2, -2, -3, 3, -3, 2, 3, 1, 1}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	result := slices.DedupBy(s, abs)
	fmt.Println(result)
	// Output: [1 2 -3 2 3 1]
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

func ExampleDeleteAll() {
	s := []int{3, 4, 5, 3, 4, 5}
	result := slices.DeleteAll(s, 3)
	fmt.Println(result)
	// Output: [4 5 4 5]
}

func ExampleDropEvery() {
	s := []int{3, 4, 5, 6, 7, 8}
	result, _ := slices.DropEvery(s, 2, 0)
	fmt.Println(result)
	// Output: [4 6 8]
}

func ExampleDropWhile() {
	s := []int{2, 4, 6, 7, 8, 9, 10}
	even := func(x int) bool { return x%2 == 0 }
	result := slices.DropWhile(s, even)
	fmt.Println(result)
	// Output: [7 8 9 10]
}

func ExampleEach() {
	s := []int{4, 5, 6}
	slices.Each(s, func(x int) {
		fmt.Println(x * 2)
	})
	// Output:
	// 8
	// 10
	// 12
}

func ExampleEndsWith() {
	s := []int{3, 4, 5, 6, 7, 8}
	result := slices.EndsWith(s, []int{7, 8})
	fmt.Println(result)
	// Output: true
}

func ExampleEqual() {
	s1 := []int{3, 4, 5}
	s2 := []int{3, 4, 5, 6}
	result := slices.Equal(s1, s2)
	fmt.Println(result)
	// Output: false
}

func ExampleFilter() {
	s := []int{4, 5, 6, 7, 8, 10, 12, 13}
	even := func(x int) bool { return x%2 == 0 }
	result := slices.Filter(s, even)
	fmt.Println(result)
	// Output: [4 6 8 10 12]
}

func ExampleFind() {
	s := []int{5, 7, 9, 4, 3, 6}
	even := func(x int) bool { return x%2 == 0 }
	result, _ := slices.Find(s, even)
	fmt.Println(result)
	// Output: 4
}

func ExampleGroupBy() {
	s := []int{3, 4, 5, 13, 14, 15, 23, 33}
	even := func(x int) int { return x % 10 }
	result := slices.GroupBy(s, even)
	fmt.Println(result)
	// Output: map[3:[3 13 23 33] 4:[4 14] 5:[5 15]]
}

func ExampleIndex() {
	s := []int{3, 4, 5}
	index, _ := slices.Index(s, 4)
	fmt.Println(index)
	// Output: 1
}

func ExampleInsertAt() {
	s := []int{3, 4, 5}
	result, _ := slices.InsertAt(s, 1, 6)
	fmt.Println(result)
	// Output: [3 6 4 5]
}

func ExampleJoin() {
	s := []int{3, 4, 5}
	result := slices.Join(s, "; ")
	fmt.Println(result)
	// Output: 3; 4; 5
}

func ExampleLast() {
	s := []int{3, 4, 5}
	result, _ := slices.Last(s)
	fmt.Println(result)
	// Output: 5
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

func ExampleReverse() {
	s := []int{3, 4, 5}
	result := slices.Reverse(s)
	fmt.Println(result)
	// Output: [5 4 3]
}

func ExampleStartsWith() {
	s := []int{3, 4, 5, 6, 7, 8}
	result := slices.StartsWith(s, []int{3, 4})
	fmt.Println(result)
	// Output: true
}

func ExampleWithout() {
	s := []int{3, 4, 5, 6, 3, 4, 5, 6, 7, 8}
	result := slices.Without(s, 4, 5)
	fmt.Println(result)
	// Output: [3 6 3 6 7 8]
}

func ExampleWrap() {
	result := slices.Wrap(4)
	fmt.Println(result)
	// Output: [4]
}
