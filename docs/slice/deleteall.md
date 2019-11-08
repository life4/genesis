# Slice.DeleteAll

```go
func (s Slice) DeleteAll(element T) []T
```

DeleteAll deletes all occurences of the element from the slice

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
// DeleteAll deletes all occurences of the element from the slice
func (s Slice) DeleteAll(element T) []T {
	if len(s.Data) == 0 {
		return s.Data
	}

	result := make([]T, 0, len(s.Data))
	for _, el := range s.Data {
		if el == element {
			continue
		}
		result = append(result, el)
	}
	return result

}
```

## Tests

```go
func TestSliceDeleteAll(t *testing.T) {
	f := func(given []T, el T, expected []T) {
		actual := Slice{given}.DeleteAll(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 1, []T{})
	f([]T{1}, 1, []T{})
	f([]T{2}, 1, []T{2})
	f([]T{1, 2}, 1, []T{2})
	f([]T{1, 2, 3}, 2, []T{1, 3})
	f([]T{1, 2, 2, 3, 2}, 2, []T{1, 3})
}
```