# Slice.DeleteAt

```go
func (s Slice) DeleteAt(indices ...int) []T
```

DeleteAt returns the slice without elements on given positions

Generic types: T.

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

## Source

```go
// DeleteAt returns the slice without elements on given positions
func (s Slice) DeleteAt(indices ...int) []T {
	result := make([]T, 0, len(s.Data)-len(indices))

	for i, el := range s.Data {
		allowed := true
		for _, j := range indices {
			if i == j {
				allowed = false
				break
			}
		}
		if allowed {
			result = append(result, el)
		}
	}
	return result
}
```

