# Slice.TakeWhile

```go
func (s Slice) TakeWhile(f func(el T) bool) []T
```

TakeWhile takes elements from arr while f returns true

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
// TakeWhile takes elements from arr while f returns true
func (s Slice) TakeWhile(f func(el T) bool) []T {
	result := make([]T, 0, len(s.Data))
	for _, el := range s.Data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}
```

## Tests

```go
func TestSliceTakeWhile(t *testing.T) {
	f := func(given []T, expected []T) {
		even := func(el T) bool { return el%2 == 0 }
		actual := Slice{given}.TakeWhile(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{})
	f([]T{2}, []T{2})
	f([]T{2, 4, 6, 1, 8}, []T{2, 4, 6})
	f([]T{1, 2, 3}, []T{})
}
```