package gchannels_test

import (
	"context"
	"testing"

	"github.com/life4/genesis/gchannels"
	"github.com/stretchr/testify/assert"
)

func TestChannelToSlice(t *testing.T) {
	f := func(given []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := gchannels.ToSlice(c)
		assert.Equal(t, given, actual)
	}
	f([]int{})
	f([]int{1})
	f([]int{1, 2, 3, 1, 2})
}

func TestChannelAny(t *testing.T) {
	f := func(given []int, expected bool) {
		even := func(t int) bool { return t%2 == 0 }
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := gchannels.Any(c, even)
		assert.Equal(t, expected, actual)
	}
	f([]int{}, false)
	f([]int{1}, false)
	f([]int{2}, true)
	f([]int{1, 2}, true)
	f([]int{1, 2, 3}, true)
	f([]int{1, 3, 5}, false)
	f([]int{1, 3, 5, 7, 9, 11}, false)
	f([]int{1, 3, 5, 7, 10, 11}, true)
}

func TestChannelAll(t *testing.T) {
	f := func(given []int, expected bool) {
		even := func(t int) bool { return t%2 == 0 }
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := gchannels.All(c, even)
		assert.Equal(t, expected, actual)
	}
	f([]int{}, true)
	f([]int{1}, false)
	f([]int{2}, true)
	f([]int{1, 2}, false)
	f([]int{2, 4}, true)
	f([]int{2, 4, 6, 8, 10, 12}, true)
	f([]int{2, 4, 6, 8, 11, 12}, false)
}

func TestChannelEach(t *testing.T) {
	f := func(given []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := make(chan int, len(given))
		mapper := func(t int) { result <- t }
		gchannels.Each(c, mapper)
		close(result)
		actual := gchannels.ToSlice(result)
		assert.Equal(t, given, actual)
	}

	f([]int{})
	f([]int{1})
	f([]int{1, 2, 3})
	f([]int{1, 2, 3, 4, 5, 6, 7})
}

func TestChannelChunkEvery(t *testing.T) {
	f := func(size int, given []int, expected [][]int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := gchannels.ChunkEvery(c, size)
		actual := make([][]int, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual)
	}
	f(2, []int{}, [][]int{})
	f(2, []int{1}, [][]int{{1}})
	f(2, []int{1, 2}, [][]int{{1, 2}})
	f(2, []int{1, 2, 3, 4}, [][]int{{1, 2}, {3, 4}})
	f(2, []int{1, 2, 3, 4, 5}, [][]int{{1, 2}, {3, 4}, {5}})
}

func TestChannelCount(t *testing.T) {
	f := func(element int, given []int, expected int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := gchannels.Count(c, element)
		assert.Equal(t, expected, actual)
	}
	f(1, []int{}, 0)
	f(1, []int{1}, 1)
	f(1, []int{2}, 0)
	f(1, []int{1, 2, 3, 1, 4}, 2)
}

func TestChannelDrop(t *testing.T) {
	f := func(count int, given []int, expected []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := gchannels.Drop(c, count)
		actual := make([]int, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual)
	}
	f(1, []int{}, []int{})
	f(1, []int{2}, []int{})
	f(1, []int{2, 3}, []int{3})
	f(1, []int{1, 2, 3}, []int{2, 3})
	f(0, []int{1, 2, 3}, []int{1, 2, 3})
	f(3, []int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6})
	f(1, []int{1, 2, 3, 4, 5, 6}, []int{2, 3, 4, 5, 6})
}

func TestChannelFilter(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(t int) bool { return t%2 == 0 }
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := gchannels.Filter(c, even)
		actual := gchannels.ToSlice(result)
		assert.Equal(t, expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{})
	f([]int{2}, []int{2})
	f([]int{1, 2, 3, 4}, []int{2, 4})
}

func TestChannelMap(t *testing.T) {
	f := func(given []int, expected []int) {
		double := func(el int) int { return (el * 2) }
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := gchannels.Map(c, double)

		// convert chan int to chan G
		c2 := make(chan int, 1)
		go func() {
			for el := range result {
				c2 <- el
			}
			close(c2)
		}()

		actual := gchannels.ToSlice(c2)
		assert.Equal(t, expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{2})
	f([]int{1, 2, 3}, []int{2, 4, 6})
}

func TestChannelMax(t *testing.T) {
	f := func(given []int, expected int, expectedErr error) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := gchannels.Max(c)
		assert.Equal(t, expected, actual)
		assert.Equal(t, expectedErr, actualErr)
	}
	f([]int{}, 0, gchannels.ErrEmpty)
	f([]int{1, 4, 2}, 4, nil)
	f([]int{1, 2, 4}, 4, nil)
	f([]int{4, 2, 1}, 4, nil)
}

func TestChannelMin(t *testing.T) {
	f := func(given []int, expected int, expectedErr error) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := gchannels.Min(c)
		assert.Equal(t, expected, actual)
		assert.Equal(t, expectedErr, actualErr)
	}
	f([]int{}, 0, gchannels.ErrEmpty)
	f([]int{4, 1, 2}, 1, nil)
	f([]int{1, 2, 4}, 1, nil)
	f([]int{4, 2, 1}, 1, nil)
}

func TestChannelReduce(t *testing.T) {
	f := func(given []int, expected int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int, acc int) int { return (el) + acc }
		actual := gchannels.Reduce(c, 0, sum)
		assert.Equal(t, expected, actual)
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3, 4, 5}, 15)
}

func TestChannelScan(t *testing.T) {
	f := func(given []int, expected []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int, acc int) int { return (el) + acc }
		result := gchannels.Scan(c, 0, sum)

		// convert chan int to chan G
		c2 := make(chan int, 1)
		go func() {
			for el := range result {
				c2 <- (el)
			}
			close(c2)
		}()

		actual := gchannels.ToSlice(c2)
		assert.Equal(t, expected, actual)
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 2}, []int{1, 3})
	f([]int{1, 2, 3, 4, 5}, []int{1, 3, 6, 10, 15})
}

func TestChannelSum(t *testing.T) {
	f := func(given []int, expected int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := gchannels.Sum(c)
		assert.Equal(t, expected, actual)
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3, 4, 5}, 15)
}

func TestChannelTake(t *testing.T) {
	f := func(count int, given int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		seq := gchannels.Repeat(ctx, given)
		seq2 := gchannels.Take(seq, count)
		actual := gchannels.ToSlice(seq2)
		cancel()
		assert.Equal(t, expected, actual)
	}
	f(0, 1, []int{})
	f(1, 1, []int{1})
	f(2, 1, []int{1, 1})
}

func TestChannelTee(t *testing.T) {
	f := func(count int, given []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		channels := gchannels.Tee(c, count)
		for _, ch := range channels {
			go func(ch chan int) {
				actual := gchannels.ToSlice(ch)
				assert.Equal(t, given, actual)
			}(ch)
		}
	}
	f(1, []int{})
	f(1, []int{1})
	f(1, []int{1, 2})
	f(1, []int{1, 2, 3})
	f(1, []int{1, 2, 3, 1, 2})

	f(2, []int{})
	f(2, []int{1})
	f(2, []int{1, 2})
	f(2, []int{1, 2, 3})
	f(2, []int{1, 2, 3, 1, 2})

	f(10, []int{1, 2, 3, 1, 2})
}
