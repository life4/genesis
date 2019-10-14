# Slice.Reduce

```go
func (s Slice) Reduce(acc G, f func(el T, acc G) G) G
```

Reduce applies F to acc and every element in slice of T and returns acc

Generic types: G, T.

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

## Functions

| Function | G type |
| -------- | ------ |
| ReduceBool | bool |
| ReduceByte | byte |
| ReduceString | string |
| ReduceFloat32 | float32 |
| ReduceFloat64 | float64 |
| ReduceInt | int |
| ReduceInt8 | int8 |
| ReduceInt16 | int16 |
| ReduceInt32 | int32 |
| ReduceInt64 | int64 |
| ReduceUint | uint |
| ReduceUint8 | uint8 |
| ReduceUint16 | uint16 |
| ReduceUint32 | uint32 |
| ReduceUint64 | uint64 |
| ReduceInterface | interface{} |

## Source

```go
// Reduce applies F to acc and every element in slice of T and returns acc
func (s Slice) Reduce(acc G, f func(el T, acc G) G) G {
	for _, el := range s.Data {
		acc = f(el, acc)
	}
	return acc
}
```

## Tests

```go
func TestSliceReduce(t *testing.T) {
	f := func(given []T, expected G) {
		sum := func(el T, acc G) G { return G(el) + acc }
		actual := Slice{given}.Reduce(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 0)
	f([]T{1}, 1)
	f([]T{1, 2}, 3)
	f([]T{1, 2, 3}, 6)
}
```