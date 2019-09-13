package genesis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterInt32(t *testing.T) {
	f := func(filter func(t int32) bool, given []int32, expected []int32) {
		actual := FilterInt32(filter, given)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	filterPositive := func(t int32) bool { return t > 0 }

	f(filterPositive, []int32{1, -1, 2, -2, 3, -3}, []int32{1, 2, 3})
	f(filterPositive, []int32{1, 2, 3}, []int32{1, 2, 3})
	f(filterPositive, []int32{-1, -2, -3}, []int32{})
}
