package curry

// region Pass

// Pass just returns the only argument.
func Pass[T0 any](arg0 T0) T0 { return arg0 }

// Pass2 just returns the both arguments.
func Pass2[T0, T1 any](arg0 T0, arg1 T1) (T0, T1) { return arg0, arg1 }

// region Return

// Return takes an argument and constructs a thunk function that returns the
// argument.
func Return[T0 any](arg0 T0) func() T0 {
	return func() T0 { return arg0 }
}

// Return2 takes two arguments and constructs a thunk function that returns
// these arguments.
func Return2[T0, T1 any](arg0 T0, arg1 T1) func() (T0, T1) {
	return func() (T0, T1) { return arg0, arg1 }
}

// ReturnS takes a variable number of arguments of any type and returns a
// thunk function, that returns a slice containing the arguments.
func ReturnS[T0 any](arg0 ...T0) func() []T0 {
	return func() []T0 { return arg0 }
}
