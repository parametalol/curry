package curry

// generic type alias requires GOEXPERIMENT=aliastypeparams
// type normalizedFuncR[R any] = func() R
// type normalizedFunc1R[T0, R any] = func(T0) R
// type normalizedFunc2R[T0, T1, R any] = func(T0, T1) R

type Func interface {
	~func()
}

type FuncR[R any] interface {
	~func() R | Func
}

type Func1R[T0, R any] interface {
	~func(T0) R | ~func(T0) |
		FuncR[R]
}

type Func1[T0 any] interface {
	~func(T0) | Func
}

type Func2[T0, T1 any] interface {
	~func(T0, T1) | Func1[T0] | Func1[T1]
}

type Func2R[T0, T1, R any] interface {
	~func(T0, T1) R | Func2[T0, T1] |
		Func1R[T0, R] | Func1R[T1, R]
}

// region Any

type FuncAnyR[R any] interface {
	FuncR[R] | FuncR[any]
}

type FuncAny1[T0 any] interface {
	Func1[T0] | Func1[any]
}

type Func1AnyR[T0, R any] interface {
	Func1R[T0, R] | Func1R[T0, any]
}

type FuncAny1AnyR[T0, R any] interface {
	Func1AnyR[T0, R] | Func1AnyR[any, R]
}

type FuncAny2[T0, T1 any] interface {
	Func2[T0, T1] | Func2[any, T1] | Func2[T0, any] | Func2[any, any]
}

type Func2AnyR[T0, T1, R any] interface {
	Func2R[T0, T1, R] | Func2R[T0, T1, any]
}

type FuncAny2AnyR[T0, T1, R any] interface {
	Func2AnyR[T0, T1, R] | Func2AnyR[any, T1, R] | Func2AnyR[T0, any, R] | Func2AnyR[any, any, R] |
		FuncAny2[T0, T1]
}

// AdaptR upgrades the fn function signature to `func() R`.
//
//	 R  0
//	 -  1
//	any 2
func AdaptR[R any, Fn FuncAnyR[R]](fn Fn) func() R {
	switch t := any(fn).(type) {
	case /* 0 */ func() R:
		return t
	case /* 1 */ func():
		return func() (r R) { t(); return }
	case /* 2 */ func() any:
		return func() R { return t().(R) }
	}
	panic("unsupported function signature")
}

// Adapt1R upgrades the fn function signature to `func(T0) R`.
//
//	 T0,  R    0
//	 T0,  -    1
//	 T0, any   2
//	 - ,  R    3
//	 - ,  -    4
//	 - , any   5
//	any,  R    6
//	any,  -    7
//	any, any   8
func Adapt1R[T0, R any, Fn FuncAny1AnyR[T0, R]](fn Fn) func(T0) R {
	switch t := any(fn).(type) {
	case /* 0 */ func(T0) R:
		return t
	case /* 1 */ func(T0):
		return func(arg0 T0) (_ R) { t(arg0); return }
	case /* 2 */ func(T0) any:
		return func(arg0 T0) R { return t(arg0).(R) }
	case /* 3 */ func() R:
		return func(T0) R { return t() }
	case /* 4 */ func():
		return func(T0) (_ R) { t(); return }
	case /* 5 */ func() any:
		return func(T0) R { return t().(R) }
	case /* 6 */ func(any) R:
		return func(arg0 T0) R { return t(arg0) }
	case /* 7 */ func(any):
		return func(arg0 T0) (_ R) { t(arg0); return }
	case /* 8 */ func(any) any:
		return func(arg0 T0) R { return t(arg0).(R) }
	}
	panic("unsupported function signature")
}

// Adapt1 upgrades the fn function signature to `func(T0)`.
//
//	 T0  0
//	 -   1
//	any  2
func Adapt1[T0 any, Fn FuncAny1[T0]](fn Fn) func(T0) {
	switch t := any(fn).(type) {
	case /* 0 */ func(T0):
		return t
	case /* 1 */ func():
		return func(T0) { t() }
	case /* 2 */ func(any):
		return func(arg0 T0) { t(arg0) }
	}
	panic("unsupported function signature")
}

// Adapt2R upgrades the fn function signature to `func(T0, T1) R`.
//
//	 T0,  T1,  R   0
//	 T0,  T1,  -   1
//	 T0,  T1, any  2
//	 T0,  - ,  R   3
//	 T0,  - ,  -   4
//	 T0,  - , any  5
//	 T0, any,  R   6
//	 T0, any,  -   7
//	 T0, any, any  8
//	 - ,  T1,  R   9
//	 - ,  T1,  -   10
//	 - ,  T1, any  11
//	 - ,  - ,  R   12
//	 - ,  - ,  -   13
//	 - ,  - , any  14
//	 - , any,  R   15
//	 - , any,  -   16
//	 - , any, any  17
//	any,  T1,  R   18
//	any,  T1,  -   19
//	any,  T1, any  20
//	any,  - ,  R   see 15
//	any,  - ,  -   see 16
//	any,  - , any  see 17
//	any, any,  R   24
//	any, any,  -   25
//	any, any, any  26
func Adapt2R[T0, T1, R any, Fn FuncAny2AnyR[T0, T1, R]](fn Fn) func(T0, T1) R {
	switch t := any(fn).(type) {
	case /* 0 */ func(T0, T1) R:
		return t
	case /* 1 */ func(T0, T1):
		return func(arg0 T0, arg1 T1) (_ R) { t(arg0, arg1); return }
	case /* 2 */ func(T0, T1) any:
		return func(arg0 T0, arg1 T1) R { return t(arg0, arg1).(R) }
	case /* 3 */ func(T0) R:
		return func(arg0 T0, _ T1) R { return t(arg0) }
	case /* 4 */ func(T0):
		return func(arg0 T0, _ T1) (_ R) { t(arg0); return }
	case /* 5 */ func(T0) any:
		return func(arg0 T0, _ T1) R { return t(arg0).(R) }
	case /* 6 */ func(T0, any) R:
		return func(arg0 T0, arg1 T1) R { return t(arg0, arg1) }
	case /* 7 */ func(T0, any):
		return func(arg0 T0, arg1 T1) (_ R) { t(arg0, arg1); return }
	case /* 8 */ func(T0, any) any:
		return func(arg0 T0, arg1 T1) R { return t(arg0, arg1).(R) }
	case /* 9 */ func(T1) R:
		return func(_ T0, arg1 T1) R { return t(arg1) }
	case /* 10 */ func(T1):
		return func(_ T0, arg1 T1) (_ R) { t(arg1); return }
	case /* 11 */ func(T1) any:
		return func(_ T0, arg1 T1) R { return t(arg1).(R) }
	case /* 12 */ func() R:
		return func(_ T0, _ T1) R { return t() }
	case /* 13 */ func():
		return func(_ T0, _ T1) (_ R) { t(); return }
	case /* 14 */ func() any:
		return func(T0, T1) R { return t().(R) }
	case /* 15, 21 */ func(any) R:
		return func(arg0 T0, _ T1) R { return t(arg0) }
	case /* 16, 22 */ func(any):
		return func(arg0 T0, _ T1) (_ R) { t(arg0); return }
	case /* 17, 23 */ func(any) any:
		return func(arg0 T0, _ T1) R { return t(arg0).(R) }
	case /* 18 */ func(any, T1) R:
		return func(arg0 T0, arg1 T1) R { return t(arg0, arg1) }
	case /* 19 */ func(any, T1):
		return func(arg0 T0, arg1 T1) (_ R) { t(arg0, arg1); return }
	case /* 20 */ func(any, T1) any:
		return func(arg0 T0, arg1 T1) R { return t(arg0, arg1).(R) }
	case /* 24 */ func(any, any) R:
		return func(arg0 T0, arg1 T1) R { return t(arg0, arg1) }
	case /* 25 */ func(any, any):
		return func(arg0 T0, arg1 T1) (_ R) { t(arg0, arg1); return }
	case /* 26 */ func(any, any) any:
		return func(arg0 T0, arg1 T1) R { return t(arg0, arg1).(R) }
	}
	panic("unsupported function signature")
}

// Adapt2 upgrades the fn function signature to `func(T0, T1)`.
//
//	 T0,  T1  0
//	 T0,  -   1
//	 T0, any  2
//	 - ,  T1  3
//	 - ,  -   4
//	any,  T1  5
//	any,  -   6
//	any, any  7
func Adapt2[T0, T1 any, Fn FuncAny2[T0, T1]](fn Fn) func(T0, T1) {
	switch t := any(fn).(type) {
	case /* 0 */ func(T0, T1):
		return t
	case /* 1 */ func(T0):
		return func(arg0 T0, _ T1) { t(arg0) }
	case /* 2 */ func(T0, any):
		return func(arg0 T0, arg1 T1) { t(arg0, arg1) }
	case /* 3 */ func(T1):
		return func(_ T0, arg1 T1) { t(arg1) }
	case /* 4 */ func():
		return func(_ T0, _ T1) { t() }
	case /* 5 */ func(any, T1):
		return func(arg0 T0, arg1 T1) { t(arg0, arg1) }
	case /* 6 */ func(any):
		return func(arg0 T0, _ T1) { t(arg0) }
	case /* 7 */ func(any, any):
		return func(arg0 T0, arg1 T1) { t(arg0, arg1) }
	}
	panic("unsupported function signature")
}

// region AdaptF

// AdaptRF upgrades the fn function signature to the signature of the first
// argument.
func AdaptRF[R any, Fn FuncAnyR[R]](_ func() R, fn Fn) func() R {
	return AdaptR[R](fn)
}

// Adapt1F upgrades the fn function signature to the signature of the first
// argument.
func Adapt1F[T0 any, Fn FuncAny1[T0]](_ func(T0), fn Fn) func(T0) {
	return Adapt1[T0](fn)
}

// Adapt1RF upgrades the fn function signature to the signature of the first
// argument.
func Adapt1RF[T0, R any, Fn FuncAny1AnyR[T0, R]](_ func(T0) R, fn Fn) func(T0) R {
	return Adapt1R[T0, R](fn)
}

// Adapt2F deduces the target function signature types from the first argument.
func Adapt2F[T0, T1 any, Fn FuncAny2[T0, T1]](_ func(T0, T1), fn Fn) func(T0, T1) {
	return Adapt2[T0, T1](fn)
}

// Adapt2RF deduces the target function signature types from the first argument.
func Adapt2RF[T0, T1, R any, Fn FuncAny2AnyR[T0, T1, R]](_ func(T0, T1) R, fn Fn) func(T0, T1) R {
	return Adapt2R[T0, T1, R](fn)
}
