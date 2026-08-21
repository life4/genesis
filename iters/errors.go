package iters

import (
	"errors"
)

// ErrNotFound is an error for case when given element is not found
var ErrNotFound = errors.New("given element is not found")

// ErrOutOfRange is an error that for index bigger than slice size
var ErrOutOfRange = errors.New("index is bigger than container size")

// ErrEmpty is an error for empty slice when it's expected to have elements
var ErrEmpty = errors.New("container is empty")
