# Slice.All

All returns true if f returns true for all elements in arr

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
// All returns true if f returns true for all elements in arr
func (s Slice) All(f func(el T) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}
```

## Tests

```go
func TestSliceAll(t *testing.T) {
	f := func(check func(t T) bool, given []T, expected bool) {
		actual := Slice{given}.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t T) bool { return (t % 2) == 0 }

	f(isEven, []T{}, true)
	f(isEven, []T{2}, true)
	f(isEven, []T{1}, false)
	f(isEven, []T{2, 4}, true)
	f(isEven, []T{2, 4, 1}, false)
	f(isEven, []T{1, 2, 4}, false)
}
```