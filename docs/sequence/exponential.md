# Sequence.Exponential

```go
func (s Sequence) Exponential(start T, factor T) chan T
```

Exponential generates elements from start with multiplication of value on factor on every step

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
// Exponential generates elements from start with
// multiplication of value on factor on every step
func (s Sequence) Exponential(start T, factor T) chan T {
	c := make(chan T, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}
```

## Tests

```go
func TestSequenceExponential(t *testing.T) {
	s := Sequence{}
	f := func(start T, factor T, count int, expected []T) {
		seq := s.Exponential(start, factor)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 1, 4, []T{1, 1, 1, 1})
	f(1, 2, 4, []T{1, 2, 4, 8})
}
```