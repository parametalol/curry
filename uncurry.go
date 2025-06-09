package curry

// region UnCurry

// UnTwo transforms a curried function with two arguments back into its original form.
func UnTwo0[A, B any](f func(A) func(B)) func(A, B) {
	return func(a A, b B) {
		f(a)(b)
	}
}

// UnThree transforms a curried function with three arguments back into its original form.
func UnThree0[A, B, C any](f func(A) func(B) func(C)) func(A, B, C) {
	return func(a A, b B, c C) {
		f(a)(b)(c)
	}
}

// UnFour transforms a curried function with four arguments back into its original form.
func UnFour0[A, B, C, D any](f func(A) func(B) func(C) func(D)) func(A, B, C, D) {
	return func(a A, b B, c C, d D) {
		f(a)(b)(c)(d)
	}
}

// UnTwo transforms a curried function with two arguments back into its original form.
func UnTwo[A, B, R any](f func(A) func(B) R) func(A, B) R {
	return func(a A, b B) R {
		return f(a)(b)
	}
}

// UnThree transforms a curried function with three arguments back into its original form.
func UnThree[A, B, C, R any](f func(A) func(B) func(C) R) func(A, B, C) R {
	return func(a A, b B, c C) R {
		return f(a)(b)(c)
	}
}

// UnFour transforms a curried function with four arguments back into its original form.
func UnFour[A, B, C, D, R any](f func(A) func(B) func(C) func(D) R) func(A, B, C, D) R {
	return func(a A, b B, c C, d D) R {
		return f(a)(b)(c)(d)
	}
}

// region UnCurry with 2 RV

// UnTwo2 transforms a curried function with two arguments and two return values
// back into its original form.
func UnTwo2[A, B, RA, RB any](f func(A) func(B) (RA, RB)) func(A, B) (RA, RB) {
	return func(a A, b B) (RA, RB) {
		return f(a)(b)
	}
}

// UnThree2 transforms a curried function with three arguments and two return values
// back into its original form.
func UnThree2[A, B, C, RA, RB any](f func(A) func(B) func(C) (RA, RB)) func(A, B, C) (RA, RB) {
	return func(a A, b B, c C) (RA, RB) {
		return f(a)(b)(c)
	}
}

// UnFour2 transforms a curried function with four arguments back into its original form.
func UnFour2[A, B, C, D, RA, RB any](f func(A) func(B) func(C) func(D) (RA, RB)) func(A, B, C, D) (RA, RB) {
	return func(a A, b B, c C, d D) (RA, RB) {
		return f(a)(b)(c)(d)
	}
}

// region UnTwoSlice

func UnTwoSlice[A, B, R any](f func(A) func(...B) R) func(A, ...B) R {
	return func(a A, b ...B) R {
		return f(a)(b...)
	}
}

func UnThreeSlice[A, B, C, R any](f func(A) func(B) func(...C) R) func(A, B, ...C) R {
	return func(a A, b B, c ...C) R {
		return f(a)(b)(c...)
	}
}

func UnTwoSlice2[A, B, RA, RB any](f func(A) func(...B) (RA, RB)) func(A, ...B) (RA, RB) {
	return func(a A, b ...B) (RA, RB) {
		return f(a)(b...)
	}
}

func UnThreeSlice2[A, B, C, RA, RB any](f func(A) func(B) func(...C) (RA, RB)) func(A, B, ...C) (RA, RB) {
	return func(a A, b B, c ...C) (RA, RB) {
		return f(a)(b)(c...)
	}
}
