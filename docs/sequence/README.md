# Sequence

Sequence is a set of operations to generate sequences

```go
// Sequence is a set of operations to generate sequences
type Sequence struct {
	ctx context.Context
}
```

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

## Functions

| Function | Description |
| -------- | ----------- |
| [Count](./count.md) | Count is like Range, but infinite |
| [Exponential](./exponential.md) | Exponential generates elements from start with multiplication of value on factor on every step |
| [Iterate](./iterate.md) | Iterate returns an infinite list of repeated applications of f to val |
| [Range](./range.md) | Range generates elements from start to end with given step |
| [Repeat](./repeat.md) | Repeat returns channel that produces val infinite times |
| [Replicate](./replicate.md) | Replicate returns channel that produces val n times |
