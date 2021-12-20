package implementation

import (
	"context"
	"sync"
)

// All returns true if f returns true for all elements in slice
func All[S ~[]T, T any](items S, workers int, f func(el T) bool) bool {
	if len(items) == 0 {
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
				if !f(items[index]) {
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
	if workers <= 0 || workers > len(items) {
		workers = len(items)
	}

	// run workers
	jobs := make(chan int, len(items))
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
	for i := 0; i < len(items); i++ {
		jobs <- i
	}
	close(jobs)

	for range result {
		return false
	}
	return true
}

// Any returns true if f returns true for any element from slice
func Any[S ~[]T, T any](items S, workers int, f func(el T) bool) bool {
	if len(items) == 0 {
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
				if f(items[index]) {
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
	if workers <= 0 || workers > len(items) {
		workers = len(items)
	}

	// run workers
	jobs := make(chan int, len(items))
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
	for i := 0; i < len(items); i++ {
		jobs <- i
	}
	close(jobs)

	for range result {
		return true
	}
	return false
}

// Each calls f for every element from slice
func Each[S ~[]T, T any](items S, workers int, f func(el T)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		defer wg.Done()
		for index := range jobs {
			f(items[index])
		}
	}

	// calculate workers count
	if workers <= 0 || workers > len(items) {
		workers = len(items)
	}

	// run workers
	jobs := make(chan int, len(items))
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(items); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

// Filter returns slice of element for which f returns true
func Filter[S ~[]T, T any](items S, workers int, f func(el T) bool) S {
	resultMap := make([]bool, len(items))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		for index := range jobs {
			if f(items[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// calculate workers count
	if workers <= 0 || workers > len(items) {
		workers = len(items)
	}

	// run workers
	jobs := make(chan int, len(items))
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(items); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]T, 0, len(items))
	for i, el := range items {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func Map[S ~[]T, T any, G any](items S, workers int, f func(el T) G) []G {
	result := make([]G, len(items))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		for index := range jobs {
			result[index] = f(items[index])
		}
		wg.Done()
	}

	// calculate workers count
	if workers <= 0 || workers > len(items) {
		workers = len(items)
	}

	// run workers
	jobs := make(chan int, len(items))
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(items); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

// Reduce reduces slice to a single value with f
func Reduce[S ~[]T, T any](items S, workers int, f func(left T, right T) T) T {
	if len(items) == 0 {
		var tmp T
		return tmp
	}

	state := make([]T, len(items))
	state = append(state, items...)
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int, result chan<- T) {
		for index := range jobs {
			result <- f(state[index], state[index+1])
		}
		wg.Done()
	}

	for len(state) > 1 {
		// calculate workers count
		if workers <= 0 || workers > len(state) {
			workers = len(state)
		}

		// run workers
		jobs := make(chan int, len(state))
		wg.Add(workers)
		result := make(chan T)
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
