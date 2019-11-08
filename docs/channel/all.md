# Channel.All

```go
func (c Channel) All(f func(el T) bool) bool
```

All returns true if f returns true for all elements in channel

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| ChannelBool | bool |
| ChannelByte | byte |
| ChannelString | string |
| ChannelRune | rune |
| ChannelError | error |
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
// All returns true if f returns true for all elements in channel
func (c Channel) All(f func(el T) bool) bool {
	for el := range c.Data {
		if !f(el) {
			return false
		}
	}
	return true
}
```

## Tests

```go
func TestChannelAll(t *testing.T) {
	f := func(given []T, expected bool) {
		even := func(t T) bool { return t%2 == 0 }
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, true)
	f([]T{1}, false)
	f([]T{2}, true)
	f([]T{1, 2}, false)
	f([]T{2, 4}, true)
	f([]T{2, 4, 6, 8, 10, 12}, true)
	f([]T{2, 4, 6, 8, 11, 12}, false)
}
```