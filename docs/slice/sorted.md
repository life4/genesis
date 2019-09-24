# Slice.Sorted

Sorted returns true if slice is sorted

Generic types: .

## Structs

| Struct | T type |
| ------ | ------ |
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

## Source

```go
// Sorted returns true if slice is sorted
func (s Slice) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}
```

