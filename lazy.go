package curry

// region Pass

func Pass[A any](a A) A                 { return a }
func PassTwo[A, B any](a A, b B) (A, B) { return a, b }

// region Return

func Return[A any](a A) func() A {
	return func() A { return a }
}

func Return2[A, B any](a A, b B) func() (A, B) {
	return func() (A, B) { return a, b }
}

func ReturnSlice[A any](a ...A) func() []A {
	return func() []A { return a }
}

// region Lazy0

// LazyOne converts the only function parameter to a thunk function.
func LazyOne0[A any](f func(A)) func(func() A) {
	return func(arg func() A) {
		f(arg())
	}
}

// LazyTwo converts function parameters to thunk functions.
func LazyTwo0[A, B any](f func(A, B)) func(func() A, func() B) {
	return func(arg1 func() A, arg2 func() B) {
		f(arg1(), arg2())
	}
}

// LazyThree converts function parameters to thunk functions.
func LazyThree0[A, B, C any](f func(A, B, C)) func(func() A, func() B, func() C) {
	return func(arg1 func() A, arg2 func() B, arg3 func() C) {
		f(arg1(), arg2(), arg3())
	}
}

// LazyFour converts function parameters to thunk functions.
func LazyFour0[A, B, C, D any](f func(A, B, C, D)) func(func() A, func() B, func() C, func() D) {
	return func(arg1 func() A, arg2 func() B, arg3 func() C, arg4 func() D) {
		f(arg1(), arg2(), arg3(), arg4())
	}
}

// LazyAll converts the only function slice parameter to a thunk function.
func LazyAll0[A any](f func(...A)) func(func() []A) {
	return func(arg func() []A) {
		f(arg()...)
	}
}

// region Lazy

// LazyOne converts the only function parameter to a thunk function.
func LazyOne[A, RV any](f func(A) RV) func(func() A) RV {
	return func(arg func() A) RV {
		return f(arg())
	}
}

// LazyTwo converts function parameters to thunk functions.
func LazyTwo[A, B, RV any](f func(A, B) RV) func(func() A, func() B) RV {
	return func(arg1 func() A, arg2 func() B) RV {
		return f(arg1(), arg2())
	}
}

// LazyThree converts function parameters to thunk functions.
func LazyThree[A, B, C, RV any](f func(A, B, C) RV) func(func() A, func() B, func() C) RV {
	return func(arg1 func() A, arg2 func() B, arg3 func() C) RV {
		return f(arg1(), arg2(), arg3())
	}
}

// LazyFour converts function parameters to thunk functions.
func LazyFour[A, B, C, D, RV any](f func(A, B, C, D) RV) func(func() A, func() B, func() C, func() D) RV {
	return func(arg1 func() A, arg2 func() B, arg3 func() C, arg4 func() D) RV {
		return f(arg1(), arg2(), arg3(), arg4())
	}
}

// LazyAll converts the only function slice parameter to a thunk function.
func LazyAll[A, RV any](f func(...A) RV) func(func() []A) RV {
	return func(arg func() []A) RV {
		return f(arg()...)
	}
}

// region Lazy2

// LazyOne2 converts the only function parameter to a thunk function.
func LazyOne2[A, RA, RB any](f func(A) (RA, RB)) func(func() A) (RA, RB) {
	return func(arg func() A) (RA, RB) {
		return f(arg())
	}
}

// LazyTwo2 converts function parameters to thunk functions.
func LazyTwo2[A, B, RA, RB any](f func(A, B) (RA, RB)) func(func() A, func() B) (RA, RB) {
	return func(arg1 func() A, arg2 func() B) (RA, RB) {
		return f(arg1(), arg2())
	}
}

// LazyThree2 converts function parameters to thunk functions.
func LazyThree2[A, B, C, RA, RB any](f func(A, B, C) (RA, RB)) func(func() A, func() B, func() C) (RA, RB) {
	return func(arg1 func() A, arg2 func() B, arg3 func() C) (RA, RB) {
		return f(arg1(), arg2(), arg3())
	}
}

// LazyFour2 converts function parameters to thunk functions.
func LazyFour2[A, B, C, D, RA, RB any](f func(A, B, C, D) (RA, RB)) func(func() A, func() B, func() C, func() D) (RA, RB) {
	return func(arg1 func() A, arg2 func() B, arg3 func() C, arg4 func() D) (RA, RB) {
		return f(arg1(), arg2(), arg3(), arg4())
	}
}

// LazyAll2 converts the only function slice parameter to a thunk function.
func LazyAll2[A, RA, RB any](f func(...A) (RA, RB)) func(func() []A) (RA, RB) {
	return func(arg func() []A) (RA, RB) {
		return f(arg()...)
	}
}
