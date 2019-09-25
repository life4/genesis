# Slice

Slice is a set of operations to work with slice

```go
// Slice is a set of operations to work with slice
type Slice struct {
	Data []T
}
```

## Structs

| Struct | T type |
| ------ | ------ |
| SliceBool | bool |
| SliceByte | byte |
| SliceString | string |
| SliceFloat32 | float32 |
| SliceFloat64 | float64 |
| SliceInt | int |
| SliceInt8 | int8 |
| SliceInt16 | int16 |
| SliceInt32 | int32 |
| SliceInt64 | int64 |
| SliceUint | uint |
| SliceUint8 | uint8 |
| SliceUint16 | uint16 |
| SliceUint32 | uint32 |
| SliceUint64 | uint64 |
| SliceInterface | interface{} |

## Functions

| Function | Description |
| -------- | ----------- |
| [Any](./any.md) | Any returns true if f returns true for any element in arr |
| [All](./all.md) | All returns true if f returns true for all elements in arr |
| [ChunkBy](./chunkby.md) | ChunkBy splits arr on every element for which f returns a new value. |
| [ChunkEvery](./chunkevery.md) | ChunkEvery returns slice of slices containing count elements each |
| [Contains](./contains.md) | Contains returns true if el in arr. |
| [Count](./count.md) | Count return count of el occurences in arr. |
| [Cycle](./cycle.md) | Cycle is an infinite loop over slice |
| [Dedup](./dedup.md) | Dedup returns a given slice without consecutive duplicated elements |
| [DedupBy](./dedupby.md) | DedupBy returns a given slice without consecutive elements For which f returns the same result |
| [DropEvery](./dropevery.md) | DropEvery returns a slice of every nth element in the enumerable dropped, starting with the first element. |
| [DropWhile](./dropwhile.md) | DropWhile drops elements from arr while f returns true |
| [Each](./each.md) | Each calls f for every element from arr |
| [Equal](./equal.md) | Equal returns true if slices are equal |
| [Filter](./filter.md) | Filter returns slice of T for which F returned true |
| [Find](./find.md) | Find returns the first element for which f returns true |
| [FindIndex](./findindex.md) | FindIndex is like Find, but return element index instead of element itself Returns -1 if element is not found |
| [Join](./join.md) | Join concatenates elements of the slice to create a single string. |
| [GroupBy](./groupby.md) | GroupBy groups element from array by value returned by f |
| [Intersperse](./intersperse.md) | Intersperse inserts el between each element of arr |
| [Map](./map.md) | Map applies F to all elements in slice of T and returns slice of results |
| [Max](./max.md) | Max returns the maximal element from arr |
| [Min](./min.md) | Min returns the minimal element from arr |
| [Permutations](./permutations.md) | Permutations returns successive size-length permutations of elements from the slice. {1, 2, 3} -> {1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2} |
| [Product](./product.md) | Product returns cortesian product of elements {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4} |
| [Reduce](./reduce.md) | Reduce applies F to acc and every element in slice of T and returns acc |
| [ReduceWhile](./reducewhile.md) | ReduceWhile is like Reduce, but stops when f returns error |
| [Reverse](./reverse.md) | Reverse returns given arr in reversed order |
| [Same](./same.md) | Same returns true if all element in arr the same |
| [Scan](./scan.md) | Scan is like Reduce, but returns slice of f results |
| [Shuffle](./shuffle.md) | Shuffle in random order arr elements |
| [Sort](./sort.md) | Sort returns sorted slice |
| [Sorted](./sorted.md) | Sorted returns true if slice is sorted |
| [Split](./split.md) | Split splits arr by sep |
| [StartsWith](./startswith.md) | StartsWith returns true if slice starts with the given prefix slice. If prefix is empty, it returns true. |
| [Sum](./sum.md) | Sum return sum of all elements from arr |
| [TakeWhile](./takewhile.md) | TakeWhile takes elements from arr while f returns true |
| [ToChannel](./tochannel.md) | ToChannel returns channel with elements from the slice |
| [Uniq](./uniq.md) | Uniq returns arr with only first occurences of every element. |
| [Window](./window.md) | Window makes sliding window for a given slice: ({1,2,3}, 2) -> (1,2), (2,3) |
| [Without](./without.md) | Without returns the slice with filtered out element |
