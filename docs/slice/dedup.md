# Slice.Dedup

```go
func (s Slice) Dedup() []T
```

Dedup returns a given slice without consecutive duplicated elements

Generic types: T.

## Examples

```go
s := []int{1, 2, 2, 3, 3, 3, 2, 3, 1, 1}
result := genesis.SliceInt{s}.Dedup()
fmt.Println(result)
// Output: [1 2 3 2 3 1]
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

## Source

```go
// Dedup returns a given slice without consecutive duplicated elements
func (s Slice) Dedup() []T {
	if len(s.Data) == 0 {
		return s.Data
	}

	result := make([]T, 0, len(s.Data))
	prev := s.Data[0]
	result = append(result, prev)
	for _, el := range s.Data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}
```

## Tests

```go
func TestSliceDedup(t *testing.T) {
	f := func(given []T, expected []T) {
		actual := Slice{given}.Dedup()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{1})
	f([]T{1, 1}, []T{1})
	f([]T{1, 2}, []T{1, 2})
	f([]T{1, 2, 3}, []T{1, 2, 3})
	f([]T{1, 2, 2, 3}, []T{1, 2, 3})
	f([]T{1, 2, 2, 3, 3, 3, 2, 1, 1}, []T{1, 2, 3, 2, 1})
}
```