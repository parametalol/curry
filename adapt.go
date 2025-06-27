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
func AdaptNone[RV any, Fn FuncNone[RV] | FuncNone[any]](f Fn) func() RV {
	switch t := any(f).(type) {
	case func() RV:
		return t
	case func() any:
		return func() RV {
			return t().(RV)
		}
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
//
//	 A , RV    0
//	 A , -     1
//	 A , any   2
//	 - , RV    3
//	 - , -     4
//	 - , any   5
//	any, RV    6
//	any, -     7
//	any, any   8
func AdaptOne[A, RV any, Fn FuncOne[A, RV] | FuncOne[any, RV] | FuncOne[A, any] | FuncOne[any, any]](f Fn) func(A) RV {
	switch t := any(f).(type) {
	case func(A) RV: // 0
		return t
	case func(A): // 1
		return func(a A) RV {
			t(a)
			var rv RV
			return rv
		}
	case func(A) any: // 2
		return func(a A) RV {
			return t(a).(RV)
		}
	case func() RV: // 3
		return func(_ A) RV {
			return t()
		}
	case func(): // 4
		return func(_ A) RV {
			t()
			var rv RV
			return rv
		}
	case func() any: // 5
		return func(_ A) RV {
			return t().(RV)
		}
	case func(any) RV: // 6
		return func(a A) RV {
			return t(a)
		}
	case func(any): // 7
		return func(a A) RV {
			t(a)
			var rv RV
			return rv
		}
	case func(any) any: // 8
		return func(a A) RV {
			return t(a).(RV)
		}
	}
	panic("unsupported function signature")
}

type funcTwoCombinations[A, B, RV any] interface {
	FuncTwo[A, B, RV] | FuncTwo[any, B, RV] | FuncTwo[A, any, RV] | FuncTwo[A, B, any] |
		FuncTwo[any, any, RV] | FuncTwo[A, any, any] | FuncTwo[any, B, any] | FuncTwo[any, any, any]
}

// AdaptTwo adapts the given function f to the `func(A, B) RV` signature.
//
//	 A ,  B , RV   0
//	 A ,  B , -    1
//	 A ,  B , any  2
//	 A ,  - , RV   3
//	 A ,  - , -    4
//	 A ,  - , any  5
//	 A , any, RV   6
//	 A , any, -    7
//	 A , any, any  8
//	 - ,  B , RV   9
//	 - ,  B , -    10
//	 - ,  B , any  11
//	 - ,  - , RV   12
//	 - ,  - , -    13
//	 - ,  - , any  14
//	 - , any, RV   15
//	 - , any, -    16
//	 - , any, any  17
//	any,  B , RV   18
//	any,  B , -    19
//	any,  B , any  20
//	any,  - , RV   see 15
//	any,  - , -    see 16
//	any,  - , any  see 17
//	any, any, RV   24
//	any, any, -    25
//	any, any, any  26
func AdaptTwo[A, B, RV any, Fn funcTwoCombinations[A, B, RV]](f Fn) func(A, B) RV {
	switch t := any(f).(type) {
	case func(A, B) RV: //0
		return t
	case func(A, B): //1
		return func(a A, b B) RV {
			t(a, b)
			var rv RV
			return rv
		}
	case func(A, B) any: // 2
		return func(a A, b B) RV {
			return t(a, b).(RV)
		}
	case func(A) RV: // 3
		return func(a A, _ B) RV {
			return t(a)
		}
	case func(A): // 4
		return func(a A, _ B) RV {
			t(a)
			var rv RV
			return rv
		}
	case func(A) any: // 5
		return func(a A, _ B) RV {
			return t(a).(RV)
		}
	case func(A, any) RV: // 6
		return func(a A, b B) RV {
			return t(a, b)
		}
	case func(A, any): // 7
		return func(a A, b B) RV {
			t(a, b)
			var rv RV
			return rv
		}
	case func(A, any) any: // 8
		return func(a A, b B) RV {
			return t(a, b).(RV)
		}
	case func(B) RV: // 9
		return func(_ A, b B) RV {
			return t(b)
		}
	case func(B): // 10
		return func(_ A, b B) RV {
			t(b)
			var rv RV
			return rv
		}
	case func(B) any: // 11
		return func(_ A, b B) RV {
			return t(b).(RV)
		}
	case func() RV: // 12
		return func(_ A, _ B) RV {
			return t()
		}
	case func(): // 13
		return func(_ A, _ B) RV {
			t()
			var rv RV
			return rv
		}
	case func() any: // 14
		return func(A, B) RV {
			return t().(RV)
		}
	case func(any) RV: // 15
		return func(a A, _ B) RV {
			return t(a)
		}
	case func(any): // 16
		return func(a A, _ B) RV {
			t(a)
			var rv RV
			return rv
		}
	case func(any) any: // 17
		return func(a A, _ B) RV {
			return t(a).(RV)
		}
	case func(any, B) RV: // 18
		return func(a A, b B) RV {
			return t(a, b)
		}
	case func(any, B): // 19
		return func(a A, b B) RV {
			t(a, b)
			var rv RV
			return rv
		}
	case func(any, B) any: // 20
		return func(a A, b B) RV {
			return t(a, b).(RV)
		}
	case func(any, any) RV: // 24
		return func(a A, b B) RV {
			return t(a, b)
		}
	case func(any, any): // 25
		return func(a A, b B) RV {
			t(a, b)
			var rv RV
			return rv
		}
	case func(any, any) any: // 26
		return func(a A, b B) RV {
			return t(a, b).(RV)
		}
	}
	panic("unsupported function signature")
}

// region AdaptF

// AdaptNoneF deduces the target function signature from the first
// unused argument.
func AdaptNoneF[RV any, Fn FuncNone[RV] | FuncNone[any]](_ func() RV, f Fn) func() RV {
	return AdaptNone[RV](f)
}

// AdaptOneF deduces the target function signature from the first
// unused argument.
func AdaptOneF[A, RV any, Fn FuncOne[A, RV] | FuncOne[any, RV] | FuncOne[A, any] | FuncOne[any, any]](_ func(A) RV, f Fn) func(A) RV {
	return AdaptOne[A, RV](f)
}

// AdaptTwoF deduces the target function signature from the first
// unused argument.
func AdaptTwoF[A, B, RV any, Fn funcTwoCombinations[A, B, RV]](_ func(A, B) RV, f Fn) func(A, B) RV {
	return AdaptTwo[A, B, RV](f)
}
