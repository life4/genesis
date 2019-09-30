# Slice.DropWhile

```go
func (s Slice) DropWhile(f func(el T) bool) []T
```

DropWhile drops elements from arr while f returns true

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
// DropWhile drops elements from arr while f returns true
func (s Slice) DropWhile(f func(el T) bool) []T {
	for i, el := range s.Data {
		if !f(el) {
			return s.Data[i:]
		}
	}
	return []T{}
}
```

## Tests

```go
func TestSliceDropWhile(t *testing.T) {
	f := func(given []T, expected []T) {
		even := func(el T) bool { return el%2 == 0 }
		actual := Slice{given}.DropWhile(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{2}, []T{})
	f([]T{1}, []T{1})
	f([]T{2, 1}, []T{1})
	f([]T{2, 1, 2}, []T{1, 2})
	f([]T{1, 2}, []T{1, 2})
	f([]T{2, 4, 6, 1, 8}, []T{1, 8})
}
```