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
		assert.Equal("abcd", DropLastOfTwo(UnFour2(curriedJoinE)("a", "b", "c", "d"))),
		assert.Equal("abcd", DropLastOfTwo(UnThree(curriedJoinE)("a", "b", "c")("d"))),
		assert.Equal("abcd", DropLastOfTwo(UnTwo(curriedJoinE)("a", "b")("c")("d"))),
	)
}

func TestUnCurrySlice(t *testing.T) {
	assert.That(t,
		assert.Equal("a-b-c",
			UnTwoSlice(TwoSlice(fmt.Sprintf))("%s-%s-%s", "a", "b", "c")),

		assert.Equal(13, DropLastOfTwo(
			UnTwoSlice2(TwoSlice2(fmt.Printf))("- %s - %s -", "abc", "def"))),

		assert.Equal("5: [a b]",
			UnThreeSlice(ThreeSlice(join2slice))("%d: %s", 5, "a", "b")),

		assert.Equal(5, DropLastOfTwo(
			UnThreeSlice2(ThreeSlice2(fmt.Fprintf))(io.Discard, "%s-%s-%s", "a", "b", "c"))),

		assert.Equal("5: [a b]",
			UnTwo(ThreeSlice(join2slice))("%d: %s", 5)("a", "b")),
	)
}
