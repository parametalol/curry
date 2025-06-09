package seq

import (
	"slices"
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestGenerate(t *testing.T) {

	g := Generate(func(i uint) int { return int(i) })

	assert.That(t, assert.EqualSlices(
		[]int{0, 1, 2, 3, 4},
		slices.Collect(Take(5, g)),
	))
}

func TestAccumulate(t *testing.T) {
	five := Take(5, Accumulate(1, func(i uint, a int) int { return a + int(i) }))

	assert.That(t, assert.EqualSlices(
		[]int{1, 2, 4, 7, 11},
		slices.Collect(five)),
	)
}
