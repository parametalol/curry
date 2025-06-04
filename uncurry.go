package curry

// region UnCurry

// UnCurry2 transforms a curried function with two arguments back into its original form.
func UnCurry2[A any, B any, R any](f func(A) func(B) R) func(A, B) R {
	return func(a A, b B) R {
		return f(a)(b)
	}
}

// UnCurry3 transforms a curried function with three arguments back into its original form.
func UnCurry3[A any, B any, C any, R any](f func(A) func(B) func(C) R) func(A, B, C) R {
	return func(a A, b B, c C) R {
		return f(a)(b)(c)
	}
}

// UnCurry4 transforms a curried function with four arguments back into its original form.
func UnCurry4[A any, B any, C any, D any, R any](f func(A) func(B) func(C) func(D) R) func(A, B, C, D) R {
	return func(a A, b B, c C, d D) R {
		return f(a)(b)(c)(d)
	}
}

// region UnCurry with error

// UnCurry2e transforms a curried function with two arguments back into its original form.
func UnCurry2e[A any, B any, R any](f func(A) func(B) (R, error)) func(A, B) (R, error) {
	return func(a A, b B) (R, error) {
		return f(a)(b)
	}
}

// UnCurry3e transforms a curried function with three arguments back into its original form.
func UnCurry3e[A any, B any, C any, R any](f func(A) func(B) func(C) (R, error)) func(A, B, C) (R, error) {
	return func(a A, b B, c C) (R, error) {
		return f(a)(b)(c)
	}
}

// UnCurry4e transforms a curried function with four arguments back into its original form.
func UnCurry4e[A any, B any, C any, D any, R any](f func(A) func(B) func(C) func(D) (R, error)) func(A, B, C, D) (R, error) {
	return func(a A, b B, c C, d D) (R, error) {
		return f(a)(b)(c)(d)
	}
}
