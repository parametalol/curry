package curry

import (
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestPass(t *testing.T) {
	assert.That(t,
		assert.Equal("abc", Return("abc")),
		assert.Equal("abc", DropLastOf2(Return2("abc", 45))),
		assert.Equal(45, DropFirstOf2(Return2("abc", 45))),
	)
}

func TestReturn(t *testing.T) {
	assert.That(t, assert.Equal("abc", Thunk("abc")()))
	a, b := Thunk2("abc", 123)()
	assert.That(t, assert.Equal("abc", a))
	assert.That(t, assert.Equal(123, b))

	i := ThunkS(1, 2, 3)
	assert.That(t, assert.Equal(1, i()[0]))
	assert.That(t, assert.Equal(2, i()[1]))
	assert.That(t, assert.Equal(3, i()[2]))
}
