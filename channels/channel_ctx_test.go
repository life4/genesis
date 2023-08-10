package channels_test

import (
	"context"
	"sync"
	"testing"

	"github.com/life4/genesis/channels"
	"github.com/matryer/is"
)

func TestFirstC(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	ctx := context.Background()
	f := func(n, i int) {
		csW := make([]chan<- int, n)
		csR := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			c := make(chan int)
			csW[i] = c
			csR[i] = c
		}
		go func() { csW[i] <- 12 }()
		v, err := channels.FirstC(ctx, csR...)
		is.NoErr(err)
		is.Equal(v, 12)
	}
	for n := 1; n < 10; n++ {
		for i := 0; i < n; i++ {
			f(n, i)
		}
	}

	_, err := channels.FirstC[int](ctx)
	is.Equal(err, channels.ErrEmpty)
}

func TestFirstC_Cancellation(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	ctx := context.Background()

	// one closed channel
	c1 := make(chan int)
	close(c1)
	_, err := channels.FirstC(ctx, c1)
	is.Equal(err, channels.ErrClosed)

	// two channels, one closed
	c2 := make(chan int)
	_, err = channels.FirstC(ctx, c2, c1)
	is.Equal(err, channels.ErrClosed)

	// one channel, context cancelled on wait
	ctx2, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = channels.FirstC(ctx2, c2)
	is.Equal(err, context.Canceled)

	// two channels, context cancelled on wait
	c3 := make(chan int)
	c4 := make(chan int)
	ctx3, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = channels.FirstC(ctx3, c3, c4)
	is.Equal(err, context.Canceled)
}

func TestFirstC_Starvation(t *testing.T) {
	t.Parallel()
	is := is.New(t)

	nChannels := 50
	nMessages := 50
	bufSize := 2

	// create channels
	csW := make([]chan<- int, nChannels)
	csR := make([]<-chan int, nChannels)
	for i := 0; i < nChannels; i++ {
		c := make(chan int, bufSize)
		csW[i] = c
		csR[i] = c
	}

	// For each channel, start a job that will emit
	// in the channel this channel's ID.
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < nChannels; i++ {
		go func(i int) {
			for m := 0; m < nMessages; m++ {
				select {
				case csW[i] <- i:
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}

	// Calculate how many messages are received from each channel.
	nReceived := make(map[int]uint32)
	for k := 0; k < nChannels*4; k++ {
		i, err := channels.FirstC(context.Background(), csR...)
		is.NoErr(err)
		nReceived[i] += 1
	}

	// Check if there are any starved channels.
	// A starved channel is the one that couldn't get
	// any messages selected from it.
	cancel()
	for i, n := range nReceived {
		if n == 0 {
			t.Errorf("channel %d is starved (n=%d)", i, n)
		}
	}
}

func TestMergeC(t *testing.T) {
	is := is.New(t)

	// no channels passed
	ctx := context.Background()
	c1 := channels.MergeC[int](ctx)
	_, more := <-c1
	is.True(!more)

	// passed channels are all closed
	c2 := make(chan int)
	c3 := make(chan int)
	close(c2)
	close(c3)
	c4 := channels.MergeC(ctx, c2, c3)
	_, more = <-c4
	is.True(!more)

	// one channel
	c5 := make(chan int)
	go func() {
		c5 <- 42
		close(c5)
	}()
	c6 := channels.MergeC(ctx, c5)
	v := <-c6
	is.Equal(v, 42)
	_, more = <-c6
	is.True(!more)

	// cancel on read
	ctx2, cancel := context.WithCancel(context.Background())
	c7 := channels.MergeC(ctx2, make(<-chan int))
	cancel()
	_, more = <-c7
	is.True(!more)

	// cancel on write
	ctx3, cancel := context.WithCancel(context.Background())
	c8 := make(chan int)
	c9 := channels.MergeC(ctx3, c8)
	c8 <- 42
	cancel()
	// Potential race? Let's see if it get flaky.
	_, more = <-c9
	is.True(!more)
}

func TestPop(t *testing.T) {
	is := is.New(t)

	// exit on canceled context
	c := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	v, ok := channels.Pop(ctx, c)
	is.True(!ok)
	is.Equal(v, 0)

	// succesful read
	ctx = context.Background()
	go func() { c <- 12 }()
	v, ok = channels.Pop(ctx, c)
	is.True(ok)
	is.Equal(v, 12)

	// exit on closed channel
	ctx = context.Background()
	close(c)
	v, ok = channels.Pop(ctx, c)
	is.True(!ok)
	is.Equal(v, 0)
}

func TestPush(t *testing.T) {
	// exit on canceled context
	c := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	channels.Push(ctx, c, 11)

	// push into the channel when there is someone to read
	ctx = context.Background()
	go channels.Push(ctx, c, 12)
	is := is.New(t)
	is.Equal(<-c, 12)
}

// WithContext should echo everything from the input channel
// and close the resulting channel when the input one is closed.
func TestWithContext_NoCancellation(t *testing.T) {
	is := is.New(t)
	c1 := make(chan int)
	c2 := channels.WithContext(c1, context.Background())

	r := make([]int, 0)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range c2 {
			r = append(r, i)
		}
	}()

	c1 <- 11
	c1 <- 12
	c1 <- 13
	close(c1)
	wg.Wait()
	is.Equal(r, []int{11, 12, 13})
}

// WithContext should exit on cancellation.
func TestWithContext_Cancellation(t *testing.T) {
	is := is.New(t)
	c1 := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	c2 := channels.WithContext(c1, ctx)

	r := make([]int, 0)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range c2 {
			r = append(r, i)
			if i == 12 {
				cancel()
			}
		}
	}()

	c1 <- 11
	c1 <- 12
	c1 <- 13
	wg.Wait()
	is.Equal(r, []int{11, 12})
	close(c1)
}

func TestWithContext_CancelWaitingInput(t *testing.T) {
	is := is.New(t)
	c1 := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	c2 := channels.WithContext(c1, ctx)
	cancel()
	_, more := <-c2
	is.True(!more)
}
