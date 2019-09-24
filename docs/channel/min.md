# Channel.Min

Min returns the minimal element from channel

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
// Min returns the minimal element from channel
func (c Channel) Min() T {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}
```

