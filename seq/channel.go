package seq

import "iter"

// ChanAsSeq converts a <-channel to a sequence.
func ChanAsSeq[Value any](ch <-chan Value) iter.Seq[Value] {
	return func(yield func(Value) bool) {
		for v := range ch {
			if !yield(v) {
				break
			}
		}
	}
}

// SeqAsChan converts a sequence to a <-channel.
func SeqAsChan[Value any](seq iter.Seq[Value]) <-chan Value {
	ch := make(chan Value)
	go func() {
		defer close(ch)
		for v := range seq {
			ch <- v
		}
	}()
	return ch
}
