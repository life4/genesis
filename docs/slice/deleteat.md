# Slice.DeleteAt

```go
func (s Slice) DeleteAt(index int) ([]T, error)
```

DeleteAt returns the slice without elements on given positions

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

## Errors

| Error | Message |
| -------- | ------ |
| ErrOutOfRange | index is bigger than container size |

## Source

```go
// DeleteAt returns the slice without elements on given positions
func (s Slice) DeleteAt(index int) ([]T, error) {
	if index >= len(s.Data) {
		return s.Data, ErrOutOfRange
	}

	result := make([]T, 0, len(s.Data)-1)
	for i, el := range s.Data {
		if i != index {
			result = append(result, el)
		}
	}
	return result, nil
}
```

