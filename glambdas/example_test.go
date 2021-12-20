package glambdas_test

import (
	"fmt"

	"github.com/life4/genesis/glambdas"
	"github.com/life4/genesis/gslices"
)

func ExampleMust() {
	res := glambdas.Must(gslices.Min([]int{42, 7, 13}))
	fmt.Println(res)
	//Output:
	// 7
}

func ExampleDefaultTo() {
	res := glambdas.DefaultTo(13)(gslices.Min([]int{}))
	fmt.Println(res)
	//Output:
	// 13
}

func ExampleSafe() {
	res := glambdas.Safe(gslices.Min([]int{}))
	fmt.Println(res)
	//Output:
	// 0
}

func ExampleAbs() {
	res := glambdas.Abs(-13)
	fmt.Println(res)
	//Output:
	// 13
}

func ExampleMin() {
	res := glambdas.Min(15, 13)
	fmt.Println(res)
	//Output:
	// 13
}

func ExampleMax() {
	res := glambdas.Max(10, 13)
	fmt.Println(res)
	//Output:
	// 13
}
