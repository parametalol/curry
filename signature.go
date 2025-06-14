package curry

type FnSigNone[RV1 any] struct {
	RV1 RV1
}
type FnSig2[RV1, RV2 any] struct {
	FnSigNone[RV1]
	RV2 RV2
}

type FnSigOne0[A any] struct {
	Arg1 A
}
type FnSigOne[A, RV any] struct {
	FnSigOne0[A]
	FnSigNone[RV]
}
type FnSigOne2[A, RV1, RV2 any] struct {
	FnSigOne0[A]
	FnSig2[RV1, RV2]
}

type FnSigTwo0[A, B any] struct {
	FnSigOne0[A]
	Arg2 B
}
type FnSigTwo[A, B, RV any] struct {
	FnSigOne[A, RV]
	Arg2 B
}
type FnSigTwo2[A, B, RV1, RV2 any] struct {
	FnSigOne2[A, RV1, RV2]
	Arg2 B
}

type FnSigThree0[A, B, C any] struct {
	FnSigTwo0[A, B]
	Arg3 C
}
type FnSigThree[A, B, C, RV any] struct {
	FnSigTwo[A, B, RV]
	Arg3 C
}
type FnSigThree2[A, B, C, RV1, RV2 any] struct {
	FnSigTwo2[A, B, RV1, RV2]
	Arg3 C
}

type FnSigFour0[A, B, C, D any] struct {
	FnSigThree0[A, B, C]
	Arg4 D
}
type FnSigFour[A, B, C, D, RV any] struct {
	FnSigThree[A, B, C, RV]
	Arg4 D
}
type FnSigFour2[A, B, C, D, RV1, RV2 any] struct {
	FnSigThree2[A, B, C, RV1, RV2]
	Arg4 D
}

func SignatureNone[RV any](f func() RV) (d FnSigNone[RV])                   { return }
func SignatureNone2[RV1, RV2 any](f func() (RV1, RV2)) (d FnSig2[RV1, RV2]) { return }

func SignatureOne0[A any](f func(A)) (d FnSigOne0[A])                                { return }
func SignatureOne[A, RV any](f func(A) RV) (d FnSigOne[A, RV])                       { return }
func SignatureOne2[A, RV1, RV2 any](f func(A) (RV1, RV2)) (d FnSigOne2[A, RV1, RV2]) { return }

func SignatureTwo0[A, B any](f func(A, B)) (d FnSigTwo0[A, B])                                { return }
func SignatureTwo[A, B, RV any](f func(A, B) RV) (d FnSigTwo[A, B, RV])                       { return }
func SignatureTwo2[A, B, RV1, RV2 any](f func(A, B) (RV1, RV2)) (d FnSigTwo2[A, B, RV1, RV2]) { return }

func SignatureThree0[A, B, C any](f func(A, B, C)) (d FnSigThree0[A, B, C])          { return }
func SignatureThree[A, B, C, RV any](f func(A, B, C) RV) (d FnSigThree[A, B, C, RV]) { return }

func SignatureFour0[A, B, C, D any](f func(A, B, C, D)) (d FnSigFour0[A, B, C, D]) { return }
func SignatureFour[A, B, C, D, RV any](f func(A, B, C, D) RV) (d FnSigFour[A, B, C, D, RV]) {
	return
}
