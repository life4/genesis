# Slice.Window

```go
func (s Slice) Window(size int) ([][]T, error)
```

Window makes sliding window for a given slice: ({1,2,3}, 2) -> (1,2), (2,3)

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

## Errors

| Error | Message |
| -------- | ------ |
| ErrNonPositiveValue | value must be positive |

## Source

```go
// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func (s Slice) Window(size int) ([][]T, error) {
	if size <= 0 {
		return nil, ErrNonPositiveValue
	}
	result := make([][]T, 0, len(s.Data)/size)
	for i := 0; i <= len(s.Data)-size; i++ {
		chunk := s.Data[i : i+size]
		result = append(result, chunk)
	}
	return result, nil
}
```

## Tests

```go
func TestSliceWindow(t *testing.T) {
	f := func(given []T, size int, expected [][]T) {
		actual, _ := Slice{given}.Window(size)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 1, [][]T{})
	f([]T{1, 2, 3, 4}, 1, [][]T{{1}, {2}, {3}, {4}})
	f([]T{1, 2, 3, 4}, 2, [][]T{{1, 2}, {2, 3}, {3, 4}})
	f([]T{1, 2, 3, 4}, 3, [][]T{{1, 2, 3}, {2, 3, 4}})
	f([]T{1, 2, 3, 4}, 4, [][]T{{1, 2, 3, 4}})
}
```