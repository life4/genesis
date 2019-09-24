# Slice.GroupBy

GroupBy groups element from array by value returned by f

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


## Source

```go
// GroupBy groups element from array by value returned by f
func (s Slice) GroupBy(f func(el T) G) map[G][]T {
	result := make(map[G][]T)
	for _, el := range s.data {
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
	f := func(mapper func(t T) G, given []T, expected map[G][]T) {
		actual := Slice{given}.GroupBy(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t T) G { return G((t % 2)) }

	f(remainder, []T{}, map[G][]T{})
	f(remainder, []T{1}, map[G][]T{1: {1}})
	f(remainder, []T{1, 3, 2, 4, 5}, map[G][]T{0: {2, 4}, 1: {1, 3, 5}})
}
```