package genesis

import "iter"

type CmpIter[T comparable] Iter[T]

func NewCmpIter[T comparable](s iter.Seq[T]) CmpIter[T] {
	return CmpIter[T](s)
}

func (it CmpIter[T]) Eq(other CmpIter[T]) bool {
	pull, stop := iter.Pull(iter.Seq[T](other))
	defer stop()
	for left := range it {
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

func (it CmpIter[T]) Contains(v T) bool {
	for left := range it {
		if left == v {
			return true
		}
	}
	return false
}
