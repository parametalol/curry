package curry

// region Bind1st

func Bind1st2[A any, B any, R any](f func(A, B) R, a A) func(B) R {
	return Curry2(f)(a)
}

func Bind1st3[A any, B any, C any, R any](f func(A, B, C) R, a A) func(B, C) R {
	return UnCurry2(Curry3(f)(a))
}

func Bind1st4[A any, B any, C any, D any, R any](f func(A, B, C, D) R, a A) func(B, C, D) R {
	return UnCurry3(Curry4(f)(a))
}

// region Bind1st with error

func Bind1st2e[A any, B any, R any](f func(A, B) (R, error), a A) func(B) (R, error) {
	return Curry2e(f)(a)
}

func Bind1st3e[A any, B any, C any, R any](f func(A, B, C) (R, error), a A) func(B, C) (R, error) {
	return UnCurry2e(Curry3e(f)(a))
}

func Bind1st4e[A any, B any, C any, D any, R any](f func(A, B, C, D) (R, error), a A) func(B, C, D) (R, error) {
	return UnCurry3e(Curry4e(f)(a))
}

// region BindLast

func BindLast2[A any, B any, R any](f func(A, B) R, b B) func(A) R {
	return func(a A) R {
		return f(a, b)
	}
}

func BindLast3[A any, B any, C any, R any](f func(A, B, C) R, c C) func(A, B) R {
	return func(a A, b B) R {
		return f(a, b, c)
	}
}

func BindLast4[A any, B any, C any, D any, R any](f func(A, B, C, D) R, d D) func(A, B, C) R {
	return func(a A, b B, c C) R {
		return f(a, b, c, d)
	}
}

// region Bind1st with error

func BindLast2e[A any, B any, R any](f func(A, B) (R, error), b B) func(A) (R, error) {
	return func(a A) (R, error) {
		return f(a, b)
	}
}

func BindLast3e[A any, B any, C any, R any](f func(A, B, C) (R, error), c C) func(A, B) (R, error) {
	return func(a A, b B) (R, error) {
		return f(a, b, c)
	}
}

func BindLast4e[A any, B any, C any, D any, R any](f func(A, B, C, D) (R, error), d D) func(A, B, C) (R, error) {
	return func(a A, b B, c C) (R, error) {
		return f(a, b, c, d)
	}
}
