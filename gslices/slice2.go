package gslices

import (
	"constraints"
)

// ChunkBy splits arr on every element for which f returns a new value.
func ChunkBy[S ~[]T, T comparable, G comparable](items S, f func(el T) G) []S {
	chunks := make([]S, 0)
	if len(items) == 0 {
		return chunks
	}

	chunk := make([]T, 0)
	prev := f(items[0])
	chunk = append(chunk, items[0])

	for _, el := range items[1:] {
		curr := f(el)
		if curr != prev {
			chunks = append(chunks, chunk)
			chunk = make([]T, 0)
			prev = curr
		}
		chunk = append(chunk, el)
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

// DedupBy returns a given slice without consecutive elements
// For which f returns the same result
func DedupBy[S ~[]T, T comparable, G comparable](items S, f func(el T) G) S {
	result := make(S, 0, len(items))
	if len(items) == 0 {
		return result
	}

	prev := f(items[0])
	result = append(result, items[0])
	for _, el := range items[1:] {
		curr := f(el)
		if curr != prev {
			result = append(result, el)
			prev = curr
		}
	}
	return result
}

// GroupBy groups element from array by value returned by f
func GroupBy[S ~[]T, T any, G constraints.Ordered](items S, f func(el T) G) map[G]S {
	result := make(map[G]S)
	for _, el := range items {
		key := f(el)
		val, ok := result[key]
		if !ok {
			result[key] = make(S, 1)
		}
		result[key] = append(val, el)
	}
	return result
}

// Map applies F to all elements in slice of T and returns slice of results
func Map[S ~[]T, T any, G any](items S, f func(el T) G) []G {
	result := make([]G, 0, len(items))
	for _, el := range items {
		result = append(result, f(el))
	}
	return result
}

// Reduce applies F to acc and every element in slice of T and returns acc
func Reduce[S ~[]T, T any, G any](items S, acc G, f func(el T, acc G) G) G {
	for _, el := range items {
		acc = f(el, acc)
	}
	return acc
}

// ReduceWhile is like Reduce, but stops when f returns error
func ReduceWhile[S ~[]T, T any, G any](items S, acc G, f func(el T, acc G) (G, error)) (G, error) {
	var err error
	for _, el := range items {
		acc, err = f(el, acc)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}

// Scan is like Reduce, but returns slice of f results
func Scan[S ~[]T, T any, G any](items S, acc G, f func(el T, acc G) G) []G {
	result := make([]G, 0, len(items))
	for _, el := range items {
		acc = f(el, acc)
		result = append(result, acc)
	}
	return result
}
