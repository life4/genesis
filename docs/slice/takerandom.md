# Slice.TakeRandom

```go
func (s Slice) TakeRandom(count int) ([]T, error)
```

TakeRandom returns slice of count random elements from the slice

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
// TakeRandom returns slice of count random elements from the slice
func (s Slice) TakeRandom(count int) ([]T, error) {
	if count > len(s.Data) {
		return nil, ErrOutOfRange
	}
	if count <= 0 {
		return nil, ErrNonPositiveValue
	}

	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
	}
	rand.Shuffle(len(s.Data), swap)
	return s.Data[:count], nil
}
```

