# Slices.product

Product returns cortesian product of elements {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4} func (s Slices) Product() chan []T { 	c := make(chan []T, 1) 	go s.product(c, []T{}, 0) 	return c } 

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| SlicesBool | bool |
| SlicesByte | byte |
| SlicesString | string |
| SlicesFloat32 | float32 |
| SlicesFloat64 | float64 |
| SlicesInt | int |
| SlicesInt8 | int8 |
| SlicesInt16 | int16 |
| SlicesInt32 | int32 |
| SlicesInt64 | int64 |
| SlicesUint | uint |
| SlicesUint8 | uint8 |
| SlicesUint16 | uint16 |
| SlicesUint32 | uint32 |
| SlicesUint64 | uint64 |
| SlicesInterface | interface{} |

## Source

```go
// Product returns cortesian product of elements
// {{1, 2}, {3, 4}} -> {1, 3}, {1, 4}, {2, 3}, {2, 4}
func (s Slices) Product() chan []T {
	c := make(chan []T, 1)
	go s.product(c, []T{}, 0)
	return c
}

func (s Slices) product(c chan []T, left []T, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]T, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]T, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}
```

