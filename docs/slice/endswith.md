# Slice.EndsWith

```go
func (s Slice) EndsWith(suffix []T) bool
```

EndsWith returns true if slice ends with the given suffix slice. If suffix is empty, it returns true.

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
// EndsWith returns true if slice ends with the given suffix slice.
// If suffix is empty, it returns true.
func (s Slice) EndsWith(suffix []T) bool {
	if len(suffix) > len(s.Data) {
		return false
	}
	start := len(s.Data) - len(suffix)
	for i, el := range suffix {
		if el != s.Data[start+i] {
			return false
		}
	}
	return true
}
```

## Tests

```go
func TestSliceEndsWith(t *testing.T) {
	f := func(given []T, suffix []T, expected bool) {
		actual := Slice{given}.EndsWith(suffix)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{}, true)
	f([]T{1}, []T{1}, true)
	f([]T{1}, []T{2}, false)
	f([]T{2, 3}, []T{1, 2, 3}, false)

	f([]T{1, 2, 3}, []T{3}, true)
	f([]T{1, 2, 3}, []T{2, 3}, true)
	f([]T{1, 2, 3}, []T{1, 2, 3}, true)

	f([]T{1, 2, 3}, []T{1}, false)
	f([]T{1, 2, 3}, []T{2}, false)
	f([]T{1, 2, 3}, []T{1, 2}, false)
	f([]T{1, 2, 3}, []T{3, 2}, false)
}
```