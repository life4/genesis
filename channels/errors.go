package channels

import (
	"errors"
)

// ErrEmpty is an error for when channel is closed without any elements being sent
var ErrEmpty = errors.New("container is empty")
