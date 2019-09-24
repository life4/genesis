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
	for val := range c.data {
		result = append(result, val)
	}
	return result
}
```

## Tests

```go
func TestChannelToSlice(t *testing.T) {
	s := Sequence{}
	f := func(start T, stop T, step T, expected []T) {
		seq := s.Range(start, stop, step)
		actual := Channel{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []T{1, 2, 3})
}
```