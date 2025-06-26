package curry

import (
	"strconv"
	"strings"
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestLazy(t *testing.T) {
	t.Run("Lazy1", func(t *testing.T) {
		var result string
		compute := func(s string) { result = s }
		Lazy1(compute)(Return("string"))
		assert.That(t, assert.Equal("string", result))
	})
	t.Run("Lazy2", func(t *testing.T) {
		var result string
		compute := func(s1, s2 string) {
			result = s1 + s2
		}
		Lazy2(compute)(Return("abc"), Return("def"))
		assert.That(t, assert.Equal("abcdef", result))
	})
	t.Run("Lazy3", func(t *testing.T) {
		var result string
		compute := func(s1, s2 string, i int) {
			result = s1 + s2 + strconv.FormatInt(int64(i), 16)
		}
		Lazy3(compute)(Return("abc"), Return("def"), Return(123))
		assert.That(t, assert.Equal("abcdef7b", result))
	})
	t.Run("Lazy4", func(t *testing.T) {
		var result string
		compute := func(s1, s2 string, i int, b bool) {
			result = s1 + s2 + strconv.FormatInt(int64(i), 16) + strconv.FormatBool(b)
		}
		Lazy4(compute)(Return("abc"), Return("def"), Return(123), Return(true))
		assert.That(t, assert.Equal("abcdef7btrue", result))
	})
	t.Run("LazyAll0", func(t *testing.T) {
		var result string
		compute := func(a ...string) {
			result = strings.Join(a, ", ")
		}
		Lazy1S(compute)(ReturnS("abc", "def"))
		assert.That(t, assert.Equal("abc, def", result))
	})
}

func TestLazyR(t *testing.T) {
	t.Run("Lazy1R", func(t *testing.T) {
		compute := func(s string) string {
			return s + s
		}
		assert.That(t, assert.Equal("stringstring",
			Lazy1R(compute)(Return("string"))))
	})
	t.Run("Lazy2R", func(t *testing.T) {
		compute := func(s1, s2 string) string {
			return s1 + s2
		}
		assert.That(t, assert.Equal("abcdef",
			Lazy2R(compute)(Return("abc"), Return("def"))))
	})
	t.Run("Lazy3R", func(t *testing.T) {
		compute := func(s1, s2 string, i int) string {
			return s1 + s2 + strconv.FormatInt(int64(i), 16)
		}
		assert.That(t, assert.Equal("abcdef7b",
			Lazy3R(compute)(Return("abc"), Return("def"), Return(123))))
	})
	t.Run("Lazy4R", func(t *testing.T) {
		compute := func(s1, s2 string, i int, b bool) string {
			return s1 + s2 + strconv.FormatInt(int64(i), 16) + strconv.FormatBool(b)
		}
		assert.That(t, assert.Equal("abcdef7btrue",
			Lazy4R(compute)(Return("abc"), Return("def"), Return(123), Return(true))))
	})
	t.Run("LazyAll", func(t *testing.T) {
		compute := func(a ...string) string {
			return strings.Join(a, ", ")
		}
		assert.That(t, assert.Equal("abc, def",
			Lazy1SR(compute)(ReturnS("abc", "def"))))
	})
}

func TestLazyR2(t *testing.T) {
	t.Run("Lazy1R2", func(t *testing.T) {
		compute := func(s string) (string, bool) {
			return s + s, true
		}
		assert.That(t, assert.True(DropFirstOf2(
			Lazy1R2(compute)(Return("string")))))
	})
	t.Run("LazyTwo2", func(t *testing.T) {
		compute := func(s1, s2 string) (string, bool) {
			return s1 + s2, true
		}
		assert.That(t, assert.True(DropFirstOf2(
			Lazy2R2(compute)(Return("abc"), Return("def")))))
	})
	t.Run("LazyThree2", func(t *testing.T) {
		compute := func(s1, s2 string, i int) (string, bool) {
			return s1 + s2 + strconv.FormatInt(int64(i), 16), true
		}
		assert.That(t, assert.True(DropFirstOf2(
			Lazy3R2(compute)(Return("abc"), Return("def"), Return(123)))))
	})
	t.Run("LazyFour2", func(t *testing.T) {
		compute := func(s1, s2 string, i int, b bool) (string, bool) {
			return s1 + s2 + strconv.FormatInt(int64(i), 16) + strconv.FormatBool(b), true
		}
		assert.That(t, assert.True(DropFirstOf2(
			Lazy4R2(compute)(Return("abc"), Return("def"), Return(123), Return(true)))))
	})
	t.Run("LazyAll2", func(t *testing.T) {
		compute := func(a ...string) (string, bool) {
			return strings.Join(a, ", "), true
		}
		assert.That(t, assert.True(DropFirstOf2(
			Lazy1SR2(compute)(ReturnS("abc", "def")))))
	})
}
