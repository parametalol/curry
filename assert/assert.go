package assert

import (
	"errors"
	"fmt"
	"slices"
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

func EqualSlices[U comparable](expected, actual []U) Assertion {
	return EqualFunc(expected, actual, slices.Equal)
}

func EqualFunc[T any](expected, actual T, cmp func(a, b T) bool) Assertion {
	return func() (bool, any, any) { return cmp(expected, actual), expected, actual }
}

func ErrorIs(err, target error) Assertion {
	return func() (bool, any, any) { return errors.Is(err, target), target, err }
}

func NoError(err error) Assertion {
	return func() (bool, any, any) { return err == nil, nil, err }
}

func Not(ass Assertion) Assertion {
	return func() (bool, any, any) {
		b, e, a := ass()
		return !b, fmt.Sprintf("not %v", e), a
	}
}

func That(t *testing.T, assertions ...Assertion) bool {
	t.Helper()
	ok := true
	for i, assertion := range assertions {
		if passed, e, a := assertion(); !passed {
			if len(assertions) > 0 {
				t.Errorf("assertion %d: expected %v, got %v", i, e, a)
			} else {
				t.Errorf("expected %v, got %v", e, a)
			}
			t.Fail()
			ok = false
		}
	}
	return ok
}
