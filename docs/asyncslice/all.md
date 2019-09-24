# AsyncSlice.All

```go
func (s AsyncSlice) All(f func(el T) bool) bool
```

All returns true if f returns true for all elements in slice

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| AsyncSliceBool | bool |
| AsyncSliceByte | byte |
| AsyncSliceString | string |
| AsyncSliceFloat32 | float32 |
| AsyncSliceFloat64 | float64 |
| AsyncSliceInt | int |
| AsyncSliceInt8 | int8 |
| AsyncSliceInt16 | int16 |
| AsyncSliceInt32 | int32 |
| AsyncSliceInt64 | int64 |
| AsyncSliceUint | uint |
| AsyncSliceUint8 | uint8 |
| AsyncSliceUint16 | uint16 |
| AsyncSliceUint32 | uint32 |
| AsyncSliceUint64 | uint64 |
| AsyncSliceInterface | interface{} |

## Source

```go
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
	workers := s.workers
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
```

## Tests

```go
func TestAsyncSliceAll(t *testing.T) {
	f := func(check func(t T) bool, given []T, expected bool) {
		s := AsyncSlice{Data: given, workers: 2}
		actual := s.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t T) bool { return (t % 2) == 0 }

	f(isEven, []T{}, true)
	f(isEven, []T{1}, false)
	f(isEven, []T{1, 3}, false)
	f(isEven, []T{2}, true)
	f(isEven, []T{2, 4}, true)
	f(isEven, []T{2, 3}, false)
	f(isEven, []T{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []T{2, 4, 6, 8, 10, 11}, false)
}
```