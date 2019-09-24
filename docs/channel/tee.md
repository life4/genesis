# Channel.Tee

Tee returns 2 channels with elements from the input channel

## Source

```go
// Tee returns 2 channels with elements from the input channel
func (c Channel) Tee(count int) []chan T {
	channels := make([]chan T, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan T, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan T) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}
```