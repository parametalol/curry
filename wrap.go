package curry

import (
	"slices"
)

// Wrap returns a function that runs start, processes the start result with
// the provided process functions, and returns the result of the last process.
func Wrap[RV, A, B any](start func(A) B, process ...func(B) RV) func(A) RV {
	return func(a A) (result RV) {
		v := start(a)
		for proc := range slices.Values(process) {
			result = proc(v)
		}
		return
	}
}
