package slices_test

import (
	"errors"
	"fmt"

	"github.com/life4/genesis/channels"
	"github.com/life4/genesis/slices"
)

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

func ExampleAllAsync() {
	even := func(item int) bool { return item%2 == 0 }
	result := slices.AllAsync([]int{2, 4, 6}, 0, even)
	fmt.Println(result)
	// Output: true
}

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

func ExampleAnyAsync() {
	even := func(item int) bool { return item%2 == 0 }
	result := slices.AnyAsync([]int{1, 2, 3}, 0, even)
	fmt.Println(result)
	// Output: true
}

func ExampleChoice() {
	result, _ := slices.Choice([]int{3, 4, 5, 6}, 13)
	fmt.Println(result)
	// Output: 3
}

func ExampleCopy() {
	orig := []int{3, 4}
	copy := slices.Copy(orig)
	orig = append(orig, 5)
	copy = append(copy, 6)
	fmt.Println(orig)
	fmt.Println(copy)
	// Output:
	// [3 4 5]
	// [3 4 6]
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

func ExampleDifference() {
	s1 := []int{3, 4, 4, 5, 6, 6, 7}
	s2 := []int{5, 5, 4, 8}
	result := slices.Difference(s1, s2)
	fmt.Println(result)
	// Output: [3 6 7]
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

func ExampleDropZero() {
	s := []int{4, 5, 0, 6, 0, 0}
	result := slices.DropZero(s)
	fmt.Println(result)
	// Output: [4 5 6]
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

func ExampleEachAsync() {
	s := []int{4, 5, 6}
	sum := 0
	slices.EachAsync(s, 0, func(x int) {
		sum += x
	})
	fmt.Println(sum)
	// Output: 15
}

func ExampleEachErr() {
	s := []int{4, 5, 6, 7, 8}
	err := slices.EachErr(s, func(x int) error {
		if x == 6 {
			return errors.New("six found")
		}
		fmt.Println(x * 2)
		return nil
	})
	fmt.Println(err)
	// Output:
	// 8
	// 10
	// six found
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

func ExampleEqualBy() {
	s1 := []int{3, 4, -5}
	s2 := []int{3, -4, 5}
	absEq := func(a, b int) bool {
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		return a == b
	}
	result := slices.EqualBy(s1, s2, absEq)
	fmt.Println(result)
	// Output: true
}

func ExampleFilter() {
	s := []int{4, 5, 6, 7, 8, 10, 12, 13}
	even := func(x int) bool { return x%2 == 0 }
	result := slices.Filter(s, even)
	fmt.Println(result)
	// Output: [4 6 8 10 12]
}

func ExampleFilterAsync() {
	s := []int{4, 5, 6, 7, 8, 10, 12, 13}
	even := func(x int) bool { return x%2 == 0 }
	result := slices.FilterAsync(s, 0, even)
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

func ExampleGrow() {
	s := make([]int, 1, 4)
	fmt.Printf("Before: len=%d, cap=%d\n", len(s), cap(s))
	r := slices.Grow(s, 5)
	fmt.Printf("After:  len=%d, cap=%d\n", len(r), cap(r))
	// Output:
	// Before: len=1, cap=4
	// After:  len=1, cap=8
}

func ExampleIndex() {
	s := []int{3, 4, 5}
	index, _ := slices.Index(s, 4)
	fmt.Println(index)
	// Output: 1
}

func ExampleIndexBy() {
	s := []int{5, 7, 9, 8, 3, 6}
	even := func(x int) bool { return x%2 == 0 }
	result, _ := slices.IndexBy(s, even)
	fmt.Println(result)
	// Output: 3
}

func ExampleInsertAt() {
	s := []int{3, 4, 5}
	result, _ := slices.InsertAt(s, 1, 6)
	fmt.Println(result)
	// Output: [3 6 4 5]
}

func ExampleIntersect() {
	s1 := []int{3, 4, 5, 5, 6, 6, 7}
	s2 := []int{6, 5, 5, 4, 8}
	result := slices.Intersect(s1, s2)
	fmt.Println(result)
	// Output: [4 5 6]
}

func ExampleUnion() {
	s1 := []int{3, 4, 5, 5, 6, 6, 7}
	s2 := []int{6, 5, 5, 4, 8}
	result := slices.Union(s1, s2)
	fmt.Println(result)
	// Output: [3 4 5 6 7 8]
}

func ExampleIntersperse() {
	s := []int{3, 4, 5}
	result := slices.Intersperse(s, 1)
	fmt.Println(result)
	// Output: [3 1 4 1 5]
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
	// Output:
	// [<web page for google.com> <web page for go.dev> <web page for golang.org>]
}

func ExamplePartition() {
	s := []int{4, 5, 6, 7, 8, 8, 7}
	isEven := func(x int) bool { return x%2 == 0 }
	even, odd := slices.Partition(s, isEven)
	fmt.Println(even)
	fmt.Println(odd)
	// Output:
	// [4 6 8 8]
	// [5 7 7]
}

func ExamplePermutations() {
	s := []int{1, 2, 3}
	ch := slices.Permutations(s, 2)
	result := make([][]int, 0)
	for x := range ch {
		result = append(result, x)
	}
	fmt.Println(result)
	// Output: [[1 2] [1 3] [2 1] [2 3] [3 1] [3 2]]
}

func ExamplePrepend() {
	s := []int{4, 5, 6}
	result := slices.Prepend(s, 2, 3)
	fmt.Println(result)
	// Output: [2 3 4 5 6]
}

func ExampleProduct() {
	s := []int{1, 2, 3}
	ch := slices.Product(s, 2)
	result := make([][]int, 0)
	for x := range ch {
		result = append(result, x)
	}
	fmt.Println(result)
	// Output: [[1 1] [1 2] [1 3] [2 1] [2 2] [2 3] [3 1] [3 2] [3 3]]
}

func ExampleProduct2() {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	ch := slices.Product2(s1, s2)
	result := make([][]int, 0)
	for x := range ch {
		result = append(result, x)
	}
	fmt.Println(result)
	// Output: [[1 3] [1 4] [2 3] [2 4]]
}

func ExampleReduce() {
	s := []int{3, 4, 5}
	sum := func(a, b int) int {
		fmt.Printf("Received %d and %d\n", a, b)
		return a + b
	}
	result := slices.Reduce(s, 0, sum)
	fmt.Printf("Result is %d\n", result)
	// Output:
	// Received 3 and 0
	// Received 4 and 3
	// Received 5 and 7
	// Result is 12
}

func ExampleReduceAsync() {
	s := []int{3, 4, 5}
	sum := func(a, b int) int { return a + b }
	result := slices.ReduceAsync(s, 0, sum)
	fmt.Println(result)
	// Output: 12
}

func ExampleReduceWhile() {
	s := []int{3, 4, 5, 6}
	sum := func(a, b int) (int, error) {
		fmt.Printf("Received %d and %d\n", a, b)
		if a == 6 {
			return b, errors.New("got six")
		}
		return a + b, nil
	}
	result, err := slices.ReduceWhile(s, 0, sum)
	fmt.Printf("Result is %d\n", result)
	fmt.Printf("Error is '%v'\n", err)
	// Output:
	// Received 3 and 0
	// Received 4 and 3
	// Received 5 and 7
	// Received 6 and 12
	// Result is 12
	// Error is 'got six'
}

func ExampleReject() {
	s := []int{4, 5, 6, 7, 8, 10, 12, 13}
	odd := func(x int) bool { return x%2 == 1 }
	result := slices.Reject(s, odd)
	fmt.Println(result)
	// Output: [4 6 8 10 12]
}

func ExampleRepeat() {
	s := []int{4, 5, 6}
	result := slices.Repeat(s, 3)
	fmt.Println(result)
	// Output: [4 5 6 4 5 6 4 5 6]
}

func ExampleReplace() {
	s := []int{4, 5, 6, 7}
	result, _ := slices.Replace(s, 1, 3, 9)
	fmt.Println(result)
	// Output: [4 9 9 7]
}

func ExampleReverse() {
	s := []int{3, 4, 5}
	result := slices.Reverse(s)
	fmt.Println(result)
	// Output: [5 4 3]
}

func ExampleSame() {
	s := []int{3, 3, 3, 3}
	result := slices.Same(s)
	fmt.Println(result)
	// Output: true
}

func ExampleScan() {
	s := []int{3, 4, 5}
	sum := func(a, b int) int { return a + b }
	result := slices.Scan(s, 0, sum)
	fmt.Println(result)
	// Output: [3 7 12]
}

func ExampleShrink() {
	s := make([]int, 1, 4)
	fmt.Printf("Before: len=%d, cap=%d\n", len(s), cap(s))
	r := slices.Shrink(s)
	fmt.Printf("After:  len=%d, cap=%d\n", len(r), cap(r))
	// Output:
	// Before: len=1, cap=4
	// After:  len=1, cap=1
}

func ExampleShuffle() {
	s := []int{3, 4, 5, 6, 7, 8}
	slices.Shuffle(s, 13)
	fmt.Println(s)
	// Output: [7 8 5 3 6 4]
}

func ExampleSort() {
	s := []int{7, 6, 8, 5, 3, 6, 4}
	result := slices.Sort(s)
	fmt.Println(result)
	// Output: [3 4 5 6 6 7 8]
}

func ExampleSortBy() {
	s := []int{7, -6, 6, 8, 5, 3, 6, 4}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	result := slices.SortBy(s, abs)
	fmt.Println(result)
	// Output: [3 4 5 -6 6 6 7 8]
}

func ExampleSorted() {
	s := []int{3, 3, 4, 5, 5, 5, 6, 7, 7, 12, 15}
	result := slices.Sorted(s)
	fmt.Println(result)
	// Output: true
}

func ExampleSortedUnique() {
	s := []int{3, 4, 5, 6, 7, 8, 12, 16, 19}
	result := slices.SortedUnique(s)
	fmt.Println(result)
	// Output: true
}

func ExampleSplit() {
	s := []int{3, 4, 1, 5, 1, 1, 6, 7, 1}
	result := slices.Split(s, 1)
	fmt.Println(result)
	// Output: [[3 4] [5] [] [6 7] []]
}

func ExampleStartsWith() {
	s := []int{3, 4, 5, 6, 7, 8}
	result := slices.StartsWith(s, []int{3, 4})
	fmt.Println(result)
	// Output: true
}

func ExampleSum() {
	s := []int{3, 4, 5}
	result := slices.Sum(s)
	fmt.Println(result)
	// Output: 12
}

func ExampleTakeEvery() {
	s := []int{3, 4, 5, 6, 7, 8}
	result, _ := slices.TakeEvery(s, 2, 0)
	fmt.Println(result)
	// Output: [3 5 7]
}

func ExampleTakeRandom() {
	s := []int{3, 4, 5, 6, 7, 8}
	result, _ := slices.TakeRandom(s, 3, 13)
	fmt.Println(result)
	// Output: [7 8 5]
}

func ExampleTakeWhile() {
	s := []int{3, 5, 7, 8, 9, 10, 11}
	odd := func(x int) bool { return x%2 == 1 }
	result := slices.TakeWhile(s, odd)
	fmt.Println(result)
	// Output: [3 5 7]
}

func ExampleToChannel() {
	s := []int{3, 4, 5}
	ch := slices.ToChannel(s)
	result := make([]int, 0)
	for x := range ch {
		result = append(result, x)
	}
	fmt.Println(result)
	// Output: [3 4 5]
}

func ExampleToKeys() {
	s := []int{3, 4, 5}
	result := slices.ToKeys(s, 2)
	fmt.Println(result)
	// Output: map[3:2 4:2 5:2]
}

func ExampleToMap() {
	s := []int{3, 4, 5}
	result := slices.ToMap(s)
	fmt.Println(result)
	// Output: map[0:3 1:4 2:5]
}

func ExampleUniq() {
	s := []int{3, 3, 4, 5, 4, 3, 3}
	result := slices.Uniq(s)
	fmt.Println(result)
	// Output: [3 4 5]
}

func ExampleUnique() {
	s := []int{3, 4, 5, 3}
	result := slices.Unique(s)
	fmt.Println(result)
	// Output: false
}

func ExampleWindow() {
	s := []int{3, 4, 5, 6}
	result, _ := slices.Window(s, 2)
	fmt.Println(result)
	// Output: [[3 4] [4 5] [5 6]]
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

func ExampleZip() {
	s1 := []int{3, 4, 5}
	s2 := []int{6, 7, 8, 9}
	ch := slices.Zip(s1, s2)
	result := make([][]int, 0)
	for x := range ch {
		result = append(result, x)
	}
	fmt.Println(result)
	// Output: [[3 6] [4 7] [5 8]]
}
