# AsyncSlice.Each

Each calls f for every element from slice

## Source

```go
// Each calls f for every element from slice
func (s AsyncSlice) Each(f func(el T)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		defer wg.Done()
		for index := range jobs {
			f(s.data[index])
		}
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
}
```