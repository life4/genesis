# Channel.Scan

```go
func (c Channel) Scan(acc G, f func(el T, acc G) G) chan G
```

Scan is like Reduce, but returns slice of f results

Generic types: G, T.

## Structs

| Struct | T type |
| ------ | ------ |
| ChannelBool | bool |
| ChannelByte | byte |
| ChannelString | string |
| ChannelFloat32 | float32 |
| ChannelFloat64 | float64 |
| ChannelInt | int |
| ChannelInt8 | int8 |
| ChannelInt16 | int16 |
| ChannelInt32 | int32 |
| ChannelInt64 | int64 |
| ChannelUint | uint |
| ChannelUint8 | uint8 |
| ChannelUint16 | uint16 |
| ChannelUint32 | uint32 |
| ChannelUint64 | uint64 |
| ChannelInterface | interface{} |

## Functions

| Function | G type |
| -------- | ------ |
| ScanBool | bool |
| ScanByte | byte |
| ScanString | string |
| ScanFloat32 | float32 |
| ScanFloat64 | float64 |
| ScanInt | int |
| ScanInt8 | int8 |
| ScanInt16 | int16 |
| ScanInt32 | int32 |
| ScanInt64 | int64 |
| ScanUint | uint |
| ScanUint8 | uint8 |
| ScanUint16 | uint16 |
| ScanUint32 | uint32 |
| ScanUint64 | uint64 |
| ScanInterface | interface{} |

## Source

```go
// Scan is like Reduce, but returns slice of f results
func (c Channel) Scan(acc G, f func(el T, acc G) G) chan G {
	result := make(chan G, 1)
	go func() {
		for el := range c.Data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}
```

