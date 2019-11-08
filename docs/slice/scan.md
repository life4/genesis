# Slice.Scan

```go
func (s Slice) Scan(acc G, f func(el T, acc G) G) []G
```

Scan is like Reduce, but returns slice of f results

Generic types: G, T.

## Structs

| Struct | T type |
| ------ | ------ |
| SliceBool | bool |
| SliceByte | byte |
| SliceString | string |
| SliceRune | rune |
| SliceError | error |
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
| ScanRune | rune |
| ScanError | error |
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
	result := make([]G, 0, len(s.Data))
	for _, el := range s.Data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}
```

## Tests

```go
func TestSliceScan(t *testing.T) {
	f := func(given []T, expected []G) {
		sum := func(el T, acc G) G { return G(el) + acc }
		actual := Slice{given}.Scan(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []G{})
	f([]T{1}, []G{1})
	f([]T{1, 2}, []G{1, 3})
	f([]T{1, 2, 3}, []G{1, 3, 6})
	f([]T{1, 2, 3, 4}, []G{1, 3, 6, 10})
}
```