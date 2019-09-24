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

## Functions

| Function | G type |
| -------- | ------ |
| MapBool | bool |
| MapByte | byte |
| MapString | string |
| MapFloat32 | float32 |
| MapFloat64 | float64 |
| MapInt | int |
| MapInt8 | int8 |
| MapInt16 | int16 |
| MapInt32 | int32 |
| MapInt64 | int64 |
| MapUint | uint |
| MapUint8 | uint8 |
| MapUint16 | uint16 |
| MapUint32 | uint32 |
| MapUint64 | uint64 |
| MapInterface | interface{} |

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

