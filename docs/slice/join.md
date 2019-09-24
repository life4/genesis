# Slice.Join

```go
func (s Slice) Join(sep string) string
```

Join concatenates elements of the slice to create a single string.

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| SliceByte | byte |
| SliceString | string |
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
// Join concatenates elements of the slice to create a single string.
func (s Slice) Join(sep string) string {
	strs := make([]string, 0, len(s.data))
	for _, el := range s.data {
		strs = append(strs, string(el))
	}
	return strings.Join(strs, sep)
}
```

