# Slice.TakeEvery

```go
func (s Slice) TakeEvery(nth int) ([]T, error)
```

TakeEvery returns slice of every nth elements

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
| ErrNonPositiveValue | value must be positive |

## Source

```go
// TakeEvery returns slice of every nth elements
func (s Slice) TakeEvery(nth int) ([]T, error) {
	if nth <= 0 {
		return s.Data, ErrNonPositiveValue
	}
	result := make([]T, 0, len(s.Data))
	for i, el := range s.Data {
		if (i+1)%nth == 0 {
			result = append(result, el)
		}
	}
	return result, nil
}
```

