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
| Concat | Concat concatenates given slices into a single slice. |
| Product | Product returns cortesian product of elements {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4} |
| product | Product returns cortesian product of elements {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4} func (s Slices) Product() chan []T { 	c := make(chan []T, 1) 	go s.product(c, []T{}, 0) 	return c }  |
| Zip | Zip returns array of arrays of elements from given arrs on the same position |
