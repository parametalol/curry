package curry

import (
	"fmt"
	"io"
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestBind(t *testing.T) {
	assert.That(t,
		assert.Equal("a", BindOne(join1, "a")()),
		assert.Equal("ab", BindFirstOfTwo(join2, "a")("b")),
		assert.Equal("abc", BindFirstOfThree(join3, "a")("b", "c")),
		assert.Equal("abcd", BindFirstOfFour(join4, "a")("b", "c", "d")),

		assert.Equal("a", DropLastOfTwo(BindOne2(join1e, "a")())),
		assert.Equal("ab", DropLastOfTwo(BindFirstOfTwo2(join2e, "a")("b"))),
		assert.Equal("abc", DropLastOfTwo(BindFirstOfThree2(join3e, "a")("b", "c"))),
		assert.Equal("abcd", DropLastOfTwo(BindFirstOfFour2(join4e, "a")("b", "c", "d"))),

		assert.Equal("ba", BindLastOfTwo(join2, "a")("b")),
		assert.Equal("bca", BindLastOfThree(join3, "a")("b", "c")),
		assert.Equal("bcda", BindLastOfFour(join4, "a")("b", "c", "d")),

		assert.Equal("ba", DropLastOfTwo(BindLastOfTwo2(join2e, "a")("b"))),
		assert.Equal("bca", DropLastOfTwo(BindLastOfThree2(join3e, "a")("b", "c"))),
		assert.Equal("bcda", DropLastOfTwo(BindLastOfFour2(join4e, "a")("b", "c", "d"))),
	)
}

func TestBindSlice(t *testing.T) {
	assert.That(t,
		assert.Equal("- abc - def -",
			BindFirstOneSlice(fmt.Sprintf, "- %s - %s -")("abc", "def")),

		assert.Equal(13, DropLastOfTwo(
			BindFirstOneSlice2(testFmt2, "- %s - %s -")("abc", "def"))),

		assert.Equal("42: a b c",
			BindFirstTwoSlice(join2slice, "%s %s %s")(42, "a", "b", "c")),

		assert.Equal(8, DropLastOfTwo(
			BindFirstTwoSlice2(fmt.Fprintf, io.Discard)("%s-%s", "qux", "quux"))),

		assert.Equal("- abc - def -",
			BindLastOneSlice(fmt.Sprintf, "abc", "def")("- %s - %s -")),

		assert.Equal(13, DropLastOfTwo(
			BindLastOneSlice2(testFmt2, "abc", "def")("- %s - %s -"))),

		assert.Equal("42: a b c",
			BindLastTwoSlice(join2slice, "a", "b", "c")("%s %s %s", 42)),

		assert.Equal(5, DropLastOfTwo(
			BindLastTwoSlice2(fmt.Fprintf, 1, 2, 3)(io.Discard, "%d-%d-%d"))),
	)
}
