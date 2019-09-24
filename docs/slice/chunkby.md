# Slice.ChunkBy

ChunkBy splits arr on every element for which f returns a new value.

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
// ChunkBy splits arr on every element for which f returns a new value.
func (s Slice) ChunkBy(f func(el T) G) [][]T {
	chunks := make([][]T, 0)
	chunk := make([]T, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]T, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}
```

## Tests

```go
func TestSliceChunkBy(t *testing.T) {
	f := func(mapper func(t T) G, given []T, expected [][]T) {
		actual := Slice{given}.ChunkBy(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t T) G { return G((t % 2)) }

	f(remainder, []T{1}, [][]T{{1}})
	f(remainder, []T{1, 2, 3}, [][]T{{1}, {2}, {3}})
	f(remainder, []T{1, 3, 2, 4, 5}, [][]T{{1, 3}, {2, 4}, {5}})
}
```