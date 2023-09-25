// 🦥 Package iters provides generic functions for lazy iteration.
//
// Iterators are useful for single-threaded processing of large amount of items.
// For a small collection, slices should give better performance for the price of
// higher memory consumption. For big collection with possibility of concurrency,
// you can get better performance by using channels and goroutines.
package iters
