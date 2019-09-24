# Channel.Tee

Tee returns 2 channels with elements from the input channel

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
// Tee returns 2 channels with elements from the input channel
func (c Channel) Tee(count int) []chan T {
	channels := make([]chan T, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan T, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan T) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}
```

