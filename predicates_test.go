package curry

import (
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestEq(t *testing.T) {
	isFive := Eq(5)

	assert.That(t,
		assert.True(isFive(5)),
		assert.False(isFive(15)),
	)
}
