# Channel

Channel is a set of operations with channel

```go
// Channel is a set of operations with channel
type Channel struct {
	data chan T
}
```

| Function | Description |
| -------- | ----------- |
| Any | Any returns true if f returns true for any element in channel |
| All | All returns true if f returns true for all elements in channel |
| ChunkEvery | ChunkEvery returns channel with slices containing count elements each |
| Count | Count return count of el occurences in channel. |
| Drop | Drop drops first n elements from channel c and returns a new channel with the rest. It returns channel do be unblocking. If you want array instead, wrap result into TakeAll. |
| Each | Each calls f for every element in the channel |
| Filter | Filter returns a new channel with elements from input channel for which f returns true |
| Map | Map applies f to all elements from channel and returns channel with results |
| Max | Max returns the maximal element from channel |
| Min | Min returns the minimal element from channel |
| Reduce | Reduce applies f to acc and every element from channel and returns acc |
| Scan | Scan is like Reduce, but returns slice of f results |
| Sum | Sum returns sum of all elements from channel |
| Take | Take takes first n elements from channel c. |
| Tee | Tee returns 2 channels with elements from the input channel |
| ToSlice | ToSlice returns slice with all elements from channel. |
