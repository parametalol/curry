package curry

import (
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestSignature(t *testing.T) {
	assert.That(t,
		assert.Equal(0,
			Signature0R(Bind1R(Return, 10)).RV0),

		assert.Equal(rune(0),
			Signature1R(Return[rune]).Arg0),
		assert.Equal(float32(0.0),
			Signature1R(Return[float32]).RV0),

		assert.Equal("",
			Signature2R(join2).Arg0),
		assert.Equal("",
			Signature2R(join2).Arg1),
		assert.Equal("",
			Signature2R(join2).RV0),
	)
}
