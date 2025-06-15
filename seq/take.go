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

func Until[Value any](seq iter.Seq[Value], f func(Value) bool) iter.Seq[Value] {
	return func(yield func(Value) bool) {
		for v := range seq {
			if f(v) || !yield(v) {
				return
			}
		}
	}
}

func Purge[Value any](seq iter.Seq[Value]) (i int) {
	for range seq {
		i++
	}
	return
}

func Last[Value any](seq iter.Seq[Value]) (result Value) {
	for v := range seq {
		result = v
	}
	return
}
