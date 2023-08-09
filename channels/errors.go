package channels

import (
	"errors"
)

// ErrEmpty is an error for when channel is closed without any elements being sent
var ErrEmpty = errors.New("container is empty")

// ErrEmpty is an error for when a channel is closed.
// Currently is used only by [First] to distinguish a closed channel
// from a zero value returned without adding a third `bool` return value.
var ErrClosed = errors.New("channel is closed")
