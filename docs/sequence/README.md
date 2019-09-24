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
| Count | Count is like Range, but infinite |
| Exponential | Exponential generates elements from start with multiplication of value on factor on every step |
| Iterate | Iterate returns an infinite list of repeated applications of f to val |
| Range | Range generates elements from start to end with given step |
| Repeat | Repeat returns channel that produces val infinite times |
| Replicate | Replicate returns channel that produces val n times |
