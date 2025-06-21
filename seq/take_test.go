package seq

import (
	"slices"
	"strconv"
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

func TestFilter(t *testing.T) {
	odd := func(i int) bool { return i&1 == 1 }

	i := []int{1, 2, 3, 4, 5}
	assert.That(t,
		assert.EqualSlices([]int{1, 3, 5},
			slices.Collect(
				Filter(slices.Values(i), odd))),

		assert.EqualSlices([]int{}, slices.Collect(
			Filter(slices.Values([]int{}), odd))),
	)
}

func TestMap(t *testing.T) {
	mul := func(i int) int { return i * 2 }

	i := []int{1, 2, 3, 4, 5}
	assert.That(t,
		assert.EqualSlices([]int{2, 4, 6, 8, 10},
			slices.Collect(
				Map(slices.Values(i), mul))),

		assert.EqualSlices([]int{},
			slices.Collect(
				Map(slices.Values([]int{}), mul))),
	)
}

func TestMap15(t *testing.T) {
	mul := func(a, b int) int { return a * b }

	i := map[int]int{2: 5}
	assert.That(t,
		assert.EqualSlices([]int{10},
			slices.Collect(Map15(FromMap(i), mul))),

		assert.EqualSlices([]int{},
			slices.Collect(
				Map15(FromMap(map[string]int{}), func(string, int) int { return 0 }))),
	)
}

func TestUntil(t *testing.T) {
	i := []int{1, 2, 3, 4, 5}
	assert.That(t,
		assert.EqualSlices([]int{1, 2, 3},
			slices.Collect(
				Until(slices.Values(i), func(i int) bool { return i > 3 }))),
	)
}

func TestAccumulate(t *testing.T) {
	length := func(_, i int) int { return i + 1 }
	assert.That(t,
		assert.Equal(4,
			Accumulate(slices.Values([]int{2, 4, 7, 9}), length)),
		assert.Equal(0,
			Accumulate(slices.Values([]int{}), length)),

		assert.EqualSlices([]byte("2479"),
			Accumulate(slices.Values([]int64{2, 4, 7, 9}),
				func(v int64, a []byte) []byte {
					return strconv.AppendInt(a, v, 10)
				})),
	)
}

func TestLast(t *testing.T) {
	assert.That(t,
		assert.Equal(4,
			Last(Take(5, Generate(Index(0))))),

		assert.Equal(0,
			Last(slices.Values([]int{}))),
	)
}
