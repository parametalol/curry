package seq

import (
	"iter"
)

func Take[A any](n uint, seq iter.Seq[A]) iter.Seq[A] {
	return func(yield func(A) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for range n {
			v, ok := next()
			if !ok || !yield(v) {
				return
			}
		}
	}
}

func Tail[A any](seq iter.Seq[A]) (A, bool, iter.Seq[A]) {
	next, stop := iter.Pull(seq)
	v, ok := next()
	return v, ok, func(yield func(A) bool) {
		defer stop()
		for ok {
			v, ok = next()
			if !ok || !yield(v) {
				return
			}
		}
	}
}

func Filter[A any](seq iter.Seq[A], f func(A) bool) iter.Seq[A] {
	return func(yield func(A) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

func Map[A any](seq iter.Seq[A], f func(A) A) iter.Seq[A] {
	return func(yield func(A) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func Until[A any](seq iter.Seq[A], f func(A) bool) iter.Seq[A] {
	return func(yield func(A) bool) {
		for v := range seq {
			if f(v) || !yield(v) {
				return
			}
		}
	}
}

func Purge[A any](seq iter.Seq[A]) (i int) {
	for range seq {
		i++
	}
	return
}
