# Channel.Scan

Scan is like Reduce, but returns slice of f results

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
// Scan is like Reduce, but returns slice of f results
func (c Channel) Scan(acc G, f func(el T, acc G) G) chan G {
	result := make(chan G, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}
```

