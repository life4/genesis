# Slice.ReduceWhile

```go
func (s Slice) ReduceWhile(acc G, f func(el T, acc G) (G, error)) (G, error)
```

ReduceWhile is like Reduce, but stops when f returns error

Generic types: G, T.

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

## Functions

| Function | G type |
| -------- | ------ |
| ReduceWhileBool | bool |
| ReduceWhileByte | byte |
| ReduceWhileString | string |
| ReduceWhileRune | rune |
| ReduceWhileError | error |
| ReduceWhileFloat32 | float32 |
| ReduceWhileFloat64 | float64 |
| ReduceWhileInt | int |
| ReduceWhileInt8 | int8 |
| ReduceWhileInt16 | int16 |
| ReduceWhileInt32 | int32 |
| ReduceWhileInt64 | int64 |
| ReduceWhileUint | uint |
| ReduceWhileUint8 | uint8 |
| ReduceWhileUint16 | uint16 |
| ReduceWhileUint32 | uint32 |
| ReduceWhileUint64 | uint64 |
| ReduceWhileInterface | interface{} |

## Source

```go
// ReduceWhile is like Reduce, but stops when f returns error
func (s Slice) ReduceWhile(acc G, f func(el T, acc G) (G, error)) (G, error) {
	var err error
	for _, el := range s.Data {
		acc, err = f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}
```

## Tests

```go
func TestSliceReduceWhile(t *testing.T) {
	f := func(given []T, expected G) {
		sum := func(el T, acc G) (G, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return G(el) + acc, nil
		}
		actual, _ := Slice{given}.ReduceWhile(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 0)
	f([]T{1}, 1)
	f([]T{1, 2}, 3)
	f([]T{1, 2, 3}, 6)
	f([]T{1, 2, 0, 3}, 3)
}
```