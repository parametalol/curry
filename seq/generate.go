package seq

import "iter"

// Generate a sequence by calling the passed generator function.
func Generate[Value any](generator func() Value) iter.Seq[Value] {
	return func(yield func(Value) bool) {
		for yield(generator()) {
		}
	}
}

// Range generates a sequnce of integer between the given begin and end,
// with given the step.
func Range(begin, end, step int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := begin; i < end; i += step {
			if !yield(i) {
				return
			}
		}
	}
}

// Generator makes a generator providing context to the given function.
func Generator[Value, Context any](c Context, f func(Context) Value) func() Value {
	return func() Value {
		return f(c)
	}
}

// Index returns a sequence index generator function.
func Index[I ~int | ~uint](i I) func() I {
	return func() I {
		i++
		return i - 1
	}
}

func Zip[A, B any](a iter.Seq[A], b iter.Seq[B]) iter.Seq2[A, B] {
	nextA, stopA := iter.Pull(a)
	nextB, stopB := iter.Pull(b)
	return func(yield func(A, B) bool) {
		defer stopA()
		defer stopB()
		for {
			a, okA := nextA()
			b, okB := nextB()
			if !okA && !okB || !yield(a, b) {
				break
			}
		}
	}
}

func ZipShort[A, B any](a iter.Seq[A], b iter.Seq[B]) iter.Seq2[A, B] {
	nextA, stopA := iter.Pull(a)
	nextB, stopB := iter.Pull(b)
	return func(yield func(A, B) bool) {
		defer stopA()
		defer stopB()
		for {
			a, ok := nextA()
			if !ok {
				break
			}
			b, ok := nextB()
			if !ok || !yield(a, b) {
				break
			}
		}
	}
}
