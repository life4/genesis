# Slice.Split

```go
func (s Slice) Split(sep T) [][]T
```

Split splits arr by sep

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
// Split splits arr by sep
func (s Slice) Split(sep T) [][]T {
	result := make([][]T, 0)
	curr := make([]T, 0)
	for _, el := range s.Data {
		if el == sep {
			result = append(result, curr)
			curr = make([]T, 0)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}
```

## Tests

```go
func TestSliceSplit(t *testing.T) {
	f := func(given []T, sep T, expected [][]T) {
		actual := Slice{given}.Split(sep)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 1, [][]T{{}})
	f([]T{2}, 1, [][]T{{2}})
	f([]T{2, 1, 3}, 1, [][]T{{2}, {3}})
	f([]T{1, 3}, 1, [][]T{{}, {3}})
	f([]T{2, 1}, 1, [][]T{{2}, {}})
	f([]T{2, 1, 3, 4, 1, 5, 6, 7}, 1, [][]T{{2}, {3, 4}, {5, 6, 7}})
}
```