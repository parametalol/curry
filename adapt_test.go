package curry

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestAdaptNone(t *testing.T) {
	assert.That(t,
		assert.NoError(AdaptNone[error](func() {})()),
		assert.Equal("string", AdaptNone[string](func() string { return "string" })()),
		assert.Equal("string", AdaptNone[string](func() any { return "string" })()),
		assert.Equal("string", AdaptNoneF(Thunk(""), func() any { return "string" })()),
	)
}

func TestAdaptOne(t *testing.T) {
	assert.That(t,
		assert.NoError(AdaptOne[string, error](func() {})("string")),
		assert.Equal("result", AdaptOne[string, string](func() string { return "result" })("ignored")),
		assert.Equal("passed result", AdaptOne[string, string](func(passed string) string { return passed + " result" })("passed")),
		assert.Equal("passed", AdaptOne[string, string](func(passed any) string { return fmt.Sprint(passed) })("passed")),
		assert.Equal("passed", AdaptOne[string, string](func(passed any) any { return fmt.Sprint(passed) })("passed")),
		assert.Equal("passed", AdaptOne[string, string](func(passed string) any { return fmt.Sprint(passed) })("passed")),
		assert.Equal("passed", AdaptOne[string, string](func() any { return "passed" })("ignored")),
		assert.Equal("", AdaptOne[string, string](func(string) {})("ignored")),
		assert.Equal("", AdaptOneF(strings.Clone, func(string) {})("ignored")),
		assert.Equal("", AdaptOneF(strings.Clone, func(any) {})("ignored")),
	)
}

func TestAdaptTwo(t *testing.T) {
	s := "string"

	testErr := errors.New("test")

	assert.That(t,
		assert.NoError(AdaptTwo[string, int, error](func() {})(s, 0)),
		assert.ErrorIs(AdaptTwo[string, int, error](func() error { return testErr })(s, 0), testErr),
		assert.Equal(5, AdaptTwo[int, int, int](func(a, b int) int { return a + b })(2, 3)),
		assert.Equal(5, AdaptTwo[int, int, int](func(a, b int) any { return a + b })(2, 3)),
		assert.Equal(5, AdaptTwo[int, int, int](func(a int, b any) int { return a + b.(int) })(2, 3)),
		assert.Equal(5, AdaptTwo[int, int, int](func(a int, b any) any { return a + b.(int) })(2, 3)),
		assert.Equal(0, AdaptTwo[int, int, int](func(_ int, _ any) {})(2, 3)),
		assert.Equal(5, AdaptTwo[int, int, int](func(a any, b int) int { return a.(int) + b })(2, 3)),
		assert.Equal(5, AdaptTwo[int, int, int](func(a any, b int) any { return a.(int) + b })(2, 3)),
		assert.Equal(0, AdaptTwo[int, int, int](func(_ any, _ int) {})(2, 3)),
		assert.Equal(5, AdaptTwo[int, int, int](func(a, b any) int { return a.(int) + b.(int) })(2, 3)),
		assert.Equal(5, AdaptTwo[int, int, int](func(a, b any) any { return a.(int) + b.(int) })(2, 3)),
		assert.Equal(0, AdaptTwo[int, int, int](func(_, _ any) {})(2, 3)),
		assert.Equal(2, AdaptTwo[int, int, int](func(a int) int { return a })(2, 3)),
		assert.Equal(2, AdaptTwo[int, int, int](func(a int) any { return a })(2, 3)),
		assert.Equal(2, AdaptTwo[int, int, int](func(a any) int { return a.(int) })(2, 3)),
		assert.Equal(2, AdaptTwo[int, int, int](func(a any) any { return a.(int) })(2, 3)),
		assert.Equal(0, AdaptTwo[int, int, int](func(_ any) {})(2, 3)),
		assert.Equal(0, AdaptTwo[int, int, int](func(_ int) {})(2, 3)),
		assert.Equal(100, AdaptTwo[int, int, int](func() any { return 100 })(2, 3)),
		assert.Equal("abc", AdaptTwo[int, string, string](func(b string) string { return b })(2, "abc")),
		assert.Equal("abc", AdaptTwo[int, string, string](func(b string) any { return b })(2, "abc")),
		assert.Equal("abc", AdaptTwoF(func(int, string) string { return "" }, func(b string) string { return b })(2, "abc")),
	)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	assert.That(t, assert.NoError(AdaptTwo[context.Context, int, error](func(context.Context) {})(ctx, 0)))
	assert.That(t, assert.ErrorIs(AdaptTwo[context.Context, int, error](func(ctx context.Context) error { return ctx.Err() })(ctx, 0), context.Canceled))

	assert.That(t, assert.NoError(AdaptTwo[context.Context, int, error](func(context.Context, int) {})(ctx, 0)))
	assert.That(t, assert.NoError(AdaptTwo[context.Context, int, error](func(int) {})(ctx, 0)))
}
