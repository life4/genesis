# AsyncSlice

AsyncSlice is a set of operations to work with slice asynchronously

```go
// AsyncSlice is a set of operations to work with slice asynchronously
type AsyncSlice struct {
	data    []T
	workers int
}
```

| Function | Description |
| -------- | ----------- |
| [All](./all.md) | All returns true if f returns true for all elements in slice |
| [Any](./any.md) | Any returns true if f returns true for any element from slice |
| [Each](./each.md) | Each calls f for every element from slice |
| [Filter](./filter.md) | Filter returns slice of element for which f returns true |
| [Map](./map.md) | Map applies F to all elements in slice of T and returns slice of results |
| [Reduce](./reduce.md) | Reduce reduces slice to a single value with f |
