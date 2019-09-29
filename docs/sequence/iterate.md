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
		for {
			c <- val
			val = f(val)
		}
	}()
	return c
}
```

