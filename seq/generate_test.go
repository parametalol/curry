package seq

import (
	"slices"
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestGenerate(t *testing.T) {
	assert.That(t, assert.EqualSlices(
		[]int{0, 1, 2, 3, 4},
		slices.Collect(Take(5, Generate(Index(0)))),
	))
}

func TestRange(t *testing.T) {
	assert.That(t,
		assert.EqualSlices(
			[]int{5, 7, 9, 11, 13},
			slices.Collect(Range(5, 15, 2)),
		),
		assert.EqualSlices(
			[]int{},
			slices.Collect(Range(5, 5, 1)),
		),
	)
}

func TestGenerator(t *testing.T) {
	runes := []rune{}
	five := Take(5, Generate(
		Generator(&runes,
			func(a *[]rune) int {
				(*a) = append((*a), '.')
				return len(*a)
			})))

	assert.That(t,
		assert.EqualSlices(
			[]int{1, 2, 3, 4, 5},
			slices.Collect(five)),

		assert.Equal(".....", string(runes)),
	)
}

func TestZip(t *testing.T) {
	type pair struct{ A, B int }

	t.Run("equal length", func(t *testing.T) {
		a := Take(5, Generate(Index(1))) // 1,2,3,4,5
		b := Range(10, 15, 1)            // 10,11,12,13,14

		zipped := Zip(a, b)
		got := slices.Collect(Map15(zipped, func(a, b int) pair { return pair{a, b} }))

		want := []pair{
			{1, 10},
			{2, 11},
			{3, 12},
			{4, 13},
			{5, 14},
		}
		assert.That(t, assert.EqualSlices(want, got))
	})

	t.Run("different length", func(t *testing.T) {
		aShort := Range(1, 4, 1)  // 1,2,3
		bLong := Range(10, 20, 2) // 10,12,14,16,18

		zipped2 := Zip(aShort, bLong)
		got2 := slices.Collect(Map15(zipped2, func(a, b int) pair { return pair{a, b} }))
		want2 := []pair{
			{1, 10},
			{2, 12},
			{3, 14},
			{0, 16},
			{0, 18},
		}
		assert.That(t, assert.EqualSlices(want2, got2))
	})

	t.Run("empty sequences", func(t *testing.T) {

		a := Range(0, 0, 1)
		b := Range(0, 0, 1)
		zippedEmpty := Zip(a, b)
		gotEmpty := slices.Collect(Map15(zippedEmpty, func(a, b int) pair { return pair{a, b} }))
		assert.That(t, assert.EqualSlices([]pair{}, gotEmpty))
	})
}

func TestZipShort(t *testing.T) {
	type pair struct{ A, B int }

	t.Run("equal length", func(t *testing.T) {
		a := Take(5, Generate(Index(1))) // 1,2,3,4,5
		b := Range(10, 15, 1)            // 10,11,12,13,14

		zipped := ZipShort(a, b)
		got := slices.Collect(Map15(zipped, func(a, b int) pair { return pair{a, b} }))

		want := []pair{
			{1, 10},
			{2, 11},
			{3, 12},
			{4, 13},
			{5, 14},
		}
		assert.That(t, assert.EqualSlices(want, got))
	})

	t.Run("different length", func(t *testing.T) {
		aShort := Range(1, 4, 1)  // 1,2,3
		bLong := Range(10, 20, 2) // 10,12,14,16,18

		zipped2 := ZipShort(aShort, bLong)
		got2 := slices.Collect(Map15(zipped2, func(a, b int) pair { return pair{a, b} }))
		want2 := []pair{
			{1, 10},
			{2, 12},
			{3, 14},
		}
		assert.That(t, assert.EqualSlices(want2, got2))
	})

	t.Run("empty sequences", func(t *testing.T) {
		zippedEmpty := ZipShort(Range(0, 0, 1), Range(0, 0, 1))
		gotEmpty := slices.Collect(Map15(zippedEmpty, func(a, b int) pair { return pair{a, b} }))
		assert.That(t, assert.EqualSlices([]pair{}, gotEmpty))
	})
}
