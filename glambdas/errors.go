package glambdas

// Must wraps a function invicotaion and panic if it returned an error.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// Safe wraps a function invicotaion and returns the empty value if it returned an error.
func Safe[T any](val T, err error) T {
	if err != nil {
		var res T
		return res
	}
	return val
}

// DefaultTo wraps a function invicotaion and returns the specified default value
// if it returned an error.
func DefaultTo[T any](def T) func(val T, err error) T {
	return func(val T, err error) T {
		if err != nil {
			return def
		}
		return val
	}
}
