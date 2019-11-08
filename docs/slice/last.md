# Slice.Last

```go
func (s Slice) Last() (T, error)
```

Last returns the last element from the slice

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
| ErrEmpty | container is empty |

## Source

```go
// Last returns the last element from the slice
func (s Slice) Last() (T, error) {
	if len(s.Data) == 0 {
		var tmp T
		return tmp, ErrEmpty
	}
	return s.Data[len(s.Data)-1], nil
}
```

## Tests

```go
func TestSliceLast(t *testing.T) {
	f := func(given []T, expectedEl T, expectedErr error) {
		el, err := Slice{given}.Last()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]T{}, 0, ErrEmpty)
	f([]T{1}, 1, nil)
	f([]T{1, 2, 3}, 3, nil)
}
```