# Slice.StartsWith

```go
func (s Slice) StartsWith(prefix []T) bool
```

StartsWith returns true if slice starts with the given prefix slice. If prefix is empty, it returns true.

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
// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func (s Slice) StartsWith(prefix []T) bool {
	if len(prefix) > len(s.Data) {
		return false
	}
	for i, el := range prefix {
		if el != s.Data[i] {
			return false
		}
	}
	return true
}
```

## Tests

```go
func TestSliceStartsWith(t *testing.T) {
	f := func(given []T, suffix []T, expected bool) {
		actual := Slice{given}.StartsWith(suffix)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{}, true)
	f([]T{1}, []T{1}, true)
	f([]T{1}, []T{2}, false)
	f([]T{1, 2}, []T{1, 2, 3}, false)

	f([]T{1, 2, 3}, []T{1}, true)
	f([]T{1, 2, 3}, []T{1, 2}, true)
	f([]T{1, 2, 3}, []T{1, 2, 3}, true)

	f([]T{1, 2, 3}, []T{2}, false)
	f([]T{1, 2, 3}, []T{3}, false)
	f([]T{1, 2, 3}, []T{2, 3}, false)
	f([]T{1, 2, 3}, []T{3, 2}, false)
	f([]T{1, 2, 3}, []T{2, 1}, false)
}
```