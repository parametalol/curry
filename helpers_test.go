package curry

import (
	"errors"
	"testing"
)

func True(actual bool) func() (bool, bool, bool) {
	return Equal(true, actual)
}

func False(actual bool) func() (bool, bool, bool) {
	return Equal(false, actual)
}

func Equal[T comparable](expected, actual T) func() (bool, T, T) {
	return func() (bool, T, T) { return expected == actual, expected, actual }
}

func ErrorIs(err, target error) func() (bool, error, error) {
	return func() (bool, error, error) { return errors.Is(err, target), target, err }
}

func NoError(err error) func() (bool, error, error) {
	return func() (bool, error, error) { return err == nil, nil, err }
}

func Assert[T comparable](t *testing.T, assertion func() (bool, T, T)) bool {
	t.Helper()
	if ok, e, a := assertion(); !ok {
		t.Errorf("expected = %v, got = %v", e, a)
		t.Fail()
		return false
	}
	return true
}
