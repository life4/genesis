# AsyncSlice.Map

```go
func (s AsyncSlice) Map(f func(el T) G) []G
```

Map applies F to all elements in slice of T and returns slice of results

Generic types: G, T.

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

## Functions

| Function | G type |
| -------- | ------ |
| MapBool | bool |
| MapByte | byte |
| MapString | string |
| MapFloat32 | float32 |
| MapFloat64 | float64 |
| MapInt | int |
| MapInt8 | int8 |
| MapInt16 | int16 |
| MapInt32 | int32 |
| MapInt64 | int64 |
| MapUint | uint |
| MapUint8 | uint8 |
| MapUint16 | uint16 |
| MapUint32 | uint32 |
| MapUint64 | uint64 |
| MapInterface | interface{} |

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

## Tests

```go
func TestAsyncSliceMap(t *testing.T) {
	f := func(mapper func(t T) G, given []T, expected []G) {
		s := AsyncSlice{data: given, workers: 2}
		actual := s.Map(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t T) G { return G((t * 2)) }

	f(double, []T{}, []G{})
	f(double, []T{1}, []G{2})
	f(double, []T{1, 2, 3}, []G{2, 4, 6})
}
```