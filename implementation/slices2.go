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

// ReduceWhile2 is like Reduce, but stops when f returns error
func ReduceWhile2(arr []T, acc G, f func(el T, acc G) (G, error)) (G, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan2 is like Reduce2, but returns slice of f results
func Scan2(arr []T, acc G, f func(el T, acc G) G) []G {
	result := make([]G, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}
