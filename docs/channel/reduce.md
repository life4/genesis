# Channel.Reduce

Reduce applies f to acc and every element from channel and returns acc

Generic types: G, T.

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

## Functions

| Function | G type |
| -------- | ------ |
| ReduceBool | bool |
| ReduceByte | byte |
| ReduceString | string |
| ReduceFloat32 | float32 |
| ReduceFloat64 | float64 |
| ReduceInt | int |
| ReduceInt8 | int8 |
| ReduceInt16 | int16 |
| ReduceInt32 | int32 |
| ReduceInt64 | int64 |
| ReduceUint | uint |
| ReduceUint8 | uint8 |
| ReduceUint16 | uint16 |
| ReduceUint32 | uint32 |
| ReduceUint64 | uint64 |
| ReduceInterface | interface{} |

## Source

```go
// Reduce applies f to acc and every element from channel and returns acc
func (c Channel) Reduce(acc G, f func(el T, acc G) G) G {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}
```

