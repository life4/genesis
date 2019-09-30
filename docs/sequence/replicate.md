# Sequence.Replicate

```go
func (s Sequence) Replicate(val T, n int) chan T
```

Replicate returns channel that produces val n times

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| SequenceBool | bool |
| SequenceByte | byte |
| SequenceString | string |
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
// Replicate returns channel that produces val n times
func (s Sequence) Replicate(val T, n int) chan T {
	c := make(chan T, 1)
	go func() {
		for i := 0; i < n; i++ {
			c <- val
		}
		close(c)
	}()
	return c
}
```

## Tests

```go
func TestSequenceReplicate(t *testing.T) {
	f := func(count int, given T, expected []T) {
		ctx, cancel := context.WithCancel(context.Background())
		s := Sequence{ctx: ctx}
		seq := s.Replicate(given, count)
		actual := Channel{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []T{})
	f(1, 1, []T{1})
	f(5, 1, []T{1, 1, 1, 1, 1})
}
```