# Slice.Find

```go
func (s Slice) Find(f func(el T) bool) (T, error)
```

Find returns the first element for which f returns true

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
// Find returns the first element for which f returns true
func (s Slice) Find(f func(el T) bool) (T, error) {
	for _, el := range s.Data {
		if f(el) {
			return el, nil
		}
	}
	var tmp T
	return tmp, ErrNotFound
}
```

## Tests

```go
func TestSliceFind(t *testing.T) {
	f := func(given []T, expectedEl T, expectedErr error) {
		even := func(t T) bool { return (t % 2) == 0 }
		el, err := Slice{given}.Find(even)
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]T{}, 0, ErrNotFound)
	f([]T{1}, 0, ErrNotFound)
	f([]T{1}, 0, ErrNotFound)
	f([]T{2}, 2, nil)
	f([]T{1, 2}, 2, nil)
	f([]T{1, 2, 3}, 2, nil)
	f([]T{1, 3, 5}, 0, ErrNotFound)
}
```