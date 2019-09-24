# Slice.Dedup

Dedup returns a given slice without consecutive duplicated elements

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
// Dedup returns a given slice without consecutive duplicated elements
func (s Slice) Dedup() []T {
	result := make([]T, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}
```

