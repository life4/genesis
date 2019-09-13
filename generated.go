package genesis

// Filter returns slice of T for which F returned true
func FilterBool(f func(el bool) bool, arr []bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, e := range arr {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBool(f func(el bool) bool, arr []bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, e := range arr {
		result = append(result, f(e))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBool(f func(el bool, acc bool) bool, arr []bool, acc bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Filter returns slice of T for which F returned true
func FilterFloat32(f func(el float32) bool, arr []float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, e := range arr {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32(f func(el float32) float32, arr []float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, e := range arr {
		result = append(result, f(e))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32(f func(el float32, acc float32) float32, arr []float32, acc float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Filter returns slice of T for which F returned true
func FilterFloat64(f func(el float64) bool, arr []float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, e := range arr {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64(f func(el float64) float64, arr []float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, e := range arr {
		result = append(result, f(e))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64(f func(el float64, acc float64) float64, arr []float64, acc float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Filter returns slice of T for which F returned true
func FilterInt32(f func(el int32) bool, arr []int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, e := range arr {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32(f func(el int32) int32, arr []int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, e := range arr {
		result = append(result, f(e))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32(f func(el int32, acc int32) int32, arr []int32, acc int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Filter returns slice of T for which F returned true
func FilterInt64(f func(el int64) bool, arr []int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, e := range arr {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64(f func(el int64) int64, arr []int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, e := range arr {
		result = append(result, f(e))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64(f func(el int64, acc int64) int64, arr []int64, acc int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Filter returns slice of T for which F returned true
func FilterString(f func(el string) bool, arr []string) []string {
	result := make([]string, 0, len(arr))
	for _, e := range arr {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapString(f func(el string) string, arr []string) []string {
	result := make([]string, 0, len(arr))
	for _, e := range arr {
		result = append(result, f(e))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceString(f func(el string, acc string) string, arr []string, acc string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}
