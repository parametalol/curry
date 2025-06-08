package curry

// region DropFirst

func DropFirstOfTwo[A, B any](_ A, b B) B {
	return b
}

func DropFirstOfThree[A, B, C any](_ A, b B, c C) (B, C) {
	return b, c
}

func DropFirstOfFour[A, B, C, D any](_ A, b B, c C, d D) (B, C, D) {
	return b, c, d
}

// region DropLast

func DropLastOfTwo[A, B any](a A, _ B) A {
	return a
}

func DropLastOfThree[A, B, C any](a A, b B, _ C) (A, B) {
	return a, b
}

func DropLastOfFour[A, B, C, D any](a A, b B, c C, _ D) (A, B, C) {
	return a, b, c
}
