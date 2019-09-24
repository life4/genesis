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
	f := func(filter func(t T) bool, given []T, expected []T) {
		actual := Slice{given}.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t T) bool { return t > 0 }

	f(filterPositive, []T{1, -1, 2, -2, 3, -3}, []T{1, 2, 3})
	f(filterPositive, []T{1, 2, 3}, []T{1, 2, 3})
	f(filterPositive, []T{-1, -2, -3}, []T{})
}
```