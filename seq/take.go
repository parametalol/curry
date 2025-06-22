package seq

import (
	"iter"
)

// Take first n values from the sequence.
func Take[Value any](n uint, seq iter.Seq[Value]) iter.Seq[Value] {
	return func(yield func(Value) bool) {
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

// Tail reads the first value from the sequence, returns whether it managed to
// get it, and the iterator to the rest of the sequence.
func Tail[Value any](seq iter.Seq[Value]) (Value, bool, iter.Seq[Value]) {
	next, stop := iter.Pull(seq)
	v, ok := next()
	return v, ok, func(yield func(Value) bool) {
		defer stop()
		for ok {
			v, ok = next()
			if !ok || !yield(v) {
				return
			}
		}
	}
}

// Filter returns a sequence with the values for which f returns true.
func Filter[Value any](seq iter.Seq[Value], f func(Value) bool) iter.Seq[Value] {
	return func(yield func(Value) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

// Map translates sequence values with the provided function.
func Map[ValueIn, ValueOut any](seq iter.Seq[ValueIn], f func(ValueIn) ValueOut) iter.Seq[ValueOut] {
	return func(yield func(ValueOut) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// Map15 translates key-value sequence to a single value sequence with the
// provided function.
func Map15[K, V, Value any](seq iter.Seq2[K, V], f func(K, V) Value) iter.Seq[Value] {
	return func(yield func(Value) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

// Map2 translates one key-value sequence to another with the provided function.
func Map2[KeyIn, ValueIn, KeyOut, ValueOut any](seq iter.Seq2[KeyIn, ValueIn], f func(KeyIn, ValueIn) (KeyOut, ValueOut)) iter.Seq2[KeyOut, ValueOut] {
	return func(yield func(KeyOut, ValueOut) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

// Until passes values from the sequence until f returns true.
func Until[Value any](seq iter.Seq[Value], f func(Value) bool) iter.Seq[Value] {
	return func(yield func(Value) bool) {
		for v := range seq {
			if f(v) || !yield(v) {
				return
			}
		}
	}
}

// Last consumes the sequence and returns the last value.
func Last[Value any](seq iter.Seq[Value]) (result Value) {
	for v := range seq {
		result = v
	}
	return
}

// Accumulate passes sequence values through the accumulator function and
// returns the accumulated value.
func Accumulate[A, Value any](seq iter.Seq[Value], acc func(Value, A) A) A {
	var a A
	for v := range seq {
		a = acc(v, a)
	}
	return a
}
