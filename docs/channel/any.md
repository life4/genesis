# Channel.Any

```go
func (c Channel) Any(f func(el T) bool) bool
```

Any returns true if f returns true for any element in channel

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
// Any returns true if f returns true for any element in channel
func (c Channel) Any(f func(el T) bool) bool {
	for el := range c.Data {
		if f(el) {
			return true
		}
	}
	return false
}
```

## Tests

```go
func TestChannelAny(t *testing.T) {
	f := func(given []T, expected bool) {
		even := func(t T) bool { return t%2 == 0 }
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, false)
	f([]T{1}, false)
	f([]T{2}, true)
	f([]T{1, 2}, true)
	f([]T{1, 2, 3}, true)
	f([]T{1, 3, 5}, false)
	f([]T{1, 3, 5, 7, 9, 11}, false)
	f([]T{1, 3, 5, 7, 10, 11}, true)
}
```