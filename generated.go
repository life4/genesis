package genesis

// Cycle is an infinite loop over arr
func CycleBool(arr []bool) chan bool {
	c := make(chan bool, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatBool(val bool) chan bool {
	c := make(chan bool, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeBool(c chan bool, n int) []bool {
	result := make([]bool, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllBool(c chan bool) []bool {
	result := make([]bool, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Cycle is an infinite loop over arr
func CycleByte(arr []byte) chan byte {
	c := make(chan byte, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatByte(val byte) chan byte {
	c := make(chan byte, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeByte(c chan byte, n int) []byte {
	result := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllByte(c chan byte) []byte {
	result := make([]byte, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Count is like Range, but infinite
func CountFloat32(start float32, step float32) chan float32 {
	c := make(chan float32, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func CycleFloat32(arr []float32) chan float32 {
	c := make(chan float32, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func RangeFloat32(start float32, end float32, step float32) chan float32 {
	c := make(chan float32, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatFloat32(val float32) chan float32 {
	c := make(chan float32, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeFloat32(c chan float32, n int) []float32 {
	result := make([]float32, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllFloat32(c chan float32) []float32 {
	result := make([]float32, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Count is like Range, but infinite
func CountFloat64(start float64, step float64) chan float64 {
	c := make(chan float64, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func CycleFloat64(arr []float64) chan float64 {
	c := make(chan float64, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func RangeFloat64(start float64, end float64, step float64) chan float64 {
	c := make(chan float64, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatFloat64(val float64) chan float64 {
	c := make(chan float64, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeFloat64(c chan float64, n int) []float64 {
	result := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllFloat64(c chan float64) []float64 {
	result := make([]float64, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Count is like Range, but infinite
func CountInt(start int, step int) chan int {
	c := make(chan int, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func CycleInt(arr []int) chan int {
	c := make(chan int, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func RangeInt(start int, end int, step int) chan int {
	c := make(chan int, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatInt(val int) chan int {
	c := make(chan int, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeInt(c chan int, n int) []int {
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllInt(c chan int) []int {
	result := make([]int, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Count is like Range, but infinite
func CountInt16(start int16, step int16) chan int16 {
	c := make(chan int16, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func CycleInt16(arr []int16) chan int16 {
	c := make(chan int16, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func RangeInt16(start int16, end int16, step int16) chan int16 {
	c := make(chan int16, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatInt16(val int16) chan int16 {
	c := make(chan int16, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeInt16(c chan int16, n int) []int16 {
	result := make([]int16, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllInt16(c chan int16) []int16 {
	result := make([]int16, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Count is like Range, but infinite
func CountInt32(start int32, step int32) chan int32 {
	c := make(chan int32, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func CycleInt32(arr []int32) chan int32 {
	c := make(chan int32, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func RangeInt32(start int32, end int32, step int32) chan int32 {
	c := make(chan int32, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatInt32(val int32) chan int32 {
	c := make(chan int32, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeInt32(c chan int32, n int) []int32 {
	result := make([]int32, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllInt32(c chan int32) []int32 {
	result := make([]int32, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Count is like Range, but infinite
func CountInt64(start int64, step int64) chan int64 {
	c := make(chan int64, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func CycleInt64(arr []int64) chan int64 {
	c := make(chan int64, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func RangeInt64(start int64, end int64, step int64) chan int64 {
	c := make(chan int64, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatInt64(val int64) chan int64 {
	c := make(chan int64, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeInt64(c chan int64, n int) []int64 {
	result := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllInt64(c chan int64) []int64 {
	result := make([]int64, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Count is like Range, but infinite
func CountInt8(start int8, step int8) chan int8 {
	c := make(chan int8, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func CycleInt8(arr []int8) chan int8 {
	c := make(chan int8, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func RangeInt8(start int8, end int8, step int8) chan int8 {
	c := make(chan int8, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatInt8(val int8) chan int8 {
	c := make(chan int8, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeInt8(c chan int8, n int) []int8 {
	result := make([]int8, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllInt8(c chan int8) []int8 {
	result := make([]int8, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Cycle is an infinite loop over arr
func Cycle(arr []interface{}) chan interface{} {
	c := make(chan interface{}, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func Repeat(val interface{}) chan interface{} {
	c := make(chan interface{}, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func Take(c chan interface{}, n int) []interface{} {
	result := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAll(c chan interface{}) []interface{} {
	result := make([]interface{}, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Cycle is an infinite loop over arr
func CycleString(arr []string) chan string {
	c := make(chan string, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatString(val string) chan string {
	c := make(chan string, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeString(c chan string, n int) []string {
	result := make([]string, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllString(c chan string) []string {
	result := make([]string, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Count is like Range, but infinite
func CountUint(start uint, step uint) chan uint {
	c := make(chan uint, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func CycleUint(arr []uint) chan uint {
	c := make(chan uint, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func RangeUint(start uint, end uint, step uint) chan uint {
	c := make(chan uint, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatUint(val uint) chan uint {
	c := make(chan uint, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeUint(c chan uint, n int) []uint {
	result := make([]uint, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllUint(c chan uint) []uint {
	result := make([]uint, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Count is like Range, but infinite
func CountUint16(start uint16, step uint16) chan uint16 {
	c := make(chan uint16, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func CycleUint16(arr []uint16) chan uint16 {
	c := make(chan uint16, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func RangeUint16(start uint16, end uint16, step uint16) chan uint16 {
	c := make(chan uint16, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatUint16(val uint16) chan uint16 {
	c := make(chan uint16, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeUint16(c chan uint16, n int) []uint16 {
	result := make([]uint16, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllUint16(c chan uint16) []uint16 {
	result := make([]uint16, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Count is like Range, but infinite
func CountUint32(start uint32, step uint32) chan uint32 {
	c := make(chan uint32, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func CycleUint32(arr []uint32) chan uint32 {
	c := make(chan uint32, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func RangeUint32(start uint32, end uint32, step uint32) chan uint32 {
	c := make(chan uint32, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatUint32(val uint32) chan uint32 {
	c := make(chan uint32, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeUint32(c chan uint32, n int) []uint32 {
	result := make([]uint32, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllUint32(c chan uint32) []uint32 {
	result := make([]uint32, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Count is like Range, but infinite
func CountUint64(start uint64, step uint64) chan uint64 {
	c := make(chan uint64, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func CycleUint64(arr []uint64) chan uint64 {
	c := make(chan uint64, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func RangeUint64(start uint64, end uint64, step uint64) chan uint64 {
	c := make(chan uint64, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatUint64(val uint64) chan uint64 {
	c := make(chan uint64, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeUint64(c chan uint64, n int) []uint64 {
	result := make([]uint64, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllUint64(c chan uint64) []uint64 {
	result := make([]uint64, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Count is like Range, but infinite
func CountUint8(start uint8, step uint8) chan uint8 {
	c := make(chan uint8, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

// Cycle is an infinite loop over arr
func CycleUint8(arr []uint8) chan uint8 {
	c := make(chan uint8, 1)
	go func() {
		for {
			for _, val := range arr {
				c <- val
			}
		}
	}()
	return c
}

// Range generates elements from start to end with given step
func RangeUint8(start uint8, end uint8, step uint8) chan uint8 {
	c := make(chan uint8, 1)
	pos := start <= end
	go func() {
		for pos && (start < end) || !pos && (start > end) {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Repeat returns channel that produces val infinite times
func RepeatUint8(val uint8) chan uint8 {
	c := make(chan uint8, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

// Take takes first n elements from channel c.
func TakeUint8(c chan uint8, n int) []uint8 {
	result := make([]uint8, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c)
	}
	return result
}

// TakeAll takes all elements from channel c.
func TakeAllUint8(c chan uint8) []uint8 {
	result := make([]uint8, 0)
	for val := range c {
		result = append(result, val)
	}
	return result
}

// Identity accepts one argument and returns it
func IdentityBool(t bool) bool { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionBool(fs ...func(val bool) bool) func(val bool) bool {
	return func(val bool) bool {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityByte(t byte) byte { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionByte(fs ...func(val byte) byte) func(val byte) byte {
	return func(val byte) byte {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityFloat32(t float32) float32 { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionFloat32(fs ...func(val float32) float32) func(val float32) float32 {
	return func(val float32) float32 {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityFloat64(t float64) float64 { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionFloat64(fs ...func(val float64) float64) func(val float64) float64 {
	return func(val float64) float64 {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityInt(t int) int { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionInt(fs ...func(val int) int) func(val int) int {
	return func(val int) int {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityInt16(t int16) int16 { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionInt16(fs ...func(val int16) int16) func(val int16) int16 {
	return func(val int16) int16 {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityInt32(t int32) int32 { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionInt32(fs ...func(val int32) int32) func(val int32) int32 {
	return func(val int32) int32 {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityInt64(t int64) int64 { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionInt64(fs ...func(val int64) int64) func(val int64) int64 {
	return func(val int64) int64 {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityInt8(t int8) int8 { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionInt8(fs ...func(val int8) int8) func(val int8) int8 {
	return func(val int8) int8 {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func Identity(t interface{}) interface{} { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func Composition(fs ...func(val interface{}) interface{}) func(val interface{}) interface{} {
	return func(val interface{}) interface{} {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityString(t string) string { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionString(fs ...func(val string) string) func(val string) string {
	return func(val string) string {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityUint(t uint) uint { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionUint(fs ...func(val uint) uint) func(val uint) uint {
	return func(val uint) uint {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityUint16(t uint16) uint16 { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionUint16(fs ...func(val uint16) uint16) func(val uint16) uint16 {
	return func(val uint16) uint16 {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityUint32(t uint32) uint32 { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionUint32(fs ...func(val uint32) uint32) func(val uint32) uint32 {
	return func(val uint32) uint32 {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityUint64(t uint64) uint64 { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionUint64(fs ...func(val uint64) uint64) func(val uint64) uint64 {
	return func(val uint64) uint64 {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

// Identity accepts one argument and returns it
func IdentityUint8(t uint8) uint8 { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func CompositionUint8(fs ...func(val uint8) uint8) func(val uint8) uint8 {
	return func(val uint8) uint8 {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}

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

// Each calls f for every element from arr
func EachBool(arr []bool, f func(el bool)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryBool(arr []bool, nth int) []bool {
	result := make([]bool, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileBool(arr []bool, f func(arr bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameBool(arr []bool) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileBool(arr []bool, f func(el bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowBool(arr []bool, size int) [][]bool {
	result := make([][]bool, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachByte(arr []byte, f func(el byte)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryByte(arr []byte, nth int) []byte {
	result := make([]byte, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileByte(arr []byte, f func(arr byte) bool) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameByte(arr []byte) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileByte(arr []byte, f func(el byte) bool) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowByte(arr []byte, size int) [][]byte {
	result := make([][]byte, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachFloat32(arr []float32, f func(el float32)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryFloat32(arr []float32, nth int) []float32 {
	result := make([]float32, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileFloat32(arr []float32, f func(arr float32) bool) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameFloat32(arr []float32) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileFloat32(arr []float32, f func(el float32) bool) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowFloat32(arr []float32, size int) [][]float32 {
	result := make([][]float32, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachFloat64(arr []float64, f func(el float64)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryFloat64(arr []float64, nth int) []float64 {
	result := make([]float64, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileFloat64(arr []float64, f func(arr float64) bool) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameFloat64(arr []float64) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileFloat64(arr []float64, f func(el float64) bool) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowFloat64(arr []float64, size int) [][]float64 {
	result := make([][]float64, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachInt(arr []int, f func(el int)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryInt(arr []int, nth int) []int {
	result := make([]int, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileInt(arr []int, f func(arr int) bool) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameInt(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileInt(arr []int, f func(el int) bool) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowInt(arr []int, size int) [][]int {
	result := make([][]int, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachInt16(arr []int16, f func(el int16)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryInt16(arr []int16, nth int) []int16 {
	result := make([]int16, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileInt16(arr []int16, f func(arr int16) bool) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameInt16(arr []int16) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileInt16(arr []int16, f func(el int16) bool) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowInt16(arr []int16, size int) [][]int16 {
	result := make([][]int16, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachInt32(arr []int32, f func(el int32)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryInt32(arr []int32, nth int) []int32 {
	result := make([]int32, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileInt32(arr []int32, f func(arr int32) bool) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameInt32(arr []int32) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileInt32(arr []int32, f func(el int32) bool) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowInt32(arr []int32, size int) [][]int32 {
	result := make([][]int32, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachInt64(arr []int64, f func(el int64)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryInt64(arr []int64, nth int) []int64 {
	result := make([]int64, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileInt64(arr []int64, f func(arr int64) bool) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameInt64(arr []int64) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileInt64(arr []int64, f func(el int64) bool) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowInt64(arr []int64, size int) [][]int64 {
	result := make([][]int64, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachInt8(arr []int8, f func(el int8)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryInt8(arr []int8, nth int) []int8 {
	result := make([]int8, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileInt8(arr []int8, f func(arr int8) bool) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameInt8(arr []int8) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileInt8(arr []int8, f func(el int8) bool) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowInt8(arr []int8, size int) [][]int8 {
	result := make([][]int8, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func Each(arr []interface{}, f func(el interface{})) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEvery(arr []interface{}, nth int) []interface{} {
	result := make([]interface{}, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhile(arr []interface{}, f func(arr interface{}) bool) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func Same(arr []interface{}) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhile(arr []interface{}, f func(el interface{}) bool) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func Window(arr []interface{}, size int) [][]interface{} {
	result := make([][]interface{}, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachString(arr []string, f func(el string)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryString(arr []string, nth int) []string {
	result := make([]string, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileString(arr []string, f func(arr string) bool) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameString(arr []string) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileString(arr []string, f func(el string) bool) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowString(arr []string, size int) [][]string {
	result := make([][]string, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachUint(arr []uint, f func(el uint)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryUint(arr []uint, nth int) []uint {
	result := make([]uint, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileUint(arr []uint, f func(arr uint) bool) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameUint(arr []uint) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileUint(arr []uint, f func(el uint) bool) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowUint(arr []uint, size int) [][]uint {
	result := make([][]uint, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachUint16(arr []uint16, f func(el uint16)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryUint16(arr []uint16, nth int) []uint16 {
	result := make([]uint16, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileUint16(arr []uint16, f func(arr uint16) bool) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameUint16(arr []uint16) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileUint16(arr []uint16, f func(el uint16) bool) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowUint16(arr []uint16, size int) [][]uint16 {
	result := make([][]uint16, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachUint32(arr []uint32, f func(el uint32)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryUint32(arr []uint32, nth int) []uint32 {
	result := make([]uint32, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileUint32(arr []uint32, f func(arr uint32) bool) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameUint32(arr []uint32) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileUint32(arr []uint32, f func(el uint32) bool) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowUint32(arr []uint32, size int) [][]uint32 {
	result := make([][]uint32, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachUint64(arr []uint64, f func(el uint64)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryUint64(arr []uint64, nth int) []uint64 {
	result := make([]uint64, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileUint64(arr []uint64, f func(arr uint64) bool) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameUint64(arr []uint64) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileUint64(arr []uint64, f func(el uint64) bool) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowUint64(arr []uint64, size int) [][]uint64 {
	result := make([][]uint64, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
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

// Each calls f for every element from arr
func EachUint8(arr []uint8, f func(el uint8)) {
	for _, el := range arr {
		f(el)
	}
}

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

// DropEvery returns a slice of every nth element in the enumerable dropped,
// starting with the first element.
func DropEveryUint8(arr []uint8, nth int) []uint8 {
	result := make([]uint8, 0, len(arr)/nth)
	for i, el := range arr {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

// DropWhile drops elements from arr while f returns true
func DropWhileUint8(arr []uint8, f func(arr uint8) bool) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
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

// Same returns true if all element in arr the same
func SameUint8(arr []uint8) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			return false
		}
	}
	return true
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

// TakeWhile takes elements from arr while f returns true
func TakeWhileUint8(arr []uint8, f func(el uint8) bool) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

// Window makes sliding window for a given slice:
// ({1,2,3}, 2) -> (1,2), (2,3)
func WindowUint8(arr []uint8, size int) [][]uint8 {
	result := make([][]uint8, 0, len(arr)/size)
	for i := 0; i <= len(arr)-size; i++ {
		chunk := arr[i : i+size]
		result = append(result, chunk)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBool(arr []bool, f func(el bool) bool) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBool(arr []bool, f func(el bool) bool) map[bool][]bool {
	result := make(map[bool][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanBool(arr []bool, acc bool, f func(el bool, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolByte(arr []bool, f func(el bool) byte) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolByte(arr []bool, f func(el bool) byte) map[byte][]bool {
	result := make(map[byte][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolByte(arr []bool, f func(el bool) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolByte(arr []bool, acc byte, f func(el bool, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolByte(arr []bool, acc byte, f func(el bool, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolByte(arr []bool, acc byte, f func(el bool, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolFloat32(arr []bool, f func(el bool) float32) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolFloat32(arr []bool, f func(el bool) float32) map[float32][]bool {
	result := make(map[float32][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolFloat32(arr []bool, f func(el bool) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolFloat32(arr []bool, acc float32, f func(el bool, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolFloat32(arr []bool, acc float32, f func(el bool, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolFloat32(arr []bool, acc float32, f func(el bool, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolFloat64(arr []bool, f func(el bool) float64) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolFloat64(arr []bool, f func(el bool) float64) map[float64][]bool {
	result := make(map[float64][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolFloat64(arr []bool, f func(el bool) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolFloat64(arr []bool, acc float64, f func(el bool, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolFloat64(arr []bool, acc float64, f func(el bool, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolFloat64(arr []bool, acc float64, f func(el bool, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolInt(arr []bool, f func(el bool) int) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolInt(arr []bool, f func(el bool) int) map[int][]bool {
	result := make(map[int][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolInt(arr []bool, f func(el bool) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolInt(arr []bool, acc int, f func(el bool, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolInt(arr []bool, acc int, f func(el bool, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolInt(arr []bool, acc int, f func(el bool, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolInt16(arr []bool, f func(el bool) int16) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolInt16(arr []bool, f func(el bool) int16) map[int16][]bool {
	result := make(map[int16][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolInt16(arr []bool, f func(el bool) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolInt16(arr []bool, acc int16, f func(el bool, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolInt16(arr []bool, acc int16, f func(el bool, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolInt16(arr []bool, acc int16, f func(el bool, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolInt32(arr []bool, f func(el bool) int32) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolInt32(arr []bool, f func(el bool) int32) map[int32][]bool {
	result := make(map[int32][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolInt32(arr []bool, f func(el bool) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolInt32(arr []bool, acc int32, f func(el bool, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolInt32(arr []bool, acc int32, f func(el bool, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolInt32(arr []bool, acc int32, f func(el bool, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolInt64(arr []bool, f func(el bool) int64) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolInt64(arr []bool, f func(el bool) int64) map[int64][]bool {
	result := make(map[int64][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolInt64(arr []bool, f func(el bool) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolInt64(arr []bool, acc int64, f func(el bool, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolInt64(arr []bool, acc int64, f func(el bool, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolInt64(arr []bool, acc int64, f func(el bool, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolInt8(arr []bool, f func(el bool) int8) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolInt8(arr []bool, f func(el bool) int8) map[int8][]bool {
	result := make(map[int8][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolInt8(arr []bool, f func(el bool) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolInt8(arr []bool, acc int8, f func(el bool, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolInt8(arr []bool, acc int8, f func(el bool, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolInt8(arr []bool, acc int8, f func(el bool, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolInterface(arr []bool, f func(el bool) interface{}) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolInterface(arr []bool, f func(el bool) interface{}) map[interface{}][]bool {
	result := make(map[interface{}][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolInterface(arr []bool, f func(el bool) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolInterface(arr []bool, acc interface{}, f func(el bool, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolInterface(arr []bool, acc interface{}, f func(el bool, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolInterface(arr []bool, acc interface{}, f func(el bool, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolString(arr []bool, f func(el bool) string) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolString(arr []bool, f func(el bool) string) map[string][]bool {
	result := make(map[string][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolString(arr []bool, f func(el bool) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolString(arr []bool, acc string, f func(el bool, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolString(arr []bool, acc string, f func(el bool, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolString(arr []bool, acc string, f func(el bool, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolUint(arr []bool, f func(el bool) uint) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolUint(arr []bool, f func(el bool) uint) map[uint][]bool {
	result := make(map[uint][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolUint(arr []bool, f func(el bool) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolUint(arr []bool, acc uint, f func(el bool, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolUint(arr []bool, acc uint, f func(el bool, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolUint(arr []bool, acc uint, f func(el bool, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolUint16(arr []bool, f func(el bool) uint16) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolUint16(arr []bool, f func(el bool) uint16) map[uint16][]bool {
	result := make(map[uint16][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolUint16(arr []bool, f func(el bool) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolUint16(arr []bool, acc uint16, f func(el bool, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolUint16(arr []bool, acc uint16, f func(el bool, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolUint16(arr []bool, acc uint16, f func(el bool, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolUint32(arr []bool, f func(el bool) uint32) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolUint32(arr []bool, f func(el bool) uint32) map[uint32][]bool {
	result := make(map[uint32][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolUint32(arr []bool, f func(el bool) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolUint32(arr []bool, acc uint32, f func(el bool, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolUint32(arr []bool, acc uint32, f func(el bool, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolUint32(arr []bool, acc uint32, f func(el bool, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolUint64(arr []bool, f func(el bool) uint64) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolUint64(arr []bool, f func(el bool) uint64) map[uint64][]bool {
	result := make(map[uint64][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolUint64(arr []bool, f func(el bool) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolUint64(arr []bool, acc uint64, f func(el bool, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolUint64(arr []bool, acc uint64, f func(el bool, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolUint64(arr []bool, acc uint64, f func(el bool, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByBoolUint8(arr []bool, f func(el bool) uint8) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByBoolUint8(arr []bool, f func(el bool) uint8) map[uint8][]bool {
	result := make(map[uint8][]bool)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapBoolUint8(arr []bool, f func(el bool) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceBoolUint8(arr []bool, acc uint8, f func(el bool, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileBoolUint8(arr []bool, acc uint8, f func(el bool, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanBoolUint8(arr []bool, acc uint8, f func(el bool, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteBool(arr []byte, f func(el byte) bool) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteBool(arr []byte, f func(el byte) bool) map[bool][]byte {
	result := make(map[bool][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteBool(arr []byte, f func(el byte) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteBool(arr []byte, acc bool, f func(el byte, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteBool(arr []byte, acc bool, f func(el byte, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteBool(arr []byte, acc bool, f func(el byte, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByte(arr []byte, f func(el byte) byte) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByte(arr []byte, f func(el byte) byte) map[byte][]byte {
	result := make(map[byte][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanByte(arr []byte, acc byte, f func(el byte, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteFloat32(arr []byte, f func(el byte) float32) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteFloat32(arr []byte, f func(el byte) float32) map[float32][]byte {
	result := make(map[float32][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteFloat32(arr []byte, f func(el byte) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteFloat32(arr []byte, acc float32, f func(el byte, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteFloat32(arr []byte, acc float32, f func(el byte, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteFloat32(arr []byte, acc float32, f func(el byte, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteFloat64(arr []byte, f func(el byte) float64) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteFloat64(arr []byte, f func(el byte) float64) map[float64][]byte {
	result := make(map[float64][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteFloat64(arr []byte, f func(el byte) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteFloat64(arr []byte, acc float64, f func(el byte, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteFloat64(arr []byte, acc float64, f func(el byte, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteFloat64(arr []byte, acc float64, f func(el byte, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteInt(arr []byte, f func(el byte) int) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteInt(arr []byte, f func(el byte) int) map[int][]byte {
	result := make(map[int][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteInt(arr []byte, f func(el byte) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteInt(arr []byte, acc int, f func(el byte, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteInt(arr []byte, acc int, f func(el byte, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteInt(arr []byte, acc int, f func(el byte, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteInt16(arr []byte, f func(el byte) int16) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteInt16(arr []byte, f func(el byte) int16) map[int16][]byte {
	result := make(map[int16][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteInt16(arr []byte, f func(el byte) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteInt16(arr []byte, acc int16, f func(el byte, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteInt16(arr []byte, acc int16, f func(el byte, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteInt16(arr []byte, acc int16, f func(el byte, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteInt32(arr []byte, f func(el byte) int32) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteInt32(arr []byte, f func(el byte) int32) map[int32][]byte {
	result := make(map[int32][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteInt32(arr []byte, f func(el byte) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteInt32(arr []byte, acc int32, f func(el byte, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteInt32(arr []byte, acc int32, f func(el byte, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteInt32(arr []byte, acc int32, f func(el byte, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteInt64(arr []byte, f func(el byte) int64) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteInt64(arr []byte, f func(el byte) int64) map[int64][]byte {
	result := make(map[int64][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteInt64(arr []byte, f func(el byte) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteInt64(arr []byte, acc int64, f func(el byte, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteInt64(arr []byte, acc int64, f func(el byte, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteInt64(arr []byte, acc int64, f func(el byte, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteInt8(arr []byte, f func(el byte) int8) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteInt8(arr []byte, f func(el byte) int8) map[int8][]byte {
	result := make(map[int8][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteInt8(arr []byte, f func(el byte) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteInt8(arr []byte, acc int8, f func(el byte, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteInt8(arr []byte, acc int8, f func(el byte, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteInt8(arr []byte, acc int8, f func(el byte, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteInterface(arr []byte, f func(el byte) interface{}) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteInterface(arr []byte, f func(el byte) interface{}) map[interface{}][]byte {
	result := make(map[interface{}][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteInterface(arr []byte, f func(el byte) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteInterface(arr []byte, acc interface{}, f func(el byte, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteInterface(arr []byte, acc interface{}, f func(el byte, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteInterface(arr []byte, acc interface{}, f func(el byte, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteString(arr []byte, f func(el byte) string) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteString(arr []byte, f func(el byte) string) map[string][]byte {
	result := make(map[string][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteString(arr []byte, f func(el byte) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteString(arr []byte, acc string, f func(el byte, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteString(arr []byte, acc string, f func(el byte, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteString(arr []byte, acc string, f func(el byte, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteUint(arr []byte, f func(el byte) uint) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteUint(arr []byte, f func(el byte) uint) map[uint][]byte {
	result := make(map[uint][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteUint(arr []byte, f func(el byte) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteUint(arr []byte, acc uint, f func(el byte, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteUint(arr []byte, acc uint, f func(el byte, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteUint(arr []byte, acc uint, f func(el byte, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteUint16(arr []byte, f func(el byte) uint16) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteUint16(arr []byte, f func(el byte) uint16) map[uint16][]byte {
	result := make(map[uint16][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteUint16(arr []byte, f func(el byte) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteUint16(arr []byte, acc uint16, f func(el byte, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteUint16(arr []byte, acc uint16, f func(el byte, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteUint16(arr []byte, acc uint16, f func(el byte, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteUint32(arr []byte, f func(el byte) uint32) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteUint32(arr []byte, f func(el byte) uint32) map[uint32][]byte {
	result := make(map[uint32][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteUint32(arr []byte, f func(el byte) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteUint32(arr []byte, acc uint32, f func(el byte, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteUint32(arr []byte, acc uint32, f func(el byte, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteUint32(arr []byte, acc uint32, f func(el byte, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteUint64(arr []byte, f func(el byte) uint64) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteUint64(arr []byte, f func(el byte) uint64) map[uint64][]byte {
	result := make(map[uint64][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteUint64(arr []byte, f func(el byte) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteUint64(arr []byte, acc uint64, f func(el byte, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteUint64(arr []byte, acc uint64, f func(el byte, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteUint64(arr []byte, acc uint64, f func(el byte, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByByteUint8(arr []byte, f func(el byte) uint8) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByByteUint8(arr []byte, f func(el byte) uint8) map[uint8][]byte {
	result := make(map[uint8][]byte)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapByteUint8(arr []byte, f func(el byte) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceByteUint8(arr []byte, acc uint8, f func(el byte, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileByteUint8(arr []byte, acc uint8, f func(el byte, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanByteUint8(arr []byte, acc uint8, f func(el byte, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Bool(arr []float32, f func(el float32) bool) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Bool(arr []float32, f func(el float32) bool) map[bool][]float32 {
	result := make(map[bool][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Bool(arr []float32, f func(el float32) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Bool(arr []float32, acc bool, f func(el float32, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Bool(arr []float32, acc bool, f func(el float32, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Bool(arr []float32, acc bool, f func(el float32, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Byte(arr []float32, f func(el float32) byte) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Byte(arr []float32, f func(el float32) byte) map[byte][]float32 {
	result := make(map[byte][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Byte(arr []float32, f func(el float32) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Byte(arr []float32, acc byte, f func(el float32, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Byte(arr []float32, acc byte, f func(el float32, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Byte(arr []float32, acc byte, f func(el float32, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32(arr []float32, f func(el float32) float32) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32(arr []float32, f func(el float32) float32) map[float32][]float32 {
	result := make(map[float32][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32(arr []float32, acc float32, f func(el float32, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Float64(arr []float32, f func(el float32) float64) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Float64(arr []float32, f func(el float32) float64) map[float64][]float32 {
	result := make(map[float64][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Float64(arr []float32, f func(el float32) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Float64(arr []float32, acc float64, f func(el float32, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Float64(arr []float32, acc float64, f func(el float32, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Float64(arr []float32, acc float64, f func(el float32, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Int(arr []float32, f func(el float32) int) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Int(arr []float32, f func(el float32) int) map[int][]float32 {
	result := make(map[int][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Int(arr []float32, f func(el float32) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Int(arr []float32, acc int, f func(el float32, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Int(arr []float32, acc int, f func(el float32, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Int(arr []float32, acc int, f func(el float32, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Int16(arr []float32, f func(el float32) int16) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Int16(arr []float32, f func(el float32) int16) map[int16][]float32 {
	result := make(map[int16][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Int16(arr []float32, f func(el float32) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Int16(arr []float32, acc int16, f func(el float32, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Int16(arr []float32, acc int16, f func(el float32, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Int16(arr []float32, acc int16, f func(el float32, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Int32(arr []float32, f func(el float32) int32) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Int32(arr []float32, f func(el float32) int32) map[int32][]float32 {
	result := make(map[int32][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Int32(arr []float32, f func(el float32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Int32(arr []float32, acc int32, f func(el float32, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Int32(arr []float32, acc int32, f func(el float32, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Int32(arr []float32, acc int32, f func(el float32, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Int64(arr []float32, f func(el float32) int64) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Int64(arr []float32, f func(el float32) int64) map[int64][]float32 {
	result := make(map[int64][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Int64(arr []float32, f func(el float32) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Int64(arr []float32, acc int64, f func(el float32, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Int64(arr []float32, acc int64, f func(el float32, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Int64(arr []float32, acc int64, f func(el float32, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Int8(arr []float32, f func(el float32) int8) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Int8(arr []float32, f func(el float32) int8) map[int8][]float32 {
	result := make(map[int8][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Int8(arr []float32, f func(el float32) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Int8(arr []float32, acc int8, f func(el float32, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Int8(arr []float32, acc int8, f func(el float32, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Int8(arr []float32, acc int8, f func(el float32, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Interface(arr []float32, f func(el float32) interface{}) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Interface(arr []float32, f func(el float32) interface{}) map[interface{}][]float32 {
	result := make(map[interface{}][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Interface(arr []float32, f func(el float32) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Interface(arr []float32, acc interface{}, f func(el float32, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Interface(arr []float32, acc interface{}, f func(el float32, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Interface(arr []float32, acc interface{}, f func(el float32, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32String(arr []float32, f func(el float32) string) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32String(arr []float32, f func(el float32) string) map[string][]float32 {
	result := make(map[string][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32String(arr []float32, f func(el float32) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32String(arr []float32, acc string, f func(el float32, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32String(arr []float32, acc string, f func(el float32, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32String(arr []float32, acc string, f func(el float32, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Uint(arr []float32, f func(el float32) uint) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Uint(arr []float32, f func(el float32) uint) map[uint][]float32 {
	result := make(map[uint][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Uint(arr []float32, f func(el float32) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Uint(arr []float32, acc uint, f func(el float32, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Uint(arr []float32, acc uint, f func(el float32, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Uint(arr []float32, acc uint, f func(el float32, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Uint16(arr []float32, f func(el float32) uint16) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Uint16(arr []float32, f func(el float32) uint16) map[uint16][]float32 {
	result := make(map[uint16][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Uint16(arr []float32, f func(el float32) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Uint16(arr []float32, acc uint16, f func(el float32, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Uint16(arr []float32, acc uint16, f func(el float32, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Uint16(arr []float32, acc uint16, f func(el float32, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Uint32(arr []float32, f func(el float32) uint32) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Uint32(arr []float32, f func(el float32) uint32) map[uint32][]float32 {
	result := make(map[uint32][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Uint32(arr []float32, f func(el float32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Uint32(arr []float32, acc uint32, f func(el float32, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Uint32(arr []float32, acc uint32, f func(el float32, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Uint32(arr []float32, acc uint32, f func(el float32, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Uint64(arr []float32, f func(el float32) uint64) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Uint64(arr []float32, f func(el float32) uint64) map[uint64][]float32 {
	result := make(map[uint64][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Uint64(arr []float32, f func(el float32) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Uint64(arr []float32, acc uint64, f func(el float32, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Uint64(arr []float32, acc uint64, f func(el float32, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Uint64(arr []float32, acc uint64, f func(el float32, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat32Uint8(arr []float32, f func(el float32) uint8) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat32Uint8(arr []float32, f func(el float32) uint8) map[uint8][]float32 {
	result := make(map[uint8][]float32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat32Uint8(arr []float32, f func(el float32) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat32Uint8(arr []float32, acc uint8, f func(el float32, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat32Uint8(arr []float32, acc uint8, f func(el float32, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat32Uint8(arr []float32, acc uint8, f func(el float32, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Bool(arr []float64, f func(el float64) bool) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Bool(arr []float64, f func(el float64) bool) map[bool][]float64 {
	result := make(map[bool][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Bool(arr []float64, f func(el float64) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Bool(arr []float64, acc bool, f func(el float64, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Bool(arr []float64, acc bool, f func(el float64, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Bool(arr []float64, acc bool, f func(el float64, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Byte(arr []float64, f func(el float64) byte) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Byte(arr []float64, f func(el float64) byte) map[byte][]float64 {
	result := make(map[byte][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Byte(arr []float64, f func(el float64) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Byte(arr []float64, acc byte, f func(el float64, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Byte(arr []float64, acc byte, f func(el float64, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Byte(arr []float64, acc byte, f func(el float64, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Float32(arr []float64, f func(el float64) float32) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Float32(arr []float64, f func(el float64) float32) map[float32][]float64 {
	result := make(map[float32][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Float32(arr []float64, f func(el float64) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Float32(arr []float64, acc float32, f func(el float64, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Float32(arr []float64, acc float32, f func(el float64, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Float32(arr []float64, acc float32, f func(el float64, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64(arr []float64, f func(el float64) float64) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64(arr []float64, f func(el float64) float64) map[float64][]float64 {
	result := make(map[float64][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64(arr []float64, acc float64, f func(el float64, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Int(arr []float64, f func(el float64) int) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Int(arr []float64, f func(el float64) int) map[int][]float64 {
	result := make(map[int][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Int(arr []float64, f func(el float64) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Int(arr []float64, acc int, f func(el float64, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Int(arr []float64, acc int, f func(el float64, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Int(arr []float64, acc int, f func(el float64, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Int16(arr []float64, f func(el float64) int16) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Int16(arr []float64, f func(el float64) int16) map[int16][]float64 {
	result := make(map[int16][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Int16(arr []float64, f func(el float64) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Int16(arr []float64, acc int16, f func(el float64, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Int16(arr []float64, acc int16, f func(el float64, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Int16(arr []float64, acc int16, f func(el float64, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Int32(arr []float64, f func(el float64) int32) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Int32(arr []float64, f func(el float64) int32) map[int32][]float64 {
	result := make(map[int32][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Int32(arr []float64, f func(el float64) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Int32(arr []float64, acc int32, f func(el float64, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Int32(arr []float64, acc int32, f func(el float64, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Int32(arr []float64, acc int32, f func(el float64, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Int64(arr []float64, f func(el float64) int64) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Int64(arr []float64, f func(el float64) int64) map[int64][]float64 {
	result := make(map[int64][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Int64(arr []float64, f func(el float64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Int64(arr []float64, acc int64, f func(el float64, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Int64(arr []float64, acc int64, f func(el float64, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Int64(arr []float64, acc int64, f func(el float64, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Int8(arr []float64, f func(el float64) int8) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Int8(arr []float64, f func(el float64) int8) map[int8][]float64 {
	result := make(map[int8][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Int8(arr []float64, f func(el float64) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Int8(arr []float64, acc int8, f func(el float64, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Int8(arr []float64, acc int8, f func(el float64, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Int8(arr []float64, acc int8, f func(el float64, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Interface(arr []float64, f func(el float64) interface{}) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Interface(arr []float64, f func(el float64) interface{}) map[interface{}][]float64 {
	result := make(map[interface{}][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Interface(arr []float64, f func(el float64) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Interface(arr []float64, acc interface{}, f func(el float64, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Interface(arr []float64, acc interface{}, f func(el float64, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Interface(arr []float64, acc interface{}, f func(el float64, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64String(arr []float64, f func(el float64) string) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64String(arr []float64, f func(el float64) string) map[string][]float64 {
	result := make(map[string][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64String(arr []float64, f func(el float64) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64String(arr []float64, acc string, f func(el float64, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64String(arr []float64, acc string, f func(el float64, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64String(arr []float64, acc string, f func(el float64, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Uint(arr []float64, f func(el float64) uint) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Uint(arr []float64, f func(el float64) uint) map[uint][]float64 {
	result := make(map[uint][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Uint(arr []float64, f func(el float64) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Uint(arr []float64, acc uint, f func(el float64, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Uint(arr []float64, acc uint, f func(el float64, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Uint(arr []float64, acc uint, f func(el float64, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Uint16(arr []float64, f func(el float64) uint16) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Uint16(arr []float64, f func(el float64) uint16) map[uint16][]float64 {
	result := make(map[uint16][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Uint16(arr []float64, f func(el float64) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Uint16(arr []float64, acc uint16, f func(el float64, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Uint16(arr []float64, acc uint16, f func(el float64, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Uint16(arr []float64, acc uint16, f func(el float64, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Uint32(arr []float64, f func(el float64) uint32) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Uint32(arr []float64, f func(el float64) uint32) map[uint32][]float64 {
	result := make(map[uint32][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Uint32(arr []float64, f func(el float64) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Uint32(arr []float64, acc uint32, f func(el float64, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Uint32(arr []float64, acc uint32, f func(el float64, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Uint32(arr []float64, acc uint32, f func(el float64, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Uint64(arr []float64, f func(el float64) uint64) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Uint64(arr []float64, f func(el float64) uint64) map[uint64][]float64 {
	result := make(map[uint64][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Uint64(arr []float64, f func(el float64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Uint64(arr []float64, acc uint64, f func(el float64, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Uint64(arr []float64, acc uint64, f func(el float64, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Uint64(arr []float64, acc uint64, f func(el float64, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByFloat64Uint8(arr []float64, f func(el float64) uint8) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByFloat64Uint8(arr []float64, f func(el float64) uint8) map[uint8][]float64 {
	result := make(map[uint8][]float64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapFloat64Uint8(arr []float64, f func(el float64) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceFloat64Uint8(arr []float64, acc uint8, f func(el float64, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileFloat64Uint8(arr []float64, acc uint8, f func(el float64, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanFloat64Uint8(arr []float64, acc uint8, f func(el float64, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Bool(arr []int16, f func(el int16) bool) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Bool(arr []int16, f func(el int16) bool) map[bool][]int16 {
	result := make(map[bool][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Bool(arr []int16, f func(el int16) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Bool(arr []int16, acc bool, f func(el int16, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Bool(arr []int16, acc bool, f func(el int16, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Bool(arr []int16, acc bool, f func(el int16, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Byte(arr []int16, f func(el int16) byte) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Byte(arr []int16, f func(el int16) byte) map[byte][]int16 {
	result := make(map[byte][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Byte(arr []int16, f func(el int16) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Byte(arr []int16, acc byte, f func(el int16, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Byte(arr []int16, acc byte, f func(el int16, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Byte(arr []int16, acc byte, f func(el int16, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Float32(arr []int16, f func(el int16) float32) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Float32(arr []int16, f func(el int16) float32) map[float32][]int16 {
	result := make(map[float32][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Float32(arr []int16, f func(el int16) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Float32(arr []int16, acc float32, f func(el int16, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Float32(arr []int16, acc float32, f func(el int16, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Float32(arr []int16, acc float32, f func(el int16, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Float64(arr []int16, f func(el int16) float64) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Float64(arr []int16, f func(el int16) float64) map[float64][]int16 {
	result := make(map[float64][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Float64(arr []int16, f func(el int16) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Float64(arr []int16, acc float64, f func(el int16, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Float64(arr []int16, acc float64, f func(el int16, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Float64(arr []int16, acc float64, f func(el int16, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Int(arr []int16, f func(el int16) int) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Int(arr []int16, f func(el int16) int) map[int][]int16 {
	result := make(map[int][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Int(arr []int16, f func(el int16) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Int(arr []int16, acc int, f func(el int16, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Int(arr []int16, acc int, f func(el int16, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Int(arr []int16, acc int, f func(el int16, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16(arr []int16, f func(el int16) int16) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16(arr []int16, f func(el int16) int16) map[int16][]int16 {
	result := make(map[int16][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanInt16(arr []int16, acc int16, f func(el int16, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Int32(arr []int16, f func(el int16) int32) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Int32(arr []int16, f func(el int16) int32) map[int32][]int16 {
	result := make(map[int32][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Int32(arr []int16, f func(el int16) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Int32(arr []int16, acc int32, f func(el int16, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Int32(arr []int16, acc int32, f func(el int16, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Int32(arr []int16, acc int32, f func(el int16, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Int64(arr []int16, f func(el int16) int64) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Int64(arr []int16, f func(el int16) int64) map[int64][]int16 {
	result := make(map[int64][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Int64(arr []int16, f func(el int16) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Int64(arr []int16, acc int64, f func(el int16, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Int64(arr []int16, acc int64, f func(el int16, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Int64(arr []int16, acc int64, f func(el int16, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Int8(arr []int16, f func(el int16) int8) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Int8(arr []int16, f func(el int16) int8) map[int8][]int16 {
	result := make(map[int8][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Int8(arr []int16, f func(el int16) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Int8(arr []int16, acc int8, f func(el int16, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Int8(arr []int16, acc int8, f func(el int16, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Int8(arr []int16, acc int8, f func(el int16, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Interface(arr []int16, f func(el int16) interface{}) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Interface(arr []int16, f func(el int16) interface{}) map[interface{}][]int16 {
	result := make(map[interface{}][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Interface(arr []int16, f func(el int16) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Interface(arr []int16, acc interface{}, f func(el int16, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Interface(arr []int16, acc interface{}, f func(el int16, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Interface(arr []int16, acc interface{}, f func(el int16, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16String(arr []int16, f func(el int16) string) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16String(arr []int16, f func(el int16) string) map[string][]int16 {
	result := make(map[string][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16String(arr []int16, f func(el int16) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16String(arr []int16, acc string, f func(el int16, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16String(arr []int16, acc string, f func(el int16, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16String(arr []int16, acc string, f func(el int16, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Uint(arr []int16, f func(el int16) uint) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Uint(arr []int16, f func(el int16) uint) map[uint][]int16 {
	result := make(map[uint][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Uint(arr []int16, f func(el int16) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Uint(arr []int16, acc uint, f func(el int16, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Uint(arr []int16, acc uint, f func(el int16, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Uint(arr []int16, acc uint, f func(el int16, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Uint16(arr []int16, f func(el int16) uint16) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Uint16(arr []int16, f func(el int16) uint16) map[uint16][]int16 {
	result := make(map[uint16][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Uint16(arr []int16, f func(el int16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Uint16(arr []int16, acc uint16, f func(el int16, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Uint16(arr []int16, acc uint16, f func(el int16, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Uint16(arr []int16, acc uint16, f func(el int16, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Uint32(arr []int16, f func(el int16) uint32) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Uint32(arr []int16, f func(el int16) uint32) map[uint32][]int16 {
	result := make(map[uint32][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Uint32(arr []int16, f func(el int16) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Uint32(arr []int16, acc uint32, f func(el int16, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Uint32(arr []int16, acc uint32, f func(el int16, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Uint32(arr []int16, acc uint32, f func(el int16, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Uint64(arr []int16, f func(el int16) uint64) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Uint64(arr []int16, f func(el int16) uint64) map[uint64][]int16 {
	result := make(map[uint64][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Uint64(arr []int16, f func(el int16) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Uint64(arr []int16, acc uint64, f func(el int16, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Uint64(arr []int16, acc uint64, f func(el int16, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Uint64(arr []int16, acc uint64, f func(el int16, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt16Uint8(arr []int16, f func(el int16) uint8) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt16Uint8(arr []int16, f func(el int16) uint8) map[uint8][]int16 {
	result := make(map[uint8][]int16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt16Uint8(arr []int16, f func(el int16) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt16Uint8(arr []int16, acc uint8, f func(el int16, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt16Uint8(arr []int16, acc uint8, f func(el int16, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt16Uint8(arr []int16, acc uint8, f func(el int16, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Bool(arr []int32, f func(el int32) bool) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Bool(arr []int32, f func(el int32) bool) map[bool][]int32 {
	result := make(map[bool][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Bool(arr []int32, f func(el int32) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Bool(arr []int32, acc bool, f func(el int32, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Bool(arr []int32, acc bool, f func(el int32, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Bool(arr []int32, acc bool, f func(el int32, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Byte(arr []int32, f func(el int32) byte) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Byte(arr []int32, f func(el int32) byte) map[byte][]int32 {
	result := make(map[byte][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Byte(arr []int32, f func(el int32) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Byte(arr []int32, acc byte, f func(el int32, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Byte(arr []int32, acc byte, f func(el int32, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Byte(arr []int32, acc byte, f func(el int32, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Float32(arr []int32, f func(el int32) float32) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Float32(arr []int32, f func(el int32) float32) map[float32][]int32 {
	result := make(map[float32][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Float32(arr []int32, f func(el int32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Float32(arr []int32, acc float32, f func(el int32, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Float32(arr []int32, acc float32, f func(el int32, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Float32(arr []int32, acc float32, f func(el int32, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Float64(arr []int32, f func(el int32) float64) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Float64(arr []int32, f func(el int32) float64) map[float64][]int32 {
	result := make(map[float64][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Float64(arr []int32, f func(el int32) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Float64(arr []int32, acc float64, f func(el int32, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Float64(arr []int32, acc float64, f func(el int32, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Float64(arr []int32, acc float64, f func(el int32, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Int(arr []int32, f func(el int32) int) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Int(arr []int32, f func(el int32) int) map[int][]int32 {
	result := make(map[int][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Int(arr []int32, f func(el int32) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Int(arr []int32, acc int, f func(el int32, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Int(arr []int32, acc int, f func(el int32, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Int(arr []int32, acc int, f func(el int32, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Int16(arr []int32, f func(el int32) int16) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Int16(arr []int32, f func(el int32) int16) map[int16][]int32 {
	result := make(map[int16][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Int16(arr []int32, f func(el int32) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Int16(arr []int32, acc int16, f func(el int32, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Int16(arr []int32, acc int16, f func(el int32, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Int16(arr []int32, acc int16, f func(el int32, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32(arr []int32, f func(el int32) int32) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32(arr []int32, f func(el int32) int32) map[int32][]int32 {
	result := make(map[int32][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanInt32(arr []int32, acc int32, f func(el int32, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Int64(arr []int32, f func(el int32) int64) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Int64(arr []int32, f func(el int32) int64) map[int64][]int32 {
	result := make(map[int64][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Int64(arr []int32, f func(el int32) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Int64(arr []int32, acc int64, f func(el int32, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Int64(arr []int32, acc int64, f func(el int32, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Int64(arr []int32, acc int64, f func(el int32, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Int8(arr []int32, f func(el int32) int8) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Int8(arr []int32, f func(el int32) int8) map[int8][]int32 {
	result := make(map[int8][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Int8(arr []int32, f func(el int32) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Int8(arr []int32, acc int8, f func(el int32, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Int8(arr []int32, acc int8, f func(el int32, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Int8(arr []int32, acc int8, f func(el int32, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Interface(arr []int32, f func(el int32) interface{}) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Interface(arr []int32, f func(el int32) interface{}) map[interface{}][]int32 {
	result := make(map[interface{}][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Interface(arr []int32, f func(el int32) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Interface(arr []int32, acc interface{}, f func(el int32, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Interface(arr []int32, acc interface{}, f func(el int32, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Interface(arr []int32, acc interface{}, f func(el int32, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32String(arr []int32, f func(el int32) string) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32String(arr []int32, f func(el int32) string) map[string][]int32 {
	result := make(map[string][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32String(arr []int32, f func(el int32) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32String(arr []int32, acc string, f func(el int32, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32String(arr []int32, acc string, f func(el int32, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32String(arr []int32, acc string, f func(el int32, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Uint(arr []int32, f func(el int32) uint) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Uint(arr []int32, f func(el int32) uint) map[uint][]int32 {
	result := make(map[uint][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Uint(arr []int32, f func(el int32) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Uint(arr []int32, acc uint, f func(el int32, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Uint(arr []int32, acc uint, f func(el int32, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Uint(arr []int32, acc uint, f func(el int32, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Uint16(arr []int32, f func(el int32) uint16) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Uint16(arr []int32, f func(el int32) uint16) map[uint16][]int32 {
	result := make(map[uint16][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Uint16(arr []int32, f func(el int32) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Uint16(arr []int32, acc uint16, f func(el int32, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Uint16(arr []int32, acc uint16, f func(el int32, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Uint16(arr []int32, acc uint16, f func(el int32, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Uint32(arr []int32, f func(el int32) uint32) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Uint32(arr []int32, f func(el int32) uint32) map[uint32][]int32 {
	result := make(map[uint32][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Uint32(arr []int32, f func(el int32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Uint32(arr []int32, acc uint32, f func(el int32, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Uint32(arr []int32, acc uint32, f func(el int32, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Uint32(arr []int32, acc uint32, f func(el int32, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Uint64(arr []int32, f func(el int32) uint64) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Uint64(arr []int32, f func(el int32) uint64) map[uint64][]int32 {
	result := make(map[uint64][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Uint64(arr []int32, f func(el int32) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Uint64(arr []int32, acc uint64, f func(el int32, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Uint64(arr []int32, acc uint64, f func(el int32, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Uint64(arr []int32, acc uint64, f func(el int32, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt32Uint8(arr []int32, f func(el int32) uint8) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt32Uint8(arr []int32, f func(el int32) uint8) map[uint8][]int32 {
	result := make(map[uint8][]int32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt32Uint8(arr []int32, f func(el int32) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt32Uint8(arr []int32, acc uint8, f func(el int32, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt32Uint8(arr []int32, acc uint8, f func(el int32, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt32Uint8(arr []int32, acc uint8, f func(el int32, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Bool(arr []int64, f func(el int64) bool) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Bool(arr []int64, f func(el int64) bool) map[bool][]int64 {
	result := make(map[bool][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Bool(arr []int64, f func(el int64) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Bool(arr []int64, acc bool, f func(el int64, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Bool(arr []int64, acc bool, f func(el int64, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Bool(arr []int64, acc bool, f func(el int64, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Byte(arr []int64, f func(el int64) byte) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Byte(arr []int64, f func(el int64) byte) map[byte][]int64 {
	result := make(map[byte][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Byte(arr []int64, f func(el int64) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Byte(arr []int64, acc byte, f func(el int64, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Byte(arr []int64, acc byte, f func(el int64, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Byte(arr []int64, acc byte, f func(el int64, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Float32(arr []int64, f func(el int64) float32) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Float32(arr []int64, f func(el int64) float32) map[float32][]int64 {
	result := make(map[float32][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Float32(arr []int64, f func(el int64) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Float32(arr []int64, acc float32, f func(el int64, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Float32(arr []int64, acc float32, f func(el int64, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Float32(arr []int64, acc float32, f func(el int64, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Float64(arr []int64, f func(el int64) float64) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Float64(arr []int64, f func(el int64) float64) map[float64][]int64 {
	result := make(map[float64][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Float64(arr []int64, f func(el int64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Float64(arr []int64, acc float64, f func(el int64, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Float64(arr []int64, acc float64, f func(el int64, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Float64(arr []int64, acc float64, f func(el int64, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Int(arr []int64, f func(el int64) int) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Int(arr []int64, f func(el int64) int) map[int][]int64 {
	result := make(map[int][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Int(arr []int64, f func(el int64) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Int(arr []int64, acc int, f func(el int64, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Int(arr []int64, acc int, f func(el int64, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Int(arr []int64, acc int, f func(el int64, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Int16(arr []int64, f func(el int64) int16) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Int16(arr []int64, f func(el int64) int16) map[int16][]int64 {
	result := make(map[int16][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Int16(arr []int64, f func(el int64) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Int16(arr []int64, acc int16, f func(el int64, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Int16(arr []int64, acc int16, f func(el int64, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Int16(arr []int64, acc int16, f func(el int64, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Int32(arr []int64, f func(el int64) int32) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Int32(arr []int64, f func(el int64) int32) map[int32][]int64 {
	result := make(map[int32][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Int32(arr []int64, f func(el int64) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Int32(arr []int64, acc int32, f func(el int64, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Int32(arr []int64, acc int32, f func(el int64, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Int32(arr []int64, acc int32, f func(el int64, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64(arr []int64, f func(el int64) int64) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64(arr []int64, f func(el int64) int64) map[int64][]int64 {
	result := make(map[int64][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanInt64(arr []int64, acc int64, f func(el int64, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Int8(arr []int64, f func(el int64) int8) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Int8(arr []int64, f func(el int64) int8) map[int8][]int64 {
	result := make(map[int8][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Int8(arr []int64, f func(el int64) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Int8(arr []int64, acc int8, f func(el int64, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Int8(arr []int64, acc int8, f func(el int64, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Int8(arr []int64, acc int8, f func(el int64, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Interface(arr []int64, f func(el int64) interface{}) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Interface(arr []int64, f func(el int64) interface{}) map[interface{}][]int64 {
	result := make(map[interface{}][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Interface(arr []int64, f func(el int64) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Interface(arr []int64, acc interface{}, f func(el int64, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Interface(arr []int64, acc interface{}, f func(el int64, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Interface(arr []int64, acc interface{}, f func(el int64, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64String(arr []int64, f func(el int64) string) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64String(arr []int64, f func(el int64) string) map[string][]int64 {
	result := make(map[string][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64String(arr []int64, f func(el int64) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64String(arr []int64, acc string, f func(el int64, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64String(arr []int64, acc string, f func(el int64, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64String(arr []int64, acc string, f func(el int64, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Uint(arr []int64, f func(el int64) uint) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Uint(arr []int64, f func(el int64) uint) map[uint][]int64 {
	result := make(map[uint][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Uint(arr []int64, f func(el int64) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Uint(arr []int64, acc uint, f func(el int64, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Uint(arr []int64, acc uint, f func(el int64, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Uint(arr []int64, acc uint, f func(el int64, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Uint16(arr []int64, f func(el int64) uint16) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Uint16(arr []int64, f func(el int64) uint16) map[uint16][]int64 {
	result := make(map[uint16][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Uint16(arr []int64, f func(el int64) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Uint16(arr []int64, acc uint16, f func(el int64, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Uint16(arr []int64, acc uint16, f func(el int64, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Uint16(arr []int64, acc uint16, f func(el int64, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Uint32(arr []int64, f func(el int64) uint32) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Uint32(arr []int64, f func(el int64) uint32) map[uint32][]int64 {
	result := make(map[uint32][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Uint32(arr []int64, f func(el int64) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Uint32(arr []int64, acc uint32, f func(el int64, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Uint32(arr []int64, acc uint32, f func(el int64, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Uint32(arr []int64, acc uint32, f func(el int64, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Uint64(arr []int64, f func(el int64) uint64) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Uint64(arr []int64, f func(el int64) uint64) map[uint64][]int64 {
	result := make(map[uint64][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Uint64(arr []int64, f func(el int64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Uint64(arr []int64, acc uint64, f func(el int64, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Uint64(arr []int64, acc uint64, f func(el int64, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Uint64(arr []int64, acc uint64, f func(el int64, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt64Uint8(arr []int64, f func(el int64) uint8) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt64Uint8(arr []int64, f func(el int64) uint8) map[uint8][]int64 {
	result := make(map[uint8][]int64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt64Uint8(arr []int64, f func(el int64) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt64Uint8(arr []int64, acc uint8, f func(el int64, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt64Uint8(arr []int64, acc uint8, f func(el int64, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt64Uint8(arr []int64, acc uint8, f func(el int64, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Bool(arr []int8, f func(el int8) bool) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Bool(arr []int8, f func(el int8) bool) map[bool][]int8 {
	result := make(map[bool][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Bool(arr []int8, f func(el int8) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Bool(arr []int8, acc bool, f func(el int8, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Bool(arr []int8, acc bool, f func(el int8, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Bool(arr []int8, acc bool, f func(el int8, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Byte(arr []int8, f func(el int8) byte) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Byte(arr []int8, f func(el int8) byte) map[byte][]int8 {
	result := make(map[byte][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Byte(arr []int8, f func(el int8) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Byte(arr []int8, acc byte, f func(el int8, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Byte(arr []int8, acc byte, f func(el int8, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Byte(arr []int8, acc byte, f func(el int8, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Float32(arr []int8, f func(el int8) float32) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Float32(arr []int8, f func(el int8) float32) map[float32][]int8 {
	result := make(map[float32][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Float32(arr []int8, f func(el int8) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Float32(arr []int8, acc float32, f func(el int8, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Float32(arr []int8, acc float32, f func(el int8, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Float32(arr []int8, acc float32, f func(el int8, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Float64(arr []int8, f func(el int8) float64) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Float64(arr []int8, f func(el int8) float64) map[float64][]int8 {
	result := make(map[float64][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Float64(arr []int8, f func(el int8) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Float64(arr []int8, acc float64, f func(el int8, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Float64(arr []int8, acc float64, f func(el int8, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Float64(arr []int8, acc float64, f func(el int8, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Int(arr []int8, f func(el int8) int) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Int(arr []int8, f func(el int8) int) map[int][]int8 {
	result := make(map[int][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Int(arr []int8, f func(el int8) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Int(arr []int8, acc int, f func(el int8, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Int(arr []int8, acc int, f func(el int8, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Int(arr []int8, acc int, f func(el int8, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Int16(arr []int8, f func(el int8) int16) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Int16(arr []int8, f func(el int8) int16) map[int16][]int8 {
	result := make(map[int16][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Int16(arr []int8, f func(el int8) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Int16(arr []int8, acc int16, f func(el int8, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Int16(arr []int8, acc int16, f func(el int8, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Int16(arr []int8, acc int16, f func(el int8, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Int32(arr []int8, f func(el int8) int32) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Int32(arr []int8, f func(el int8) int32) map[int32][]int8 {
	result := make(map[int32][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Int32(arr []int8, f func(el int8) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Int32(arr []int8, acc int32, f func(el int8, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Int32(arr []int8, acc int32, f func(el int8, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Int32(arr []int8, acc int32, f func(el int8, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Int64(arr []int8, f func(el int8) int64) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Int64(arr []int8, f func(el int8) int64) map[int64][]int8 {
	result := make(map[int64][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Int64(arr []int8, f func(el int8) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Int64(arr []int8, acc int64, f func(el int8, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Int64(arr []int8, acc int64, f func(el int8, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Int64(arr []int8, acc int64, f func(el int8, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8(arr []int8, f func(el int8) int8) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8(arr []int8, f func(el int8) int8) map[int8][]int8 {
	result := make(map[int8][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanInt8(arr []int8, acc int8, f func(el int8, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Interface(arr []int8, f func(el int8) interface{}) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Interface(arr []int8, f func(el int8) interface{}) map[interface{}][]int8 {
	result := make(map[interface{}][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Interface(arr []int8, f func(el int8) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Interface(arr []int8, acc interface{}, f func(el int8, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Interface(arr []int8, acc interface{}, f func(el int8, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Interface(arr []int8, acc interface{}, f func(el int8, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8String(arr []int8, f func(el int8) string) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8String(arr []int8, f func(el int8) string) map[string][]int8 {
	result := make(map[string][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8String(arr []int8, f func(el int8) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8String(arr []int8, acc string, f func(el int8, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8String(arr []int8, acc string, f func(el int8, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8String(arr []int8, acc string, f func(el int8, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Uint(arr []int8, f func(el int8) uint) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Uint(arr []int8, f func(el int8) uint) map[uint][]int8 {
	result := make(map[uint][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Uint(arr []int8, f func(el int8) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Uint(arr []int8, acc uint, f func(el int8, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Uint(arr []int8, acc uint, f func(el int8, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Uint(arr []int8, acc uint, f func(el int8, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Uint16(arr []int8, f func(el int8) uint16) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Uint16(arr []int8, f func(el int8) uint16) map[uint16][]int8 {
	result := make(map[uint16][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Uint16(arr []int8, f func(el int8) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Uint16(arr []int8, acc uint16, f func(el int8, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Uint16(arr []int8, acc uint16, f func(el int8, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Uint16(arr []int8, acc uint16, f func(el int8, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Uint32(arr []int8, f func(el int8) uint32) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Uint32(arr []int8, f func(el int8) uint32) map[uint32][]int8 {
	result := make(map[uint32][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Uint32(arr []int8, f func(el int8) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Uint32(arr []int8, acc uint32, f func(el int8, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Uint32(arr []int8, acc uint32, f func(el int8, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Uint32(arr []int8, acc uint32, f func(el int8, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Uint64(arr []int8, f func(el int8) uint64) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Uint64(arr []int8, f func(el int8) uint64) map[uint64][]int8 {
	result := make(map[uint64][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Uint64(arr []int8, f func(el int8) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Uint64(arr []int8, acc uint64, f func(el int8, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Uint64(arr []int8, acc uint64, f func(el int8, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Uint64(arr []int8, acc uint64, f func(el int8, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt8Uint8(arr []int8, f func(el int8) uint8) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt8Uint8(arr []int8, f func(el int8) uint8) map[uint8][]int8 {
	result := make(map[uint8][]int8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInt8Uint8(arr []int8, f func(el int8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInt8Uint8(arr []int8, acc uint8, f func(el int8, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInt8Uint8(arr []int8, acc uint8, f func(el int8, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInt8Uint8(arr []int8, acc uint8, f func(el int8, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntBool(arr []int, f func(el int) bool) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntBool(arr []int, f func(el int) bool) map[bool][]int {
	result := make(map[bool][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntBool(arr []int, f func(el int) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntBool(arr []int, acc bool, f func(el int, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntBool(arr []int, acc bool, f func(el int, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntBool(arr []int, acc bool, f func(el int, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntByte(arr []int, f func(el int) byte) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntByte(arr []int, f func(el int) byte) map[byte][]int {
	result := make(map[byte][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntByte(arr []int, f func(el int) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntByte(arr []int, acc byte, f func(el int, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntByte(arr []int, acc byte, f func(el int, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntByte(arr []int, acc byte, f func(el int, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceBool(arr []interface{}, f func(el interface{}) bool) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceBool(arr []interface{}, f func(el interface{}) bool) map[bool][]interface{} {
	result := make(map[bool][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceBool(arr []interface{}, f func(el interface{}) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceBool(arr []interface{}, acc bool, f func(el interface{}, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceBool(arr []interface{}, acc bool, f func(el interface{}, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceBool(arr []interface{}, acc bool, f func(el interface{}, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceByte(arr []interface{}, f func(el interface{}) byte) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceByte(arr []interface{}, f func(el interface{}) byte) map[byte][]interface{} {
	result := make(map[byte][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceByte(arr []interface{}, f func(el interface{}) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceByte(arr []interface{}, acc byte, f func(el interface{}, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceByte(arr []interface{}, acc byte, f func(el interface{}, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceByte(arr []interface{}, acc byte, f func(el interface{}, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceFloat32(arr []interface{}, f func(el interface{}) float32) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceFloat32(arr []interface{}, f func(el interface{}) float32) map[float32][]interface{} {
	result := make(map[float32][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceFloat32(arr []interface{}, f func(el interface{}) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceFloat32(arr []interface{}, acc float32, f func(el interface{}, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceFloat32(arr []interface{}, acc float32, f func(el interface{}, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceFloat32(arr []interface{}, acc float32, f func(el interface{}, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceFloat64(arr []interface{}, f func(el interface{}) float64) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceFloat64(arr []interface{}, f func(el interface{}) float64) map[float64][]interface{} {
	result := make(map[float64][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceFloat64(arr []interface{}, f func(el interface{}) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceFloat64(arr []interface{}, acc float64, f func(el interface{}, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceFloat64(arr []interface{}, acc float64, f func(el interface{}, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceFloat64(arr []interface{}, acc float64, f func(el interface{}, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceInt(arr []interface{}, f func(el interface{}) int) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceInt(arr []interface{}, f func(el interface{}) int) map[int][]interface{} {
	result := make(map[int][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceInt(arr []interface{}, f func(el interface{}) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceInt(arr []interface{}, acc int, f func(el interface{}, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceInt(arr []interface{}, acc int, f func(el interface{}, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceInt(arr []interface{}, acc int, f func(el interface{}, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceInt16(arr []interface{}, f func(el interface{}) int16) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceInt16(arr []interface{}, f func(el interface{}) int16) map[int16][]interface{} {
	result := make(map[int16][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceInt16(arr []interface{}, f func(el interface{}) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceInt16(arr []interface{}, acc int16, f func(el interface{}, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceInt16(arr []interface{}, acc int16, f func(el interface{}, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceInt16(arr []interface{}, acc int16, f func(el interface{}, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceInt32(arr []interface{}, f func(el interface{}) int32) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceInt32(arr []interface{}, f func(el interface{}) int32) map[int32][]interface{} {
	result := make(map[int32][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceInt32(arr []interface{}, f func(el interface{}) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceInt32(arr []interface{}, acc int32, f func(el interface{}, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceInt32(arr []interface{}, acc int32, f func(el interface{}, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceInt32(arr []interface{}, acc int32, f func(el interface{}, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceInt64(arr []interface{}, f func(el interface{}) int64) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceInt64(arr []interface{}, f func(el interface{}) int64) map[int64][]interface{} {
	result := make(map[int64][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceInt64(arr []interface{}, f func(el interface{}) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceInt64(arr []interface{}, acc int64, f func(el interface{}, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceInt64(arr []interface{}, acc int64, f func(el interface{}, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceInt64(arr []interface{}, acc int64, f func(el interface{}, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceInt8(arr []interface{}, f func(el interface{}) int8) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceInt8(arr []interface{}, f func(el interface{}) int8) map[int8][]interface{} {
	result := make(map[int8][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceInt8(arr []interface{}, f func(el interface{}) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceInt8(arr []interface{}, acc int8, f func(el interface{}, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceInt8(arr []interface{}, acc int8, f func(el interface{}, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceInt8(arr []interface{}, acc int8, f func(el interface{}, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterface(arr []interface{}, f func(el interface{}) interface{}) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterface(arr []interface{}, f func(el interface{}) interface{}) map[interface{}][]interface{} {
	result := make(map[interface{}][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterface(arr []interface{}, f func(el interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterface(arr []interface{}, acc interface{}, f func(el interface{}, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterface(arr []interface{}, acc interface{}, f func(el interface{}, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterface(arr []interface{}, acc interface{}, f func(el interface{}, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceString(arr []interface{}, f func(el interface{}) string) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceString(arr []interface{}, f func(el interface{}) string) map[string][]interface{} {
	result := make(map[string][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceString(arr []interface{}, f func(el interface{}) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceString(arr []interface{}, acc string, f func(el interface{}, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceString(arr []interface{}, acc string, f func(el interface{}, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceString(arr []interface{}, acc string, f func(el interface{}, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceUint(arr []interface{}, f func(el interface{}) uint) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceUint(arr []interface{}, f func(el interface{}) uint) map[uint][]interface{} {
	result := make(map[uint][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceUint(arr []interface{}, f func(el interface{}) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceUint(arr []interface{}, acc uint, f func(el interface{}, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceUint(arr []interface{}, acc uint, f func(el interface{}, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceUint(arr []interface{}, acc uint, f func(el interface{}, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceUint16(arr []interface{}, f func(el interface{}) uint16) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceUint16(arr []interface{}, f func(el interface{}) uint16) map[uint16][]interface{} {
	result := make(map[uint16][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceUint16(arr []interface{}, f func(el interface{}) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceUint16(arr []interface{}, acc uint16, f func(el interface{}, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceUint16(arr []interface{}, acc uint16, f func(el interface{}, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceUint16(arr []interface{}, acc uint16, f func(el interface{}, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceUint32(arr []interface{}, f func(el interface{}) uint32) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceUint32(arr []interface{}, f func(el interface{}) uint32) map[uint32][]interface{} {
	result := make(map[uint32][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceUint32(arr []interface{}, f func(el interface{}) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceUint32(arr []interface{}, acc uint32, f func(el interface{}, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceUint32(arr []interface{}, acc uint32, f func(el interface{}, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceUint32(arr []interface{}, acc uint32, f func(el interface{}, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceUint64(arr []interface{}, f func(el interface{}) uint64) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceUint64(arr []interface{}, f func(el interface{}) uint64) map[uint64][]interface{} {
	result := make(map[uint64][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceUint64(arr []interface{}, f func(el interface{}) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceUint64(arr []interface{}, acc uint64, f func(el interface{}, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceUint64(arr []interface{}, acc uint64, f func(el interface{}, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceUint64(arr []interface{}, acc uint64, f func(el interface{}, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInterfaceUint8(arr []interface{}, f func(el interface{}) uint8) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInterfaceUint8(arr []interface{}, f func(el interface{}) uint8) map[uint8][]interface{} {
	result := make(map[uint8][]interface{})
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapInterfaceUint8(arr []interface{}, f func(el interface{}) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceInterfaceUint8(arr []interface{}, acc uint8, f func(el interface{}, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileInterfaceUint8(arr []interface{}, acc uint8, f func(el interface{}, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanInterfaceUint8(arr []interface{}, acc uint8, f func(el interface{}, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntFloat32(arr []int, f func(el int) float32) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntFloat32(arr []int, f func(el int) float32) map[float32][]int {
	result := make(map[float32][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntFloat32(arr []int, f func(el int) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntFloat32(arr []int, acc float32, f func(el int, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntFloat32(arr []int, acc float32, f func(el int, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntFloat32(arr []int, acc float32, f func(el int, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntFloat64(arr []int, f func(el int) float64) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntFloat64(arr []int, f func(el int) float64) map[float64][]int {
	result := make(map[float64][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntFloat64(arr []int, f func(el int) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntFloat64(arr []int, acc float64, f func(el int, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntFloat64(arr []int, acc float64, f func(el int, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntFloat64(arr []int, acc float64, f func(el int, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByInt(arr []int, f func(el int) int) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByInt(arr []int, f func(el int) int) map[int][]int {
	result := make(map[int][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanInt(arr []int, acc int, f func(el int, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntInt16(arr []int, f func(el int) int16) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntInt16(arr []int, f func(el int) int16) map[int16][]int {
	result := make(map[int16][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntInt16(arr []int, f func(el int) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntInt16(arr []int, acc int16, f func(el int, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntInt16(arr []int, acc int16, f func(el int, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntInt16(arr []int, acc int16, f func(el int, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntInt32(arr []int, f func(el int) int32) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntInt32(arr []int, f func(el int) int32) map[int32][]int {
	result := make(map[int32][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntInt32(arr []int, f func(el int) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntInt32(arr []int, acc int32, f func(el int, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntInt32(arr []int, acc int32, f func(el int, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntInt32(arr []int, acc int32, f func(el int, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntInt64(arr []int, f func(el int) int64) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntInt64(arr []int, f func(el int) int64) map[int64][]int {
	result := make(map[int64][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntInt64(arr []int, f func(el int) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntInt64(arr []int, acc int64, f func(el int, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntInt64(arr []int, acc int64, f func(el int, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntInt64(arr []int, acc int64, f func(el int, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntInt8(arr []int, f func(el int) int8) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntInt8(arr []int, f func(el int) int8) map[int8][]int {
	result := make(map[int8][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntInt8(arr []int, f func(el int) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntInt8(arr []int, acc int8, f func(el int, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntInt8(arr []int, acc int8, f func(el int, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntInt8(arr []int, acc int8, f func(el int, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntInterface(arr []int, f func(el int) interface{}) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntInterface(arr []int, f func(el int) interface{}) map[interface{}][]int {
	result := make(map[interface{}][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntInterface(arr []int, f func(el int) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntInterface(arr []int, acc interface{}, f func(el int, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntInterface(arr []int, acc interface{}, f func(el int, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntInterface(arr []int, acc interface{}, f func(el int, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntString(arr []int, f func(el int) string) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntString(arr []int, f func(el int) string) map[string][]int {
	result := make(map[string][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntString(arr []int, f func(el int) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntString(arr []int, acc string, f func(el int, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntString(arr []int, acc string, f func(el int, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntString(arr []int, acc string, f func(el int, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntUint(arr []int, f func(el int) uint) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntUint(arr []int, f func(el int) uint) map[uint][]int {
	result := make(map[uint][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntUint(arr []int, f func(el int) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntUint(arr []int, acc uint, f func(el int, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntUint(arr []int, acc uint, f func(el int, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntUint(arr []int, acc uint, f func(el int, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntUint16(arr []int, f func(el int) uint16) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntUint16(arr []int, f func(el int) uint16) map[uint16][]int {
	result := make(map[uint16][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntUint16(arr []int, f func(el int) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntUint16(arr []int, acc uint16, f func(el int, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntUint16(arr []int, acc uint16, f func(el int, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntUint16(arr []int, acc uint16, f func(el int, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntUint32(arr []int, f func(el int) uint32) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntUint32(arr []int, f func(el int) uint32) map[uint32][]int {
	result := make(map[uint32][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntUint32(arr []int, f func(el int) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntUint32(arr []int, acc uint32, f func(el int, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntUint32(arr []int, acc uint32, f func(el int, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntUint32(arr []int, acc uint32, f func(el int, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntUint64(arr []int, f func(el int) uint64) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntUint64(arr []int, f func(el int) uint64) map[uint64][]int {
	result := make(map[uint64][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntUint64(arr []int, f func(el int) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntUint64(arr []int, acc uint64, f func(el int, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntUint64(arr []int, acc uint64, f func(el int, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntUint64(arr []int, acc uint64, f func(el int, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByIntUint8(arr []int, f func(el int) uint8) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByIntUint8(arr []int, f func(el int) uint8) map[uint8][]int {
	result := make(map[uint8][]int)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapIntUint8(arr []int, f func(el int) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceIntUint8(arr []int, acc uint8, f func(el int, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileIntUint8(arr []int, acc uint8, f func(el int, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanIntUint8(arr []int, acc uint8, f func(el int, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringBool(arr []string, f func(el string) bool) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringBool(arr []string, f func(el string) bool) map[bool][]string {
	result := make(map[bool][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringBool(arr []string, f func(el string) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringBool(arr []string, acc bool, f func(el string, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringBool(arr []string, acc bool, f func(el string, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringBool(arr []string, acc bool, f func(el string, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringByte(arr []string, f func(el string) byte) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringByte(arr []string, f func(el string) byte) map[byte][]string {
	result := make(map[byte][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringByte(arr []string, f func(el string) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringByte(arr []string, acc byte, f func(el string, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringByte(arr []string, acc byte, f func(el string, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringByte(arr []string, acc byte, f func(el string, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringFloat32(arr []string, f func(el string) float32) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringFloat32(arr []string, f func(el string) float32) map[float32][]string {
	result := make(map[float32][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringFloat32(arr []string, f func(el string) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringFloat32(arr []string, acc float32, f func(el string, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringFloat32(arr []string, acc float32, f func(el string, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringFloat32(arr []string, acc float32, f func(el string, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringFloat64(arr []string, f func(el string) float64) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringFloat64(arr []string, f func(el string) float64) map[float64][]string {
	result := make(map[float64][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringFloat64(arr []string, f func(el string) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringFloat64(arr []string, acc float64, f func(el string, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringFloat64(arr []string, acc float64, f func(el string, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringFloat64(arr []string, acc float64, f func(el string, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringInt(arr []string, f func(el string) int) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringInt(arr []string, f func(el string) int) map[int][]string {
	result := make(map[int][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringInt(arr []string, f func(el string) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringInt(arr []string, acc int, f func(el string, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringInt(arr []string, acc int, f func(el string, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringInt(arr []string, acc int, f func(el string, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringInt16(arr []string, f func(el string) int16) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringInt16(arr []string, f func(el string) int16) map[int16][]string {
	result := make(map[int16][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringInt16(arr []string, f func(el string) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringInt16(arr []string, acc int16, f func(el string, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringInt16(arr []string, acc int16, f func(el string, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringInt16(arr []string, acc int16, f func(el string, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringInt32(arr []string, f func(el string) int32) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringInt32(arr []string, f func(el string) int32) map[int32][]string {
	result := make(map[int32][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringInt32(arr []string, f func(el string) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringInt32(arr []string, acc int32, f func(el string, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringInt32(arr []string, acc int32, f func(el string, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringInt32(arr []string, acc int32, f func(el string, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringInt64(arr []string, f func(el string) int64) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringInt64(arr []string, f func(el string) int64) map[int64][]string {
	result := make(map[int64][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringInt64(arr []string, f func(el string) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringInt64(arr []string, acc int64, f func(el string, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringInt64(arr []string, acc int64, f func(el string, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringInt64(arr []string, acc int64, f func(el string, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringInt8(arr []string, f func(el string) int8) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringInt8(arr []string, f func(el string) int8) map[int8][]string {
	result := make(map[int8][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringInt8(arr []string, f func(el string) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringInt8(arr []string, acc int8, f func(el string, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringInt8(arr []string, acc int8, f func(el string, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringInt8(arr []string, acc int8, f func(el string, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringInterface(arr []string, f func(el string) interface{}) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringInterface(arr []string, f func(el string) interface{}) map[interface{}][]string {
	result := make(map[interface{}][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringInterface(arr []string, f func(el string) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringInterface(arr []string, acc interface{}, f func(el string, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringInterface(arr []string, acc interface{}, f func(el string, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringInterface(arr []string, acc interface{}, f func(el string, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByString(arr []string, f func(el string) string) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByString(arr []string, f func(el string) string) map[string][]string {
	result := make(map[string][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanString(arr []string, acc string, f func(el string, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringUint(arr []string, f func(el string) uint) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringUint(arr []string, f func(el string) uint) map[uint][]string {
	result := make(map[uint][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringUint(arr []string, f func(el string) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringUint(arr []string, acc uint, f func(el string, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringUint(arr []string, acc uint, f func(el string, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringUint(arr []string, acc uint, f func(el string, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringUint16(arr []string, f func(el string) uint16) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringUint16(arr []string, f func(el string) uint16) map[uint16][]string {
	result := make(map[uint16][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringUint16(arr []string, f func(el string) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringUint16(arr []string, acc uint16, f func(el string, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringUint16(arr []string, acc uint16, f func(el string, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringUint16(arr []string, acc uint16, f func(el string, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringUint32(arr []string, f func(el string) uint32) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringUint32(arr []string, f func(el string) uint32) map[uint32][]string {
	result := make(map[uint32][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringUint32(arr []string, f func(el string) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringUint32(arr []string, acc uint32, f func(el string, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringUint32(arr []string, acc uint32, f func(el string, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringUint32(arr []string, acc uint32, f func(el string, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringUint64(arr []string, f func(el string) uint64) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringUint64(arr []string, f func(el string) uint64) map[uint64][]string {
	result := make(map[uint64][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringUint64(arr []string, f func(el string) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringUint64(arr []string, acc uint64, f func(el string, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringUint64(arr []string, acc uint64, f func(el string, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringUint64(arr []string, acc uint64, f func(el string, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByStringUint8(arr []string, f func(el string) uint8) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByStringUint8(arr []string, f func(el string) uint8) map[uint8][]string {
	result := make(map[uint8][]string)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapStringUint8(arr []string, f func(el string) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceStringUint8(arr []string, acc uint8, f func(el string, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileStringUint8(arr []string, acc uint8, f func(el string, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanStringUint8(arr []string, acc uint8, f func(el string, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Bool(arr []uint16, f func(el uint16) bool) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Bool(arr []uint16, f func(el uint16) bool) map[bool][]uint16 {
	result := make(map[bool][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Bool(arr []uint16, f func(el uint16) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Bool(arr []uint16, acc bool, f func(el uint16, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Bool(arr []uint16, acc bool, f func(el uint16, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Bool(arr []uint16, acc bool, f func(el uint16, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Byte(arr []uint16, f func(el uint16) byte) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Byte(arr []uint16, f func(el uint16) byte) map[byte][]uint16 {
	result := make(map[byte][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Byte(arr []uint16, f func(el uint16) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Byte(arr []uint16, acc byte, f func(el uint16, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Byte(arr []uint16, acc byte, f func(el uint16, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Byte(arr []uint16, acc byte, f func(el uint16, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Float32(arr []uint16, f func(el uint16) float32) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Float32(arr []uint16, f func(el uint16) float32) map[float32][]uint16 {
	result := make(map[float32][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Float32(arr []uint16, f func(el uint16) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Float32(arr []uint16, acc float32, f func(el uint16, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Float32(arr []uint16, acc float32, f func(el uint16, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Float32(arr []uint16, acc float32, f func(el uint16, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Float64(arr []uint16, f func(el uint16) float64) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Float64(arr []uint16, f func(el uint16) float64) map[float64][]uint16 {
	result := make(map[float64][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Float64(arr []uint16, f func(el uint16) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Float64(arr []uint16, acc float64, f func(el uint16, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Float64(arr []uint16, acc float64, f func(el uint16, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Float64(arr []uint16, acc float64, f func(el uint16, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Int(arr []uint16, f func(el uint16) int) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Int(arr []uint16, f func(el uint16) int) map[int][]uint16 {
	result := make(map[int][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Int(arr []uint16, f func(el uint16) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Int(arr []uint16, acc int, f func(el uint16, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Int(arr []uint16, acc int, f func(el uint16, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Int(arr []uint16, acc int, f func(el uint16, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Int16(arr []uint16, f func(el uint16) int16) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Int16(arr []uint16, f func(el uint16) int16) map[int16][]uint16 {
	result := make(map[int16][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Int16(arr []uint16, f func(el uint16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Int16(arr []uint16, acc int16, f func(el uint16, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Int16(arr []uint16, acc int16, f func(el uint16, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Int16(arr []uint16, acc int16, f func(el uint16, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Int32(arr []uint16, f func(el uint16) int32) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Int32(arr []uint16, f func(el uint16) int32) map[int32][]uint16 {
	result := make(map[int32][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Int32(arr []uint16, f func(el uint16) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Int32(arr []uint16, acc int32, f func(el uint16, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Int32(arr []uint16, acc int32, f func(el uint16, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Int32(arr []uint16, acc int32, f func(el uint16, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Int64(arr []uint16, f func(el uint16) int64) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Int64(arr []uint16, f func(el uint16) int64) map[int64][]uint16 {
	result := make(map[int64][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Int64(arr []uint16, f func(el uint16) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Int64(arr []uint16, acc int64, f func(el uint16, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Int64(arr []uint16, acc int64, f func(el uint16, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Int64(arr []uint16, acc int64, f func(el uint16, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Int8(arr []uint16, f func(el uint16) int8) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Int8(arr []uint16, f func(el uint16) int8) map[int8][]uint16 {
	result := make(map[int8][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Int8(arr []uint16, f func(el uint16) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Int8(arr []uint16, acc int8, f func(el uint16, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Int8(arr []uint16, acc int8, f func(el uint16, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Int8(arr []uint16, acc int8, f func(el uint16, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Interface(arr []uint16, f func(el uint16) interface{}) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Interface(arr []uint16, f func(el uint16) interface{}) map[interface{}][]uint16 {
	result := make(map[interface{}][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Interface(arr []uint16, f func(el uint16) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Interface(arr []uint16, acc interface{}, f func(el uint16, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Interface(arr []uint16, acc interface{}, f func(el uint16, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Interface(arr []uint16, acc interface{}, f func(el uint16, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16String(arr []uint16, f func(el uint16) string) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16String(arr []uint16, f func(el uint16) string) map[string][]uint16 {
	result := make(map[string][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16String(arr []uint16, f func(el uint16) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16String(arr []uint16, acc string, f func(el uint16, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16String(arr []uint16, acc string, f func(el uint16, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16String(arr []uint16, acc string, f func(el uint16, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Uint(arr []uint16, f func(el uint16) uint) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Uint(arr []uint16, f func(el uint16) uint) map[uint][]uint16 {
	result := make(map[uint][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Uint(arr []uint16, f func(el uint16) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Uint(arr []uint16, acc uint, f func(el uint16, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Uint(arr []uint16, acc uint, f func(el uint16, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Uint(arr []uint16, acc uint, f func(el uint16, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16(arr []uint16, f func(el uint16) uint16) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16(arr []uint16, f func(el uint16) uint16) map[uint16][]uint16 {
	result := make(map[uint16][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanUint16(arr []uint16, acc uint16, f func(el uint16, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Uint32(arr []uint16, f func(el uint16) uint32) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Uint32(arr []uint16, f func(el uint16) uint32) map[uint32][]uint16 {
	result := make(map[uint32][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Uint32(arr []uint16, f func(el uint16) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Uint32(arr []uint16, acc uint32, f func(el uint16, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Uint32(arr []uint16, acc uint32, f func(el uint16, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Uint32(arr []uint16, acc uint32, f func(el uint16, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Uint64(arr []uint16, f func(el uint16) uint64) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Uint64(arr []uint16, f func(el uint16) uint64) map[uint64][]uint16 {
	result := make(map[uint64][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Uint64(arr []uint16, f func(el uint16) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Uint64(arr []uint16, acc uint64, f func(el uint16, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Uint64(arr []uint16, acc uint64, f func(el uint16, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Uint64(arr []uint16, acc uint64, f func(el uint16, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint16Uint8(arr []uint16, f func(el uint16) uint8) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint16Uint8(arr []uint16, f func(el uint16) uint8) map[uint8][]uint16 {
	result := make(map[uint8][]uint16)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint16Uint8(arr []uint16, f func(el uint16) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint16Uint8(arr []uint16, acc uint8, f func(el uint16, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint16Uint8(arr []uint16, acc uint8, f func(el uint16, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint16Uint8(arr []uint16, acc uint8, f func(el uint16, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Bool(arr []uint32, f func(el uint32) bool) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Bool(arr []uint32, f func(el uint32) bool) map[bool][]uint32 {
	result := make(map[bool][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Bool(arr []uint32, f func(el uint32) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Bool(arr []uint32, acc bool, f func(el uint32, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Bool(arr []uint32, acc bool, f func(el uint32, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Bool(arr []uint32, acc bool, f func(el uint32, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Byte(arr []uint32, f func(el uint32) byte) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Byte(arr []uint32, f func(el uint32) byte) map[byte][]uint32 {
	result := make(map[byte][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Byte(arr []uint32, f func(el uint32) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Byte(arr []uint32, acc byte, f func(el uint32, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Byte(arr []uint32, acc byte, f func(el uint32, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Byte(arr []uint32, acc byte, f func(el uint32, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Float32(arr []uint32, f func(el uint32) float32) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Float32(arr []uint32, f func(el uint32) float32) map[float32][]uint32 {
	result := make(map[float32][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Float32(arr []uint32, f func(el uint32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Float32(arr []uint32, acc float32, f func(el uint32, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Float32(arr []uint32, acc float32, f func(el uint32, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Float32(arr []uint32, acc float32, f func(el uint32, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Float64(arr []uint32, f func(el uint32) float64) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Float64(arr []uint32, f func(el uint32) float64) map[float64][]uint32 {
	result := make(map[float64][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Float64(arr []uint32, f func(el uint32) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Float64(arr []uint32, acc float64, f func(el uint32, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Float64(arr []uint32, acc float64, f func(el uint32, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Float64(arr []uint32, acc float64, f func(el uint32, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Int(arr []uint32, f func(el uint32) int) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Int(arr []uint32, f func(el uint32) int) map[int][]uint32 {
	result := make(map[int][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Int(arr []uint32, f func(el uint32) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Int(arr []uint32, acc int, f func(el uint32, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Int(arr []uint32, acc int, f func(el uint32, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Int(arr []uint32, acc int, f func(el uint32, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Int16(arr []uint32, f func(el uint32) int16) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Int16(arr []uint32, f func(el uint32) int16) map[int16][]uint32 {
	result := make(map[int16][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Int16(arr []uint32, f func(el uint32) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Int16(arr []uint32, acc int16, f func(el uint32, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Int16(arr []uint32, acc int16, f func(el uint32, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Int16(arr []uint32, acc int16, f func(el uint32, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Int32(arr []uint32, f func(el uint32) int32) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Int32(arr []uint32, f func(el uint32) int32) map[int32][]uint32 {
	result := make(map[int32][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Int32(arr []uint32, f func(el uint32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Int32(arr []uint32, acc int32, f func(el uint32, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Int32(arr []uint32, acc int32, f func(el uint32, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Int32(arr []uint32, acc int32, f func(el uint32, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Int64(arr []uint32, f func(el uint32) int64) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Int64(arr []uint32, f func(el uint32) int64) map[int64][]uint32 {
	result := make(map[int64][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Int64(arr []uint32, f func(el uint32) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Int64(arr []uint32, acc int64, f func(el uint32, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Int64(arr []uint32, acc int64, f func(el uint32, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Int64(arr []uint32, acc int64, f func(el uint32, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Int8(arr []uint32, f func(el uint32) int8) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Int8(arr []uint32, f func(el uint32) int8) map[int8][]uint32 {
	result := make(map[int8][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Int8(arr []uint32, f func(el uint32) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Int8(arr []uint32, acc int8, f func(el uint32, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Int8(arr []uint32, acc int8, f func(el uint32, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Int8(arr []uint32, acc int8, f func(el uint32, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Interface(arr []uint32, f func(el uint32) interface{}) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Interface(arr []uint32, f func(el uint32) interface{}) map[interface{}][]uint32 {
	result := make(map[interface{}][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Interface(arr []uint32, f func(el uint32) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Interface(arr []uint32, acc interface{}, f func(el uint32, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Interface(arr []uint32, acc interface{}, f func(el uint32, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Interface(arr []uint32, acc interface{}, f func(el uint32, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32String(arr []uint32, f func(el uint32) string) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32String(arr []uint32, f func(el uint32) string) map[string][]uint32 {
	result := make(map[string][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32String(arr []uint32, f func(el uint32) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32String(arr []uint32, acc string, f func(el uint32, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32String(arr []uint32, acc string, f func(el uint32, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32String(arr []uint32, acc string, f func(el uint32, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Uint(arr []uint32, f func(el uint32) uint) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Uint(arr []uint32, f func(el uint32) uint) map[uint][]uint32 {
	result := make(map[uint][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Uint(arr []uint32, f func(el uint32) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Uint(arr []uint32, acc uint, f func(el uint32, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Uint(arr []uint32, acc uint, f func(el uint32, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Uint(arr []uint32, acc uint, f func(el uint32, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Uint16(arr []uint32, f func(el uint32) uint16) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Uint16(arr []uint32, f func(el uint32) uint16) map[uint16][]uint32 {
	result := make(map[uint16][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Uint16(arr []uint32, f func(el uint32) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Uint16(arr []uint32, acc uint16, f func(el uint32, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Uint16(arr []uint32, acc uint16, f func(el uint32, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Uint16(arr []uint32, acc uint16, f func(el uint32, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32(arr []uint32, f func(el uint32) uint32) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32(arr []uint32, f func(el uint32) uint32) map[uint32][]uint32 {
	result := make(map[uint32][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanUint32(arr []uint32, acc uint32, f func(el uint32, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Uint64(arr []uint32, f func(el uint32) uint64) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Uint64(arr []uint32, f func(el uint32) uint64) map[uint64][]uint32 {
	result := make(map[uint64][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Uint64(arr []uint32, f func(el uint32) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Uint64(arr []uint32, acc uint64, f func(el uint32, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Uint64(arr []uint32, acc uint64, f func(el uint32, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Uint64(arr []uint32, acc uint64, f func(el uint32, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint32Uint8(arr []uint32, f func(el uint32) uint8) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint32Uint8(arr []uint32, f func(el uint32) uint8) map[uint8][]uint32 {
	result := make(map[uint8][]uint32)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint32Uint8(arr []uint32, f func(el uint32) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint32Uint8(arr []uint32, acc uint8, f func(el uint32, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint32Uint8(arr []uint32, acc uint8, f func(el uint32, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint32Uint8(arr []uint32, acc uint8, f func(el uint32, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Bool(arr []uint64, f func(el uint64) bool) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Bool(arr []uint64, f func(el uint64) bool) map[bool][]uint64 {
	result := make(map[bool][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Bool(arr []uint64, f func(el uint64) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Bool(arr []uint64, acc bool, f func(el uint64, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Bool(arr []uint64, acc bool, f func(el uint64, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Bool(arr []uint64, acc bool, f func(el uint64, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Byte(arr []uint64, f func(el uint64) byte) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Byte(arr []uint64, f func(el uint64) byte) map[byte][]uint64 {
	result := make(map[byte][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Byte(arr []uint64, f func(el uint64) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Byte(arr []uint64, acc byte, f func(el uint64, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Byte(arr []uint64, acc byte, f func(el uint64, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Byte(arr []uint64, acc byte, f func(el uint64, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Float32(arr []uint64, f func(el uint64) float32) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Float32(arr []uint64, f func(el uint64) float32) map[float32][]uint64 {
	result := make(map[float32][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Float32(arr []uint64, f func(el uint64) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Float32(arr []uint64, acc float32, f func(el uint64, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Float32(arr []uint64, acc float32, f func(el uint64, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Float32(arr []uint64, acc float32, f func(el uint64, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Float64(arr []uint64, f func(el uint64) float64) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Float64(arr []uint64, f func(el uint64) float64) map[float64][]uint64 {
	result := make(map[float64][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Float64(arr []uint64, f func(el uint64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Float64(arr []uint64, acc float64, f func(el uint64, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Float64(arr []uint64, acc float64, f func(el uint64, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Float64(arr []uint64, acc float64, f func(el uint64, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Int(arr []uint64, f func(el uint64) int) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Int(arr []uint64, f func(el uint64) int) map[int][]uint64 {
	result := make(map[int][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Int(arr []uint64, f func(el uint64) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Int(arr []uint64, acc int, f func(el uint64, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Int(arr []uint64, acc int, f func(el uint64, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Int(arr []uint64, acc int, f func(el uint64, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Int16(arr []uint64, f func(el uint64) int16) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Int16(arr []uint64, f func(el uint64) int16) map[int16][]uint64 {
	result := make(map[int16][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Int16(arr []uint64, f func(el uint64) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Int16(arr []uint64, acc int16, f func(el uint64, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Int16(arr []uint64, acc int16, f func(el uint64, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Int16(arr []uint64, acc int16, f func(el uint64, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Int32(arr []uint64, f func(el uint64) int32) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Int32(arr []uint64, f func(el uint64) int32) map[int32][]uint64 {
	result := make(map[int32][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Int32(arr []uint64, f func(el uint64) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Int32(arr []uint64, acc int32, f func(el uint64, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Int32(arr []uint64, acc int32, f func(el uint64, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Int32(arr []uint64, acc int32, f func(el uint64, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Int64(arr []uint64, f func(el uint64) int64) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Int64(arr []uint64, f func(el uint64) int64) map[int64][]uint64 {
	result := make(map[int64][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Int64(arr []uint64, f func(el uint64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Int64(arr []uint64, acc int64, f func(el uint64, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Int64(arr []uint64, acc int64, f func(el uint64, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Int64(arr []uint64, acc int64, f func(el uint64, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Int8(arr []uint64, f func(el uint64) int8) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Int8(arr []uint64, f func(el uint64) int8) map[int8][]uint64 {
	result := make(map[int8][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Int8(arr []uint64, f func(el uint64) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Int8(arr []uint64, acc int8, f func(el uint64, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Int8(arr []uint64, acc int8, f func(el uint64, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Int8(arr []uint64, acc int8, f func(el uint64, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Interface(arr []uint64, f func(el uint64) interface{}) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Interface(arr []uint64, f func(el uint64) interface{}) map[interface{}][]uint64 {
	result := make(map[interface{}][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Interface(arr []uint64, f func(el uint64) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Interface(arr []uint64, acc interface{}, f func(el uint64, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Interface(arr []uint64, acc interface{}, f func(el uint64, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Interface(arr []uint64, acc interface{}, f func(el uint64, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64String(arr []uint64, f func(el uint64) string) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64String(arr []uint64, f func(el uint64) string) map[string][]uint64 {
	result := make(map[string][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64String(arr []uint64, f func(el uint64) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64String(arr []uint64, acc string, f func(el uint64, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64String(arr []uint64, acc string, f func(el uint64, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64String(arr []uint64, acc string, f func(el uint64, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Uint(arr []uint64, f func(el uint64) uint) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Uint(arr []uint64, f func(el uint64) uint) map[uint][]uint64 {
	result := make(map[uint][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Uint(arr []uint64, f func(el uint64) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Uint(arr []uint64, acc uint, f func(el uint64, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Uint(arr []uint64, acc uint, f func(el uint64, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Uint(arr []uint64, acc uint, f func(el uint64, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Uint16(arr []uint64, f func(el uint64) uint16) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Uint16(arr []uint64, f func(el uint64) uint16) map[uint16][]uint64 {
	result := make(map[uint16][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Uint16(arr []uint64, f func(el uint64) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Uint16(arr []uint64, acc uint16, f func(el uint64, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Uint16(arr []uint64, acc uint16, f func(el uint64, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Uint16(arr []uint64, acc uint16, f func(el uint64, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Uint32(arr []uint64, f func(el uint64) uint32) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Uint32(arr []uint64, f func(el uint64) uint32) map[uint32][]uint64 {
	result := make(map[uint32][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Uint32(arr []uint64, f func(el uint64) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Uint32(arr []uint64, acc uint32, f func(el uint64, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Uint32(arr []uint64, acc uint32, f func(el uint64, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Uint32(arr []uint64, acc uint32, f func(el uint64, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64(arr []uint64, f func(el uint64) uint64) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64(arr []uint64, f func(el uint64) uint64) map[uint64][]uint64 {
	result := make(map[uint64][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanUint64(arr []uint64, acc uint64, f func(el uint64, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint64Uint8(arr []uint64, f func(el uint64) uint8) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint64Uint8(arr []uint64, f func(el uint64) uint8) map[uint8][]uint64 {
	result := make(map[uint8][]uint64)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint64Uint8(arr []uint64, f func(el uint64) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint64Uint8(arr []uint64, acc uint8, f func(el uint64, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint64Uint8(arr []uint64, acc uint8, f func(el uint64, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint64Uint8(arr []uint64, acc uint8, f func(el uint64, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Bool(arr []uint8, f func(el uint8) bool) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Bool(arr []uint8, f func(el uint8) bool) map[bool][]uint8 {
	result := make(map[bool][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Bool(arr []uint8, f func(el uint8) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Bool(arr []uint8, acc bool, f func(el uint8, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Bool(arr []uint8, acc bool, f func(el uint8, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Bool(arr []uint8, acc bool, f func(el uint8, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Byte(arr []uint8, f func(el uint8) byte) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Byte(arr []uint8, f func(el uint8) byte) map[byte][]uint8 {
	result := make(map[byte][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Byte(arr []uint8, f func(el uint8) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Byte(arr []uint8, acc byte, f func(el uint8, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Byte(arr []uint8, acc byte, f func(el uint8, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Byte(arr []uint8, acc byte, f func(el uint8, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Float32(arr []uint8, f func(el uint8) float32) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Float32(arr []uint8, f func(el uint8) float32) map[float32][]uint8 {
	result := make(map[float32][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Float32(arr []uint8, f func(el uint8) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Float32(arr []uint8, acc float32, f func(el uint8, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Float32(arr []uint8, acc float32, f func(el uint8, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Float32(arr []uint8, acc float32, f func(el uint8, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Float64(arr []uint8, f func(el uint8) float64) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Float64(arr []uint8, f func(el uint8) float64) map[float64][]uint8 {
	result := make(map[float64][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Float64(arr []uint8, f func(el uint8) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Float64(arr []uint8, acc float64, f func(el uint8, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Float64(arr []uint8, acc float64, f func(el uint8, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Float64(arr []uint8, acc float64, f func(el uint8, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Int(arr []uint8, f func(el uint8) int) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Int(arr []uint8, f func(el uint8) int) map[int][]uint8 {
	result := make(map[int][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Int(arr []uint8, f func(el uint8) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Int(arr []uint8, acc int, f func(el uint8, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Int(arr []uint8, acc int, f func(el uint8, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Int(arr []uint8, acc int, f func(el uint8, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Int16(arr []uint8, f func(el uint8) int16) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Int16(arr []uint8, f func(el uint8) int16) map[int16][]uint8 {
	result := make(map[int16][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Int16(arr []uint8, f func(el uint8) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Int16(arr []uint8, acc int16, f func(el uint8, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Int16(arr []uint8, acc int16, f func(el uint8, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Int16(arr []uint8, acc int16, f func(el uint8, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Int32(arr []uint8, f func(el uint8) int32) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Int32(arr []uint8, f func(el uint8) int32) map[int32][]uint8 {
	result := make(map[int32][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Int32(arr []uint8, f func(el uint8) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Int32(arr []uint8, acc int32, f func(el uint8, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Int32(arr []uint8, acc int32, f func(el uint8, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Int32(arr []uint8, acc int32, f func(el uint8, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Int64(arr []uint8, f func(el uint8) int64) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Int64(arr []uint8, f func(el uint8) int64) map[int64][]uint8 {
	result := make(map[int64][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Int64(arr []uint8, f func(el uint8) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Int64(arr []uint8, acc int64, f func(el uint8, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Int64(arr []uint8, acc int64, f func(el uint8, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Int64(arr []uint8, acc int64, f func(el uint8, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Int8(arr []uint8, f func(el uint8) int8) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Int8(arr []uint8, f func(el uint8) int8) map[int8][]uint8 {
	result := make(map[int8][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Int8(arr []uint8, f func(el uint8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Int8(arr []uint8, acc int8, f func(el uint8, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Int8(arr []uint8, acc int8, f func(el uint8, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Int8(arr []uint8, acc int8, f func(el uint8, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Interface(arr []uint8, f func(el uint8) interface{}) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Interface(arr []uint8, f func(el uint8) interface{}) map[interface{}][]uint8 {
	result := make(map[interface{}][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Interface(arr []uint8, f func(el uint8) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Interface(arr []uint8, acc interface{}, f func(el uint8, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Interface(arr []uint8, acc interface{}, f func(el uint8, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Interface(arr []uint8, acc interface{}, f func(el uint8, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8String(arr []uint8, f func(el uint8) string) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8String(arr []uint8, f func(el uint8) string) map[string][]uint8 {
	result := make(map[string][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8String(arr []uint8, f func(el uint8) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8String(arr []uint8, acc string, f func(el uint8, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8String(arr []uint8, acc string, f func(el uint8, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8String(arr []uint8, acc string, f func(el uint8, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Uint(arr []uint8, f func(el uint8) uint) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Uint(arr []uint8, f func(el uint8) uint) map[uint][]uint8 {
	result := make(map[uint][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Uint(arr []uint8, f func(el uint8) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Uint(arr []uint8, acc uint, f func(el uint8, acc uint) uint) uint {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Uint(arr []uint8, acc uint, f func(el uint8, acc uint) (uint, error)) (uint, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Uint(arr []uint8, acc uint, f func(el uint8, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Uint16(arr []uint8, f func(el uint8) uint16) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Uint16(arr []uint8, f func(el uint8) uint16) map[uint16][]uint8 {
	result := make(map[uint16][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Uint16(arr []uint8, f func(el uint8) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Uint16(arr []uint8, acc uint16, f func(el uint8, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Uint16(arr []uint8, acc uint16, f func(el uint8, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Uint16(arr []uint8, acc uint16, f func(el uint8, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Uint32(arr []uint8, f func(el uint8) uint32) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Uint32(arr []uint8, f func(el uint8) uint32) map[uint32][]uint8 {
	result := make(map[uint32][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Uint32(arr []uint8, f func(el uint8) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Uint32(arr []uint8, acc uint32, f func(el uint8, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Uint32(arr []uint8, acc uint32, f func(el uint8, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Uint32(arr []uint8, acc uint32, f func(el uint8, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8Uint64(arr []uint8, f func(el uint8) uint64) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8Uint64(arr []uint8, f func(el uint8) uint64) map[uint64][]uint8 {
	result := make(map[uint64][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUint8Uint64(arr []uint8, f func(el uint8) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUint8Uint64(arr []uint8, acc uint64, f func(el uint8, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUint8Uint64(arr []uint8, acc uint64, f func(el uint8, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUint8Uint64(arr []uint8, acc uint64, f func(el uint8, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint8(arr []uint8, f func(el uint8) uint8) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint8(arr []uint8, f func(el uint8) uint8) map[uint8][]uint8 {
	result := make(map[uint8][]uint8)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanUint8(arr []uint8, acc uint8, f func(el uint8, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintBool(arr []uint, f func(el uint) bool) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintBool(arr []uint, f func(el uint) bool) map[bool][]uint {
	result := make(map[bool][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintBool(arr []uint, f func(el uint) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintBool(arr []uint, acc bool, f func(el uint, acc bool) bool) bool {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintBool(arr []uint, acc bool, f func(el uint, acc bool) (bool, error)) (bool, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintBool(arr []uint, acc bool, f func(el uint, acc bool) bool) []bool {
	result := make([]bool, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintByte(arr []uint, f func(el uint) byte) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintByte(arr []uint, f func(el uint) byte) map[byte][]uint {
	result := make(map[byte][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintByte(arr []uint, f func(el uint) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintByte(arr []uint, acc byte, f func(el uint, acc byte) byte) byte {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintByte(arr []uint, acc byte, f func(el uint, acc byte) (byte, error)) (byte, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintByte(arr []uint, acc byte, f func(el uint, acc byte) byte) []byte {
	result := make([]byte, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintFloat32(arr []uint, f func(el uint) float32) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintFloat32(arr []uint, f func(el uint) float32) map[float32][]uint {
	result := make(map[float32][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintFloat32(arr []uint, f func(el uint) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintFloat32(arr []uint, acc float32, f func(el uint, acc float32) float32) float32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintFloat32(arr []uint, acc float32, f func(el uint, acc float32) (float32, error)) (float32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintFloat32(arr []uint, acc float32, f func(el uint, acc float32) float32) []float32 {
	result := make([]float32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintFloat64(arr []uint, f func(el uint) float64) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintFloat64(arr []uint, f func(el uint) float64) map[float64][]uint {
	result := make(map[float64][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintFloat64(arr []uint, f func(el uint) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintFloat64(arr []uint, acc float64, f func(el uint, acc float64) float64) float64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintFloat64(arr []uint, acc float64, f func(el uint, acc float64) (float64, error)) (float64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintFloat64(arr []uint, acc float64, f func(el uint, acc float64) float64) []float64 {
	result := make([]float64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintInt(arr []uint, f func(el uint) int) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintInt(arr []uint, f func(el uint) int) map[int][]uint {
	result := make(map[int][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintInt(arr []uint, f func(el uint) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintInt(arr []uint, acc int, f func(el uint, acc int) int) int {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintInt(arr []uint, acc int, f func(el uint, acc int) (int, error)) (int, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintInt(arr []uint, acc int, f func(el uint, acc int) int) []int {
	result := make([]int, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintInt16(arr []uint, f func(el uint) int16) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintInt16(arr []uint, f func(el uint) int16) map[int16][]uint {
	result := make(map[int16][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintInt16(arr []uint, f func(el uint) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintInt16(arr []uint, acc int16, f func(el uint, acc int16) int16) int16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintInt16(arr []uint, acc int16, f func(el uint, acc int16) (int16, error)) (int16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintInt16(arr []uint, acc int16, f func(el uint, acc int16) int16) []int16 {
	result := make([]int16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintInt32(arr []uint, f func(el uint) int32) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintInt32(arr []uint, f func(el uint) int32) map[int32][]uint {
	result := make(map[int32][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintInt32(arr []uint, f func(el uint) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintInt32(arr []uint, acc int32, f func(el uint, acc int32) int32) int32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintInt32(arr []uint, acc int32, f func(el uint, acc int32) (int32, error)) (int32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintInt32(arr []uint, acc int32, f func(el uint, acc int32) int32) []int32 {
	result := make([]int32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintInt64(arr []uint, f func(el uint) int64) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintInt64(arr []uint, f func(el uint) int64) map[int64][]uint {
	result := make(map[int64][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintInt64(arr []uint, f func(el uint) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintInt64(arr []uint, acc int64, f func(el uint, acc int64) int64) int64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintInt64(arr []uint, acc int64, f func(el uint, acc int64) (int64, error)) (int64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintInt64(arr []uint, acc int64, f func(el uint, acc int64) int64) []int64 {
	result := make([]int64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintInt8(arr []uint, f func(el uint) int8) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintInt8(arr []uint, f func(el uint) int8) map[int8][]uint {
	result := make(map[int8][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintInt8(arr []uint, f func(el uint) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintInt8(arr []uint, acc int8, f func(el uint, acc int8) int8) int8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintInt8(arr []uint, acc int8, f func(el uint, acc int8) (int8, error)) (int8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintInt8(arr []uint, acc int8, f func(el uint, acc int8) int8) []int8 {
	result := make([]int8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintInterface(arr []uint, f func(el uint) interface{}) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintInterface(arr []uint, f func(el uint) interface{}) map[interface{}][]uint {
	result := make(map[interface{}][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintInterface(arr []uint, f func(el uint) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintInterface(arr []uint, acc interface{}, f func(el uint, acc interface{}) interface{}) interface{} {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintInterface(arr []uint, acc interface{}, f func(el uint, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintInterface(arr []uint, acc interface{}, f func(el uint, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintString(arr []uint, f func(el uint) string) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintString(arr []uint, f func(el uint) string) map[string][]uint {
	result := make(map[string][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintString(arr []uint, f func(el uint) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintString(arr []uint, acc string, f func(el uint, acc string) string) string {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintString(arr []uint, acc string, f func(el uint, acc string) (string, error)) (string, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintString(arr []uint, acc string, f func(el uint, acc string) string) []string {
	result := make([]string, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUint(arr []uint, f func(el uint) uint) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUint(arr []uint, f func(el uint) uint) map[uint][]uint {
	result := make(map[uint][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
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

// Scan is like Reduce2, but returns slice of f results
func ScanUint(arr []uint, acc uint, f func(el uint, acc uint) uint) []uint {
	result := make([]uint, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintUint16(arr []uint, f func(el uint) uint16) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintUint16(arr []uint, f func(el uint) uint16) map[uint16][]uint {
	result := make(map[uint16][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintUint16(arr []uint, f func(el uint) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintUint16(arr []uint, acc uint16, f func(el uint, acc uint16) uint16) uint16 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintUint16(arr []uint, acc uint16, f func(el uint, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintUint16(arr []uint, acc uint16, f func(el uint, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintUint32(arr []uint, f func(el uint) uint32) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintUint32(arr []uint, f func(el uint) uint32) map[uint32][]uint {
	result := make(map[uint32][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintUint32(arr []uint, f func(el uint) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintUint32(arr []uint, acc uint32, f func(el uint, acc uint32) uint32) uint32 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintUint32(arr []uint, acc uint32, f func(el uint, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintUint32(arr []uint, acc uint32, f func(el uint, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintUint64(arr []uint, f func(el uint) uint64) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintUint64(arr []uint, f func(el uint) uint64) map[uint64][]uint {
	result := make(map[uint64][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintUint64(arr []uint, f func(el uint) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintUint64(arr []uint, acc uint64, f func(el uint, acc uint64) uint64) uint64 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintUint64(arr []uint, acc uint64, f func(el uint, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintUint64(arr []uint, acc uint64, f func(el uint, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkByUintUint8(arr []uint, f func(el uint) uint8) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(arr[0])
	chunk = append(chunk, arr[0])

	for _, el := range arr[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// GroupBy groups element from array by value returned by f
func GroupByUintUint8(arr []uint, f func(el uint) uint8) map[uint8][]uint {
	result := make(map[uint8][]uint)
	for _, el := range arr {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func MapUintUint8(arr []uint, f func(el uint) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func ReduceUintUint8(arr []uint, acc uint8, f func(el uint, acc uint8) uint8) uint8 {
	for _, el := range arr {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhileUintUint8(arr []uint, acc uint8, f func(el uint, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range arr {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce2, but returns slice of f results
func ScanUintUint8(arr []uint, acc uint8, f func(el uint, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(arr))
	for _, el := range arr {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}
