# Channel.Min

```go
func (c Channel) Min() (T, error)
```

Min returns the minimal element from channel

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
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

## Errors

| Error | Message |
| -------- | ------ |
| ErrEmpty | container is empty |

## Source

```go
// Min returns the minimal element from channel
func (c Channel) Min() (T, error) {
	min, ok := <-c.Data
	if !ok {
		return min, ErrEmpty
	}
	for el := range c.Data {
		if el < min {
			min = el
		}
	}
	return min, nil
}
```

## Tests

```go
func TestChannelMin(t *testing.T) {
	f := func(given []T, expected T, expectedErr error) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := Channel{c}.Min()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]T{}, 0, ErrEmpty)
	f([]T{4, 1, 2}, 1, nil)
	f([]T{1, 2, 4}, 1, nil)
	f([]T{4, 2, 1}, 1, nil)
}
```