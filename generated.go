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

// Contains returns true if el in arr.
func ContainsBool(arr []bool, el bool) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindBool(arr []bool, def bool, f func(el bool) bool) bool {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexBool(arr []bool, f func(el bool) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseBool(arr []bool, el bool) []bool {
	result := make([]bool, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
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

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBool(arr []bool, acc bool, f func(el bool, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanBool(arr []bool, acc bool, f func(el bool, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithBool(arr []bool, prefix []bool) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeBool(arr []bool, f func(el bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsByte(arr []byte, el byte) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindByte(arr []byte, def byte, f func(el byte) bool) byte {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexByte(arr []byte, f func(el byte) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseByte(arr []byte, el byte) []byte {
	result := make([]byte, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxByte(arr []byte) byte {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinByte(arr []byte) byte {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByte(arr []byte, acc byte, f func(el byte, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByte(arr []byte, acc byte, f func(el byte, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanByte(arr []byte, acc byte, f func(el byte, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithByte(arr []byte, prefix []byte) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeByte(arr []byte, f func(el byte) bool) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsFloat32(arr []float32, el float32) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindFloat32(arr []float32, def float32, f func(el float32) bool) float32 {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexFloat32(arr []float32, f func(el float32) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseFloat32(arr []float32, el float32) []float32 {
	result := make([]float32, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxFloat32(arr []float32) float32 {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinFloat32(arr []float32) float32 {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32(arr []float32, acc float32, f func(el float32, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32(arr []float32, acc float32, f func(el float32, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanFloat32(arr []float32, acc float32, f func(el float32, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithFloat32(arr []float32, prefix []float32) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeFloat32(arr []float32, f func(el float32) bool) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsFloat64(arr []float64, el float64) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindFloat64(arr []float64, def float64, f func(el float64) bool) float64 {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexFloat64(arr []float64, f func(el float64) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseFloat64(arr []float64, el float64) []float64 {
	result := make([]float64, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxFloat64(arr []float64) float64 {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinFloat64(arr []float64) float64 {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64(arr []float64, acc float64, f func(el float64, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64(arr []float64, acc float64, f func(el float64, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanFloat64(arr []float64, acc float64, f func(el float64, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithFloat64(arr []float64, prefix []float64) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeFloat64(arr []float64, f func(el float64) bool) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsInt(arr []int, el int) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindInt(arr []int, def int, f func(el int) bool) int {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexInt(arr []int, f func(el int) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseInt(arr []int, el int) []int {
	result := make([]int, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxInt(arr []int) int {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinInt(arr []int) int {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt(arr []int, acc int, f func(el int, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt(arr []int, acc int, f func(el int, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanInt(arr []int, acc int, f func(el int, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithInt(arr []int, prefix []int) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeInt(arr []int, f func(el int) bool) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsInt16(arr []int16, el int16) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindInt16(arr []int16, def int16, f func(el int16) bool) int16 {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexInt16(arr []int16, f func(el int16) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseInt16(arr []int16, el int16) []int16 {
	result := make([]int16, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxInt16(arr []int16) int16 {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinInt16(arr []int16) int16 {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16(arr []int16, acc int16, f func(el int16, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16(arr []int16, acc int16, f func(el int16, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanInt16(arr []int16, acc int16, f func(el int16, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithInt16(arr []int16, prefix []int16) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeInt16(arr []int16, f func(el int16) bool) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsInt32(arr []int32, el int32) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindInt32(arr []int32, def int32, f func(el int32) bool) int32 {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexInt32(arr []int32, f func(el int32) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseInt32(arr []int32, el int32) []int32 {
	result := make([]int32, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxInt32(arr []int32) int32 {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinInt32(arr []int32) int32 {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32(arr []int32, acc int32, f func(el int32, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32(arr []int32, acc int32, f func(el int32, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanInt32(arr []int32, acc int32, f func(el int32, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithInt32(arr []int32, prefix []int32) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeInt32(arr []int32, f func(el int32) bool) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsInt64(arr []int64, el int64) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindInt64(arr []int64, def int64, f func(el int64) bool) int64 {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexInt64(arr []int64, f func(el int64) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseInt64(arr []int64, el int64) []int64 {
	result := make([]int64, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxInt64(arr []int64) int64 {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinInt64(arr []int64) int64 {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64(arr []int64, acc int64, f func(el int64, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64(arr []int64, acc int64, f func(el int64, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanInt64(arr []int64, acc int64, f func(el int64, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithInt64(arr []int64, prefix []int64) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeInt64(arr []int64, f func(el int64) bool) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsInt8(arr []int8, el int8) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindInt8(arr []int8, def int8, f func(el int8) bool) int8 {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexInt8(arr []int8, f func(el int8) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseInt8(arr []int8, el int8) []int8 {
	result := make([]int8, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxInt8(arr []int8) int8 {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinInt8(arr []int8) int8 {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8(arr []int8, acc int8, f func(el int8, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8(arr []int8, acc int8, f func(el int8, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanInt8(arr []int8, acc int8, f func(el int8, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithInt8(arr []int8, prefix []int8) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeInt8(arr []int8, f func(el int8) bool) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Filter returns slice of T for which F returned true
func Filter(arr []interface{}, f func(el interface{}) bool) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func Map(arr []interface{}, f func(el interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Each calls f for every element from arr
func Each(arr []interface{}, f func(el interface{})) {
	for _, el := range arr {
		f(el)
	}
}

// Identity accepts one argument and returns it
func Identity(t interface{}) interface{} { return t }

// Any returns true if f returns true for any element in arr
func Any(arr []interface{}, f func(el interface{}) bool) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements in arr
func All(arr []interface{}, f func(el interface{}) bool) bool {
	for _, el := range arr {
		if !f(el) {
			return false
		}
	}
	return true
}

// ChunkEvery returns slice of slices containing count elements each
func ChunkEvery(arr []interface{}, count int) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)
	for i, el := range arr {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// Concat concatenates given slices into a single slice.
func Concat(arrs ...[]interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	return result
}

// Contains returns true if el in arr.
func Contains(arr []interface{}, el interface{}) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
}

// Dedup returns a given slice without consecutive duplicated elements
func Dedup(arr []interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))

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

// Find returns the first element for which f returns true
func Find(arr []interface{}, def interface{}, f func(el interface{}) bool) interface{} {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndex(arr []interface{}, f func(el interface{}) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func Intersperse(arr []interface{}, el interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func Reduce(arr []interface{}, acc interface{}, f func(el interface{}, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhile(arr []interface{}, acc interface{}, f func(el interface{}, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func Scan(arr []interface{}, acc interface{}, f func(el interface{}, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWith(arr []interface{}, prefix []interface{}) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func Take(arr []interface{}, f func(el interface{}) bool) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsString(arr []string, el string) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindString(arr []string, def string, f func(el string) bool) string {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexString(arr []string, f func(el string) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseString(arr []string, el string) []string {
	result := make([]string, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxString(arr []string) string {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinString(arr []string) string {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceString(arr []string, acc string, f func(el string, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileString(arr []string, acc string, f func(el string, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanString(arr []string, acc string, f func(el string, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithString(arr []string, prefix []string) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeString(arr []string, f func(el string) bool) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsUint(arr []uint, el uint) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindUint(arr []uint, def uint, f func(el uint) bool) uint {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexUint(arr []uint, f func(el uint) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseUint(arr []uint, el uint) []uint {
	result := make([]uint, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxUint(arr []uint) uint {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinUint(arr []uint) uint {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint(arr []uint, acc uint, f func(el uint, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint(arr []uint, acc uint, f func(el uint, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanUint(arr []uint, acc uint, f func(el uint, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithUint(arr []uint, prefix []uint) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeUint(arr []uint, f func(el uint) bool) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsUint16(arr []uint16, el uint16) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindUint16(arr []uint16, def uint16, f func(el uint16) bool) uint16 {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexUint16(arr []uint16, f func(el uint16) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseUint16(arr []uint16, el uint16) []uint16 {
	result := make([]uint16, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxUint16(arr []uint16) uint16 {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinUint16(arr []uint16) uint16 {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16(arr []uint16, acc uint16, f func(el uint16, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16(arr []uint16, acc uint16, f func(el uint16, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanUint16(arr []uint16, acc uint16, f func(el uint16, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithUint16(arr []uint16, prefix []uint16) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeUint16(arr []uint16, f func(el uint16) bool) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsUint32(arr []uint32, el uint32) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindUint32(arr []uint32, def uint32, f func(el uint32) bool) uint32 {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexUint32(arr []uint32, f func(el uint32) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseUint32(arr []uint32, el uint32) []uint32 {
	result := make([]uint32, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxUint32(arr []uint32) uint32 {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinUint32(arr []uint32) uint32 {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32(arr []uint32, acc uint32, f func(el uint32, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32(arr []uint32, acc uint32, f func(el uint32, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanUint32(arr []uint32, acc uint32, f func(el uint32, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithUint32(arr []uint32, prefix []uint32) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeUint32(arr []uint32, f func(el uint32) bool) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsUint64(arr []uint64, el uint64) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindUint64(arr []uint64, def uint64, f func(el uint64) bool) uint64 {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexUint64(arr []uint64, f func(el uint64) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseUint64(arr []uint64, el uint64) []uint64 {
	result := make([]uint64, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxUint64(arr []uint64) uint64 {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinUint64(arr []uint64) uint64 {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64(arr []uint64, acc uint64, f func(el uint64, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64(arr []uint64, acc uint64, f func(el uint64, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanUint64(arr []uint64, acc uint64, f func(el uint64, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithUint64(arr []uint64, prefix []uint64) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeUint64(arr []uint64, f func(el uint64) bool) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Contains returns true if el in arr.
func ContainsUint8(arr []uint8, el uint8) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
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

// Find returns the first element for which f returns true
func FindUint8(arr []uint8, def uint8, f func(el uint8) bool) uint8 {
	for _, el := range arr {
		if f(el) {
			return el
		}
	}
	return def
}

// FindIndex is like Find, but return element index instead of element itself
// Returns -1 if element is not found
func FindIndexUint8(arr []uint8, f func(el uint8) bool) int {
	for i, el := range arr {
		if f(el) {
			return i
		}
	}
	return -1
}

// Intersperse inserts el between each element of arr
func IntersperseUint8(arr []uint8, el uint8) []uint8 {
	result := make([]uint8, 0, len(arr)*2-1)
	result = append(result, arr[0])
	for _, val := range arr[1:] {
		result = append(result, el, val)
	}
	return result
}

// Max returns the maximal element from arr
func MaxUint8(arr []uint8) uint8 {
	max := arr[0]
	for _, el := range arr[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

// Min returns the minimal element from arr
func MinUint8(arr []uint8) uint8 {
	min := arr[0]
	for _, el := range arr[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8(arr []uint8, acc uint8, f func(el uint8, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8(arr []uint8, acc uint8, f func(el uint8, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func ScanUint8(arr []uint8, acc uint8, f func(el uint8, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// StartsWith returns true if slice starts with the given prefix slice.
// If prefix is empty, it returns true.
func StartsWithUint8(arr []uint8, prefix []uint8) bool {
	for i, el := range prefix {
		if el != arr[i] {
			return false
		}
	}
	return true
}

// Take takes elements from arr while f returns true
func TakeUint8(arr []uint8, f func(el uint8) bool) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolBool(arr []bool, f func(el bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolByte(arr []bool, f func(el bool) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolFloat32(arr []bool, f func(el bool) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolFloat64(arr []bool, f func(el bool) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolInt(arr []bool, f func(el bool) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolInt16(arr []bool, f func(el bool) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolInt32(arr []bool, f func(el bool) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolInt64(arr []bool, f func(el bool) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolInt8(arr []bool, f func(el bool) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolInterface(arr []bool, f func(el bool) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolString(arr []bool, f func(el bool) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolUint(arr []bool, f func(el bool) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolUint16(arr []bool, f func(el bool) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolUint32(arr []bool, f func(el bool) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolUint64(arr []bool, f func(el bool) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2BoolUint8(arr []bool, f func(el bool) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteBool(arr []byte, f func(el byte) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteByte(arr []byte, f func(el byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteFloat32(arr []byte, f func(el byte) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteFloat64(arr []byte, f func(el byte) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteInt(arr []byte, f func(el byte) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteInt16(arr []byte, f func(el byte) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteInt32(arr []byte, f func(el byte) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteInt64(arr []byte, f func(el byte) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteInt8(arr []byte, f func(el byte) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteInterface(arr []byte, f func(el byte) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteString(arr []byte, f func(el byte) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteUint(arr []byte, f func(el byte) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteUint16(arr []byte, f func(el byte) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteUint32(arr []byte, f func(el byte) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteUint64(arr []byte, f func(el byte) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2ByteUint8(arr []byte, f func(el byte) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Bool(arr []float32, f func(el float32) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Byte(arr []float32, f func(el float32) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Float32(arr []float32, f func(el float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Float64(arr []float32, f func(el float32) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Int(arr []float32, f func(el float32) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Int16(arr []float32, f func(el float32) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Int32(arr []float32, f func(el float32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Int64(arr []float32, f func(el float32) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Int8(arr []float32, f func(el float32) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Interface(arr []float32, f func(el float32) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32String(arr []float32, f func(el float32) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Uint(arr []float32, f func(el float32) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Uint16(arr []float32, f func(el float32) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Uint32(arr []float32, f func(el float32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Uint64(arr []float32, f func(el float32) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float32Uint8(arr []float32, f func(el float32) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Bool(arr []float64, f func(el float64) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Byte(arr []float64, f func(el float64) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Float32(arr []float64, f func(el float64) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Float64(arr []float64, f func(el float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Int(arr []float64, f func(el float64) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Int16(arr []float64, f func(el float64) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Int32(arr []float64, f func(el float64) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Int64(arr []float64, f func(el float64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Int8(arr []float64, f func(el float64) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Interface(arr []float64, f func(el float64) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64String(arr []float64, f func(el float64) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Uint(arr []float64, f func(el float64) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Uint16(arr []float64, f func(el float64) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Uint32(arr []float64, f func(el float64) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Uint64(arr []float64, f func(el float64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Float64Uint8(arr []float64, f func(el float64) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Bool(arr []int16, f func(el int16) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Byte(arr []int16, f func(el int16) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Float32(arr []int16, f func(el int16) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Float64(arr []int16, f func(el int16) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Int(arr []int16, f func(el int16) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Int16(arr []int16, f func(el int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Int32(arr []int16, f func(el int16) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Int64(arr []int16, f func(el int16) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Int8(arr []int16, f func(el int16) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Interface(arr []int16, f func(el int16) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16String(arr []int16, f func(el int16) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Uint(arr []int16, f func(el int16) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Uint16(arr []int16, f func(el int16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Uint32(arr []int16, f func(el int16) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Uint64(arr []int16, f func(el int16) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int16Uint8(arr []int16, f func(el int16) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Bool(arr []int32, f func(el int32) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Byte(arr []int32, f func(el int32) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Float32(arr []int32, f func(el int32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Float64(arr []int32, f func(el int32) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Int(arr []int32, f func(el int32) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Int16(arr []int32, f func(el int32) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Int32(arr []int32, f func(el int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Int64(arr []int32, f func(el int32) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Int8(arr []int32, f func(el int32) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Interface(arr []int32, f func(el int32) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32String(arr []int32, f func(el int32) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Uint(arr []int32, f func(el int32) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Uint16(arr []int32, f func(el int32) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Uint32(arr []int32, f func(el int32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Uint64(arr []int32, f func(el int32) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int32Uint8(arr []int32, f func(el int32) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Bool(arr []int64, f func(el int64) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Byte(arr []int64, f func(el int64) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Float32(arr []int64, f func(el int64) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Float64(arr []int64, f func(el int64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Int(arr []int64, f func(el int64) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Int16(arr []int64, f func(el int64) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Int32(arr []int64, f func(el int64) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Int64(arr []int64, f func(el int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Int8(arr []int64, f func(el int64) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Interface(arr []int64, f func(el int64) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64String(arr []int64, f func(el int64) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Uint(arr []int64, f func(el int64) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Uint16(arr []int64, f func(el int64) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Uint32(arr []int64, f func(el int64) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Uint64(arr []int64, f func(el int64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int64Uint8(arr []int64, f func(el int64) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Bool(arr []int8, f func(el int8) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Byte(arr []int8, f func(el int8) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Float32(arr []int8, f func(el int8) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Float64(arr []int8, f func(el int8) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Int(arr []int8, f func(el int8) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Int16(arr []int8, f func(el int8) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Int32(arr []int8, f func(el int8) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Int64(arr []int8, f func(el int8) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Int8(arr []int8, f func(el int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Interface(arr []int8, f func(el int8) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8String(arr []int8, f func(el int8) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Uint(arr []int8, f func(el int8) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Uint16(arr []int8, f func(el int8) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Uint32(arr []int8, f func(el int8) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Uint64(arr []int8, f func(el int8) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Int8Uint8(arr []int8, f func(el int8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntBool(arr []int, f func(el int) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntByte(arr []int, f func(el int) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceBool(arr []interface{}, f func(el interface{}) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceByte(arr []interface{}, f func(el interface{}) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceFloat32(arr []interface{}, f func(el interface{}) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceFloat64(arr []interface{}, f func(el interface{}) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceInt(arr []interface{}, f func(el interface{}) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceInt16(arr []interface{}, f func(el interface{}) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceInt32(arr []interface{}, f func(el interface{}) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceInt64(arr []interface{}, f func(el interface{}) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceInt8(arr []interface{}, f func(el interface{}) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceInterface(arr []interface{}, f func(el interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceString(arr []interface{}, f func(el interface{}) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceUint(arr []interface{}, f func(el interface{}) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceUint16(arr []interface{}, f func(el interface{}) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceUint32(arr []interface{}, f func(el interface{}) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceUint64(arr []interface{}, f func(el interface{}) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2InterfaceUint8(arr []interface{}, f func(el interface{}) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntFloat32(arr []int, f func(el int) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntFloat64(arr []int, f func(el int) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntInt(arr []int, f func(el int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntInt16(arr []int, f func(el int) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntInt32(arr []int, f func(el int) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntInt64(arr []int, f func(el int) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntInt8(arr []int, f func(el int) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntInterface(arr []int, f func(el int) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntString(arr []int, f func(el int) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntUint(arr []int, f func(el int) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntUint16(arr []int, f func(el int) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntUint32(arr []int, f func(el int) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntUint64(arr []int, f func(el int) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2IntUint8(arr []int, f func(el int) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringBool(arr []string, f func(el string) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringByte(arr []string, f func(el string) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringFloat32(arr []string, f func(el string) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringFloat64(arr []string, f func(el string) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringInt(arr []string, f func(el string) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringInt16(arr []string, f func(el string) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringInt32(arr []string, f func(el string) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringInt64(arr []string, f func(el string) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringInt8(arr []string, f func(el string) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringInterface(arr []string, f func(el string) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringString(arr []string, f func(el string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringUint(arr []string, f func(el string) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringUint16(arr []string, f func(el string) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringUint32(arr []string, f func(el string) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringUint64(arr []string, f func(el string) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2StringUint8(arr []string, f func(el string) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Bool(arr []uint16, f func(el uint16) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Byte(arr []uint16, f func(el uint16) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Float32(arr []uint16, f func(el uint16) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Float64(arr []uint16, f func(el uint16) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Int(arr []uint16, f func(el uint16) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Int16(arr []uint16, f func(el uint16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Int32(arr []uint16, f func(el uint16) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Int64(arr []uint16, f func(el uint16) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Int8(arr []uint16, f func(el uint16) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Interface(arr []uint16, f func(el uint16) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16String(arr []uint16, f func(el uint16) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Uint(arr []uint16, f func(el uint16) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Uint16(arr []uint16, f func(el uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Uint32(arr []uint16, f func(el uint16) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Uint64(arr []uint16, f func(el uint16) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint16Uint8(arr []uint16, f func(el uint16) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Bool(arr []uint32, f func(el uint32) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Byte(arr []uint32, f func(el uint32) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Float32(arr []uint32, f func(el uint32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Float64(arr []uint32, f func(el uint32) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Int(arr []uint32, f func(el uint32) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Int16(arr []uint32, f func(el uint32) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Int32(arr []uint32, f func(el uint32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Int64(arr []uint32, f func(el uint32) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Int8(arr []uint32, f func(el uint32) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Interface(arr []uint32, f func(el uint32) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32String(arr []uint32, f func(el uint32) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Uint(arr []uint32, f func(el uint32) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Uint16(arr []uint32, f func(el uint32) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Uint32(arr []uint32, f func(el uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Uint64(arr []uint32, f func(el uint32) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint32Uint8(arr []uint32, f func(el uint32) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Bool(arr []uint64, f func(el uint64) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Byte(arr []uint64, f func(el uint64) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Float32(arr []uint64, f func(el uint64) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Float64(arr []uint64, f func(el uint64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Int(arr []uint64, f func(el uint64) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Int16(arr []uint64, f func(el uint64) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Int32(arr []uint64, f func(el uint64) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Int64(arr []uint64, f func(el uint64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Int8(arr []uint64, f func(el uint64) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Interface(arr []uint64, f func(el uint64) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64String(arr []uint64, f func(el uint64) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Uint(arr []uint64, f func(el uint64) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Uint16(arr []uint64, f func(el uint64) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Uint32(arr []uint64, f func(el uint64) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Uint64(arr []uint64, f func(el uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint64Uint8(arr []uint64, f func(el uint64) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Bool(arr []uint8, f func(el uint8) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Byte(arr []uint8, f func(el uint8) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Float32(arr []uint8, f func(el uint8) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Float64(arr []uint8, f func(el uint8) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Int(arr []uint8, f func(el uint8) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Int16(arr []uint8, f func(el uint8) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Int32(arr []uint8, f func(el uint8) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Int64(arr []uint8, f func(el uint8) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Int8(arr []uint8, f func(el uint8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Interface(arr []uint8, f func(el uint8) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8String(arr []uint8, f func(el uint8) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Uint(arr []uint8, f func(el uint8) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Uint16(arr []uint8, f func(el uint8) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Uint32(arr []uint8, f func(el uint8) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Uint64(arr []uint8, f func(el uint8) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2Uint8Uint8(arr []uint8, f func(el uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintBool(arr []uint, f func(el uint) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintByte(arr []uint, f func(el uint) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintFloat32(arr []uint, f func(el uint) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintFloat64(arr []uint, f func(el uint) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintInt(arr []uint, f func(el uint) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintInt16(arr []uint, f func(el uint) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintInt32(arr []uint, f func(el uint) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintInt64(arr []uint, f func(el uint) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintInt8(arr []uint, f func(el uint) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintInterface(arr []uint, f func(el uint) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintString(arr []uint, f func(el uint) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintUint(arr []uint, f func(el uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintUint16(arr []uint, f func(el uint) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintUint32(arr []uint, f func(el uint) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintUint64(arr []uint, f func(el uint) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Map2 applies F to all elements in slice of T and returns slice of results
func Map2UintUint8(arr []uint, f func(el uint) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}
