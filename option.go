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
	return Option[T]{def, false}
}

func (opt Option[T]) Tuple() (T, bool) {
	return opt.val, opt.ok
}

func (opt Option[T]) IsSome() bool {
	return opt.ok
}

func (opt Option[T]) IsNone() bool {
	return !opt.ok
}

func (opt Option[T]) IsSomeAnd(f func(T) bool) bool {
	return opt.ok && f(opt.val)
}

func (opt Option[T]) IsNoneOr(f func(T) bool) bool {
	return !opt.ok || f(opt.val)
}

func (opt Option[T]) Slice(f func(T) bool) []T {
	if opt.ok {
		return []T{opt.val}
	}
	return []T{}
}

func (opt Option[T]) Expect(msg string) T {
	if !opt.ok {
		panic(msg)
	}
	return opt.val
}

func (opt Option[T]) Unwrap() T {
	if !opt.ok {
		panic("expected Some, got None")
	}
	return opt.val
}

func (opt Option[T]) UnwrapOr(def T) T {
	if !opt.ok {
		return def
	}
	return opt.val
}

func (opt Option[T]) UnwrapOrElse(f func() T) T {
	if !opt.ok {
		return f()
	}
	return opt.val
}

func (opt Option[T]) UnwrapOrDefault() T {
	if !opt.ok {
		var def T
		return def
	}
	return opt.val
}

func (opt Option[T]) UnwrapUnchecked() T {
	return opt.val
}

func (opt Option[T]) Map[R any](f func(T) R) Option[R] {
	if opt.ok {
		return Some(f(opt.val))
	}
	return None[R]()
}

func (opt Option[T]) Inspect(f func(T)) Option[T] {
	if opt.ok {
		f(opt.val)
	}
	return opt
}

func (opt Option[T]) MapOr[R any](def R, f func(T) R) R {
	if opt.ok {
		return f(opt.val)
	}
	return def
}

func (opt Option[T]) MapOrElse[R any](def func() R, f func(T) R) R {
	if opt.ok {
		return f(opt.val)
	}
	return def()
}

func (opt Option[T]) MapOrDefault[R any](f func(T) R) R {
	if opt.ok {
		return f(opt.val)
	}
	var def R
	return def
}

func (opt Option[T]) And[R any](other Option[R]) Option[R] {
	if opt.ok {
		return other
	}
	return None[R]()
}
