# Slice.ReplaceAt

```go
func (s Slice) ReplaceAt(index int, element T) []T
```

ReplaceAt returns the slice with element replaced at given index.

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
// ReplaceAt returns the slice with element replaced at given index.
func (s Slice) ReplaceAt(index int, element T) []T {
	result := make([]T, 0, len(s.Data)+1)
	for i, el := range s.Data {
		if i == index {
			result = append(result, element)
		} else {
			result = append(result, el)
		}
	}
	return result
}
```

