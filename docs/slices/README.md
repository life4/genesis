# Slices

Slices is a set of operations to work with slice of slices

```go
// Slices is a set of operations to work with slice of slices
type Slices struct {
	data [][]T
}
```

| Function | Description |
| -------- | ----------- |
| [Concat](./concat.md) | Concat concatenates given slices into a single slice. |
| [Product](./product.md) | Product returns cortesian product of elements {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4} |
| [Zip](./zip.md) | Zip returns array of arrays of elements from given arrs on the same position |
