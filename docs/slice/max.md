# Slice.Max

```go
func (s Slice) Max() T
```

Max returns the maximal element from arr

Generic types: T.

## Examples

```go
s := []int{7, 42, 13}
max := genesis.SliceInt{s}.Max()
fmt.Println(max)
// Output: 42
```

## Structs

| Struct | T type |
| ------ | ------ |
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

## Source

```go
// Max returns the maximal element from arr
func (s Slice) Max() T {
	max := s.Data[0]
	for _, el := range s.Data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}
```

