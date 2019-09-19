package genesis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkByIntInt(t *testing.T) {
	f := func(mapper func(t int) int, given []int, expected [][]int) {
		actual := SliceInt{given}.ChunkByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int { return int((t % 2)) }

	f(remainder, []int{1}, [][]int{{1}})
	f(remainder, []int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f(remainder, []int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestChunkByIntInt8(t *testing.T) {
	f := func(mapper func(t int) int8, given []int, expected [][]int) {
		actual := SliceInt{given}.ChunkByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int8 { return int8((t % 2)) }

	f(remainder, []int{1}, [][]int{{1}})
	f(remainder, []int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f(remainder, []int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestChunkByIntInt16(t *testing.T) {
	f := func(mapper func(t int) int16, given []int, expected [][]int) {
		actual := SliceInt{given}.ChunkByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int16 { return int16((t % 2)) }

	f(remainder, []int{1}, [][]int{{1}})
	f(remainder, []int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f(remainder, []int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestChunkByIntInt32(t *testing.T) {
	f := func(mapper func(t int) int32, given []int, expected [][]int) {
		actual := SliceInt{given}.ChunkByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int32 { return int32((t % 2)) }

	f(remainder, []int{1}, [][]int{{1}})
	f(remainder, []int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f(remainder, []int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestChunkByIntInt64(t *testing.T) {
	f := func(mapper func(t int) int64, given []int, expected [][]int) {
		actual := SliceInt{given}.ChunkByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int64 { return int64((t % 2)) }

	f(remainder, []int{1}, [][]int{{1}})
	f(remainder, []int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f(remainder, []int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestChunkEveryInt(t *testing.T) {
	f := func(count int, given []int, expected [][]int) {
		actual := SliceInt{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int{1, 2, 3, 4}, [][]int{[]int{1, 2}, []int{3, 4}})
	f(2, []int{1, 2, 3, 4, 5}, [][]int{[]int{1, 2}, []int{3, 4}, []int{5}})
}

func TestFilterInt(t *testing.T) {
	f := func(filter func(t int) bool, given []int, expected []int) {
		actual := SliceInt{given}.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t int) bool { return t > 0 }

	f(filterPositive, []int{1, -1, 2, -2, 3, -3}, []int{1, 2, 3})
	f(filterPositive, []int{1, 2, 3}, []int{1, 2, 3})
	f(filterPositive, []int{-1, -2, -3}, []int{})
}

func TestGroupByIntInt(t *testing.T) {
	f := func(mapper func(t int) int, given []int, expected map[int][]int) {
		actual := SliceInt{given}.GroupByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int { return int((t % 2)) }

	f(remainder, []int{}, map[int][]int{})
	f(remainder, []int{1}, map[int][]int{1: {1}})
	f(remainder, []int{1, 3, 2, 4, 5}, map[int][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByIntInt8(t *testing.T) {
	f := func(mapper func(t int) int8, given []int, expected map[int8][]int) {
		actual := SliceInt{given}.GroupByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int8 { return int8((t % 2)) }

	f(remainder, []int{}, map[int8][]int{})
	f(remainder, []int{1}, map[int8][]int{1: {1}})
	f(remainder, []int{1, 3, 2, 4, 5}, map[int8][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByIntInt16(t *testing.T) {
	f := func(mapper func(t int) int16, given []int, expected map[int16][]int) {
		actual := SliceInt{given}.GroupByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int16 { return int16((t % 2)) }

	f(remainder, []int{}, map[int16][]int{})
	f(remainder, []int{1}, map[int16][]int{1: {1}})
	f(remainder, []int{1, 3, 2, 4, 5}, map[int16][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByIntInt32(t *testing.T) {
	f := func(mapper func(t int) int32, given []int, expected map[int32][]int) {
		actual := SliceInt{given}.GroupByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int32 { return int32((t % 2)) }

	f(remainder, []int{}, map[int32][]int{})
	f(remainder, []int{1}, map[int32][]int{1: {1}})
	f(remainder, []int{1, 3, 2, 4, 5}, map[int32][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByIntInt64(t *testing.T) {
	f := func(mapper func(t int) int64, given []int, expected map[int64][]int) {
		actual := SliceInt{given}.GroupByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int) int64 { return int64((t % 2)) }

	f(remainder, []int{}, map[int64][]int{})
	f(remainder, []int{1}, map[int64][]int{1: {1}})
	f(remainder, []int{1, 3, 2, 4, 5}, map[int64][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestIntersperseInt(t *testing.T) {
	f := func(el int, given []int, expected []int) {
		actual := SliceInt{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int{1, 2, 3}, []int{1, 0, 2, 0, 3})
}

func TestMapIntInt(t *testing.T) {
	f := func(mapper func(t int) int, given []int, expected []int) {
		actual := SliceInt{given}.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int { return int((t * 2)) }

	f(double, []int{1, 2, 3}, []int{2, 4, 6})
}

func TestMapIntInt8(t *testing.T) {
	f := func(mapper func(t int) int8, given []int, expected []int8) {
		actual := SliceInt{given}.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int8 { return int8((t * 2)) }

	f(double, []int{1, 2, 3}, []int8{2, 4, 6})
}

func TestMapIntInt16(t *testing.T) {
	f := func(mapper func(t int) int16, given []int, expected []int16) {
		actual := SliceInt{given}.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int16 { return int16((t * 2)) }

	f(double, []int{1, 2, 3}, []int16{2, 4, 6})
}

func TestMapIntInt32(t *testing.T) {
	f := func(mapper func(t int) int32, given []int, expected []int32) {
		actual := SliceInt{given}.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int32 { return int32((t * 2)) }

	f(double, []int{1, 2, 3}, []int32{2, 4, 6})
}

func TestMapIntInt64(t *testing.T) {
	f := func(mapper func(t int) int64, given []int, expected []int64) {
		actual := SliceInt{given}.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int64 { return int64((t * 2)) }

	f(double, []int{1, 2, 3}, []int64{2, 4, 6})
}

func TestChunkByInt8Int(t *testing.T) {
	f := func(mapper func(t int8) int, given []int8, expected [][]int8) {
		actual := SliceInt8{given}.ChunkByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int { return int((t % 2)) }

	f(remainder, []int8{1}, [][]int8{{1}})
	f(remainder, []int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f(remainder, []int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt8Int8(t *testing.T) {
	f := func(mapper func(t int8) int8, given []int8, expected [][]int8) {
		actual := SliceInt8{given}.ChunkByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int8 { return int8((t % 2)) }

	f(remainder, []int8{1}, [][]int8{{1}})
	f(remainder, []int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f(remainder, []int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt8Int16(t *testing.T) {
	f := func(mapper func(t int8) int16, given []int8, expected [][]int8) {
		actual := SliceInt8{given}.ChunkByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int16 { return int16((t % 2)) }

	f(remainder, []int8{1}, [][]int8{{1}})
	f(remainder, []int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f(remainder, []int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt8Int32(t *testing.T) {
	f := func(mapper func(t int8) int32, given []int8, expected [][]int8) {
		actual := SliceInt8{given}.ChunkByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int32 { return int32((t % 2)) }

	f(remainder, []int8{1}, [][]int8{{1}})
	f(remainder, []int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f(remainder, []int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt8Int64(t *testing.T) {
	f := func(mapper func(t int8) int64, given []int8, expected [][]int8) {
		actual := SliceInt8{given}.ChunkByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int64 { return int64((t % 2)) }

	f(remainder, []int8{1}, [][]int8{{1}})
	f(remainder, []int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f(remainder, []int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestChunkEveryInt8(t *testing.T) {
	f := func(count int, given []int8, expected [][]int8) {
		actual := SliceInt8{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int8{1, 2, 3, 4}, [][]int8{[]int8{1, 2}, []int8{3, 4}})
	f(2, []int8{1, 2, 3, 4, 5}, [][]int8{[]int8{1, 2}, []int8{3, 4}, []int8{5}})
}

func TestFilterInt8(t *testing.T) {
	f := func(filter func(t int8) bool, given []int8, expected []int8) {
		actual := SliceInt8{given}.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t int8) bool { return t > 0 }

	f(filterPositive, []int8{1, -1, 2, -2, 3, -3}, []int8{1, 2, 3})
	f(filterPositive, []int8{1, 2, 3}, []int8{1, 2, 3})
	f(filterPositive, []int8{-1, -2, -3}, []int8{})
}

func TestGroupByInt8Int(t *testing.T) {
	f := func(mapper func(t int8) int, given []int8, expected map[int][]int8) {
		actual := SliceInt8{given}.GroupByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int { return int((t % 2)) }

	f(remainder, []int8{}, map[int][]int8{})
	f(remainder, []int8{1}, map[int][]int8{1: {1}})
	f(remainder, []int8{1, 3, 2, 4, 5}, map[int][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt8Int8(t *testing.T) {
	f := func(mapper func(t int8) int8, given []int8, expected map[int8][]int8) {
		actual := SliceInt8{given}.GroupByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int8 { return int8((t % 2)) }

	f(remainder, []int8{}, map[int8][]int8{})
	f(remainder, []int8{1}, map[int8][]int8{1: {1}})
	f(remainder, []int8{1, 3, 2, 4, 5}, map[int8][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt8Int16(t *testing.T) {
	f := func(mapper func(t int8) int16, given []int8, expected map[int16][]int8) {
		actual := SliceInt8{given}.GroupByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int16 { return int16((t % 2)) }

	f(remainder, []int8{}, map[int16][]int8{})
	f(remainder, []int8{1}, map[int16][]int8{1: {1}})
	f(remainder, []int8{1, 3, 2, 4, 5}, map[int16][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt8Int32(t *testing.T) {
	f := func(mapper func(t int8) int32, given []int8, expected map[int32][]int8) {
		actual := SliceInt8{given}.GroupByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int32 { return int32((t % 2)) }

	f(remainder, []int8{}, map[int32][]int8{})
	f(remainder, []int8{1}, map[int32][]int8{1: {1}})
	f(remainder, []int8{1, 3, 2, 4, 5}, map[int32][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt8Int64(t *testing.T) {
	f := func(mapper func(t int8) int64, given []int8, expected map[int64][]int8) {
		actual := SliceInt8{given}.GroupByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int8) int64 { return int64((t % 2)) }

	f(remainder, []int8{}, map[int64][]int8{})
	f(remainder, []int8{1}, map[int64][]int8{1: {1}})
	f(remainder, []int8{1, 3, 2, 4, 5}, map[int64][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestIntersperseInt8(t *testing.T) {
	f := func(el int8, given []int8, expected []int8) {
		actual := SliceInt8{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int8{1, 2, 3}, []int8{1, 0, 2, 0, 3})
}

func TestMapInt8Int(t *testing.T) {
	f := func(mapper func(t int8) int, given []int8, expected []int) {
		actual := SliceInt8{given}.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int { return int((t * 2)) }

	f(double, []int8{1, 2, 3}, []int{2, 4, 6})
}

func TestMapInt8Int8(t *testing.T) {
	f := func(mapper func(t int8) int8, given []int8, expected []int8) {
		actual := SliceInt8{given}.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int8 { return int8((t * 2)) }

	f(double, []int8{1, 2, 3}, []int8{2, 4, 6})
}

func TestMapInt8Int16(t *testing.T) {
	f := func(mapper func(t int8) int16, given []int8, expected []int16) {
		actual := SliceInt8{given}.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int16 { return int16((t * 2)) }

	f(double, []int8{1, 2, 3}, []int16{2, 4, 6})
}

func TestMapInt8Int32(t *testing.T) {
	f := func(mapper func(t int8) int32, given []int8, expected []int32) {
		actual := SliceInt8{given}.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int32 { return int32((t * 2)) }

	f(double, []int8{1, 2, 3}, []int32{2, 4, 6})
}

func TestMapInt8Int64(t *testing.T) {
	f := func(mapper func(t int8) int64, given []int8, expected []int64) {
		actual := SliceInt8{given}.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int64 { return int64((t * 2)) }

	f(double, []int8{1, 2, 3}, []int64{2, 4, 6})
}

func TestChunkByInt16Int(t *testing.T) {
	f := func(mapper func(t int16) int, given []int16, expected [][]int16) {
		actual := SliceInt16{given}.ChunkByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int { return int((t % 2)) }

	f(remainder, []int16{1}, [][]int16{{1}})
	f(remainder, []int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f(remainder, []int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt16Int8(t *testing.T) {
	f := func(mapper func(t int16) int8, given []int16, expected [][]int16) {
		actual := SliceInt16{given}.ChunkByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int8 { return int8((t % 2)) }

	f(remainder, []int16{1}, [][]int16{{1}})
	f(remainder, []int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f(remainder, []int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt16Int16(t *testing.T) {
	f := func(mapper func(t int16) int16, given []int16, expected [][]int16) {
		actual := SliceInt16{given}.ChunkByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int16 { return int16((t % 2)) }

	f(remainder, []int16{1}, [][]int16{{1}})
	f(remainder, []int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f(remainder, []int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt16Int32(t *testing.T) {
	f := func(mapper func(t int16) int32, given []int16, expected [][]int16) {
		actual := SliceInt16{given}.ChunkByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int32 { return int32((t % 2)) }

	f(remainder, []int16{1}, [][]int16{{1}})
	f(remainder, []int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f(remainder, []int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt16Int64(t *testing.T) {
	f := func(mapper func(t int16) int64, given []int16, expected [][]int16) {
		actual := SliceInt16{given}.ChunkByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int64 { return int64((t % 2)) }

	f(remainder, []int16{1}, [][]int16{{1}})
	f(remainder, []int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f(remainder, []int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestChunkEveryInt16(t *testing.T) {
	f := func(count int, given []int16, expected [][]int16) {
		actual := SliceInt16{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int16{1, 2, 3, 4}, [][]int16{[]int16{1, 2}, []int16{3, 4}})
	f(2, []int16{1, 2, 3, 4, 5}, [][]int16{[]int16{1, 2}, []int16{3, 4}, []int16{5}})
}

func TestFilterInt16(t *testing.T) {
	f := func(filter func(t int16) bool, given []int16, expected []int16) {
		actual := SliceInt16{given}.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t int16) bool { return t > 0 }

	f(filterPositive, []int16{1, -1, 2, -2, 3, -3}, []int16{1, 2, 3})
	f(filterPositive, []int16{1, 2, 3}, []int16{1, 2, 3})
	f(filterPositive, []int16{-1, -2, -3}, []int16{})
}

func TestGroupByInt16Int(t *testing.T) {
	f := func(mapper func(t int16) int, given []int16, expected map[int][]int16) {
		actual := SliceInt16{given}.GroupByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int { return int((t % 2)) }

	f(remainder, []int16{}, map[int][]int16{})
	f(remainder, []int16{1}, map[int][]int16{1: {1}})
	f(remainder, []int16{1, 3, 2, 4, 5}, map[int][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt16Int8(t *testing.T) {
	f := func(mapper func(t int16) int8, given []int16, expected map[int8][]int16) {
		actual := SliceInt16{given}.GroupByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int8 { return int8((t % 2)) }

	f(remainder, []int16{}, map[int8][]int16{})
	f(remainder, []int16{1}, map[int8][]int16{1: {1}})
	f(remainder, []int16{1, 3, 2, 4, 5}, map[int8][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt16Int16(t *testing.T) {
	f := func(mapper func(t int16) int16, given []int16, expected map[int16][]int16) {
		actual := SliceInt16{given}.GroupByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int16 { return int16((t % 2)) }

	f(remainder, []int16{}, map[int16][]int16{})
	f(remainder, []int16{1}, map[int16][]int16{1: {1}})
	f(remainder, []int16{1, 3, 2, 4, 5}, map[int16][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt16Int32(t *testing.T) {
	f := func(mapper func(t int16) int32, given []int16, expected map[int32][]int16) {
		actual := SliceInt16{given}.GroupByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int32 { return int32((t % 2)) }

	f(remainder, []int16{}, map[int32][]int16{})
	f(remainder, []int16{1}, map[int32][]int16{1: {1}})
	f(remainder, []int16{1, 3, 2, 4, 5}, map[int32][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt16Int64(t *testing.T) {
	f := func(mapper func(t int16) int64, given []int16, expected map[int64][]int16) {
		actual := SliceInt16{given}.GroupByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int16) int64 { return int64((t % 2)) }

	f(remainder, []int16{}, map[int64][]int16{})
	f(remainder, []int16{1}, map[int64][]int16{1: {1}})
	f(remainder, []int16{1, 3, 2, 4, 5}, map[int64][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestIntersperseInt16(t *testing.T) {
	f := func(el int16, given []int16, expected []int16) {
		actual := SliceInt16{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int16{1, 2, 3}, []int16{1, 0, 2, 0, 3})
}

func TestMapInt16Int(t *testing.T) {
	f := func(mapper func(t int16) int, given []int16, expected []int) {
		actual := SliceInt16{given}.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int { return int((t * 2)) }

	f(double, []int16{1, 2, 3}, []int{2, 4, 6})
}

func TestMapInt16Int8(t *testing.T) {
	f := func(mapper func(t int16) int8, given []int16, expected []int8) {
		actual := SliceInt16{given}.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int8 { return int8((t * 2)) }

	f(double, []int16{1, 2, 3}, []int8{2, 4, 6})
}

func TestMapInt16Int16(t *testing.T) {
	f := func(mapper func(t int16) int16, given []int16, expected []int16) {
		actual := SliceInt16{given}.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int16 { return int16((t * 2)) }

	f(double, []int16{1, 2, 3}, []int16{2, 4, 6})
}

func TestMapInt16Int32(t *testing.T) {
	f := func(mapper func(t int16) int32, given []int16, expected []int32) {
		actual := SliceInt16{given}.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int32 { return int32((t * 2)) }

	f(double, []int16{1, 2, 3}, []int32{2, 4, 6})
}

func TestMapInt16Int64(t *testing.T) {
	f := func(mapper func(t int16) int64, given []int16, expected []int64) {
		actual := SliceInt16{given}.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int64 { return int64((t * 2)) }

	f(double, []int16{1, 2, 3}, []int64{2, 4, 6})
}

func TestChunkByInt32Int(t *testing.T) {
	f := func(mapper func(t int32) int, given []int32, expected [][]int32) {
		actual := SliceInt32{given}.ChunkByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int { return int((t % 2)) }

	f(remainder, []int32{1}, [][]int32{{1}})
	f(remainder, []int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f(remainder, []int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt32Int8(t *testing.T) {
	f := func(mapper func(t int32) int8, given []int32, expected [][]int32) {
		actual := SliceInt32{given}.ChunkByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int8 { return int8((t % 2)) }

	f(remainder, []int32{1}, [][]int32{{1}})
	f(remainder, []int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f(remainder, []int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt32Int16(t *testing.T) {
	f := func(mapper func(t int32) int16, given []int32, expected [][]int32) {
		actual := SliceInt32{given}.ChunkByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int16 { return int16((t % 2)) }

	f(remainder, []int32{1}, [][]int32{{1}})
	f(remainder, []int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f(remainder, []int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt32Int32(t *testing.T) {
	f := func(mapper func(t int32) int32, given []int32, expected [][]int32) {
		actual := SliceInt32{given}.ChunkByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int32 { return int32((t % 2)) }

	f(remainder, []int32{1}, [][]int32{{1}})
	f(remainder, []int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f(remainder, []int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt32Int64(t *testing.T) {
	f := func(mapper func(t int32) int64, given []int32, expected [][]int32) {
		actual := SliceInt32{given}.ChunkByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int64 { return int64((t % 2)) }

	f(remainder, []int32{1}, [][]int32{{1}})
	f(remainder, []int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f(remainder, []int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestChunkEveryInt32(t *testing.T) {
	f := func(count int, given []int32, expected [][]int32) {
		actual := SliceInt32{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int32{1, 2, 3, 4}, [][]int32{[]int32{1, 2}, []int32{3, 4}})
	f(2, []int32{1, 2, 3, 4, 5}, [][]int32{[]int32{1, 2}, []int32{3, 4}, []int32{5}})
}

func TestFilterInt32(t *testing.T) {
	f := func(filter func(t int32) bool, given []int32, expected []int32) {
		actual := SliceInt32{given}.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t int32) bool { return t > 0 }

	f(filterPositive, []int32{1, -1, 2, -2, 3, -3}, []int32{1, 2, 3})
	f(filterPositive, []int32{1, 2, 3}, []int32{1, 2, 3})
	f(filterPositive, []int32{-1, -2, -3}, []int32{})
}

func TestGroupByInt32Int(t *testing.T) {
	f := func(mapper func(t int32) int, given []int32, expected map[int][]int32) {
		actual := SliceInt32{given}.GroupByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int { return int((t % 2)) }

	f(remainder, []int32{}, map[int][]int32{})
	f(remainder, []int32{1}, map[int][]int32{1: {1}})
	f(remainder, []int32{1, 3, 2, 4, 5}, map[int][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt32Int8(t *testing.T) {
	f := func(mapper func(t int32) int8, given []int32, expected map[int8][]int32) {
		actual := SliceInt32{given}.GroupByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int8 { return int8((t % 2)) }

	f(remainder, []int32{}, map[int8][]int32{})
	f(remainder, []int32{1}, map[int8][]int32{1: {1}})
	f(remainder, []int32{1, 3, 2, 4, 5}, map[int8][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt32Int16(t *testing.T) {
	f := func(mapper func(t int32) int16, given []int32, expected map[int16][]int32) {
		actual := SliceInt32{given}.GroupByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int16 { return int16((t % 2)) }

	f(remainder, []int32{}, map[int16][]int32{})
	f(remainder, []int32{1}, map[int16][]int32{1: {1}})
	f(remainder, []int32{1, 3, 2, 4, 5}, map[int16][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt32Int32(t *testing.T) {
	f := func(mapper func(t int32) int32, given []int32, expected map[int32][]int32) {
		actual := SliceInt32{given}.GroupByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int32 { return int32((t % 2)) }

	f(remainder, []int32{}, map[int32][]int32{})
	f(remainder, []int32{1}, map[int32][]int32{1: {1}})
	f(remainder, []int32{1, 3, 2, 4, 5}, map[int32][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt32Int64(t *testing.T) {
	f := func(mapper func(t int32) int64, given []int32, expected map[int64][]int32) {
		actual := SliceInt32{given}.GroupByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int32) int64 { return int64((t % 2)) }

	f(remainder, []int32{}, map[int64][]int32{})
	f(remainder, []int32{1}, map[int64][]int32{1: {1}})
	f(remainder, []int32{1, 3, 2, 4, 5}, map[int64][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestIntersperseInt32(t *testing.T) {
	f := func(el int32, given []int32, expected []int32) {
		actual := SliceInt32{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int32{1, 2, 3}, []int32{1, 0, 2, 0, 3})
}

func TestMapInt32Int(t *testing.T) {
	f := func(mapper func(t int32) int, given []int32, expected []int) {
		actual := SliceInt32{given}.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int { return int((t * 2)) }

	f(double, []int32{1, 2, 3}, []int{2, 4, 6})
}

func TestMapInt32Int8(t *testing.T) {
	f := func(mapper func(t int32) int8, given []int32, expected []int8) {
		actual := SliceInt32{given}.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int8 { return int8((t * 2)) }

	f(double, []int32{1, 2, 3}, []int8{2, 4, 6})
}

func TestMapInt32Int16(t *testing.T) {
	f := func(mapper func(t int32) int16, given []int32, expected []int16) {
		actual := SliceInt32{given}.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int16 { return int16((t * 2)) }

	f(double, []int32{1, 2, 3}, []int16{2, 4, 6})
}

func TestMapInt32Int32(t *testing.T) {
	f := func(mapper func(t int32) int32, given []int32, expected []int32) {
		actual := SliceInt32{given}.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int32 { return int32((t * 2)) }

	f(double, []int32{1, 2, 3}, []int32{2, 4, 6})
}

func TestMapInt32Int64(t *testing.T) {
	f := func(mapper func(t int32) int64, given []int32, expected []int64) {
		actual := SliceInt32{given}.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int64 { return int64((t * 2)) }

	f(double, []int32{1, 2, 3}, []int64{2, 4, 6})
}

func TestChunkByInt64Int(t *testing.T) {
	f := func(mapper func(t int64) int, given []int64, expected [][]int64) {
		actual := SliceInt64{given}.ChunkByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int { return int((t % 2)) }

	f(remainder, []int64{1}, [][]int64{{1}})
	f(remainder, []int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f(remainder, []int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt64Int8(t *testing.T) {
	f := func(mapper func(t int64) int8, given []int64, expected [][]int64) {
		actual := SliceInt64{given}.ChunkByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int8 { return int8((t % 2)) }

	f(remainder, []int64{1}, [][]int64{{1}})
	f(remainder, []int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f(remainder, []int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt64Int16(t *testing.T) {
	f := func(mapper func(t int64) int16, given []int64, expected [][]int64) {
		actual := SliceInt64{given}.ChunkByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int16 { return int16((t % 2)) }

	f(remainder, []int64{1}, [][]int64{{1}})
	f(remainder, []int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f(remainder, []int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt64Int32(t *testing.T) {
	f := func(mapper func(t int64) int32, given []int64, expected [][]int64) {
		actual := SliceInt64{given}.ChunkByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int32 { return int32((t % 2)) }

	f(remainder, []int64{1}, [][]int64{{1}})
	f(remainder, []int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f(remainder, []int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestChunkByInt64Int64(t *testing.T) {
	f := func(mapper func(t int64) int64, given []int64, expected [][]int64) {
		actual := SliceInt64{given}.ChunkByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int64 { return int64((t % 2)) }

	f(remainder, []int64{1}, [][]int64{{1}})
	f(remainder, []int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f(remainder, []int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestChunkEveryInt64(t *testing.T) {
	f := func(count int, given []int64, expected [][]int64) {
		actual := SliceInt64{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int64{1, 2, 3, 4}, [][]int64{[]int64{1, 2}, []int64{3, 4}})
	f(2, []int64{1, 2, 3, 4, 5}, [][]int64{[]int64{1, 2}, []int64{3, 4}, []int64{5}})
}

func TestFilterInt64(t *testing.T) {
	f := func(filter func(t int64) bool, given []int64, expected []int64) {
		actual := SliceInt64{given}.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t int64) bool { return t > 0 }

	f(filterPositive, []int64{1, -1, 2, -2, 3, -3}, []int64{1, 2, 3})
	f(filterPositive, []int64{1, 2, 3}, []int64{1, 2, 3})
	f(filterPositive, []int64{-1, -2, -3}, []int64{})
}

func TestGroupByInt64Int(t *testing.T) {
	f := func(mapper func(t int64) int, given []int64, expected map[int][]int64) {
		actual := SliceInt64{given}.GroupByInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int { return int((t % 2)) }

	f(remainder, []int64{}, map[int][]int64{})
	f(remainder, []int64{1}, map[int][]int64{1: {1}})
	f(remainder, []int64{1, 3, 2, 4, 5}, map[int][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt64Int8(t *testing.T) {
	f := func(mapper func(t int64) int8, given []int64, expected map[int8][]int64) {
		actual := SliceInt64{given}.GroupByInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int8 { return int8((t % 2)) }

	f(remainder, []int64{}, map[int8][]int64{})
	f(remainder, []int64{1}, map[int8][]int64{1: {1}})
	f(remainder, []int64{1, 3, 2, 4, 5}, map[int8][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt64Int16(t *testing.T) {
	f := func(mapper func(t int64) int16, given []int64, expected map[int16][]int64) {
		actual := SliceInt64{given}.GroupByInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int16 { return int16((t % 2)) }

	f(remainder, []int64{}, map[int16][]int64{})
	f(remainder, []int64{1}, map[int16][]int64{1: {1}})
	f(remainder, []int64{1, 3, 2, 4, 5}, map[int16][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt64Int32(t *testing.T) {
	f := func(mapper func(t int64) int32, given []int64, expected map[int32][]int64) {
		actual := SliceInt64{given}.GroupByInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int32 { return int32((t % 2)) }

	f(remainder, []int64{}, map[int32][]int64{})
	f(remainder, []int64{1}, map[int32][]int64{1: {1}})
	f(remainder, []int64{1, 3, 2, 4, 5}, map[int32][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestGroupByInt64Int64(t *testing.T) {
	f := func(mapper func(t int64) int64, given []int64, expected map[int64][]int64) {
		actual := SliceInt64{given}.GroupByInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	remainder := func(t int64) int64 { return int64((t % 2)) }

	f(remainder, []int64{}, map[int64][]int64{})
	f(remainder, []int64{1}, map[int64][]int64{1: {1}})
	f(remainder, []int64{1, 3, 2, 4, 5}, map[int64][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestIntersperseInt64(t *testing.T) {
	f := func(el int64, given []int64, expected []int64) {
		actual := SliceInt64{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int64{1, 2, 3}, []int64{1, 0, 2, 0, 3})
}

func TestMapInt64Int(t *testing.T) {
	f := func(mapper func(t int64) int, given []int64, expected []int) {
		actual := SliceInt64{given}.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int { return int((t * 2)) }

	f(double, []int64{1, 2, 3}, []int{2, 4, 6})
}

func TestMapInt64Int8(t *testing.T) {
	f := func(mapper func(t int64) int8, given []int64, expected []int8) {
		actual := SliceInt64{given}.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int8 { return int8((t * 2)) }

	f(double, []int64{1, 2, 3}, []int8{2, 4, 6})
}

func TestMapInt64Int16(t *testing.T) {
	f := func(mapper func(t int64) int16, given []int64, expected []int16) {
		actual := SliceInt64{given}.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int16 { return int16((t * 2)) }

	f(double, []int64{1, 2, 3}, []int16{2, 4, 6})
}

func TestMapInt64Int32(t *testing.T) {
	f := func(mapper func(t int64) int32, given []int64, expected []int32) {
		actual := SliceInt64{given}.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int32 { return int32((t * 2)) }

	f(double, []int64{1, 2, 3}, []int32{2, 4, 6})
}

func TestMapInt64Int64(t *testing.T) {
	f := func(mapper func(t int64) int64, given []int64, expected []int64) {
		actual := SliceInt64{given}.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int64 { return int64((t * 2)) }

	f(double, []int64{1, 2, 3}, []int64{2, 4, 6})
}
