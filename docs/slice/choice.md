# Slice.Choice

```go
func (s Slice) Choice(seed int64) (T, error)
```

Choice chooses a random element from the slice. If seed is zero, UNIX timestamp will be used.

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
| ErrEmpty | container is empty |

## Source

```go
// Choice chooses a random element from the slice.
// If seed is zero, UNIX timestamp will be used.
func (s Slice) Choice(seed int64) (T, error) {
	if len(s.Data) == 0 {
		var tmp T
		return tmp, ErrEmpty
	}

	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)
	i := rand.Intn(len(s.Data))
	return s.Data[i], nil
}
```

## Tests

```go
func TestSliceChoice(t *testing.T) {
	f := func(given []T, seed int64, expected T) {
		actual, _ := Slice{given}.Choice(seed)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{1}, 0, 1)
	f([]T{1, 2, 3}, 1, 3)
	f([]T{1, 2, 3}, 2, 2)
}
```