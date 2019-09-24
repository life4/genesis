# Channel.Filter

```go
func (c Channel) Filter(f func(el T) bool) chan T
```

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
		for el := range c.Data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}
```

## Tests

```go
func TestChannelFilter(t *testing.T) {
	f := func(given []T, expected []T) {
		even := func(t T) bool { return t%2 == 0 }
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := Channel{c}.Filter(even)
		actual := Channel{result}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{})
	f([]T{2}, []T{2})
	f([]T{1, 2, 3, 4}, []T{2, 4})
}
```