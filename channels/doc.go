// Package channels provides generic functions for channels.
//
// # ☎️ Naming
//
// Most of the functions provide two version: a regular one and one with suffix C.
// The latter also accepts a [context.Context] value as the first argument, and you should
// always prefer it over the regular function. The reason is that if the function
// accepts a channel and returns another one and you close the input channel,
// the internal goroutine that the function starts might be blocked trying to write
// into the output channel, and if you never read values from it, you'll have
// a goroutine leak.
//
// Generator functions producing infinite sequence of values (like [Counter])
// have only one version and always require a context because that's the only way
// how you can cancel them.
//
// # ⏹️ Function termination
//
// Most of the functions in the package accept a channel, return a channel,
// and create a goroutine to read messages from the input channel and propagate them
// (usually with some changes) into the output one. Some (like [TakeC]) do all three of these,
// some (like [ToSliceC]) only accept a channel, some (like [Counter]) only start a goroutine
// and return a channel, etc. Which of these apply to the function is described in the
// function's documentation.
//
// Here is what you need to know:
//
//   - If a function starts a goroutine, it returns immediately, and all termination rules
//     described below apply to the goroutine instead of the original function.
//   - If the function creates and returns a channel, this channel is closed
//     when the goroutine terminates.
//   - If the function accepts a [context.Context], the goroutine is terminated
//     when the context is cancelled.
//   - If the input channel is closed and the goroutine tries to read from it,
//     the goroutine is terminated.
//   - ⚠️ IMPORTANT: If the goroutine is blocked trying to write into the output channel
//     and the input channel is closed, the goroutine will not terminate until
//     another goroutine reads from the output channel. This is why you should always
//     either use context and cancell the goroutine through it or make sure to read
//     everything from the output channel.
//
// # 🍾 Buffered channels
//
// All functions in the package (except [WithBuffer]) return a unbuffered channel.
// That means, they won't do anything and won't read values from the input channel
// if you don't read values from their output channel.
// If you want to make the output channel buffered, use [WithBuffer].
//
// # 😱 Error handling
//
// The only functions that return an error on the channel cancellation
// (either [context.Canceled] or [context.Cause]) are [FirstC], [MaxC], and [MinC].
// All other functions simply terminate on cancellation, and it's on you to check
// the context for errors (using [context.Context.Err]).
//
// The reason for this is that most of the function start a goroutine producing values,
// and that's the only communication channel the goroutine has. So, the only way for
// the goroutine to return an error would be to emit it in the output channel,
// and handling such errors would be tedious.
//
// Available errors:
//
//   - [ErrEmpty] is returned when channel is closed without any elements being sent.
//   - [ErrClosed] is returned by [FirstC] (and [First]) when one of the passed channels
//     is closed.
//
// # 🖨 Sequence generators
//
// These are the functions that make a channel and emit in it values until canceled.
// In other languages this is done using iterators but in Go the only infinite iterator
// we have is channel.
//
// If you want to stop a generator (terminate the goroutine it started
// and close the channel), cancel the context you have passed in it.
//
// Available functions:
//
//   - [Counter](ctx, start, step)
//   - [Exponential](ctx, start, factor)
//   - [Iterate](ctx, value, func)
//   - [Range](ctx, start, end, step)
//   - [Repeat](ctx, value)
//   - [Replicate](ctx, value, n)
//
// # 📥 Working with channels
//
// All these functions accept a channel (or multiple) as an input.
//
// 🏃 Functions that return a channel and start an internal goroutine
// producing values into the channel:
//
//   - [ChunkEvery] and [ChunkEveryC]
//   - [Drop] and [DropC]
//   - [Filter] and [FilterC]
//   - [Flatten] and [FlattenC]
//   - [Map] and [MapC]
//   - [Merge] and [MergeC]
//   - [Scan] and [ScanC]
//   - [Take] and [TakeC]
//   - [Tee] and [TeeC]
//   - [WithBuffer] and [WithBufferC]
//   - [WithContext]
//
// ✋ Functions that block until they produce a result (or until they are canceled):
//
//   - [Any] and [AnyC]
//   - [All] and [AllC]
//   - [Count] and [CountC]
//   - [Each] and [EachC]
//   - [Echo] and [EchoC]
//   - [First] and [FirstC]
//   - [Max] and [MaxC]
//   - [Min] and [MinC]
//   - [Reduce] and [ReduceC]
//   - [Sum] and [SumC]
//   - [ToSlice] and [ToSliceC]
//
// # 🛟 Helpers
//
// These are little functions that help working with channels.
// They don't iterate over values in the channel.
//
// Available functions:
//
//   - [BufferSize] returns the channel's buffer size.
//   - [Close] closes the channel and never panics.
//   - [IsEmpty] tells if there are no messages in the channel.
//   - [IsFull] tells if the channel has reached its capacity.
//   - [Pop] is a blocking read from the channel that can be canceled.
//   - [Push] is a blocking write into a channel that can be canceled.
package channels
