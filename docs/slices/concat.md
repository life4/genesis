# Slices.Concat

```go
func (s Slices) Concat() []T
```

Concat concatenates given slices into a single slice.

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| SlicesBool | bool |
| SlicesByte | byte |
| SlicesString | string |
| SlicesRune | rune |
| SlicesError | error |
| SlicesFloat32 | float32 |
| SlicesFloat64 | float64 |
| SlicesInt | int |
| SlicesInt8 | int8 |
| SlicesInt16 | int16 |
| SlicesInt32 | int32 |
| SlicesInt64 | int64 |
| SlicesUint | uint |
| SlicesUint8 | uint8 |
| SlicesUint16 | uint16 |
| SlicesUint32 | uint32 |
| SlicesUint64 | uint64 |
| SlicesInterface | interface{} |

## Source

```go
// Concat concatenates given slices into a single slice.
func (s Slices) Concat() []T {
	result := make([]T, 0)
	for _, arr := range s.Data {
		result = append(result, arr...)
	}
	return result
}
```

## Tests

```go
func TestSlicesConcat(t *testing.T) {
	f := func(given [][]T, expected []T) {
		actual := Slices{given}.Concat()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]T{}, []T{})
	f([][]T{{}}, []T{})
	f([][]T{{1}}, []T{1})
	f([][]T{{1}, {}}, []T{1})
	f([][]T{{}, {1}}, []T{1})
	f([][]T{{1, 2}, {3, 4, 5}}, []T{1, 2, 3, 4, 5})
}
```