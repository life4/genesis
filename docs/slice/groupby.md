# Slice.GroupBy

```go
func (s Slice) GroupBy(f func(el T) G) map[G][]T
```

GroupBy groups element from array by value returned by f

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
| GroupByBool | bool |
| GroupByByte | byte |
| GroupByString | string |
| GroupByRune | rune |
| GroupByError | error |
| GroupByFloat32 | float32 |
| GroupByFloat64 | float64 |
| GroupByInt | int |
| GroupByInt8 | int8 |
| GroupByInt16 | int16 |
| GroupByInt32 | int32 |
| GroupByInt64 | int64 |
| GroupByUint | uint |
| GroupByUint8 | uint8 |
| GroupByUint16 | uint16 |
| GroupByUint32 | uint32 |
| GroupByUint64 | uint64 |
| GroupByInterface | interface{} |

## Source

```go
// GroupBy groups element from array by value returned by f
func (s Slice) GroupBy(f func(el T) G) map[G][]T {
	result := make(map[G][]T)
	for _, el := range s.Data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]T, 1)
		}
		result[key] = append(val, el)
	}
	return result
}
```

## Tests

```go
func TestSliceGroupBy(t *testing.T) {
	f := func(given []T, expected map[G][]T) {
		reminder := func(t T) G { return G((t % 2)) }
		actual := Slice{given}.GroupBy(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, map[G][]T{})
	f([]T{1}, map[G][]T{1: {1}})
	f([]T{1, 3, 2, 4, 5}, map[G][]T{0: {2, 4}, 1: {1, 3, 5}})
}
```