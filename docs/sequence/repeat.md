# Sequence.Repeat

```go
func (s Sequence) Repeat(val T) chan T
```

Repeat returns channel that produces val infinite times

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| SequenceBool | bool |
| SequenceByte | byte |
| SequenceString | string |
| SequenceRune | rune |
| SequenceError | error |
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
| SequenceInterface | interface{} |

## Source

```go
// Repeat returns channel that produces val infinite times
func (s Sequence) Repeat(val T) chan T {
	c := make(chan T, 1)
	go func() {
		defer close(c)
		for {
			select {
			case <-s.ctx.Done():
				return
			case c <- val:
				continue
			}
		}
	}()
	return c
}
```

## Tests

```go
func TestSequenceRepeat(t *testing.T) {
	f := func(count int, given T, expected []T) {
		ctx, cancel := context.WithCancel(context.Background())
		s := Sequence{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []T{1, 1})
}
```