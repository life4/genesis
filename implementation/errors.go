package implementation

import (
	"errors"
)

// ErrNegativeIndex is an error for passed index <0
var ErrNegativeIndex = errors.New("negative index passed")

// ErrNotFound is an error for case when given element is not found
var ErrNotFound = errors.New("negative index passed")

// ErrNonPositiveStep is an error for passed step <=0
var ErrNonPositiveStep = errors.New("step must be positive")

// ErrIndexOutOfRange is an error that for index bigger than slice size
var ErrIndexOutOfRange = errors.New("index is bigger than slice size")

// ErrEmptySlice is an error for empty slice when it's expected to have elements
var ErrEmptySlice = errors.New("slice is empty")
