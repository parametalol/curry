package curry

import (
	"slices"
	"strings"
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestWrap(t *testing.T) {
	assert.That(t,
		assert.Equal("A,B,C,DE",
			Wrap(Wrap(Wrap(Pass(
				strings.ToUpper,
			), strings.Fields,
			), slices.Compact[[]string],
			), BindLastOfTwo(strings.Join, ","),
			)(" a a a b c c c c de")),

		assert.Equal(true,
			Wrap(Wrap(Wrap(Pass(
				strings.ToUpper,
			), strings.Fields,
			), slices.Compact[[]string],
			), slices.IsSorted,
			)("a a b c d d e")),

		assert.Equal(8,
			Wrap(
				Number[int],
				(*NumberT[int]).Inc,
				(*NumberT[int]).Inc,
				(*NumberT[int]).Inc,
			)(5),
		),
	)
}
