# AsyncSlice.Each

```go
func (s AsyncSlice) Each(f func(el T))
```

Each calls f for every element from slice

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
```

## Tests

```go
func TestAsyncSliceEach(t *testing.T) {
	f := func(given []T) {
		s := AsyncSlice{Data: given, Workers: 2}
		result := make(chan T, len(given))
		mapper := func(t T) { result <- t }
		s.Each(mapper)
		close(result)
		actual := Channel{result}.ToSlice()
		sorted := Slice{actual}.Sort()
		assert.Equal(t, given, sorted, "they should be equal")
	}

	f([]T{})
	f([]T{1})
	f([]T{1, 2, 3})
	f([]T{1, 2, 3, 4, 5, 6, 7})
}
```