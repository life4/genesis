package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
