package implementation

// T is a generic type
type T struct{}

// Filter returns slice of T for which F returned true
func Filter(f func(el T) bool, arr []T) []T {
	result := make([]T, 0, len(arr))
	for _, e := range arr {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func Map(f func(el T) T, arr []T) []T {
	result := make([]T, 0, len(arr))
	for _, e := range arr {
		result = append(result, f(e))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func Reduce(f func(el T, acc T) T, arr []T, acc T) T {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}
