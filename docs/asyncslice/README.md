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
| All | All returns true if f returns true for all elements in slice |
| Any | Any returns true if f returns true for any element from slice |
| Each | Each calls f for every element from slice |
| Filter | Filter returns slice of element for which f returns true |
| Map | Map applies F to all elements in slice of T and returns slice of results |
| Reduce | Reduce reduces slice to a single value with f |
