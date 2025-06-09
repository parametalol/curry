package seq

import "iter"

func Generate[A any](f func(uint) A) iter.Seq[A] {
	return func(yield func(A) bool) {
		for i := uint(0); yield(f(i)); i++ {
		}
	}
}

func Accumulate[A any](a A, f func(uint, A) A) iter.Seq[A] {
	return func(yield func(A) bool) {
		for i := uint(0); ; i++ {
			a = f(i, a)
			if !yield(a) {
				break
			}
		}
	}
}
