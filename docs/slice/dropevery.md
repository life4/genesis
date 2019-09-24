# Slice.DropEvery

DropEvery returns a slice of every nth element in the enumerable dropped, starting with the first element.

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
// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func (s Slice) DropEvery(nth int) []T {
	result := make([]T, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}
```

