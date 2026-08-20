package genesis

type Option[T any] struct {
	val T
	ok  bool
}

func Some[T any](val T) Option[T] {
	return Option[T]{val, true}
}

func None[T any]() Option[T] {
	var def T
	return Option[T]{def, true}
}

func (opt Option[T]) Tuple() (T, bool) {
	return opt.val, opt.ok
}

func (opt Option[T]) Must() T {
	if !opt.ok {
		panic("expected Some, got None")
	}
	return opt.val
}
