package curry

import (
	"context"
	"errors"
	"testing"

	"github.com/parametalol/curry/assert"
)

func TestAdaptNone(t *testing.T) {
	assert.That(t,
		assert.NoError(AdaptNone[error](func() {})()),
		assert.Equal("string", AdaptNone[string](func() string { return "string" })()),
	)
}

func TestAdaptOne(t *testing.T) {
	assert.That(t,
		assert.NoError(AdaptOne[string, error](func() {})("string")),
		assert.Equal("result", AdaptOne[string, string](func() string { return "result" })("ignored")),
		assert.Equal("passed result", AdaptOne[string, string](func(passed string) string { return passed + " result" })("passed")),
		assert.Equal("", AdaptOne[string, string](func(string) {})("ignored")),
	)
}

func TestAdaptTwo(t *testing.T) {
	s := "string"

	testErr := errors.New("test")

	assert.That(t,
		assert.NoError(AdaptTwo[string, int, error](func() {})(s, 0)),
		assert.ErrorIs(AdaptTwo[string, int, error](func() error { return testErr })(s, 0), testErr),
		assert.Equal(5, AdaptTwo[int, int, int](func(a, b int) int { return a + b })(2, 3)),
		assert.Equal("abc", AdaptTwo[int, string, string](func(b string) string { return b })(2, "abc")),
	)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	assert.That(t, assert.NoError(AdaptTwo[context.Context, int, error](func(context.Context) {})(ctx, 0)))
	assert.That(t, assert.ErrorIs(AdaptTwo[context.Context, int, error](func(ctx context.Context) error { return ctx.Err() })(ctx, 0), context.Canceled))

	assert.That(t, assert.NoError(AdaptTwo[context.Context, int, error](func(context.Context, int) {})(ctx, 0)))
	assert.That(t, assert.NoError(AdaptTwo[context.Context, int, error](func(int) {})(ctx, 0)))
}
