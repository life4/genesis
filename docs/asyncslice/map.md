# AsyncSlice.Map

Map applies F to all elements in slice of T and returns slice of results

## Source

```go
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

	// calculate workers count
	workers := s.workers
	if workers == 0 || workers > len(s.data) {
		workers = len(s.data)
	}

	// run workers
	jobs := make(chan int, len(s.data))
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
```