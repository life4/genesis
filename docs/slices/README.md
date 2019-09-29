# Slices

Slices is a set of operations to work with slice of slices

```go
// Slices is a set of operations to work with slice of slices
type Slices struct {
	Data [][]T
}
```

## Structs

| Struct | T type |
| ------ | ------ |
| SlicesBool | bool |
| SlicesByte | byte |
| SlicesString | string |
| SlicesFloat32 | float32 |
| SlicesFloat64 | float64 |
| SlicesInt | int |
| SlicesInt8 | int8 |
| SlicesInt16 | int16 |
| SlicesInt32 | int32 |
| SlicesInt64 | int64 |
| SlicesUint | uint |
| SlicesUint8 | uint8 |
| SlicesUint16 | uint16 |
| SlicesUint32 | uint32 |
| SlicesUint64 | uint64 |
| SlicesInterface | interface{} |

## Functions

| Function | Description |
| -------- | ----------- |
| [Concat](./concat.md) | Concat concatenates given slices into a single slice. |
| [Product](./product.md) | Product returns cortesian product of elements {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4} |
| [Zip](./zip.md) | Zip returns array of arrays of elements from given arrs on the same position |
