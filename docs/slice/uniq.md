# Slice.Uniq

```go
func (s Slice) Uniq() []T
```

Uniq returns arr with only first occurrences of every element.

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
// Uniq returns arr with only first occurrences of every element.
func (s Slice) Uniq() []T {
	if len(s.Data) <= 1 {
		return s.Data
	}
	added := make(map[T]struct{})
	nothing := struct{}{}
	result := make([]T, 0, len(s.Data))
	for _, el := range s.Data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}
```

## Tests

```go
func TestSliceUniq(t *testing.T) {
	f := func(given []T, expected []T) {
		actual := Slice{given}.Uniq()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{1})
	f([]T{1, 1}, []T{1})
	f([]T{1, 2}, []T{1, 2})
	f([]T{1, 2, 1}, []T{1, 2})
	f([]T{1, 2, 1, 2}, []T{1, 2})
	f([]T{1, 2, 1, 2, 3, 2, 1, 1}, []T{1, 2, 3})
}
```