# Channel.ChunkEvery

ChunkEvery returns channel with slices containing count elements each

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| ChannelBool | bool |
| ChannelByte | byte |
| ChannelString | string |
| ChannelFloat32 | float32 |
| ChannelFloat64 | float64 |
| ChannelInt | int |
| ChannelInt8 | int8 |
| ChannelInt16 | int16 |
| ChannelInt32 | int32 |
| ChannelInt64 | int64 |
| ChannelUint | uint |
| ChannelUint8 | uint8 |
| ChannelUint16 | uint16 |
| ChannelUint32 | uint32 |
| ChannelUint64 | uint64 |
| ChannelInterface | interface{} |


## Source

```go
// ChunkEvery returns channel with slices containing count elements each
func (c Channel) ChunkEvery(count int) chan []T {
	chunks := make(chan []T, 1)
	go func() {
		chunk := make([]T, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]T, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}
```

## Tests

```go
func TestChannelChunkEvery(t *testing.T) {
	f := func(size int, given []T, expected [][]T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := Channel{c}.ChunkEvery(size)
		actual := make([][]T, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []T{}, [][]T{})
	f(2, []T{1}, [][]T{{1}})
	f(2, []T{1, 2}, [][]T{{1, 2}})
	f(2, []T{1, 2, 3, 4}, [][]T{{1, 2}, {3, 4}})
	f(2, []T{1, 2, 3, 4, 5}, [][]T{{1, 2}, {3, 4}, {5}})
}
```