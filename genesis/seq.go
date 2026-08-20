package genesis

import (
	"iter"
	"slices"
)

type Seq[T any] iter.Seq[T]

func NewSeq[T any](s iter.Seq[T]) Seq[T] {
	return Seq[T](s)
}

func (it Seq[T]) Slice() Slice[T] {
	return NewSlice(slices.Collect(iter.Seq[T](it)))
}

func (it Seq[T]) Count() int {
	cnt := 0
	for range it {
		cnt++
	}
	return cnt
}

func (it Seq[T]) Last() Option[T] {
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

func (it Seq[T]) Nth(n int) Option[T] {
	i := 0
	for v := range it {
		if i >= n {
			return Some(v)
		}
		i++
	}
	return None[T]()
}

func (it Seq[T]) StepBy(n int) Seq[T] {
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

func (it Seq[T]) Chain(other Seq[T]) Seq[T] {
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

func (it Seq[T]) Intersperse(other T) Seq[T] {
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

func (it Seq[T]) Map[R any](f func(T) R) Seq[R] {
	return func(yield func(R) bool) {
		for v := range it {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func (it Seq[T]) Each(f func(T)) {
	for v := range it {
		f(v)
	}
}

func (it Seq[T]) Filter(f func(T) bool) Seq[T] {
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

func (it Seq[T]) FilterMap[R any](f func(T) Option[R]) Seq[R] {
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

func (it Seq[T]) Skip(n int) Seq[T] {
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

func (it Seq[T]) Take(n int) Seq[T] {
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

func (it Seq[T]) FlatMap[R any](f func(T) Seq[R]) Seq[R] {
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

func (it Seq[T]) Fuse() Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if !yield(v) {
				break
			}
		}
	}
}

func (it Seq[T]) Inspect(f func(T)) Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			f(v)
			if !yield(v) {
				return
			}
		}
	}
}

func (it Seq[T]) Fold[R any](init R, f func(R, T) R) R {
	for v := range it {
		init = f(init, v)
	}
	return init
}

func (it Seq[T]) Reduce(f func(T, T) T) Option[T] {
	acc := None[T]()
	for v := range it {
		if !acc.ok {
			acc = Some(v)
		} else {
			acc = Some(f(acc.val, v))
		}
	}
	return acc
}

func (it Seq[T]) All(f func(T) bool) bool {
	for v := range it {
		if !f(v) {
			return false
		}
	}
	return true
}

func (it Seq[T]) Any(f func(T) bool) bool {
	for v := range it {
		if f(v) {
			return true
		}
	}
	return false
}

func (it Seq[T]) Find(f func(T) bool) Option[T] {
	for v := range it {
		if f(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (it Seq[T]) FindMap(f func(T) Option[T]) Option[T] {
	for v := range it {
		r := f(v)
		if r.ok {
			return r
		}
	}
	return None[T]()
}

func (it Seq[T]) Index(f func(T) bool) Option[int] {
	i := 0
	for v := range it {
		if f(v) {
			return Some(i)
		}
		i++
	}
	return None[int]()
}
