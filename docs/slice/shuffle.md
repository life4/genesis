# Slice.Shuffle

```go
func (s Slice) Shuffle(seed int64) []T
```

Shuffle in random order arr elements

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
// Shuffle in random order arr elements
func (s Slice) Shuffle(seed int64) []T {
	if len(s.Data) <= 1 {
		return s.Data
	}
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)
	swap := func(i, j int) {
		s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
	}
	rand.Shuffle(len(s.Data), swap)
	return s.Data
}
```

## Tests

```go
func TestSliceShuffle(t *testing.T) {
	f := func(given []T, seed int64, expected []T) {
		actual := Slice{given}.Shuffle(seed)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 0, []T{})
	f([]T{1}, 0, []T{1})
	f([]T{1, 2, 3, 4, 5, 6}, 2, []T{3, 5, 4, 1, 6, 2})
	f([]T{1, 2, 2, 3, 3}, 2, []T{3, 2, 3, 2, 1})
}
```