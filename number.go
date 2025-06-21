package curry

// TODO: replace eventually with the standart Number constraint.
type number interface {
	~uintptr |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

type NumberT[N number] struct {
	N N
}

func Number[N number](n N) *NumberT[N] {
	return &NumberT[N]{n}
}

func (i *NumberT[N]) Inc() N {
	i.N++
	return i.N
}

func (i *NumberT[N]) Dec() N {
	i.N--
	return i.N
}

func (i *NumberT[N]) Add(n N) N {
	i.N += n
	return i.N
}
