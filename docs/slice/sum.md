# Slice.Sum

Sum return sum of all elements from arr

## Source

```go
// Sum return sum of all elements from arr
func (s Slice) Sum() T {
	var sum T
	for _, el := range s.data {
		sum += el
	}
	return sum
}
```