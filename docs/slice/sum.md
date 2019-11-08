# Slice.Sum

```go
func (s Slice) Sum() T
```

Sum return sum of all elements from arr

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| SliceByte | byte |
| SliceString | string |
| SliceRune | rune |
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

## Source

```go
// Sum return sum of all elements from arr
func (s Slice) Sum() T {
	var sum T
	for _, el := range s.Data {
		sum += el
	}
	return sum
}
```

## Tests

```go
func TestSliceSum(t *testing.T) {
	f := func(given []T, expected T) {
		actual := Slice{given}.Sum()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 0)
	f([]T{1}, 1)
	f([]T{1, 2}, 3)
	f([]T{1, 2, 3}, 6)
}
```