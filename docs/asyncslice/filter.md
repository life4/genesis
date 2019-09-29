# AsyncSlice.Filter

```go
func (s AsyncSlice) Filter(f func(el T) bool) []T
```

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
```

## Tests

```go
func TestAsyncSliceFilter(t *testing.T) {
	f := func(given []T, expected []T) {
		filter := func(t T) bool { return t > 10 }
		s := AsyncSlice{data: given, workers: 2}
		actual := s.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]T{}, []T{})
	f([]T{5}, []T{})
	f([]T{15}, []T{15})
	f([]T{9, 11, 12, 13, 6}, []T{11, 12, 13})
}
```