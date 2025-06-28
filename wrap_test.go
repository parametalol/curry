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
			Wrap(Wrap(Wrap(Return(
				strings.ToUpper,
			), strings.Fields,
			), slices.Compact[[]string],
			), BindLastOf2R(strings.Join, ","),
			)(" a a a b c c c c de")),

		assert.Equal(true,
			Wrap(Wrap(Wrap(Return(
				strings.ToUpper,
			), strings.Fields,
			), slices.Compact[[]string],
			), slices.IsSorted,
			)("a a b c d d e")),
	)
}
