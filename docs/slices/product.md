# Slices.Product

```go
func (s Slices) Product() chan []T
```

Product returns cortesian product of elements {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4}

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| SlicesBool | bool |
| SlicesByte | byte |
| SlicesString | string |
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
// Product returns cortesian product of elements
// {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4}
func (s Slices) Product() chan []T {
	c := make(chan []T, 1)
	go s.product(c, []T{}, 0)
	return c
}
```

## Tests

```go
func TestSlicesProduct(t *testing.T) {
	f := func(given [][]T, expected [][]T) {
		actual := make([][]T, 0)
		i := 0
		s := Slices{given}
		for el := range s.Product() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]T{{1, 2}, {3, 4}}, [][]T{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]T{{1, 2}, {3}, {4, 5}}, [][]T{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}
```