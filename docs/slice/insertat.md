# Slice.InsertAt

```go
func (s Slice) InsertAt(index int, element T) ([]T, error)
```

InsertAt returns the slice with element inserted at given index.

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
// InsertAt returns the slice with element inserted at given index.
func (s Slice) InsertAt(index int, element T) ([]T, error) {
	result := make([]T, 0, len(s.Data)+1)

	// insert at the end
	if index == len(s.Data) || index == -1 {
		result = append(result, s.Data...)
		result = append(result, element)
		return result, nil
	}

	if index > len(s.Data) {
		return s.Data, ErrOutOfRange
	}
	if index < 0 {
		return s.Data, ErrNegativeValue
	}

	for i, el := range s.Data {
		if i == index {
			result = append(result, element)
		}
		result = append(result, el)
	}
	return result, nil
}
```

