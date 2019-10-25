# Slice.TakeRandom

```go
func (s Slice) TakeRandom(count int, seed int64) ([]T, error)
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

## Errors

| Error | Message |
| -------- | ------ |
| ErrNonPositiveValue | value must be positive |
| ErrOutOfRange | index is bigger than container size |

## Source

```go
// TakeRandom returns slice of count random elements from the slice
func (s Slice) TakeRandom(count int, seed int64) ([]T, error) {
	if count > len(s.Data) {
		return nil, ErrOutOfRange
	}
	if count <= 0 {
		return nil, ErrNonPositiveValue
	}

	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)
	swap := func(i, j int) {
		s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
	}
	rand.Shuffle(len(s.Data), swap)
	return s.Data[:count], nil
}
```

## Tests

```go
func TestSliceTakeRandom(t *testing.T) {
	f := func(given []T, count int, seed int64, expected []T) {
		actual, _ := Slice{given}.TakeRandom(count, seed)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{1}, 1, 0, []T{1})
	f([]T{1, 2, 3, 4, 5}, 3, 1, []T{3, 1, 2})
	f([]T{1, 2, 3, 4, 5}, 5, 1, []T{3, 1, 2, 5, 4})
}
```