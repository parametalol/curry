package curry

import (
	"fmt"
	"io"
	"testing"
)

func TestBind(t *testing.T) {
	Assert(t,
		Equal("a", BindOne(join1, "a")()),
		Equal("ab", BindFirstOfTwo(join2, "a")("b")),
		Equal("abc", BindFirstOfThree(join3, "a")("b", "c")),
		Equal("abcd", BindFirstOfFour(join4, "a")("b", "c", "d")),

		Equal("a", DropLastOfTwo(BindOne2(join1e, "a")())),
		Equal("ab", DropLastOfTwo(BindFirstOfTwo2(join2e, "a")("b"))),
		Equal("abc", DropLastOfTwo(BindFirstOfThree2(join3e, "a")("b", "c"))),
		Equal("abcd", DropLastOfTwo(BindFirstOfFour2(join4e, "a")("b", "c", "d"))),

		Equal("ba", BindLastOfTwo(join2, "a")("b")),
		Equal("bca", BindLastOfThree(join3, "a")("b", "c")),
		Equal("bcda", BindLastOfFour(join4, "a")("b", "c", "d")),

		Equal("ba", DropLastOfTwo(BindLastOfTwo2(join2e, "a")("b"))),
		Equal("bca", DropLastOfTwo(BindLastOfThree2(join3e, "a")("b", "c"))),
		Equal("bcda", DropLastOfTwo(BindLastOfFour2(join4e, "a")("b", "c", "d"))),
	)
}

func TestBindSlice(t *testing.T) {
	Assert(t, Equal("- abc - def -",
		BindFirstOneSlice(fmt.Sprintf, "- %s - %s -")("abc", "def")))

	Assert(t, Equal(13, DropLastOfTwo(
		BindFirstOneSlice2(fmt.Printf, "- %s - %s -")("abc", "def"))))

	Assert(t, Equal("prefix-numbers:[1 2 3]",
		BindFirstTwoSlice(func(a, b string, c ...int) string {
			return a + "-" + b + ":" + fmt.Sprint(c)
		}, "prefix")("numbers", 1, 2, 3)))

	Assert(t, Equal(8, DropLastOfTwo(
		BindFirstTwoSlice2(fmt.Fprintf, io.Discard)("%s-%s", "qux", "quux"))))

	Assert(t, Equal("- abc - def -",
		BindLastOneSlice(fmt.Sprintf, "abc", "def")("- %s - %s -")))

	Assert(t, Equal(13, DropLastOfTwo(
		BindLastOneSlice2(fmt.Printf, "abc", "def")("- %s - %s -"))))

	Assert(t, Equal("foo-bar-[1 2 3]",
		BindLastTwoSlice(func(a, b string, c ...int) string {
			return a + "-" + b + "-" + fmt.Sprint(c)
		}, 1, 2, 3)("foo", "bar")))

	Assert(t, Equal(5, DropLastOfTwo(
		BindLastTwoSlice2(fmt.Fprintf, 1, 2, 3)(io.Discard, "%d-%d-%d"))))
}
