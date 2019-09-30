# Slice.EndsWith

```go
func (s Slice) EndsWith(suffix []T) bool
```

EndsWith returns true if slice ends with the given suffix slice. If suffix is empty, it returns true.

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
// EndsWith returns true if slice ends with the given suffix slice.
// If suffix is empty, it returns true.
func (s Slice) EndsWith(suffix []T) bool {
	if len(suffix) > len(s.Data) {
		return false
	}
	for i, el := range suffix {
		if el != s.Data[len(s.Data)-i] {
			return false
		}
	}
	return true
}
```

