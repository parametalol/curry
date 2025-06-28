package curry

// region Pass

// Return just returns the only argument.
func Return[T0 any](arg0 T0) T0 { return arg0 }

// Return2 just returns the both arguments.
func Return2[T0, T1 any](arg0 T0, arg1 T1) (T0, T1) { return arg0, arg1 }

// Pass just returns the only argument.
func ReturnS[T0 any](arg0 ...T0) []T0 { return arg0 }

// region Return

// Thunk takes an argument and constructs a function that returns the argument.
func Thunk[T0 any](arg0 T0) func() T0 {
	return Bind1R(Return, arg0)
}

// Thunk2 takes two arguments and constructs a function that returns these
// arguments.
func Thunk2[T0, T1 any](arg0 T0, arg1 T1) func() (T0, T1) {
	return Bind1R2(BindFirstOf2R2(Return2[T0, T1], arg0), arg1)
}

// ThunkS takes a variable number of arguments of any type and returns a
// function, that returns a slice containing the arguments.
func ThunkS[T0 any](arg0 ...T0) func() []T0 {
	return Bind1SR(ReturnS, arg0...)
}
