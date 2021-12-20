package channels

import (
	"errors"
)

// ErrEmpty is an error for empty slice when it's expected to have elements
var ErrEmpty = errors.New("container is empty")
