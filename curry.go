// curry is a generic function currying framework in Go.
// It takes a function with multiple arguments and returns a sequence of functions,
// each taking a single argument.
package curry

// region Curry

// Two curries a function with two arguments.
func Two0[A, B any](f func(A, B)) func(A) func(B) {
	return func(a A) func(B) {
		return func(b B) {
			f(a, b)
		}
	}
}

// Three curries a function with three arguments.
func Three0[A, B, C any](f func(A, B, C)) func(A) func(B) func(C) {
	return func(a A) func(B) func(C) {
		return func(b B) func(C) {
			return func(c C) {
				f(a, b, c)
			}
		}
	}
}

// Four curries a function with four arguments.
func Four0[A, B, C, D any](f func(A, B, C, D)) func(A) func(B) func(C) func(D) {
	return func(a A) func(B) func(C) func(D) {
		return func(b B) func(C) func(D) {
			return func(c C) func(D) {
				return func(d D) {
					f(a, b, c, d)
				}
			}
		}
	}
}

// Two curries a function with two arguments.
func Two[A, B, R any](f func(A, B) R) func(A) func(B) R {
	return func(a A) func(B) R {
		return func(b B) R {
			return f(a, b)
		}
	}
}

// Three curries a function with three arguments.
func Three[A, B, C, R any](f func(A, B, C) R) func(A) func(B) func(C) R {
	return func(a A) func(B) func(C) R {
		return func(b B) func(C) R {
			return func(c C) R {
				return f(a, b, c)
			}
		}
	}
}

// Four curries a function with four arguments.
func Four[A, B, C, D, R any](f func(A, B, C, D) R) func(A) func(B) func(C) func(D) R {
	return func(a A) func(B) func(C) func(D) R {
		return func(b B) func(C) func(D) R {
			return func(c C) func(D) R {
				return func(d D) R {
					return f(a, b, c, d)
				}
			}
		}
	}
}

// region Curry with 2 RV

// Two2 curries a function with two arguments and two return values.
func Two2[A, B, RA, RB any](f func(A, B) (RA, RB)) func(A) func(B) (RA, RB) {
	return func(a A) func(B) (RA, RB) {
		return func(b B) (RA, RB) {
			return f(a, b)
		}
	}
}

// Three2 curries a function with three arguments and two return values.
func Three2[A, B, C, RA, RB any](f func(A, B, C) (RA, RB)) func(A) func(B) func(C) (RA, RB) {
	return func(a A) func(B) func(C) (RA, RB) {
		return func(b B) func(C) (RA, RB) {
			return func(c C) (RA, RB) {
				return f(a, b, c)
			}
		}
	}
}

// Four2 curries a function with four arguments and two return values.
func Four2[A, B, C, D, RA, RB any](f func(A, B, C, D) (RA, RB)) func(A) func(B) func(C) func(D) (RA, RB) {
	return func(a A) func(B) func(C) func(D) (RA, RB) {
		return func(b B) func(C) func(D) (RA, RB) {
			return func(c C) func(D) (RA, RB) {
				return func(d D) (RA, RB) {
					return f(a, b, c, d)
				}
			}
		}
	}
}

// region Curry*Slice

func TwoSlice[A, B, R any](f func(A, ...B) R) func(A) func(...B) R {
	return func(a A) func(...B) R {
		return func(b ...B) R {
			return f(a, b...)
		}
	}
}

func ThreeSlice[A, B, C, R any](f func(A, B, ...C) R) func(A) func(B) func(...C) R {
	return func(a A) func(B) func(...C) R {
		return func(b B) func(...C) R {
			return func(c ...C) R {
				return f(a, b, c...)
			}
		}
	}
}

func TwoSlice2[A, B, RA, RB any](f func(A, ...B) (RA, RB)) func(A) func(...B) (RA, RB) {
	return func(a A) func(...B) (RA, RB) {
		return func(b ...B) (RA, RB) {
			return f(a, b...)
		}
	}
}

func ThreeSlice2[A, B, C, RA, RB any](f func(A, B, ...C) (RA, RB)) func(A) func(B) func(...C) (RA, RB) {
	return func(a A) func(B) func(...C) (RA, RB) {
		return func(b B) func(...C) (RA, RB) {
			return func(c ...C) (RA, RB) {
				return f(a, b, c...)
			}
		}
	}
}
