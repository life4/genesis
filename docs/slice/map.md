# Slice.Map

```go
func (s Slice) Map(f func(el T) G) []G
```

Map applies F to all elements in slice of T and returns slice of results

Generic types: G, T.

## Examples

```go
s := []int{4, 8, 15, 16, 23, 42}
double := func(el int) int { return el * 2 }
doubled := genesis.SliceInt{s}.MapInt(double)
fmt.Println(doubled)
// Output: [8 16 30 32 46 84]
```

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
| MapBool | bool |
| MapByte | byte |
| MapString | string |
| MapRune | rune |
| MapError | error |
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
	result := make([]G, 0, len(s.Data))
	for _, el := range s.Data {
		result = append(result, f(el))
	}
	return result
}
```

## Tests

```go
func TestSliceMap(t *testing.T) {
	f := func(given []T, expected []G) {
		double := func(t T) G { return G((t * 2)) }
		actual := Slice{given}.Map(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []G{})
	f([]T{1}, []G{2})
	f([]T{1, 2, 3}, []G{2, 4, 6})
}
```