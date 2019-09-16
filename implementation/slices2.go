package implementation

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2(arr []T, f func(el T) G) []G {
	result := make([]G, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}
