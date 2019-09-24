# Slice.Reverse

Reverse returns given arr in reversed order

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
// Reverse returns given arr in reversed order
func (s Slice) Reverse() []T {
	result := make([]T, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}
```

