# Slice.Equal

```go
func (s Slice) Equal(other []T) bool
```

Equal returns true if slices are equal

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
// Equal returns true if slices are equal
func (s Slice) Equal(other []T) bool {
	if len(s.Data) != len(other) {
		return false
	}
	for i, el := range other {
		if s.Data[i] != el {
			return false
		}
	}
	return true
}
```

