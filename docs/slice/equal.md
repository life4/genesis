# Slice.Equal

```go
func (s Slice) Equal(other []T) bool
```

Equal returns true if slices are equal

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
// Equal returns true if slices are equal
func (s Slice) Equal(other []T) bool {
	if len(s.Data) != len(other) {
		return false
	}
	for i, el := range other {
		if s.Data[i] != el {
			return false
		}
	}
	return true
}
```

## Tests

```go
func TestSliceEqual(t *testing.T) {
	f := func(left []T, right []T, expected bool) {
		actual := Slice{left}.Equal(right)
		assert.Equal(t, expected, actual, "they should be equal")

		actual = Slice{right}.Equal(left)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{}, true)
	f([]T{1}, []T{1}, true)
	f([]T{1}, []T{2}, false)
	f([]T{1, 2, 3, 3}, []T{1, 2, 3, 3}, true)
	f([]T{1, 2, 3, 3}, []T{1, 2, 2, 3}, false)
	f([]T{1, 2, 3, 3}, []T{1, 2, 4, 3}, false)

	// different len
	f([]T{1, 2, 3}, []T{1, 2}, false)
	f([]T{1, 2}, []T{1, 2, 3}, false)
	f([]T{}, []T{1, 2, 3}, false)
	f([]T{1, 2, 3}, []T{}, false)
}
```