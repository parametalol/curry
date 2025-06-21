package curry

import (
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestNumber(t *testing.T) {
	n := Number(5)
	inc := n.Inc
	inc()
	inc()
	inc()
	assert.That(t, assert.Equal(9, inc()))
	dec := n.Dec
	assert.That(t, assert.Equal(8, dec()))
	assert.That(t, assert.Equal(16, n.Add(8)))
	assert.That(t, assert.Equal(16, n.N))
}
