# AsyncSlice.Reduce

```go
func (s AsyncSlice) Reduce(f func(left T, right T) T) T
```

Reduce reduces slice to a single value with f

Generic types: T.

## Structs

| Struct | T type |
| ------ | ------ |
| AsyncSliceBool | bool |
| AsyncSliceByte | byte |
| AsyncSliceString | string |
| AsyncSliceRune | rune |
| AsyncSliceError | error |
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
// Reduce reduces slice to a single value with f
func (s AsyncSlice) Reduce(f func(left T, right T) T) T {
	if len(s.Data) == 0 {
		var tmp T
		return tmp
	}

	state := make([]T, len(s.Data))
	state = append(state, s.Data...)
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int, result chan<- T) {
		for index := range jobs {
			result <- f(state[index], state[index+1])
		}
		wg.Done()
	}

	for len(state) > 1 {
		// calculate workers count
		workers := s.Workers
		if workers == 0 || workers > len(state) {
			workers = len(state)
		}

		// run workers
		jobs := make(chan int, len(state))
		wg.Add(workers)
		result := make(chan T, 1)
		for i := 0; i < workers; i++ {
			go worker(jobs, result)
		}

		go func() {
			wg.Wait()
			close(result)
		}()

		// add indices into jobs for workers
		for i := 0; i < len(state)-1; i += 2 {
			jobs <- i
		}
		close(jobs)

		// collect new state
		newState := make([]T, 0, len(state)/2+len(state)%2)
		for el := range result {
			newState = append(newState, el)
		}
		if len(state)%2 == 1 {
			newState = append(newState, state[len(state)-1])
		}
		// put new state as current state after all
		state = newState
	}

	return state[0]
}
```

## Tests

```go
func TestAsyncSliceReduce(t *testing.T) {
	f := func(reducer func(a T, b T) T, given []T, expected T) {
		s := AsyncSlice{Data: given, Workers: 4}
		actual := s.Reduce(reducer)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	sum := func(a T, b T) T { return a + b }

	f(sum, []T{}, 0)
	f(sum, []T{1}, 1)
	f(sum, []T{1, 2}, 3)
	f(sum, []T{1, 2, 3}, 6)
	f(sum, []T{1, 2, 3, 4}, 10)
	f(sum, []T{1, 2, 3, 4, 5}, 15)
}
```