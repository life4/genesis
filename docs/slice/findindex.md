# Slice.FindIndex

```go
func (s Slice) FindIndex(f func(el T) bool) (int, error)
```

FindIndex is like Find, but return element index instead of element itself

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
| ErrNotFound | given element is not found |

## Source

```go
// FindIndex is like Find, but return element index instead of element itself
func (s Slice) FindIndex(f func(el T) bool) (int, error) {
	for i, el := range s.Data {
		if f(el) {
			return i, nil
		}
	}
	return 0, ErrNotFound
}
```

## Tests

```go
func TestSliceFindIndex(t *testing.T) {
	f := func(given []T, expectedInd int, expectedErr error) {
		even := func(t T) bool { return (t % 2) == 0 }
		index, err := Slice{given}.FindIndex(even)
		assert.Equal(t, expectedInd, index, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]T{}, 0, ErrNotFound)
	f([]T{1}, 0, ErrNotFound)
	f([]T{1}, 0, ErrNotFound)
	f([]T{2}, 0, nil)
	f([]T{1, 2}, 1, nil)
	f([]T{1, 2, 3}, 1, nil)
	f([]T{1, 3, 5, 7, 9, 2}, 5, nil)
	f([]T{1, 3, 5}, 0, ErrNotFound)
}
```