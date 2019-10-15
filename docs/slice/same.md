# Slice.Same

```go
func (s Slice) Same() bool
```

Same returns true if all element in arr the same

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
// Same returns true if all element in arr the same
func (s Slice) Same() bool {
	if len(s.Data) <= 1 {
		return true
	}
	for i := 0; i < len(s.Data)-1; i++ {
		if s.Data[i] != s.Data[i+1] {
			return false
		}
	}
	return true
}
```

## Tests

```go
func TestSliceSame(t *testing.T) {
	f := func(given []T, expected bool) {
		actual := Slice{given}.Same()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, true)
	f([]T{1}, true)
	f([]T{1, 1}, true)
	f([]T{1, 1, 1}, true)

	f([]T{1, 2, 1}, false)
	f([]T{1, 2, 2}, false)
	f([]T{1, 1, 2}, false)
}
```