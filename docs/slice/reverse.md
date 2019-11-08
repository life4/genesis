# Slice.Reverse

```go
func (s Slice) Reverse() []T
```

Reverse returns given arr in reversed order

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| SliceBool | bool |
| SliceByte | byte |
| SliceString | string |
| SliceRune | rune |
| SliceError | error |
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
// Reverse returns given arr in reversed order
func (s Slice) Reverse() []T {
	if len(s.Data) <= 1 {
		return s.Data
	}
	result := make([]T, 0, len(s.Data))
	for i := len(s.Data) - 1; i >= 0; i-- {
		result = append(result, s.Data[i])
	}
	return result
}
```

## Tests

```go
func TestSliceReverse(t *testing.T) {
	f := func(given []T, expected []T) {
		actual := Slice{given}.Reverse()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{1})
	f([]T{1, 2}, []T{2, 1})
	f([]T{1, 2, 3}, []T{3, 2, 1})
	f([]T{1, 2, 2, 3, 3}, []T{3, 3, 2, 2, 1})
}
```