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

## Errors

| Error | Message |
| -------- | ------ |
| ErrNegativeValue | negative value passed |
| ErrOutOfRange | index is bigger than container size |

## Source

```go
// InsertAt returns the slice with element inserted at given index.
func (s Slice) InsertAt(index int, element T) ([]T, error) {
	result := make([]T, 0, len(s.Data)+1)

	// insert at the end
	if index == len(s.Data) {
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

## Tests

```go
func TestSliceInsertAt(t *testing.T) {
	f := func(given []T, index int, expected []T, expectedErr error) {
		actual, err := Slice{given}.InsertAt(index, 10)
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]T{}, -1, []T{}, ErrNegativeValue)
	f([]T{}, 0, []T{10}, nil)
	f([]T{}, 1, []T{}, ErrOutOfRange)

	f([]T{1, 2, 3}, -1, []T{1, 2, 3}, ErrNegativeValue)
	f([]T{1, 2, 3}, 0, []T{10, 1, 2, 3}, nil)
	f([]T{1, 2, 3}, 1, []T{1, 10, 2, 3}, nil)
	f([]T{1, 2, 3}, 3, []T{1, 2, 3, 10}, nil)
	f([]T{1, 2, 3}, 4, []T{1, 2, 3}, ErrOutOfRange)
}
```