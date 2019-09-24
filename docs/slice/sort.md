# Slice.Sort

Sort returns sorted slice

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
// Sort returns sorted slice
func (s Slice) Sort() []T {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}
```