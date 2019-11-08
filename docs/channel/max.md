# Channel.Max

```go
func (c Channel) Max() (T, error)
```

Max returns the maximal element from channel

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| ChannelByte | byte |
| ChannelString | string |
| ChannelRune | rune |
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
// Max returns the maximal element from channel
func (c Channel) Max() (T, error) {
	max, ok := <-c.Data
	if !ok {
		return max, ErrEmpty
	}
	for el := range c.Data {
		if el > max {
			max = el
		}
	}
	return max, nil
}
```

## Tests

```go
func TestChannelMax(t *testing.T) {
	f := func(given []T, expected T, expectedErr error) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := Channel{c}.Max()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]T{}, 0, ErrEmpty)
	f([]T{1, 4, 2}, 4, nil)
	f([]T{1, 2, 4}, 4, nil)
	f([]T{4, 2, 1}, 4, nil)
}
```