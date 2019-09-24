# Slice.Window

Window makes sliding window for a given slice: ({1,2,3}, 2) -> (1,2), (2,3)

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
// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func (s Slice) Window(size int) [][]T {
	result := make([][]T, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}
```