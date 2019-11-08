# Slice.Without

```go
func (s Slice) Without(elements ...T) []T
```

Without returns the slice with filtered out element

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| SliceBool | bool |
| SliceByte | byte |
| SliceString | string |
| SliceRune | rune |
| SliceError | error |
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
// Without returns the slice with filtered out element
func (s Slice) Without(elements ...T) []T {
	result := make([]T, 0, len(s.Data))
	for _, el := range s.Data {
		allowed := true
		for _, other := range elements {
			if el == other {
				allowed = false
			}
		}
		if allowed {
			result = append(result, el)
		}
	}
	return result
}
```

## Tests

```go
func TestSliceWithout(t *testing.T) {
	f := func(given []T, items []T, expected []T) {
		actual := Slice{given}.Without(items...)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{}, []T{})
	f([]T{}, []T{1, 2}, []T{})

	f([]T{1}, []T{1}, []T{})
	f([]T{1}, []T{1, 2}, []T{})
	f([]T{1}, []T{2}, []T{1})

	f([]T{1, 2, 3, 4}, []T{1}, []T{2, 3, 4})
	f([]T{1, 2, 3, 4}, []T{2}, []T{1, 3, 4})
	f([]T{1, 2, 3, 4}, []T{4}, []T{1, 2, 3})

	f([]T{1, 2, 3, 4}, []T{1, 2}, []T{3, 4})
	f([]T{1, 2, 3, 4}, []T{1, 2, 4}, []T{3})
	f([]T{1, 2, 3, 4}, []T{1, 2, 3, 4}, []T{})
	f([]T{1, 2, 3, 4}, []T{2, 4}, []T{1, 3})

	f([]T{1, 1, 2, 3, 1, 4, 1}, []T{1}, []T{2, 3, 4})
}
```