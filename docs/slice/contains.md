# Slice.Contains

```go
func (s Slice) Contains(el T) bool
```

Contains returns true if el in arr.

Generic types: T.

## Examples

```go
s := []int{2, 4, 6, 8}
result := genesis.SliceInt{s}.Contains(4)
fmt.Println(result)

result = genesis.SliceInt{s}.Contains(3)
fmt.Println(result)

// Output:
// true
// false
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
// Contains returns true if el in arr.
func (s Slice) Contains(el T) bool {
	for _, val := range s.Data {
		if val == el {
			return true
		}
	}
	return false
}
```

## Tests

```go
func TestSliceContains(t *testing.T) {
	f := func(el T, given []T, expected bool) {
		actual := Slice{given}.Contains(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []T{}, false)
	f(1, []T{1}, true)
	f(1, []T{2}, false)
	f(1, []T{2, 3, 4, 5}, false)
	f(1, []T{2, 3, 1, 4, 5}, true)
	f(1, []T{2, 3, 1, 1, 4, 5}, true)
}
```