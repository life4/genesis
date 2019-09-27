# Slice.Choice

```go
func (s Slice) Choice() (T, error)
```

Choice chooses a random element from the slice

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
// Choice chooses a random element from the slice
func (s Slice) Choice() (T, error) {
	if len(s.Data) == 0 {
		var tmp T
		return tmp, ErrEmptySlice
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(s.Data))
	return s.Data[i], nil
}
```

