# Sequence.Count

```go
func (s Sequence) Count(start T, step T) chan T
```

Count is like Range, but infinite

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| SequenceRune | rune |
| SequenceFloat32 | float32 |
| SequenceFloat64 | float64 |
| SequenceInt | int |
| SequenceInt8 | int8 |
| SequenceInt16 | int16 |
| SequenceInt32 | int32 |
| SequenceInt64 | int64 |
| SequenceUint | uint |
| SequenceUint8 | uint8 |
| SequenceUint16 | uint16 |
| SequenceUint32 | uint32 |
| SequenceUint64 | uint64 |

## Source

```go
// Count is like Range, but infinite
func (s Sequence) Count(start T, step T) chan T {
	c := make(chan T, 1)
	go func() {
		defer close(c)
		for {
			select {
			case <-s.ctx.Done():
				return
			case c <- start:
				start += step
			}
		}
	}()
	return c
}
```

## Tests

```go
func TestSequenceCount(t *testing.T) {
	f := func(start T, step T, count int, expected []T) {
		ctx, cancel := context.WithCancel(context.Background())
		s := Sequence{ctx: ctx}
		seq := s.Count(start, step)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []T{1, 3, 5, 7})
}
```