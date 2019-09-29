# Sequence.Replicate

```go
func (Sequence) Replicate(val T, n int) chan T
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
func (Sequence) Replicate(val T, n int) chan T {
	c := make(chan T, 1)
	go func() {
		for i := 0; i < n; i++ {
			c <- val
		}
	}()
	return c
}
```

