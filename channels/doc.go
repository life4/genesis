// Package channels provides generic functions for channels.
//
// The symbol ⏹️ indicates paragraph describing cancellation
// of goroutines and channels that the function creates.
// Most of the functions in the package are cancelled
// when the passed in channel is closed.
// If you want to cancel them using a Context instead,
// wrap the passed in channel into [WithContext].
//
// The symbol ⏸️ indicates a notice about unbuffered channels.
// All functions in the package return unbuffered channels,
// regardless of if the input channel is buffered or not.
// If you want to make the returned channel buffered,
// wrap it into [WithBuffer].
package channels
