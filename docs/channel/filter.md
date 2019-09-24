# Channel.Filter

Filter returns a new channel with elements from input channel for which f returns true

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
// Filter returns a new channel with elements from input channel
// for which f returns true
func (c Channel) Filter(f func(el T) bool) chan T {
	result := make(chan T, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}
```