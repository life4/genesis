# Slice.Filter

```go
func (s Slice) Filter(f func(el T) bool) []T
```

Filter returns slice of T for which F returned true

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
// Filter returns slice of T for which F returned true
func (s Slice) Filter(f func(el T) bool) []T {
	result := make([]T, 0, len(s.Data))
	for _, el := range s.Data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}
```

## Tests

```go
func TestSliceFilter(t *testing.T) {
	f := func(given []T, expected []T) {
		even := func(t T) bool { return (t % 2) == 0 }
		actual := Slice{given}.Filter(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1, 2, 3, 4}, []T{2, 4})
	f([]T{1, 3}, []T{})
	f([]T{2, 4}, []T{2, 4})
}
```