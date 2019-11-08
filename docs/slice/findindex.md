# Slice.FindIndex

```go
func (s Slice) FindIndex(f func(el T) bool) int
```

FindIndex is like Find, but return element index instead of element itself. Returns -1 if element not found

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
// FindIndex is like Find, but return element index instead of element itself.
// Returns -1 if element not found
func (s Slice) FindIndex(f func(el T) bool) int {
	for i, el := range s.Data {
		if f(el) {
			return i
		}
	}
	return -1
}
```

## Tests

```go
func TestSliceFindIndex(t *testing.T) {
	f := func(given []T, expectedInd int) {
		even := func(t T) bool { return (t % 2) == 0 }
		index := Slice{given}.FindIndex(even)
		assert.Equal(t, expectedInd, index, "they should be equal")
	}
	f([]T{}, -1)
	f([]T{1}, -1)
	f([]T{1}, -1)
	f([]T{2}, 0)
	f([]T{1, 2}, 1)
	f([]T{1, 2, 3}, 1)
	f([]T{1, 3, 5, 7, 9, 2}, 5)
	f([]T{1, 3, 5}, -1)
}
```