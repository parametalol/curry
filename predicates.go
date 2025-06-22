package curry

func Eq[Value comparable](a, b Value) bool {
	return a == b
}

func Not(b bool) bool {
	return !b
}

func LenString[Value ~string](v Value) int                    { return len(v) }
func LenSlice[Value ~[]V, V any](v Value) int                 { return len(v) }
func LenMap[Value ~map[K]V, V any, K comparable](v Value) int { return len(v) }
