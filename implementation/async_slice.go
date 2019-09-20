package implementation

import "sync"

// AsyncSlice is a set of operations to work with slice asynchronously
type AsyncSlice struct {
	data    []T
	workers int
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
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	wg.Wait()
}

// Map applies F to all elements in slice of T and returns slice of results
func (s AsyncSlice) Map(f func(el T) G) []G {
	result := make([]G, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	wg.Wait()
	return result
}
