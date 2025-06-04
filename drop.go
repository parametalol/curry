package curry

// region DropFirst

func DropFirstOfTwo[A any, B any](_ A, b B) B {
	return b
}

func DropFirstOfThree[A any, B any, C any](_ A, b B, c C) (B, C) {
	return b, c
}

func DropFirstOfFour[A any, B any, C any, D any](_ A, b B, c C, d D) (B, C, D) {
	return b, c, d
}

// region DropLast

func DropLastOfTwo[A any, B any](a A, _ B) A {
	return a
}

func DropLastOfThree[A any, B any, C any](a A, b B, _ C) (A, B) {
	return a, b
}

func DropLastOfFour[A any, B any, C any, D any](a A, b B, c C, _ D) (A, B, C) {
	return a, b, c
}
