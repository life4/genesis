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

// Each calls f for every element from arr
func EachBool(arr []bool, f func(el bool)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryBool(arr []bool, count int) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatBool(arrs ...[]bool) []bool {
	result := make([]bool, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupBool(arr []bool) []bool {
	result := make([]bool, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachByte(arr []byte, f func(el byte)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryByte(arr []byte, count int) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatByte(arrs ...[]byte) []byte {
	result := make([]byte, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupByte(arr []byte) []byte {
	result := make([]byte, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachFloat32(arr []float32, f func(el float32)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryFloat32(arr []float32, count int) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatFloat32(arrs ...[]float32) []float32 {
	result := make([]float32, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupFloat32(arr []float32) []float32 {
	result := make([]float32, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachFloat64(arr []float64, f func(el float64)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryFloat64(arr []float64, count int) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatFloat64(arrs ...[]float64) []float64 {
	result := make([]float64, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupFloat64(arr []float64) []float64 {
	result := make([]float64, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachInt(arr []int, f func(el int)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryInt(arr []int, count int) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatInt(arrs ...[]int) []int {
	result := make([]int, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupInt(arr []int) []int {
	result := make([]int, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachInt16(arr []int16, f func(el int16)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryInt16(arr []int16, count int) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatInt16(arrs ...[]int16) []int16 {
	result := make([]int16, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupInt16(arr []int16) []int16 {
	result := make([]int16, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachInt32(arr []int32, f func(el int32)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryInt32(arr []int32, count int) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatInt32(arrs ...[]int32) []int32 {
	result := make([]int32, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupInt32(arr []int32) []int32 {
	result := make([]int32, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachInt64(arr []int64, f func(el int64)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryInt64(arr []int64, count int) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatInt64(arrs ...[]int64) []int64 {
	result := make([]int64, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupInt64(arr []int64) []int64 {
	result := make([]int64, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachInt8(arr []int8, f func(el int8)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryInt8(arr []int8, count int) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatInt8(arrs ...[]int8) []int8 {
	result := make([]int8, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupInt8(arr []int8) []int8 {
	result := make([]int8, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachString(arr []string, f func(el string)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryString(arr []string, count int) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatString(arrs ...[]string) []string {
	result := make([]string, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupString(arr []string) []string {
	result := make([]string, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachUint(arr []uint, f func(el uint)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryUint(arr []uint, count int) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatUint(arrs ...[]uint) []uint {
	result := make([]uint, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupUint(arr []uint) []uint {
	result := make([]uint, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachUint16(arr []uint16, f func(el uint16)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryUint16(arr []uint16, count int) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatUint16(arrs ...[]uint16) []uint16 {
	result := make([]uint16, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupUint16(arr []uint16) []uint16 {
	result := make([]uint16, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachUint32(arr []uint32, f func(el uint32)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryUint32(arr []uint32, count int) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatUint32(arrs ...[]uint32) []uint32 {
	result := make([]uint32, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupUint32(arr []uint32) []uint32 {
	result := make([]uint32, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachUint64(arr []uint64, f func(el uint64)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryUint64(arr []uint64, count int) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatUint64(arrs ...[]uint64) []uint64 {
	result := make([]uint64, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupUint64(arr []uint64) []uint64 {
	result := make([]uint64, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
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

// Each calls f for every element from arr
func EachUint8(arr []uint8, f func(el uint8)) {
	for _, el := range arr {
		f(el)
	}
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

// ChunkEvery returns slice of slices containing count elements each
func ChunkEveryUint8(arr []uint8, count int) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func ConcatUint8(arrs ...[]uint8) []uint8 {
	result := make([]uint8, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Dedup returns a given slice without consecutive duplicated elements
func DedupUint8(arr []uint8) []uint8 {
	result := make([]uint8, 0, len(arr))

	prev := arr[0]
	result = append(result, prev)
	for _, el := range arr[1:] {
		if el != prev {
			result = append(result, el)
		}
		prev = el
	}
	return result
}
