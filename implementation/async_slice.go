package implementation

import (
	"context"
	"sync"
)

// AsyncSlice is a set of operations to work with slice asynchronously
type AsyncSlice struct {
	data    []T
	workers int
}

// All returns true if f returns true for all elements in slice
func (s AsyncSlice) All(f func(el T) bool) bool {
	if len(s.data) == 0 {
		return true
	}

	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int, result chan<- bool, ctx context.Context) {
		// enter
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case index, ok := <-jobs:
				if !ok {
					return
				}
				if !f(s.data[index]) {
					result <- false
					return
				}
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	// when we're returning the result, cancel all workers
	defer cancel()

	// calculate workers count
	workers := s.workers
	if workers == 0 || workers > len(s.data) {
		workers = len(s.data)
	}

	// run workers
	jobs := make(chan int, len(s.data))
	result := make(chan bool, workers)
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(jobs, result, ctx)
	}

	// close the result channel when all workers have done
	go func() {
		wg.Wait()
		close(result)
	}()

	// schedule the jobs: indices to check
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)

	for range result {
		return false
	}
	return true
}

// Any returns true if f returns true for any element from slice
func (s AsyncSlice) Any(f func(el T) bool) bool {
	if len(s.data) == 0 {
		return false
	}

	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int, result chan<- bool, ctx context.Context) {
		// enter
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case index, ok := <-jobs:
				if !ok {
					return
				}
				if f(s.data[index]) {
					result <- true
					return
				}
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	// when we're returning the result, cancel all workers
	defer cancel()

	// calculate workers count
	workers := s.workers
	if workers == 0 || workers > len(s.data) {
		workers = len(s.data)
	}

	// run workers
	jobs := make(chan int, len(s.data))
	result := make(chan bool, workers)
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(jobs, result, ctx)
	}

	// close the result channel when all workers have done
	go func() {
		wg.Wait()
		close(result)
	}()

	// schedule the jobs: indices to check
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)

	for range result {
		return true
	}
	return false
}

// Each calls f for every element from slice
func (s AsyncSlice) Each(f func(el T)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

// Filter returns slice of element for which f returns true
func (s AsyncSlice) Filter(f func(el T) bool) []T {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]T, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func (s AsyncSlice) Map(f func(el T) G) []G {
	result := make([]G, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 || workers > len(s.data) {
		workers = len(s.data)
	}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}
