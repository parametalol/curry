package curry

import (
	"errors"
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

func TestDrop(t *testing.T) {
	Assert(t, Equal(1, DropLastOfTwo(1, 2)))
	Assert(t, Equal(1, DropLastOfTwo(DropLastOfThree(1, 2, 3))))
	Assert(t, Equal(1, DropLastOfTwo(DropLastOfThree(DropLastOfFour(1, 2, 3, 4)))))

	Assert(t, Equal(2, DropFirstOfTwo(1, 2)))
	Assert(t, Equal(3, DropFirstOfTwo(DropFirstOfThree(1, 2, 3))))
	Assert(t, Equal(4, DropFirstOfTwo(DropFirstOfThree(DropFirstOfFour(1, 2, 3, 4)))))
}

func TestCurry(t *testing.T) {
	Assert(t, Equal("ab", Two(join2)("a")("b")))
	Assert(t, Equal("abc", Three(join3)("a")("b")("c")))
	Assert(t, Equal("abcd", Four(join4)("a")("b")("c")("d")))

	Assert(t, Equal("ab", DropLastOfTwo(Two2(join2e)("a")("b"))))
	Assert(t, Equal("abc", DropLastOfTwo(Three2(join3e)("a")("b")("c"))))
	Assert(t, Equal("abcd", DropLastOfTwo(Four2(join4e)("a")("b")("c")("d"))))
}

func TestUnCurry(t *testing.T) {
	curriedJoin := Four(join4)
	Assert(t, Equal("abcd", UnFour(curriedJoin)("a", "b", "c", "d")))
	Assert(t, Equal("abcd", UnThree(curriedJoin)("a", "b", "c")("d")))
	Assert(t, Equal("abcd", UnTwo(curriedJoin)("a", "b")("c")("d")))

	curriedJoinE := Four2(join4e)
	Assert(t, Equal("abcd", DropLastOfTwo(UnFour2(curriedJoinE)("a", "b", "c", "d"))))
	Assert(t, Equal("abcd", DropLastOfTwo(UnThree(curriedJoinE)("a", "b", "c")("d"))))
	Assert(t, Equal("abcd", DropLastOfTwo(UnTwo(curriedJoinE)("a", "b")("c")("d"))))
}

func TestBind(t *testing.T) {
	Assert(t, Equal("a", BindOne(join1, "a")()))
	Assert(t, Equal("ab", BindFirstOfTwo(join2, "a")("b")))
	Assert(t, Equal("abc", BindFirstOfThree(join3, "a")("b", "c")))
	Assert(t, Equal("abcd", BindFirstOfFour(join4, "a")("b", "c", "d")))

	Assert(t, Equal("a", DropLastOfTwo(BindOne2(join1e, "a")())))
	Assert(t, Equal("ab", DropLastOfTwo(BindFirstOfTwo2(join2e, "a")("b"))))
	Assert(t, Equal("abc", DropLastOfTwo(BindFirstOfThree2(join3e, "a")("b", "c"))))
	Assert(t, Equal("abcd", DropLastOfTwo(BindFirstOfFour2(join4e, "a")("b", "c", "d"))))

	Assert(t, Equal("ba", BindLastOfTwo(join2, "a")("b")))
	Assert(t, Equal("bca", BindLastOfThree(join3, "a")("b", "c")))
	Assert(t, Equal("bcda", BindLastOfFour(join4, "a")("b", "c", "d")))

	Assert(t, Equal("ba", DropLastOfTwo(BindLastOfTwo2(join2e, "a")("b"))))
	Assert(t, Equal("bca", DropLastOfTwo(BindLastOfThree2(join3e, "a")("b", "c"))))
	Assert(t, Equal("bcda", DropLastOfTwo(BindLastOfFour2(join4e, "a")("b", "c", "d"))))
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
