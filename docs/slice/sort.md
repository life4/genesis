# Slice.Sort

```go
func (s Slice) Sort() []T
```

Sort returns sorted slice

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| SliceByte | byte |
| SliceString | string |
| SliceRune | rune |
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
	if len(s.Data) <= 1 {
		return s.Data
	}
	less := func(i int, j int) bool {
		return s.Data[i] < s.Data[j]
	}
	sort.SliceStable(s.Data, less)
	return s.Data
}
```

