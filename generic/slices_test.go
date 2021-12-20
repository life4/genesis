package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicesConcat(t *testing.T) {
	f := func(given [][]int, expected []int) {
		actual := Slices[int]{given}.Concat()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]int{}, []int{})
	f([][]int{{}}, []int{})
	f([][]int{{1}}, []int{1})
	f([][]int{{1}, {}}, []int{1})
	f([][]int{{}, {1}}, []int{1})
	f([][]int{{1, 2}, {3, 4, 5}}, []int{1, 2, 3, 4, 5})
}

func TestSlicesProduct(t *testing.T) {
	f := func(given [][]int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		s := Slices[int]{given}
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

func TestSlicesZip(t *testing.T) {
	f := func(given [][]int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		s := Slices[int]{given}
		for el := range s.Zip() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]int{}, [][]int{})
	f([][]int{{1}, {2}}, [][]int{{1, 2}})
	f([][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {2, 4}})
	f([][]int{{1, 2, 3}, {4, 5}}, [][]int{{1, 4}, {2, 5}})
}
