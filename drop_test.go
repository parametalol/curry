package curry

import (
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestDrop(t *testing.T) {
	DropOne("dropped")
	assert.That(t,
		assert.Equal(1, DropLastOfTwo(1, 2)),
		assert.Equal(1, DropLastOfTwo(DropLastOfThree(1, 2, 3))),
		assert.Equal(1, DropLastOfTwo(DropLastOfThree(DropLastOfFour(1, 2, 3, 4)))),

		assert.Equal(2, DropFirstOfTwo(1, 2)),
		assert.Equal(3, DropFirstOfTwo(DropFirstOfThree(1, 2, 3))),
		assert.Equal(4, DropFirstOfTwo(DropFirstOfThree(DropFirstOfFour(1, 2, 3, 4)))),
	)
}
