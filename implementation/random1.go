package implementation

import (
	"math/rand"
	"time"
)

// Shuffle in random order arr elements
func Shuffle(arr []T) []T {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}
	rand.Shuffle(len(arr), swap)
	return arr
}
