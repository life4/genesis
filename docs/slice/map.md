# Slice.Map

```go
func (s Slice) Map(f func(el T) G) []G
```

Map applies F to all elements in slice of T and returns slice of results

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
| MapBool | bool |
| MapByte | byte |
| MapString | string |
| MapFloat32 | float32 |
| MapFloat64 | float64 |
| MapInt | int |
| MapInt8 | int8 |
| MapInt16 | int16 |
| MapInt32 | int32 |
| MapInt64 | int64 |
| MapUint | uint |
| MapUint8 | uint8 |
| MapUint16 | uint16 |
| MapUint32 | uint32 |
| MapUint64 | uint64 |
| MapInterface | interface{} |

## Source

```go
// Map applies F to all elements in slice of T and returns slice of results
func (s Slice) Map(f func(el T) G) []G {
	result := make([]G, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}
```

## Tests

```go
func TestSliceMap(t *testing.T) {
	f := func(mapper func(t T) G, given []T, expected []G) {
		actual := Slice{given}.Map(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t T) G { return G((t * 2)) }

	f(double, []T{1, 2, 3}, []G{2, 4, 6})
}
```