# Slice.DedupBy

```go
func (s Slice) DedupBy(f func(el T) G) []T
```

DedupBy returns a given slice without consecutive elements For which f returns the same result

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
| DedupByBool | bool |
| DedupByByte | byte |
| DedupByString | string |
| DedupByRune | rune |
| DedupByError | error |
| DedupByFloat32 | float32 |
| DedupByFloat64 | float64 |
| DedupByInt | int |
| DedupByInt8 | int8 |
| DedupByInt16 | int16 |
| DedupByInt32 | int32 |
| DedupByInt64 | int64 |
| DedupByUint | uint |
| DedupByUint8 | uint8 |
| DedupByUint16 | uint16 |
| DedupByUint32 | uint32 |
| DedupByUint64 | uint64 |
| DedupByInterface | interface{} |

## Source

```go
// DedupBy returns a given slice without consecutive elements
// For which f returns the same result
func (s Slice) DedupBy(f func(el T) G) []T {
	result := make([]T, 0, len(s.Data))
	if len(s.Data) == 0 {
		return result
	}

	prev := f(s.Data[0])
	result = append(result, s.Data[0])
	for _, el := range s.Data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}
```

## Tests

```go
func TestSliceDedupBy(t *testing.T) {
	f := func(given []T, expected []T) {
		even := func(el T) G { return G(el % 2) }
		actual := Slice{given}.DedupBy(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{1})
	f([]T{1, 1}, []T{1})
	f([]T{1, 2}, []T{1, 2})
	f([]T{1, 2, 3}, []T{1, 2, 3})
	f([]T{1, 2, 2, 3}, []T{1, 2, 3})
	f([]T{1, 2, 4, 3, 5, 7, 10}, []T{1, 2, 3, 10})
}
```