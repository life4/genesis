# Slice.ReduceWhile

ReduceWhile is like Reduce, but stops when f returns error

Generic types: G, T.

## Structs

| Struct | T type |
| ------ | ------ |
| SliceBool | bool |
| SliceByte | byte |
| SliceString | string |
| SliceFloat32 | float32 |
| SliceFloat64 | float64 |
| SliceInt | int |
| SliceInt8 | int8 |
| SliceInt16 | int16 |
| SliceInt32 | int32 |
| SliceInt64 | int64 |
| SliceUint | uint |
| SliceUint8 | uint8 |
| SliceUint16 | uint16 |
| SliceUint32 | uint32 |
| SliceUint64 | uint64 |
| SliceInterface | interface{} |


## Source

```go
// ReduceWhile is like Reduce, but stops when f returns error
func (s Slice) ReduceWhile(acc G, f func(el T, acc G) (G, error)) (G, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}
```

