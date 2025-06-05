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
	equal(t, 1, DropLastOfTwo(1, 2))
	equal(t, 1, DropLastOfTwo(DropLastOfThree(1, 2, 3)))
	equal(t, 1, DropLastOfTwo(DropLastOfThree(DropLastOfFour(1, 2, 3, 4))))

	equal(t, 2, DropFirstOfTwo(1, 2))
	equal(t, 3, DropFirstOfTwo(DropFirstOfThree(1, 2, 3)))
	equal(t, 4, DropFirstOfTwo(DropFirstOfThree(DropFirstOfFour(1, 2, 3, 4))))
}

func TestCurry(t *testing.T) {
	equal(t, "ab", Two(join2)("a")("b"))
	equal(t, "abc", Three(join3)("a")("b")("c"))
	equal(t, "abcd", Four(join4)("a")("b")("c")("d"))

	equal(t, "ab", DropLastOfTwo(Two2(join2e)("a")("b")))
	equal(t, "abc", DropLastOfTwo(Three2(join3e)("a")("b")("c")))
	equal(t, "abcd", DropLastOfTwo(Four2(join4e)("a")("b")("c")("d")))
}

func TestUnCurry(t *testing.T) {
	curriedJoin := Four(join4)
	equal(t, "abcd", UnFour(curriedJoin)("a", "b", "c", "d"))
	equal(t, "abcd", UnThree(curriedJoin)("a", "b", "c")("d"))
	equal(t, "abcd", UnTwo(curriedJoin)("a", "b")("c")("d"))

	curriedJoinE := Four2(join4e)
	equal(t, "abcd", DropLastOfTwo(UnFour2(curriedJoinE)("a", "b", "c", "d")))
	equal(t, "abcd", DropLastOfTwo(UnThree(curriedJoinE)("a", "b", "c")("d")))
	equal(t, "abcd", DropLastOfTwo(UnTwo(curriedJoinE)("a", "b")("c")("d")))
}

func TestBind(t *testing.T) {
	equal(t, "a", BindOne(join1, "a")())
	equal(t, "ab", BindFirstOfTwo(join2, "a")("b"))
	equal(t, "abc", BindFirstOfThree(join3, "a")("b", "c"))
	equal(t, "abcd", BindFirstOfFour(join4, "a")("b", "c", "d"))

	equal(t, "a", DropLastOfTwo(BindOne2(join1e, "a")()))
	equal(t, "ab", DropLastOfTwo(BindFirstOfTwo2(join2e, "a")("b")))
	equal(t, "abc", DropLastOfTwo(BindFirstOfThree2(join3e, "a")("b", "c")))
	equal(t, "abcd", DropLastOfTwo(BindFirstOfFour2(join4e, "a")("b", "c", "d")))

	equal(t, "ba", BindLastOfTwo(join2, "a")("b"))
	equal(t, "bca", BindLastOfThree(join3, "a")("b", "c"))
	equal(t, "bcda", BindLastOfFour(join4, "a")("b", "c", "d"))

	equal(t, "ba", DropLastOfTwo(BindLastOfTwo2(join2e, "a")("b")))
	equal(t, "bca", DropLastOfTwo(BindLastOfThree2(join3e, "a")("b", "c")))
	equal(t, "bcda", DropLastOfTwo(BindLastOfFour2(join4e, "a")("b", "c", "d")))
}

func TestCombinations(t *testing.T) {
	bindFirstTwoOf3 := UnTwo(Three(join3))

	t.Run("bind 2nd of three", func(t *testing.T) {
		boundSecondOfThree := UnTwo(BindLastOfTwo(bindFirstTwoOf3, "b"))
		equal(t, "abc", boundSecondOfThree("a", "c"))
	})

	t.Run("bind first two of three", func(t *testing.T) {
		boundFirstTwo := bindFirstTwoOf3("a", "b")
		equal(t, "abc", boundFirstTwo("c"))
	})

}

func TestDifferentTypes(t *testing.T) {

	f1 := func(int, string) bool {
		return true
	}
	var b bool
	b = Two(f1)(42)("abc")
	equal(t, true, b)

	testErr := errors.New("err")
	f1e := func(int, string) (bool, error) {
		return true, testErr
	}
	var err error
	b, err = Two2(f1e)(42)("abc")
	equal(t, true, b)
	equal(t, true, errors.Is(err, testErr))

	b = UnTwo(Two(f1))(42, "abc")
	equal(t, true, b)

	b, err = UnTwo2(Two2(f1e))(42, "abc")
	equal(t, true, b)
	equal(t, true, errors.Is(err, testErr))

	err = DropFirstOfTwo(f1e(42, "abc"))
	equal(t, true, errors.Is(err, testErr))
	b = DropLastOfTwo(f1e(42, "abc"))
	equal(t, true, b)
}

func equal[C comparable](t *testing.T, expected, actual C) {
	t.Helper()
	if expected == actual {
		return
	}
	t.Errorf("expected = %v, got = %v", expected, actual)
	t.Fail()
}
