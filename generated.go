package genesis

// Filter returns slice of T for which F returned true
func FilterBool(f func(el bool) bool, arr []bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBool(f func(el bool) bool, arr []bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
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

// Identity accepts one argument and returns it
func IdentityBool(t bool) bool { return t }

// Any returns true if f returns true for any element in arr
func AnyBool(f func(el bool) bool, arr []bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllBool(f func(el bool) bool, arr []bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterByte(f func(el byte) bool, arr []byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByte(f func(el byte) byte, arr []byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByte(f func(el byte, acc byte) byte, arr []byte, acc byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityByte(t byte) byte { return t }

// Any returns true if f returns true for any element in arr
func AnyByte(f func(el byte) bool, arr []byte) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllByte(f func(el byte) bool, arr []byte) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterFloat32(f func(el float32) bool, arr []float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32(f func(el float32) float32, arr []float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
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

// Identity accepts one argument and returns it
func IdentityFloat32(t float32) float32 { return t }

// Any returns true if f returns true for any element in arr
func AnyFloat32(f func(el float32) bool, arr []float32) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllFloat32(f func(el float32) bool, arr []float32) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterFloat64(f func(el float64) bool, arr []float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64(f func(el float64) float64, arr []float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
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

// Identity accepts one argument and returns it
func IdentityFloat64(t float64) float64 { return t }

// Any returns true if f returns true for any element in arr
func AnyFloat64(f func(el float64) bool, arr []float64) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllFloat64(f func(el float64) bool, arr []float64) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterInt(f func(el int) bool, arr []int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt(f func(el int) int, arr []int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt(f func(el int, acc int) int, arr []int, acc int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityInt(t int) int { return t }

// Any returns true if f returns true for any element in arr
func AnyInt(f func(el int) bool, arr []int) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllInt(f func(el int) bool, arr []int) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterInt16(f func(el int16) bool, arr []int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16(f func(el int16) int16, arr []int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16(f func(el int16, acc int16) int16, arr []int16, acc int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityInt16(t int16) int16 { return t }

// Any returns true if f returns true for any element in arr
func AnyInt16(f func(el int16) bool, arr []int16) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllInt16(f func(el int16) bool, arr []int16) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterInt32(f func(el int32) bool, arr []int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32(f func(el int32) int32, arr []int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
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

// Identity accepts one argument and returns it
func IdentityInt32(t int32) int32 { return t }

// Any returns true if f returns true for any element in arr
func AnyInt32(f func(el int32) bool, arr []int32) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllInt32(f func(el int32) bool, arr []int32) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterInt64(f func(el int64) bool, arr []int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64(f func(el int64) int64, arr []int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
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

// Identity accepts one argument and returns it
func IdentityInt64(t int64) int64 { return t }

// Any returns true if f returns true for any element in arr
func AnyInt64(f func(el int64) bool, arr []int64) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllInt64(f func(el int64) bool, arr []int64) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterInt8(f func(el int8) bool, arr []int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8(f func(el int8) int8, arr []int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8(f func(el int8, acc int8) int8, arr []int8, acc int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityInt8(t int8) int8 { return t }

// Any returns true if f returns true for any element in arr
func AnyInt8(f func(el int8) bool, arr []int8) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllInt8(f func(el int8) bool, arr []int8) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterString(f func(el string) bool, arr []string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapString(f func(el string) string, arr []string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
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

// Identity accepts one argument and returns it
func IdentityString(t string) string { return t }

// Any returns true if f returns true for any element in arr
func AnyString(f func(el string) bool, arr []string) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllString(f func(el string) bool, arr []string) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterUint(f func(el uint) bool, arr []uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint(f func(el uint) uint, arr []uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint(f func(el uint, acc uint) uint, arr []uint, acc uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityUint(t uint) uint { return t }

// Any returns true if f returns true for any element in arr
func AnyUint(f func(el uint) bool, arr []uint) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllUint(f func(el uint) bool, arr []uint) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterUint16(f func(el uint16) bool, arr []uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16(f func(el uint16) uint16, arr []uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16(f func(el uint16, acc uint16) uint16, arr []uint16, acc uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityUint16(t uint16) uint16 { return t }

// Any returns true if f returns true for any element in arr
func AnyUint16(f func(el uint16) bool, arr []uint16) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllUint16(f func(el uint16) bool, arr []uint16) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterUint32(f func(el uint32) bool, arr []uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32(f func(el uint32) uint32, arr []uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32(f func(el uint32, acc uint32) uint32, arr []uint32, acc uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityUint32(t uint32) uint32 { return t }

// Any returns true if f returns true for any element in arr
func AnyUint32(f func(el uint32) bool, arr []uint32) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllUint32(f func(el uint32) bool, arr []uint32) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterUint64(f func(el uint64) bool, arr []uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64(f func(el uint64) uint64, arr []uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64(f func(el uint64, acc uint64) uint64, arr []uint64, acc uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityUint64(t uint64) uint64 { return t }

// Any returns true if f returns true for any element in arr
func AnyUint64(f func(el uint64) bool, arr []uint64) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllUint64(f func(el uint64) bool, arr []uint64) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// Filter returns slice of T for which F returned true
func FilterUint8(f func(el uint8) bool, arr []uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8(f func(el uint8) uint8, arr []uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8(f func(el uint8, acc uint8) uint8, arr []uint8, acc uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// Identity accepts one argument and returns it
func IdentityUint8(t uint8) uint8 { return t }

// Any returns true if f returns true for any element in arr
func AnyUint8(f func(el uint8) bool, arr []uint8) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func AllUint8(f func(el uint8) bool, arr []uint8) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}
