package implementation

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2(arr []T, f func(el T) G) []G {
	result := make([]G, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce2 applies F to acc and every element in slice of T and returns acc
func Reduce2(arr []T, acc G, f func(el T, acc G) G) G {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}
