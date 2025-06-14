package curry

// generic type alias requires GOEXPERIMENT=aliastypeparams
// type normalizedFuncNone[RV any] = func() RV
// type normalizedFuncOne[A, RV any] = func(A) RV
// type normalizedFuncTwo[A, B, RV any] = func(A, B) RV

type FuncNone[RV any] interface {
	~func() RV | ~func()
}

type FuncOne[A, RV any] interface {
	~func(A) RV | ~func(A) |
		FuncNone[RV]
}

type FuncTwo[A, B, RV any] interface {
	~func(A, B) RV | ~func(A, B) |
		FuncOne[A, RV] | FuncOne[B, RV]
}

// AdaptNone adapts the given function f to the `func() RV` signature.
func AdaptNone[RV any, Fn FuncNone[RV]](f Fn) func() RV {
	switch t := any(f).(type) {
	case func() RV:
		return t
	case func():
		return func() RV {
			t()
			var rv RV
			return rv
		}
	}
	panic("unsupported function signature")
}

// AdaptOne adapts the given function f to the `func(A) RV` signature.
func AdaptOne[A, RV any, Fn FuncOne[A, RV]](f Fn) func(A) RV {
	switch t := any(f).(type) {
	case func(A) RV:
		return t
	case func():
		return func(_ A) RV {
			t()
			var rv RV
			return rv
		}
	case func() RV:
		return func(_ A) RV {
			return t()
		}
	case func(A):
		return func(a A) RV {
			t(a)
			var rv RV
			return rv
		}
	}
	panic("unsupported function signature")
}

// AdaptTwo adapts the given function f to the `func(A, B) RV` signature.
func AdaptTwo[A, B, RV any, Fn FuncTwo[A, B, RV]](f Fn) func(A, B) RV {
	switch t := any(f).(type) {
	case func(A, B) RV:
		return t
	case func():
		return func(_ A, _ B) RV {
			t()
			var rv RV
			return rv
		}
	case func() RV:
		return func(_ A, _ B) RV {
			return t()
		}
	case func(A):
		return func(a A, _ B) RV {
			t(a)
			var rv RV
			return rv
		}
	case func(A) RV:
		return func(a A, _ B) RV {
			return t(a)
		}
	case func(B):
		return func(_ A, b B) RV {
			t(b)
			var rv RV
			return rv
		}
	case func(B) RV:
		return func(_ A, b B) RV {
			return t(b)
		}
	case func(A, B):
		return func(a A, b B) RV {
			t(a, b)
			var rv RV
			return rv
		}
	}
	panic("unsupported function signature")
}

// region AdaptF

// AdaptNoneF deduces the target function signature from the first
// unused argument.
func AdaptNoneF[RV any, Fn FuncNone[RV]](_ func() RV, f Fn) func() RV {
	return AdaptNone[RV](f)
}

// AdaptOneF deduces the target function signature from the first
// unused argument.
func AdaptOneF[A, RV any, Fn FuncOne[A, RV]](_ func(A) RV, f Fn) func(A) RV {
	return AdaptOne[A, RV](f)
}

// AdaptTwoF deduces the target function signature from the first
// unused argument.
func AdaptTwoF[A, B, RV any, Fn FuncTwo[A, B, RV]](_ func(A, B) RV, f Fn) func(A, B) RV {
	return AdaptTwo[A, B, RV](f)
}
