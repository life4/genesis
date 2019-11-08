# Channel.Each

```go
func (c Channel) Each(f func(el T))
```

Each calls f for every element in the channel

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| ChannelBool | bool |
| ChannelByte | byte |
| ChannelString | string |
| ChannelRune | rune |
| ChannelError | error |
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
// Each calls f for every element in the channel
func (c Channel) Each(f func(el T)) {
	for el := range c.Data {
		f(el)
	}
}
```

## Tests

```go
func TestChannelEach(t *testing.T) {
	f := func(given []T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := make(chan T, len(given))
		mapper := func(t T) { result <- t }
		Channel{c}.Each(mapper)
		close(result)
		actual := Channel{result}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}

	f([]T{})
	f([]T{1})
	f([]T{1, 2, 3})
	f([]T{1, 2, 3, 4, 5, 6, 7})
}
```