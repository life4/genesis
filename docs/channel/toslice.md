# Channel.ToSlice

```go
func (c Channel) ToSlice() []T
```

ToSlice returns slice with all elements from channel.

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
// ToSlice returns slice with all elements from channel.
func (c Channel) ToSlice() []T {
	result := make([]T, 0)
	for val := range c.Data {
		result = append(result, val)
	}
	return result
}
```

## Tests

```go
func TestChannelToSlice(t *testing.T) {
	f := func(given []T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}
	f([]T{})
	f([]T{1})
	f([]T{1, 2, 3, 1, 2})
}
```