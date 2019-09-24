# Slice.Map

Map applies F to all elements in slice of T and returns slice of results

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
// Map applies F to all elements in slice of T and returns slice of results
func (s Slice) Map(f func(el T) G) []G {
	result := make([]G, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}
```