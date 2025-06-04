package curry

// region DropFirst

func DropFirst2[A any, B any](_ A, b B) B {
	return b
}

func DropFirst3[A any, B any, C any](_ A, b B, c C) (B, C) {
	return b, c
}

func DropFirst4[A any, B any, C any, D any](_ A, b B, c C, d D) (B, C, D) {
	return b, c, d
}

// region DropLast

func DropLast2[A any, B any](a A, _ B) A {
	return a
}

func DropLast3[A any, B any, C any](a A, b B, _ C) (A, B) {
	return a, b
}

func DropLast4[A any, B any, C any, D any](a A, b B, c C, _ D) (A, B, C) {
	return a, b, c
}
