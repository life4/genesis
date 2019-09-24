package implementation

import (
	"context"
	"sync"
)

// AsyncSlice is a set of operations to work with slice asynchronously
type AsyncSlice struct {
	Data    []T
	Workers int
}

// All returns true if f returns true for all elements in slice
func (s AsyncSlice) All(f func(el T) bool) bool {
	if len(s.Data) == 0 {
		return true
	}

	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int, result chan<- bool, ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case index, ok := <-jobs:
				if !ok {
					return
				}
				if !f(s.Data[index]) {
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
	workers := s.Workers
	if workers == 0 || workers > len(s.Data) {
		workers = len(s.Data)
	}

	// run workers
	jobs := make(chan int, len(s.Data))
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
	for i := 0; i < len(s.Data); i++ {
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
	if len(s.Data) == 0 {
		return false
	}

	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int, result chan<- bool, ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case index, ok := <-jobs:
				if !ok {
					return
				}
				if f(s.Data[index]) {
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
	workers := s.Workers
	if workers == 0 || workers > len(s.Data) {
		workers = len(s.Data)
	}

	// run workers
	jobs := make(chan int, len(s.Data))
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
	for i := 0; i < len(s.Data); i++ {
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
		defer wg.Done()
		for index := range jobs {
			f(s.Data[index])
		}
	}

	// calculate workers count
	workers := s.Workers
	if workers == 0 || workers > len(s.Data) {
		workers = len(s.Data)
	}

	// run workers
	jobs := make(chan int, len(s.Data))
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.Data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

// Filter returns slice of element for which f returns true
func (s AsyncSlice) Filter(f func(el T) bool) []T {
	resultMap := make([]bool, len(s.Data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		for index := range jobs {
			if f(s.Data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// calculate workers count
	workers := s.Workers
	if workers == 0 || workers > len(s.Data) {
		workers = len(s.Data)
	}

	// run workers
	jobs := make(chan int, len(s.Data))
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.Data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]T, 0, len(s.Data))
	for i, el := range s.Data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func (s AsyncSlice) Map(f func(el T) G) []G {
	result := make([]G, len(s.Data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		for index := range jobs {
			result[index] = f(s.Data[index])
		}
		wg.Done()
	}

	// calculate workers count
	workers := s.Workers
	if workers == 0 || workers > len(s.Data) {
		workers = len(s.Data)
	}

	// run workers
	jobs := make(chan int, len(s.Data))
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.Data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

// Reduce reduces slice to a single value with f
func (s AsyncSlice) Reduce(f func(left T, right T) T) T {
	if len(s.Data) == 0 {
		var tmp T
		return tmp
	}

	state := make([]T, len(s.Data))
	state = append(state, s.Data...)
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int, result chan<- T) {
		for index := range jobs {
			result <- f(state[index], state[index+1])
		}
		wg.Done()
	}

	for len(state) > 1 {
		// calculate workers count
		workers := s.Workers
		if workers == 0 || workers > len(state) {
			workers = len(state)
		}

		// run workers
		jobs := make(chan int, len(state))
		wg.Add(workers)
		result := make(chan T, 1)
		for i := 0; i < workers; i++ {
			go worker(jobs, result)
		}

		go func() {
			wg.Wait()
			close(result)
		}()

		// add indices into jobs for workers
		for i := 0; i < len(state)-1; i += 2 {
			jobs <- i
		}
		close(jobs)

		// collect new state
		newState := make([]T, 0, len(state)/2+len(state)%2)
		for el := range result {
			newState = append(newState, el)
		}
		if len(state)%2 == 1 {
			newState = append(newState, state[len(state)-1])
		}
		// put new state as current state after all
		state = newState
	}

	return state[0]
}
