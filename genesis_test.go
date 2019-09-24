package genesis_test

import (
	"fmt"

	"github.com/life4/genesis"
)

func ExampleSliceMin() {
	s := []int{42, 7, 13}
	min := genesis.SliceInt{s}.Min()
	fmt.Println(min)
	// Output: 7
}

func ExampleSliceMax() {
	s := []int{7, 42, 13}
	max := genesis.SliceInt{s}.Max()
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
