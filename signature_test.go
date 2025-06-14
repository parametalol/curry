package curry

import (
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestSignature(t *testing.T) {
	assert.That(t,
		assert.Equal(0,
			SignatureNone(BindOne(Pass, 10)).RV1),

		assert.Equal(rune(0),
			SignatureOne(Pass[rune]).Arg1),
		assert.Equal(float32(0.0),
			SignatureOne(Pass[float32]).RV1),

		assert.Equal("",
			SignatureTwo(join2).Arg1),
		assert.Equal("",
			SignatureTwo(join2).Arg2),
		assert.Equal("",
			SignatureTwo(join2).RV1),
	)
}
