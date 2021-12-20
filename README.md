# Genesis

[![Build Status](https://travis-ci.org/life4/genesis.svg?branch=master)](https://travis-ci.org/life4/genesis)

Generic functions for Go. Bringing the beauty of functional programming in Go 1.18+.

**What's inside:**

+ Filter, Map, Reduce.
+ Min, Max, Sum.
+ Permutations, Product.
+ Any, All.
+ Contains, Find.
+ Shuffle, Sort.
+ Range, Count, Cycle.

And much more.

**Features:**

+ Uses the power of Go 1.18+ generics.
+ No code generation.
+ Sync and async versions.
+ For slices and channels.

```bash
go get github.com/life4/genesis
```

## Examples

Find the minimal value in a slice of ints:

```go
lambdas.Must(slices.Min([]int{42, 7, 13})) == 7
```

Double values in a slice of ints:

```go
slices.Map([]int{4, 8, 15}, func(el int) int { return el * 2 })
```

## Usage

Genesis contains the following packages:

+ [slices](https://pkg.go.dev/github.com/life4/genesis/slices): generic functions for slices.
+ [aslices](https://pkg.go.dev/github.com/life4/genesis/aslices): asynchronous versions of some slice functions. You specify how many goroutines to run, genesis does the magic.
+ [channels](https://pkg.go.dev/github.com/life4/genesis/channels): generic function for channels.
+ [lambdas](https://pkg.go.dev/github.com/life4/genesis/lambdas): helper generic functions to work with `slices.Map` and similar.

See [documentation](https://pkg.go.dev/github.com/life4/genesis) for more information.

![mascot image](./gopher.png)
