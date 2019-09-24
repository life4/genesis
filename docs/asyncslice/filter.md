# AsyncSlice.Filter

Filter returns slice of element for which f returns true

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

