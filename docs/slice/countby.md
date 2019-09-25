# Slice.CountBy

```go
func (s Slice) CountBy(f func(el T) bool) int
```

CountBy returns how many times f returns true.

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
// CountBy returns how many times f returns true.
func (s Slice) CountBy(f func(el T) bool) int {
	count := 0
	for _, el := range s.Data {
		if f(el) {
			count++
		}
	}
	return count
}
```

## Tests

```go
func TestSliceCountBy(t *testing.T) {
	f := func(given []T, expected int) {
		even := func(t T) bool { return (t % 2) == 0 }
		actual := Slice{given}.CountBy(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 0)
	f([]T{1}, 0)
	f([]T{2}, 1)
	f([]T{1, 2, 3, 4, 5}, 2)
	f([]T{1, 2, 3, 4, 5, 6}, 3)
}
```