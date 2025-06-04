// curry is a generic function currying framework in Go.
// It takes a function with multiple arguments and returns a sequence of functions,
// each taking a single argument.
package curry

// region Curry

// Curry2 curries a function with two arguments.
func Curry2[A any, B any, R any](f func(A, B) R) func(A) func(B) R {
	return func(a A) func(B) R {
		return func(b B) R {
			return f(a, b)
		}
	}
}

// Curry3 curries a function with three arguments.
func Curry3[A any, B any, C any, R any](f func(A, B, C) R) func(A) func(B) func(C) R {
	return func(a A) func(B) func(C) R {
		return func(b B) func(C) R {
			return func(c C) R {
				return f(a, b, c)
			}
		}
	}
}

// Curry4 curries a function with four arguments.
func Curry4[A any, B any, C any, D any, R any](f func(A, B, C, D) R) func(A) func(B) func(C) func(D) R {
	return func(a A) func(B) func(C) func(D) R {
		return func(b B) func(C) func(D) R {
			return func(c C) func(D) R {
				return func(d D) R {
					return f(a, b, c, d)
				}
			}
		}
	}
}

// region Curry with error

// Curry2e curries a function with two arguments.
func Curry2e[A any, B any, R any](f func(A, B) (R, error)) func(A) func(B) (R, error) {
	return func(a A) func(B) (R, error) {
		return func(b B) (R, error) {
			return f(a, b)
		}
	}
}

// Curry3e curries a function with three arguments.
func Curry3e[A any, B any, C any, R any](f func(A, B, C) (R, error)) func(A) func(B) func(C) (R, error) {
	return func(a A) func(B) func(C) (R, error) {
		return func(b B) func(C) (R, error) {
			return func(c C) (R, error) {
				return f(a, b, c)
			}
		}
	}
}

// Curry4e curries a function with four arguments.
func Curry4e[A any, B any, C any, D any, R any](f func(A, B, C, D) (R, error)) func(A) func(B) func(C) func(D) (R, error) {
	return func(a A) func(B) func(C) func(D) (R, error) {
		return func(b B) func(C) func(D) (R, error) {
			return func(c C) func(D) (R, error) {
				return func(d D) (R, error) {
					return f(a, b, c, d)
				}
			}
		}
	}
}
