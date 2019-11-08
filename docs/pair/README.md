# Pair

Pair is a set of functions for 2 values that you can pass into reduce-like funcs

```go
// Pair is a set of functions for 2 values that you can pass into reduce-like funcs
type Pair struct {
	// empty
}
```

## Structs

| Struct | T type |
| ------ | ------ |
| PairBool | bool |
| PairByte | byte |
| PairString | string |
| PairRune | rune |
| PairError | error |
| PairFloat32 | float32 |
| PairFloat64 | float64 |
| PairInt | int |
| PairInt8 | int8 |
| PairInt16 | int16 |
| PairInt32 | int32 |
| PairInt64 | int64 |
| PairUint | uint |
| PairUint8 | uint8 |
| PairUint16 | uint16 |
| PairUint32 | uint32 |
| PairUint64 | uint64 |
| PairInterface | interface{} |

## Functions

| Function | Description |
| -------- | ----------- |
| [Min](./min.md) | Min returns minimal value |
| [Max](./max.md) | Max returns maximal value |
