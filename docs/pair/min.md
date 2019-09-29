# Pair.Min

```go
func (Pair) Min(a T, b T) T
```

Min returns minimal value

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| PairByte | byte |
| PairString | string |
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

## Source

```go
// Min returns minimal value
func (Pair) Min(a T, b T) T {
	if a <= b {
		return a
	}
	return b
}
```

