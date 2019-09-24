# AsyncSlice

AsyncSlice is a set of operations to work with slice asynchronously

```go
// AsyncSlice is a set of operations to work with slice asynchronously
type AsyncSlice struct {
	Data    []T
	Workers int
}
```

## Structs

| Struct | T type |
| ------ | ------ |
| AsyncSliceBool | bool |
| AsyncSliceByte | byte |
| AsyncSliceString | string |
| AsyncSliceFloat32 | float32 |
| AsyncSliceFloat64 | float64 |
| AsyncSliceInt | int |
| AsyncSliceInt8 | int8 |
| AsyncSliceInt16 | int16 |
| AsyncSliceInt32 | int32 |
| AsyncSliceInt64 | int64 |
| AsyncSliceUint | uint |
| AsyncSliceUint8 | uint8 |
| AsyncSliceUint16 | uint16 |
| AsyncSliceUint32 | uint32 |
| AsyncSliceUint64 | uint64 |
| AsyncSliceInterface | interface{} |

## Functions

| Function | Description |
| -------- | ----------- |
| [All](./all.md) | All returns true if f returns true for all elements in slice |
| [Any](./any.md) | Any returns true if f returns true for any element from slice |
| [Each](./each.md) | Each calls f for every element from slice |
| [Filter](./filter.md) | Filter returns slice of element for which f returns true |
| [Map](./map.md) | Map applies F to all elements in slice of T and returns slice of results |
| [Reduce](./reduce.md) | Reduce reduces slice to a single value with f |
