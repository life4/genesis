package channels

import (
	"context"

	"github.com/life4/genesis/constraints"
)

// All is an alias for [AllC] without a context.
func All[T any](c <-chan T, f func(el T) bool) bool {
	return AllC(context.Background(), c, f)
}

// Any returns true if f returns true for any element in channel.
func Any[T any](c <-chan T, f func(el T) bool) bool {
	return AnyC(context.Background(), c, f)
}

// BufferSize returns how many messages a channel can hold before being blocked.
//
// When you push that many messages in the channel, the next attempt
// to write in it will block until another goroutine reads a message.
//
// For unbuffered channels the result is 0.
func BufferSize[T any](c <-chan T) int {
	return cap(c)
}

// ChunkEvery is an alias for [ChunkEveryC] without a context.
func ChunkEvery[T any](c <-chan T, count int) chan []T {
	return ChunkEveryC(context.Background(), c, count)
}

// Close safely closes the given channel.
//
// This is a safer version of built-in close. The built-in close function
// will panic if the given channel is already closed or nil.
// This function in both cases returns false.
//
// There is a reason why the built-in function might panic.
// The best practice of working with channels is to close the channel
// only within the context ("owner") that created it.
// In such cases, you can be 100% sure that the channel is non-nil
// and closed only once. However, if you found yourself in a different
// situation and no refactoring can guarantee these rules,
// the safer Close function is here to help.
func Close[T any](c chan<- T) bool {
	defer func() {
		_ = recover()
	}()
	close(c)
	return true
}

// Count is an alias for [CountC] without a context.
func Count[T comparable](c <-chan T, el T) int {
	return CountC(context.Background(), c, el)
}

// Drop is an alias for [DropC] without a context.
func Drop[T any](c <-chan T, n int) chan T {
	return DropC(context.Background(), c, n)
}

// Each is an alias for [EachC] without a context.
func Each[T any](c <-chan T, f func(el T)) {
	for el := range c {
		f(el)
	}
}

// Echo is an alias for [EchoC] without a context.
func Echo[T any](from <-chan T, to chan<- T) {
	EchoC(context.Background(), from, to)
}

// Filter is an alias for [FilterC] without a context.
func Filter[T any](c <-chan T, f func(el T) bool) chan T {
	return FilterC(context.Background(), c, f)
}

// First is an alias for [FirstC] without a context.
func First[T any](cs ...<-chan T) (T, error) {
	return FirstC(context.Background(), cs...)
}

// Flatten is an alias for [FlattenC] without a context.
func Flatten[T any](c <-chan <-chan T) chan T {
	return FlattenC(context.Background(), c)
}

// IsEmpty returns true if there are no messages in the channel.
//
// For unbuffered channels, the result is always true.
func IsEmpty[T any](c <-chan T) bool {
	return len(c) == 0
}

// IsFull returns true if the channel's buffer is full.
//
// Attempts to write into a full channel will block
// until another goroutine reads a message from it.
//
// For unbuffered channels, the result is always true.
func IsFull[T any](c <-chan T) bool {
	return len(c) == cap(c)
}

// Map is an alias for [MapC] without a context.
func Map[T any, G any](c <-chan T, f func(el T) G) chan G {
	return MapC(context.Background(), c, f)
}

// Max is an alias for [MaxC] without a context.
func Max[T constraints.Ordered](c <-chan T) (T, error) {
	return MaxC(context.Background(), c)
}

// Merge is an alias for [MergeC] without a context.
func Merge[T any](cs ...<-chan T) chan T {
	return MergeC(context.Background(), cs...)
}

// Min is an alias for [MinC] without a context.
func Min[T constraints.Ordered](c <-chan T) (T, error) {
	return MinC(context.Background(), c)
}

// Reduce is an alias for [ReduceC] without a context.
func Reduce[T any, G any](c <-chan T, acc G, f func(el T, acc G) G) G {
	return ReduceC(context.Background(), c, acc, f)
}

// Scan is an alias for [ScanC] without a context.
func Scan[T any, G any](c <-chan T, acc G, f func(el T, acc G) G) chan G {
	return ScanC(context.Background(), c, acc, f)
}

// Sum is an alias for [SumC] without a context.
func Sum[T constraints.Ordered](c <-chan T) T {
	return SumC(context.Background(), c)
}

// Take is an alias for [TakeC] without a context.
func Take[T any](c <-chan T, count int) chan T {
	return TakeC(context.Background(), c, count)
}

// Tee is an alias for [TeeC] without a context.
func Tee[T any](c <-chan T, count int) []chan T {
	return TeeC(context.Background(), c, count)
}

// ToSlice is an alias for [ToSliceC] without a context.
func ToSlice[T any](c <-chan T) []T {
	return ToSliceC(context.Background(), c)
}

// WithBuffer is an alias for [WithBufferC] without a context.
func WithBuffer[T any](c <-chan T, bufSize int) chan T {
	return WithBufferC(context.Background(), c, bufSize)
}
