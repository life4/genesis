package genesis

import (
	"iter"
	"slices"
)

type Iter[T any] iter.Seq[T]

func NewIter[T any](s iter.Seq[T]) Iter[T] {
	return Iter[T](s)
}

func (it Iter[T]) Seq() iter.Seq[T] {
	return iter.Seq[T](it)
}

func (it Iter[T]) Slice() Slice[T] {
	return NewSlice(slices.Collect(it.Seq()))
}

func (it Iter[T]) Next() Option[T] {
	res := None[T]()
	it(func(v T) bool {
		res = Some(v)
		return true
	})
	return res
}

func (it Iter[T]) Count() int {
	cnt := 0
	for range it {
		cnt++
	}
	return cnt
}

func (it Iter[T]) Last() Option[T] {
	var res T
	empty := true
	for v := range it {
		res = v
		empty = false
	}
	if empty {
		return None[T]()
	}
	return Some(res)
}

func (it Iter[T]) AdvanceBy(n int) int {
	i := 0
	if n == 0 {
		return 0
	}
	for range it {
		i++
		if i >= n {
			break
		}
	}
	return i
}

func (it Iter[T]) Nth(n int) Option[T] {
	i := 0
	for v := range it {
		if i >= n {
			return Some(v)
		}
		i++
	}
	return None[T]()
}

func (it Iter[T]) StepBy(n int) Iter[T] {
	return func(yield func(T) bool) {
	loop:
		i := 0
		for v := range it {
			if i >= n {
				if !yield(v) {
					return
				}
				goto loop
			}
			i++
		}
	}
}

func (it Iter[T]) Chain(other Iter[T]) Iter[T] {
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

// func (it Iter[T]) Zip[R any](other Iter[R]) Iter[Pair[T, R]] {
// 	return func(yield func(Pair[T, R]) bool) {
// 		for left := range it {
// 			var right R
// 			other(func(v R) bool {
// 				right = v
// 				return true
// 			})
// 			cont := yield(Pair[T, R]{left, right})
// 			if !cont {
// 				break
// 			}
// 		}
// 	}
// }

func (it Iter[T]) Intersperse(other T) Iter[T] {
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

func (it Iter[T]) Map[R any](f func(T) R) Iter[R] {
	return func(yield func(R) bool) {
		for v := range it {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func (it Iter[T]) Each(f func(T)) {
	for v := range it {
		f(v)
	}
}

func (it Iter[T]) Filter(f func(T) bool) Iter[T] {
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

func (it Iter[T]) FilterMap[R any](f func(T) Option[R]) Iter[R] {
	return func(yield func(R) bool) {
		for v := range it {
			res := f(v)
			if res.ok {
				if !yield(res.val) {
					return
				}
			}
		}
	}
}

func (it Iter[T]) Skip(n int) Iter[T] {
	return func(yield func(T) bool) {
		i := 0
		if n == 0 {
			return
		}
		for range it {
			i++
			if i >= n {
				break
			}
		}
		for v := range it {
			if !yield(v) {
				break
			}
		}
	}
}

func (it Iter[T]) Take(n int) Iter[T] {
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

func (it Iter[T]) FlatMap[R any](f func(T) Iter[R]) Iter[R] {
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

func (it Iter[T]) Fuse() Iter[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if !yield(v) {
				break
			}
		}
	}
}

func (it Iter[T]) Inspect(f func(T)) Iter[T] {
	return func(yield func(T) bool) {
		for v := range it {
			f(v)
			if !yield(v) {
				return
			}
		}
	}
}

func (it Iter[T]) Fold[R any](init R, f func(R, T) R) R {
	for v := range it {
		init = f(init, v)
	}
	return init
}

func (it Iter[T]) Reduce(f func(T, T) T) Option[T] {
	opt := it.Next()
	if !opt.ok {
		return opt
	}
	init := opt.val
	for v := range it {
		init = f(init, v)
	}
	return Some(init)
}

func (it Iter[T]) All(f func(T) bool) bool {
	for v := range it {
		if !f(v) {
			return false
		}
	}
	return true
}

func (it Iter[T]) Any(f func(T) bool) bool {
	for v := range it {
		if f(v) {
			return true
		}
	}
	return false
}

func (it Iter[T]) Find(f func(T) bool) Option[T] {
	for v := range it {
		if f(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (it Iter[T]) FindMap(f func(T) Option[T]) Option[T] {
	for v := range it {
		r := f(v)
		if r.ok {
			return r
		}
	}
	return None[T]()
}

func (it Iter[T]) Index(f func(T) bool) Option[int] {
	i := 0
	for v := range it {
		if f(v) {
			return Some(i)
		}
		i++
	}
	return None[int]()
}
