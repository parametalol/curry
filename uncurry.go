package curry

// region UnCurry

// UnTwo transforms a curried function with two arguments back into its original form.
func UnTwo[A any, B any, R any](f func(A) func(B) R) func(A, B) R {
	return func(a A, b B) R {
		return f(a)(b)
	}
}

// UnThree transforms a curried function with three arguments back into its original form.
func UnThree[A any, B any, C any, R any](f func(A) func(B) func(C) R) func(A, B, C) R {
	return func(a A, b B, c C) R {
		return f(a)(b)(c)
	}
}

// UnFour transforms a curried function with four arguments back into its original form.
func UnFour[A any, B any, C any, D any, R any](f func(A) func(B) func(C) func(D) R) func(A, B, C, D) R {
	return func(a A, b B, c C, d D) R {
		return f(a)(b)(c)(d)
	}
}

// region UnCurry with 2 RV

// UnTwo2 transforms a curried function with two arguments and two return values
// back into its original form.
func UnTwo2[A any, B any, RA any, RB any](f func(A) func(B) (RA, RB)) func(A, B) (RA, RB) {
	return func(a A, b B) (RA, RB) {
		return f(a)(b)
	}
}

// UnThree2 transforms a curried function with three arguments and two return values
// back into its original form.
func UnThree2[A any, B any, C any, RA any, RB any](f func(A) func(B) func(C) (RA, RB)) func(A, B, C) (RA, RB) {
	return func(a A, b B, c C) (RA, RB) {
		return f(a)(b)(c)
	}
}

// UnFour2 transforms a curried function with four arguments back into its original form.
func UnFour2[A any, B any, C any, D any, RA any, RB any](f func(A) func(B) func(C) func(D) (RA, RB)) func(A, B, C, D) (RA, RB) {
	return func(a A, b B, c C, d D) (RA, RB) {
		return f(a)(b)(c)(d)
	}
}
