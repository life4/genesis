package implementation

// T is a generic type
type T struct{}

// Filter returns slice of T for which F returned true
func Filter(f func(el T) bool, arr []T) []T {
	result := make([]T, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func Map(f func(el T) T, arr []T) []T {
	result := make([]T, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
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

// Identity accepts one argument and returns it
func Identity(t T) T { return t }

// Any returns true if f returns true for any element in arr
func Any(f func(el T) bool, arr []T) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func All(f func(el T) bool, arr []T) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}
