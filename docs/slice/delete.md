# Slice.Delete

```go
func (s Slice) Delete(element T) []T
```

Delete deletes the first occurence of the element from the slice

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
// Delete deletes the first occurence of the element from the slice
func (s Slice) Delete(element T) []T {
	result := make([]T, 0, len(s.Data)-1)
	deleted := false
	for _, el := range s.Data {
		if !deleted && el == element {
			continue
		}
		result = append(result, el)
	}
	return result

}
```

