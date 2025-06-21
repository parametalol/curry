package curry

func Eq[A comparable](a A) func(A) bool {
	return func(b A) bool {
		return a == b
	}
}
