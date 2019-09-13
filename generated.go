package genesis

// Filter returns slice of T for which F returned true
func FilterBool(arr []bool, f func(el bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBool(arr []bool, f func(el bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBool(arr []bool, acc bool, f func(el bool, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityBool(t bool) bool { return t }

// Any returns true if f returns true for any element in arr
func AnyBool(arr []bool, f func(el bool) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllBool(arr []bool, f func(el bool) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterByte(arr []byte, f func(el byte) bool) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByte(arr []byte, f func(el byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByte(arr []byte, acc byte, f func(el byte, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityByte(t byte) byte { return t }

// Any returns true if f returns true for any element in arr
func AnyByte(arr []byte, f func(el byte) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllByte(arr []byte, f func(el byte) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterFloat32(arr []float32, f func(el float32) bool) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32(arr []float32, f func(el float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32(arr []float32, acc float32, f func(el float32, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityFloat32(t float32) float32 { return t }

// Any returns true if f returns true for any element in arr
func AnyFloat32(arr []float32, f func(el float32) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllFloat32(arr []float32, f func(el float32) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterFloat64(arr []float64, f func(el float64) bool) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64(arr []float64, f func(el float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64(arr []float64, acc float64, f func(el float64, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityFloat64(t float64) float64 { return t }

// Any returns true if f returns true for any element in arr
func AnyFloat64(arr []float64, f func(el float64) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllFloat64(arr []float64, f func(el float64) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterInt(arr []int, f func(el int) bool) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt(arr []int, f func(el int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt(arr []int, acc int, f func(el int, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityInt(t int) int { return t }

// Any returns true if f returns true for any element in arr
func AnyInt(arr []int, f func(el int) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllInt(arr []int, f func(el int) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterInt16(arr []int16, f func(el int16) bool) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16(arr []int16, f func(el int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16(arr []int16, acc int16, f func(el int16, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityInt16(t int16) int16 { return t }

// Any returns true if f returns true for any element in arr
func AnyInt16(arr []int16, f func(el int16) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllInt16(arr []int16, f func(el int16) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterInt32(arr []int32, f func(el int32) bool) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32(arr []int32, f func(el int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32(arr []int32, acc int32, f func(el int32, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityInt32(t int32) int32 { return t }

// Any returns true if f returns true for any element in arr
func AnyInt32(arr []int32, f func(el int32) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllInt32(arr []int32, f func(el int32) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterInt64(arr []int64, f func(el int64) bool) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64(arr []int64, f func(el int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64(arr []int64, acc int64, f func(el int64, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityInt64(t int64) int64 { return t }

// Any returns true if f returns true for any element in arr
func AnyInt64(arr []int64, f func(el int64) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllInt64(arr []int64, f func(el int64) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterInt8(arr []int8, f func(el int8) bool) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8(arr []int8, f func(el int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8(arr []int8, acc int8, f func(el int8, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityInt8(t int8) int8 { return t }

// Any returns true if f returns true for any element in arr
func AnyInt8(arr []int8, f func(el int8) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllInt8(arr []int8, f func(el int8) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterString(arr []string, f func(el string) bool) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapString(arr []string, f func(el string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceString(arr []string, acc string, f func(el string, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityString(t string) string { return t }

// Any returns true if f returns true for any element in arr
func AnyString(arr []string, f func(el string) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllString(arr []string, f func(el string) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterUint(arr []uint, f func(el uint) bool) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint(arr []uint, f func(el uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint(arr []uint, acc uint, f func(el uint, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityUint(t uint) uint { return t }

// Any returns true if f returns true for any element in arr
func AnyUint(arr []uint, f func(el uint) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllUint(arr []uint, f func(el uint) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterUint16(arr []uint16, f func(el uint16) bool) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16(arr []uint16, f func(el uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16(arr []uint16, acc uint16, f func(el uint16, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityUint16(t uint16) uint16 { return t }

// Any returns true if f returns true for any element in arr
func AnyUint16(arr []uint16, f func(el uint16) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllUint16(arr []uint16, f func(el uint16) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterUint32(arr []uint32, f func(el uint32) bool) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32(arr []uint32, f func(el uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32(arr []uint32, acc uint32, f func(el uint32, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityUint32(t uint32) uint32 { return t }

// Any returns true if f returns true for any element in arr
func AnyUint32(arr []uint32, f func(el uint32) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllUint32(arr []uint32, f func(el uint32) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterUint64(arr []uint64, f func(el uint64) bool) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64(arr []uint64, f func(el uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64(arr []uint64, acc uint64, f func(el uint64, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityUint64(t uint64) uint64 { return t }

// Any returns true if f returns true for any element in arr
func AnyUint64(arr []uint64, f func(el uint64) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllUint64(arr []uint64, f func(el uint64) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterUint8(arr []uint8, f func(el uint8) bool) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8(arr []uint8, f func(el uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8(arr []uint8, acc uint8, f func(el uint8, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityUint8(t uint8) uint8 { return t }

// Any returns true if f returns true for any element in arr
func AnyUint8(arr []uint8, f func(el uint8) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllUint8(arr []uint8, f func(el uint8) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}
