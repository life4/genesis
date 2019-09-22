package genesis

import (
	"github.com/stretchr/testify/assert"
	"testing")

func TestAsyncSliceAnyInt(t *testing.T) {
	f := func(check func(t int) bool, given []int, expected bool) {
		s := AsyncSliceInt{data: given, workers: 2}
		actual := s.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int) bool { return (t % 2) == 0 }

	f(isEven, []int{}, false)
	f(isEven, []int{1}, false)
	f(isEven, []int{1, 3}, false)
	f(isEven, []int{2}, true)
	f(isEven, []int{1, 2}, true)
	f(isEven, []int{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []int{1, 3, 5, 7, 9, 12}, true)
}

func TestAsyncSliceAllInt(t *testing.T) {
	f := func(check func(t int) bool, given []int, expected bool) {
		s := AsyncSliceInt{data: given, workers: 2}
		actual := s.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int) bool { return (t % 2) == 0 }

	f(isEven, []int{}, true)
	f(isEven, []int{1}, false)
	f(isEven, []int{1, 3}, false)
	f(isEven, []int{2}, true)
	f(isEven, []int{2, 4}, true)
	f(isEven, []int{2, 3}, false)
	f(isEven, []int{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []int{2, 4, 6, 8, 10, 11}, false)
}

func TestAsyncSliceEachInt(t *testing.T) {
	f := func(given []int) {
		s := AsyncSliceInt{data: given, workers: 2}
		result := make(chan int, len(given))
		mapper := func(t int) { result <- t }
		s.Each(mapper)
		close(result)
		actual := ChannelInt{result}.ToSlice()
		sorted := SliceInt{actual}.Sort()
		assert.Equal(t, given, sorted, "they should be equal")
	}

	f([]int{})
	f([]int{1})
	f([]int{1, 2, 3})
	f([]int{1, 2, 3, 4, 5, 6, 7})
}

func TestAsyncSliceFilterInt(t *testing.T) {
	f := func(given []int, expected []int) {
		filter := func(t int) bool { return t > 10 }
		s := AsyncSliceInt{data: given, workers: 2}
		actual := s.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int{}, []int{})
	f([]int{5}, []int{})
	f([]int{15}, []int{15})
	f([]int{9, 11, 12, 13, 6}, []int{11, 12, 13})
}

func TestAsyncSliceMapIntInt(t *testing.T) {
	f := func(mapper func(t int) int, given []int, expected []int) {
		s := AsyncSliceInt{data: given, workers: 2}
		actual := s.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int { return int((t * 2)) }

	f(double, []int{}, []int{})
	f(double, []int{1}, []int{2})
	f(double, []int{1, 2, 3}, []int{2, 4, 6})
}

func TestAsyncSliceMapIntInt8(t *testing.T) {
	f := func(mapper func(t int) int8, given []int, expected []int8) {
		s := AsyncSliceInt{data: given, workers: 2}
		actual := s.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int8 { return int8((t * 2)) }

	f(double, []int{}, []int8{})
	f(double, []int{1}, []int8{2})
	f(double, []int{1, 2, 3}, []int8{2, 4, 6})
}

func TestAsyncSliceMapIntInt16(t *testing.T) {
	f := func(mapper func(t int) int16, given []int, expected []int16) {
		s := AsyncSliceInt{data: given, workers: 2}
		actual := s.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int16 { return int16((t * 2)) }

	f(double, []int{}, []int16{})
	f(double, []int{1}, []int16{2})
	f(double, []int{1, 2, 3}, []int16{2, 4, 6})
}

func TestAsyncSliceMapIntInt32(t *testing.T) {
	f := func(mapper func(t int) int32, given []int, expected []int32) {
		s := AsyncSliceInt{data: given, workers: 2}
		actual := s.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int32 { return int32((t * 2)) }

	f(double, []int{}, []int32{})
	f(double, []int{1}, []int32{2})
	f(double, []int{1, 2, 3}, []int32{2, 4, 6})
}

func TestAsyncSliceMapIntInt64(t *testing.T) {
	f := func(mapper func(t int) int64, given []int, expected []int64) {
		s := AsyncSliceInt{data: given, workers: 2}
		actual := s.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int64 { return int64((t * 2)) }

	f(double, []int{}, []int64{})
	f(double, []int{1}, []int64{2})
	f(double, []int{1, 2, 3}, []int64{2, 4, 6})
}

func TestChannelToSliceInt(t *testing.T) {
	s := SequenceInt{}
	f := func(start int, stop int, step int, expected []int) {
		seq := s.Range(start, stop, step)
		actual := ChannelInt{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int{1, 2, 3})
}

func TestSequenceExponentialInt(t *testing.T) {
	s := SequenceInt{}
	f := func(start int, factor int, count int, expected []int) {
		seq := s.Exponential(start, factor)
		actual := ChannelInt{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []int{1, 2, 4, 8})
}

func TestSequenceRangeInt(t *testing.T) {
	s := SequenceInt{}
	f := func(start int, stop int, step int, expected []int) {
		seq := s.Range(start, stop, step)
		actual := ChannelInt{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int{1, 2, 3})
	f(3, 0, -1, []int{3, 2, 1})
}

func TestSequenceRepeatInt(t *testing.T) {
	s := SequenceInt{}
	f := func(count int, given int, expected []int) {
		seq := s.Repeat(given)
		actual := ChannelInt{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []int{1, 1})
}

func TestSequenceTakeInt(t *testing.T) {
	s := SequenceInt{}
	f := func(count int, given int, expected []int) {
		seq := s.Repeat(given)
		actual := ChannelInt{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int{})
	f(1, 1, []int{1})
	f(2, 1, []int{1, 1})
}

func TestSliceAnyInt(t *testing.T) {
	f := func(check func(t int) bool, given []int, expected bool) {
		actual := SliceInt{given}.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int) bool { return (t % 2) == 0 }

	f(isEven, []int{}, false)
	f(isEven, []int{1, 3}, false)
	f(isEven, []int{2}, true)
	f(isEven, []int{1, 2}, true)
}

func TestSliceAllInt(t *testing.T) {
	f := func(check func(t int) bool, given []int, expected bool) {
		actual := SliceInt{given}.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int) bool { return (t % 2) == 0 }

	f(isEven, []int{}, true)
	f(isEven, []int{2}, true)
	f(isEven, []int{1}, false)
	f(isEven, []int{2, 4}, true)
	f(isEven, []int{2, 4, 1}, false)
	f(isEven, []int{1, 2, 4}, false)
}

func TestSliceChunkByIntInt(t *testing.T) {
	f := func(mapper func(t int) int, given []int, expected [][]int) {
		actual := SliceInt{given}.ChunkByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int { return int((t % 2)) }

	f(remainder, []int{1}, [][]int{{1}})
	f(remainder, []int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f(remainder, []int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByIntInt8(t *testing.T) {
	f := func(mapper func(t int) int8, given []int, expected [][]int) {
		actual := SliceInt{given}.ChunkByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int8 { return int8((t % 2)) }

	f(remainder, []int{1}, [][]int{{1}})
	f(remainder, []int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f(remainder, []int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByIntInt16(t *testing.T) {
	f := func(mapper func(t int) int16, given []int, expected [][]int) {
		actual := SliceInt{given}.ChunkByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int16 { return int16((t % 2)) }

	f(remainder, []int{1}, [][]int{{1}})
	f(remainder, []int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f(remainder, []int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByIntInt32(t *testing.T) {
	f := func(mapper func(t int) int32, given []int, expected [][]int) {
		actual := SliceInt{given}.ChunkByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int32 { return int32((t % 2)) }

	f(remainder, []int{1}, [][]int{{1}})
	f(remainder, []int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f(remainder, []int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByIntInt64(t *testing.T) {
	f := func(mapper func(t int) int64, given []int, expected [][]int) {
		actual := SliceInt{given}.ChunkByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int64 { return int64((t % 2)) }

	f(remainder, []int{1}, [][]int{{1}})
	f(remainder, []int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f(remainder, []int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEveryInt(t *testing.T) {
	f := func(count int, given []int, expected [][]int) {
		actual := SliceInt{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int{1, 2, 3, 4}, [][]int{[]int{1, 2}, []int{3, 4}})
	f(2, []int{1, 2, 3, 4, 5}, [][]int{[]int{1, 2}, []int{3, 4}, []int{5}})
}

func TestSliceFilterInt(t *testing.T) {
	f := func(filter func(t int) bool, given []int, expected []int) {
		actual := SliceInt{given}.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t int) bool { return t > 0 }

	f(filterPositive, []int{1, -1, 2, -2, 3, -3}, []int{1, 2, 3})
	f(filterPositive, []int{1, 2, 3}, []int{1, 2, 3})
	f(filterPositive, []int{-1, -2, -3}, []int{})
}

func TestSliceGroupByIntInt(t *testing.T) {
	f := func(mapper func(t int) int, given []int, expected map[int][]int) {
		actual := SliceInt{given}.GroupByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int { return int((t % 2)) }

	f(remainder, []int{}, map[int][]int{})
	f(remainder, []int{1}, map[int][]int{1: {1}})
	f(remainder, []int{1, 3, 2, 4, 5}, map[int][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByIntInt8(t *testing.T) {
	f := func(mapper func(t int) int8, given []int, expected map[int8][]int) {
		actual := SliceInt{given}.GroupByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int8 { return int8((t % 2)) }

	f(remainder, []int{}, map[int8][]int{})
	f(remainder, []int{1}, map[int8][]int{1: {1}})
	f(remainder, []int{1, 3, 2, 4, 5}, map[int8][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByIntInt16(t *testing.T) {
	f := func(mapper func(t int) int16, given []int, expected map[int16][]int) {
		actual := SliceInt{given}.GroupByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int16 { return int16((t % 2)) }

	f(remainder, []int{}, map[int16][]int{})
	f(remainder, []int{1}, map[int16][]int{1: {1}})
	f(remainder, []int{1, 3, 2, 4, 5}, map[int16][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByIntInt32(t *testing.T) {
	f := func(mapper func(t int) int32, given []int, expected map[int32][]int) {
		actual := SliceInt{given}.GroupByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int32 { return int32((t % 2)) }

	f(remainder, []int{}, map[int32][]int{})
	f(remainder, []int{1}, map[int32][]int{1: {1}})
	f(remainder, []int{1, 3, 2, 4, 5}, map[int32][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByIntInt64(t *testing.T) {
	f := func(mapper func(t int) int64, given []int, expected map[int64][]int) {
		actual := SliceInt{given}.GroupByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int64 { return int64((t % 2)) }

	f(remainder, []int{}, map[int64][]int{})
	f(remainder, []int{1}, map[int64][]int{1: {1}})
	f(remainder, []int{1, 3, 2, 4, 5}, map[int64][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceIntersperseInt(t *testing.T) {
	f := func(el int, given []int, expected []int) {
		actual := SliceInt{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int{1, 2, 3}, []int{1, 0, 2, 0, 3})
}

func TestSliceMapIntInt(t *testing.T) {
	f := func(mapper func(t int) int, given []int, expected []int) {
		actual := SliceInt{given}.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int { return int((t * 2)) }

	f(double, []int{1, 2, 3}, []int{2, 4, 6})
}

func TestSliceMapIntInt8(t *testing.T) {
	f := func(mapper func(t int) int8, given []int, expected []int8) {
		actual := SliceInt{given}.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int8 { return int8((t * 2)) }

	f(double, []int{1, 2, 3}, []int8{2, 4, 6})
}

func TestSliceMapIntInt16(t *testing.T) {
	f := func(mapper func(t int) int16, given []int, expected []int16) {
		actual := SliceInt{given}.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int16 { return int16((t * 2)) }

	f(double, []int{1, 2, 3}, []int16{2, 4, 6})
}

func TestSliceMapIntInt32(t *testing.T) {
	f := func(mapper func(t int) int32, given []int, expected []int32) {
		actual := SliceInt{given}.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int32 { return int32((t * 2)) }

	f(double, []int{1, 2, 3}, []int32{2, 4, 6})
}

func TestSliceMapIntInt64(t *testing.T) {
	f := func(mapper func(t int) int64, given []int, expected []int64) {
		actual := SliceInt{given}.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int64 { return int64((t * 2)) }

	f(double, []int{1, 2, 3}, []int64{2, 4, 6})
}

func TestSlicesProductInt(t *testing.T) {
	f := func(given [][]int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		s := SlicesInt{given}
		for el := range s.Product() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]int{{1, 2}, {3}, {4, 5}}, [][]int{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}

func TestAsyncSliceAnyInt8(t *testing.T) {
	f := func(check func(t int8) bool, given []int8, expected bool) {
		s := AsyncSliceInt8{data: given, workers: 2}
		actual := s.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int8) bool { return (t % 2) == 0 }

	f(isEven, []int8{}, false)
	f(isEven, []int8{1}, false)
	f(isEven, []int8{1, 3}, false)
	f(isEven, []int8{2}, true)
	f(isEven, []int8{1, 2}, true)
	f(isEven, []int8{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []int8{1, 3, 5, 7, 9, 12}, true)
}

func TestAsyncSliceAllInt8(t *testing.T) {
	f := func(check func(t int8) bool, given []int8, expected bool) {
		s := AsyncSliceInt8{data: given, workers: 2}
		actual := s.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int8) bool { return (t % 2) == 0 }

	f(isEven, []int8{}, true)
	f(isEven, []int8{1}, false)
	f(isEven, []int8{1, 3}, false)
	f(isEven, []int8{2}, true)
	f(isEven, []int8{2, 4}, true)
	f(isEven, []int8{2, 3}, false)
	f(isEven, []int8{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []int8{2, 4, 6, 8, 10, 11}, false)
}

func TestAsyncSliceEachInt8(t *testing.T) {
	f := func(given []int8) {
		s := AsyncSliceInt8{data: given, workers: 2}
		result := make(chan int8, len(given))
		mapper := func(t int8) { result <- t }
		s.Each(mapper)
		close(result)
		actual := ChannelInt8{result}.ToSlice()
		sorted := SliceInt8{actual}.Sort()
		assert.Equal(t, given, sorted, "they should be equal")
	}

	f([]int8{})
	f([]int8{1})
	f([]int8{1, 2, 3})
	f([]int8{1, 2, 3, 4, 5, 6, 7})
}

func TestAsyncSliceFilterInt8(t *testing.T) {
	f := func(given []int8, expected []int8) {
		filter := func(t int8) bool { return t > 10 }
		s := AsyncSliceInt8{data: given, workers: 2}
		actual := s.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int8{}, []int8{})
	f([]int8{5}, []int8{})
	f([]int8{15}, []int8{15})
	f([]int8{9, 11, 12, 13, 6}, []int8{11, 12, 13})
}

func TestAsyncSliceMapInt8Int(t *testing.T) {
	f := func(mapper func(t int8) int, given []int8, expected []int) {
		s := AsyncSliceInt8{data: given, workers: 2}
		actual := s.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int { return int((t * 2)) }

	f(double, []int8{}, []int{})
	f(double, []int8{1}, []int{2})
	f(double, []int8{1, 2, 3}, []int{2, 4, 6})
}

func TestAsyncSliceMapInt8Int8(t *testing.T) {
	f := func(mapper func(t int8) int8, given []int8, expected []int8) {
		s := AsyncSliceInt8{data: given, workers: 2}
		actual := s.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int8 { return int8((t * 2)) }

	f(double, []int8{}, []int8{})
	f(double, []int8{1}, []int8{2})
	f(double, []int8{1, 2, 3}, []int8{2, 4, 6})
}

func TestAsyncSliceMapInt8Int16(t *testing.T) {
	f := func(mapper func(t int8) int16, given []int8, expected []int16) {
		s := AsyncSliceInt8{data: given, workers: 2}
		actual := s.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int16 { return int16((t * 2)) }

	f(double, []int8{}, []int16{})
	f(double, []int8{1}, []int16{2})
	f(double, []int8{1, 2, 3}, []int16{2, 4, 6})
}

func TestAsyncSliceMapInt8Int32(t *testing.T) {
	f := func(mapper func(t int8) int32, given []int8, expected []int32) {
		s := AsyncSliceInt8{data: given, workers: 2}
		actual := s.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int32 { return int32((t * 2)) }

	f(double, []int8{}, []int32{})
	f(double, []int8{1}, []int32{2})
	f(double, []int8{1, 2, 3}, []int32{2, 4, 6})
}

func TestAsyncSliceMapInt8Int64(t *testing.T) {
	f := func(mapper func(t int8) int64, given []int8, expected []int64) {
		s := AsyncSliceInt8{data: given, workers: 2}
		actual := s.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int64 { return int64((t * 2)) }

	f(double, []int8{}, []int64{})
	f(double, []int8{1}, []int64{2})
	f(double, []int8{1, 2, 3}, []int64{2, 4, 6})
}

func TestChannelToSliceInt8(t *testing.T) {
	s := SequenceInt8{}
	f := func(start int8, stop int8, step int8, expected []int8) {
		seq := s.Range(start, stop, step)
		actual := ChannelInt8{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int8{1, 2, 3})
}

func TestSequenceExponentialInt8(t *testing.T) {
	s := SequenceInt8{}
	f := func(start int8, factor int8, count int, expected []int8) {
		seq := s.Exponential(start, factor)
		actual := ChannelInt8{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []int8{1, 2, 4, 8})
}

func TestSequenceRangeInt8(t *testing.T) {
	s := SequenceInt8{}
	f := func(start int8, stop int8, step int8, expected []int8) {
		seq := s.Range(start, stop, step)
		actual := ChannelInt8{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int8{1, 2, 3})
	f(3, 0, -1, []int8{3, 2, 1})
}

func TestSequenceRepeatInt8(t *testing.T) {
	s := SequenceInt8{}
	f := func(count int, given int8, expected []int8) {
		seq := s.Repeat(given)
		actual := ChannelInt8{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []int8{1, 1})
}

func TestSequenceTakeInt8(t *testing.T) {
	s := SequenceInt8{}
	f := func(count int, given int8, expected []int8) {
		seq := s.Repeat(given)
		actual := ChannelInt8{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int8{})
	f(1, 1, []int8{1})
	f(2, 1, []int8{1, 1})
}

func TestSliceAnyInt8(t *testing.T) {
	f := func(check func(t int8) bool, given []int8, expected bool) {
		actual := SliceInt8{given}.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int8) bool { return (t % 2) == 0 }

	f(isEven, []int8{}, false)
	f(isEven, []int8{1, 3}, false)
	f(isEven, []int8{2}, true)
	f(isEven, []int8{1, 2}, true)
}

func TestSliceAllInt8(t *testing.T) {
	f := func(check func(t int8) bool, given []int8, expected bool) {
		actual := SliceInt8{given}.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int8) bool { return (t % 2) == 0 }

	f(isEven, []int8{}, true)
	f(isEven, []int8{2}, true)
	f(isEven, []int8{1}, false)
	f(isEven, []int8{2, 4}, true)
	f(isEven, []int8{2, 4, 1}, false)
	f(isEven, []int8{1, 2, 4}, false)
}

func TestSliceChunkByInt8Int(t *testing.T) {
	f := func(mapper func(t int8) int, given []int8, expected [][]int8) {
		actual := SliceInt8{given}.ChunkByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int { return int((t % 2)) }

	f(remainder, []int8{1}, [][]int8{{1}})
	f(remainder, []int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f(remainder, []int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt8Int8(t *testing.T) {
	f := func(mapper func(t int8) int8, given []int8, expected [][]int8) {
		actual := SliceInt8{given}.ChunkByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int8 { return int8((t % 2)) }

	f(remainder, []int8{1}, [][]int8{{1}})
	f(remainder, []int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f(remainder, []int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt8Int16(t *testing.T) {
	f := func(mapper func(t int8) int16, given []int8, expected [][]int8) {
		actual := SliceInt8{given}.ChunkByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int16 { return int16((t % 2)) }

	f(remainder, []int8{1}, [][]int8{{1}})
	f(remainder, []int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f(remainder, []int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt8Int32(t *testing.T) {
	f := func(mapper func(t int8) int32, given []int8, expected [][]int8) {
		actual := SliceInt8{given}.ChunkByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int32 { return int32((t % 2)) }

	f(remainder, []int8{1}, [][]int8{{1}})
	f(remainder, []int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f(remainder, []int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt8Int64(t *testing.T) {
	f := func(mapper func(t int8) int64, given []int8, expected [][]int8) {
		actual := SliceInt8{given}.ChunkByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int64 { return int64((t % 2)) }

	f(remainder, []int8{1}, [][]int8{{1}})
	f(remainder, []int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f(remainder, []int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEveryInt8(t *testing.T) {
	f := func(count int, given []int8, expected [][]int8) {
		actual := SliceInt8{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int8{1, 2, 3, 4}, [][]int8{[]int8{1, 2}, []int8{3, 4}})
	f(2, []int8{1, 2, 3, 4, 5}, [][]int8{[]int8{1, 2}, []int8{3, 4}, []int8{5}})
}

func TestSliceFilterInt8(t *testing.T) {
	f := func(filter func(t int8) bool, given []int8, expected []int8) {
		actual := SliceInt8{given}.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t int8) bool { return t > 0 }

	f(filterPositive, []int8{1, -1, 2, -2, 3, -3}, []int8{1, 2, 3})
	f(filterPositive, []int8{1, 2, 3}, []int8{1, 2, 3})
	f(filterPositive, []int8{-1, -2, -3}, []int8{})
}

func TestSliceGroupByInt8Int(t *testing.T) {
	f := func(mapper func(t int8) int, given []int8, expected map[int][]int8) {
		actual := SliceInt8{given}.GroupByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int { return int((t % 2)) }

	f(remainder, []int8{}, map[int][]int8{})
	f(remainder, []int8{1}, map[int][]int8{1: {1}})
	f(remainder, []int8{1, 3, 2, 4, 5}, map[int][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt8Int8(t *testing.T) {
	f := func(mapper func(t int8) int8, given []int8, expected map[int8][]int8) {
		actual := SliceInt8{given}.GroupByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int8 { return int8((t % 2)) }

	f(remainder, []int8{}, map[int8][]int8{})
	f(remainder, []int8{1}, map[int8][]int8{1: {1}})
	f(remainder, []int8{1, 3, 2, 4, 5}, map[int8][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt8Int16(t *testing.T) {
	f := func(mapper func(t int8) int16, given []int8, expected map[int16][]int8) {
		actual := SliceInt8{given}.GroupByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int16 { return int16((t % 2)) }

	f(remainder, []int8{}, map[int16][]int8{})
	f(remainder, []int8{1}, map[int16][]int8{1: {1}})
	f(remainder, []int8{1, 3, 2, 4, 5}, map[int16][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt8Int32(t *testing.T) {
	f := func(mapper func(t int8) int32, given []int8, expected map[int32][]int8) {
		actual := SliceInt8{given}.GroupByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int32 { return int32((t % 2)) }

	f(remainder, []int8{}, map[int32][]int8{})
	f(remainder, []int8{1}, map[int32][]int8{1: {1}})
	f(remainder, []int8{1, 3, 2, 4, 5}, map[int32][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt8Int64(t *testing.T) {
	f := func(mapper func(t int8) int64, given []int8, expected map[int64][]int8) {
		actual := SliceInt8{given}.GroupByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int64 { return int64((t % 2)) }

	f(remainder, []int8{}, map[int64][]int8{})
	f(remainder, []int8{1}, map[int64][]int8{1: {1}})
	f(remainder, []int8{1, 3, 2, 4, 5}, map[int64][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceIntersperseInt8(t *testing.T) {
	f := func(el int8, given []int8, expected []int8) {
		actual := SliceInt8{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int8{1, 2, 3}, []int8{1, 0, 2, 0, 3})
}

func TestSliceMapInt8Int(t *testing.T) {
	f := func(mapper func(t int8) int, given []int8, expected []int) {
		actual := SliceInt8{given}.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int { return int((t * 2)) }

	f(double, []int8{1, 2, 3}, []int{2, 4, 6})
}

func TestSliceMapInt8Int8(t *testing.T) {
	f := func(mapper func(t int8) int8, given []int8, expected []int8) {
		actual := SliceInt8{given}.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int8 { return int8((t * 2)) }

	f(double, []int8{1, 2, 3}, []int8{2, 4, 6})
}

func TestSliceMapInt8Int16(t *testing.T) {
	f := func(mapper func(t int8) int16, given []int8, expected []int16) {
		actual := SliceInt8{given}.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int16 { return int16((t * 2)) }

	f(double, []int8{1, 2, 3}, []int16{2, 4, 6})
}

func TestSliceMapInt8Int32(t *testing.T) {
	f := func(mapper func(t int8) int32, given []int8, expected []int32) {
		actual := SliceInt8{given}.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int32 { return int32((t * 2)) }

	f(double, []int8{1, 2, 3}, []int32{2, 4, 6})
}

func TestSliceMapInt8Int64(t *testing.T) {
	f := func(mapper func(t int8) int64, given []int8, expected []int64) {
		actual := SliceInt8{given}.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int64 { return int64((t * 2)) }

	f(double, []int8{1, 2, 3}, []int64{2, 4, 6})
}

func TestSlicesProductInt8(t *testing.T) {
	f := func(given [][]int8, expected [][]int8) {
		actual := make([][]int8, 0)
		i := 0
		s := SlicesInt8{given}
		for el := range s.Product() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]int8{{1, 2}, {3, 4}}, [][]int8{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]int8{{1, 2}, {3}, {4, 5}}, [][]int8{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}

func TestAsyncSliceAnyInt16(t *testing.T) {
	f := func(check func(t int16) bool, given []int16, expected bool) {
		s := AsyncSliceInt16{data: given, workers: 2}
		actual := s.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int16) bool { return (t % 2) == 0 }

	f(isEven, []int16{}, false)
	f(isEven, []int16{1}, false)
	f(isEven, []int16{1, 3}, false)
	f(isEven, []int16{2}, true)
	f(isEven, []int16{1, 2}, true)
	f(isEven, []int16{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []int16{1, 3, 5, 7, 9, 12}, true)
}

func TestAsyncSliceAllInt16(t *testing.T) {
	f := func(check func(t int16) bool, given []int16, expected bool) {
		s := AsyncSliceInt16{data: given, workers: 2}
		actual := s.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int16) bool { return (t % 2) == 0 }

	f(isEven, []int16{}, true)
	f(isEven, []int16{1}, false)
	f(isEven, []int16{1, 3}, false)
	f(isEven, []int16{2}, true)
	f(isEven, []int16{2, 4}, true)
	f(isEven, []int16{2, 3}, false)
	f(isEven, []int16{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []int16{2, 4, 6, 8, 10, 11}, false)
}

func TestAsyncSliceEachInt16(t *testing.T) {
	f := func(given []int16) {
		s := AsyncSliceInt16{data: given, workers: 2}
		result := make(chan int16, len(given))
		mapper := func(t int16) { result <- t }
		s.Each(mapper)
		close(result)
		actual := ChannelInt16{result}.ToSlice()
		sorted := SliceInt16{actual}.Sort()
		assert.Equal(t, given, sorted, "they should be equal")
	}

	f([]int16{})
	f([]int16{1})
	f([]int16{1, 2, 3})
	f([]int16{1, 2, 3, 4, 5, 6, 7})
}

func TestAsyncSliceFilterInt16(t *testing.T) {
	f := func(given []int16, expected []int16) {
		filter := func(t int16) bool { return t > 10 }
		s := AsyncSliceInt16{data: given, workers: 2}
		actual := s.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int16{}, []int16{})
	f([]int16{5}, []int16{})
	f([]int16{15}, []int16{15})
	f([]int16{9, 11, 12, 13, 6}, []int16{11, 12, 13})
}

func TestAsyncSliceMapInt16Int(t *testing.T) {
	f := func(mapper func(t int16) int, given []int16, expected []int) {
		s := AsyncSliceInt16{data: given, workers: 2}
		actual := s.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int { return int((t * 2)) }

	f(double, []int16{}, []int{})
	f(double, []int16{1}, []int{2})
	f(double, []int16{1, 2, 3}, []int{2, 4, 6})
}

func TestAsyncSliceMapInt16Int8(t *testing.T) {
	f := func(mapper func(t int16) int8, given []int16, expected []int8) {
		s := AsyncSliceInt16{data: given, workers: 2}
		actual := s.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int8 { return int8((t * 2)) }

	f(double, []int16{}, []int8{})
	f(double, []int16{1}, []int8{2})
	f(double, []int16{1, 2, 3}, []int8{2, 4, 6})
}

func TestAsyncSliceMapInt16Int16(t *testing.T) {
	f := func(mapper func(t int16) int16, given []int16, expected []int16) {
		s := AsyncSliceInt16{data: given, workers: 2}
		actual := s.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int16 { return int16((t * 2)) }

	f(double, []int16{}, []int16{})
	f(double, []int16{1}, []int16{2})
	f(double, []int16{1, 2, 3}, []int16{2, 4, 6})
}

func TestAsyncSliceMapInt16Int32(t *testing.T) {
	f := func(mapper func(t int16) int32, given []int16, expected []int32) {
		s := AsyncSliceInt16{data: given, workers: 2}
		actual := s.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int32 { return int32((t * 2)) }

	f(double, []int16{}, []int32{})
	f(double, []int16{1}, []int32{2})
	f(double, []int16{1, 2, 3}, []int32{2, 4, 6})
}

func TestAsyncSliceMapInt16Int64(t *testing.T) {
	f := func(mapper func(t int16) int64, given []int16, expected []int64) {
		s := AsyncSliceInt16{data: given, workers: 2}
		actual := s.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int64 { return int64((t * 2)) }

	f(double, []int16{}, []int64{})
	f(double, []int16{1}, []int64{2})
	f(double, []int16{1, 2, 3}, []int64{2, 4, 6})
}

func TestChannelToSliceInt16(t *testing.T) {
	s := SequenceInt16{}
	f := func(start int16, stop int16, step int16, expected []int16) {
		seq := s.Range(start, stop, step)
		actual := ChannelInt16{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int16{1, 2, 3})
}

func TestSequenceExponentialInt16(t *testing.T) {
	s := SequenceInt16{}
	f := func(start int16, factor int16, count int, expected []int16) {
		seq := s.Exponential(start, factor)
		actual := ChannelInt16{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []int16{1, 2, 4, 8})
}

func TestSequenceRangeInt16(t *testing.T) {
	s := SequenceInt16{}
	f := func(start int16, stop int16, step int16, expected []int16) {
		seq := s.Range(start, stop, step)
		actual := ChannelInt16{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int16{1, 2, 3})
	f(3, 0, -1, []int16{3, 2, 1})
}

func TestSequenceRepeatInt16(t *testing.T) {
	s := SequenceInt16{}
	f := func(count int, given int16, expected []int16) {
		seq := s.Repeat(given)
		actual := ChannelInt16{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []int16{1, 1})
}

func TestSequenceTakeInt16(t *testing.T) {
	s := SequenceInt16{}
	f := func(count int, given int16, expected []int16) {
		seq := s.Repeat(given)
		actual := ChannelInt16{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int16{})
	f(1, 1, []int16{1})
	f(2, 1, []int16{1, 1})
}

func TestSliceAnyInt16(t *testing.T) {
	f := func(check func(t int16) bool, given []int16, expected bool) {
		actual := SliceInt16{given}.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int16) bool { return (t % 2) == 0 }

	f(isEven, []int16{}, false)
	f(isEven, []int16{1, 3}, false)
	f(isEven, []int16{2}, true)
	f(isEven, []int16{1, 2}, true)
}

func TestSliceAllInt16(t *testing.T) {
	f := func(check func(t int16) bool, given []int16, expected bool) {
		actual := SliceInt16{given}.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int16) bool { return (t % 2) == 0 }

	f(isEven, []int16{}, true)
	f(isEven, []int16{2}, true)
	f(isEven, []int16{1}, false)
	f(isEven, []int16{2, 4}, true)
	f(isEven, []int16{2, 4, 1}, false)
	f(isEven, []int16{1, 2, 4}, false)
}

func TestSliceChunkByInt16Int(t *testing.T) {
	f := func(mapper func(t int16) int, given []int16, expected [][]int16) {
		actual := SliceInt16{given}.ChunkByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int { return int((t % 2)) }

	f(remainder, []int16{1}, [][]int16{{1}})
	f(remainder, []int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f(remainder, []int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt16Int8(t *testing.T) {
	f := func(mapper func(t int16) int8, given []int16, expected [][]int16) {
		actual := SliceInt16{given}.ChunkByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int8 { return int8((t % 2)) }

	f(remainder, []int16{1}, [][]int16{{1}})
	f(remainder, []int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f(remainder, []int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt16Int16(t *testing.T) {
	f := func(mapper func(t int16) int16, given []int16, expected [][]int16) {
		actual := SliceInt16{given}.ChunkByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int16 { return int16((t % 2)) }

	f(remainder, []int16{1}, [][]int16{{1}})
	f(remainder, []int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f(remainder, []int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt16Int32(t *testing.T) {
	f := func(mapper func(t int16) int32, given []int16, expected [][]int16) {
		actual := SliceInt16{given}.ChunkByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int32 { return int32((t % 2)) }

	f(remainder, []int16{1}, [][]int16{{1}})
	f(remainder, []int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f(remainder, []int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt16Int64(t *testing.T) {
	f := func(mapper func(t int16) int64, given []int16, expected [][]int16) {
		actual := SliceInt16{given}.ChunkByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int64 { return int64((t % 2)) }

	f(remainder, []int16{1}, [][]int16{{1}})
	f(remainder, []int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f(remainder, []int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEveryInt16(t *testing.T) {
	f := func(count int, given []int16, expected [][]int16) {
		actual := SliceInt16{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int16{1, 2, 3, 4}, [][]int16{[]int16{1, 2}, []int16{3, 4}})
	f(2, []int16{1, 2, 3, 4, 5}, [][]int16{[]int16{1, 2}, []int16{3, 4}, []int16{5}})
}

func TestSliceFilterInt16(t *testing.T) {
	f := func(filter func(t int16) bool, given []int16, expected []int16) {
		actual := SliceInt16{given}.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t int16) bool { return t > 0 }

	f(filterPositive, []int16{1, -1, 2, -2, 3, -3}, []int16{1, 2, 3})
	f(filterPositive, []int16{1, 2, 3}, []int16{1, 2, 3})
	f(filterPositive, []int16{-1, -2, -3}, []int16{})
}

func TestSliceGroupByInt16Int(t *testing.T) {
	f := func(mapper func(t int16) int, given []int16, expected map[int][]int16) {
		actual := SliceInt16{given}.GroupByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int { return int((t % 2)) }

	f(remainder, []int16{}, map[int][]int16{})
	f(remainder, []int16{1}, map[int][]int16{1: {1}})
	f(remainder, []int16{1, 3, 2, 4, 5}, map[int][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt16Int8(t *testing.T) {
	f := func(mapper func(t int16) int8, given []int16, expected map[int8][]int16) {
		actual := SliceInt16{given}.GroupByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int8 { return int8((t % 2)) }

	f(remainder, []int16{}, map[int8][]int16{})
	f(remainder, []int16{1}, map[int8][]int16{1: {1}})
	f(remainder, []int16{1, 3, 2, 4, 5}, map[int8][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt16Int16(t *testing.T) {
	f := func(mapper func(t int16) int16, given []int16, expected map[int16][]int16) {
		actual := SliceInt16{given}.GroupByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int16 { return int16((t % 2)) }

	f(remainder, []int16{}, map[int16][]int16{})
	f(remainder, []int16{1}, map[int16][]int16{1: {1}})
	f(remainder, []int16{1, 3, 2, 4, 5}, map[int16][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt16Int32(t *testing.T) {
	f := func(mapper func(t int16) int32, given []int16, expected map[int32][]int16) {
		actual := SliceInt16{given}.GroupByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int32 { return int32((t % 2)) }

	f(remainder, []int16{}, map[int32][]int16{})
	f(remainder, []int16{1}, map[int32][]int16{1: {1}})
	f(remainder, []int16{1, 3, 2, 4, 5}, map[int32][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt16Int64(t *testing.T) {
	f := func(mapper func(t int16) int64, given []int16, expected map[int64][]int16) {
		actual := SliceInt16{given}.GroupByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int64 { return int64((t % 2)) }

	f(remainder, []int16{}, map[int64][]int16{})
	f(remainder, []int16{1}, map[int64][]int16{1: {1}})
	f(remainder, []int16{1, 3, 2, 4, 5}, map[int64][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceIntersperseInt16(t *testing.T) {
	f := func(el int16, given []int16, expected []int16) {
		actual := SliceInt16{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int16{1, 2, 3}, []int16{1, 0, 2, 0, 3})
}

func TestSliceMapInt16Int(t *testing.T) {
	f := func(mapper func(t int16) int, given []int16, expected []int) {
		actual := SliceInt16{given}.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int { return int((t * 2)) }

	f(double, []int16{1, 2, 3}, []int{2, 4, 6})
}

func TestSliceMapInt16Int8(t *testing.T) {
	f := func(mapper func(t int16) int8, given []int16, expected []int8) {
		actual := SliceInt16{given}.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int8 { return int8((t * 2)) }

	f(double, []int16{1, 2, 3}, []int8{2, 4, 6})
}

func TestSliceMapInt16Int16(t *testing.T) {
	f := func(mapper func(t int16) int16, given []int16, expected []int16) {
		actual := SliceInt16{given}.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int16 { return int16((t * 2)) }

	f(double, []int16{1, 2, 3}, []int16{2, 4, 6})
}

func TestSliceMapInt16Int32(t *testing.T) {
	f := func(mapper func(t int16) int32, given []int16, expected []int32) {
		actual := SliceInt16{given}.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int32 { return int32((t * 2)) }

	f(double, []int16{1, 2, 3}, []int32{2, 4, 6})
}

func TestSliceMapInt16Int64(t *testing.T) {
	f := func(mapper func(t int16) int64, given []int16, expected []int64) {
		actual := SliceInt16{given}.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int64 { return int64((t * 2)) }

	f(double, []int16{1, 2, 3}, []int64{2, 4, 6})
}

func TestSlicesProductInt16(t *testing.T) {
	f := func(given [][]int16, expected [][]int16) {
		actual := make([][]int16, 0)
		i := 0
		s := SlicesInt16{given}
		for el := range s.Product() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]int16{{1, 2}, {3, 4}}, [][]int16{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]int16{{1, 2}, {3}, {4, 5}}, [][]int16{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}

func TestAsyncSliceAnyInt32(t *testing.T) {
	f := func(check func(t int32) bool, given []int32, expected bool) {
		s := AsyncSliceInt32{data: given, workers: 2}
		actual := s.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int32) bool { return (t % 2) == 0 }

	f(isEven, []int32{}, false)
	f(isEven, []int32{1}, false)
	f(isEven, []int32{1, 3}, false)
	f(isEven, []int32{2}, true)
	f(isEven, []int32{1, 2}, true)
	f(isEven, []int32{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []int32{1, 3, 5, 7, 9, 12}, true)
}

func TestAsyncSliceAllInt32(t *testing.T) {
	f := func(check func(t int32) bool, given []int32, expected bool) {
		s := AsyncSliceInt32{data: given, workers: 2}
		actual := s.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int32) bool { return (t % 2) == 0 }

	f(isEven, []int32{}, true)
	f(isEven, []int32{1}, false)
	f(isEven, []int32{1, 3}, false)
	f(isEven, []int32{2}, true)
	f(isEven, []int32{2, 4}, true)
	f(isEven, []int32{2, 3}, false)
	f(isEven, []int32{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []int32{2, 4, 6, 8, 10, 11}, false)
}

func TestAsyncSliceEachInt32(t *testing.T) {
	f := func(given []int32) {
		s := AsyncSliceInt32{data: given, workers: 2}
		result := make(chan int32, len(given))
		mapper := func(t int32) { result <- t }
		s.Each(mapper)
		close(result)
		actual := ChannelInt32{result}.ToSlice()
		sorted := SliceInt32{actual}.Sort()
		assert.Equal(t, given, sorted, "they should be equal")
	}

	f([]int32{})
	f([]int32{1})
	f([]int32{1, 2, 3})
	f([]int32{1, 2, 3, 4, 5, 6, 7})
}

func TestAsyncSliceFilterInt32(t *testing.T) {
	f := func(given []int32, expected []int32) {
		filter := func(t int32) bool { return t > 10 }
		s := AsyncSliceInt32{data: given, workers: 2}
		actual := s.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int32{}, []int32{})
	f([]int32{5}, []int32{})
	f([]int32{15}, []int32{15})
	f([]int32{9, 11, 12, 13, 6}, []int32{11, 12, 13})
}

func TestAsyncSliceMapInt32Int(t *testing.T) {
	f := func(mapper func(t int32) int, given []int32, expected []int) {
		s := AsyncSliceInt32{data: given, workers: 2}
		actual := s.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int { return int((t * 2)) }

	f(double, []int32{}, []int{})
	f(double, []int32{1}, []int{2})
	f(double, []int32{1, 2, 3}, []int{2, 4, 6})
}

func TestAsyncSliceMapInt32Int8(t *testing.T) {
	f := func(mapper func(t int32) int8, given []int32, expected []int8) {
		s := AsyncSliceInt32{data: given, workers: 2}
		actual := s.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int8 { return int8((t * 2)) }

	f(double, []int32{}, []int8{})
	f(double, []int32{1}, []int8{2})
	f(double, []int32{1, 2, 3}, []int8{2, 4, 6})
}

func TestAsyncSliceMapInt32Int16(t *testing.T) {
	f := func(mapper func(t int32) int16, given []int32, expected []int16) {
		s := AsyncSliceInt32{data: given, workers: 2}
		actual := s.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int16 { return int16((t * 2)) }

	f(double, []int32{}, []int16{})
	f(double, []int32{1}, []int16{2})
	f(double, []int32{1, 2, 3}, []int16{2, 4, 6})
}

func TestAsyncSliceMapInt32Int32(t *testing.T) {
	f := func(mapper func(t int32) int32, given []int32, expected []int32) {
		s := AsyncSliceInt32{data: given, workers: 2}
		actual := s.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int32 { return int32((t * 2)) }

	f(double, []int32{}, []int32{})
	f(double, []int32{1}, []int32{2})
	f(double, []int32{1, 2, 3}, []int32{2, 4, 6})
}

func TestAsyncSliceMapInt32Int64(t *testing.T) {
	f := func(mapper func(t int32) int64, given []int32, expected []int64) {
		s := AsyncSliceInt32{data: given, workers: 2}
		actual := s.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int64 { return int64((t * 2)) }

	f(double, []int32{}, []int64{})
	f(double, []int32{1}, []int64{2})
	f(double, []int32{1, 2, 3}, []int64{2, 4, 6})
}

func TestChannelToSliceInt32(t *testing.T) {
	s := SequenceInt32{}
	f := func(start int32, stop int32, step int32, expected []int32) {
		seq := s.Range(start, stop, step)
		actual := ChannelInt32{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int32{1, 2, 3})
}

func TestSequenceExponentialInt32(t *testing.T) {
	s := SequenceInt32{}
	f := func(start int32, factor int32, count int, expected []int32) {
		seq := s.Exponential(start, factor)
		actual := ChannelInt32{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []int32{1, 2, 4, 8})
}

func TestSequenceRangeInt32(t *testing.T) {
	s := SequenceInt32{}
	f := func(start int32, stop int32, step int32, expected []int32) {
		seq := s.Range(start, stop, step)
		actual := ChannelInt32{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int32{1, 2, 3})
	f(3, 0, -1, []int32{3, 2, 1})
}

func TestSequenceRepeatInt32(t *testing.T) {
	s := SequenceInt32{}
	f := func(count int, given int32, expected []int32) {
		seq := s.Repeat(given)
		actual := ChannelInt32{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []int32{1, 1})
}

func TestSequenceTakeInt32(t *testing.T) {
	s := SequenceInt32{}
	f := func(count int, given int32, expected []int32) {
		seq := s.Repeat(given)
		actual := ChannelInt32{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int32{})
	f(1, 1, []int32{1})
	f(2, 1, []int32{1, 1})
}

func TestSliceAnyInt32(t *testing.T) {
	f := func(check func(t int32) bool, given []int32, expected bool) {
		actual := SliceInt32{given}.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int32) bool { return (t % 2) == 0 }

	f(isEven, []int32{}, false)
	f(isEven, []int32{1, 3}, false)
	f(isEven, []int32{2}, true)
	f(isEven, []int32{1, 2}, true)
}

func TestSliceAllInt32(t *testing.T) {
	f := func(check func(t int32) bool, given []int32, expected bool) {
		actual := SliceInt32{given}.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int32) bool { return (t % 2) == 0 }

	f(isEven, []int32{}, true)
	f(isEven, []int32{2}, true)
	f(isEven, []int32{1}, false)
	f(isEven, []int32{2, 4}, true)
	f(isEven, []int32{2, 4, 1}, false)
	f(isEven, []int32{1, 2, 4}, false)
}

func TestSliceChunkByInt32Int(t *testing.T) {
	f := func(mapper func(t int32) int, given []int32, expected [][]int32) {
		actual := SliceInt32{given}.ChunkByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int { return int((t % 2)) }

	f(remainder, []int32{1}, [][]int32{{1}})
	f(remainder, []int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f(remainder, []int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt32Int8(t *testing.T) {
	f := func(mapper func(t int32) int8, given []int32, expected [][]int32) {
		actual := SliceInt32{given}.ChunkByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int8 { return int8((t % 2)) }

	f(remainder, []int32{1}, [][]int32{{1}})
	f(remainder, []int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f(remainder, []int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt32Int16(t *testing.T) {
	f := func(mapper func(t int32) int16, given []int32, expected [][]int32) {
		actual := SliceInt32{given}.ChunkByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int16 { return int16((t % 2)) }

	f(remainder, []int32{1}, [][]int32{{1}})
	f(remainder, []int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f(remainder, []int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt32Int32(t *testing.T) {
	f := func(mapper func(t int32) int32, given []int32, expected [][]int32) {
		actual := SliceInt32{given}.ChunkByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int32 { return int32((t % 2)) }

	f(remainder, []int32{1}, [][]int32{{1}})
	f(remainder, []int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f(remainder, []int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt32Int64(t *testing.T) {
	f := func(mapper func(t int32) int64, given []int32, expected [][]int32) {
		actual := SliceInt32{given}.ChunkByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int64 { return int64((t % 2)) }

	f(remainder, []int32{1}, [][]int32{{1}})
	f(remainder, []int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f(remainder, []int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEveryInt32(t *testing.T) {
	f := func(count int, given []int32, expected [][]int32) {
		actual := SliceInt32{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int32{1, 2, 3, 4}, [][]int32{[]int32{1, 2}, []int32{3, 4}})
	f(2, []int32{1, 2, 3, 4, 5}, [][]int32{[]int32{1, 2}, []int32{3, 4}, []int32{5}})
}

func TestSliceFilterInt32(t *testing.T) {
	f := func(filter func(t int32) bool, given []int32, expected []int32) {
		actual := SliceInt32{given}.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t int32) bool { return t > 0 }

	f(filterPositive, []int32{1, -1, 2, -2, 3, -3}, []int32{1, 2, 3})
	f(filterPositive, []int32{1, 2, 3}, []int32{1, 2, 3})
	f(filterPositive, []int32{-1, -2, -3}, []int32{})
}

func TestSliceGroupByInt32Int(t *testing.T) {
	f := func(mapper func(t int32) int, given []int32, expected map[int][]int32) {
		actual := SliceInt32{given}.GroupByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int { return int((t % 2)) }

	f(remainder, []int32{}, map[int][]int32{})
	f(remainder, []int32{1}, map[int][]int32{1: {1}})
	f(remainder, []int32{1, 3, 2, 4, 5}, map[int][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt32Int8(t *testing.T) {
	f := func(mapper func(t int32) int8, given []int32, expected map[int8][]int32) {
		actual := SliceInt32{given}.GroupByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int8 { return int8((t % 2)) }

	f(remainder, []int32{}, map[int8][]int32{})
	f(remainder, []int32{1}, map[int8][]int32{1: {1}})
	f(remainder, []int32{1, 3, 2, 4, 5}, map[int8][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt32Int16(t *testing.T) {
	f := func(mapper func(t int32) int16, given []int32, expected map[int16][]int32) {
		actual := SliceInt32{given}.GroupByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int16 { return int16((t % 2)) }

	f(remainder, []int32{}, map[int16][]int32{})
	f(remainder, []int32{1}, map[int16][]int32{1: {1}})
	f(remainder, []int32{1, 3, 2, 4, 5}, map[int16][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt32Int32(t *testing.T) {
	f := func(mapper func(t int32) int32, given []int32, expected map[int32][]int32) {
		actual := SliceInt32{given}.GroupByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int32 { return int32((t % 2)) }

	f(remainder, []int32{}, map[int32][]int32{})
	f(remainder, []int32{1}, map[int32][]int32{1: {1}})
	f(remainder, []int32{1, 3, 2, 4, 5}, map[int32][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt32Int64(t *testing.T) {
	f := func(mapper func(t int32) int64, given []int32, expected map[int64][]int32) {
		actual := SliceInt32{given}.GroupByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int64 { return int64((t % 2)) }

	f(remainder, []int32{}, map[int64][]int32{})
	f(remainder, []int32{1}, map[int64][]int32{1: {1}})
	f(remainder, []int32{1, 3, 2, 4, 5}, map[int64][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceIntersperseInt32(t *testing.T) {
	f := func(el int32, given []int32, expected []int32) {
		actual := SliceInt32{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int32{1, 2, 3}, []int32{1, 0, 2, 0, 3})
}

func TestSliceMapInt32Int(t *testing.T) {
	f := func(mapper func(t int32) int, given []int32, expected []int) {
		actual := SliceInt32{given}.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int { return int((t * 2)) }

	f(double, []int32{1, 2, 3}, []int{2, 4, 6})
}

func TestSliceMapInt32Int8(t *testing.T) {
	f := func(mapper func(t int32) int8, given []int32, expected []int8) {
		actual := SliceInt32{given}.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int8 { return int8((t * 2)) }

	f(double, []int32{1, 2, 3}, []int8{2, 4, 6})
}

func TestSliceMapInt32Int16(t *testing.T) {
	f := func(mapper func(t int32) int16, given []int32, expected []int16) {
		actual := SliceInt32{given}.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int16 { return int16((t * 2)) }

	f(double, []int32{1, 2, 3}, []int16{2, 4, 6})
}

func TestSliceMapInt32Int32(t *testing.T) {
	f := func(mapper func(t int32) int32, given []int32, expected []int32) {
		actual := SliceInt32{given}.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int32 { return int32((t * 2)) }

	f(double, []int32{1, 2, 3}, []int32{2, 4, 6})
}

func TestSliceMapInt32Int64(t *testing.T) {
	f := func(mapper func(t int32) int64, given []int32, expected []int64) {
		actual := SliceInt32{given}.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int64 { return int64((t * 2)) }

	f(double, []int32{1, 2, 3}, []int64{2, 4, 6})
}

func TestSlicesProductInt32(t *testing.T) {
	f := func(given [][]int32, expected [][]int32) {
		actual := make([][]int32, 0)
		i := 0
		s := SlicesInt32{given}
		for el := range s.Product() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]int32{{1, 2}, {3, 4}}, [][]int32{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]int32{{1, 2}, {3}, {4, 5}}, [][]int32{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}

func TestAsyncSliceAnyInt64(t *testing.T) {
	f := func(check func(t int64) bool, given []int64, expected bool) {
		s := AsyncSliceInt64{data: given, workers: 2}
		actual := s.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int64) bool { return (t % 2) == 0 }

	f(isEven, []int64{}, false)
	f(isEven, []int64{1}, false)
	f(isEven, []int64{1, 3}, false)
	f(isEven, []int64{2}, true)
	f(isEven, []int64{1, 2}, true)
	f(isEven, []int64{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []int64{1, 3, 5, 7, 9, 12}, true)
}

func TestAsyncSliceAllInt64(t *testing.T) {
	f := func(check func(t int64) bool, given []int64, expected bool) {
		s := AsyncSliceInt64{data: given, workers: 2}
		actual := s.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int64) bool { return (t % 2) == 0 }

	f(isEven, []int64{}, true)
	f(isEven, []int64{1}, false)
	f(isEven, []int64{1, 3}, false)
	f(isEven, []int64{2}, true)
	f(isEven, []int64{2, 4}, true)
	f(isEven, []int64{2, 3}, false)
	f(isEven, []int64{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []int64{2, 4, 6, 8, 10, 11}, false)
}

func TestAsyncSliceEachInt64(t *testing.T) {
	f := func(given []int64) {
		s := AsyncSliceInt64{data: given, workers: 2}
		result := make(chan int64, len(given))
		mapper := func(t int64) { result <- t }
		s.Each(mapper)
		close(result)
		actual := ChannelInt64{result}.ToSlice()
		sorted := SliceInt64{actual}.Sort()
		assert.Equal(t, given, sorted, "they should be equal")
	}

	f([]int64{})
	f([]int64{1})
	f([]int64{1, 2, 3})
	f([]int64{1, 2, 3, 4, 5, 6, 7})
}

func TestAsyncSliceFilterInt64(t *testing.T) {
	f := func(given []int64, expected []int64) {
		filter := func(t int64) bool { return t > 10 }
		s := AsyncSliceInt64{data: given, workers: 2}
		actual := s.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int64{}, []int64{})
	f([]int64{5}, []int64{})
	f([]int64{15}, []int64{15})
	f([]int64{9, 11, 12, 13, 6}, []int64{11, 12, 13})
}

func TestAsyncSliceMapInt64Int(t *testing.T) {
	f := func(mapper func(t int64) int, given []int64, expected []int) {
		s := AsyncSliceInt64{data: given, workers: 2}
		actual := s.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int { return int((t * 2)) }

	f(double, []int64{}, []int{})
	f(double, []int64{1}, []int{2})
	f(double, []int64{1, 2, 3}, []int{2, 4, 6})
}

func TestAsyncSliceMapInt64Int8(t *testing.T) {
	f := func(mapper func(t int64) int8, given []int64, expected []int8) {
		s := AsyncSliceInt64{data: given, workers: 2}
		actual := s.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int8 { return int8((t * 2)) }

	f(double, []int64{}, []int8{})
	f(double, []int64{1}, []int8{2})
	f(double, []int64{1, 2, 3}, []int8{2, 4, 6})
}

func TestAsyncSliceMapInt64Int16(t *testing.T) {
	f := func(mapper func(t int64) int16, given []int64, expected []int16) {
		s := AsyncSliceInt64{data: given, workers: 2}
		actual := s.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int16 { return int16((t * 2)) }

	f(double, []int64{}, []int16{})
	f(double, []int64{1}, []int16{2})
	f(double, []int64{1, 2, 3}, []int16{2, 4, 6})
}

func TestAsyncSliceMapInt64Int32(t *testing.T) {
	f := func(mapper func(t int64) int32, given []int64, expected []int32) {
		s := AsyncSliceInt64{data: given, workers: 2}
		actual := s.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int32 { return int32((t * 2)) }

	f(double, []int64{}, []int32{})
	f(double, []int64{1}, []int32{2})
	f(double, []int64{1, 2, 3}, []int32{2, 4, 6})
}

func TestAsyncSliceMapInt64Int64(t *testing.T) {
	f := func(mapper func(t int64) int64, given []int64, expected []int64) {
		s := AsyncSliceInt64{data: given, workers: 2}
		actual := s.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int64 { return int64((t * 2)) }

	f(double, []int64{}, []int64{})
	f(double, []int64{1}, []int64{2})
	f(double, []int64{1, 2, 3}, []int64{2, 4, 6})
}

func TestChannelToSliceInt64(t *testing.T) {
	s := SequenceInt64{}
	f := func(start int64, stop int64, step int64, expected []int64) {
		seq := s.Range(start, stop, step)
		actual := ChannelInt64{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int64{1, 2, 3})
}

func TestSequenceExponentialInt64(t *testing.T) {
	s := SequenceInt64{}
	f := func(start int64, factor int64, count int, expected []int64) {
		seq := s.Exponential(start, factor)
		actual := ChannelInt64{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []int64{1, 2, 4, 8})
}

func TestSequenceRangeInt64(t *testing.T) {
	s := SequenceInt64{}
	f := func(start int64, stop int64, step int64, expected []int64) {
		seq := s.Range(start, stop, step)
		actual := ChannelInt64{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int64{1, 2, 3})
	f(3, 0, -1, []int64{3, 2, 1})
}

func TestSequenceRepeatInt64(t *testing.T) {
	s := SequenceInt64{}
	f := func(count int, given int64, expected []int64) {
		seq := s.Repeat(given)
		actual := ChannelInt64{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []int64{1, 1})
}

func TestSequenceTakeInt64(t *testing.T) {
	s := SequenceInt64{}
	f := func(count int, given int64, expected []int64) {
		seq := s.Repeat(given)
		actual := ChannelInt64{seq}.Take(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int64{})
	f(1, 1, []int64{1})
	f(2, 1, []int64{1, 1})
}

func TestSliceAnyInt64(t *testing.T) {
	f := func(check func(t int64) bool, given []int64, expected bool) {
		actual := SliceInt64{given}.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int64) bool { return (t % 2) == 0 }

	f(isEven, []int64{}, false)
	f(isEven, []int64{1, 3}, false)
	f(isEven, []int64{2}, true)
	f(isEven, []int64{1, 2}, true)
}

func TestSliceAllInt64(t *testing.T) {
	f := func(check func(t int64) bool, given []int64, expected bool) {
		actual := SliceInt64{given}.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int64) bool { return (t % 2) == 0 }

	f(isEven, []int64{}, true)
	f(isEven, []int64{2}, true)
	f(isEven, []int64{1}, false)
	f(isEven, []int64{2, 4}, true)
	f(isEven, []int64{2, 4, 1}, false)
	f(isEven, []int64{1, 2, 4}, false)
}

func TestSliceChunkByInt64Int(t *testing.T) {
	f := func(mapper func(t int64) int, given []int64, expected [][]int64) {
		actual := SliceInt64{given}.ChunkByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int { return int((t % 2)) }

	f(remainder, []int64{1}, [][]int64{{1}})
	f(remainder, []int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f(remainder, []int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt64Int8(t *testing.T) {
	f := func(mapper func(t int64) int8, given []int64, expected [][]int64) {
		actual := SliceInt64{given}.ChunkByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int8 { return int8((t % 2)) }

	f(remainder, []int64{1}, [][]int64{{1}})
	f(remainder, []int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f(remainder, []int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt64Int16(t *testing.T) {
	f := func(mapper func(t int64) int16, given []int64, expected [][]int64) {
		actual := SliceInt64{given}.ChunkByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int16 { return int16((t % 2)) }

	f(remainder, []int64{1}, [][]int64{{1}})
	f(remainder, []int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f(remainder, []int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt64Int32(t *testing.T) {
	f := func(mapper func(t int64) int32, given []int64, expected [][]int64) {
		actual := SliceInt64{given}.ChunkByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int32 { return int32((t % 2)) }

	f(remainder, []int64{1}, [][]int64{{1}})
	f(remainder, []int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f(remainder, []int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt64Int64(t *testing.T) {
	f := func(mapper func(t int64) int64, given []int64, expected [][]int64) {
		actual := SliceInt64{given}.ChunkByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int64 { return int64((t % 2)) }

	f(remainder, []int64{1}, [][]int64{{1}})
	f(remainder, []int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f(remainder, []int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEveryInt64(t *testing.T) {
	f := func(count int, given []int64, expected [][]int64) {
		actual := SliceInt64{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int64{1, 2, 3, 4}, [][]int64{[]int64{1, 2}, []int64{3, 4}})
	f(2, []int64{1, 2, 3, 4, 5}, [][]int64{[]int64{1, 2}, []int64{3, 4}, []int64{5}})
}

func TestSliceFilterInt64(t *testing.T) {
	f := func(filter func(t int64) bool, given []int64, expected []int64) {
		actual := SliceInt64{given}.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t int64) bool { return t > 0 }

	f(filterPositive, []int64{1, -1, 2, -2, 3, -3}, []int64{1, 2, 3})
	f(filterPositive, []int64{1, 2, 3}, []int64{1, 2, 3})
	f(filterPositive, []int64{-1, -2, -3}, []int64{})
}

func TestSliceGroupByInt64Int(t *testing.T) {
	f := func(mapper func(t int64) int, given []int64, expected map[int][]int64) {
		actual := SliceInt64{given}.GroupByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int { return int((t % 2)) }

	f(remainder, []int64{}, map[int][]int64{})
	f(remainder, []int64{1}, map[int][]int64{1: {1}})
	f(remainder, []int64{1, 3, 2, 4, 5}, map[int][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt64Int8(t *testing.T) {
	f := func(mapper func(t int64) int8, given []int64, expected map[int8][]int64) {
		actual := SliceInt64{given}.GroupByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int8 { return int8((t % 2)) }

	f(remainder, []int64{}, map[int8][]int64{})
	f(remainder, []int64{1}, map[int8][]int64{1: {1}})
	f(remainder, []int64{1, 3, 2, 4, 5}, map[int8][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt64Int16(t *testing.T) {
	f := func(mapper func(t int64) int16, given []int64, expected map[int16][]int64) {
		actual := SliceInt64{given}.GroupByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int16 { return int16((t % 2)) }

	f(remainder, []int64{}, map[int16][]int64{})
	f(remainder, []int64{1}, map[int16][]int64{1: {1}})
	f(remainder, []int64{1, 3, 2, 4, 5}, map[int16][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt64Int32(t *testing.T) {
	f := func(mapper func(t int64) int32, given []int64, expected map[int32][]int64) {
		actual := SliceInt64{given}.GroupByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int32 { return int32((t % 2)) }

	f(remainder, []int64{}, map[int32][]int64{})
	f(remainder, []int64{1}, map[int32][]int64{1: {1}})
	f(remainder, []int64{1, 3, 2, 4, 5}, map[int32][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt64Int64(t *testing.T) {
	f := func(mapper func(t int64) int64, given []int64, expected map[int64][]int64) {
		actual := SliceInt64{given}.GroupByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int64 { return int64((t % 2)) }

	f(remainder, []int64{}, map[int64][]int64{})
	f(remainder, []int64{1}, map[int64][]int64{1: {1}})
	f(remainder, []int64{1, 3, 2, 4, 5}, map[int64][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceIntersperseInt64(t *testing.T) {
	f := func(el int64, given []int64, expected []int64) {
		actual := SliceInt64{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int64{1, 2, 3}, []int64{1, 0, 2, 0, 3})
}

func TestSliceMapInt64Int(t *testing.T) {
	f := func(mapper func(t int64) int, given []int64, expected []int) {
		actual := SliceInt64{given}.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int { return int((t * 2)) }

	f(double, []int64{1, 2, 3}, []int{2, 4, 6})
}

func TestSliceMapInt64Int8(t *testing.T) {
	f := func(mapper func(t int64) int8, given []int64, expected []int8) {
		actual := SliceInt64{given}.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int8 { return int8((t * 2)) }

	f(double, []int64{1, 2, 3}, []int8{2, 4, 6})
}

func TestSliceMapInt64Int16(t *testing.T) {
	f := func(mapper func(t int64) int16, given []int64, expected []int16) {
		actual := SliceInt64{given}.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int16 { return int16((t * 2)) }

	f(double, []int64{1, 2, 3}, []int16{2, 4, 6})
}

func TestSliceMapInt64Int32(t *testing.T) {
	f := func(mapper func(t int64) int32, given []int64, expected []int32) {
		actual := SliceInt64{given}.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int32 { return int32((t * 2)) }

	f(double, []int64{1, 2, 3}, []int32{2, 4, 6})
}

func TestSliceMapInt64Int64(t *testing.T) {
	f := func(mapper func(t int64) int64, given []int64, expected []int64) {
		actual := SliceInt64{given}.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int64 { return int64((t * 2)) }

	f(double, []int64{1, 2, 3}, []int64{2, 4, 6})
}

func TestSlicesProductInt64(t *testing.T) {
	f := func(given [][]int64, expected [][]int64) {
		actual := make([][]int64, 0)
		i := 0
		s := SlicesInt64{given}
		for el := range s.Product() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]int64{{1, 2}, {3, 4}}, [][]int64{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]int64{{1, 2}, {3}, {4, 5}}, [][]int64{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}
