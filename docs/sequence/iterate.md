# Sequence.Iterate

```go
func (s Sequence) Iterate(val T, f func(val T) T) chan T
```

Iterate returns an infinite list of repeated applications of f to val

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
// Iterate returns an infinite list of repeated applications of f to val
func (s Sequence) Iterate(val T, f func(val T) T) chan T {
	c := make(chan T, 1)
	go func() {
		defer close(c)
		for {
			select {
			case <-s.ctx.Done():
				return
			case c <- val:
				val = f(val)
			}
		}
	}()
	return c
}
```

## Tests

```go
func TestSequenceIterate(t *testing.T) {
	f := func(start T, count int, expected []T) {
		ctx, cancel := context.WithCancel(context.Background())
		s := Sequence{ctx: ctx}
		double := func(val T) T { return val * 2 }
		seq := s.Iterate(start, double)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, []T{1, 2, 4, 8})
}
```