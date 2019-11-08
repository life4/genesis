package genesis_test

import (
	"fmt"

	"github.com/life4/genesis"
)

func ExampleInterface() {
	type UserId int
	ids := []UserId{1, 2, 3, 4, 5}
	// https://github.com/golang/go/wiki/InterfaceSlice
	idsInterface := make([]interface{}, len(ids), len(ids))
	for i := range ids {
		idsInterface[i] = ids[i]
	}
	index := genesis.SliceInterface{idsInterface}.FindIndex(
		func(el interface{}) bool { return el.(UserId) == 3 },
	)
	fmt.Println(index)
	// Output: 2
}

func ExampleSliceAny() {
	even := func(item int) bool { return item%2 == 0 }

	s := []int{1, 2, 3}
	result := genesis.SliceInt{s}.Any(even)
	fmt.Println(result)

	s = []int{1, 3, 5}
	result = genesis.SliceInt{s}.Any(even)
	fmt.Println(result)

	// Output:
	// true
	// false
}

func ExampleSliceAll() {
	even := func(item int) bool { return item%2 == 0 }

	s := []int{2, 4, 6}
	result := genesis.SliceInt{s}.All(even)
	fmt.Println(result)

	s = []int{2, 4, 5}
	result = genesis.SliceInt{s}.All(even)
	fmt.Println(result)

	// Output:
	// true
	// false
}

func ExampleSliceChunkBy() {
	s := []int{1, 3, 4, 6, 8, 9}
	remainder := func(item int) int { return item % 2 }
	result := genesis.SliceInt{s}.ChunkByInt(remainder)
	fmt.Println(result)
	// Output: [[1 3] [4 6 8] [9]]
}

func ExampleSliceChunkEvery() {
	s := []int{1, 1, 2, 3, 5, 8, 13}
	result, _ := genesis.SliceInt{s}.ChunkEvery(3)
	fmt.Println(result)
	// Output: [[1 1 2] [3 5 8] [13]]
}

func ExampleSliceContains() {
	s := []int{2, 4, 6, 8}
	result := genesis.SliceInt{s}.Contains(4)
	fmt.Println(result)

	result = genesis.SliceInt{s}.Contains(3)
	fmt.Println(result)

	// Output:
	// true
	// false
}

func ExampleSliceCount() {
	s := []int{1, 0, 1, 0, 0, 1, 1, 0, 1, 0}
	result := genesis.SliceInt{s}.Count(1)
	fmt.Println(result)
	// Output: 5
}

func ExampleSliceCountBy() {
	s := []int{1, 2, 3, 4, 5, 6}
	even := func(item int) bool { return item%2 == 0 }
	result := genesis.SliceInt{s}.CountBy(even)
	fmt.Println(result)
	// Output: 3
}

func ExampleSliceCycle() {
	s := []int{1, 2, 3}
	channel := genesis.SliceInt{s}.Cycle()
	limited := genesis.ChannelInt{channel}.Take(5)
	result := genesis.ChannelInt{limited}.ToSlice()
	fmt.Println(result)
	// Output: [1 2 3 1 2]
}

func ExampleSliceDedup() {
	s := []int{1, 2, 2, 3, 3, 3, 2, 3, 1, 1}
	result := genesis.SliceInt{s}.Dedup()
	fmt.Println(result)
	// Output: [1 2 3 2 3 1]
}

func ExampleSliceMin() {
	s := []int{42, 7, 13}
	min, _ := genesis.SliceInt{s}.Min()
	fmt.Println(min)
	// Output: 7
}

func ExampleSliceMax() {
	s := []int{7, 42, 13}
	max, _ := genesis.SliceInt{s}.Max()
	fmt.Println(max)
	// Output: 42
}

func ExampleSliceMap() {
	s := []int{4, 8, 15, 16, 23, 42}
	double := func(el int) int { return el * 2 }
	doubled := genesis.SliceInt{s}.MapInt(double)
	fmt.Println(doubled)
	// Output: [8 16 30 32 46 84]
}
