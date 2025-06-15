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
