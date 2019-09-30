# Sequence.Range

```go
func (s Sequence) Range(start T, end T, step T) chan T
```

Range generates elements from start to end with given step

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
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
// Range generates elements from start to end with given step
func (s Sequence) Range(start T, end T, step T) chan T {
	c := make(chan T, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}
```

## Tests

```go
func TestSequenceRange(t *testing.T) {
	f := func(start T, stop T, step T, expected []T) {
		ctx, cancel := context.WithCancel(context.Background())
		s := Sequence{ctx: ctx}
		seq := s.Range(start, stop, step)
		actual := Channel{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []T{1, 2, 3})
	f(3, 0, -1, []T{3, 2, 1})
	f(1, 1, 1, []T{})
	f(1, 2, 1, []T{1})
}
```