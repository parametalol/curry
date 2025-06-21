package seq

import (
	"iter"
	"slices"
	"testing"
	"time"

	"github.com/parametalol/curry/assert"
)

func TestChanAsSeq(t *testing.T) {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()

	got := slices.Collect(ChanAsSeq(ch))
	want := []int{1, 2, 3}
	assert.That(t,
		assert.EqualSlices(want, got),
	)
}

func TestChanEarlyStop(t *testing.T) {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()

	seq := ChanAsSeq(ch)
	var got []int
	seq(func(v int) bool {
		got = append(got, v)
		return v < 2 // stop after first two values
	})

	want := []int{1, 2}
	assert.That(t,
		assert.EqualSlices(want, got),
	)
}

func TestSeqAsChan(t *testing.T) {
	// Create a sequence using iter.Seq
	seq := iter.Seq[int](func(yield func(int) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i) {
				break
			}
		}
	})

	ch := SeqAsChan(seq)
	var got []int
	for v := range ch {
		got = append(got, v)
	}

	want := []int{1, 2, 3}
	assert.That(t,
		assert.EqualSlices(want, got),
	)
}

func TestSeqEmpty(t *testing.T) {
	seq := iter.Seq[int](func(yield func(int) bool) {})
	ch := SeqAsChan(seq)
	select {
	case v, ok := <-ch:
		if ok {
			t.Errorf("Seq() expected closed channel, got value %v", v)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Seq() did not close channel in time")
	}
}
