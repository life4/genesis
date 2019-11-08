# Slice.Product

```go
func (s Slice) Product(repeat int) chan []T
```

Product returns cortesian product of elements {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4}

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
// Product returns cortesian product of elements
// {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4}
func (s Slice) Product(repeat int) chan []T {
	c := make(chan []T, 1)
	go func() {
		defer close(c)
		if repeat < 1 {
			return
		}
		s.product(c, repeat, []T{}, 0)
	}()
	return c
}
```

## Tests

```go
func TestSliceProduct(t *testing.T) {
	f := func(given []T, repeat int, expected [][]T) {
		actual := make([][]T, 0)
		i := 0
		s := Slice{given}
		for el := range s.Product(repeat) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]T{1, 2}, 0, [][]T{})
	f([]T{}, 2, [][]T{})
	f([]T{1}, 2, [][]T{{1, 1}})

	f([]T{1, 2}, 1, [][]T{{1}, {2}})
	f([]T{1, 2}, 2, [][]T{{1, 1}, {1, 2}, {2, 1}, {2, 2}})
	f([]T{1, 2}, 3, [][]T{
		{1, 1, 1}, {1, 1, 2}, {1, 2, 1}, {1, 2, 2},
		{2, 1, 1}, {2, 1, 2}, {2, 2, 1}, {2, 2, 2},
	})
}
```