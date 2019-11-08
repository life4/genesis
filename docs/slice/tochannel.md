# Slice.ToChannel

```go
func (s Slice) ToChannel() chan T
```

ToChannel returns channel with elements from the slice

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

## Source

```go
// ToChannel returns channel with elements from the slice
func (s Slice) ToChannel() chan T {
	c := make(chan T, 1)
	go func() {
		for _, el := range s.Data {
			c <- el
		}
		close(c)
	}()
	return c
}
```

