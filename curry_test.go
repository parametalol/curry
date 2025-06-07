package curry

import (
	"errors"
	"fmt"
	"io"
	"testing"
)

func join1(a string) string {
	return a
}

func join2(a, b string) string {
	return a + b
}

func join3(a, b, c string) string {
	return a + b + c
}

func join4(a, b, c, d string) string {
	return a + b + c + d
}

func join1e(a string) (string, error) {
	return a, nil
}

func join2e(a, b string) (string, error) {
	return a + b, nil
}

func join3e(a, b, c string) (string, error) {
	return a + b + c, nil
}

func join4e(a, b, c, d string) (string, error) {
	return a + b + c + d, nil
}

func join2slice(a string, b int, c ...string) string {
	return fmt.Sprintf(a, b, c)
}

func TestCurry(t *testing.T) {
	Assert(t,
		Equal("ab", Two(join2)("a")("b")),
		Equal("abc", Three(join3)("a")("b")("c")),
		Equal("abcd", Four(join4)("a")("b")("c")("d")),

		Equal("ab", DropLastOfTwo(Two2(join2e)("a")("b"))),
		Equal("abc", DropLastOfTwo(Three2(join3e)("a")("b")("c"))),
		Equal("abcd", DropLastOfTwo(Four2(join4e)("a")("b")("c")("d"))),
	)
}

func TestCurrySlice(t *testing.T) {
	Assert(t,
		Equal("- abc - def -",
			TwoSlice(fmt.Sprintf)("- %s - %s -")("abc", "def")),

		Equal(13, DropLastOfTwo(
			TwoSlice2(fmt.Printf)("- %s - %s -")("abc", "def"))),

		Equal("5: [abc def]",
			ThreeSlice(join2slice)("%d: %v")(5)("abc", "def")),

		Equal(13, DropLastOfTwo(
			ThreeSlice2(fmt.Fprintf)(io.Discard)("- %s - %s -")("abc", "def"))),
	)
}

func TestCombinations(t *testing.T) {
	bindFirstTwoOf3 := UnTwo(Three(join3))

	t.Run("bind 2nd of three", func(t *testing.T) {
		boundSecondOfThree := UnTwo(BindLastOfTwo(bindFirstTwoOf3, "b"))
		Assert(t, Equal("abc", boundSecondOfThree("a", "c")))
	})

	t.Run("bind first two of three", func(t *testing.T) {
		boundFirstTwo := bindFirstTwoOf3("a", "b")
		Assert(t, Equal("abc", boundFirstTwo("c")))
	})

}

func TestDifferentTypes(t *testing.T) {

	f1 := func(int, string) bool {
		return true
	}
	var b bool
	b = Two(f1)(42)("abc")
	Assert(t, True(b))

	testErr := errors.New("err")
	f1e := func(int, string) (bool, error) {
		return true, testErr
	}
	var err error
	b, err = Two2(f1e)(42)("abc")
	Assert(t, True(b))
	Assert(t, ErrorIs(err, testErr))

	b = UnTwo(Two(f1))(42, "abc")
	Assert(t, True(b))

	b, err = UnTwo2(Two2(f1e))(42, "abc")
	Assert(t, True(b))
	Assert(t, ErrorIs(err, testErr))

	err = DropFirstOfTwo(f1e(42, "abc"))
	Assert(t, ErrorIs(err, testErr))
	b = DropLastOfTwo(f1e(42, "abc"))
	Assert(t, True(b))
}
