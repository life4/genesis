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

## Functions

| Function | G type |
| -------- | ------ |
| ScanBool | bool |
| ScanByte | byte |
| ScanString | string |
| ScanFloat32 | float32 |
| ScanFloat64 | float64 |
| ScanInt | int |
| ScanInt8 | int8 |
| ScanInt16 | int16 |
| ScanInt32 | int32 |
| ScanInt64 | int64 |
| ScanUint | uint |
| ScanUint8 | uint8 |
| ScanUint16 | uint16 |
| ScanUint32 | uint32 |
| ScanUint64 | uint64 |
| ScanInterface | interface{} |

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

