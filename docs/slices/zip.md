# Slices.Zip

Zip returns array of arrays of elements from given arrs on the same position

Generic types: T.

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

## Source

```go
// Zip returns array of arrays of elements from given arrs
// on the same position
func (s Slices) Zip() [][]T {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]T, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]T, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}
```

