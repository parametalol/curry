package curry

// region DropFirst

// DropOne
func DropOne(any) {}

func DropFirstOfTwo[B any](_ any, b B) B {
	return b
}

func DropFirstOfThree[B, C any](_ any, b B, c C) (B, C) {
	return b, c
}

func DropFirstOfFour[B, C, D any](_ any, b B, c C, d D) (B, C, D) {
	return b, c, d
}

// region DropLast

func DropLastOfTwo[A any](a A, _ any) A {
	return a
}

func DropLastOfThree[A, B any](a A, b B, _ any) (A, B) {
	return a, b
}

func DropLastOfFour[A, B, C any](a A, b B, c C, _ any) (A, B, C) {
	return a, b, c
}
