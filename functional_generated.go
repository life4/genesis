package genesis

// Filter returns slice of T for which F returned true
func FilterInt32(f func(t int32) bool, t []int32) []int32 {
	result := make([]int32, 0, len(t))
	for _, e := range t {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}
