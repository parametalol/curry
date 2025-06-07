package curry

import (
	"context"
	"errors"
	"testing"
)

func TestAdaptNone(t *testing.T) {
	Assert(t,
		NoError(AdaptNone[error](func() {})()),
		Equal("string", AdaptNone[string](func() string { return "string" })()),
	)
}

func TestAdaptOne(t *testing.T) {
	Assert(t,
		NoError(AdaptOne[string, error](func() {})("string")),
		Equal("result", AdaptOne[string, string](func() string { return "result" })("ignored")),
		Equal("passed result", AdaptOne[string, string](func(passed string) string { return passed + " result" })("passed")),
		Equal("", AdaptOne[string, string](func(string) {})("ignored")),
	)
}

func TestAdaptTwo(t *testing.T) {
	s := "string"

	testErr := errors.New("test")

	Assert(t,
		NoError(AdaptTwo[string, int, error](func() {})(s, 0)),
		ErrorIs(AdaptTwo[string, int, error](func() error { return testErr })(s, 0), testErr),
		Equal(5, AdaptTwo[int, int, int](func(a, b int) int { return a + b })(2, 3)),
		Equal("abc", AdaptTwo[int, string, string](func(b string) string { return b })(2, "abc")),
	)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	Assert(t, NoError(AdaptTwo[context.Context, int, error](func(context.Context) {})(ctx, 0)))
	Assert(t, ErrorIs(AdaptTwo[context.Context, int, error](func(ctx context.Context) error { return ctx.Err() })(ctx, 0), context.Canceled))

	Assert(t, NoError(AdaptTwo[context.Context, int, error](func(context.Context, int) {})(ctx, 0)))
	Assert(t, NoError(AdaptTwo[context.Context, int, error](func(int) {})(ctx, 0)))
}
