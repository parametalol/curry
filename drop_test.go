package curry

import (
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestDrop(t *testing.T) {
	DropFirstOf1("dropped")
	assert.That(t,
		assert.Equal(1, DropLastOf2(1, 2)),
		assert.Equal(1, DropLastOf2(DropLastOf3(1, 2, 3))),
		assert.Equal(1, DropLastOf2(DropLastOf3(DropLastOf4(1, 2, 3, 4)))),

		assert.Equal(2, DropFirstOf2(1, 2)),
		assert.Equal(3, DropFirstOf2(DropFirstOf3(1, 2, 3))),
		assert.Equal(4, DropFirstOf2(DropFirstOf3(DropFirstOf4(1, 2, 3, 4)))),
	)
}
