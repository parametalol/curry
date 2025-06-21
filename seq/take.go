package seq

import (
	"iter"
)

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

func Filter[Value any](seq iter.Seq[Value], f func(Value) bool) iter.Seq[Value] {
	return func(yield func(Value) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

func Map[Value any](seq iter.Seq[Value], f func(Value) Value) iter.Seq[Value] {
	return func(yield func(Value) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func Map15[K, V, Value any](seq iter.Seq2[K, V], f func(K, V) Value) iter.Seq[Value] {
	return func(yield func(Value) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

func Map2[Key, Value any](seq iter.Seq2[Key, Value], f func(Key, Value) (Key, Value)) iter.Seq2[Key, Value] {
	return func(yield func(Key, Value) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

func FromMap[Key comparable, Value any](m map[Key]Value) iter.Seq2[Key, Value] {
	return func(yield func(Key, Value) bool) {
		for k, v := range m {
			if !yield(k, v) {
				break
			}
		}
	}
}

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
