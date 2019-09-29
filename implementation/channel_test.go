package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelToSlice(t *testing.T) {
	f := func(given []T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}
	f([]T{})
	f([]T{1})
	f([]T{1, 2, 3, 1, 2})
}

func TestChannelAny(t *testing.T) {
	f := func(given []T, expected bool) {
		even := func(t T) bool { return t%2 == 0 }
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, false)
	f([]T{1}, false)
	f([]T{2}, true)
	f([]T{1, 2}, true)
	f([]T{1, 2, 3}, true)
	f([]T{1, 3, 5}, false)
	f([]T{1, 3, 5, 7, 9, 11}, false)
	f([]T{1, 3, 5, 7, 10, 11}, true)
}

func TestChannelAll(t *testing.T) {
	f := func(given []T, expected bool) {
		even := func(t T) bool { return t%2 == 0 }
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, true)
	f([]T{1}, false)
	f([]T{2}, true)
	f([]T{1, 2}, false)
	f([]T{2, 4}, true)
	f([]T{2, 4, 6, 8, 10, 12}, true)
	f([]T{2, 4, 6, 8, 11, 12}, false)
}

func TestChannelEach(t *testing.T) {
	f := func(given []T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := make(chan T, len(given))
		mapper := func(t T) { result <- t }
		Channel{c}.Each(mapper)
		close(result)
		actual := Channel{result}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}

	f([]T{})
	f([]T{1})
	f([]T{1, 2, 3})
	f([]T{1, 2, 3, 4, 5, 6, 7})
}

func TestChannelChunkEvery(t *testing.T) {
	f := func(size int, given []T, expected [][]T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := Channel{c}.ChunkEvery(size)
		actual := make([][]T, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []T{}, [][]T{})
	f(2, []T{1}, [][]T{{1}})
	f(2, []T{1, 2}, [][]T{{1, 2}})
	f(2, []T{1, 2, 3, 4}, [][]T{{1, 2}, {3, 4}})
	f(2, []T{1, 2, 3, 4, 5}, [][]T{{1, 2}, {3, 4}, {5}})
}

func TestChannelCount(t *testing.T) {
	f := func(element T, given []T, expected int) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.Count(element)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []T{}, 0)
	f(1, []T{1}, 1)
	f(1, []T{2}, 0)
	f(1, []T{1, 2, 3, 1, 4}, 2)
}

func TestChannelDrop(t *testing.T) {
	f := func(count int, given []T, expected []T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := Channel{c}.Drop(count)
		actual := make([]T, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []T{}, []T{})
	f(1, []T{2}, []T{})
	f(1, []T{2, 3}, []T{3})
	f(1, []T{1, 2, 3}, []T{2, 3})
	f(0, []T{1, 2, 3}, []T{1, 2, 3})
	f(3, []T{1, 2, 3, 4, 5, 6}, []T{4, 5, 6})
	f(1, []T{1, 2, 3, 4, 5, 6}, []T{2, 3, 4, 5, 6})
}

func TestChannelFilter(t *testing.T) {
	f := func(given []T, expected []T) {
		even := func(t T) bool { return t%2 == 0 }
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := Channel{c}.Filter(even)
		actual := Channel{result}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{})
	f([]T{2}, []T{2})
	f([]T{1, 2, 3, 4}, []T{2, 4})
}

func TestChannelMap(t *testing.T) {
	f := func(given []T, expected []T) {
		double := func(el T) G { return G(el * 2) }
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := Channel{c}.Map(double)

		// convert chan T to chan G
		c2 := make(chan T, 1)
		go func() {
			for el := range result {
				c2 <- T(el)
			}
			close(c2)
		}()

		actual := Channel{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{2})
	f([]T{1, 2, 3}, []T{2, 4, 6})
}

func TestChannelMax(t *testing.T) {
	f := func(given []T, expected T, expectedErr error) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := Channel{c}.Max()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]T{}, 0, ErrEmpty)
	f([]T{1, 4, 2}, 4, nil)
	f([]T{1, 2, 4}, 4, nil)
	f([]T{4, 2, 1}, 4, nil)
}

func TestChannelMin(t *testing.T) {
	f := func(given []T, expected T, expectedErr error) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := Channel{c}.Min()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]T{}, 0, ErrEmpty)
	f([]T{4, 1, 2}, 1, nil)
	f([]T{1, 2, 4}, 1, nil)
	f([]T{4, 2, 1}, 1, nil)
}

func TestChannelReduce(t *testing.T) {
	f := func(given []T, expected G) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el T, acc G) G { return G(el) + acc }
		actual := Channel{c}.Reduce(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 0)
	f([]T{1}, 1)
	f([]T{1, 2}, 3)
	f([]T{1, 2, 3, 4, 5}, 15)
}

func TestChannelScan(t *testing.T) {
	f := func(given []T, expected []T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el T, acc G) G { return G(el) + acc }
		result := Channel{c}.Scan(0, sum)

		// convert chan T to chan G
		c2 := make(chan T, 1)
		go func() {
			for el := range result {
				c2 <- T(el)
			}
			close(c2)
		}()

		actual := Channel{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{1})
	f([]T{1, 2}, []T{1, 3})
	f([]T{1, 2, 3, 4, 5}, []T{1, 3, 6, 10, 15})
}

func TestChannelSum(t *testing.T) {
	f := func(given []T, expected T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := Channel{c}.Sum()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 0)
	f([]T{1}, 1)
	f([]T{1, 2}, 3)
	f([]T{1, 2, 3, 4, 5}, 15)
}

func TestChannelTake(t *testing.T) {
	s := Sequence{}
	f := func(count int, given T, expected []T) {
		seq := s.Repeat(given)
		seq2 := Channel{seq}.Take(count)
		actual := Channel{seq2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []T{})
	f(1, 1, []T{1})
	f(2, 1, []T{1, 1})
}

func TestChannelTee(t *testing.T) {
	f := func(count int, given []T) {
		c := make(chan T, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		channels := Channel{c}.Tee(count)
		for _, ch := range channels {
			go func(ch chan T) {
				actual := Channel{ch}.ToSlice()
				assert.Equal(t, given, actual, "they should be equal")
			}(ch)
		}
	}
	f(1, []T{})
	f(1, []T{1})
	f(1, []T{1, 2})
	f(1, []T{1, 2, 3})
	f(1, []T{1, 2, 3, 1, 2})

	f(2, []T{})
	f(2, []T{1})
	f(2, []T{1, 2})
	f(2, []T{1, 2, 3})
	f(2, []T{1, 2, 3, 1, 2})

	f(10, []T{1, 2, 3, 1, 2})
}
