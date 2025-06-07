package curry

// region BindOne

func BindOne[A, R any](f func(A) R, a A) func() R {
	return func() R {
		return f(a)
	}
}

func BindOne2[A, RA, RB any](f func(A) (RA, RB), a A) func() (RA, RB) {
	return func() (RA, RB) {
		return f(a)
	}
}

// region BindFirst

func BindFirstOfTwo[A, B, R any](f func(A, B) R, a A) func(B) R {
	return Two(f)(a)
}

func BindFirstOfThree[A, B, C, R any](f func(A, B, C) R, a A) func(B, C) R {
	return UnTwo(Three(f)(a))
}

func BindFirstOfFour[A, B, C, D, R any](f func(A, B, C, D) R, a A) func(B, C, D) R {
	return UnThree(Four(f)(a))
}

// region BindFirst with 2 RV

func BindFirstOfTwo2[A, B, RA, RB any](f func(A, B) (RA, RB), a A) func(B) (RA, RB) {
	return Two2(f)(a)
}

func BindFirstOfThree2[A, B, C, RA, RB any](f func(A, B, C) (RA, RB), a A) func(B, C) (RA, RB) {
	return UnTwo2(Three2(f)(a))
}

func BindFirstOfFour2[A, B, C, D, RA, RB any](f func(A, B, C, D) (RA, RB), a A) func(B, C, D) (RA, RB) {
	return UnThree2(Four2(f)(a))
}

// region BindLast

func BindLastOfTwo[A, B, R any](f func(A, B) R, b B) func(A) R {
	return func(a A) R {
		return f(a, b)
	}
}

func BindLastOfThree[A, B, C, R any](f func(A, B, C) R, c C) func(A, B) R {
	return func(a A, b B) R {
		return f(a, b, c)
	}
}

func BindLastOfFour[A, B, C, D, R any](f func(A, B, C, D) R, d D) func(A, B, C) R {
	return func(a A, b B, c C) R {
		return f(a, b, c, d)
	}
}

// region BindLast with 2 RV

func BindLastOfTwo2[A, B, RA, RB any](f func(A, B) (RA, RB), b B) func(A) (RA, RB) {
	return func(a A) (RA, RB) {
		return f(a, b)
	}
}

func BindLastOfThree2[A, B, C, RA, RB any](f func(A, B, C) (RA, RB), c C) func(A, B) (RA, RB) {
	return func(a A, b B) (RA, RB) {
		return f(a, b, c)
	}
}

func BindLastOfFour2[A, B, C, D, RA, RB any](f func(A, B, C, D) (RA, RB), d D) func(A, B, C) (RA, RB) {
	return func(a A, b B, c C) (RA, RB) {
		return f(a, b, c, d)
	}
}

// region BindFirst*Slice

func BindFirstOneSlice[A, B, RV any](f func(A, ...B) RV, a A) func(...B) RV {
	return func(b ...B) RV {
		return f(a, b...)
	}
}

func BindFirstOneSlice2[A, B, RA, RB any](f func(A, ...B) (RA, RB), a A) func(...B) (RA, RB) {
	return func(b ...B) (RA, RB) {
		return f(a, b...)
	}
}

func BindFirstTwoSlice[A, B, C, RV any](f func(A, B, ...C) RV, a A) func(B, ...C) RV {
	return func(b B, c ...C) RV {
		return f(a, b, c...)
	}
}

func BindFirstTwoSlice2[A, B, C, RA, RB any](f func(A, B, ...C) (RA, RB), a A) func(B, ...C) (RA, RB) {
	return func(b B, c ...C) (RA, RB) {
		return f(a, b, c...)
	}
}

// region BindLast*Slice

func BindLastOneSlice[A, B, RV any](f func(A, ...B) RV, b ...B) func(A) RV {
	return func(a A) RV {
		return f(a, b...)
	}
}

func BindLastOneSlice2[A, B, RA, RB any](f func(A, ...B) (RA, RB), b ...B) func(A) (RA, RB) {
	return func(a A) (RA, RB) {
		return f(a, b...)
	}
}

func BindLastTwoSlice[A, B, C, RV any](f func(A, B, ...C) RV, c ...C) func(A, B) RV {
	return func(a A, b B) RV {
		return f(a, b, c...)
	}
}

func BindLastTwoSlice2[A, B, C, RA, RB any](f func(A, B, ...C) (RA, RB), c ...C) func(A, B) (RA, RB) {
	return func(a A, b B) (RA, RB) {
		return f(a, b, c...)
	}
}
