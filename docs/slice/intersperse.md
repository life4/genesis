# Slice.Intersperse

```go
func (s Slice) Intersperse(el T) []T
```

Intersperse inserts el between each element of arr

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
// Intersperse inserts el between each element of arr
func (s Slice) Intersperse(el T) []T {
	if len(s.Data) == 0 {
		tmp := make([]T, 0)
		return tmp
	}
	result := make([]T, 0, len(s.Data)*2-1)
	result = append(result, s.Data[0])
	for _, val := range s.Data[1:] {
		result = append(result, el, val)
	}
	return result
}
```

## Tests

```go
func TestSliceIntersperse(t *testing.T) {
	f := func(el T, given []T, expected []T) {
		actual := Slice{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []T{}, []T{})
	f(0, []T{1}, []T{1})
	f(0, []T{1, 2}, []T{1, 0, 2})
	f(0, []T{1, 2, 3}, []T{1, 0, 2, 0, 3})
}
```