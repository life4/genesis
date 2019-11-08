# Slices.Zip

```go
func (s Slices) Zip() chan []T
```

Zip returns chan of arrays of elements from given arrs on the same position.

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| SlicesBool | bool |
| SlicesByte | byte |
| SlicesString | string |
| SlicesRune | rune |
| SlicesError | error |
| SlicesFloat32 | float32 |
| SlicesFloat64 | float64 |
| SlicesInt | int |
| SlicesInt8 | int8 |
| SlicesInt16 | int16 |
| SlicesInt32 | int32 |
| SlicesInt64 | int64 |
| SlicesUint | uint |
| SlicesUint8 | uint8 |
| SlicesUint16 | uint16 |
| SlicesUint32 | uint32 |
| SlicesUint64 | uint64 |
| SlicesInterface | interface{} |

## Source

```go
// Zip returns chan of arrays of elements from given arrs on the same position.
func (s Slices) Zip() chan []T {
	if len(s.Data) == 0 {
		result := make(chan []T)
		close(result)
		return result
	}

	size := len(s.Data[0])
	for _, arr := range s.Data[1:] {
		if len(arr) < size {
			size = len(arr)
		}
	}

	result := make(chan []T, 1)
	go func() {
		for i := 0; i < size; i++ {
			chunk := make([]T, 0, len(s.Data))
			for _, arr := range s.Data {
				chunk = append(chunk, arr[i])
			}
			result <- chunk
		}
		close(result)
	}()
	return result
}
```

## Tests

```go
func TestSlicesZip(t *testing.T) {
	f := func(given [][]T, expected [][]T) {
		actual := make([][]T, 0)
		i := 0
		s := Slices{given}
		for el := range s.Zip() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]T{}, [][]T{})
	f([][]T{{1}, {2}}, [][]T{{1, 2}})
	f([][]T{{1, 2}, {3, 4}}, [][]T{{1, 3}, {2, 4}})
	f([][]T{{1, 2, 3}, {4, 5}}, [][]T{{1, 4}, {2, 5}})
}
```