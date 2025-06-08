package curry

import (
	"strconv"
	"strings"
	"testing"
)

func TestReturn(t *testing.T) {
	Assert(t, Equal("abc", Return("abc")()))
	a, b := Return2("abc", 123)()
	Assert(t, Equal("abc", a))
	Assert(t, Equal(123, b))

	i := ReturnSlice(1, 2, 3)
	Assert(t, Equal(1, i()[0]))
	Assert(t, Equal(2, i()[1]))
	Assert(t, Equal(3, i()[2]))
}

func TestLazy0(t *testing.T) {
	t.Run("LazyOne0", func(t *testing.T) {
		var result string
		compute := func(s string) { result = s }
		LazyOne0(compute)(Return("string"))
		Assert(t, Equal("string", result))
	})
	t.Run("LazyTwo0", func(t *testing.T) {
		var result string
		compute := func(s1, s2 string) {
			result = s1 + s2
		}
		LazyTwo0(compute)(Return("abc"), Return("def"))
		Assert(t, Equal("abcdef", result))
	})
	t.Run("LazyThree0", func(t *testing.T) {
		var result string
		compute := func(s1, s2 string, i int) {
			result = s1 + s2 + strconv.FormatInt(int64(i), 16)
		}
		LazyThree0(compute)(Return("abc"), Return("def"), Return(123))
		Assert(t, Equal("abcdef7b", result))
	})
	t.Run("LazyFour0", func(t *testing.T) {
		var result string
		compute := func(s1, s2 string, i int, b bool) {
			result = s1 + s2 + strconv.FormatInt(int64(i), 16) + strconv.FormatBool(b)
		}
		LazyFour0(compute)(Return("abc"), Return("def"), Return(123), Return(true))
		Assert(t, Equal("abcdef7btrue", result))
	})
	t.Run("LazyAll0", func(t *testing.T) {
		var result string
		compute := func(a ...string) {
			result = strings.Join(a, ", ")
		}
		LazyAll0(compute)(ReturnSlice("abc", "def"))
		Assert(t, Equal("abc, def", result))
	})
}

func TestLazy(t *testing.T) {
	t.Run("LazyOne", func(t *testing.T) {
		compute := func(s string) string {
			return s + s
		}
		Assert(t, Equal("stringstring",
			LazyOne(compute)(Return("string"))))
	})
	t.Run("LazyTwo", func(t *testing.T) {
		compute := func(s1, s2 string) string {
			return s1 + s2
		}
		Assert(t, Equal("abcdef",
			LazyTwo(compute)(Return("abc"), Return("def"))))
	})
	t.Run("LazyThree", func(t *testing.T) {
		compute := func(s1, s2 string, i int) string {
			return s1 + s2 + strconv.FormatInt(int64(i), 16)
		}
		Assert(t, Equal("abcdef7b",
			LazyThree(compute)(Return("abc"), Return("def"), Return(123))))
	})
	t.Run("LazyFour", func(t *testing.T) {
		compute := func(s1, s2 string, i int, b bool) string {
			return s1 + s2 + strconv.FormatInt(int64(i), 16) + strconv.FormatBool(b)
		}
		Assert(t, Equal("abcdef7btrue",
			LazyFour(compute)(Return("abc"), Return("def"), Return(123), Return(true))))
	})
	t.Run("LazyAll", func(t *testing.T) {
		compute := func(a ...string) string {
			return strings.Join(a, ", ")
		}
		Assert(t, Equal("abc, def",
			LazyAll(compute)(ReturnSlice("abc", "def"))))
	})
}

func TestLazy2(t *testing.T) {
	t.Run("LazyOne2", func(t *testing.T) {
		compute := func(s string) (string, bool) {
			return s + s, true
		}
		Assert(t, True(DropFirstOfTwo(
			LazyOne2(compute)(Return("string")))))
	})
	t.Run("LazyTwo2", func(t *testing.T) {
		compute := func(s1, s2 string) (string, bool) {
			return s1 + s2, true
		}
		Assert(t, True(DropFirstOfTwo(
			LazyTwo2(compute)(Return("abc"), Return("def")))))
	})
	t.Run("LazyThree2", func(t *testing.T) {
		compute := func(s1, s2 string, i int) (string, bool) {
			return s1 + s2 + strconv.FormatInt(int64(i), 16), true
		}
		Assert(t, True(DropFirstOfTwo(
			LazyThree2(compute)(Return("abc"), Return("def"), Return(123)))))
	})
	t.Run("LazyFour2", func(t *testing.T) {
		compute := func(s1, s2 string, i int, b bool) (string, bool) {
			return s1 + s2 + strconv.FormatInt(int64(i), 16) + strconv.FormatBool(b), true
		}
		Assert(t, True(DropFirstOfTwo(
			LazyFour2(compute)(Return("abc"), Return("def"), Return(123), Return(true)))))
	})
	t.Run("LazyAll2", func(t *testing.T) {
		compute := func(a ...string) (string, bool) {
			return strings.Join(a, ", "), true
		}
		Assert(t, True(DropFirstOfTwo(
			LazyAll2(compute)(ReturnSlice("abc", "def")))))
	})
}
