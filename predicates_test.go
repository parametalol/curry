package curry

import (
	"cmp"
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestEq(t *testing.T) {
	isFive := F2R(Eq[int])(5)

	assert.That(t,
		assert.True(isFive(5)),
		assert.False(isFive(15)),
	)
}

func TestNot(t *testing.T) {
	assert.That(t,
		assert.True(Not(false)),
		assert.False(Not(true)),

		assert.True(Wrap(
			BindLastOf2R(cmp.Less, 5),
			Not)(7)),
	)
}
