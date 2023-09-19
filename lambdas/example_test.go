package lambdas_test

import (
	"fmt"

	"github.com/life4/genesis/lambdas"
	"github.com/life4/genesis/slices"
)

func ExampleAbs() {
	res := lambdas.Abs(-13)
	fmt.Println(res)
	//Output:
	// 13
}

func ExampleDefaultTo() {
	res := lambdas.DefaultTo(13)(slices.Min([]int{}))
	fmt.Println(res)
	//Output:
	// 13
}

func ExampleMax() {
	res := lambdas.Max(10, 13)
	fmt.Println(res)
	//Output:
	// 13
}

func ExampleMin() {
	res := lambdas.Min(15, 13)
	fmt.Println(res)
	//Output:
	// 13
}

func ExampleMust() {
	res := lambdas.Must(slices.Min([]int{42, 7, 13}))
	fmt.Println(res)
	//Output:
	// 7
}

func ExampleSafe() {
	res := lambdas.Safe(slices.Min([]int{}))
	fmt.Println(res)
	//Output:
	// 0
}
