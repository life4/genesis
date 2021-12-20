// Package aslices defines concurrent generic functions for slices.
//
// Each function accepts a slice to perform operations on, and int `workers`
// indicating how many workers to spawn. The operation that the function performs
// is performed concurrently (in multiple goroutines, as specified by `workers`),
// with work distribution, wait groups, cancellation, and other best practice.
package aslices
