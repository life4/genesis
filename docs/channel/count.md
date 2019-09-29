# Channel.Count

```go
func (c Channel) Count(el T) int
```

Count return count of el occurences in channel.

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
// Count return count of el occurences in channel.
func (c Channel) Count(el T) int {
	count := 0
	for val := range c.Data {
		if val == el {
			count++
		}
	}
	return count
}
```

## Tests

```go
func TestChannelCount(t *testing.T) {
	f := func(element T, given []T, expected int) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.Count(element)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []T{}, 0)
	f(1, []T{1}, 1)
	f(1, []T{2}, 0)
	f(1, []T{1, 2, 3, 1, 4}, 2)
}
```