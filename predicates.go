package curry

func Eq[Value comparable](a, b Value) bool {
	return a == b
}

func Not(b bool) bool {
	return !b
}
