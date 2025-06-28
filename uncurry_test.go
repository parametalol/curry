package curry

import (
	"fmt"
	"io"
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestUnCurry(t *testing.T) {
	curriedJoin := F4R(join4)

	assert.That(t,
		assert.Equal("abcd", Un4R(curriedJoin)("a", "b", "c", "d")),
		assert.Equal("abcd", Un3R(curriedJoin)("a", "b", "c")("d")),
		assert.Equal("abcd", Un2R(curriedJoin)("a", "b")("c")("d")),
	)

	curriedJoinE := F4R2(join4e)
	assert.That(t,
		assert.Equal("abcd", DropLastOf2(Un4R2(curriedJoinE)("a", "b", "c", "d"))),
		assert.Equal("abcd", DropLastOf2(Un3R(curriedJoinE)("a", "b", "c")("d"))),
		assert.Equal("abcd", DropLastOf2(Un2R(curriedJoinE)("a", "b")("c")("d"))),
	)
}

func TestUnCurrySlice(t *testing.T) {
	assert.That(t,
		assert.Equal("a-b-c",
			Un2SR(F2SR(fmt.Sprintf))("%s-%s-%s", "a", "b", "c")),

		assert.Equal(13, DropLastOf2(
			Un2SR2(F2SR2(testFmt2))("- %s - %s -", "abc", "def"))),

		assert.Equal("5: a b",
			Un3SR(F3SR(join2slice))("%s %s", 5, "a", "b")),

		assert.Equal(5, DropLastOf2(
			Un3SR2(F3SR2(fmt.Fprintf))(io.Discard, "%s-%s-%s", "a", "b", "c"))),

		assert.Equal("5: a b",
			Un2R(F3SR(join2slice))("%s %s", 5)("a", "b")),
	)
}
