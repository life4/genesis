# AsyncSlice.Any

Any returns true if f returns true for any element from slice

## Source

```go
// Any returns true if f returns true for any element from slice
func (s AsyncSlice) Any(f func(el T) bool) bool {
	if len(s.data) == 0 {
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
```