# Slice.Sorted

```go
func (s Slice) Sorted() bool
```

Sorted returns true if slice is sorted

Generic types: T.

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
	if len(s.Data) <= 1 {
		return true
	}
	for i := 1; i < len(s.Data); i++ {
		if s.Data[i-1] > s.Data[i] {
			return false
		}
	}
	return true
}
```

