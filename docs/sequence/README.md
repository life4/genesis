# Sequence

Sequence is a set of operations to generate sequences

```go
// Sequence is a set of operations to generate sequences
type Sequence struct {
	data chan T
}
```

| Function | Description |
| -------- | ----------- |
| [Count](./count.md) | Count is like Range, but infinite |
| [Exponential](./exponential.md) | Exponential generates elements from start with multiplication of value on factor on every step |
| [Iterate](./iterate.md) | Iterate returns an infinite list of repeated applications of f to val |
| [Range](./range.md) | Range generates elements from start to end with given step |
| [Repeat](./repeat.md) | Repeat returns channel that produces val infinite times |
| [Replicate](./replicate.md) | Replicate returns channel that produces val n times |
