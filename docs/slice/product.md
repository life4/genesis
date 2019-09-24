# Slice.product

product is a core implementation for Product

Generic types: T.

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
// product is a core implementation for Product
func (s Slice) product(c chan []T, repeat int, left []T, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]T, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]T, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}
```

