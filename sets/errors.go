package sets

import (
	"errors"
)

// ErrEmpty is an error for empty set when it's expected to have elements
var ErrEmpty = errors.New("container is empty")
