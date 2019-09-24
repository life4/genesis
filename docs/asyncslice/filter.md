# AsyncSlice.Filter

Filter returns slice of element for which f returns true

## Source

```go
// Filter returns slice of element for which f returns true
func (s AsyncSlice) Filter(f func(el T) bool) []T {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
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

	// return filtered results
	result := make([]T, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}
```