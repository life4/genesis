# Channel

Channel is a set of operations with channel

```go
// Channel is a set of operations with channel
type Channel struct {
	Data chan T
}
```

## Structs

| Struct | T type |
| ------ | ------ |
| ChannelBool | bool |
| ChannelByte | byte |
| ChannelString | string |
| ChannelRune | rune |
| ChannelError | error |
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

| Function | Description |
| -------- | ----------- |
| [Any](./any.md) | Any returns true if f returns true for any element in channel |
| [All](./all.md) | All returns true if f returns true for all elements in channel |
| [ChunkEvery](./chunkevery.md) | ChunkEvery returns channel with slices containing count elements each |
| [Count](./count.md) | Count return count of el occurences in channel. |
| [Drop](./drop.md) | Drop drops first n elements from channel c and returns a new channel with the rest. It returns channel do be unblocking. If you want array instead, wrap result into TakeAll. |
| [Each](./each.md) | Each calls f for every element in the channel |
| [Filter](./filter.md) | Filter returns a new channel with elements from input channel for which f returns true |
| [Map](./map.md) | Map applies f to all elements from channel and returns channel with results |
| [Max](./max.md) | Max returns the maximal element from channel |
| [Min](./min.md) | Min returns the minimal element from channel |
| [Reduce](./reduce.md) | Reduce applies f to acc and every element from channel and returns acc |
| [Scan](./scan.md) | Scan is like Reduce, but returns slice of f results |
| [Sum](./sum.md) | Sum returns sum of all elements from channel |
| [Take](./take.md) | Take takes first count elements from the channel. |
| [Tee](./tee.md) | Tee returns 2 channels with elements from the input channel |
| [ToSlice](./toslice.md) | ToSlice returns slice with all elements from channel. |
