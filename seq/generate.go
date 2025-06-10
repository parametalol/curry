package seq

import "iter"

func Generate[A any](f func(uint) A) iter.Seq[A] {
	return func(yield func(A) bool) {
		for i := uint(0); yield(f(i)); i++ {
		}
	}
}

func Range(begin, end, step int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := begin; i < end; i += step {
			if !yield(i) {
				return
			}
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
