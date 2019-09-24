# Slice.ChunkEvery

```go
func (s Slice) ChunkEvery(count int) [][]T
```

ChunkEvery returns slice of slices containing count elements each

Generic types: T.

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
// ChunkEvery returns slice of slices containing count elements each
func (s Slice) ChunkEvery(count int) [][]T {
	chunks := make([][]T, 0)
	chunk := make([]T, 0, count)
	for i, el := range s.Data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]T, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}
```

## Tests

```go
func TestSliceChunkEvery(t *testing.T) {
	f := func(count int, given []T, expected [][]T) {
		actual := Slice{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []T{1, 2, 3, 4}, [][]T{[]T{1, 2}, []T{3, 4}})
	f(2, []T{1, 2, 3, 4, 5}, [][]T{[]T{1, 2}, []T{3, 4}, []T{5}})
}
```