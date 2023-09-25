package iters_test

import (
	"testing"

	"github.com/life4/genesis/iters"
	"github.com/matryer/is"
)

func assertAllocs(t *testing.T, expected uint64, f func()) {
	res := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		f()
	})
	is := is.New(t)
	is.Equal(res.MemAllocs, expected)
}

func TestFromSlice_Allocs(t *testing.T) {
	s := make([]int, 1000)
	assertAllocs(t, 2, func() {
		it := iters.FromSlice(s)
		it.Next()
		it.Next()
		it.Next()
	})
}
