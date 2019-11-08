# Slice.Max

```go
func (s Slice) Max() (T, error)
```

Max returns the maximal element from arr

Generic types: T.

## Examples

```go
s := []int{7, 42, 13}
max, _ := genesis.SliceInt{s}.Max()
fmt.Println(max)
// Output: 42
```

## Structs

| Struct | T type |
| ------ | ------ |
| SliceByte | byte |
| SliceString | string |
| SliceRune | rune |
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

## Errors

| Error | Message |
| -------- | ------ |
| ErrEmpty | container is empty |

## Source

```go
// Max returns the maximal element from arr
func (s Slice) Max() (T, error) {
	if len(s.Data) == 0 {
		var tmp T
		return tmp, ErrEmpty
	}

	max := s.Data[0]
	for _, el := range s.Data[1:] {
		if el > max {
			max = el
		}
	}
	return max, nil
}
```

## Tests

```go
func TestSliceMax(t *testing.T) {
	f := func(given []T, expectedEl T, expectedErr error) {
		el, err := Slice{given}.Max()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]T{}, 0, ErrEmpty)
	f([]T{1}, 1, nil)
	f([]T{1, 2, 3}, 3, nil)
	f([]T{1, 3, 2}, 3, nil)
	f([]T{3, 2, 1}, 3, nil)
}
```