package implementation

// T is a generic type
type T struct{}

// Filter returns slice of T for which F returned true
func Filter(f func(t T) bool, t []T) []T {
	result := make([]T, 0, len(t))
	for _, e := range t {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}
