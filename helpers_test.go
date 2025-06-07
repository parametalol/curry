package curry

import (
	"errors"
	"testing"
)

type Assertion func() (bool, any, any)

func True(actual bool) Assertion {
	return Equal(true, actual)
}

func False(actual bool) Assertion {
	return Equal(false, actual)
}

func Equal[T comparable](expected, actual T) Assertion {
	return func() (bool, any, any) { return expected == actual, expected, actual }
}

func ErrorIs(err, target error) Assertion {
	return func() (bool, any, any) { return errors.Is(err, target), target, err }
}

func NoError(err error) Assertion {
	return func() (bool, any, any) { return err == nil, nil, err }
}

func Assert(t *testing.T, assertions ...Assertion) bool {
	t.Helper()
	ok := true
	for i, assertion := range assertions {
		if passed, e, a := assertion(); !passed {
			if len(assertions) > 0 {
				t.Errorf("assertion %d: expected = %v, got = %v", i, e, a)
			} else {
				t.Errorf("expected = %v, got = %v", e, a)
			}
			t.Fail()
			ok = false
		}
	}
	return ok
}
