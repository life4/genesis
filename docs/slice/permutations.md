# Slice.permutations

permutations is a core implementation for Permutations

## Source

```go
// permutations is a core implementation for Permutations
func (s Slice) permutations(c chan []T, size int, left []T, right []T) {
	if len(left) == size {
		c <- left
		return
	}

	for i, el := range right {
		newLeft := make([]T, 0, len(left)+1)
		newLeft = append(newLeft, left...)
		newLeft = append(newLeft, el)

		newRight := make([]T, 0, len(right)-1)
		for j, other := range right {
			if j != i {
				newRight = append(newRight, other)
			}
		}
		s.permutations(c, size, newLeft, newRight)
	}

	// close channel in the first function call after all
	if len(right) == len(s.data) {
		close(c)
	}
}
```