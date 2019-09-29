# Channel.Tee

```go
func (c Channel) Tee(count int) []chan T
```

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
		for el := range c.Data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan T) {
				defer wg.Done()
				ch <- el
			}
			wg.Add(count)
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

## Tests

```go
func TestChannelTee(t *testing.T) {
	f := func(count int, given []T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		channels := Channel{c}.Tee(count)
		for _, ch := range channels {
			go func(ch chan T) {
				actual := Channel{ch}.ToSlice()
				assert.Equal(t, given, actual, "they should be equal")
			}(ch)
		}
	}
	f(1, []T{})
	f(1, []T{1})
	f(1, []T{1, 2})
	f(1, []T{1, 2, 3})
	f(1, []T{1, 2, 3, 1, 2})

	f(2, []T{})
	f(2, []T{1})
	f(2, []T{1, 2})
	f(2, []T{1, 2, 3})
	f(2, []T{1, 2, 3, 1, 2})

	f(10, []T{1, 2, 3, 1, 2})
}
```