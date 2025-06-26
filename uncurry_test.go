package curry

import (
	"fmt"
	"io"
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestUnCurry(t *testing.T) {
	curriedJoin := Four(join4)

	assert.That(t,
		assert.Equal("abcd", UnFour(curriedJoin)("a", "b", "c", "d")),
		assert.Equal("abcd", UnThree(curriedJoin)("a", "b", "c")("d")),
		assert.Equal("abcd", UnTwo(curriedJoin)("a", "b")("c")("d")),
	)

	curriedJoinE := Four2(join4e)
	assert.That(t,
		assert.Equal("abcd", DropLastOf2(UnFour2(curriedJoinE)("a", "b", "c", "d"))),
		assert.Equal("abcd", DropLastOf2(UnThree(curriedJoinE)("a", "b", "c")("d"))),
		assert.Equal("abcd", DropLastOf2(UnTwo(curriedJoinE)("a", "b")("c")("d"))),
	)
}

func TestUnCurrySlice(t *testing.T) {
	assert.That(t,
		assert.Equal("a-b-c",
			UnTwoSlice(TwoSlice(fmt.Sprintf))("%s-%s-%s", "a", "b", "c")),

		assert.Equal(13, DropLastOf2(
			UnTwoSlice2(TwoSlice2(testFmt2))("- %s - %s -", "abc", "def"))),

		assert.Equal("5: a b",
			UnThreeSlice(ThreeSlice(join2slice))("%s %s", 5, "a", "b")),

		assert.Equal(5, DropLastOf2(
			UnThreeSlice2(ThreeSlice2(fmt.Fprintf))(io.Discard, "%s-%s-%s", "a", "b", "c"))),

		assert.Equal("5: a b",
			UnTwo(ThreeSlice(join2slice))("%s %s", 5)("a", "b")),
	)
}
