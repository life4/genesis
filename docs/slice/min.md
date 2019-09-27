# Slice.Min

```go
func (s Slice) Min() (T, error)
```

Min returns the minimal element from arr

Generic types: T.

## Examples

```go
s := []int{42, 7, 13}
min, _ := genesis.SliceInt{s}.Min()
fmt.Println(min)
// Output: 7
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

## Errors

| Error | Message |
| -------- | ------ |
| ErrEmpty | container is empty |

## Source

```go
// Min returns the minimal element from arr
func (s Slice) Min() (T, error) {
	if len(s.Data) == 0 {
		var tmp T
		return tmp, ErrEmpty
	}

	min := s.Data[0]
	for _, el := range s.Data[1:] {
		if el < min {
			min = el
		}
	}
	return min, nil
}
```

