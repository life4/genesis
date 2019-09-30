# Channel.Take

```go
func (c Channel) Take(count int) chan T
```

Take takes first count elements from the channel.

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
// Take takes first count elements from the channel.
func (c Channel) Take(count int) chan T {
	result := make(chan T, 1)
	go func() {
		defer close(result)
		if count <= 0 {
			return
		}
		i := 0
		for el := range c.Data {
			result <- el
			i++
			if i == count {
				return
			}
		}
	}()
	return result
}
```

## Tests

```go
func TestChannelTake(t *testing.T) {
	f := func(count int, given T, expected []T) {
		ctx, cancel := context.WithCancel(context.Background())
		s := Sequence{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []T{})
	f(1, 1, []T{1})
	f(2, 1, []T{1, 1})
}
```