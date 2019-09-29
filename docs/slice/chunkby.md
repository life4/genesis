# Slice.ChunkBy

```go
func (s Slice) ChunkBy(f func(el T) G) [][]T
```

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

## Functions

| Function | G type |
| -------- | ------ |
| ChunkByBool | bool |
| ChunkByByte | byte |
| ChunkByString | string |
| ChunkByFloat32 | float32 |
| ChunkByFloat64 | float64 |
| ChunkByInt | int |
| ChunkByInt8 | int8 |
| ChunkByInt16 | int16 |
| ChunkByInt32 | int32 |
| ChunkByInt64 | int64 |
| ChunkByUint | uint |
| ChunkByUint8 | uint8 |
| ChunkByUint16 | uint16 |
| ChunkByUint32 | uint32 |
| ChunkByUint64 | uint64 |
| ChunkByInterface | interface{} |

## Source

```go
// ChunkBy splits arr on every element for which f returns a new value.
func (s Slice) ChunkBy(f func(el T) G) [][]T {
	chunks := make([][]T, 0)
	if len(s.Data) == 0 {
		return chunks
	}

	chunk := make([]T, 0)
	prev := f(s.Data[0])
	chunk = append(chunk, s.Data[0])

	for _, el := range s.Data[1:] {
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
	f := func(given []T, expected [][]T) {
		reminder := func(t T) G { return G((t % 2)) }
		actual := Slice{given}.ChunkBy(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, [][]T{})
	f([]T{1}, [][]T{{1}})
	f([]T{1, 2, 3}, [][]T{{1}, {2}, {3}})
	f([]T{1, 3, 2, 4, 5}, [][]T{{1, 3}, {2, 4}, {5}})
}
```