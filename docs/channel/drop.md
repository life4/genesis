# Channel.Drop

```go
func (c Channel) Drop(n int) chan T
```

Drop drops first n elements from channel c and returns a new channel with the rest. It returns channel do be unblocking. If you want array instead, wrap result into TakeAll.

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
// Drop drops first n elements from channel c and returns a new channel with the rest.
// It returns channel do be unblocking. If you want array instead, wrap result into TakeAll.
func (c Channel) Drop(n int) chan T {
	result := make(chan T, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}
```

## Tests

```go
func TestChannelDrop(t *testing.T) {
	f := func(count int, given []T, expected []T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := Channel{c}.Drop(count)
		actual := make([]T, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []T{}, []T{})
	f(1, []T{2}, []T{})
	f(1, []T{2, 3}, []T{3})
	f(1, []T{1, 2, 3}, []T{2, 3})
	f(0, []T{1, 2, 3}, []T{1, 2, 3})
	f(3, []T{1, 2, 3, 4, 5, 6}, []T{4, 5, 6})
	f(1, []T{1, 2, 3, 4, 5, 6}, []T{2, 3, 4, 5, 6})
}
```