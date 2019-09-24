# Slice.Any

```go
func (s Slice) Any(f func(el T) bool) bool
```

Any returns true if f returns true for any element in arr

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
// Any returns true if f returns true for any element in arr
func (s Slice) Any(f func(el T) bool) bool {
	for _, el := range s.Data {
		if f(el) {
			return true
		}
	}
	return false
}
```

## Tests

```go
func TestSliceAny(t *testing.T) {
	f := func(check func(t T) bool, given []T, expected bool) {
		actual := Slice{given}.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t T) bool { return (t % 2) == 0 }

	f(isEven, []T{}, false)
	f(isEven, []T{1, 3}, false)
	f(isEven, []T{2}, true)
	f(isEven, []T{1, 2}, true)
}
```