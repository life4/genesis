package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicesConcat(t *testing.T) {
	f := func(given [][]T, expected []T) {
		actual := Slices{given}.Concat()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]T{}, []T{})
	f([][]T{{}}, []T{})
	f([][]T{{1}}, []T{1})
	f([][]T{{1}, {}}, []T{1})
	f([][]T{{}, {1}}, []T{1})
	f([][]T{{1, 2}, {3, 4, 5}}, []T{1, 2, 3, 4, 5})
}

func TestSlicesProduct(t *testing.T) {
	f := func(given [][]T, expected [][]T) {
		actual := make([][]T, 0)
		i := 0
		s := Slices{given}
		for el := range s.Product() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]T{{1, 2}, {3, 4}}, [][]T{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]T{{1, 2}, {3}, {4, 5}}, [][]T{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}

func TestSlicesZip(t *testing.T) {
	f := func(given [][]T, expected [][]T) {
		actual := make([][]T, 0)
		i := 0
		s := Slices{given}
		for el := range s.Zip() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]T{}, [][]T{})
	f([][]T{{1}, {2}}, [][]T{{1, 2}})
	f([][]T{{1, 2}, {3, 4}}, [][]T{{1, 3}, {2, 4}})
	f([][]T{{1, 2, 3}, {4, 5}}, [][]T{{1, 4}, {2, 5}})
}
