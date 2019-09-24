# Channel.Map

Map applies f to all elements from channel and returns channel with results

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


## Source

```go
// Map applies f to all elements from channel and returns channel with results
func (c Channel) Map(f func(el T) G) chan G {
	result := make(chan G, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}
```

