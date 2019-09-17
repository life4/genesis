package implementation

// Identity accepts one argument and returns it
func Identity(t T) T { return t }

// Composition makes a new function that calls given functions one-by-one
// for given argument. EverySecond function accepts result of first function and so on.
func Composition(fs ...func(val T) T) func(val T) T {
	return func(val T) T {
		for _, f := range fs {
			val = f(val)
		}
		return val
	}
}
