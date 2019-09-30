# Slice.DropEvery

```go
func (s Slice) DropEvery(nth int) ([]T, error)
```

DropEvery returns a slice of every nth element in the enumerable dropped, starting with the first element.

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
// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func (s Slice) DropEvery(nth int) ([]T, error) {
	if nth <= 0 {
		return s.Data, ErrNonPositiveValue
	}
	result := make([]T, 0, len(s.Data)/nth)
	for i, el := range s.Data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result, nil
}
```

## Tests

```go
func TestSliceDropEvery(t *testing.T) {
	f := func(given []T, nth int, expected []T) {
		actual, _ := Slice{given}.DropEvery(nth)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 1, []T{})
	f([]T{1, 2, 3}, 1, []T{})

	f([]T{1, 2, 3, 4}, 2, []T{1, 3})
	f([]T{1, 2, 3, 4, 5}, 2, []T{1, 3, 5})

	f([]T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, []T{1, 3, 5, 7, 9})
	f([]T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, []T{1, 2, 4, 5, 7, 8, 10})
}
```