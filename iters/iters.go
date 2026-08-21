package iters

import (
	"cmp"
	"iter"
	"slices"
)

func FromSlice[S ~[]E, E any](s S) iter.Seq[E] {
	return slices.Values(s)
}

func FromChannel[T any](ch <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range ch {
			if !yield(v) {
				return
			}
		}
	}
}

func ToSlice[T any](it iter.Seq[T]) []T {
	return slices.Collect(it)
}

func Count[T any](it iter.Seq[T]) int {
	cnt := 0
	for range it {
		cnt++
	}
	return cnt
}

func First[T any](it iter.Seq[T]) (T, error) {
	for v := range it {
		return v, nil
	}
	var def T
	return def, ErrEmpty
}

func Last[T any](it iter.Seq[T]) (T, error) {
	var res T
	empty := true
	for v := range it {
		res = v
		empty = false
	}
	if empty {
		var def T
		return def, ErrEmpty
	}
	return res, nil
}

func Nth[T any](it iter.Seq[T], n int) (T, error) {
	i := 0
	for v := range it {
		if i >= n {
			return v, nil
		}
		i++
	}
	var def T
	return def, ErrNotFound
}

func StepBy[T any](it iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		for v := range it {
			if i == 0 {
				if !yield(v) {
					return
				}
			}
			i = (i + 1) % n
		}
	}
}

func Chain[T any](it iter.Seq[T], other iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if !yield(v) {
				return
			}
		}
		for v := range other {
			if !yield(v) {
				return
			}
		}
	}
}

func Intersperse[T any](it iter.Seq[T], other T) iter.Seq[T] {
	return func(yield func(T) bool) {
		first := true
		for v := range it {
			if !first {
				if !yield(other) {
					return
				}
			}
			first = false
			if !yield(v) {
				return
			}
		}
	}
}

func Map[T, R any](it iter.Seq[T], f func(T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range it {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func Each[T any](it iter.Seq[T], f func(T)) {
	for v := range it {
		f(v)
	}
}

func Filter[T any](it iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func FilterMap[T, R any](it iter.Seq[T], f func(T) (R, bool)) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range it {
			res, ok := f(v)
			if ok {
				if !yield(res) {
					return
				}
			}
		}
	}
}

func Skip[T any](it iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		if n == 0 {
			return
		}
		i := 0
		for v := range it {
			i++
			if i >= n {
				if !yield(v) {
					break
				}
			}
		}
	}
}

func Take[T any](it iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		if n == 0 {
			return
		}
		for v := range it {
			if !yield(v) {
				break
			}
			i++
			if i >= n {
				break
			}
		}
	}
}

func FlatMap[T, R any](it iter.Seq[T], f func(T) iter.Seq[R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range it {
			for r := range f(v) {
				if !yield(r) {
					return
				}
			}
		}
	}
}

func Fuse[T any](it iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if !yield(v) {
				break
			}
		}
	}
}

func Inspect[T any](it iter.Seq[T], f func(T)) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			f(v)
			if !yield(v) {
				return
			}
		}
	}
}

func Fold[T, R any](it iter.Seq[T], init R, f func(R, T) R) R {
	for v := range it {
		init = f(init, v)
	}
	return init
}

func Reduce[T any](it iter.Seq[T], f func(T, T) T) (T, error) {
	var acc T
	accSet := false
	for v := range it {
		if !accSet {
			acc = v
		} else {
			acc = f(acc, v)
		}
	}
	if !accSet {
		var def T
		return def, ErrEmpty
	}
	return acc, nil
}

func All[T any](it iter.Seq[T], f func(T) bool) bool {
	for v := range it {
		if !f(v) {
			return false
		}
	}
	return true
}

func Any[T any](it iter.Seq[T], f func(T) bool) bool {
	for v := range it {
		if f(v) {
			return true
		}
	}
	return false
}

func Find[T any](it iter.Seq[T], f func(T) bool) (T, error) {
	for v := range it {
		if f(v) {
			return v, nil
		}
	}
	var def T
	return def, ErrNotFound
}

func FindMap[T any](it iter.Seq[T], f func(T) (T, bool)) (T, error) {
	for v := range it {
		r, found := f(v)
		if found {
			return r, nil
		}
	}
	var def T
	return def, ErrNotFound
}

func Index[T any](it iter.Seq[T], f func(T) bool) (int, error) {
	i := 0
	for v := range it {
		if f(v) {
			return i, nil
		}
		i++
	}
	return 0, ErrOutOfRange
}

func Eq[T comparable](lhs, rhs iter.Seq[T]) bool {
	pull, stop := iter.Pull(iter.Seq[T](rhs))
	defer stop()
	for left := range lhs {
		right, more := pull()
		if !more {
			return false
		}
		if left != right {
			return false
		}
	}
	_, more := pull()
	return !more
}

func Contains[T comparable](it iter.Seq[T], v T) bool {
	for left := range it {
		if left == v {
			return true
		}
	}
	return false
}

func Max[T cmp.Ordered](it iter.Seq[T]) (T, error) {
	var max T
	first := true
	for v := range it {
		if first || v > max {
			max = v
		}
		first = false
	}
	if first {
		return max, ErrEmpty
	}
	return max, nil
}

func Min[T cmp.Ordered](it iter.Seq[T]) (T, error) {
	var min T
	first := true
	for v := range it {
		if first || v < min {
			min = v
		}
		first = false
	}
	if first {
		return min, ErrEmpty
	}
	return min, nil
}

func Sum[T cmp.Ordered](it iter.Seq[T]) T {
	var sum T
	for v := range it {
		sum += v
	}
	return sum
}

func Join[T cmp.Ordered](it iter.Seq[T], sep T) T {
	var sum T
	first := true
	for v := range it {
		if !first {
			sum += sep
		}
		sum += v
		first = false
	}
	return sum
}

func Sorted[T cmp.Ordered](it iter.Seq[T]) bool {
	var prev T
	first := true
	for v := range it {
		if !first && v < prev {
			return false
		}
		prev = v
		first = false
	}
	return true
}
