package curry

import (
	"fmt"
	"io"
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestBind(t *testing.T) {
	assert.That(t,
		assert.Equal("a", Bind1R(join1, "a")()),
		assert.Equal("ab", BindFirstOf2R(join2, "a")("b")),
		assert.Equal("abc", BindFirstOf3R(join3, "a")("b", "c")),
		assert.Equal("abcd", BindFirstOf4R(join4, "a")("b", "c", "d")),

		assert.Equal("a", DropLastOf2(Bind1R2(join1e, "a")())),
		assert.Equal("ab", DropLastOf2(BindFirstOf2R2(join2e, "a")("b"))),
		assert.Equal("abc", DropLastOf2(BindFirstOf3R2(join3e, "a")("b", "c"))),
		assert.Equal("abcd", DropLastOf2(BindFirstOf4R2(join4e, "a")("b", "c", "d"))),

		assert.Equal("ba", BindLastOf2R(join2, "a")("b")),
		assert.Equal("bca", BindLastOf3R(join3, "a")("b", "c")),
		assert.Equal("bcda", BindLastOf4R(join4, "a")("b", "c", "d")),

		assert.Equal("ba", DropLastOf2(BindLastOf2R2(join2e, "a")("b"))),
		assert.Equal("bca", DropLastOf2(BindLastOf3R2(join3e, "a")("b", "c"))),
		assert.Equal("bcda", DropLastOf2(BindLastOf4R2(join4e, "a")("b", "c", "d"))),
	)
}

func TestBindSlice(t *testing.T) {
	assert.That(t,
		assert.Equal("- abc - def -",
			BindFirstOf2SR(fmt.Sprintf, "- %s - %s -")("abc", "def")),

		assert.Equal(13, DropLastOf2(
			BindFirstOf2SR2(testFmt2, "- %s - %s -")("abc", "def"))),

		assert.Equal("42: a b c",
			BindFirstOf3SR(join2slice, "%s %s %s")(42, "a", "b", "c")),

		assert.Equal(8, DropLastOf2(
			BindFirstOf3SR2(fmt.Fprintf, io.Discard)("%s-%s", "qux", "quux"))),

		assert.Equal("- abc - def -",
			BindLastOf2SR(fmt.Sprintf, "abc", "def")("- %s - %s -")),

		assert.Equal(13, DropLastOf2(
			BindLastOf2SR2(testFmt2, "abc", "def")("- %s - %s -"))),

		assert.Equal("42: a b c",
			BindLastOf3SR(join2slice, "a", "b", "c")("%s %s %s", 42)),

		assert.Equal(5, DropLastOf2(
			BindLastOf3SR2(fmt.Fprintf, 1, 2, 3)(io.Discard, "%d-%d-%d"))),
	)
}

func TestBind1(t *testing.T) {
	i := 0
	set := func(j int) {
		i = j
	}
	set5 := Bind1(set, 5)
	set5()
	assert.That(t, assert.Equal(5, i))
}
