# Channel.Max

```go
func (c Channel) Max() T
```

Max returns the maximal element from channel

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
// Max returns the maximal element from channel
func (c Channel) Max() T {
	max := <-c.Data
	for el := range c.Data {
		if el > max {
			max = el
		}
	}
	return max
}
```

