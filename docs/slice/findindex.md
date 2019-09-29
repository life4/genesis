# Slice.FindIndex

```go
func (s Slice) FindIndex(f func(el T) bool) (int, error)
```

FindIndex is like Find, but return element index instead of element itself

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
| ErrNotFound | given element is not found |

## Source

```go
// FindIndex is like Find, but return element index instead of element itself
func (s Slice) FindIndex(f func(el T) bool) (int, error) {
	for i, el := range s.Data {
		if f(el) {
			return i, nil
		}
	}
	return 0, ErrNotFound
}
```

