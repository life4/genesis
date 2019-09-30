# Slice.DeleteAt

```go
func (s Slice) DeleteAt(indices ...int) ([]T, error)
```

DeleteAt returns the slice without elements on given positions

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

## Errors

| Error | Message |
| -------- | ------ |
| ErrOutOfRange | index is bigger than container size |

## Source

```go
// DeleteAt returns the slice without elements on given positions
func (s Slice) DeleteAt(indices ...int) ([]T, error) {
	if len(indices) == 0 || len(s.Data) == 0 {
		return s.Data, nil
	}

	for _, index := range indices {
		if index >= len(s.Data) {
			return s.Data, ErrOutOfRange
		}
	}

	result := make([]T, 0, len(s.Data)-1)
	for i, el := range s.Data {
		add := true
		for _, index := range indices {
			if i == index {
				add = false
				break
			}
		}
		if add {
			result = append(result, el)
		}
	}
	return result, nil
}
```

## Tests

```go
func TestSliceDeleteAt(t *testing.T) {
	f := func(given []T, indices []int, expected []T) {
		actual, _ := Slice{given}.DeleteAt(indices...)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []int{}, []T{})
	f([]T{1}, []int{0}, []T{})
	f([]T{1, 2}, []int{0}, []T{2})

	f([]T{1, 2, 3}, []int{0}, []T{2, 3})
	f([]T{1, 2, 3}, []int{1}, []T{1, 3})
	f([]T{1, 2, 3}, []int{2}, []T{1, 2})

	f([]T{1, 2, 3}, []int{0, 1}, []T{3})
	f([]T{1, 2, 3}, []int{0, 2}, []T{2})
	f([]T{1, 2, 3}, []int{1, 2}, []T{1})
}
```