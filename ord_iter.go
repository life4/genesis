package genesis

import (
	"iter"

	"github.com/life4/genesis/constraints"
)

type OrdIter[T constraints.Ordered] CmpIter[T]

func NewOrdIter[T constraints.Ordered](s iter.Seq[T]) OrdIter[T] {
	return OrdIter[T](s)
}

func (it OrdIter[T]) Max() Option[T] {
	max := None[T]()
	for v := range it {
		if !max.ok || v > max.val {
			max = Some(v)
		}
	}
	return max
}

func (it OrdIter[T]) Min() Option[T] {
	min := None[T]()
	for v := range it {
		if !min.ok || v < min.val {
			min = Some(v)
		}
	}
	return min
}

func (it OrdIter[T]) Sum() T {
	var sum T
	for v := range it {
		sum += v
	}
	return sum
}

func (it OrdIter[T]) Join(sep T) T {
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

func (it OrdIter[T]) Sorted() bool {
	prev := None[T]()
	for v := range it {
		if prev.ok && v < prev.val {
			return false
		}
		prev = Some(v)
	}
	return true
}
