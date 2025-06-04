package curry

// region DropFirst

func DropFirstOf2[A any, B any](_ A, b B) B {
	return b
}

func DropFirstOf3[A any, B any, C any](_ A, b B, c C) (B, C) {
	return b, c
}

func DropFirstOf4[A any, B any, C any, D any](_ A, b B, c C, d D) (B, C, D) {
	return b, c, d
}

// region DropLast

func DropLastOf2[A any, B any](a A, _ B) A {
	return a
}

func DropLastOf3[A any, B any, C any](a A, b B, _ C) (A, B) {
	return a, b
}

func DropLastOf4[A any, B any, C any, D any](a A, b B, c C, _ D) (A, B, C) {
	return a, b, c
}
