# Slice.Cycle

```go
func (s Slice) Cycle() chan T
```

Cycle is an infinite loop over slice

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
// Cycle is an infinite loop over slice
func (s Slice) Cycle() chan T {
	c := make(chan T, 1)
	go func() {
		for {
			for _, val := range s.Data {
				c <- val
			}
		}
	}()
	return c
}
```

