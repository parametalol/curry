package curry

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func join2(a, b string) string {
	return a + b
}

func join3(a, b, c string) string {
	return a + b + c
}

func join4(a, b, c, d string) string {
	return a + b + c + d
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

func TestDrop(t *testing.T) {
	assert.Equal(t, 1, DropLast2(1, 2))
	assert.Equal(t, 1, DropLast2(DropLast3(1, 2, 3)))
	assert.Equal(t, 1, DropLast2(DropLast3(DropLast4(1, 2, 3, 4))))

	assert.Equal(t, 2, DropFirst2(1, 2))
	assert.Equal(t, 3, DropFirst2(DropFirst3(1, 2, 3)))
	assert.Equal(t, 4, DropFirst2(DropFirst3(DropFirst4(1, 2, 3, 4))))
}

func TestCurry(t *testing.T) {
	assert.Equal(t, "ab", Curry2(join2)("a")("b"))
	assert.Equal(t, "abc", Curry3(join3)("a")("b")("c"))
	assert.Equal(t, "abcd", Curry4(join4)("a")("b")("c")("d"))

	assert.Equal(t, "ab", DropLast2(Curry2e(join2e)("a")("b")))
	assert.Equal(t, "abc", DropLast2(Curry3e(join3e)("a")("b")("c")))
	assert.Equal(t, "abcd", DropLast2(Curry4e(join4e)("a")("b")("c")("d")))
}

func TestUnCurry(t *testing.T) {
	curriedJoin := Curry4(join4)
	assert.Equal(t, "abcd", UnCurry4(curriedJoin)("a", "b", "c", "d"))
	assert.Equal(t, "abcd", UnCurry3(curriedJoin)("a", "b", "c")("d"))
	assert.Equal(t, "abcd", UnCurry2(curriedJoin)("a", "b")("c")("d"))

	curriedJoinE := Curry4e(join4e)
	assert.Equal(t, "abcd", DropLast2(UnCurry4e(curriedJoinE)("a", "b", "c", "d")))
	assert.Equal(t, "abcd", DropLast2(UnCurry3(curriedJoinE)("a", "b", "c")("d")))
	assert.Equal(t, "abcd", DropLast2(UnCurry2(curriedJoinE)("a", "b")("c")("d")))
}

func TestBind(t *testing.T) {
	assert.Equal(t, "ab", Bind1st2(join2, "a")("b"))
	assert.Equal(t, "abc", Bind1st3(join3, "a")("b", "c"))
	assert.Equal(t, "abcd", Bind1st4(join4, "a")("b", "c", "d"))

	assert.Equal(t, "ab", DropLast2(Bind1st2e(join2e, "a")("b")))
	assert.Equal(t, "abc", DropLast2(Bind1st3e(join3e, "a")("b", "c")))
	assert.Equal(t, "abcd", DropLast2(Bind1st4e(join4e, "a")("b", "c", "d")))

	assert.Equal(t, "ba", BindLast2(join2, "a")("b"))
	assert.Equal(t, "bca", BindLast3(join3, "a")("b", "c"))
	assert.Equal(t, "bcda", BindLast4(join4, "a")("b", "c", "d"))

	assert.Equal(t, "ba", DropLast2(BindLast2e(join2e, "a")("b")))
	assert.Equal(t, "bca", DropLast2(BindLast3e(join3e, "a")("b", "c")))
	assert.Equal(t, "bcda", DropLast2(BindLast4e(join4e, "a")("b", "c", "d")))
}

func TestCombinations(t *testing.T) {
	bindFirstTwoOf3 := UnCurry2(Curry3(join3))

	t.Run("bind 2nd of three", func(t *testing.T) {
		boundSecondOfThree := UnCurry2(BindLast2(bindFirstTwoOf3, "b"))
		assert.Equal(t, "abc", boundSecondOfThree("a", "c"))
	})

	t.Run("bind first two of three", func(t *testing.T) {
		boundFirstTwo := bindFirstTwoOf3("a", "b")
		assert.Equal(t, "abc", boundFirstTwo("c"))
	})

}

func TestDifferentTypes(t *testing.T) {

	f1 := func(int, string) bool {
		return true
	}
	var b bool = Curry2(f1)(42)("abc")
	assert.True(t, b)

	testErr := errors.New("err")
	f1e := func(int, string) (bool, error) {
		return true, testErr
	}
	var err error
	b, err = Curry2e(f1e)(42)("abc")
	assert.True(t, b)
	assert.ErrorIs(t, err, testErr)

	b = UnCurry2(Curry2(f1))(42, "abc")
	assert.True(t, b)

	b, err = UnCurry2e(Curry2e(f1e))(42, "abc")
	assert.True(t, b)
	assert.ErrorIs(t, err, testErr)

	err = DropFirst2(f1e(42, "abc"))
	assert.ErrorIs(t, err, testErr)
	b = DropLast2(f1e(42, "abc"))
	assert.True(t, b)
}
