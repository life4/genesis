# Genesis

Generic functions for Go. Bringing the beauty of functional programming in Go 1.18+.

**😎 Features:**

+ 🛠️ Over 130 generic functions for channels, maps, and slices.
+ 💪 Uses the power of Go 1.18+ generics.
+ 🧐 No code generation.
+ 🪶 No dependencies (except [is](https://github.com/matryer/is) for testing).
+ 🏃 Pure Go.
+ 🪩 Sync and async versions of all the main functions.

**🔨 When to use:**

+ 🐘 **In a big project**. More the project grows, more you find yourself writing boring generic code like "Min". Break the cycle.
+ 🤝 **In a team project**. Each line of code you write means higher maintenance cost that in turn means loosing time and money.
+ 🐶 **In a pet project**. Leave the boring stuff to us, focus on the fun parts.
+ 📚 **When readability matters**. `slices.Shrink` is a function with a human-friendly name and documentation. `s[:len(s):len(s)]` is a jibberish and black magic. Prefer the former.
+ 💔 **When you miss some conveniences** that come in other languages out-of-the-box.
+ 🐇 **When you write a highly concurrent code** and don't want to manually implement code for cancellation, results collection and ordering, worker groups, context, etc.

**📦 What's inside**:

+ `Filter`, `Map`, and `Reduce` for data processing on steroids.
+ `FilterAsync`, `MapAsync`, and `ReduceAsync` for making your code fast and concurrent with a single line of code.
+ `Grow` and `Shrink` for reducing memory allocations.
+ `Permutations` and `Product` for simple iterations.
+ `Shuffle` and `Sort` for randomization.
+ `Any` and `All` for simple flow control.
+ `Range`, `Count`, and `Cycle` for generating sequences.

And much more.

## 💾 Installation

```bash
go get github.com/life4/genesis
```

## 👀 Examples

Find the minimal value in a slice of ints:

```go
lambdas.Must(slices.Min([]int{42, 7, 13})) == 7
```

Double values in a slice of ints:

```go
slices.Map([]int{4, 8, 15}, func(el int) int { return el * 2 })
```

Concurrently check status codes for multiple URLs:

```go
urls := []string{
    "https://go.dev/",
    "https://golang.org/",
    "https://google.com/",
}
codes := slices.MapAsync(
    urls, 0,
    func(url string) int {
        return lambdas.Must(http.Get(url)).StatusCode
    },
)
```

## 🔨 Usage

Genesis contains the following packages:

+ [🍞 slices](https://pkg.go.dev/github.com/life4/genesis/slices): generic functions for slices.
+ [🗺 maps](https://pkg.go.dev/github.com/life4/genesis/maps): generic functions for maps.
+ [📺 channels](https://pkg.go.dev/github.com/life4/genesis/channels): generic function for channels.
+ [🛟 lambdas](https://pkg.go.dev/github.com/life4/genesis/lambdas): helper generic functions to work with `slices.Map` and similar.

See [📄 DOCUMENTATION](https://pkg.go.dev/github.com/life4/genesis) for more info.
