package curry

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"testing"

	"github.com/parametalol/curry/assert"
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

func testFmt2(format string, b ...any) (int, string) {
	result := fmt.Sprintf(format, b...)
	return len(result), result
}

func join2slice(format string, b int64, c ...any) string {
	return strconv.FormatInt(b, 10) + ": " + fmt.Sprintf(format, c...)
}

func TestCurry(t *testing.T) {
	assert.That(t,
		assert.Equal("ab", Two(join2)("a")("b")),
		assert.Equal("abc", Three(join3)("a")("b")("c")),
		assert.Equal("abcd", Four(join4)("a")("b")("c")("d")),

		assert.Equal("ab", DropLastOfTwo(Two2(join2e)("a")("b"))),
		assert.Equal("abc", DropLastOfTwo(Three2(join3e)("a")("b")("c"))),
		assert.Equal("abcd", DropLastOfTwo(Four2(join4e)("a")("b")("c")("d"))),
	)
}

func TestCurrySlice(t *testing.T) {
	assert.That(t,
		assert.Equal("- abc - def -",
			TwoSlice(fmt.Sprintf)("- %s - %s -")("abc", "def")),

		assert.Equal(13, DropLastOfTwo(
			TwoSlice2(testFmt2)("- %s - %s -")("abc", "def"))),

		assert.Equal("5: abc def",
			ThreeSlice(join2slice)("%s %s")(5)("abc", "def")),

		assert.Equal(13, DropLastOfTwo(
			ThreeSlice2(fmt.Fprintf)(io.Discard)("- %s - %s -")("abc", "def"))),
	)
}

func TestCombinations(t *testing.T) {
	bindFirstTwoOf3 := UnTwo(Three(join3))

	t.Run("bind 2nd of three", func(t *testing.T) {
		boundSecondOfThree := UnTwo(BindLastOfTwo(bindFirstTwoOf3, "b"))
		assert.That(t, assert.Equal("abc", boundSecondOfThree("a", "c")))
	})

	t.Run("bind first two of three", func(t *testing.T) {
		boundFirstTwo := bindFirstTwoOf3("a", "b")
		assert.That(t, assert.Equal("abc", boundFirstTwo("c")))
	})

}

func TestDifferentTypes(t *testing.T) {

	f1 := func(int, string) bool {
		return true
	}
	var b bool
	b = Two(f1)(42)("abc")
	assert.That(t, assert.True(b))

	testErr := errors.New("err")
	f1e := func(int, string) (bool, error) {
		return true, testErr
	}
	var err error
	b, err = Two2(f1e)(42)("abc")
	assert.That(t, assert.True(b))
	assert.That(t, assert.ErrorIs(err, testErr))

	b = UnTwo(Two(f1))(42, "abc")
	assert.That(t, assert.True(b))

	b, err = UnTwo2(Two2(f1e))(42, "abc")
	assert.That(t, assert.True(b))
	assert.That(t, assert.ErrorIs(err, testErr))

	err = DropFirstOfTwo(f1e(42, "abc"))
	assert.That(t, assert.ErrorIs(err, testErr))
	b = DropLastOfTwo(f1e(42, "abc"))
	assert.That(t, assert.True(b))
}
