# Slice.Scan

Scan is like Reduce, but returns slice of f results

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
// Scan is like Reduce, but returns slice of f results
func (s Slice) Scan(acc G, f func(el T, acc G) G) []G {
	result := make([]G, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}
```