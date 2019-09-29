# Genesis

Missed golang generic stdlib for slices and channels.

**Some functions:**

+ Filter, Map, Reduce.
+ Min, Max, Sum.
+ Permutations, Product.
+ Any, All.
+ Contains, Find.
+ Shuffle, Sort.
+ Range, Count, Cycle.

And much more.

**Features:**

+ Typesafe.
+ Sync and async versions.
+ For slices and channels.
+ Pre-generated for all built-in types.

```bash
go get github.com/life4/genesis
```

## Examples

Find minimal value in a slice of ints:

```go
s := []int{42, 7, 13}
min := genesis.SliceInt{s}.Min()
```

Double values in a slice of ints:

```go
s := []int{4, 8, 15, 16, 23, 42}
double := func(el int) int { return el * 2 }
doubled := genesis.SliceInt{s}.MapInt(double)
```

See [docs](./docs) to dive deeper.

## Generation

Install requirements

```bash
python3 -m pip install --user -r requirements.txt
```

Re-generate everything for built-in types:

```bash
python3 -m generate
```

Generate a new package with given types:

```bash
python3 -m generate
```
