# Slice

Slice is a set of operations to work with slice

```go
// Slice is a set of operations to work with slice
type Slice struct {
	data []T
}
```

| Function | Description |
| -------- | ----------- |
| Any | Any returns true if f returns true for any element in arr |
| All | All returns true if f returns true for all elements in arr |
| ChunkBy | ChunkBy splits arr on every element for which f returns a new value. |
| ChunkEvery | ChunkEvery returns slice of slices containing count elements each |
| Contains | Contains returns true if el in arr. |
| Count | Count return count of el occurences in arr. |
| Cycle | Cycle is an infinite loop over slice |
| Dedup | Dedup returns a given slice without consecutive duplicated elements |
| DedupBy | DedupBy returns a given slice without consecutive elements For which f returns the same result |
| DropEvery | DropEvery returns a slice of every nth element in the enumerable dropped, starting with the first element. |
| DropWhile | DropWhile drops elements from arr while f returns true |
| Each | Each calls f for every element from arr |
| Filter | Filter returns slice of T for which F returned true |
| Find | Find returns the first element for which f returns true |
| FindIndex | FindIndex is like Find, but return element index instead of element itself Returns -1 if element is not found |
| GroupBy | GroupBy groups element from array by value returned by f |
| Intersperse | Intersperse inserts el between each element of arr |
| Map | Map applies F to all elements in slice of T and returns slice of results |
| Max | Max returns the maximal element from arr |
| Min | Min returns the minimal element from arr |
| Permutations | Permutations returns successive size-length permutations of elements from the slice. {1, 2, 3} -> {1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2} |
| permutations | permutations is a core implementation for Permutations |
| Product | Product returns cortesian product of elements {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4} |
| product | product is a core implementation for Product |
| Reduce | Reduce applies F to acc and every element in slice of T and returns acc |
| ReduceWhile | ReduceWhile is like Reduce, but stops when f returns error |
| Reverse | Reverse returns given arr in reversed order |
| Same | Same returns true if all element in arr the same |
| Scan | Scan is like Reduce, but returns slice of f results |
| Shuffle | Shuffle in random order arr elements |
| Sort | Sort returns sorted slice |
| Sorted | Sorted returns true if slice is sorted |
| Split | Split splits arr by sep |
| StartsWith | StartsWith returns true if slice starts with the given prefix slice. If prefix is empty, it returns true. |
| Sum | Sum return sum of all elements from arr |
| TakeWhile | TakeWhile takes elements from arr while f returns true |
| Uniq | Uniq returns arr with only first occurences of every element. |
| Window | Window makes sliding window for a given slice: ({1,2,3}, 2) -> (1,2), (2,3) |
