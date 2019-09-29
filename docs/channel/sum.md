# Channel.Sum

```go
func (c Channel) Sum() T
```

Sum returns sum of all elements from channel

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
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

## Source

```go
// Sum returns sum of all elements from channel
func (c Channel) Sum() T {
	var sum T
	for el := range c.Data {
		sum += el
	}
	return sum
}
```

## Tests

```go
func TestChannelSum(t *testing.T) {
	f := func(given []T, expected T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.Sum()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 0)
	f([]T{1}, 1)
	f([]T{1, 2}, 3)
	f([]T{1, 2, 3, 4, 5}, 15)
}
```