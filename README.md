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
gslices.Min([]int{42, 7, 13}) == 7
```

Double values in a slice of ints:

```go
gslices.Map([]int{4, 8, 15, 16, 23, 42}, func(el int) int { return el * 2 })
```

## Usage

Genesis contains the following packages:

+ [gslices](https://pkg.go.dev/github.com/life4/genesis/gslices): generic functions for slices.
+ [gaslices](https://pkg.go.dev/github.com/life4/genesis/gaslices): asynchronous versions of some slice functions. You specify how many goroutines to run, genesis does the magic.
+ [gchannels](https://pkg.go.dev/github.com/life4/genesis/gchannels): generic function for channels.
+ [glambdas](https://pkg.go.dev/github.com/life4/genesis/glambdas): helper generic functions to work with `gslices.Map` and similar.

![mascot image](./gopher.png)
