package seq

import (
	"slices"
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestTake(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		i := slices.Collect(Take(5, slices.Values([]int{})))
		assert.That(t, assert.EqualSlices([]int{}, i))
	})

	t.Run("not empty", func(t *testing.T) {
		i := []int{15, 24, 33, 42, 51, 60, 79}

		five := slices.Collect(Take(5, slices.Values(i)))

		assert.That(t, assert.True(
			slices.Equal(i[0:5], five),
		))
	})

	t.Run("over request", func(t *testing.T) {
		i := []int{15, 24, 33}

		three := slices.Collect(Take(5, slices.Values(i)))

		assert.That(t, assert.True(
			slices.Equal(i[0:3], three),
		))
	})
}

func TestTail(t *testing.T) {
	t.Run("empty", func(t *testing.T) {

		v, ok, tail := Tail(slices.Values([]int{}))

		assert.That(t,
			assert.Equal(0, v),
			assert.False(ok),
			assert.EqualSlices([]int{}, slices.Collect(tail)),
		)
	})

	t.Run("not empty", func(t *testing.T) {
		i := []int{15, 24, 33, 42, 51, 60, 79}

		v, ok, tail := Tail(slices.Values(i))

		assert.That(t,
			assert.Equal(i[0], v),
			assert.True(ok),
			assert.EqualSlices(i[1:], slices.Collect(tail)),
		)
	})

	t.Run("one", func(t *testing.T) {
		i := []int{15}

		v, ok, tail := Tail(slices.Values(i))

		assert.That(t,
			assert.Equal(i[0], v),
			assert.True(ok),
			assert.EqualSlices(i[1:], slices.Collect(tail)),
		)
	})
}
