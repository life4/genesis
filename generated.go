package genesis

import (
	"math/rand"
	"sort"
	"sync"
	"time")

type ChannelBool struct {
	data chan bool
}

type AsyncSliceBool struct {
	data    []bool
	workers int
}

type SequenceBool struct {
	data chan bool
}

type SliceBool struct {
	data []bool
}

type SlicesBool struct {
	data [][]bool
}

func (c ChannelBool) Any(f func(el bool) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelBool) All(f func(el bool) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelBool) ChunkEvery(count int) chan []bool {
	chunks := make(chan []bool, 1)
	go func() {
		chunk := make([]bool, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]bool, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelBool) Count(el bool) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelBool) Drop(n int) chan bool {
	result := make(chan bool, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) Each(f func(el bool)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelBool) Filter(f func(el bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapBool(f func(el bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapByte(f func(el bool) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapString(f func(el bool) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapFloat32(f func(el bool) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapFloat64(f func(el bool) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapInt(f func(el bool) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapInt8(f func(el bool) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapInt16(f func(el bool) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapInt32(f func(el bool) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapInt64(f func(el bool) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapUint(f func(el bool) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapUint8(f func(el bool) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapUint16(f func(el bool) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapUint32(f func(el bool) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapUint64(f func(el bool) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) MapInterface(f func(el bool) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ReduceBool(acc bool, f func(el bool, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceByte(acc byte, f func(el bool, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceString(acc string, f func(el bool, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceFloat32(acc float32, f func(el bool, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceFloat64(acc float64, f func(el bool, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceInt(acc int, f func(el bool, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceInt8(acc int8, f func(el bool, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceInt16(acc int16, f func(el bool, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceInt32(acc int32, f func(el bool, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceInt64(acc int64, f func(el bool, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceUint(acc uint, f func(el bool, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceUint8(acc uint8, f func(el bool, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceUint16(acc uint16, f func(el bool, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceUint32(acc uint32, f func(el bool, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceUint64(acc uint64, f func(el bool, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ReduceInterface(acc interface{}, f func(el bool, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelBool) ScanBool(acc bool, f func(el bool, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanByte(acc byte, f func(el bool, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanString(acc string, f func(el bool, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanFloat32(acc float32, f func(el bool, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanFloat64(acc float64, f func(el bool, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanInt(acc int, f func(el bool, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanInt8(acc int8, f func(el bool, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanInt16(acc int16, f func(el bool, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanInt32(acc int32, f func(el bool, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanInt64(acc int64, f func(el bool, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanUint(acc uint, f func(el bool, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanUint8(acc uint8, f func(el bool, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanUint16(acc uint16, f func(el bool, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanUint32(acc uint32, f func(el bool, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanUint64(acc uint64, f func(el bool, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) ScanInterface(acc interface{}, f func(el bool, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelBool) Take(n int) []bool {
	result := make([]bool, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelBool) Tee(count int) []chan bool {
	channels := make([]chan bool, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan bool, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan bool) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelBool) ToSlice() []bool {
	result := make([]bool, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceBool) Each(f func(el bool)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceBool) Filter(f func(el bool) bool) []bool {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]bool, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceBool) MapBool(f func(el bool) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapByte(f func(el bool) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapString(f func(el bool) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapFloat32(f func(el bool) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapFloat64(f func(el bool) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapInt(f func(el bool) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapInt8(f func(el bool) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapInt16(f func(el bool) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapInt32(f func(el bool) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapInt64(f func(el bool) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapUint(f func(el bool) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapUint8(f func(el bool) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapUint16(f func(el bool) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapUint32(f func(el bool) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapUint64(f func(el bool) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceBool) MapInterface(f func(el bool) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceBool) Repeat(val bool) chan bool {
	c := make(chan bool, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceBool) Any(f func(el bool) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceBool) All(f func(el bool) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceBool) ChunkByBool(f func(el bool) bool) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByByte(f func(el bool) byte) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByString(f func(el bool) string) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByFloat32(f func(el bool) float32) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByFloat64(f func(el bool) float64) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByInt(f func(el bool) int) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByInt8(f func(el bool) int8) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByInt16(f func(el bool) int16) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByInt32(f func(el bool) int32) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByInt64(f func(el bool) int64) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByUint(f func(el bool) uint) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByUint8(f func(el bool) uint8) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByUint16(f func(el bool) uint16) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByUint32(f func(el bool) uint32) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByUint64(f func(el bool) uint64) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkByInterface(f func(el bool) interface{}) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceBool) ChunkEvery(count int) [][]bool {
	chunks := make([][]bool, 0)
	chunk := make([]bool, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]bool, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceBool) Contains(el bool) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceBool) Count(el bool) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceBool) Cycle() chan bool {
	c := make(chan bool, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceBool) Dedup() []bool {
	result := make([]bool, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceBool) DedupByBool(f func(el bool) bool) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByByte(f func(el bool) byte) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByString(f func(el bool) string) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByFloat32(f func(el bool) float32) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByFloat64(f func(el bool) float64) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByInt(f func(el bool) int) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByInt8(f func(el bool) int8) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByInt16(f func(el bool) int16) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByInt32(f func(el bool) int32) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByInt64(f func(el bool) int64) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByUint(f func(el bool) uint) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByUint8(f func(el bool) uint8) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByUint16(f func(el bool) uint16) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByUint32(f func(el bool) uint32) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByUint64(f func(el bool) uint64) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DedupByInterface(f func(el bool) interface{}) []bool {
	result := make([]bool, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceBool) DropEvery(nth int) []bool {
	result := make([]bool, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceBool) DropWhile(f func(arr bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceBool) Each(f func(el bool)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceBool) Filter(f func(el bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceBool) Find(def bool, f func(el bool) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceBool) FindIndex(f func(el bool) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceBool) GroupByBool(f func(el bool) bool) map[bool][]bool {
	result := make(map[bool][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByByte(f func(el bool) byte) map[byte][]bool {
	result := make(map[byte][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByString(f func(el bool) string) map[string][]bool {
	result := make(map[string][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByFloat32(f func(el bool) float32) map[float32][]bool {
	result := make(map[float32][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByFloat64(f func(el bool) float64) map[float64][]bool {
	result := make(map[float64][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByInt(f func(el bool) int) map[int][]bool {
	result := make(map[int][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByInt8(f func(el bool) int8) map[int8][]bool {
	result := make(map[int8][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByInt16(f func(el bool) int16) map[int16][]bool {
	result := make(map[int16][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByInt32(f func(el bool) int32) map[int32][]bool {
	result := make(map[int32][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByInt64(f func(el bool) int64) map[int64][]bool {
	result := make(map[int64][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByUint(f func(el bool) uint) map[uint][]bool {
	result := make(map[uint][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByUint8(f func(el bool) uint8) map[uint8][]bool {
	result := make(map[uint8][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByUint16(f func(el bool) uint16) map[uint16][]bool {
	result := make(map[uint16][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByUint32(f func(el bool) uint32) map[uint32][]bool {
	result := make(map[uint32][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByUint64(f func(el bool) uint64) map[uint64][]bool {
	result := make(map[uint64][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) GroupByInterface(f func(el bool) interface{}) map[interface{}][]bool {
	result := make(map[interface{}][]bool)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]bool, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceBool) Intersperse(el bool) []bool {
	result := make([]bool, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceBool) MapBool(f func(el bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapByte(f func(el bool) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapString(f func(el bool) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapFloat32(f func(el bool) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapFloat64(f func(el bool) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapInt(f func(el bool) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapInt8(f func(el bool) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapInt16(f func(el bool) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapInt32(f func(el bool) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapInt64(f func(el bool) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapUint(f func(el bool) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapUint8(f func(el bool) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapUint16(f func(el bool) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapUint32(f func(el bool) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapUint64(f func(el bool) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) MapInterface(f func(el bool) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceBool) Product(repeat int) chan []bool {
	c := make(chan []bool, 1)
	go s.product(c, repeat, []bool{}, 0)
	return c
}

func (s SliceBool) product(c chan []bool, repeat int, left []bool, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]bool, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]bool, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceBool) ReduceBool(acc bool, f func(el bool, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceByte(acc byte, f func(el bool, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceString(acc string, f func(el bool, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceFloat32(acc float32, f func(el bool, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceFloat64(acc float64, f func(el bool, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceInt(acc int, f func(el bool, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceInt8(acc int8, f func(el bool, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceInt16(acc int16, f func(el bool, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceInt32(acc int32, f func(el bool, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceInt64(acc int64, f func(el bool, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceUint(acc uint, f func(el bool, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceUint8(acc uint8, f func(el bool, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceUint16(acc uint16, f func(el bool, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceUint32(acc uint32, f func(el bool, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceUint64(acc uint64, f func(el bool, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceInterface(acc interface{}, f func(el bool, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceBool) ReduceWhileBool(acc bool, f func(el bool, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileByte(acc byte, f func(el bool, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileString(acc string, f func(el bool, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileFloat32(acc float32, f func(el bool, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileFloat64(acc float64, f func(el bool, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileInt(acc int, f func(el bool, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileInt8(acc int8, f func(el bool, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileInt16(acc int16, f func(el bool, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileInt32(acc int32, f func(el bool, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileInt64(acc int64, f func(el bool, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileUint(acc uint, f func(el bool, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileUint8(acc uint8, f func(el bool, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileUint16(acc uint16, f func(el bool, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileUint32(acc uint32, f func(el bool, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileUint64(acc uint64, f func(el bool, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) ReduceWhileInterface(acc interface{}, f func(el bool, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceBool) Reverse() []bool {
	result := make([]bool, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceBool) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceBool) ScanBool(acc bool, f func(el bool, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanByte(acc byte, f func(el bool, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanString(acc string, f func(el bool, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanFloat32(acc float32, f func(el bool, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanFloat64(acc float64, f func(el bool, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanInt(acc int, f func(el bool, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanInt8(acc int8, f func(el bool, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanInt16(acc int16, f func(el bool, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanInt32(acc int32, f func(el bool, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanInt64(acc int64, f func(el bool, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanUint(acc uint, f func(el bool, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanUint8(acc uint8, f func(el bool, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanUint16(acc uint16, f func(el bool, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanUint32(acc uint32, f func(el bool, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanUint64(acc uint64, f func(el bool, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) ScanInterface(acc interface{}, f func(el bool, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceBool) Shuffle() []bool {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceBool) Split(sep bool) [][]bool {
	result := make([][]bool, 0)
	curr := make([]bool, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceBool) StartsWith(prefix []bool) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceBool) TakeWhile(f func(el bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceBool) Uniq() []bool {
	added := make(map[bool]struct{})
	nothing := struct{}{}
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceBool) Window(size int) [][]bool {
	result := make([][]bool, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesBool) Concat() []bool {
	result := make([]bool, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesBool) Product() chan []bool {
	c := make(chan []bool, 1)
	go s.product(c, []bool{}, 0)
	return c
}

func (s SlicesBool) product(c chan []bool, left []bool, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]bool, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]bool, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesBool) Zip() [][]bool {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]bool, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]bool, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelByte struct {
	data chan byte
}

type AsyncSliceByte struct {
	data    []byte
	workers int
}

type SequenceByte struct {
	data chan byte
}

type SliceByte struct {
	data []byte
}

type SlicesByte struct {
	data [][]byte
}

func (c ChannelByte) Any(f func(el byte) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelByte) All(f func(el byte) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelByte) ChunkEvery(count int) chan []byte {
	chunks := make(chan []byte, 1)
	go func() {
		chunk := make([]byte, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]byte, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelByte) Count(el byte) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelByte) Drop(n int) chan byte {
	result := make(chan byte, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) Each(f func(el byte)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelByte) Filter(f func(el byte) bool) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapBool(f func(el byte) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapByte(f func(el byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapString(f func(el byte) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapFloat32(f func(el byte) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapFloat64(f func(el byte) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapInt(f func(el byte) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapInt8(f func(el byte) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapInt16(f func(el byte) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapInt32(f func(el byte) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapInt64(f func(el byte) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapUint(f func(el byte) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapUint8(f func(el byte) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapUint16(f func(el byte) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapUint32(f func(el byte) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapUint64(f func(el byte) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) MapInterface(f func(el byte) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) Max() byte {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelByte) Min() byte {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelByte) ReduceBool(acc bool, f func(el byte, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceByte(acc byte, f func(el byte, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceString(acc string, f func(el byte, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceFloat32(acc float32, f func(el byte, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceFloat64(acc float64, f func(el byte, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceInt(acc int, f func(el byte, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceInt8(acc int8, f func(el byte, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceInt16(acc int16, f func(el byte, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceInt32(acc int32, f func(el byte, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceInt64(acc int64, f func(el byte, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceUint(acc uint, f func(el byte, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceUint8(acc uint8, f func(el byte, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceUint16(acc uint16, f func(el byte, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceUint32(acc uint32, f func(el byte, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceUint64(acc uint64, f func(el byte, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ReduceInterface(acc interface{}, f func(el byte, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelByte) ScanBool(acc bool, f func(el byte, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanByte(acc byte, f func(el byte, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanString(acc string, f func(el byte, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanFloat32(acc float32, f func(el byte, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanFloat64(acc float64, f func(el byte, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanInt(acc int, f func(el byte, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanInt8(acc int8, f func(el byte, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanInt16(acc int16, f func(el byte, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanInt32(acc int32, f func(el byte, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanInt64(acc int64, f func(el byte, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanUint(acc uint, f func(el byte, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanUint8(acc uint8, f func(el byte, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanUint16(acc uint16, f func(el byte, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanUint32(acc uint32, f func(el byte, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanUint64(acc uint64, f func(el byte, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) ScanInterface(acc interface{}, f func(el byte, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelByte) Sum() byte {
	var sum byte
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelByte) Take(n int) []byte {
	result := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelByte) Tee(count int) []chan byte {
	channels := make([]chan byte, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan byte, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan byte) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelByte) ToSlice() []byte {
	result := make([]byte, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceByte) Each(f func(el byte)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceByte) Filter(f func(el byte) bool) []byte {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]byte, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceByte) MapBool(f func(el byte) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapByte(f func(el byte) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapString(f func(el byte) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapFloat32(f func(el byte) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapFloat64(f func(el byte) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapInt(f func(el byte) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapInt8(f func(el byte) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapInt16(f func(el byte) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapInt32(f func(el byte) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapInt64(f func(el byte) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapUint(f func(el byte) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapUint8(f func(el byte) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapUint16(f func(el byte) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapUint32(f func(el byte) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapUint64(f func(el byte) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceByte) MapInterface(f func(el byte) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceByte) Repeat(val byte) chan byte {
	c := make(chan byte, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceByte) Any(f func(el byte) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceByte) All(f func(el byte) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceByte) ChunkByBool(f func(el byte) bool) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByByte(f func(el byte) byte) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByString(f func(el byte) string) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByFloat32(f func(el byte) float32) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByFloat64(f func(el byte) float64) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByInt(f func(el byte) int) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByInt8(f func(el byte) int8) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByInt16(f func(el byte) int16) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByInt32(f func(el byte) int32) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByInt64(f func(el byte) int64) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByUint(f func(el byte) uint) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByUint8(f func(el byte) uint8) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByUint16(f func(el byte) uint16) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByUint32(f func(el byte) uint32) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByUint64(f func(el byte) uint64) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkByInterface(f func(el byte) interface{}) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceByte) ChunkEvery(count int) [][]byte {
	chunks := make([][]byte, 0)
	chunk := make([]byte, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]byte, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceByte) Contains(el byte) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceByte) Count(el byte) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceByte) Cycle() chan byte {
	c := make(chan byte, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceByte) Dedup() []byte {
	result := make([]byte, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceByte) DedupByBool(f func(el byte) bool) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByByte(f func(el byte) byte) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByString(f func(el byte) string) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByFloat32(f func(el byte) float32) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByFloat64(f func(el byte) float64) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByInt(f func(el byte) int) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByInt8(f func(el byte) int8) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByInt16(f func(el byte) int16) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByInt32(f func(el byte) int32) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByInt64(f func(el byte) int64) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByUint(f func(el byte) uint) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByUint8(f func(el byte) uint8) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByUint16(f func(el byte) uint16) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByUint32(f func(el byte) uint32) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByUint64(f func(el byte) uint64) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DedupByInterface(f func(el byte) interface{}) []byte {
	result := make([]byte, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceByte) DropEvery(nth int) []byte {
	result := make([]byte, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceByte) DropWhile(f func(arr byte) bool) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceByte) Each(f func(el byte)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceByte) Filter(f func(el byte) bool) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceByte) Find(def byte, f func(el byte) bool) byte {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceByte) FindIndex(f func(el byte) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceByte) GroupByBool(f func(el byte) bool) map[bool][]byte {
	result := make(map[bool][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByByte(f func(el byte) byte) map[byte][]byte {
	result := make(map[byte][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByString(f func(el byte) string) map[string][]byte {
	result := make(map[string][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByFloat32(f func(el byte) float32) map[float32][]byte {
	result := make(map[float32][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByFloat64(f func(el byte) float64) map[float64][]byte {
	result := make(map[float64][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByInt(f func(el byte) int) map[int][]byte {
	result := make(map[int][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByInt8(f func(el byte) int8) map[int8][]byte {
	result := make(map[int8][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByInt16(f func(el byte) int16) map[int16][]byte {
	result := make(map[int16][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByInt32(f func(el byte) int32) map[int32][]byte {
	result := make(map[int32][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByInt64(f func(el byte) int64) map[int64][]byte {
	result := make(map[int64][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByUint(f func(el byte) uint) map[uint][]byte {
	result := make(map[uint][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByUint8(f func(el byte) uint8) map[uint8][]byte {
	result := make(map[uint8][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByUint16(f func(el byte) uint16) map[uint16][]byte {
	result := make(map[uint16][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByUint32(f func(el byte) uint32) map[uint32][]byte {
	result := make(map[uint32][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByUint64(f func(el byte) uint64) map[uint64][]byte {
	result := make(map[uint64][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) GroupByInterface(f func(el byte) interface{}) map[interface{}][]byte {
	result := make(map[interface{}][]byte)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]byte, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceByte) Intersperse(el byte) []byte {
	result := make([]byte, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceByte) MapBool(f func(el byte) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapByte(f func(el byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapString(f func(el byte) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapFloat32(f func(el byte) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapFloat64(f func(el byte) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapInt(f func(el byte) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapInt8(f func(el byte) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapInt16(f func(el byte) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapInt32(f func(el byte) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapInt64(f func(el byte) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapUint(f func(el byte) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapUint8(f func(el byte) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapUint16(f func(el byte) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapUint32(f func(el byte) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapUint64(f func(el byte) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) MapInterface(f func(el byte) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceByte) Max() byte {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceByte) Min() byte {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceByte) Product(repeat int) chan []byte {
	c := make(chan []byte, 1)
	go s.product(c, repeat, []byte{}, 0)
	return c
}

func (s SliceByte) product(c chan []byte, repeat int, left []byte, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]byte, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]byte, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceByte) ReduceBool(acc bool, f func(el byte, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceByte(acc byte, f func(el byte, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceString(acc string, f func(el byte, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceFloat32(acc float32, f func(el byte, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceFloat64(acc float64, f func(el byte, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceInt(acc int, f func(el byte, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceInt8(acc int8, f func(el byte, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceInt16(acc int16, f func(el byte, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceInt32(acc int32, f func(el byte, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceInt64(acc int64, f func(el byte, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceUint(acc uint, f func(el byte, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceUint8(acc uint8, f func(el byte, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceUint16(acc uint16, f func(el byte, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceUint32(acc uint32, f func(el byte, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceUint64(acc uint64, f func(el byte, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceInterface(acc interface{}, f func(el byte, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceByte) ReduceWhileBool(acc bool, f func(el byte, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileByte(acc byte, f func(el byte, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileString(acc string, f func(el byte, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileFloat32(acc float32, f func(el byte, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileFloat64(acc float64, f func(el byte, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileInt(acc int, f func(el byte, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileInt8(acc int8, f func(el byte, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileInt16(acc int16, f func(el byte, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileInt32(acc int32, f func(el byte, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileInt64(acc int64, f func(el byte, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileUint(acc uint, f func(el byte, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileUint8(acc uint8, f func(el byte, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileUint16(acc uint16, f func(el byte, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileUint32(acc uint32, f func(el byte, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileUint64(acc uint64, f func(el byte, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) ReduceWhileInterface(acc interface{}, f func(el byte, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceByte) Reverse() []byte {
	result := make([]byte, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceByte) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceByte) ScanBool(acc bool, f func(el byte, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanByte(acc byte, f func(el byte, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanString(acc string, f func(el byte, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanFloat32(acc float32, f func(el byte, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanFloat64(acc float64, f func(el byte, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanInt(acc int, f func(el byte, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanInt8(acc int8, f func(el byte, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanInt16(acc int16, f func(el byte, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanInt32(acc int32, f func(el byte, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanInt64(acc int64, f func(el byte, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanUint(acc uint, f func(el byte, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanUint8(acc uint8, f func(el byte, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanUint16(acc uint16, f func(el byte, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanUint32(acc uint32, f func(el byte, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanUint64(acc uint64, f func(el byte, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) ScanInterface(acc interface{}, f func(el byte, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceByte) Shuffle() []byte {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceByte) Sort() []byte {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceByte) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceByte) Split(sep byte) [][]byte {
	result := make([][]byte, 0)
	curr := make([]byte, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceByte) StartsWith(prefix []byte) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceByte) Sum() byte {
	var sum byte
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceByte) TakeWhile(f func(el byte) bool) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceByte) Uniq() []byte {
	added := make(map[byte]struct{})
	nothing := struct{}{}
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceByte) Window(size int) [][]byte {
	result := make([][]byte, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesByte) Concat() []byte {
	result := make([]byte, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesByte) Product() chan []byte {
	c := make(chan []byte, 1)
	go s.product(c, []byte{}, 0)
	return c
}

func (s SlicesByte) product(c chan []byte, left []byte, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]byte, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]byte, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesByte) Zip() [][]byte {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]byte, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]byte, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelString struct {
	data chan string
}

type AsyncSliceString struct {
	data    []string
	workers int
}

type SequenceString struct {
	data chan string
}

type SliceString struct {
	data []string
}

type SlicesString struct {
	data [][]string
}

func (c ChannelString) Any(f func(el string) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelString) All(f func(el string) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelString) ChunkEvery(count int) chan []string {
	chunks := make(chan []string, 1)
	go func() {
		chunk := make([]string, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]string, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelString) Count(el string) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelString) Drop(n int) chan string {
	result := make(chan string, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelString) Each(f func(el string)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelString) Filter(f func(el string) bool) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapBool(f func(el string) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapByte(f func(el string) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapString(f func(el string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapFloat32(f func(el string) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapFloat64(f func(el string) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapInt(f func(el string) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapInt8(f func(el string) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapInt16(f func(el string) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapInt32(f func(el string) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapInt64(f func(el string) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapUint(f func(el string) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapUint8(f func(el string) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapUint16(f func(el string) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapUint32(f func(el string) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapUint64(f func(el string) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) MapInterface(f func(el string) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelString) Max() string {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelString) Min() string {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelString) ReduceBool(acc bool, f func(el string, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceByte(acc byte, f func(el string, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceString(acc string, f func(el string, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceFloat32(acc float32, f func(el string, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceFloat64(acc float64, f func(el string, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceInt(acc int, f func(el string, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceInt8(acc int8, f func(el string, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceInt16(acc int16, f func(el string, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceInt32(acc int32, f func(el string, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceInt64(acc int64, f func(el string, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceUint(acc uint, f func(el string, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceUint8(acc uint8, f func(el string, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceUint16(acc uint16, f func(el string, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceUint32(acc uint32, f func(el string, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceUint64(acc uint64, f func(el string, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ReduceInterface(acc interface{}, f func(el string, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelString) ScanBool(acc bool, f func(el string, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanByte(acc byte, f func(el string, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanString(acc string, f func(el string, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanFloat32(acc float32, f func(el string, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanFloat64(acc float64, f func(el string, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanInt(acc int, f func(el string, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanInt8(acc int8, f func(el string, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanInt16(acc int16, f func(el string, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanInt32(acc int32, f func(el string, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanInt64(acc int64, f func(el string, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanUint(acc uint, f func(el string, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanUint8(acc uint8, f func(el string, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanUint16(acc uint16, f func(el string, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanUint32(acc uint32, f func(el string, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanUint64(acc uint64, f func(el string, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) ScanInterface(acc interface{}, f func(el string, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelString) Sum() string {
	var sum string
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelString) Take(n int) []string {
	result := make([]string, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelString) Tee(count int) []chan string {
	channels := make([]chan string, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan string, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan string) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelString) ToSlice() []string {
	result := make([]string, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceString) Each(f func(el string)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceString) Filter(f func(el string) bool) []string {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]string, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceString) MapBool(f func(el string) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapByte(f func(el string) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapString(f func(el string) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapFloat32(f func(el string) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapFloat64(f func(el string) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapInt(f func(el string) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapInt8(f func(el string) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapInt16(f func(el string) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapInt32(f func(el string) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapInt64(f func(el string) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapUint(f func(el string) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapUint8(f func(el string) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapUint16(f func(el string) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapUint32(f func(el string) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapUint64(f func(el string) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceString) MapInterface(f func(el string) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceString) Repeat(val string) chan string {
	c := make(chan string, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceString) Any(f func(el string) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceString) All(f func(el string) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceString) ChunkByBool(f func(el string) bool) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByByte(f func(el string) byte) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByString(f func(el string) string) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByFloat32(f func(el string) float32) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByFloat64(f func(el string) float64) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByInt(f func(el string) int) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByInt8(f func(el string) int8) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByInt16(f func(el string) int16) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByInt32(f func(el string) int32) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByInt64(f func(el string) int64) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByUint(f func(el string) uint) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByUint8(f func(el string) uint8) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByUint16(f func(el string) uint16) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByUint32(f func(el string) uint32) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByUint64(f func(el string) uint64) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkByInterface(f func(el string) interface{}) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceString) ChunkEvery(count int) [][]string {
	chunks := make([][]string, 0)
	chunk := make([]string, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceString) Contains(el string) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceString) Count(el string) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceString) Cycle() chan string {
	c := make(chan string, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceString) Dedup() []string {
	result := make([]string, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceString) DedupByBool(f func(el string) bool) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByByte(f func(el string) byte) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByString(f func(el string) string) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByFloat32(f func(el string) float32) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByFloat64(f func(el string) float64) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByInt(f func(el string) int) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByInt8(f func(el string) int8) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByInt16(f func(el string) int16) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByInt32(f func(el string) int32) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByInt64(f func(el string) int64) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByUint(f func(el string) uint) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByUint8(f func(el string) uint8) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByUint16(f func(el string) uint16) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByUint32(f func(el string) uint32) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByUint64(f func(el string) uint64) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DedupByInterface(f func(el string) interface{}) []string {
	result := make([]string, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceString) DropEvery(nth int) []string {
	result := make([]string, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceString) DropWhile(f func(arr string) bool) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceString) Each(f func(el string)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceString) Filter(f func(el string) bool) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceString) Find(def string, f func(el string) bool) string {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceString) FindIndex(f func(el string) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceString) GroupByBool(f func(el string) bool) map[bool][]string {
	result := make(map[bool][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByByte(f func(el string) byte) map[byte][]string {
	result := make(map[byte][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByString(f func(el string) string) map[string][]string {
	result := make(map[string][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByFloat32(f func(el string) float32) map[float32][]string {
	result := make(map[float32][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByFloat64(f func(el string) float64) map[float64][]string {
	result := make(map[float64][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByInt(f func(el string) int) map[int][]string {
	result := make(map[int][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByInt8(f func(el string) int8) map[int8][]string {
	result := make(map[int8][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByInt16(f func(el string) int16) map[int16][]string {
	result := make(map[int16][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByInt32(f func(el string) int32) map[int32][]string {
	result := make(map[int32][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByInt64(f func(el string) int64) map[int64][]string {
	result := make(map[int64][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByUint(f func(el string) uint) map[uint][]string {
	result := make(map[uint][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByUint8(f func(el string) uint8) map[uint8][]string {
	result := make(map[uint8][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByUint16(f func(el string) uint16) map[uint16][]string {
	result := make(map[uint16][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByUint32(f func(el string) uint32) map[uint32][]string {
	result := make(map[uint32][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByUint64(f func(el string) uint64) map[uint64][]string {
	result := make(map[uint64][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) GroupByInterface(f func(el string) interface{}) map[interface{}][]string {
	result := make(map[interface{}][]string)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]string, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceString) Intersperse(el string) []string {
	result := make([]string, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceString) MapBool(f func(el string) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapByte(f func(el string) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapString(f func(el string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapFloat32(f func(el string) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapFloat64(f func(el string) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapInt(f func(el string) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapInt8(f func(el string) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapInt16(f func(el string) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapInt32(f func(el string) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapInt64(f func(el string) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapUint(f func(el string) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapUint8(f func(el string) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapUint16(f func(el string) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapUint32(f func(el string) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapUint64(f func(el string) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) MapInterface(f func(el string) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceString) Max() string {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceString) Min() string {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceString) Product(repeat int) chan []string {
	c := make(chan []string, 1)
	go s.product(c, repeat, []string{}, 0)
	return c
}

func (s SliceString) product(c chan []string, repeat int, left []string, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]string, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]string, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceString) ReduceBool(acc bool, f func(el string, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceByte(acc byte, f func(el string, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceString(acc string, f func(el string, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceFloat32(acc float32, f func(el string, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceFloat64(acc float64, f func(el string, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceInt(acc int, f func(el string, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceInt8(acc int8, f func(el string, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceInt16(acc int16, f func(el string, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceInt32(acc int32, f func(el string, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceInt64(acc int64, f func(el string, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceUint(acc uint, f func(el string, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceUint8(acc uint8, f func(el string, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceUint16(acc uint16, f func(el string, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceUint32(acc uint32, f func(el string, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceUint64(acc uint64, f func(el string, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceInterface(acc interface{}, f func(el string, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceString) ReduceWhileBool(acc bool, f func(el string, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileByte(acc byte, f func(el string, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileString(acc string, f func(el string, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileFloat32(acc float32, f func(el string, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileFloat64(acc float64, f func(el string, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileInt(acc int, f func(el string, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileInt8(acc int8, f func(el string, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileInt16(acc int16, f func(el string, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileInt32(acc int32, f func(el string, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileInt64(acc int64, f func(el string, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileUint(acc uint, f func(el string, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileUint8(acc uint8, f func(el string, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileUint16(acc uint16, f func(el string, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileUint32(acc uint32, f func(el string, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileUint64(acc uint64, f func(el string, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) ReduceWhileInterface(acc interface{}, f func(el string, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceString) Reverse() []string {
	result := make([]string, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceString) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceString) ScanBool(acc bool, f func(el string, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanByte(acc byte, f func(el string, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanString(acc string, f func(el string, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanFloat32(acc float32, f func(el string, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanFloat64(acc float64, f func(el string, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanInt(acc int, f func(el string, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanInt8(acc int8, f func(el string, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanInt16(acc int16, f func(el string, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanInt32(acc int32, f func(el string, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanInt64(acc int64, f func(el string, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanUint(acc uint, f func(el string, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanUint8(acc uint8, f func(el string, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanUint16(acc uint16, f func(el string, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanUint32(acc uint32, f func(el string, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanUint64(acc uint64, f func(el string, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) ScanInterface(acc interface{}, f func(el string, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceString) Shuffle() []string {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceString) Sort() []string {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceString) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceString) Split(sep string) [][]string {
	result := make([][]string, 0)
	curr := make([]string, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceString) StartsWith(prefix []string) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceString) Sum() string {
	var sum string
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceString) TakeWhile(f func(el string) bool) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceString) Uniq() []string {
	added := make(map[string]struct{})
	nothing := struct{}{}
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceString) Window(size int) [][]string {
	result := make([][]string, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesString) Concat() []string {
	result := make([]string, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesString) Product() chan []string {
	c := make(chan []string, 1)
	go s.product(c, []string{}, 0)
	return c
}

func (s SlicesString) product(c chan []string, left []string, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]string, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]string, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesString) Zip() [][]string {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]string, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]string, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelFloat32 struct {
	data chan float32
}

type AsyncSliceFloat32 struct {
	data    []float32
	workers int
}

type SequenceFloat32 struct {
	data chan float32
}

type SliceFloat32 struct {
	data []float32
}

type SlicesFloat32 struct {
	data [][]float32
}

func (c ChannelFloat32) Any(f func(el float32) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelFloat32) All(f func(el float32) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelFloat32) ChunkEvery(count int) chan []float32 {
	chunks := make(chan []float32, 1)
	go func() {
		chunk := make([]float32, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]float32, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelFloat32) Count(el float32) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelFloat32) Drop(n int) chan float32 {
	result := make(chan float32, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) Each(f func(el float32)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelFloat32) Filter(f func(el float32) bool) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapBool(f func(el float32) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapByte(f func(el float32) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapString(f func(el float32) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapFloat32(f func(el float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapFloat64(f func(el float32) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapInt(f func(el float32) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapInt8(f func(el float32) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapInt16(f func(el float32) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapInt32(f func(el float32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapInt64(f func(el float32) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapUint(f func(el float32) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapUint8(f func(el float32) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapUint16(f func(el float32) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapUint32(f func(el float32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapUint64(f func(el float32) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) MapInterface(f func(el float32) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) Max() float32 {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelFloat32) Min() float32 {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelFloat32) ReduceBool(acc bool, f func(el float32, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceByte(acc byte, f func(el float32, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceString(acc string, f func(el float32, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceFloat32(acc float32, f func(el float32, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceFloat64(acc float64, f func(el float32, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceInt(acc int, f func(el float32, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceInt8(acc int8, f func(el float32, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceInt16(acc int16, f func(el float32, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceInt32(acc int32, f func(el float32, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceInt64(acc int64, f func(el float32, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceUint(acc uint, f func(el float32, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceUint8(acc uint8, f func(el float32, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceUint16(acc uint16, f func(el float32, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceUint32(acc uint32, f func(el float32, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceUint64(acc uint64, f func(el float32, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ReduceInterface(acc interface{}, f func(el float32, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat32) ScanBool(acc bool, f func(el float32, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanByte(acc byte, f func(el float32, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanString(acc string, f func(el float32, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanFloat32(acc float32, f func(el float32, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanFloat64(acc float64, f func(el float32, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanInt(acc int, f func(el float32, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanInt8(acc int8, f func(el float32, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanInt16(acc int16, f func(el float32, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanInt32(acc int32, f func(el float32, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanInt64(acc int64, f func(el float32, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanUint(acc uint, f func(el float32, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanUint8(acc uint8, f func(el float32, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanUint16(acc uint16, f func(el float32, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanUint32(acc uint32, f func(el float32, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanUint64(acc uint64, f func(el float32, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) ScanInterface(acc interface{}, f func(el float32, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat32) Sum() float32 {
	var sum float32
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelFloat32) Take(n int) []float32 {
	result := make([]float32, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelFloat32) Tee(count int) []chan float32 {
	channels := make([]chan float32, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan float32, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan float32) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelFloat32) ToSlice() []float32 {
	result := make([]float32, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceFloat32) Each(f func(el float32)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceFloat32) Filter(f func(el float32) bool) []float32 {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]float32, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceFloat32) MapBool(f func(el float32) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapByte(f func(el float32) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapString(f func(el float32) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapFloat32(f func(el float32) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapFloat64(f func(el float32) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapInt(f func(el float32) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapInt8(f func(el float32) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapInt16(f func(el float32) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapInt32(f func(el float32) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapInt64(f func(el float32) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapUint(f func(el float32) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapUint8(f func(el float32) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapUint16(f func(el float32) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapUint32(f func(el float32) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapUint64(f func(el float32) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat32) MapInterface(f func(el float32) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceFloat32) Count(start float32, step float32) chan float32 {
	c := make(chan float32, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

func (s SequenceFloat32) Exponential(start float32, factor float32) chan float32 {
	c := make(chan float32, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

func (s SequenceFloat32) Range(start float32, end float32, step float32) chan float32 {
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

func (s SequenceFloat32) Repeat(val float32) chan float32 {
	c := make(chan float32, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceFloat32) Any(f func(el float32) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceFloat32) All(f func(el float32) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceFloat32) ChunkByBool(f func(el float32) bool) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByByte(f func(el float32) byte) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByString(f func(el float32) string) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByFloat32(f func(el float32) float32) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByFloat64(f func(el float32) float64) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByInt(f func(el float32) int) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByInt8(f func(el float32) int8) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByInt16(f func(el float32) int16) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByInt32(f func(el float32) int32) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByInt64(f func(el float32) int64) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByUint(f func(el float32) uint) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByUint8(f func(el float32) uint8) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByUint16(f func(el float32) uint16) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByUint32(f func(el float32) uint32) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByUint64(f func(el float32) uint64) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkByInterface(f func(el float32) interface{}) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat32) ChunkEvery(count int) [][]float32 {
	chunks := make([][]float32, 0)
	chunk := make([]float32, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]float32, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceFloat32) Contains(el float32) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceFloat32) Count(el float32) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceFloat32) Cycle() chan float32 {
	c := make(chan float32, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceFloat32) Dedup() []float32 {
	result := make([]float32, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceFloat32) DedupByBool(f func(el float32) bool) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByByte(f func(el float32) byte) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByString(f func(el float32) string) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByFloat32(f func(el float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByFloat64(f func(el float32) float64) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByInt(f func(el float32) int) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByInt8(f func(el float32) int8) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByInt16(f func(el float32) int16) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByInt32(f func(el float32) int32) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByInt64(f func(el float32) int64) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByUint(f func(el float32) uint) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByUint8(f func(el float32) uint8) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByUint16(f func(el float32) uint16) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByUint32(f func(el float32) uint32) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByUint64(f func(el float32) uint64) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DedupByInterface(f func(el float32) interface{}) []float32 {
	result := make([]float32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat32) DropEvery(nth int) []float32 {
	result := make([]float32, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceFloat32) DropWhile(f func(arr float32) bool) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceFloat32) Each(f func(el float32)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceFloat32) Filter(f func(el float32) bool) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceFloat32) Find(def float32, f func(el float32) bool) float32 {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceFloat32) FindIndex(f func(el float32) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceFloat32) GroupByBool(f func(el float32) bool) map[bool][]float32 {
	result := make(map[bool][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByByte(f func(el float32) byte) map[byte][]float32 {
	result := make(map[byte][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByString(f func(el float32) string) map[string][]float32 {
	result := make(map[string][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByFloat32(f func(el float32) float32) map[float32][]float32 {
	result := make(map[float32][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByFloat64(f func(el float32) float64) map[float64][]float32 {
	result := make(map[float64][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByInt(f func(el float32) int) map[int][]float32 {
	result := make(map[int][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByInt8(f func(el float32) int8) map[int8][]float32 {
	result := make(map[int8][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByInt16(f func(el float32) int16) map[int16][]float32 {
	result := make(map[int16][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByInt32(f func(el float32) int32) map[int32][]float32 {
	result := make(map[int32][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByInt64(f func(el float32) int64) map[int64][]float32 {
	result := make(map[int64][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByUint(f func(el float32) uint) map[uint][]float32 {
	result := make(map[uint][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByUint8(f func(el float32) uint8) map[uint8][]float32 {
	result := make(map[uint8][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByUint16(f func(el float32) uint16) map[uint16][]float32 {
	result := make(map[uint16][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByUint32(f func(el float32) uint32) map[uint32][]float32 {
	result := make(map[uint32][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByUint64(f func(el float32) uint64) map[uint64][]float32 {
	result := make(map[uint64][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) GroupByInterface(f func(el float32) interface{}) map[interface{}][]float32 {
	result := make(map[interface{}][]float32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat32) Intersperse(el float32) []float32 {
	result := make([]float32, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceFloat32) MapBool(f func(el float32) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapByte(f func(el float32) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapString(f func(el float32) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapFloat32(f func(el float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapFloat64(f func(el float32) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapInt(f func(el float32) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapInt8(f func(el float32) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapInt16(f func(el float32) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapInt32(f func(el float32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapInt64(f func(el float32) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapUint(f func(el float32) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapUint8(f func(el float32) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapUint16(f func(el float32) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapUint32(f func(el float32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapUint64(f func(el float32) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) MapInterface(f func(el float32) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat32) Max() float32 {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceFloat32) Min() float32 {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceFloat32) Product(repeat int) chan []float32 {
	c := make(chan []float32, 1)
	go s.product(c, repeat, []float32{}, 0)
	return c
}

func (s SliceFloat32) product(c chan []float32, repeat int, left []float32, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]float32, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]float32, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceFloat32) ReduceBool(acc bool, f func(el float32, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceByte(acc byte, f func(el float32, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceString(acc string, f func(el float32, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceFloat32(acc float32, f func(el float32, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceFloat64(acc float64, f func(el float32, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceInt(acc int, f func(el float32, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceInt8(acc int8, f func(el float32, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceInt16(acc int16, f func(el float32, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceInt32(acc int32, f func(el float32, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceInt64(acc int64, f func(el float32, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceUint(acc uint, f func(el float32, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceUint8(acc uint8, f func(el float32, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceUint16(acc uint16, f func(el float32, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceUint32(acc uint32, f func(el float32, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceUint64(acc uint64, f func(el float32, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceInterface(acc interface{}, f func(el float32, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat32) ReduceWhileBool(acc bool, f func(el float32, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileByte(acc byte, f func(el float32, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileString(acc string, f func(el float32, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileFloat32(acc float32, f func(el float32, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileFloat64(acc float64, f func(el float32, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileInt(acc int, f func(el float32, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileInt8(acc int8, f func(el float32, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileInt16(acc int16, f func(el float32, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileInt32(acc int32, f func(el float32, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileInt64(acc int64, f func(el float32, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileUint(acc uint, f func(el float32, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileUint8(acc uint8, f func(el float32, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileUint16(acc uint16, f func(el float32, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileUint32(acc uint32, f func(el float32, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileUint64(acc uint64, f func(el float32, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) ReduceWhileInterface(acc interface{}, f func(el float32, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat32) Reverse() []float32 {
	result := make([]float32, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceFloat32) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceFloat32) ScanBool(acc bool, f func(el float32, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanByte(acc byte, f func(el float32, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanString(acc string, f func(el float32, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanFloat32(acc float32, f func(el float32, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanFloat64(acc float64, f func(el float32, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanInt(acc int, f func(el float32, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanInt8(acc int8, f func(el float32, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanInt16(acc int16, f func(el float32, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanInt32(acc int32, f func(el float32, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanInt64(acc int64, f func(el float32, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanUint(acc uint, f func(el float32, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanUint8(acc uint8, f func(el float32, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanUint16(acc uint16, f func(el float32, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanUint32(acc uint32, f func(el float32, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanUint64(acc uint64, f func(el float32, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) ScanInterface(acc interface{}, f func(el float32, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat32) Shuffle() []float32 {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceFloat32) Sort() []float32 {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceFloat32) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceFloat32) Split(sep float32) [][]float32 {
	result := make([][]float32, 0)
	curr := make([]float32, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceFloat32) StartsWith(prefix []float32) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceFloat32) Sum() float32 {
	var sum float32
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceFloat32) TakeWhile(f func(el float32) bool) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceFloat32) Uniq() []float32 {
	added := make(map[float32]struct{})
	nothing := struct{}{}
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceFloat32) Window(size int) [][]float32 {
	result := make([][]float32, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesFloat32) Concat() []float32 {
	result := make([]float32, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesFloat32) Product() chan []float32 {
	c := make(chan []float32, 1)
	go s.product(c, []float32{}, 0)
	return c
}

func (s SlicesFloat32) product(c chan []float32, left []float32, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]float32, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]float32, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesFloat32) Zip() [][]float32 {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]float32, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]float32, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelFloat64 struct {
	data chan float64
}

type AsyncSliceFloat64 struct {
	data    []float64
	workers int
}

type SequenceFloat64 struct {
	data chan float64
}

type SliceFloat64 struct {
	data []float64
}

type SlicesFloat64 struct {
	data [][]float64
}

func (c ChannelFloat64) Any(f func(el float64) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelFloat64) All(f func(el float64) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelFloat64) ChunkEvery(count int) chan []float64 {
	chunks := make(chan []float64, 1)
	go func() {
		chunk := make([]float64, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]float64, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelFloat64) Count(el float64) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelFloat64) Drop(n int) chan float64 {
	result := make(chan float64, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) Each(f func(el float64)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelFloat64) Filter(f func(el float64) bool) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapBool(f func(el float64) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapByte(f func(el float64) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapString(f func(el float64) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapFloat32(f func(el float64) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapFloat64(f func(el float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapInt(f func(el float64) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapInt8(f func(el float64) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapInt16(f func(el float64) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapInt32(f func(el float64) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapInt64(f func(el float64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapUint(f func(el float64) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapUint8(f func(el float64) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapUint16(f func(el float64) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapUint32(f func(el float64) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapUint64(f func(el float64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) MapInterface(f func(el float64) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) Max() float64 {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelFloat64) Min() float64 {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelFloat64) ReduceBool(acc bool, f func(el float64, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceByte(acc byte, f func(el float64, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceString(acc string, f func(el float64, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceFloat32(acc float32, f func(el float64, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceFloat64(acc float64, f func(el float64, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceInt(acc int, f func(el float64, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceInt8(acc int8, f func(el float64, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceInt16(acc int16, f func(el float64, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceInt32(acc int32, f func(el float64, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceInt64(acc int64, f func(el float64, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceUint(acc uint, f func(el float64, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceUint8(acc uint8, f func(el float64, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceUint16(acc uint16, f func(el float64, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceUint32(acc uint32, f func(el float64, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceUint64(acc uint64, f func(el float64, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ReduceInterface(acc interface{}, f func(el float64, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelFloat64) ScanBool(acc bool, f func(el float64, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanByte(acc byte, f func(el float64, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanString(acc string, f func(el float64, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanFloat32(acc float32, f func(el float64, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanFloat64(acc float64, f func(el float64, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanInt(acc int, f func(el float64, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanInt8(acc int8, f func(el float64, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanInt16(acc int16, f func(el float64, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanInt32(acc int32, f func(el float64, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanInt64(acc int64, f func(el float64, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanUint(acc uint, f func(el float64, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanUint8(acc uint8, f func(el float64, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanUint16(acc uint16, f func(el float64, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanUint32(acc uint32, f func(el float64, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanUint64(acc uint64, f func(el float64, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) ScanInterface(acc interface{}, f func(el float64, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelFloat64) Sum() float64 {
	var sum float64
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelFloat64) Take(n int) []float64 {
	result := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelFloat64) Tee(count int) []chan float64 {
	channels := make([]chan float64, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan float64, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan float64) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelFloat64) ToSlice() []float64 {
	result := make([]float64, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceFloat64) Each(f func(el float64)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceFloat64) Filter(f func(el float64) bool) []float64 {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]float64, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceFloat64) MapBool(f func(el float64) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapByte(f func(el float64) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapString(f func(el float64) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapFloat32(f func(el float64) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapFloat64(f func(el float64) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapInt(f func(el float64) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapInt8(f func(el float64) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapInt16(f func(el float64) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapInt32(f func(el float64) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapInt64(f func(el float64) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapUint(f func(el float64) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapUint8(f func(el float64) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapUint16(f func(el float64) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapUint32(f func(el float64) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapUint64(f func(el float64) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceFloat64) MapInterface(f func(el float64) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceFloat64) Count(start float64, step float64) chan float64 {
	c := make(chan float64, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

func (s SequenceFloat64) Exponential(start float64, factor float64) chan float64 {
	c := make(chan float64, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

func (s SequenceFloat64) Range(start float64, end float64, step float64) chan float64 {
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

func (s SequenceFloat64) Repeat(val float64) chan float64 {
	c := make(chan float64, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceFloat64) Any(f func(el float64) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceFloat64) All(f func(el float64) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceFloat64) ChunkByBool(f func(el float64) bool) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByByte(f func(el float64) byte) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByString(f func(el float64) string) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByFloat32(f func(el float64) float32) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByFloat64(f func(el float64) float64) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByInt(f func(el float64) int) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByInt8(f func(el float64) int8) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByInt16(f func(el float64) int16) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByInt32(f func(el float64) int32) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByInt64(f func(el float64) int64) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByUint(f func(el float64) uint) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByUint8(f func(el float64) uint8) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByUint16(f func(el float64) uint16) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByUint32(f func(el float64) uint32) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByUint64(f func(el float64) uint64) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkByInterface(f func(el float64) interface{}) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceFloat64) ChunkEvery(count int) [][]float64 {
	chunks := make([][]float64, 0)
	chunk := make([]float64, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]float64, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceFloat64) Contains(el float64) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceFloat64) Count(el float64) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceFloat64) Cycle() chan float64 {
	c := make(chan float64, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceFloat64) Dedup() []float64 {
	result := make([]float64, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceFloat64) DedupByBool(f func(el float64) bool) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByByte(f func(el float64) byte) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByString(f func(el float64) string) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByFloat32(f func(el float64) float32) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByFloat64(f func(el float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByInt(f func(el float64) int) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByInt8(f func(el float64) int8) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByInt16(f func(el float64) int16) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByInt32(f func(el float64) int32) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByInt64(f func(el float64) int64) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByUint(f func(el float64) uint) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByUint8(f func(el float64) uint8) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByUint16(f func(el float64) uint16) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByUint32(f func(el float64) uint32) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByUint64(f func(el float64) uint64) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DedupByInterface(f func(el float64) interface{}) []float64 {
	result := make([]float64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceFloat64) DropEvery(nth int) []float64 {
	result := make([]float64, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceFloat64) DropWhile(f func(arr float64) bool) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceFloat64) Each(f func(el float64)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceFloat64) Filter(f func(el float64) bool) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceFloat64) Find(def float64, f func(el float64) bool) float64 {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceFloat64) FindIndex(f func(el float64) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceFloat64) GroupByBool(f func(el float64) bool) map[bool][]float64 {
	result := make(map[bool][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByByte(f func(el float64) byte) map[byte][]float64 {
	result := make(map[byte][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByString(f func(el float64) string) map[string][]float64 {
	result := make(map[string][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByFloat32(f func(el float64) float32) map[float32][]float64 {
	result := make(map[float32][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByFloat64(f func(el float64) float64) map[float64][]float64 {
	result := make(map[float64][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByInt(f func(el float64) int) map[int][]float64 {
	result := make(map[int][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByInt8(f func(el float64) int8) map[int8][]float64 {
	result := make(map[int8][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByInt16(f func(el float64) int16) map[int16][]float64 {
	result := make(map[int16][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByInt32(f func(el float64) int32) map[int32][]float64 {
	result := make(map[int32][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByInt64(f func(el float64) int64) map[int64][]float64 {
	result := make(map[int64][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByUint(f func(el float64) uint) map[uint][]float64 {
	result := make(map[uint][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByUint8(f func(el float64) uint8) map[uint8][]float64 {
	result := make(map[uint8][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByUint16(f func(el float64) uint16) map[uint16][]float64 {
	result := make(map[uint16][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByUint32(f func(el float64) uint32) map[uint32][]float64 {
	result := make(map[uint32][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByUint64(f func(el float64) uint64) map[uint64][]float64 {
	result := make(map[uint64][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) GroupByInterface(f func(el float64) interface{}) map[interface{}][]float64 {
	result := make(map[interface{}][]float64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]float64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceFloat64) Intersperse(el float64) []float64 {
	result := make([]float64, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceFloat64) MapBool(f func(el float64) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapByte(f func(el float64) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapString(f func(el float64) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapFloat32(f func(el float64) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapFloat64(f func(el float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapInt(f func(el float64) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapInt8(f func(el float64) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapInt16(f func(el float64) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapInt32(f func(el float64) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapInt64(f func(el float64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapUint(f func(el float64) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapUint8(f func(el float64) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapUint16(f func(el float64) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapUint32(f func(el float64) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapUint64(f func(el float64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) MapInterface(f func(el float64) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceFloat64) Max() float64 {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceFloat64) Min() float64 {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceFloat64) Product(repeat int) chan []float64 {
	c := make(chan []float64, 1)
	go s.product(c, repeat, []float64{}, 0)
	return c
}

func (s SliceFloat64) product(c chan []float64, repeat int, left []float64, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]float64, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]float64, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceFloat64) ReduceBool(acc bool, f func(el float64, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceByte(acc byte, f func(el float64, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceString(acc string, f func(el float64, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceFloat32(acc float32, f func(el float64, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceFloat64(acc float64, f func(el float64, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceInt(acc int, f func(el float64, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceInt8(acc int8, f func(el float64, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceInt16(acc int16, f func(el float64, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceInt32(acc int32, f func(el float64, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceInt64(acc int64, f func(el float64, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceUint(acc uint, f func(el float64, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceUint8(acc uint8, f func(el float64, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceUint16(acc uint16, f func(el float64, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceUint32(acc uint32, f func(el float64, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceUint64(acc uint64, f func(el float64, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceInterface(acc interface{}, f func(el float64, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceFloat64) ReduceWhileBool(acc bool, f func(el float64, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileByte(acc byte, f func(el float64, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileString(acc string, f func(el float64, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileFloat32(acc float32, f func(el float64, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileFloat64(acc float64, f func(el float64, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileInt(acc int, f func(el float64, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileInt8(acc int8, f func(el float64, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileInt16(acc int16, f func(el float64, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileInt32(acc int32, f func(el float64, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileInt64(acc int64, f func(el float64, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileUint(acc uint, f func(el float64, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileUint8(acc uint8, f func(el float64, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileUint16(acc uint16, f func(el float64, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileUint32(acc uint32, f func(el float64, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileUint64(acc uint64, f func(el float64, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) ReduceWhileInterface(acc interface{}, f func(el float64, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceFloat64) Reverse() []float64 {
	result := make([]float64, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceFloat64) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceFloat64) ScanBool(acc bool, f func(el float64, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanByte(acc byte, f func(el float64, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanString(acc string, f func(el float64, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanFloat32(acc float32, f func(el float64, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanFloat64(acc float64, f func(el float64, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanInt(acc int, f func(el float64, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanInt8(acc int8, f func(el float64, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanInt16(acc int16, f func(el float64, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanInt32(acc int32, f func(el float64, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanInt64(acc int64, f func(el float64, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanUint(acc uint, f func(el float64, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanUint8(acc uint8, f func(el float64, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanUint16(acc uint16, f func(el float64, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanUint32(acc uint32, f func(el float64, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanUint64(acc uint64, f func(el float64, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) ScanInterface(acc interface{}, f func(el float64, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceFloat64) Shuffle() []float64 {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceFloat64) Sort() []float64 {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceFloat64) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceFloat64) Split(sep float64) [][]float64 {
	result := make([][]float64, 0)
	curr := make([]float64, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceFloat64) StartsWith(prefix []float64) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceFloat64) Sum() float64 {
	var sum float64
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceFloat64) TakeWhile(f func(el float64) bool) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceFloat64) Uniq() []float64 {
	added := make(map[float64]struct{})
	nothing := struct{}{}
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceFloat64) Window(size int) [][]float64 {
	result := make([][]float64, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesFloat64) Concat() []float64 {
	result := make([]float64, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesFloat64) Product() chan []float64 {
	c := make(chan []float64, 1)
	go s.product(c, []float64{}, 0)
	return c
}

func (s SlicesFloat64) product(c chan []float64, left []float64, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]float64, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]float64, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesFloat64) Zip() [][]float64 {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]float64, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]float64, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelInt struct {
	data chan int
}

type AsyncSliceInt struct {
	data    []int
	workers int
}

type SequenceInt struct {
	data chan int
}

type SliceInt struct {
	data []int
}

type SlicesInt struct {
	data [][]int
}

func (c ChannelInt) Any(f func(el int) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelInt) All(f func(el int) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelInt) ChunkEvery(count int) chan []int {
	chunks := make(chan []int, 1)
	go func() {
		chunk := make([]int, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]int, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelInt) Count(el int) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelInt) Drop(n int) chan int {
	result := make(chan int, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) Each(f func(el int)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelInt) Filter(f func(el int) bool) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapBool(f func(el int) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapByte(f func(el int) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapString(f func(el int) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapFloat32(f func(el int) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapFloat64(f func(el int) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapInt(f func(el int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapInt8(f func(el int) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapInt16(f func(el int) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapInt32(f func(el int) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapInt64(f func(el int) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapUint(f func(el int) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapUint8(f func(el int) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapUint16(f func(el int) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapUint32(f func(el int) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapUint64(f func(el int) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) MapInterface(f func(el int) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) Max() int {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelInt) Min() int {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelInt) ReduceBool(acc bool, f func(el int, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceByte(acc byte, f func(el int, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceString(acc string, f func(el int, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceFloat32(acc float32, f func(el int, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceFloat64(acc float64, f func(el int, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceInt(acc int, f func(el int, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceInt8(acc int8, f func(el int, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceInt16(acc int16, f func(el int, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceInt32(acc int32, f func(el int, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceInt64(acc int64, f func(el int, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceUint(acc uint, f func(el int, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceUint8(acc uint8, f func(el int, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceUint16(acc uint16, f func(el int, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceUint32(acc uint32, f func(el int, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceUint64(acc uint64, f func(el int, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ReduceInterface(acc interface{}, f func(el int, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt) ScanBool(acc bool, f func(el int, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanByte(acc byte, f func(el int, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanString(acc string, f func(el int, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanFloat32(acc float32, f func(el int, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanFloat64(acc float64, f func(el int, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanInt(acc int, f func(el int, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanInt8(acc int8, f func(el int, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanInt16(acc int16, f func(el int, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanInt32(acc int32, f func(el int, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanInt64(acc int64, f func(el int, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanUint(acc uint, f func(el int, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanUint8(acc uint8, f func(el int, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanUint16(acc uint16, f func(el int, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanUint32(acc uint32, f func(el int, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanUint64(acc uint64, f func(el int, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) ScanInterface(acc interface{}, f func(el int, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt) Sum() int {
	var sum int
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelInt) Take(n int) []int {
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelInt) Tee(count int) []chan int {
	channels := make([]chan int, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan int, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan int) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelInt) ToSlice() []int {
	result := make([]int, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceInt) Each(f func(el int)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceInt) Filter(f func(el int) bool) []int {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]int, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceInt) MapBool(f func(el int) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapByte(f func(el int) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapString(f func(el int) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapFloat32(f func(el int) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapFloat64(f func(el int) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapInt(f func(el int) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapInt8(f func(el int) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapInt16(f func(el int) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapInt32(f func(el int) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapInt64(f func(el int) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapUint(f func(el int) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapUint8(f func(el int) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapUint16(f func(el int) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapUint32(f func(el int) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapUint64(f func(el int) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt) MapInterface(f func(el int) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceInt) Count(start int, step int) chan int {
	c := make(chan int, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

func (s SequenceInt) Exponential(start int, factor int) chan int {
	c := make(chan int, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

func (s SequenceInt) Range(start int, end int, step int) chan int {
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

func (s SequenceInt) Repeat(val int) chan int {
	c := make(chan int, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceInt) Any(f func(el int) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceInt) All(f func(el int) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceInt) ChunkByBool(f func(el int) bool) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByByte(f func(el int) byte) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByString(f func(el int) string) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByFloat32(f func(el int) float32) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByFloat64(f func(el int) float64) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByInt(f func(el int) int) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByInt8(f func(el int) int8) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByInt16(f func(el int) int16) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByInt32(f func(el int) int32) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByInt64(f func(el int) int64) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByUint(f func(el int) uint) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByUint8(f func(el int) uint8) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByUint16(f func(el int) uint16) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByUint32(f func(el int) uint32) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByUint64(f func(el int) uint64) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkByInterface(f func(el int) interface{}) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt) ChunkEvery(count int) [][]int {
	chunks := make([][]int, 0)
	chunk := make([]int, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceInt) Contains(el int) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceInt) Count(el int) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceInt) Cycle() chan int {
	c := make(chan int, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceInt) Dedup() []int {
	result := make([]int, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceInt) DedupByBool(f func(el int) bool) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByByte(f func(el int) byte) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByString(f func(el int) string) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByFloat32(f func(el int) float32) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByFloat64(f func(el int) float64) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByInt(f func(el int) int) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByInt8(f func(el int) int8) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByInt16(f func(el int) int16) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByInt32(f func(el int) int32) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByInt64(f func(el int) int64) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByUint(f func(el int) uint) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByUint8(f func(el int) uint8) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByUint16(f func(el int) uint16) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByUint32(f func(el int) uint32) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByUint64(f func(el int) uint64) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DedupByInterface(f func(el int) interface{}) []int {
	result := make([]int, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt) DropEvery(nth int) []int {
	result := make([]int, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceInt) DropWhile(f func(arr int) bool) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceInt) Each(f func(el int)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceInt) Filter(f func(el int) bool) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceInt) Find(def int, f func(el int) bool) int {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceInt) FindIndex(f func(el int) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceInt) GroupByBool(f func(el int) bool) map[bool][]int {
	result := make(map[bool][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByByte(f func(el int) byte) map[byte][]int {
	result := make(map[byte][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByString(f func(el int) string) map[string][]int {
	result := make(map[string][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByFloat32(f func(el int) float32) map[float32][]int {
	result := make(map[float32][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByFloat64(f func(el int) float64) map[float64][]int {
	result := make(map[float64][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByInt(f func(el int) int) map[int][]int {
	result := make(map[int][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByInt8(f func(el int) int8) map[int8][]int {
	result := make(map[int8][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByInt16(f func(el int) int16) map[int16][]int {
	result := make(map[int16][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByInt32(f func(el int) int32) map[int32][]int {
	result := make(map[int32][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByInt64(f func(el int) int64) map[int64][]int {
	result := make(map[int64][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByUint(f func(el int) uint) map[uint][]int {
	result := make(map[uint][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByUint8(f func(el int) uint8) map[uint8][]int {
	result := make(map[uint8][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByUint16(f func(el int) uint16) map[uint16][]int {
	result := make(map[uint16][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByUint32(f func(el int) uint32) map[uint32][]int {
	result := make(map[uint32][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByUint64(f func(el int) uint64) map[uint64][]int {
	result := make(map[uint64][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) GroupByInterface(f func(el int) interface{}) map[interface{}][]int {
	result := make(map[interface{}][]int)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt) Intersperse(el int) []int {
	result := make([]int, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceInt) MapBool(f func(el int) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapByte(f func(el int) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapString(f func(el int) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapFloat32(f func(el int) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapFloat64(f func(el int) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapInt(f func(el int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapInt8(f func(el int) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapInt16(f func(el int) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapInt32(f func(el int) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapInt64(f func(el int) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapUint(f func(el int) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapUint8(f func(el int) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapUint16(f func(el int) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapUint32(f func(el int) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapUint64(f func(el int) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) MapInterface(f func(el int) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt) Max() int {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceInt) Min() int {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceInt) Product(repeat int) chan []int {
	c := make(chan []int, 1)
	go s.product(c, repeat, []int{}, 0)
	return c
}

func (s SliceInt) product(c chan []int, repeat int, left []int, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]int, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]int, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceInt) ReduceBool(acc bool, f func(el int, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceByte(acc byte, f func(el int, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceString(acc string, f func(el int, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceFloat32(acc float32, f func(el int, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceFloat64(acc float64, f func(el int, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceInt(acc int, f func(el int, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceInt8(acc int8, f func(el int, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceInt16(acc int16, f func(el int, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceInt32(acc int32, f func(el int, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceInt64(acc int64, f func(el int, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceUint(acc uint, f func(el int, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceUint8(acc uint8, f func(el int, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceUint16(acc uint16, f func(el int, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceUint32(acc uint32, f func(el int, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceUint64(acc uint64, f func(el int, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceInterface(acc interface{}, f func(el int, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt) ReduceWhileBool(acc bool, f func(el int, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileByte(acc byte, f func(el int, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileString(acc string, f func(el int, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileFloat32(acc float32, f func(el int, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileFloat64(acc float64, f func(el int, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileInt(acc int, f func(el int, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileInt8(acc int8, f func(el int, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileInt16(acc int16, f func(el int, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileInt32(acc int32, f func(el int, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileInt64(acc int64, f func(el int, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileUint(acc uint, f func(el int, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileUint8(acc uint8, f func(el int, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileUint16(acc uint16, f func(el int, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileUint32(acc uint32, f func(el int, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileUint64(acc uint64, f func(el int, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) ReduceWhileInterface(acc interface{}, f func(el int, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt) Reverse() []int {
	result := make([]int, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceInt) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceInt) ScanBool(acc bool, f func(el int, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanByte(acc byte, f func(el int, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanString(acc string, f func(el int, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanFloat32(acc float32, f func(el int, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanFloat64(acc float64, f func(el int, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanInt(acc int, f func(el int, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanInt8(acc int8, f func(el int, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanInt16(acc int16, f func(el int, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanInt32(acc int32, f func(el int, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanInt64(acc int64, f func(el int, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanUint(acc uint, f func(el int, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanUint8(acc uint8, f func(el int, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanUint16(acc uint16, f func(el int, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanUint32(acc uint32, f func(el int, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanUint64(acc uint64, f func(el int, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) ScanInterface(acc interface{}, f func(el int, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt) Shuffle() []int {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceInt) Sort() []int {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceInt) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceInt) Split(sep int) [][]int {
	result := make([][]int, 0)
	curr := make([]int, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceInt) StartsWith(prefix []int) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceInt) Sum() int {
	var sum int
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceInt) TakeWhile(f func(el int) bool) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceInt) Uniq() []int {
	added := make(map[int]struct{})
	nothing := struct{}{}
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceInt) Window(size int) [][]int {
	result := make([][]int, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesInt) Concat() []int {
	result := make([]int, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesInt) Product() chan []int {
	c := make(chan []int, 1)
	go s.product(c, []int{}, 0)
	return c
}

func (s SlicesInt) product(c chan []int, left []int, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]int, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]int, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesInt) Zip() [][]int {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]int, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]int, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelInt8 struct {
	data chan int8
}

type AsyncSliceInt8 struct {
	data    []int8
	workers int
}

type SequenceInt8 struct {
	data chan int8
}

type SliceInt8 struct {
	data []int8
}

type SlicesInt8 struct {
	data [][]int8
}

func (c ChannelInt8) Any(f func(el int8) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelInt8) All(f func(el int8) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelInt8) ChunkEvery(count int) chan []int8 {
	chunks := make(chan []int8, 1)
	go func() {
		chunk := make([]int8, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]int8, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelInt8) Count(el int8) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelInt8) Drop(n int) chan int8 {
	result := make(chan int8, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) Each(f func(el int8)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelInt8) Filter(f func(el int8) bool) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapBool(f func(el int8) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapByte(f func(el int8) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapString(f func(el int8) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapFloat32(f func(el int8) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapFloat64(f func(el int8) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapInt(f func(el int8) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapInt8(f func(el int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapInt16(f func(el int8) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapInt32(f func(el int8) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapInt64(f func(el int8) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapUint(f func(el int8) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapUint8(f func(el int8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapUint16(f func(el int8) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapUint32(f func(el int8) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapUint64(f func(el int8) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) MapInterface(f func(el int8) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) Max() int8 {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelInt8) Min() int8 {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelInt8) ReduceBool(acc bool, f func(el int8, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceByte(acc byte, f func(el int8, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceString(acc string, f func(el int8, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceFloat32(acc float32, f func(el int8, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceFloat64(acc float64, f func(el int8, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceInt(acc int, f func(el int8, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceInt8(acc int8, f func(el int8, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceInt16(acc int16, f func(el int8, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceInt32(acc int32, f func(el int8, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceInt64(acc int64, f func(el int8, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceUint(acc uint, f func(el int8, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceUint8(acc uint8, f func(el int8, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceUint16(acc uint16, f func(el int8, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceUint32(acc uint32, f func(el int8, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceUint64(acc uint64, f func(el int8, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ReduceInterface(acc interface{}, f func(el int8, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt8) ScanBool(acc bool, f func(el int8, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanByte(acc byte, f func(el int8, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanString(acc string, f func(el int8, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanFloat32(acc float32, f func(el int8, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanFloat64(acc float64, f func(el int8, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanInt(acc int, f func(el int8, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanInt8(acc int8, f func(el int8, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanInt16(acc int16, f func(el int8, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanInt32(acc int32, f func(el int8, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanInt64(acc int64, f func(el int8, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanUint(acc uint, f func(el int8, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanUint8(acc uint8, f func(el int8, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanUint16(acc uint16, f func(el int8, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanUint32(acc uint32, f func(el int8, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanUint64(acc uint64, f func(el int8, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) ScanInterface(acc interface{}, f func(el int8, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt8) Sum() int8 {
	var sum int8
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelInt8) Take(n int) []int8 {
	result := make([]int8, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelInt8) Tee(count int) []chan int8 {
	channels := make([]chan int8, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan int8, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan int8) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelInt8) ToSlice() []int8 {
	result := make([]int8, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceInt8) Each(f func(el int8)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceInt8) Filter(f func(el int8) bool) []int8 {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]int8, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceInt8) MapBool(f func(el int8) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapByte(f func(el int8) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapString(f func(el int8) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapFloat32(f func(el int8) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapFloat64(f func(el int8) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapInt(f func(el int8) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapInt8(f func(el int8) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapInt16(f func(el int8) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapInt32(f func(el int8) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapInt64(f func(el int8) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapUint(f func(el int8) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapUint8(f func(el int8) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapUint16(f func(el int8) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapUint32(f func(el int8) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapUint64(f func(el int8) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt8) MapInterface(f func(el int8) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceInt8) Count(start int8, step int8) chan int8 {
	c := make(chan int8, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

func (s SequenceInt8) Exponential(start int8, factor int8) chan int8 {
	c := make(chan int8, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

func (s SequenceInt8) Range(start int8, end int8, step int8) chan int8 {
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

func (s SequenceInt8) Repeat(val int8) chan int8 {
	c := make(chan int8, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceInt8) Any(f func(el int8) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceInt8) All(f func(el int8) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceInt8) ChunkByBool(f func(el int8) bool) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByByte(f func(el int8) byte) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByString(f func(el int8) string) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByFloat32(f func(el int8) float32) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByFloat64(f func(el int8) float64) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByInt(f func(el int8) int) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByInt8(f func(el int8) int8) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByInt16(f func(el int8) int16) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByInt32(f func(el int8) int32) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByInt64(f func(el int8) int64) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByUint(f func(el int8) uint) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByUint8(f func(el int8) uint8) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByUint16(f func(el int8) uint16) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByUint32(f func(el int8) uint32) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByUint64(f func(el int8) uint64) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkByInterface(f func(el int8) interface{}) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt8) ChunkEvery(count int) [][]int8 {
	chunks := make([][]int8, 0)
	chunk := make([]int8, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]int8, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceInt8) Contains(el int8) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceInt8) Count(el int8) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceInt8) Cycle() chan int8 {
	c := make(chan int8, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceInt8) Dedup() []int8 {
	result := make([]int8, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceInt8) DedupByBool(f func(el int8) bool) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByByte(f func(el int8) byte) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByString(f func(el int8) string) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByFloat32(f func(el int8) float32) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByFloat64(f func(el int8) float64) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByInt(f func(el int8) int) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByInt8(f func(el int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByInt16(f func(el int8) int16) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByInt32(f func(el int8) int32) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByInt64(f func(el int8) int64) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByUint(f func(el int8) uint) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByUint8(f func(el int8) uint8) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByUint16(f func(el int8) uint16) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByUint32(f func(el int8) uint32) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByUint64(f func(el int8) uint64) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DedupByInterface(f func(el int8) interface{}) []int8 {
	result := make([]int8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt8) DropEvery(nth int) []int8 {
	result := make([]int8, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceInt8) DropWhile(f func(arr int8) bool) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceInt8) Each(f func(el int8)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceInt8) Filter(f func(el int8) bool) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceInt8) Find(def int8, f func(el int8) bool) int8 {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceInt8) FindIndex(f func(el int8) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceInt8) GroupByBool(f func(el int8) bool) map[bool][]int8 {
	result := make(map[bool][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByByte(f func(el int8) byte) map[byte][]int8 {
	result := make(map[byte][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByString(f func(el int8) string) map[string][]int8 {
	result := make(map[string][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByFloat32(f func(el int8) float32) map[float32][]int8 {
	result := make(map[float32][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByFloat64(f func(el int8) float64) map[float64][]int8 {
	result := make(map[float64][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByInt(f func(el int8) int) map[int][]int8 {
	result := make(map[int][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByInt8(f func(el int8) int8) map[int8][]int8 {
	result := make(map[int8][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByInt16(f func(el int8) int16) map[int16][]int8 {
	result := make(map[int16][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByInt32(f func(el int8) int32) map[int32][]int8 {
	result := make(map[int32][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByInt64(f func(el int8) int64) map[int64][]int8 {
	result := make(map[int64][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByUint(f func(el int8) uint) map[uint][]int8 {
	result := make(map[uint][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByUint8(f func(el int8) uint8) map[uint8][]int8 {
	result := make(map[uint8][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByUint16(f func(el int8) uint16) map[uint16][]int8 {
	result := make(map[uint16][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByUint32(f func(el int8) uint32) map[uint32][]int8 {
	result := make(map[uint32][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByUint64(f func(el int8) uint64) map[uint64][]int8 {
	result := make(map[uint64][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) GroupByInterface(f func(el int8) interface{}) map[interface{}][]int8 {
	result := make(map[interface{}][]int8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt8) Intersperse(el int8) []int8 {
	result := make([]int8, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceInt8) MapBool(f func(el int8) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapByte(f func(el int8) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapString(f func(el int8) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapFloat32(f func(el int8) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapFloat64(f func(el int8) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapInt(f func(el int8) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapInt8(f func(el int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapInt16(f func(el int8) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapInt32(f func(el int8) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapInt64(f func(el int8) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapUint(f func(el int8) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapUint8(f func(el int8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapUint16(f func(el int8) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapUint32(f func(el int8) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapUint64(f func(el int8) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) MapInterface(f func(el int8) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt8) Max() int8 {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceInt8) Min() int8 {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceInt8) Product(repeat int) chan []int8 {
	c := make(chan []int8, 1)
	go s.product(c, repeat, []int8{}, 0)
	return c
}

func (s SliceInt8) product(c chan []int8, repeat int, left []int8, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]int8, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]int8, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceInt8) ReduceBool(acc bool, f func(el int8, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceByte(acc byte, f func(el int8, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceString(acc string, f func(el int8, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceFloat32(acc float32, f func(el int8, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceFloat64(acc float64, f func(el int8, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceInt(acc int, f func(el int8, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceInt8(acc int8, f func(el int8, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceInt16(acc int16, f func(el int8, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceInt32(acc int32, f func(el int8, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceInt64(acc int64, f func(el int8, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceUint(acc uint, f func(el int8, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceUint8(acc uint8, f func(el int8, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceUint16(acc uint16, f func(el int8, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceUint32(acc uint32, f func(el int8, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceUint64(acc uint64, f func(el int8, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceInterface(acc interface{}, f func(el int8, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt8) ReduceWhileBool(acc bool, f func(el int8, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileByte(acc byte, f func(el int8, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileString(acc string, f func(el int8, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileFloat32(acc float32, f func(el int8, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileFloat64(acc float64, f func(el int8, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileInt(acc int, f func(el int8, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileInt8(acc int8, f func(el int8, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileInt16(acc int16, f func(el int8, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileInt32(acc int32, f func(el int8, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileInt64(acc int64, f func(el int8, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileUint(acc uint, f func(el int8, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileUint8(acc uint8, f func(el int8, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileUint16(acc uint16, f func(el int8, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileUint32(acc uint32, f func(el int8, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileUint64(acc uint64, f func(el int8, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) ReduceWhileInterface(acc interface{}, f func(el int8, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt8) Reverse() []int8 {
	result := make([]int8, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceInt8) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceInt8) ScanBool(acc bool, f func(el int8, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanByte(acc byte, f func(el int8, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanString(acc string, f func(el int8, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanFloat32(acc float32, f func(el int8, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanFloat64(acc float64, f func(el int8, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanInt(acc int, f func(el int8, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanInt8(acc int8, f func(el int8, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanInt16(acc int16, f func(el int8, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanInt32(acc int32, f func(el int8, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanInt64(acc int64, f func(el int8, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanUint(acc uint, f func(el int8, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanUint8(acc uint8, f func(el int8, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanUint16(acc uint16, f func(el int8, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanUint32(acc uint32, f func(el int8, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanUint64(acc uint64, f func(el int8, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) ScanInterface(acc interface{}, f func(el int8, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt8) Shuffle() []int8 {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceInt8) Sort() []int8 {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceInt8) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceInt8) Split(sep int8) [][]int8 {
	result := make([][]int8, 0)
	curr := make([]int8, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceInt8) StartsWith(prefix []int8) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceInt8) Sum() int8 {
	var sum int8
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceInt8) TakeWhile(f func(el int8) bool) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceInt8) Uniq() []int8 {
	added := make(map[int8]struct{})
	nothing := struct{}{}
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceInt8) Window(size int) [][]int8 {
	result := make([][]int8, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesInt8) Concat() []int8 {
	result := make([]int8, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesInt8) Product() chan []int8 {
	c := make(chan []int8, 1)
	go s.product(c, []int8{}, 0)
	return c
}

func (s SlicesInt8) product(c chan []int8, left []int8, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]int8, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]int8, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesInt8) Zip() [][]int8 {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]int8, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]int8, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelInt16 struct {
	data chan int16
}

type AsyncSliceInt16 struct {
	data    []int16
	workers int
}

type SequenceInt16 struct {
	data chan int16
}

type SliceInt16 struct {
	data []int16
}

type SlicesInt16 struct {
	data [][]int16
}

func (c ChannelInt16) Any(f func(el int16) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelInt16) All(f func(el int16) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelInt16) ChunkEvery(count int) chan []int16 {
	chunks := make(chan []int16, 1)
	go func() {
		chunk := make([]int16, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]int16, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelInt16) Count(el int16) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelInt16) Drop(n int) chan int16 {
	result := make(chan int16, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) Each(f func(el int16)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelInt16) Filter(f func(el int16) bool) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapBool(f func(el int16) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapByte(f func(el int16) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapString(f func(el int16) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapFloat32(f func(el int16) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapFloat64(f func(el int16) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapInt(f func(el int16) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapInt8(f func(el int16) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapInt16(f func(el int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapInt32(f func(el int16) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapInt64(f func(el int16) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapUint(f func(el int16) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapUint8(f func(el int16) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapUint16(f func(el int16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapUint32(f func(el int16) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapUint64(f func(el int16) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) MapInterface(f func(el int16) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) Max() int16 {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelInt16) Min() int16 {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelInt16) ReduceBool(acc bool, f func(el int16, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceByte(acc byte, f func(el int16, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceString(acc string, f func(el int16, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceFloat32(acc float32, f func(el int16, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceFloat64(acc float64, f func(el int16, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceInt(acc int, f func(el int16, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceInt8(acc int8, f func(el int16, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceInt16(acc int16, f func(el int16, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceInt32(acc int32, f func(el int16, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceInt64(acc int64, f func(el int16, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceUint(acc uint, f func(el int16, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceUint8(acc uint8, f func(el int16, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceUint16(acc uint16, f func(el int16, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceUint32(acc uint32, f func(el int16, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceUint64(acc uint64, f func(el int16, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ReduceInterface(acc interface{}, f func(el int16, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt16) ScanBool(acc bool, f func(el int16, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanByte(acc byte, f func(el int16, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanString(acc string, f func(el int16, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanFloat32(acc float32, f func(el int16, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanFloat64(acc float64, f func(el int16, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanInt(acc int, f func(el int16, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanInt8(acc int8, f func(el int16, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanInt16(acc int16, f func(el int16, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanInt32(acc int32, f func(el int16, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanInt64(acc int64, f func(el int16, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanUint(acc uint, f func(el int16, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanUint8(acc uint8, f func(el int16, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanUint16(acc uint16, f func(el int16, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanUint32(acc uint32, f func(el int16, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanUint64(acc uint64, f func(el int16, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) ScanInterface(acc interface{}, f func(el int16, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt16) Sum() int16 {
	var sum int16
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelInt16) Take(n int) []int16 {
	result := make([]int16, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelInt16) Tee(count int) []chan int16 {
	channels := make([]chan int16, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan int16, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan int16) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelInt16) ToSlice() []int16 {
	result := make([]int16, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceInt16) Each(f func(el int16)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceInt16) Filter(f func(el int16) bool) []int16 {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]int16, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceInt16) MapBool(f func(el int16) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapByte(f func(el int16) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapString(f func(el int16) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapFloat32(f func(el int16) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapFloat64(f func(el int16) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapInt(f func(el int16) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapInt8(f func(el int16) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapInt16(f func(el int16) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapInt32(f func(el int16) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapInt64(f func(el int16) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapUint(f func(el int16) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapUint8(f func(el int16) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapUint16(f func(el int16) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapUint32(f func(el int16) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapUint64(f func(el int16) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt16) MapInterface(f func(el int16) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceInt16) Count(start int16, step int16) chan int16 {
	c := make(chan int16, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

func (s SequenceInt16) Exponential(start int16, factor int16) chan int16 {
	c := make(chan int16, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

func (s SequenceInt16) Range(start int16, end int16, step int16) chan int16 {
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

func (s SequenceInt16) Repeat(val int16) chan int16 {
	c := make(chan int16, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceInt16) Any(f func(el int16) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceInt16) All(f func(el int16) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceInt16) ChunkByBool(f func(el int16) bool) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByByte(f func(el int16) byte) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByString(f func(el int16) string) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByFloat32(f func(el int16) float32) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByFloat64(f func(el int16) float64) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByInt(f func(el int16) int) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByInt8(f func(el int16) int8) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByInt16(f func(el int16) int16) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByInt32(f func(el int16) int32) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByInt64(f func(el int16) int64) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByUint(f func(el int16) uint) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByUint8(f func(el int16) uint8) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByUint16(f func(el int16) uint16) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByUint32(f func(el int16) uint32) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByUint64(f func(el int16) uint64) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkByInterface(f func(el int16) interface{}) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt16) ChunkEvery(count int) [][]int16 {
	chunks := make([][]int16, 0)
	chunk := make([]int16, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]int16, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceInt16) Contains(el int16) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceInt16) Count(el int16) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceInt16) Cycle() chan int16 {
	c := make(chan int16, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceInt16) Dedup() []int16 {
	result := make([]int16, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceInt16) DedupByBool(f func(el int16) bool) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByByte(f func(el int16) byte) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByString(f func(el int16) string) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByFloat32(f func(el int16) float32) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByFloat64(f func(el int16) float64) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByInt(f func(el int16) int) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByInt8(f func(el int16) int8) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByInt16(f func(el int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByInt32(f func(el int16) int32) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByInt64(f func(el int16) int64) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByUint(f func(el int16) uint) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByUint8(f func(el int16) uint8) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByUint16(f func(el int16) uint16) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByUint32(f func(el int16) uint32) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByUint64(f func(el int16) uint64) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DedupByInterface(f func(el int16) interface{}) []int16 {
	result := make([]int16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt16) DropEvery(nth int) []int16 {
	result := make([]int16, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceInt16) DropWhile(f func(arr int16) bool) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceInt16) Each(f func(el int16)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceInt16) Filter(f func(el int16) bool) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceInt16) Find(def int16, f func(el int16) bool) int16 {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceInt16) FindIndex(f func(el int16) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceInt16) GroupByBool(f func(el int16) bool) map[bool][]int16 {
	result := make(map[bool][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByByte(f func(el int16) byte) map[byte][]int16 {
	result := make(map[byte][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByString(f func(el int16) string) map[string][]int16 {
	result := make(map[string][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByFloat32(f func(el int16) float32) map[float32][]int16 {
	result := make(map[float32][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByFloat64(f func(el int16) float64) map[float64][]int16 {
	result := make(map[float64][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByInt(f func(el int16) int) map[int][]int16 {
	result := make(map[int][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByInt8(f func(el int16) int8) map[int8][]int16 {
	result := make(map[int8][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByInt16(f func(el int16) int16) map[int16][]int16 {
	result := make(map[int16][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByInt32(f func(el int16) int32) map[int32][]int16 {
	result := make(map[int32][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByInt64(f func(el int16) int64) map[int64][]int16 {
	result := make(map[int64][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByUint(f func(el int16) uint) map[uint][]int16 {
	result := make(map[uint][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByUint8(f func(el int16) uint8) map[uint8][]int16 {
	result := make(map[uint8][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByUint16(f func(el int16) uint16) map[uint16][]int16 {
	result := make(map[uint16][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByUint32(f func(el int16) uint32) map[uint32][]int16 {
	result := make(map[uint32][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByUint64(f func(el int16) uint64) map[uint64][]int16 {
	result := make(map[uint64][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) GroupByInterface(f func(el int16) interface{}) map[interface{}][]int16 {
	result := make(map[interface{}][]int16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt16) Intersperse(el int16) []int16 {
	result := make([]int16, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceInt16) MapBool(f func(el int16) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapByte(f func(el int16) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapString(f func(el int16) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapFloat32(f func(el int16) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapFloat64(f func(el int16) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapInt(f func(el int16) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapInt8(f func(el int16) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapInt16(f func(el int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapInt32(f func(el int16) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapInt64(f func(el int16) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapUint(f func(el int16) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapUint8(f func(el int16) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapUint16(f func(el int16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapUint32(f func(el int16) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapUint64(f func(el int16) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) MapInterface(f func(el int16) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt16) Max() int16 {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceInt16) Min() int16 {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceInt16) Product(repeat int) chan []int16 {
	c := make(chan []int16, 1)
	go s.product(c, repeat, []int16{}, 0)
	return c
}

func (s SliceInt16) product(c chan []int16, repeat int, left []int16, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]int16, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]int16, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceInt16) ReduceBool(acc bool, f func(el int16, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceByte(acc byte, f func(el int16, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceString(acc string, f func(el int16, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceFloat32(acc float32, f func(el int16, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceFloat64(acc float64, f func(el int16, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceInt(acc int, f func(el int16, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceInt8(acc int8, f func(el int16, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceInt16(acc int16, f func(el int16, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceInt32(acc int32, f func(el int16, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceInt64(acc int64, f func(el int16, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceUint(acc uint, f func(el int16, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceUint8(acc uint8, f func(el int16, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceUint16(acc uint16, f func(el int16, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceUint32(acc uint32, f func(el int16, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceUint64(acc uint64, f func(el int16, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceInterface(acc interface{}, f func(el int16, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt16) ReduceWhileBool(acc bool, f func(el int16, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileByte(acc byte, f func(el int16, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileString(acc string, f func(el int16, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileFloat32(acc float32, f func(el int16, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileFloat64(acc float64, f func(el int16, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileInt(acc int, f func(el int16, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileInt8(acc int8, f func(el int16, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileInt16(acc int16, f func(el int16, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileInt32(acc int32, f func(el int16, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileInt64(acc int64, f func(el int16, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileUint(acc uint, f func(el int16, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileUint8(acc uint8, f func(el int16, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileUint16(acc uint16, f func(el int16, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileUint32(acc uint32, f func(el int16, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileUint64(acc uint64, f func(el int16, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) ReduceWhileInterface(acc interface{}, f func(el int16, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt16) Reverse() []int16 {
	result := make([]int16, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceInt16) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceInt16) ScanBool(acc bool, f func(el int16, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanByte(acc byte, f func(el int16, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanString(acc string, f func(el int16, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanFloat32(acc float32, f func(el int16, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanFloat64(acc float64, f func(el int16, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanInt(acc int, f func(el int16, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanInt8(acc int8, f func(el int16, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanInt16(acc int16, f func(el int16, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanInt32(acc int32, f func(el int16, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanInt64(acc int64, f func(el int16, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanUint(acc uint, f func(el int16, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanUint8(acc uint8, f func(el int16, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanUint16(acc uint16, f func(el int16, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanUint32(acc uint32, f func(el int16, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanUint64(acc uint64, f func(el int16, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) ScanInterface(acc interface{}, f func(el int16, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt16) Shuffle() []int16 {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceInt16) Sort() []int16 {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceInt16) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceInt16) Split(sep int16) [][]int16 {
	result := make([][]int16, 0)
	curr := make([]int16, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceInt16) StartsWith(prefix []int16) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceInt16) Sum() int16 {
	var sum int16
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceInt16) TakeWhile(f func(el int16) bool) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceInt16) Uniq() []int16 {
	added := make(map[int16]struct{})
	nothing := struct{}{}
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceInt16) Window(size int) [][]int16 {
	result := make([][]int16, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesInt16) Concat() []int16 {
	result := make([]int16, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesInt16) Product() chan []int16 {
	c := make(chan []int16, 1)
	go s.product(c, []int16{}, 0)
	return c
}

func (s SlicesInt16) product(c chan []int16, left []int16, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]int16, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]int16, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesInt16) Zip() [][]int16 {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]int16, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]int16, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelInt32 struct {
	data chan int32
}

type AsyncSliceInt32 struct {
	data    []int32
	workers int
}

type SequenceInt32 struct {
	data chan int32
}

type SliceInt32 struct {
	data []int32
}

type SlicesInt32 struct {
	data [][]int32
}

func (c ChannelInt32) Any(f func(el int32) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelInt32) All(f func(el int32) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelInt32) ChunkEvery(count int) chan []int32 {
	chunks := make(chan []int32, 1)
	go func() {
		chunk := make([]int32, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]int32, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelInt32) Count(el int32) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelInt32) Drop(n int) chan int32 {
	result := make(chan int32, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) Each(f func(el int32)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelInt32) Filter(f func(el int32) bool) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapBool(f func(el int32) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapByte(f func(el int32) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapString(f func(el int32) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapFloat32(f func(el int32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapFloat64(f func(el int32) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapInt(f func(el int32) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapInt8(f func(el int32) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapInt16(f func(el int32) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapInt32(f func(el int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapInt64(f func(el int32) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapUint(f func(el int32) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapUint8(f func(el int32) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapUint16(f func(el int32) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapUint32(f func(el int32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapUint64(f func(el int32) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) MapInterface(f func(el int32) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) Max() int32 {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelInt32) Min() int32 {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelInt32) ReduceBool(acc bool, f func(el int32, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceByte(acc byte, f func(el int32, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceString(acc string, f func(el int32, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceFloat32(acc float32, f func(el int32, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceFloat64(acc float64, f func(el int32, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceInt(acc int, f func(el int32, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceInt8(acc int8, f func(el int32, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceInt16(acc int16, f func(el int32, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceInt32(acc int32, f func(el int32, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceInt64(acc int64, f func(el int32, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceUint(acc uint, f func(el int32, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceUint8(acc uint8, f func(el int32, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceUint16(acc uint16, f func(el int32, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceUint32(acc uint32, f func(el int32, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceUint64(acc uint64, f func(el int32, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ReduceInterface(acc interface{}, f func(el int32, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt32) ScanBool(acc bool, f func(el int32, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanByte(acc byte, f func(el int32, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanString(acc string, f func(el int32, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanFloat32(acc float32, f func(el int32, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanFloat64(acc float64, f func(el int32, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanInt(acc int, f func(el int32, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanInt8(acc int8, f func(el int32, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanInt16(acc int16, f func(el int32, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanInt32(acc int32, f func(el int32, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanInt64(acc int64, f func(el int32, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanUint(acc uint, f func(el int32, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanUint8(acc uint8, f func(el int32, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanUint16(acc uint16, f func(el int32, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanUint32(acc uint32, f func(el int32, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanUint64(acc uint64, f func(el int32, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) ScanInterface(acc interface{}, f func(el int32, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt32) Sum() int32 {
	var sum int32
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelInt32) Take(n int) []int32 {
	result := make([]int32, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelInt32) Tee(count int) []chan int32 {
	channels := make([]chan int32, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan int32, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan int32) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelInt32) ToSlice() []int32 {
	result := make([]int32, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceInt32) Each(f func(el int32)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceInt32) Filter(f func(el int32) bool) []int32 {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]int32, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceInt32) MapBool(f func(el int32) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapByte(f func(el int32) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapString(f func(el int32) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapFloat32(f func(el int32) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapFloat64(f func(el int32) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapInt(f func(el int32) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapInt8(f func(el int32) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapInt16(f func(el int32) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapInt32(f func(el int32) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapInt64(f func(el int32) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapUint(f func(el int32) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapUint8(f func(el int32) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapUint16(f func(el int32) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapUint32(f func(el int32) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapUint64(f func(el int32) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt32) MapInterface(f func(el int32) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceInt32) Count(start int32, step int32) chan int32 {
	c := make(chan int32, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

func (s SequenceInt32) Exponential(start int32, factor int32) chan int32 {
	c := make(chan int32, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

func (s SequenceInt32) Range(start int32, end int32, step int32) chan int32 {
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

func (s SequenceInt32) Repeat(val int32) chan int32 {
	c := make(chan int32, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceInt32) Any(f func(el int32) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceInt32) All(f func(el int32) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceInt32) ChunkByBool(f func(el int32) bool) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByByte(f func(el int32) byte) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByString(f func(el int32) string) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByFloat32(f func(el int32) float32) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByFloat64(f func(el int32) float64) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByInt(f func(el int32) int) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByInt8(f func(el int32) int8) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByInt16(f func(el int32) int16) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByInt32(f func(el int32) int32) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByInt64(f func(el int32) int64) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByUint(f func(el int32) uint) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByUint8(f func(el int32) uint8) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByUint16(f func(el int32) uint16) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByUint32(f func(el int32) uint32) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByUint64(f func(el int32) uint64) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkByInterface(f func(el int32) interface{}) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt32) ChunkEvery(count int) [][]int32 {
	chunks := make([][]int32, 0)
	chunk := make([]int32, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]int32, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceInt32) Contains(el int32) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceInt32) Count(el int32) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceInt32) Cycle() chan int32 {
	c := make(chan int32, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceInt32) Dedup() []int32 {
	result := make([]int32, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceInt32) DedupByBool(f func(el int32) bool) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByByte(f func(el int32) byte) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByString(f func(el int32) string) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByFloat32(f func(el int32) float32) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByFloat64(f func(el int32) float64) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByInt(f func(el int32) int) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByInt8(f func(el int32) int8) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByInt16(f func(el int32) int16) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByInt32(f func(el int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByInt64(f func(el int32) int64) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByUint(f func(el int32) uint) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByUint8(f func(el int32) uint8) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByUint16(f func(el int32) uint16) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByUint32(f func(el int32) uint32) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByUint64(f func(el int32) uint64) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DedupByInterface(f func(el int32) interface{}) []int32 {
	result := make([]int32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt32) DropEvery(nth int) []int32 {
	result := make([]int32, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceInt32) DropWhile(f func(arr int32) bool) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceInt32) Each(f func(el int32)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceInt32) Filter(f func(el int32) bool) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceInt32) Find(def int32, f func(el int32) bool) int32 {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceInt32) FindIndex(f func(el int32) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceInt32) GroupByBool(f func(el int32) bool) map[bool][]int32 {
	result := make(map[bool][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByByte(f func(el int32) byte) map[byte][]int32 {
	result := make(map[byte][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByString(f func(el int32) string) map[string][]int32 {
	result := make(map[string][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByFloat32(f func(el int32) float32) map[float32][]int32 {
	result := make(map[float32][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByFloat64(f func(el int32) float64) map[float64][]int32 {
	result := make(map[float64][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByInt(f func(el int32) int) map[int][]int32 {
	result := make(map[int][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByInt8(f func(el int32) int8) map[int8][]int32 {
	result := make(map[int8][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByInt16(f func(el int32) int16) map[int16][]int32 {
	result := make(map[int16][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByInt32(f func(el int32) int32) map[int32][]int32 {
	result := make(map[int32][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByInt64(f func(el int32) int64) map[int64][]int32 {
	result := make(map[int64][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByUint(f func(el int32) uint) map[uint][]int32 {
	result := make(map[uint][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByUint8(f func(el int32) uint8) map[uint8][]int32 {
	result := make(map[uint8][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByUint16(f func(el int32) uint16) map[uint16][]int32 {
	result := make(map[uint16][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByUint32(f func(el int32) uint32) map[uint32][]int32 {
	result := make(map[uint32][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByUint64(f func(el int32) uint64) map[uint64][]int32 {
	result := make(map[uint64][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) GroupByInterface(f func(el int32) interface{}) map[interface{}][]int32 {
	result := make(map[interface{}][]int32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt32) Intersperse(el int32) []int32 {
	result := make([]int32, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceInt32) MapBool(f func(el int32) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapByte(f func(el int32) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapString(f func(el int32) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapFloat32(f func(el int32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapFloat64(f func(el int32) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapInt(f func(el int32) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapInt8(f func(el int32) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapInt16(f func(el int32) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapInt32(f func(el int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapInt64(f func(el int32) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapUint(f func(el int32) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapUint8(f func(el int32) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapUint16(f func(el int32) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapUint32(f func(el int32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapUint64(f func(el int32) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) MapInterface(f func(el int32) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt32) Max() int32 {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceInt32) Min() int32 {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceInt32) Product(repeat int) chan []int32 {
	c := make(chan []int32, 1)
	go s.product(c, repeat, []int32{}, 0)
	return c
}

func (s SliceInt32) product(c chan []int32, repeat int, left []int32, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]int32, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]int32, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceInt32) ReduceBool(acc bool, f func(el int32, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceByte(acc byte, f func(el int32, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceString(acc string, f func(el int32, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceFloat32(acc float32, f func(el int32, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceFloat64(acc float64, f func(el int32, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceInt(acc int, f func(el int32, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceInt8(acc int8, f func(el int32, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceInt16(acc int16, f func(el int32, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceInt32(acc int32, f func(el int32, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceInt64(acc int64, f func(el int32, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceUint(acc uint, f func(el int32, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceUint8(acc uint8, f func(el int32, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceUint16(acc uint16, f func(el int32, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceUint32(acc uint32, f func(el int32, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceUint64(acc uint64, f func(el int32, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceInterface(acc interface{}, f func(el int32, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt32) ReduceWhileBool(acc bool, f func(el int32, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileByte(acc byte, f func(el int32, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileString(acc string, f func(el int32, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileFloat32(acc float32, f func(el int32, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileFloat64(acc float64, f func(el int32, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileInt(acc int, f func(el int32, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileInt8(acc int8, f func(el int32, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileInt16(acc int16, f func(el int32, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileInt32(acc int32, f func(el int32, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileInt64(acc int64, f func(el int32, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileUint(acc uint, f func(el int32, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileUint8(acc uint8, f func(el int32, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileUint16(acc uint16, f func(el int32, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileUint32(acc uint32, f func(el int32, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileUint64(acc uint64, f func(el int32, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) ReduceWhileInterface(acc interface{}, f func(el int32, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt32) Reverse() []int32 {
	result := make([]int32, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceInt32) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceInt32) ScanBool(acc bool, f func(el int32, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanByte(acc byte, f func(el int32, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanString(acc string, f func(el int32, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanFloat32(acc float32, f func(el int32, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanFloat64(acc float64, f func(el int32, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanInt(acc int, f func(el int32, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanInt8(acc int8, f func(el int32, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanInt16(acc int16, f func(el int32, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanInt32(acc int32, f func(el int32, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanInt64(acc int64, f func(el int32, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanUint(acc uint, f func(el int32, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanUint8(acc uint8, f func(el int32, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanUint16(acc uint16, f func(el int32, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanUint32(acc uint32, f func(el int32, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanUint64(acc uint64, f func(el int32, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) ScanInterface(acc interface{}, f func(el int32, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt32) Shuffle() []int32 {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceInt32) Sort() []int32 {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceInt32) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceInt32) Split(sep int32) [][]int32 {
	result := make([][]int32, 0)
	curr := make([]int32, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceInt32) StartsWith(prefix []int32) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceInt32) Sum() int32 {
	var sum int32
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceInt32) TakeWhile(f func(el int32) bool) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceInt32) Uniq() []int32 {
	added := make(map[int32]struct{})
	nothing := struct{}{}
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceInt32) Window(size int) [][]int32 {
	result := make([][]int32, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesInt32) Concat() []int32 {
	result := make([]int32, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesInt32) Product() chan []int32 {
	c := make(chan []int32, 1)
	go s.product(c, []int32{}, 0)
	return c
}

func (s SlicesInt32) product(c chan []int32, left []int32, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]int32, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]int32, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesInt32) Zip() [][]int32 {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]int32, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]int32, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelInt64 struct {
	data chan int64
}

type AsyncSliceInt64 struct {
	data    []int64
	workers int
}

type SequenceInt64 struct {
	data chan int64
}

type SliceInt64 struct {
	data []int64
}

type SlicesInt64 struct {
	data [][]int64
}

func (c ChannelInt64) Any(f func(el int64) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelInt64) All(f func(el int64) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelInt64) ChunkEvery(count int) chan []int64 {
	chunks := make(chan []int64, 1)
	go func() {
		chunk := make([]int64, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]int64, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelInt64) Count(el int64) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelInt64) Drop(n int) chan int64 {
	result := make(chan int64, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) Each(f func(el int64)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelInt64) Filter(f func(el int64) bool) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapBool(f func(el int64) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapByte(f func(el int64) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapString(f func(el int64) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapFloat32(f func(el int64) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapFloat64(f func(el int64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapInt(f func(el int64) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapInt8(f func(el int64) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapInt16(f func(el int64) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapInt32(f func(el int64) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapInt64(f func(el int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapUint(f func(el int64) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapUint8(f func(el int64) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapUint16(f func(el int64) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapUint32(f func(el int64) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapUint64(f func(el int64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) MapInterface(f func(el int64) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) Max() int64 {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelInt64) Min() int64 {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelInt64) ReduceBool(acc bool, f func(el int64, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceByte(acc byte, f func(el int64, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceString(acc string, f func(el int64, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceFloat32(acc float32, f func(el int64, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceFloat64(acc float64, f func(el int64, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceInt(acc int, f func(el int64, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceInt8(acc int8, f func(el int64, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceInt16(acc int16, f func(el int64, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceInt32(acc int32, f func(el int64, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceInt64(acc int64, f func(el int64, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceUint(acc uint, f func(el int64, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceUint8(acc uint8, f func(el int64, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceUint16(acc uint16, f func(el int64, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceUint32(acc uint32, f func(el int64, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceUint64(acc uint64, f func(el int64, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ReduceInterface(acc interface{}, f func(el int64, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInt64) ScanBool(acc bool, f func(el int64, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanByte(acc byte, f func(el int64, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanString(acc string, f func(el int64, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanFloat32(acc float32, f func(el int64, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanFloat64(acc float64, f func(el int64, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanInt(acc int, f func(el int64, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanInt8(acc int8, f func(el int64, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanInt16(acc int16, f func(el int64, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanInt32(acc int32, f func(el int64, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanInt64(acc int64, f func(el int64, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanUint(acc uint, f func(el int64, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanUint8(acc uint8, f func(el int64, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanUint16(acc uint16, f func(el int64, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanUint32(acc uint32, f func(el int64, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanUint64(acc uint64, f func(el int64, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) ScanInterface(acc interface{}, f func(el int64, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInt64) Sum() int64 {
	var sum int64
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelInt64) Take(n int) []int64 {
	result := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelInt64) Tee(count int) []chan int64 {
	channels := make([]chan int64, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan int64, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan int64) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelInt64) ToSlice() []int64 {
	result := make([]int64, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceInt64) Each(f func(el int64)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceInt64) Filter(f func(el int64) bool) []int64 {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]int64, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceInt64) MapBool(f func(el int64) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapByte(f func(el int64) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapString(f func(el int64) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapFloat32(f func(el int64) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapFloat64(f func(el int64) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapInt(f func(el int64) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapInt8(f func(el int64) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapInt16(f func(el int64) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapInt32(f func(el int64) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapInt64(f func(el int64) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapUint(f func(el int64) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapUint8(f func(el int64) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapUint16(f func(el int64) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapUint32(f func(el int64) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapUint64(f func(el int64) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInt64) MapInterface(f func(el int64) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceInt64) Count(start int64, step int64) chan int64 {
	c := make(chan int64, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

func (s SequenceInt64) Exponential(start int64, factor int64) chan int64 {
	c := make(chan int64, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

func (s SequenceInt64) Range(start int64, end int64, step int64) chan int64 {
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

func (s SequenceInt64) Repeat(val int64) chan int64 {
	c := make(chan int64, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceInt64) Any(f func(el int64) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceInt64) All(f func(el int64) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceInt64) ChunkByBool(f func(el int64) bool) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByByte(f func(el int64) byte) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByString(f func(el int64) string) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByFloat32(f func(el int64) float32) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByFloat64(f func(el int64) float64) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByInt(f func(el int64) int) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByInt8(f func(el int64) int8) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByInt16(f func(el int64) int16) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByInt32(f func(el int64) int32) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByInt64(f func(el int64) int64) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByUint(f func(el int64) uint) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByUint8(f func(el int64) uint8) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByUint16(f func(el int64) uint16) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByUint32(f func(el int64) uint32) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByUint64(f func(el int64) uint64) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkByInterface(f func(el int64) interface{}) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInt64) ChunkEvery(count int) [][]int64 {
	chunks := make([][]int64, 0)
	chunk := make([]int64, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]int64, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceInt64) Contains(el int64) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceInt64) Count(el int64) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceInt64) Cycle() chan int64 {
	c := make(chan int64, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceInt64) Dedup() []int64 {
	result := make([]int64, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceInt64) DedupByBool(f func(el int64) bool) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByByte(f func(el int64) byte) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByString(f func(el int64) string) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByFloat32(f func(el int64) float32) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByFloat64(f func(el int64) float64) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByInt(f func(el int64) int) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByInt8(f func(el int64) int8) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByInt16(f func(el int64) int16) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByInt32(f func(el int64) int32) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByInt64(f func(el int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByUint(f func(el int64) uint) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByUint8(f func(el int64) uint8) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByUint16(f func(el int64) uint16) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByUint32(f func(el int64) uint32) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByUint64(f func(el int64) uint64) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DedupByInterface(f func(el int64) interface{}) []int64 {
	result := make([]int64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInt64) DropEvery(nth int) []int64 {
	result := make([]int64, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceInt64) DropWhile(f func(arr int64) bool) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceInt64) Each(f func(el int64)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceInt64) Filter(f func(el int64) bool) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceInt64) Find(def int64, f func(el int64) bool) int64 {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceInt64) FindIndex(f func(el int64) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceInt64) GroupByBool(f func(el int64) bool) map[bool][]int64 {
	result := make(map[bool][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByByte(f func(el int64) byte) map[byte][]int64 {
	result := make(map[byte][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByString(f func(el int64) string) map[string][]int64 {
	result := make(map[string][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByFloat32(f func(el int64) float32) map[float32][]int64 {
	result := make(map[float32][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByFloat64(f func(el int64) float64) map[float64][]int64 {
	result := make(map[float64][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByInt(f func(el int64) int) map[int][]int64 {
	result := make(map[int][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByInt8(f func(el int64) int8) map[int8][]int64 {
	result := make(map[int8][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByInt16(f func(el int64) int16) map[int16][]int64 {
	result := make(map[int16][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByInt32(f func(el int64) int32) map[int32][]int64 {
	result := make(map[int32][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByInt64(f func(el int64) int64) map[int64][]int64 {
	result := make(map[int64][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByUint(f func(el int64) uint) map[uint][]int64 {
	result := make(map[uint][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByUint8(f func(el int64) uint8) map[uint8][]int64 {
	result := make(map[uint8][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByUint16(f func(el int64) uint16) map[uint16][]int64 {
	result := make(map[uint16][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByUint32(f func(el int64) uint32) map[uint32][]int64 {
	result := make(map[uint32][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByUint64(f func(el int64) uint64) map[uint64][]int64 {
	result := make(map[uint64][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) GroupByInterface(f func(el int64) interface{}) map[interface{}][]int64 {
	result := make(map[interface{}][]int64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]int64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInt64) Intersperse(el int64) []int64 {
	result := make([]int64, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceInt64) MapBool(f func(el int64) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapByte(f func(el int64) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapString(f func(el int64) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapFloat32(f func(el int64) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapFloat64(f func(el int64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapInt(f func(el int64) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapInt8(f func(el int64) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapInt16(f func(el int64) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapInt32(f func(el int64) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapInt64(f func(el int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapUint(f func(el int64) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapUint8(f func(el int64) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapUint16(f func(el int64) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapUint32(f func(el int64) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapUint64(f func(el int64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) MapInterface(f func(el int64) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInt64) Max() int64 {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceInt64) Min() int64 {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceInt64) Product(repeat int) chan []int64 {
	c := make(chan []int64, 1)
	go s.product(c, repeat, []int64{}, 0)
	return c
}

func (s SliceInt64) product(c chan []int64, repeat int, left []int64, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]int64, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]int64, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceInt64) ReduceBool(acc bool, f func(el int64, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceByte(acc byte, f func(el int64, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceString(acc string, f func(el int64, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceFloat32(acc float32, f func(el int64, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceFloat64(acc float64, f func(el int64, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceInt(acc int, f func(el int64, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceInt8(acc int8, f func(el int64, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceInt16(acc int16, f func(el int64, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceInt32(acc int32, f func(el int64, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceInt64(acc int64, f func(el int64, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceUint(acc uint, f func(el int64, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceUint8(acc uint8, f func(el int64, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceUint16(acc uint16, f func(el int64, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceUint32(acc uint32, f func(el int64, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceUint64(acc uint64, f func(el int64, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceInterface(acc interface{}, f func(el int64, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInt64) ReduceWhileBool(acc bool, f func(el int64, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileByte(acc byte, f func(el int64, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileString(acc string, f func(el int64, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileFloat32(acc float32, f func(el int64, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileFloat64(acc float64, f func(el int64, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileInt(acc int, f func(el int64, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileInt8(acc int8, f func(el int64, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileInt16(acc int16, f func(el int64, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileInt32(acc int32, f func(el int64, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileInt64(acc int64, f func(el int64, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileUint(acc uint, f func(el int64, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileUint8(acc uint8, f func(el int64, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileUint16(acc uint16, f func(el int64, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileUint32(acc uint32, f func(el int64, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileUint64(acc uint64, f func(el int64, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) ReduceWhileInterface(acc interface{}, f func(el int64, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInt64) Reverse() []int64 {
	result := make([]int64, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceInt64) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceInt64) ScanBool(acc bool, f func(el int64, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanByte(acc byte, f func(el int64, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanString(acc string, f func(el int64, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanFloat32(acc float32, f func(el int64, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanFloat64(acc float64, f func(el int64, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanInt(acc int, f func(el int64, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanInt8(acc int8, f func(el int64, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanInt16(acc int16, f func(el int64, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanInt32(acc int32, f func(el int64, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanInt64(acc int64, f func(el int64, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanUint(acc uint, f func(el int64, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanUint8(acc uint8, f func(el int64, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanUint16(acc uint16, f func(el int64, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanUint32(acc uint32, f func(el int64, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanUint64(acc uint64, f func(el int64, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) ScanInterface(acc interface{}, f func(el int64, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInt64) Shuffle() []int64 {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceInt64) Sort() []int64 {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceInt64) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceInt64) Split(sep int64) [][]int64 {
	result := make([][]int64, 0)
	curr := make([]int64, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceInt64) StartsWith(prefix []int64) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceInt64) Sum() int64 {
	var sum int64
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceInt64) TakeWhile(f func(el int64) bool) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceInt64) Uniq() []int64 {
	added := make(map[int64]struct{})
	nothing := struct{}{}
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceInt64) Window(size int) [][]int64 {
	result := make([][]int64, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesInt64) Concat() []int64 {
	result := make([]int64, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesInt64) Product() chan []int64 {
	c := make(chan []int64, 1)
	go s.product(c, []int64{}, 0)
	return c
}

func (s SlicesInt64) product(c chan []int64, left []int64, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]int64, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]int64, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesInt64) Zip() [][]int64 {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]int64, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]int64, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelUint struct {
	data chan uint
}

type AsyncSliceUint struct {
	data    []uint
	workers int
}

type SequenceUint struct {
	data chan uint
}

type SliceUint struct {
	data []uint
}

type SlicesUint struct {
	data [][]uint
}

func (c ChannelUint) Any(f func(el uint) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelUint) All(f func(el uint) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelUint) ChunkEvery(count int) chan []uint {
	chunks := make(chan []uint, 1)
	go func() {
		chunk := make([]uint, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]uint, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelUint) Count(el uint) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelUint) Drop(n int) chan uint {
	result := make(chan uint, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) Each(f func(el uint)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelUint) Filter(f func(el uint) bool) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapBool(f func(el uint) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapByte(f func(el uint) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapString(f func(el uint) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapFloat32(f func(el uint) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapFloat64(f func(el uint) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapInt(f func(el uint) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapInt8(f func(el uint) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapInt16(f func(el uint) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapInt32(f func(el uint) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapInt64(f func(el uint) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapUint(f func(el uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapUint8(f func(el uint) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapUint16(f func(el uint) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapUint32(f func(el uint) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapUint64(f func(el uint) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) MapInterface(f func(el uint) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) Max() uint {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelUint) Min() uint {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelUint) ReduceBool(acc bool, f func(el uint, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceByte(acc byte, f func(el uint, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceString(acc string, f func(el uint, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceFloat32(acc float32, f func(el uint, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceFloat64(acc float64, f func(el uint, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceInt(acc int, f func(el uint, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceInt8(acc int8, f func(el uint, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceInt16(acc int16, f func(el uint, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceInt32(acc int32, f func(el uint, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceInt64(acc int64, f func(el uint, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceUint(acc uint, f func(el uint, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceUint8(acc uint8, f func(el uint, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceUint16(acc uint16, f func(el uint, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceUint32(acc uint32, f func(el uint, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceUint64(acc uint64, f func(el uint, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ReduceInterface(acc interface{}, f func(el uint, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint) ScanBool(acc bool, f func(el uint, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanByte(acc byte, f func(el uint, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanString(acc string, f func(el uint, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanFloat32(acc float32, f func(el uint, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanFloat64(acc float64, f func(el uint, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanInt(acc int, f func(el uint, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanInt8(acc int8, f func(el uint, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanInt16(acc int16, f func(el uint, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanInt32(acc int32, f func(el uint, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanInt64(acc int64, f func(el uint, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanUint(acc uint, f func(el uint, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanUint8(acc uint8, f func(el uint, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanUint16(acc uint16, f func(el uint, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanUint32(acc uint32, f func(el uint, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanUint64(acc uint64, f func(el uint, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) ScanInterface(acc interface{}, f func(el uint, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint) Sum() uint {
	var sum uint
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelUint) Take(n int) []uint {
	result := make([]uint, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelUint) Tee(count int) []chan uint {
	channels := make([]chan uint, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan uint, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan uint) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelUint) ToSlice() []uint {
	result := make([]uint, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceUint) Each(f func(el uint)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceUint) Filter(f func(el uint) bool) []uint {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]uint, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceUint) MapBool(f func(el uint) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapByte(f func(el uint) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapString(f func(el uint) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapFloat32(f func(el uint) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapFloat64(f func(el uint) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapInt(f func(el uint) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapInt8(f func(el uint) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapInt16(f func(el uint) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapInt32(f func(el uint) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapInt64(f func(el uint) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapUint(f func(el uint) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapUint8(f func(el uint) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapUint16(f func(el uint) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapUint32(f func(el uint) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapUint64(f func(el uint) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint) MapInterface(f func(el uint) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceUint) Count(start uint, step uint) chan uint {
	c := make(chan uint, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

func (s SequenceUint) Exponential(start uint, factor uint) chan uint {
	c := make(chan uint, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

func (s SequenceUint) Range(start uint, end uint, step uint) chan uint {
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

func (s SequenceUint) Repeat(val uint) chan uint {
	c := make(chan uint, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceUint) Any(f func(el uint) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceUint) All(f func(el uint) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceUint) ChunkByBool(f func(el uint) bool) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByByte(f func(el uint) byte) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByString(f func(el uint) string) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByFloat32(f func(el uint) float32) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByFloat64(f func(el uint) float64) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByInt(f func(el uint) int) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByInt8(f func(el uint) int8) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByInt16(f func(el uint) int16) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByInt32(f func(el uint) int32) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByInt64(f func(el uint) int64) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByUint(f func(el uint) uint) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByUint8(f func(el uint) uint8) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByUint16(f func(el uint) uint16) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByUint32(f func(el uint) uint32) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByUint64(f func(el uint) uint64) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkByInterface(f func(el uint) interface{}) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint) ChunkEvery(count int) [][]uint {
	chunks := make([][]uint, 0)
	chunk := make([]uint, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]uint, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceUint) Contains(el uint) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceUint) Count(el uint) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceUint) Cycle() chan uint {
	c := make(chan uint, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceUint) Dedup() []uint {
	result := make([]uint, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceUint) DedupByBool(f func(el uint) bool) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByByte(f func(el uint) byte) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByString(f func(el uint) string) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByFloat32(f func(el uint) float32) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByFloat64(f func(el uint) float64) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByInt(f func(el uint) int) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByInt8(f func(el uint) int8) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByInt16(f func(el uint) int16) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByInt32(f func(el uint) int32) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByInt64(f func(el uint) int64) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByUint(f func(el uint) uint) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByUint8(f func(el uint) uint8) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByUint16(f func(el uint) uint16) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByUint32(f func(el uint) uint32) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByUint64(f func(el uint) uint64) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DedupByInterface(f func(el uint) interface{}) []uint {
	result := make([]uint, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint) DropEvery(nth int) []uint {
	result := make([]uint, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceUint) DropWhile(f func(arr uint) bool) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceUint) Each(f func(el uint)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceUint) Filter(f func(el uint) bool) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceUint) Find(def uint, f func(el uint) bool) uint {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceUint) FindIndex(f func(el uint) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceUint) GroupByBool(f func(el uint) bool) map[bool][]uint {
	result := make(map[bool][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByByte(f func(el uint) byte) map[byte][]uint {
	result := make(map[byte][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByString(f func(el uint) string) map[string][]uint {
	result := make(map[string][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByFloat32(f func(el uint) float32) map[float32][]uint {
	result := make(map[float32][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByFloat64(f func(el uint) float64) map[float64][]uint {
	result := make(map[float64][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByInt(f func(el uint) int) map[int][]uint {
	result := make(map[int][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByInt8(f func(el uint) int8) map[int8][]uint {
	result := make(map[int8][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByInt16(f func(el uint) int16) map[int16][]uint {
	result := make(map[int16][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByInt32(f func(el uint) int32) map[int32][]uint {
	result := make(map[int32][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByInt64(f func(el uint) int64) map[int64][]uint {
	result := make(map[int64][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByUint(f func(el uint) uint) map[uint][]uint {
	result := make(map[uint][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByUint8(f func(el uint) uint8) map[uint8][]uint {
	result := make(map[uint8][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByUint16(f func(el uint) uint16) map[uint16][]uint {
	result := make(map[uint16][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByUint32(f func(el uint) uint32) map[uint32][]uint {
	result := make(map[uint32][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByUint64(f func(el uint) uint64) map[uint64][]uint {
	result := make(map[uint64][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) GroupByInterface(f func(el uint) interface{}) map[interface{}][]uint {
	result := make(map[interface{}][]uint)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint) Intersperse(el uint) []uint {
	result := make([]uint, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceUint) MapBool(f func(el uint) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapByte(f func(el uint) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapString(f func(el uint) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapFloat32(f func(el uint) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapFloat64(f func(el uint) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapInt(f func(el uint) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapInt8(f func(el uint) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapInt16(f func(el uint) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapInt32(f func(el uint) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapInt64(f func(el uint) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapUint(f func(el uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapUint8(f func(el uint) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapUint16(f func(el uint) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapUint32(f func(el uint) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapUint64(f func(el uint) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) MapInterface(f func(el uint) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint) Max() uint {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceUint) Min() uint {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceUint) Product(repeat int) chan []uint {
	c := make(chan []uint, 1)
	go s.product(c, repeat, []uint{}, 0)
	return c
}

func (s SliceUint) product(c chan []uint, repeat int, left []uint, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]uint, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]uint, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceUint) ReduceBool(acc bool, f func(el uint, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceByte(acc byte, f func(el uint, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceString(acc string, f func(el uint, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceFloat32(acc float32, f func(el uint, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceFloat64(acc float64, f func(el uint, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceInt(acc int, f func(el uint, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceInt8(acc int8, f func(el uint, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceInt16(acc int16, f func(el uint, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceInt32(acc int32, f func(el uint, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceInt64(acc int64, f func(el uint, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceUint(acc uint, f func(el uint, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceUint8(acc uint8, f func(el uint, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceUint16(acc uint16, f func(el uint, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceUint32(acc uint32, f func(el uint, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceUint64(acc uint64, f func(el uint, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceInterface(acc interface{}, f func(el uint, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint) ReduceWhileBool(acc bool, f func(el uint, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileByte(acc byte, f func(el uint, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileString(acc string, f func(el uint, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileFloat32(acc float32, f func(el uint, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileFloat64(acc float64, f func(el uint, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileInt(acc int, f func(el uint, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileInt8(acc int8, f func(el uint, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileInt16(acc int16, f func(el uint, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileInt32(acc int32, f func(el uint, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileInt64(acc int64, f func(el uint, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileUint(acc uint, f func(el uint, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileUint8(acc uint8, f func(el uint, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileUint16(acc uint16, f func(el uint, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileUint32(acc uint32, f func(el uint, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileUint64(acc uint64, f func(el uint, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) ReduceWhileInterface(acc interface{}, f func(el uint, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint) Reverse() []uint {
	result := make([]uint, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceUint) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceUint) ScanBool(acc bool, f func(el uint, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanByte(acc byte, f func(el uint, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanString(acc string, f func(el uint, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanFloat32(acc float32, f func(el uint, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanFloat64(acc float64, f func(el uint, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanInt(acc int, f func(el uint, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanInt8(acc int8, f func(el uint, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanInt16(acc int16, f func(el uint, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanInt32(acc int32, f func(el uint, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanInt64(acc int64, f func(el uint, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanUint(acc uint, f func(el uint, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanUint8(acc uint8, f func(el uint, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanUint16(acc uint16, f func(el uint, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanUint32(acc uint32, f func(el uint, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanUint64(acc uint64, f func(el uint, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) ScanInterface(acc interface{}, f func(el uint, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint) Shuffle() []uint {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceUint) Sort() []uint {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceUint) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceUint) Split(sep uint) [][]uint {
	result := make([][]uint, 0)
	curr := make([]uint, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceUint) StartsWith(prefix []uint) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceUint) Sum() uint {
	var sum uint
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceUint) TakeWhile(f func(el uint) bool) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceUint) Uniq() []uint {
	added := make(map[uint]struct{})
	nothing := struct{}{}
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceUint) Window(size int) [][]uint {
	result := make([][]uint, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesUint) Concat() []uint {
	result := make([]uint, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesUint) Product() chan []uint {
	c := make(chan []uint, 1)
	go s.product(c, []uint{}, 0)
	return c
}

func (s SlicesUint) product(c chan []uint, left []uint, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]uint, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]uint, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesUint) Zip() [][]uint {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]uint, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]uint, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelUint8 struct {
	data chan uint8
}

type AsyncSliceUint8 struct {
	data    []uint8
	workers int
}

type SequenceUint8 struct {
	data chan uint8
}

type SliceUint8 struct {
	data []uint8
}

type SlicesUint8 struct {
	data [][]uint8
}

func (c ChannelUint8) Any(f func(el uint8) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelUint8) All(f func(el uint8) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelUint8) ChunkEvery(count int) chan []uint8 {
	chunks := make(chan []uint8, 1)
	go func() {
		chunk := make([]uint8, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]uint8, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelUint8) Count(el uint8) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelUint8) Drop(n int) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) Each(f func(el uint8)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelUint8) Filter(f func(el uint8) bool) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapBool(f func(el uint8) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapByte(f func(el uint8) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapString(f func(el uint8) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapFloat32(f func(el uint8) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapFloat64(f func(el uint8) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapInt(f func(el uint8) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapInt8(f func(el uint8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapInt16(f func(el uint8) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapInt32(f func(el uint8) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapInt64(f func(el uint8) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapUint(f func(el uint8) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapUint8(f func(el uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapUint16(f func(el uint8) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapUint32(f func(el uint8) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapUint64(f func(el uint8) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) MapInterface(f func(el uint8) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) Max() uint8 {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelUint8) Min() uint8 {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelUint8) ReduceBool(acc bool, f func(el uint8, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceByte(acc byte, f func(el uint8, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceString(acc string, f func(el uint8, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceFloat32(acc float32, f func(el uint8, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceFloat64(acc float64, f func(el uint8, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceInt(acc int, f func(el uint8, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceInt8(acc int8, f func(el uint8, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceInt16(acc int16, f func(el uint8, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceInt32(acc int32, f func(el uint8, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceInt64(acc int64, f func(el uint8, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceUint(acc uint, f func(el uint8, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceUint8(acc uint8, f func(el uint8, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceUint16(acc uint16, f func(el uint8, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceUint32(acc uint32, f func(el uint8, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceUint64(acc uint64, f func(el uint8, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ReduceInterface(acc interface{}, f func(el uint8, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint8) ScanBool(acc bool, f func(el uint8, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanByte(acc byte, f func(el uint8, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanString(acc string, f func(el uint8, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanFloat32(acc float32, f func(el uint8, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanFloat64(acc float64, f func(el uint8, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanInt(acc int, f func(el uint8, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanInt8(acc int8, f func(el uint8, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanInt16(acc int16, f func(el uint8, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanInt32(acc int32, f func(el uint8, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanInt64(acc int64, f func(el uint8, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanUint(acc uint, f func(el uint8, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanUint8(acc uint8, f func(el uint8, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanUint16(acc uint16, f func(el uint8, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanUint32(acc uint32, f func(el uint8, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanUint64(acc uint64, f func(el uint8, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) ScanInterface(acc interface{}, f func(el uint8, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint8) Sum() uint8 {
	var sum uint8
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelUint8) Take(n int) []uint8 {
	result := make([]uint8, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelUint8) Tee(count int) []chan uint8 {
	channels := make([]chan uint8, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan uint8, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan uint8) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelUint8) ToSlice() []uint8 {
	result := make([]uint8, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceUint8) Each(f func(el uint8)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceUint8) Filter(f func(el uint8) bool) []uint8 {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]uint8, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceUint8) MapBool(f func(el uint8) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapByte(f func(el uint8) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapString(f func(el uint8) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapFloat32(f func(el uint8) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapFloat64(f func(el uint8) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapInt(f func(el uint8) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapInt8(f func(el uint8) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapInt16(f func(el uint8) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapInt32(f func(el uint8) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapInt64(f func(el uint8) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapUint(f func(el uint8) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapUint8(f func(el uint8) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapUint16(f func(el uint8) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapUint32(f func(el uint8) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapUint64(f func(el uint8) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint8) MapInterface(f func(el uint8) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceUint8) Count(start uint8, step uint8) chan uint8 {
	c := make(chan uint8, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

func (s SequenceUint8) Exponential(start uint8, factor uint8) chan uint8 {
	c := make(chan uint8, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

func (s SequenceUint8) Range(start uint8, end uint8, step uint8) chan uint8 {
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

func (s SequenceUint8) Repeat(val uint8) chan uint8 {
	c := make(chan uint8, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceUint8) Any(f func(el uint8) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceUint8) All(f func(el uint8) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceUint8) ChunkByBool(f func(el uint8) bool) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByByte(f func(el uint8) byte) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByString(f func(el uint8) string) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByFloat32(f func(el uint8) float32) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByFloat64(f func(el uint8) float64) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByInt(f func(el uint8) int) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByInt8(f func(el uint8) int8) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByInt16(f func(el uint8) int16) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByInt32(f func(el uint8) int32) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByInt64(f func(el uint8) int64) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByUint(f func(el uint8) uint) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByUint8(f func(el uint8) uint8) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByUint16(f func(el uint8) uint16) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByUint32(f func(el uint8) uint32) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByUint64(f func(el uint8) uint64) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkByInterface(f func(el uint8) interface{}) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint8) ChunkEvery(count int) [][]uint8 {
	chunks := make([][]uint8, 0)
	chunk := make([]uint8, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]uint8, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceUint8) Contains(el uint8) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceUint8) Count(el uint8) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceUint8) Cycle() chan uint8 {
	c := make(chan uint8, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceUint8) Dedup() []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceUint8) DedupByBool(f func(el uint8) bool) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByByte(f func(el uint8) byte) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByString(f func(el uint8) string) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByFloat32(f func(el uint8) float32) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByFloat64(f func(el uint8) float64) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByInt(f func(el uint8) int) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByInt8(f func(el uint8) int8) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByInt16(f func(el uint8) int16) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByInt32(f func(el uint8) int32) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByInt64(f func(el uint8) int64) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByUint(f func(el uint8) uint) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByUint8(f func(el uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByUint16(f func(el uint8) uint16) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByUint32(f func(el uint8) uint32) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByUint64(f func(el uint8) uint64) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DedupByInterface(f func(el uint8) interface{}) []uint8 {
	result := make([]uint8, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint8) DropEvery(nth int) []uint8 {
	result := make([]uint8, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceUint8) DropWhile(f func(arr uint8) bool) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceUint8) Each(f func(el uint8)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceUint8) Filter(f func(el uint8) bool) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceUint8) Find(def uint8, f func(el uint8) bool) uint8 {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceUint8) FindIndex(f func(el uint8) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceUint8) GroupByBool(f func(el uint8) bool) map[bool][]uint8 {
	result := make(map[bool][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByByte(f func(el uint8) byte) map[byte][]uint8 {
	result := make(map[byte][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByString(f func(el uint8) string) map[string][]uint8 {
	result := make(map[string][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByFloat32(f func(el uint8) float32) map[float32][]uint8 {
	result := make(map[float32][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByFloat64(f func(el uint8) float64) map[float64][]uint8 {
	result := make(map[float64][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByInt(f func(el uint8) int) map[int][]uint8 {
	result := make(map[int][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByInt8(f func(el uint8) int8) map[int8][]uint8 {
	result := make(map[int8][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByInt16(f func(el uint8) int16) map[int16][]uint8 {
	result := make(map[int16][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByInt32(f func(el uint8) int32) map[int32][]uint8 {
	result := make(map[int32][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByInt64(f func(el uint8) int64) map[int64][]uint8 {
	result := make(map[int64][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByUint(f func(el uint8) uint) map[uint][]uint8 {
	result := make(map[uint][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByUint8(f func(el uint8) uint8) map[uint8][]uint8 {
	result := make(map[uint8][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByUint16(f func(el uint8) uint16) map[uint16][]uint8 {
	result := make(map[uint16][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByUint32(f func(el uint8) uint32) map[uint32][]uint8 {
	result := make(map[uint32][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByUint64(f func(el uint8) uint64) map[uint64][]uint8 {
	result := make(map[uint64][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) GroupByInterface(f func(el uint8) interface{}) map[interface{}][]uint8 {
	result := make(map[interface{}][]uint8)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint8, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint8) Intersperse(el uint8) []uint8 {
	result := make([]uint8, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceUint8) MapBool(f func(el uint8) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapByte(f func(el uint8) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapString(f func(el uint8) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapFloat32(f func(el uint8) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapFloat64(f func(el uint8) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapInt(f func(el uint8) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapInt8(f func(el uint8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapInt16(f func(el uint8) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapInt32(f func(el uint8) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapInt64(f func(el uint8) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapUint(f func(el uint8) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapUint8(f func(el uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapUint16(f func(el uint8) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapUint32(f func(el uint8) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapUint64(f func(el uint8) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) MapInterface(f func(el uint8) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint8) Max() uint8 {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceUint8) Min() uint8 {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceUint8) Product(repeat int) chan []uint8 {
	c := make(chan []uint8, 1)
	go s.product(c, repeat, []uint8{}, 0)
	return c
}

func (s SliceUint8) product(c chan []uint8, repeat int, left []uint8, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]uint8, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]uint8, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceUint8) ReduceBool(acc bool, f func(el uint8, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceByte(acc byte, f func(el uint8, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceString(acc string, f func(el uint8, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceFloat32(acc float32, f func(el uint8, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceFloat64(acc float64, f func(el uint8, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceInt(acc int, f func(el uint8, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceInt8(acc int8, f func(el uint8, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceInt16(acc int16, f func(el uint8, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceInt32(acc int32, f func(el uint8, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceInt64(acc int64, f func(el uint8, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceUint(acc uint, f func(el uint8, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceUint8(acc uint8, f func(el uint8, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceUint16(acc uint16, f func(el uint8, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceUint32(acc uint32, f func(el uint8, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceUint64(acc uint64, f func(el uint8, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceInterface(acc interface{}, f func(el uint8, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint8) ReduceWhileBool(acc bool, f func(el uint8, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileByte(acc byte, f func(el uint8, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileString(acc string, f func(el uint8, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileFloat32(acc float32, f func(el uint8, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileFloat64(acc float64, f func(el uint8, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileInt(acc int, f func(el uint8, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileInt8(acc int8, f func(el uint8, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileInt16(acc int16, f func(el uint8, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileInt32(acc int32, f func(el uint8, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileInt64(acc int64, f func(el uint8, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileUint(acc uint, f func(el uint8, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileUint8(acc uint8, f func(el uint8, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileUint16(acc uint16, f func(el uint8, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileUint32(acc uint32, f func(el uint8, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileUint64(acc uint64, f func(el uint8, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) ReduceWhileInterface(acc interface{}, f func(el uint8, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint8) Reverse() []uint8 {
	result := make([]uint8, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceUint8) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceUint8) ScanBool(acc bool, f func(el uint8, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanByte(acc byte, f func(el uint8, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanString(acc string, f func(el uint8, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanFloat32(acc float32, f func(el uint8, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanFloat64(acc float64, f func(el uint8, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanInt(acc int, f func(el uint8, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanInt8(acc int8, f func(el uint8, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanInt16(acc int16, f func(el uint8, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanInt32(acc int32, f func(el uint8, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanInt64(acc int64, f func(el uint8, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanUint(acc uint, f func(el uint8, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanUint8(acc uint8, f func(el uint8, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanUint16(acc uint16, f func(el uint8, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanUint32(acc uint32, f func(el uint8, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanUint64(acc uint64, f func(el uint8, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) ScanInterface(acc interface{}, f func(el uint8, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint8) Shuffle() []uint8 {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceUint8) Sort() []uint8 {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceUint8) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceUint8) Split(sep uint8) [][]uint8 {
	result := make([][]uint8, 0)
	curr := make([]uint8, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceUint8) StartsWith(prefix []uint8) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceUint8) Sum() uint8 {
	var sum uint8
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceUint8) TakeWhile(f func(el uint8) bool) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceUint8) Uniq() []uint8 {
	added := make(map[uint8]struct{})
	nothing := struct{}{}
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceUint8) Window(size int) [][]uint8 {
	result := make([][]uint8, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesUint8) Concat() []uint8 {
	result := make([]uint8, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesUint8) Product() chan []uint8 {
	c := make(chan []uint8, 1)
	go s.product(c, []uint8{}, 0)
	return c
}

func (s SlicesUint8) product(c chan []uint8, left []uint8, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]uint8, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]uint8, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesUint8) Zip() [][]uint8 {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]uint8, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]uint8, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelUint16 struct {
	data chan uint16
}

type AsyncSliceUint16 struct {
	data    []uint16
	workers int
}

type SequenceUint16 struct {
	data chan uint16
}

type SliceUint16 struct {
	data []uint16
}

type SlicesUint16 struct {
	data [][]uint16
}

func (c ChannelUint16) Any(f func(el uint16) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelUint16) All(f func(el uint16) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelUint16) ChunkEvery(count int) chan []uint16 {
	chunks := make(chan []uint16, 1)
	go func() {
		chunk := make([]uint16, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]uint16, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelUint16) Count(el uint16) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelUint16) Drop(n int) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) Each(f func(el uint16)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelUint16) Filter(f func(el uint16) bool) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapBool(f func(el uint16) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapByte(f func(el uint16) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapString(f func(el uint16) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapFloat32(f func(el uint16) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapFloat64(f func(el uint16) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapInt(f func(el uint16) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapInt8(f func(el uint16) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapInt16(f func(el uint16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapInt32(f func(el uint16) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapInt64(f func(el uint16) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapUint(f func(el uint16) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapUint8(f func(el uint16) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapUint16(f func(el uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapUint32(f func(el uint16) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapUint64(f func(el uint16) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) MapInterface(f func(el uint16) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) Max() uint16 {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelUint16) Min() uint16 {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelUint16) ReduceBool(acc bool, f func(el uint16, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceByte(acc byte, f func(el uint16, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceString(acc string, f func(el uint16, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceFloat32(acc float32, f func(el uint16, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceFloat64(acc float64, f func(el uint16, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceInt(acc int, f func(el uint16, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceInt8(acc int8, f func(el uint16, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceInt16(acc int16, f func(el uint16, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceInt32(acc int32, f func(el uint16, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceInt64(acc int64, f func(el uint16, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceUint(acc uint, f func(el uint16, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceUint8(acc uint8, f func(el uint16, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceUint16(acc uint16, f func(el uint16, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceUint32(acc uint32, f func(el uint16, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceUint64(acc uint64, f func(el uint16, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ReduceInterface(acc interface{}, f func(el uint16, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint16) ScanBool(acc bool, f func(el uint16, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanByte(acc byte, f func(el uint16, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanString(acc string, f func(el uint16, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanFloat32(acc float32, f func(el uint16, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanFloat64(acc float64, f func(el uint16, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanInt(acc int, f func(el uint16, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanInt8(acc int8, f func(el uint16, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanInt16(acc int16, f func(el uint16, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanInt32(acc int32, f func(el uint16, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanInt64(acc int64, f func(el uint16, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanUint(acc uint, f func(el uint16, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanUint8(acc uint8, f func(el uint16, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanUint16(acc uint16, f func(el uint16, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanUint32(acc uint32, f func(el uint16, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanUint64(acc uint64, f func(el uint16, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) ScanInterface(acc interface{}, f func(el uint16, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint16) Sum() uint16 {
	var sum uint16
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelUint16) Take(n int) []uint16 {
	result := make([]uint16, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelUint16) Tee(count int) []chan uint16 {
	channels := make([]chan uint16, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan uint16, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan uint16) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelUint16) ToSlice() []uint16 {
	result := make([]uint16, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceUint16) Each(f func(el uint16)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceUint16) Filter(f func(el uint16) bool) []uint16 {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]uint16, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceUint16) MapBool(f func(el uint16) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapByte(f func(el uint16) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapString(f func(el uint16) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapFloat32(f func(el uint16) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapFloat64(f func(el uint16) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapInt(f func(el uint16) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapInt8(f func(el uint16) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapInt16(f func(el uint16) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapInt32(f func(el uint16) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapInt64(f func(el uint16) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapUint(f func(el uint16) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapUint8(f func(el uint16) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapUint16(f func(el uint16) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapUint32(f func(el uint16) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapUint64(f func(el uint16) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint16) MapInterface(f func(el uint16) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceUint16) Count(start uint16, step uint16) chan uint16 {
	c := make(chan uint16, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

func (s SequenceUint16) Exponential(start uint16, factor uint16) chan uint16 {
	c := make(chan uint16, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

func (s SequenceUint16) Range(start uint16, end uint16, step uint16) chan uint16 {
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

func (s SequenceUint16) Repeat(val uint16) chan uint16 {
	c := make(chan uint16, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceUint16) Any(f func(el uint16) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceUint16) All(f func(el uint16) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceUint16) ChunkByBool(f func(el uint16) bool) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByByte(f func(el uint16) byte) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByString(f func(el uint16) string) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByFloat32(f func(el uint16) float32) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByFloat64(f func(el uint16) float64) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByInt(f func(el uint16) int) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByInt8(f func(el uint16) int8) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByInt16(f func(el uint16) int16) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByInt32(f func(el uint16) int32) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByInt64(f func(el uint16) int64) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByUint(f func(el uint16) uint) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByUint8(f func(el uint16) uint8) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByUint16(f func(el uint16) uint16) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByUint32(f func(el uint16) uint32) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByUint64(f func(el uint16) uint64) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkByInterface(f func(el uint16) interface{}) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint16) ChunkEvery(count int) [][]uint16 {
	chunks := make([][]uint16, 0)
	chunk := make([]uint16, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]uint16, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceUint16) Contains(el uint16) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceUint16) Count(el uint16) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceUint16) Cycle() chan uint16 {
	c := make(chan uint16, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceUint16) Dedup() []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceUint16) DedupByBool(f func(el uint16) bool) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByByte(f func(el uint16) byte) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByString(f func(el uint16) string) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByFloat32(f func(el uint16) float32) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByFloat64(f func(el uint16) float64) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByInt(f func(el uint16) int) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByInt8(f func(el uint16) int8) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByInt16(f func(el uint16) int16) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByInt32(f func(el uint16) int32) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByInt64(f func(el uint16) int64) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByUint(f func(el uint16) uint) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByUint8(f func(el uint16) uint8) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByUint16(f func(el uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByUint32(f func(el uint16) uint32) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByUint64(f func(el uint16) uint64) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DedupByInterface(f func(el uint16) interface{}) []uint16 {
	result := make([]uint16, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint16) DropEvery(nth int) []uint16 {
	result := make([]uint16, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceUint16) DropWhile(f func(arr uint16) bool) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceUint16) Each(f func(el uint16)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceUint16) Filter(f func(el uint16) bool) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceUint16) Find(def uint16, f func(el uint16) bool) uint16 {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceUint16) FindIndex(f func(el uint16) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceUint16) GroupByBool(f func(el uint16) bool) map[bool][]uint16 {
	result := make(map[bool][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByByte(f func(el uint16) byte) map[byte][]uint16 {
	result := make(map[byte][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByString(f func(el uint16) string) map[string][]uint16 {
	result := make(map[string][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByFloat32(f func(el uint16) float32) map[float32][]uint16 {
	result := make(map[float32][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByFloat64(f func(el uint16) float64) map[float64][]uint16 {
	result := make(map[float64][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByInt(f func(el uint16) int) map[int][]uint16 {
	result := make(map[int][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByInt8(f func(el uint16) int8) map[int8][]uint16 {
	result := make(map[int8][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByInt16(f func(el uint16) int16) map[int16][]uint16 {
	result := make(map[int16][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByInt32(f func(el uint16) int32) map[int32][]uint16 {
	result := make(map[int32][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByInt64(f func(el uint16) int64) map[int64][]uint16 {
	result := make(map[int64][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByUint(f func(el uint16) uint) map[uint][]uint16 {
	result := make(map[uint][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByUint8(f func(el uint16) uint8) map[uint8][]uint16 {
	result := make(map[uint8][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByUint16(f func(el uint16) uint16) map[uint16][]uint16 {
	result := make(map[uint16][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByUint32(f func(el uint16) uint32) map[uint32][]uint16 {
	result := make(map[uint32][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByUint64(f func(el uint16) uint64) map[uint64][]uint16 {
	result := make(map[uint64][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) GroupByInterface(f func(el uint16) interface{}) map[interface{}][]uint16 {
	result := make(map[interface{}][]uint16)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint16, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint16) Intersperse(el uint16) []uint16 {
	result := make([]uint16, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceUint16) MapBool(f func(el uint16) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapByte(f func(el uint16) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapString(f func(el uint16) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapFloat32(f func(el uint16) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapFloat64(f func(el uint16) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapInt(f func(el uint16) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapInt8(f func(el uint16) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapInt16(f func(el uint16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapInt32(f func(el uint16) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapInt64(f func(el uint16) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapUint(f func(el uint16) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapUint8(f func(el uint16) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapUint16(f func(el uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapUint32(f func(el uint16) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapUint64(f func(el uint16) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) MapInterface(f func(el uint16) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint16) Max() uint16 {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceUint16) Min() uint16 {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceUint16) Product(repeat int) chan []uint16 {
	c := make(chan []uint16, 1)
	go s.product(c, repeat, []uint16{}, 0)
	return c
}

func (s SliceUint16) product(c chan []uint16, repeat int, left []uint16, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]uint16, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]uint16, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceUint16) ReduceBool(acc bool, f func(el uint16, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceByte(acc byte, f func(el uint16, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceString(acc string, f func(el uint16, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceFloat32(acc float32, f func(el uint16, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceFloat64(acc float64, f func(el uint16, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceInt(acc int, f func(el uint16, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceInt8(acc int8, f func(el uint16, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceInt16(acc int16, f func(el uint16, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceInt32(acc int32, f func(el uint16, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceInt64(acc int64, f func(el uint16, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceUint(acc uint, f func(el uint16, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceUint8(acc uint8, f func(el uint16, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceUint16(acc uint16, f func(el uint16, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceUint32(acc uint32, f func(el uint16, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceUint64(acc uint64, f func(el uint16, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceInterface(acc interface{}, f func(el uint16, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint16) ReduceWhileBool(acc bool, f func(el uint16, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileByte(acc byte, f func(el uint16, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileString(acc string, f func(el uint16, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileFloat32(acc float32, f func(el uint16, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileFloat64(acc float64, f func(el uint16, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileInt(acc int, f func(el uint16, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileInt8(acc int8, f func(el uint16, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileInt16(acc int16, f func(el uint16, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileInt32(acc int32, f func(el uint16, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileInt64(acc int64, f func(el uint16, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileUint(acc uint, f func(el uint16, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileUint8(acc uint8, f func(el uint16, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileUint16(acc uint16, f func(el uint16, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileUint32(acc uint32, f func(el uint16, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileUint64(acc uint64, f func(el uint16, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) ReduceWhileInterface(acc interface{}, f func(el uint16, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint16) Reverse() []uint16 {
	result := make([]uint16, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceUint16) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceUint16) ScanBool(acc bool, f func(el uint16, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanByte(acc byte, f func(el uint16, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanString(acc string, f func(el uint16, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanFloat32(acc float32, f func(el uint16, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanFloat64(acc float64, f func(el uint16, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanInt(acc int, f func(el uint16, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanInt8(acc int8, f func(el uint16, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanInt16(acc int16, f func(el uint16, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanInt32(acc int32, f func(el uint16, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanInt64(acc int64, f func(el uint16, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanUint(acc uint, f func(el uint16, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanUint8(acc uint8, f func(el uint16, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanUint16(acc uint16, f func(el uint16, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanUint32(acc uint32, f func(el uint16, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanUint64(acc uint64, f func(el uint16, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) ScanInterface(acc interface{}, f func(el uint16, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint16) Shuffle() []uint16 {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceUint16) Sort() []uint16 {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceUint16) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceUint16) Split(sep uint16) [][]uint16 {
	result := make([][]uint16, 0)
	curr := make([]uint16, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceUint16) StartsWith(prefix []uint16) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceUint16) Sum() uint16 {
	var sum uint16
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceUint16) TakeWhile(f func(el uint16) bool) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceUint16) Uniq() []uint16 {
	added := make(map[uint16]struct{})
	nothing := struct{}{}
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceUint16) Window(size int) [][]uint16 {
	result := make([][]uint16, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesUint16) Concat() []uint16 {
	result := make([]uint16, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesUint16) Product() chan []uint16 {
	c := make(chan []uint16, 1)
	go s.product(c, []uint16{}, 0)
	return c
}

func (s SlicesUint16) product(c chan []uint16, left []uint16, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]uint16, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]uint16, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesUint16) Zip() [][]uint16 {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]uint16, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]uint16, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelUint32 struct {
	data chan uint32
}

type AsyncSliceUint32 struct {
	data    []uint32
	workers int
}

type SequenceUint32 struct {
	data chan uint32
}

type SliceUint32 struct {
	data []uint32
}

type SlicesUint32 struct {
	data [][]uint32
}

func (c ChannelUint32) Any(f func(el uint32) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelUint32) All(f func(el uint32) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelUint32) ChunkEvery(count int) chan []uint32 {
	chunks := make(chan []uint32, 1)
	go func() {
		chunk := make([]uint32, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]uint32, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelUint32) Count(el uint32) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelUint32) Drop(n int) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) Each(f func(el uint32)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelUint32) Filter(f func(el uint32) bool) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapBool(f func(el uint32) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapByte(f func(el uint32) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapString(f func(el uint32) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapFloat32(f func(el uint32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapFloat64(f func(el uint32) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapInt(f func(el uint32) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapInt8(f func(el uint32) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapInt16(f func(el uint32) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapInt32(f func(el uint32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapInt64(f func(el uint32) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapUint(f func(el uint32) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapUint8(f func(el uint32) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapUint16(f func(el uint32) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapUint32(f func(el uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapUint64(f func(el uint32) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) MapInterface(f func(el uint32) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) Max() uint32 {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelUint32) Min() uint32 {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelUint32) ReduceBool(acc bool, f func(el uint32, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceByte(acc byte, f func(el uint32, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceString(acc string, f func(el uint32, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceFloat32(acc float32, f func(el uint32, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceFloat64(acc float64, f func(el uint32, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceInt(acc int, f func(el uint32, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceInt8(acc int8, f func(el uint32, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceInt16(acc int16, f func(el uint32, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceInt32(acc int32, f func(el uint32, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceInt64(acc int64, f func(el uint32, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceUint(acc uint, f func(el uint32, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceUint8(acc uint8, f func(el uint32, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceUint16(acc uint16, f func(el uint32, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceUint32(acc uint32, f func(el uint32, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceUint64(acc uint64, f func(el uint32, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ReduceInterface(acc interface{}, f func(el uint32, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint32) ScanBool(acc bool, f func(el uint32, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanByte(acc byte, f func(el uint32, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanString(acc string, f func(el uint32, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanFloat32(acc float32, f func(el uint32, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanFloat64(acc float64, f func(el uint32, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanInt(acc int, f func(el uint32, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanInt8(acc int8, f func(el uint32, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanInt16(acc int16, f func(el uint32, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanInt32(acc int32, f func(el uint32, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanInt64(acc int64, f func(el uint32, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanUint(acc uint, f func(el uint32, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanUint8(acc uint8, f func(el uint32, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanUint16(acc uint16, f func(el uint32, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanUint32(acc uint32, f func(el uint32, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanUint64(acc uint64, f func(el uint32, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) ScanInterface(acc interface{}, f func(el uint32, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint32) Sum() uint32 {
	var sum uint32
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelUint32) Take(n int) []uint32 {
	result := make([]uint32, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelUint32) Tee(count int) []chan uint32 {
	channels := make([]chan uint32, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan uint32, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan uint32) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelUint32) ToSlice() []uint32 {
	result := make([]uint32, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceUint32) Each(f func(el uint32)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceUint32) Filter(f func(el uint32) bool) []uint32 {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]uint32, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceUint32) MapBool(f func(el uint32) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapByte(f func(el uint32) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapString(f func(el uint32) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapFloat32(f func(el uint32) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapFloat64(f func(el uint32) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapInt(f func(el uint32) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapInt8(f func(el uint32) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapInt16(f func(el uint32) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapInt32(f func(el uint32) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapInt64(f func(el uint32) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapUint(f func(el uint32) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapUint8(f func(el uint32) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapUint16(f func(el uint32) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapUint32(f func(el uint32) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapUint64(f func(el uint32) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint32) MapInterface(f func(el uint32) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceUint32) Count(start uint32, step uint32) chan uint32 {
	c := make(chan uint32, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

func (s SequenceUint32) Exponential(start uint32, factor uint32) chan uint32 {
	c := make(chan uint32, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

func (s SequenceUint32) Range(start uint32, end uint32, step uint32) chan uint32 {
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

func (s SequenceUint32) Repeat(val uint32) chan uint32 {
	c := make(chan uint32, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceUint32) Any(f func(el uint32) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceUint32) All(f func(el uint32) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceUint32) ChunkByBool(f func(el uint32) bool) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByByte(f func(el uint32) byte) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByString(f func(el uint32) string) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByFloat32(f func(el uint32) float32) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByFloat64(f func(el uint32) float64) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByInt(f func(el uint32) int) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByInt8(f func(el uint32) int8) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByInt16(f func(el uint32) int16) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByInt32(f func(el uint32) int32) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByInt64(f func(el uint32) int64) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByUint(f func(el uint32) uint) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByUint8(f func(el uint32) uint8) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByUint16(f func(el uint32) uint16) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByUint32(f func(el uint32) uint32) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByUint64(f func(el uint32) uint64) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkByInterface(f func(el uint32) interface{}) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint32) ChunkEvery(count int) [][]uint32 {
	chunks := make([][]uint32, 0)
	chunk := make([]uint32, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]uint32, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceUint32) Contains(el uint32) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceUint32) Count(el uint32) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceUint32) Cycle() chan uint32 {
	c := make(chan uint32, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceUint32) Dedup() []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceUint32) DedupByBool(f func(el uint32) bool) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByByte(f func(el uint32) byte) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByString(f func(el uint32) string) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByFloat32(f func(el uint32) float32) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByFloat64(f func(el uint32) float64) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByInt(f func(el uint32) int) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByInt8(f func(el uint32) int8) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByInt16(f func(el uint32) int16) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByInt32(f func(el uint32) int32) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByInt64(f func(el uint32) int64) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByUint(f func(el uint32) uint) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByUint8(f func(el uint32) uint8) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByUint16(f func(el uint32) uint16) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByUint32(f func(el uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByUint64(f func(el uint32) uint64) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DedupByInterface(f func(el uint32) interface{}) []uint32 {
	result := make([]uint32, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint32) DropEvery(nth int) []uint32 {
	result := make([]uint32, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceUint32) DropWhile(f func(arr uint32) bool) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceUint32) Each(f func(el uint32)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceUint32) Filter(f func(el uint32) bool) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceUint32) Find(def uint32, f func(el uint32) bool) uint32 {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceUint32) FindIndex(f func(el uint32) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceUint32) GroupByBool(f func(el uint32) bool) map[bool][]uint32 {
	result := make(map[bool][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByByte(f func(el uint32) byte) map[byte][]uint32 {
	result := make(map[byte][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByString(f func(el uint32) string) map[string][]uint32 {
	result := make(map[string][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByFloat32(f func(el uint32) float32) map[float32][]uint32 {
	result := make(map[float32][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByFloat64(f func(el uint32) float64) map[float64][]uint32 {
	result := make(map[float64][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByInt(f func(el uint32) int) map[int][]uint32 {
	result := make(map[int][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByInt8(f func(el uint32) int8) map[int8][]uint32 {
	result := make(map[int8][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByInt16(f func(el uint32) int16) map[int16][]uint32 {
	result := make(map[int16][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByInt32(f func(el uint32) int32) map[int32][]uint32 {
	result := make(map[int32][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByInt64(f func(el uint32) int64) map[int64][]uint32 {
	result := make(map[int64][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByUint(f func(el uint32) uint) map[uint][]uint32 {
	result := make(map[uint][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByUint8(f func(el uint32) uint8) map[uint8][]uint32 {
	result := make(map[uint8][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByUint16(f func(el uint32) uint16) map[uint16][]uint32 {
	result := make(map[uint16][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByUint32(f func(el uint32) uint32) map[uint32][]uint32 {
	result := make(map[uint32][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByUint64(f func(el uint32) uint64) map[uint64][]uint32 {
	result := make(map[uint64][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) GroupByInterface(f func(el uint32) interface{}) map[interface{}][]uint32 {
	result := make(map[interface{}][]uint32)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint32, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint32) Intersperse(el uint32) []uint32 {
	result := make([]uint32, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceUint32) MapBool(f func(el uint32) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapByte(f func(el uint32) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapString(f func(el uint32) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapFloat32(f func(el uint32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapFloat64(f func(el uint32) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapInt(f func(el uint32) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapInt8(f func(el uint32) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapInt16(f func(el uint32) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapInt32(f func(el uint32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapInt64(f func(el uint32) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapUint(f func(el uint32) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapUint8(f func(el uint32) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapUint16(f func(el uint32) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapUint32(f func(el uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapUint64(f func(el uint32) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) MapInterface(f func(el uint32) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint32) Max() uint32 {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceUint32) Min() uint32 {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceUint32) Product(repeat int) chan []uint32 {
	c := make(chan []uint32, 1)
	go s.product(c, repeat, []uint32{}, 0)
	return c
}

func (s SliceUint32) product(c chan []uint32, repeat int, left []uint32, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]uint32, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]uint32, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceUint32) ReduceBool(acc bool, f func(el uint32, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceByte(acc byte, f func(el uint32, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceString(acc string, f func(el uint32, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceFloat32(acc float32, f func(el uint32, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceFloat64(acc float64, f func(el uint32, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceInt(acc int, f func(el uint32, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceInt8(acc int8, f func(el uint32, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceInt16(acc int16, f func(el uint32, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceInt32(acc int32, f func(el uint32, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceInt64(acc int64, f func(el uint32, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceUint(acc uint, f func(el uint32, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceUint8(acc uint8, f func(el uint32, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceUint16(acc uint16, f func(el uint32, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceUint32(acc uint32, f func(el uint32, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceUint64(acc uint64, f func(el uint32, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceInterface(acc interface{}, f func(el uint32, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint32) ReduceWhileBool(acc bool, f func(el uint32, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileByte(acc byte, f func(el uint32, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileString(acc string, f func(el uint32, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileFloat32(acc float32, f func(el uint32, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileFloat64(acc float64, f func(el uint32, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileInt(acc int, f func(el uint32, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileInt8(acc int8, f func(el uint32, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileInt16(acc int16, f func(el uint32, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileInt32(acc int32, f func(el uint32, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileInt64(acc int64, f func(el uint32, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileUint(acc uint, f func(el uint32, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileUint8(acc uint8, f func(el uint32, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileUint16(acc uint16, f func(el uint32, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileUint32(acc uint32, f func(el uint32, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileUint64(acc uint64, f func(el uint32, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) ReduceWhileInterface(acc interface{}, f func(el uint32, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint32) Reverse() []uint32 {
	result := make([]uint32, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceUint32) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceUint32) ScanBool(acc bool, f func(el uint32, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanByte(acc byte, f func(el uint32, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanString(acc string, f func(el uint32, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanFloat32(acc float32, f func(el uint32, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanFloat64(acc float64, f func(el uint32, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanInt(acc int, f func(el uint32, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanInt8(acc int8, f func(el uint32, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanInt16(acc int16, f func(el uint32, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanInt32(acc int32, f func(el uint32, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanInt64(acc int64, f func(el uint32, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanUint(acc uint, f func(el uint32, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanUint8(acc uint8, f func(el uint32, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanUint16(acc uint16, f func(el uint32, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanUint32(acc uint32, f func(el uint32, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanUint64(acc uint64, f func(el uint32, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) ScanInterface(acc interface{}, f func(el uint32, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint32) Shuffle() []uint32 {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceUint32) Sort() []uint32 {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceUint32) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceUint32) Split(sep uint32) [][]uint32 {
	result := make([][]uint32, 0)
	curr := make([]uint32, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceUint32) StartsWith(prefix []uint32) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceUint32) Sum() uint32 {
	var sum uint32
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceUint32) TakeWhile(f func(el uint32) bool) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceUint32) Uniq() []uint32 {
	added := make(map[uint32]struct{})
	nothing := struct{}{}
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceUint32) Window(size int) [][]uint32 {
	result := make([][]uint32, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesUint32) Concat() []uint32 {
	result := make([]uint32, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesUint32) Product() chan []uint32 {
	c := make(chan []uint32, 1)
	go s.product(c, []uint32{}, 0)
	return c
}

func (s SlicesUint32) product(c chan []uint32, left []uint32, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]uint32, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]uint32, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesUint32) Zip() [][]uint32 {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]uint32, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]uint32, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelUint64 struct {
	data chan uint64
}

type AsyncSliceUint64 struct {
	data    []uint64
	workers int
}

type SequenceUint64 struct {
	data chan uint64
}

type SliceUint64 struct {
	data []uint64
}

type SlicesUint64 struct {
	data [][]uint64
}

func (c ChannelUint64) Any(f func(el uint64) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelUint64) All(f func(el uint64) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelUint64) ChunkEvery(count int) chan []uint64 {
	chunks := make(chan []uint64, 1)
	go func() {
		chunk := make([]uint64, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]uint64, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelUint64) Count(el uint64) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelUint64) Drop(n int) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) Each(f func(el uint64)) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelUint64) Filter(f func(el uint64) bool) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapBool(f func(el uint64) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapByte(f func(el uint64) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapString(f func(el uint64) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapFloat32(f func(el uint64) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapFloat64(f func(el uint64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapInt(f func(el uint64) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapInt8(f func(el uint64) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapInt16(f func(el uint64) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapInt32(f func(el uint64) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapInt64(f func(el uint64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapUint(f func(el uint64) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapUint8(f func(el uint64) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapUint16(f func(el uint64) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapUint32(f func(el uint64) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapUint64(f func(el uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) MapInterface(f func(el uint64) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) Max() uint64 {
	max := <-c.data
	for el := range c.data {
		if el > max {
			max = el
		}
	}
	return max
}

func (c ChannelUint64) Min() uint64 {
	min := <-c.data
	for el := range c.data {
		if el < min {
			min = el
		}
	}
	return min
}

func (c ChannelUint64) ReduceBool(acc bool, f func(el uint64, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceByte(acc byte, f func(el uint64, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceString(acc string, f func(el uint64, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceFloat32(acc float32, f func(el uint64, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceFloat64(acc float64, f func(el uint64, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceInt(acc int, f func(el uint64, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceInt8(acc int8, f func(el uint64, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceInt16(acc int16, f func(el uint64, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceInt32(acc int32, f func(el uint64, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceInt64(acc int64, f func(el uint64, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceUint(acc uint, f func(el uint64, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceUint8(acc uint8, f func(el uint64, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceUint16(acc uint16, f func(el uint64, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceUint32(acc uint32, f func(el uint64, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceUint64(acc uint64, f func(el uint64, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ReduceInterface(acc interface{}, f func(el uint64, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelUint64) ScanBool(acc bool, f func(el uint64, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanByte(acc byte, f func(el uint64, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanString(acc string, f func(el uint64, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanFloat32(acc float32, f func(el uint64, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanFloat64(acc float64, f func(el uint64, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanInt(acc int, f func(el uint64, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanInt8(acc int8, f func(el uint64, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanInt16(acc int16, f func(el uint64, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanInt32(acc int32, f func(el uint64, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanInt64(acc int64, f func(el uint64, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanUint(acc uint, f func(el uint64, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanUint8(acc uint8, f func(el uint64, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanUint16(acc uint16, f func(el uint64, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanUint32(acc uint32, f func(el uint64, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanUint64(acc uint64, f func(el uint64, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) ScanInterface(acc interface{}, f func(el uint64, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelUint64) Sum() uint64 {
	var sum uint64
	for el := range c.data {
		sum += el
	}
	return sum
}

func (c ChannelUint64) Take(n int) []uint64 {
	result := make([]uint64, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelUint64) Tee(count int) []chan uint64 {
	channels := make([]chan uint64, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan uint64, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan uint64) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelUint64) ToSlice() []uint64 {
	result := make([]uint64, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceUint64) Each(f func(el uint64)) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceUint64) Filter(f func(el uint64) bool) []uint64 {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]uint64, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceUint64) MapBool(f func(el uint64) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapByte(f func(el uint64) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapString(f func(el uint64) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapFloat32(f func(el uint64) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapFloat64(f func(el uint64) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapInt(f func(el uint64) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapInt8(f func(el uint64) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapInt16(f func(el uint64) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapInt32(f func(el uint64) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapInt64(f func(el uint64) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapUint(f func(el uint64) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapUint8(f func(el uint64) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapUint16(f func(el uint64) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapUint32(f func(el uint64) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapUint64(f func(el uint64) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceUint64) MapInterface(f func(el uint64) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceUint64) Count(start uint64, step uint64) chan uint64 {
	c := make(chan uint64, 1)
	go func() {
		for {
			c <- start
			start += step
		}
	}()
	return c
}

func (s SequenceUint64) Exponential(start uint64, factor uint64) chan uint64 {
	c := make(chan uint64, 1)
	go func() {
		for {
			c <- start
			start *= factor
		}
	}()
	return c
}

func (s SequenceUint64) Range(start uint64, end uint64, step uint64) chan uint64 {
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

func (s SequenceUint64) Repeat(val uint64) chan uint64 {
	c := make(chan uint64, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceUint64) Any(f func(el uint64) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceUint64) All(f func(el uint64) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceUint64) ChunkByBool(f func(el uint64) bool) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByByte(f func(el uint64) byte) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByString(f func(el uint64) string) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByFloat32(f func(el uint64) float32) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByFloat64(f func(el uint64) float64) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByInt(f func(el uint64) int) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByInt8(f func(el uint64) int8) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByInt16(f func(el uint64) int16) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByInt32(f func(el uint64) int32) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByInt64(f func(el uint64) int64) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByUint(f func(el uint64) uint) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByUint8(f func(el uint64) uint8) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByUint16(f func(el uint64) uint16) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByUint32(f func(el uint64) uint32) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByUint64(f func(el uint64) uint64) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkByInterface(f func(el uint64) interface{}) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceUint64) ChunkEvery(count int) [][]uint64 {
	chunks := make([][]uint64, 0)
	chunk := make([]uint64, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]uint64, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceUint64) Contains(el uint64) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceUint64) Count(el uint64) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceUint64) Cycle() chan uint64 {
	c := make(chan uint64, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceUint64) Dedup() []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceUint64) DedupByBool(f func(el uint64) bool) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByByte(f func(el uint64) byte) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByString(f func(el uint64) string) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByFloat32(f func(el uint64) float32) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByFloat64(f func(el uint64) float64) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByInt(f func(el uint64) int) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByInt8(f func(el uint64) int8) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByInt16(f func(el uint64) int16) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByInt32(f func(el uint64) int32) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByInt64(f func(el uint64) int64) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByUint(f func(el uint64) uint) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByUint8(f func(el uint64) uint8) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByUint16(f func(el uint64) uint16) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByUint32(f func(el uint64) uint32) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByUint64(f func(el uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DedupByInterface(f func(el uint64) interface{}) []uint64 {
	result := make([]uint64, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceUint64) DropEvery(nth int) []uint64 {
	result := make([]uint64, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceUint64) DropWhile(f func(arr uint64) bool) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceUint64) Each(f func(el uint64)) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceUint64) Filter(f func(el uint64) bool) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceUint64) Find(def uint64, f func(el uint64) bool) uint64 {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceUint64) FindIndex(f func(el uint64) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceUint64) GroupByBool(f func(el uint64) bool) map[bool][]uint64 {
	result := make(map[bool][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByByte(f func(el uint64) byte) map[byte][]uint64 {
	result := make(map[byte][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByString(f func(el uint64) string) map[string][]uint64 {
	result := make(map[string][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByFloat32(f func(el uint64) float32) map[float32][]uint64 {
	result := make(map[float32][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByFloat64(f func(el uint64) float64) map[float64][]uint64 {
	result := make(map[float64][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByInt(f func(el uint64) int) map[int][]uint64 {
	result := make(map[int][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByInt8(f func(el uint64) int8) map[int8][]uint64 {
	result := make(map[int8][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByInt16(f func(el uint64) int16) map[int16][]uint64 {
	result := make(map[int16][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByInt32(f func(el uint64) int32) map[int32][]uint64 {
	result := make(map[int32][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByInt64(f func(el uint64) int64) map[int64][]uint64 {
	result := make(map[int64][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByUint(f func(el uint64) uint) map[uint][]uint64 {
	result := make(map[uint][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByUint8(f func(el uint64) uint8) map[uint8][]uint64 {
	result := make(map[uint8][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByUint16(f func(el uint64) uint16) map[uint16][]uint64 {
	result := make(map[uint16][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByUint32(f func(el uint64) uint32) map[uint32][]uint64 {
	result := make(map[uint32][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByUint64(f func(el uint64) uint64) map[uint64][]uint64 {
	result := make(map[uint64][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) GroupByInterface(f func(el uint64) interface{}) map[interface{}][]uint64 {
	result := make(map[interface{}][]uint64)
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]uint64, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceUint64) Intersperse(el uint64) []uint64 {
	result := make([]uint64, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceUint64) MapBool(f func(el uint64) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapByte(f func(el uint64) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapString(f func(el uint64) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapFloat32(f func(el uint64) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapFloat64(f func(el uint64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapInt(f func(el uint64) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapInt8(f func(el uint64) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapInt16(f func(el uint64) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapInt32(f func(el uint64) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapInt64(f func(el uint64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapUint(f func(el uint64) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapUint8(f func(el uint64) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapUint16(f func(el uint64) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapUint32(f func(el uint64) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapUint64(f func(el uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) MapInterface(f func(el uint64) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceUint64) Max() uint64 {
	max := s.data[0]
	for _, el := range s.data[1:] {
		if el > max {
			max = el
		}
	}
	return max
}

func (s SliceUint64) Min() uint64 {
	min := s.data[0]
	for _, el := range s.data[1:] {
		if el < min {
			min = el
		}
	}
	return min
}

func (s SliceUint64) Product(repeat int) chan []uint64 {
	c := make(chan []uint64, 1)
	go s.product(c, repeat, []uint64{}, 0)
	return c
}

func (s SliceUint64) product(c chan []uint64, repeat int, left []uint64, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]uint64, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]uint64, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceUint64) ReduceBool(acc bool, f func(el uint64, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceByte(acc byte, f func(el uint64, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceString(acc string, f func(el uint64, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceFloat32(acc float32, f func(el uint64, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceFloat64(acc float64, f func(el uint64, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceInt(acc int, f func(el uint64, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceInt8(acc int8, f func(el uint64, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceInt16(acc int16, f func(el uint64, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceInt32(acc int32, f func(el uint64, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceInt64(acc int64, f func(el uint64, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceUint(acc uint, f func(el uint64, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceUint8(acc uint8, f func(el uint64, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceUint16(acc uint16, f func(el uint64, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceUint32(acc uint32, f func(el uint64, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceUint64(acc uint64, f func(el uint64, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceInterface(acc interface{}, f func(el uint64, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceUint64) ReduceWhileBool(acc bool, f func(el uint64, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileByte(acc byte, f func(el uint64, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileString(acc string, f func(el uint64, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileFloat32(acc float32, f func(el uint64, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileFloat64(acc float64, f func(el uint64, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileInt(acc int, f func(el uint64, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileInt8(acc int8, f func(el uint64, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileInt16(acc int16, f func(el uint64, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileInt32(acc int32, f func(el uint64, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileInt64(acc int64, f func(el uint64, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileUint(acc uint, f func(el uint64, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileUint8(acc uint8, f func(el uint64, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileUint16(acc uint16, f func(el uint64, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileUint32(acc uint32, f func(el uint64, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileUint64(acc uint64, f func(el uint64, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) ReduceWhileInterface(acc interface{}, f func(el uint64, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceUint64) Reverse() []uint64 {
	result := make([]uint64, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceUint64) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceUint64) ScanBool(acc bool, f func(el uint64, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanByte(acc byte, f func(el uint64, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanString(acc string, f func(el uint64, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanFloat32(acc float32, f func(el uint64, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanFloat64(acc float64, f func(el uint64, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanInt(acc int, f func(el uint64, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanInt8(acc int8, f func(el uint64, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanInt16(acc int16, f func(el uint64, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanInt32(acc int32, f func(el uint64, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanInt64(acc int64, f func(el uint64, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanUint(acc uint, f func(el uint64, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanUint8(acc uint8, f func(el uint64, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanUint16(acc uint16, f func(el uint64, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanUint32(acc uint32, f func(el uint64, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanUint64(acc uint64, f func(el uint64, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) ScanInterface(acc interface{}, f func(el uint64, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceUint64) Shuffle() []uint64 {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceUint64) Sort() []uint64 {
	less := func(i int, j int) bool {
		return s.data[i] < s.data[j]
	}
	sort.SliceStable(s.data, less)
	return s.data
}

func (s SliceUint64) Sorted() bool {
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceUint64) Split(sep uint64) [][]uint64 {
	result := make([][]uint64, 0)
	curr := make([]uint64, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceUint64) StartsWith(prefix []uint64) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceUint64) Sum() uint64 {
	var sum uint64
	for _, el := range s.data {
		sum += el
	}
	return sum
}

func (s SliceUint64) TakeWhile(f func(el uint64) bool) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceUint64) Uniq() []uint64 {
	added := make(map[uint64]struct{})
	nothing := struct{}{}
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceUint64) Window(size int) [][]uint64 {
	result := make([][]uint64, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesUint64) Concat() []uint64 {
	result := make([]uint64, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesUint64) Product() chan []uint64 {
	c := make(chan []uint64, 1)
	go s.product(c, []uint64{}, 0)
	return c
}

func (s SlicesUint64) product(c chan []uint64, left []uint64, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]uint64, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]uint64, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesUint64) Zip() [][]uint64 {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]uint64, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]uint64, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}

type ChannelInterface struct {
	data chan interface{}
}

type AsyncSliceInterface struct {
	data    []interface{}
	workers int
}

type SequenceInterface struct {
	data chan interface{}
}

type SliceInterface struct {
	data []interface{}
}

type SlicesInterface struct {
	data [][]interface{}
}

func (c ChannelInterface) Any(f func(el interface{}) bool) bool {
	for el := range c.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (c ChannelInterface) All(f func(el interface{}) bool) bool {
	for el := range c.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (c ChannelInterface) ChunkEvery(count int) chan []interface{} {
	chunks := make(chan []interface{}, 1)
	go func() {
		chunk := make([]interface{}, 0, count)
		i := 0
		for el := range c.data {
			chunk = append(chunk, el)
			i++
			if i%count == 0 {
				i = 0
				chunks <- chunk
				chunk = make([]interface{}, 0, count)
			}
		}
		if len(chunk) > 0 {
			chunks <- chunk
		}
		close(chunks)
	}()
	return chunks
}

func (c ChannelInterface) Count(el interface{}) int {
	count := 0
	for val := range c.data {
		if val == el {
			count++
		}
	}
	return count
}

func (c ChannelInterface) Drop(n int) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) Each(f func(el interface{})) {
	for el := range c.data {
		f(el)
	}
}

func (c ChannelInterface) Filter(f func(el interface{}) bool) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapBool(f func(el interface{}) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapByte(f func(el interface{}) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapString(f func(el interface{}) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapFloat32(f func(el interface{}) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapFloat64(f func(el interface{}) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapInt(f func(el interface{}) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapInt8(f func(el interface{}) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapInt16(f func(el interface{}) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapInt32(f func(el interface{}) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapInt64(f func(el interface{}) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapUint(f func(el interface{}) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapUint8(f func(el interface{}) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapUint16(f func(el interface{}) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapUint32(f func(el interface{}) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapUint64(f func(el interface{}) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) MapInterface(f func(el interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			result <- f(el)
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ReduceBool(acc bool, f func(el interface{}, acc bool) bool) bool {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceByte(acc byte, f func(el interface{}, acc byte) byte) byte {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceString(acc string, f func(el interface{}, acc string) string) string {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceFloat32(acc float32, f func(el interface{}, acc float32) float32) float32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceFloat64(acc float64, f func(el interface{}, acc float64) float64) float64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceInt(acc int, f func(el interface{}, acc int) int) int {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceInt8(acc int8, f func(el interface{}, acc int8) int8) int8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceInt16(acc int16, f func(el interface{}, acc int16) int16) int16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceInt32(acc int32, f func(el interface{}, acc int32) int32) int32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceInt64(acc int64, f func(el interface{}, acc int64) int64) int64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceUint(acc uint, f func(el interface{}, acc uint) uint) uint {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceUint8(acc uint8, f func(el interface{}, acc uint8) uint8) uint8 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceUint16(acc uint16, f func(el interface{}, acc uint16) uint16) uint16 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceUint32(acc uint32, f func(el interface{}, acc uint32) uint32) uint32 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceUint64(acc uint64, f func(el interface{}, acc uint64) uint64) uint64 {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ReduceInterface(acc interface{}, f func(el interface{}, acc interface{}) interface{}) interface{} {
	for el := range c.data {
		acc = f(el, acc)
	}
	return acc
}

func (c ChannelInterface) ScanBool(acc bool, f func(el interface{}, acc bool) bool) chan bool {
	result := make(chan bool, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanByte(acc byte, f func(el interface{}, acc byte) byte) chan byte {
	result := make(chan byte, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanString(acc string, f func(el interface{}, acc string) string) chan string {
	result := make(chan string, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanFloat32(acc float32, f func(el interface{}, acc float32) float32) chan float32 {
	result := make(chan float32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanFloat64(acc float64, f func(el interface{}, acc float64) float64) chan float64 {
	result := make(chan float64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanInt(acc int, f func(el interface{}, acc int) int) chan int {
	result := make(chan int, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanInt8(acc int8, f func(el interface{}, acc int8) int8) chan int8 {
	result := make(chan int8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanInt16(acc int16, f func(el interface{}, acc int16) int16) chan int16 {
	result := make(chan int16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanInt32(acc int32, f func(el interface{}, acc int32) int32) chan int32 {
	result := make(chan int32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanInt64(acc int64, f func(el interface{}, acc int64) int64) chan int64 {
	result := make(chan int64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanUint(acc uint, f func(el interface{}, acc uint) uint) chan uint {
	result := make(chan uint, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanUint8(acc uint8, f func(el interface{}, acc uint8) uint8) chan uint8 {
	result := make(chan uint8, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanUint16(acc uint16, f func(el interface{}, acc uint16) uint16) chan uint16 {
	result := make(chan uint16, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanUint32(acc uint32, f func(el interface{}, acc uint32) uint32) chan uint32 {
	result := make(chan uint32, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanUint64(acc uint64, f func(el interface{}, acc uint64) uint64) chan uint64 {
	result := make(chan uint64, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) ScanInterface(acc interface{}, f func(el interface{}, acc interface{}) interface{}) chan interface{} {
	result := make(chan interface{}, 1)
	go func() {
		for el := range c.data {
			acc = f(el, acc)
			result <- acc
		}
		close(result)
	}()
	return result
}

func (c ChannelInterface) Take(n int) []interface{} {
	result := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

func (c ChannelInterface) Tee(count int) []chan interface{} {
	channels := make([]chan interface{}, 0, count)
	for i := 0; i < count; i++ {
		channels = append(channels, make(chan interface{}, 1))
	}
	go func() {
		for el := range c.data {
			wg := sync.WaitGroup{}
			putInto := func(ch chan interface{}) {
				wg.Add(1)
				defer wg.Done()
				ch <- el
			}
			for _, ch := range channels {
				putInto(ch)
			}
			wg.Wait()
		}
		for _, ch := range channels {
			close(ch)
		}
	}()
	return channels
}

func (c ChannelInterface) ToSlice() []interface{} {
	result := make([]interface{}, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}

func (s AsyncSliceInterface) Each(f func(el interface{})) {
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func (s AsyncSliceInterface) Filter(f func(el interface{}) bool) []interface{} {
	resultMap := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			if f(s.data[index]) {
				resultMap[index] = true
			}
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()

	// return filtered results
	result := make([]interface{}, 0, len(s.data))
	for i, el := range s.data {
		if resultMap[i] {
			result = append(result, el)
		}
	}
	return result
}

func (s AsyncSliceInterface) MapBool(f func(el interface{}) bool) []bool {
	result := make([]bool, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapByte(f func(el interface{}) byte) []byte {
	result := make([]byte, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapString(f func(el interface{}) string) []string {
	result := make([]string, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapFloat32(f func(el interface{}) float32) []float32 {
	result := make([]float32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapFloat64(f func(el interface{}) float64) []float64 {
	result := make([]float64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapInt(f func(el interface{}) int) []int {
	result := make([]int, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapInt8(f func(el interface{}) int8) []int8 {
	result := make([]int8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapInt16(f func(el interface{}) int16) []int16 {
	result := make([]int16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapInt32(f func(el interface{}) int32) []int32 {
	result := make([]int32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapInt64(f func(el interface{}) int64) []int64 {
	result := make([]int64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapUint(f func(el interface{}) uint) []uint {
	result := make([]uint, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapUint8(f func(el interface{}) uint8) []uint8 {
	result := make([]uint8, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapUint16(f func(el interface{}) uint16) []uint16 {
	result := make([]uint16, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapUint32(f func(el interface{}) uint32) []uint32 {
	result := make([]uint32, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapUint64(f func(el interface{}) uint64) []uint64 {
	result := make([]uint64, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s AsyncSliceInterface) MapInterface(f func(el interface{}) interface{}) []interface{} {
	result := make([]interface{}, len(s.data))
	wg := sync.WaitGroup{}

	worker := func(jobs <-chan int) {
		wg.Add(1)
		for index := range jobs {
			result[index] = f(s.data[index])
		}
		wg.Done()
	}

	// run workers
	jobs := make(chan int, len(s.data))
	workers := s.workers
	if workers == 0 {
		workers = len(s.data)
	}
	for i := 0; i < s.workers; i++ {
		go worker(jobs)
	}

	// add indices into jobs for workers
	for i := 0; i < len(s.data); i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func (s SequenceInterface) Repeat(val interface{}) chan interface{} {
	c := make(chan interface{}, 1)
	go func() {
		for {
			c <- val
		}
	}()
	return c
}

func (s SliceInterface) Any(f func(el interface{}) bool) bool {
	for _, el := range s.data {
		if f(el) {
			return true
		}
	}
	return false
}

func (s SliceInterface) All(f func(el interface{}) bool) bool {
	for _, el := range s.data {
		if !f(el) {
			return false
		}
	}
	return true
}

func (s SliceInterface) ChunkByBool(f func(el interface{}) bool) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByByte(f func(el interface{}) byte) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByString(f func(el interface{}) string) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByFloat32(f func(el interface{}) float32) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByFloat64(f func(el interface{}) float64) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByInt(f func(el interface{}) int) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByInt8(f func(el interface{}) int8) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByInt16(f func(el interface{}) int16) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByInt32(f func(el interface{}) int32) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByInt64(f func(el interface{}) int64) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByUint(f func(el interface{}) uint) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByUint8(f func(el interface{}) uint8) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByUint16(f func(el interface{}) uint16) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByUint32(f func(el interface{}) uint32) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByUint64(f func(el interface{}) uint64) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkByInterface(f func(el interface{}) interface{}) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0)

	prev := f(s.data[0])
	chunk = append(chunk, s.data[0])

	for _, el := range s.data[1:] {
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

func (s SliceInterface) ChunkEvery(count int) [][]interface{} {
	chunks := make([][]interface{}, 0)
	chunk := make([]interface{}, 0, count)
	for i, el := range s.data {
		chunk = append(chunk, el)
		if (i+1)%count == 0 {
			chunks = append(chunks, chunk)
			chunk = make([]interface{}, 0, count)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (s SliceInterface) Contains(el interface{}) bool {
	for _, val := range s.data {
		if val == el {
			return true
		}
	}
	return false
}

func (s SliceInterface) Count(el interface{}) int {
	count := 0
	for _, val := range s.data {
		if val == el {
			count++
		}
	}
	return count
}

func (s SliceInterface) Cycle() chan interface{} {
	c := make(chan interface{}, 1)
	go func() {
		for {
			for _, val := range s.data {
				c <- val
			}
		}
	}()
	return c
}

func (s SliceInterface) Dedup() []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := s.data[0]
	result = append(result, prev)
	for _, el := range s.data[1:] {
		if el != prev {
			result = append(result, el)
			prev = el
		}
	}
	return result
}

func (s SliceInterface) DedupByBool(f func(el interface{}) bool) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByByte(f func(el interface{}) byte) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByString(f func(el interface{}) string) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByFloat32(f func(el interface{}) float32) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByFloat64(f func(el interface{}) float64) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByInt(f func(el interface{}) int) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByInt8(f func(el interface{}) int8) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByInt16(f func(el interface{}) int16) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByInt32(f func(el interface{}) int32) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByInt64(f func(el interface{}) int64) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByUint(f func(el interface{}) uint) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByUint8(f func(el interface{}) uint8) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByUint16(f func(el interface{}) uint16) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByUint32(f func(el interface{}) uint32) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByUint64(f func(el interface{}) uint64) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DedupByInterface(f func(el interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))

	prev := f(s.data[0])
	result = append(result, s.data[0])
	for _, el := range s.data[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

func (s SliceInterface) DropEvery(nth int) []interface{} {
	result := make([]interface{}, 0, len(s.data)/nth)
	for i, el := range s.data {
		if (i+1)%nth != 0 {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceInterface) DropWhile(f func(arr interface{}) bool) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceInterface) Each(f func(el interface{})) {
	for _, el := range s.data {
		f(el)
	}
}

func (s SliceInterface) Filter(f func(el interface{}) bool) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		if f(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s SliceInterface) Find(def interface{}, f func(el interface{}) bool) interface{} {
	for _, el := range s.data {
		if f(el) {
			return el
		}
	}
	return def
}

func (s SliceInterface) FindIndex(f func(el interface{}) bool) int {
	for i, el := range s.data {
		if f(el) {
			return i
		}
	}
	return -1
}

func (s SliceInterface) GroupByBool(f func(el interface{}) bool) map[bool][]interface{} {
	result := make(map[bool][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByByte(f func(el interface{}) byte) map[byte][]interface{} {
	result := make(map[byte][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByString(f func(el interface{}) string) map[string][]interface{} {
	result := make(map[string][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByFloat32(f func(el interface{}) float32) map[float32][]interface{} {
	result := make(map[float32][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByFloat64(f func(el interface{}) float64) map[float64][]interface{} {
	result := make(map[float64][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByInt(f func(el interface{}) int) map[int][]interface{} {
	result := make(map[int][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByInt8(f func(el interface{}) int8) map[int8][]interface{} {
	result := make(map[int8][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByInt16(f func(el interface{}) int16) map[int16][]interface{} {
	result := make(map[int16][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByInt32(f func(el interface{}) int32) map[int32][]interface{} {
	result := make(map[int32][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByInt64(f func(el interface{}) int64) map[int64][]interface{} {
	result := make(map[int64][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByUint(f func(el interface{}) uint) map[uint][]interface{} {
	result := make(map[uint][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByUint8(f func(el interface{}) uint8) map[uint8][]interface{} {
	result := make(map[uint8][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByUint16(f func(el interface{}) uint16) map[uint16][]interface{} {
	result := make(map[uint16][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByUint32(f func(el interface{}) uint32) map[uint32][]interface{} {
	result := make(map[uint32][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByUint64(f func(el interface{}) uint64) map[uint64][]interface{} {
	result := make(map[uint64][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) GroupByInterface(f func(el interface{}) interface{}) map[interface{}][]interface{} {
	result := make(map[interface{}][]interface{})
	for _, el := range s.data {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make([]interface{}, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

func (s SliceInterface) Intersperse(el interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data)*2-1)
	result = append(result, s.data[0])
	for _, val := range s.data[1:] {
		result = append(result, el, val)
	}
	return result
}

func (s SliceInterface) MapBool(f func(el interface{}) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapByte(f func(el interface{}) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapString(f func(el interface{}) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapFloat32(f func(el interface{}) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapFloat64(f func(el interface{}) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapInt(f func(el interface{}) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapInt8(f func(el interface{}) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapInt16(f func(el interface{}) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapInt32(f func(el interface{}) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapInt64(f func(el interface{}) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapUint(f func(el interface{}) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapUint8(f func(el interface{}) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapUint16(f func(el interface{}) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapUint32(f func(el interface{}) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapUint64(f func(el interface{}) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) MapInterface(f func(el interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		result = append(result, f(el))
	}
	return result
}

func (s SliceInterface) Product(repeat int) chan []interface{} {
	c := make(chan []interface{}, 1)
	go s.product(c, repeat, []interface{}{}, 0)
	return c
}

func (s SliceInterface) product(c chan []interface{}, repeat int, left []interface{}, pos int) {
	// iterate over the last array
	if pos == repeat-1 {
		for _, el := range s.data {
			result := make([]interface{}, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data {
		result := make([]interface{}, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, repeat, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SliceInterface) ReduceBool(acc bool, f func(el interface{}, acc bool) bool) bool {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceByte(acc byte, f func(el interface{}, acc byte) byte) byte {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceString(acc string, f func(el interface{}, acc string) string) string {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceFloat32(acc float32, f func(el interface{}, acc float32) float32) float32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceFloat64(acc float64, f func(el interface{}, acc float64) float64) float64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceInt(acc int, f func(el interface{}, acc int) int) int {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceInt8(acc int8, f func(el interface{}, acc int8) int8) int8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceInt16(acc int16, f func(el interface{}, acc int16) int16) int16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceInt32(acc int32, f func(el interface{}, acc int32) int32) int32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceInt64(acc int64, f func(el interface{}, acc int64) int64) int64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceUint(acc uint, f func(el interface{}, acc uint) uint) uint {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceUint8(acc uint8, f func(el interface{}, acc uint8) uint8) uint8 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceUint16(acc uint16, f func(el interface{}, acc uint16) uint16) uint16 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceUint32(acc uint32, f func(el interface{}, acc uint32) uint32) uint32 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceUint64(acc uint64, f func(el interface{}, acc uint64) uint64) uint64 {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceInterface(acc interface{}, f func(el interface{}, acc interface{}) interface{}) interface{} {
	for _, el := range s.data {
		acc = f(el, acc)
	}
	return acc
}

func (s SliceInterface) ReduceWhileBool(acc bool, f func(el interface{}, acc bool) (bool, error)) (bool, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileByte(acc byte, f func(el interface{}, acc byte) (byte, error)) (byte, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileString(acc string, f func(el interface{}, acc string) (string, error)) (string, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileFloat32(acc float32, f func(el interface{}, acc float32) (float32, error)) (float32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileFloat64(acc float64, f func(el interface{}, acc float64) (float64, error)) (float64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileInt(acc int, f func(el interface{}, acc int) (int, error)) (int, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileInt8(acc int8, f func(el interface{}, acc int8) (int8, error)) (int8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileInt16(acc int16, f func(el interface{}, acc int16) (int16, error)) (int16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileInt32(acc int32, f func(el interface{}, acc int32) (int32, error)) (int32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileInt64(acc int64, f func(el interface{}, acc int64) (int64, error)) (int64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileUint(acc uint, f func(el interface{}, acc uint) (uint, error)) (uint, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileUint8(acc uint8, f func(el interface{}, acc uint8) (uint8, error)) (uint8, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileUint16(acc uint16, f func(el interface{}, acc uint16) (uint16, error)) (uint16, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileUint32(acc uint32, f func(el interface{}, acc uint32) (uint32, error)) (uint32, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileUint64(acc uint64, f func(el interface{}, acc uint64) (uint64, error)) (uint64, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) ReduceWhileInterface(acc interface{}, f func(el interface{}, acc interface{}) (interface{}, error)) (interface{}, error) {
	for _, el := range s.data {
		acc, err := f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

func (s SliceInterface) Reverse() []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		result = append(result, s.data[i])
	}
	return result
}

func (s SliceInterface) Same() bool {
	for i := 0; i < len(s.data)-1; i++ {
		if s.data[i] != s.data[i+1] {
			return false
		}
	}
	return true
}

func (s SliceInterface) ScanBool(acc bool, f func(el interface{}, acc bool) bool) []bool {
	result := make([]bool, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanByte(acc byte, f func(el interface{}, acc byte) byte) []byte {
	result := make([]byte, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanString(acc string, f func(el interface{}, acc string) string) []string {
	result := make([]string, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanFloat32(acc float32, f func(el interface{}, acc float32) float32) []float32 {
	result := make([]float32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanFloat64(acc float64, f func(el interface{}, acc float64) float64) []float64 {
	result := make([]float64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanInt(acc int, f func(el interface{}, acc int) int) []int {
	result := make([]int, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanInt8(acc int8, f func(el interface{}, acc int8) int8) []int8 {
	result := make([]int8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanInt16(acc int16, f func(el interface{}, acc int16) int16) []int16 {
	result := make([]int16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanInt32(acc int32, f func(el interface{}, acc int32) int32) []int32 {
	result := make([]int32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanInt64(acc int64, f func(el interface{}, acc int64) int64) []int64 {
	result := make([]int64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanUint(acc uint, f func(el interface{}, acc uint) uint) []uint {
	result := make([]uint, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanUint8(acc uint8, f func(el interface{}, acc uint8) uint8) []uint8 {
	result := make([]uint8, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanUint16(acc uint16, f func(el interface{}, acc uint16) uint16) []uint16 {
	result := make([]uint16, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanUint32(acc uint32, f func(el interface{}, acc uint32) uint32) []uint32 {
	result := make([]uint32, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanUint64(acc uint64, f func(el interface{}, acc uint64) uint64) []uint64 {
	result := make([]uint64, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) ScanInterface(acc interface{}, f func(el interface{}, acc interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}

func (s SliceInterface) Shuffle() []interface{} {
	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	rand.Shuffle(len(s.data), swap)
	return s.data
}

func (s SliceInterface) Split(sep interface{}) [][]interface{} {
	result := make([][]interface{}, 0)
	curr := make([]interface{}, 0)
	for _, el := range s.data {
		if el == sep {
			result = append(result, curr)
		} else {
			curr = append(curr, el)
		}
	}
	result = append(result, curr)
	return result
}

func (s SliceInterface) StartsWith(prefix []interface{}) bool {
	for i, el := range prefix {
		if el != s.data[i] {
			return false
		}
	}
	return true
}

func (s SliceInterface) TakeWhile(f func(el interface{}) bool) []interface{} {
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		if !f(el) {
			return result
		}
		result = append(result, el)
	}
	return result
}

func (s SliceInterface) Uniq() []interface{} {
	added := make(map[interface{}]struct{})
	nothing := struct{}{}
	result := make([]interface{}, 0, len(s.data))
	for _, el := range s.data {
		_, exists := added[el]
		if !exists {
			result = append(result, el)
			added[el] = nothing
		}
	}
	return result

}

func (s SliceInterface) Window(size int) [][]interface{} {
	result := make([][]interface{}, 0, len(s.data)/size)
	for i := 0; i <= len(s.data)-size; i++ {
		chunk := s.data[i : i+size]
		result = append(result, chunk)
	}
	return result
}

func (s SlicesInterface) Concat() []interface{} {
	result := make([]interface{}, 0)
	for _, arr := range s.data {
		result = append(result, arr...)
	}
	return result
}

func (s SlicesInterface) Product() chan []interface{} {
	c := make(chan []interface{}, 1)
	go s.product(c, []interface{}{}, 0)
	return c
}

func (s SlicesInterface) product(c chan []interface{}, left []interface{}, pos int) {
	// iterate over the last array
	if pos == len(s.data)-1 {
		for _, el := range s.data[pos] {
			result := make([]interface{}, 0, len(left)+1)
			result = append(result, left...)
			result = append(result, el)
			c <- result
		}
		return
	}

	for _, el := range s.data[pos] {
		result := make([]interface{}, 0, len(left)+1)
		result = append(result, left...)
		result = append(result, el)
		s.product(c, result, pos+1)
	}

	if pos == 0 {
		close(c)
	}
}

func (s SlicesInterface) Zip() [][]interface{} {
	size := len(s.data[0])
	for _, arr := range s.data[1:] {
		if len(arr) > size {
			size = len(arr)
		}
	}

	result := make([][]interface{}, 0, size)
	for i := 0; i <= size; i++ {
		chunk := make([]interface{}, 0, len(s.data))
		for _, arr := range s.data {
			chunk = append(chunk, arr[i])
		}
		result = append(result, chunk)
	}
	return result
}
