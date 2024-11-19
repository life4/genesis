package lambdas

// DefaultTo wraps a function invocation and returns the specified default value
// if it returned an error.
func DefaultTo[T any](def T) func(val T, err error) T {
	return func(val T, err error) T {
		if err != nil {
			return def
		}
		return val
	}
}

// Ensure wraps a function invocation and panic if it returned an error.
// Compared to Must, in Ensure the called function only returns an error.
func Ensure(err error) {
	if err != nil {
		panic(err)
	}
}

// Must wraps a function invocation and panic if it returned an error.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// Safe wraps a function invocation and returns the empty value if it returned an error.
func Safe[T any](val T, err error) T {
	if err != nil {
		var res T
		return res
	}
	return val
}
