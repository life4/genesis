# Slice.Count

```go
func (s Slice) Count(el T) int
```

Count return count of el occurences in arr.

Generic types: T.

## Examples

```go
s := []int{1, 0, 1, 0, 0, 1, 1, 0, 1, 0}
result := genesis.SliceInt{s}.Count(1)
fmt.Println(result)
// Output: 5
```

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
// Count return count of el occurences in arr.
func (s Slice) Count(el T) int {
	count := 0
	for _, val := range s.Data {
		if val == el {
			count++
		}
	}
	return count
}
```

## Tests

```go
func TestSliceCount(t *testing.T) {
	f := func(el T, given []T, expected int) {
		actual := Slice{given}.Count(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []T{}, 0)
	f(1, []T{1}, 1)
	f(1, []T{2}, 0)
	f(1, []T{2, 3, 4, 5}, 0)
	f(1, []T{2, 3, 1, 4, 5}, 1)
	f(1, []T{2, 3, 1, 1, 4, 5}, 2)
	f(1, []T{1, 1, 1, 1, 1}, 5)
}
```