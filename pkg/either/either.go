package either

type discriminator uint

const (
	left = iota
	right
)

func NewLeft[L any, R any](l L) Either[L, R] {
	return Either[L, R]{
		l: l,
		d: left,
	}
}

func NewRight[L any, R any](r R) Either[L, R] {
	return Either[L, R]{
		r: r,
		d: right,
	}
}

type Either[L any, R any] struct {
	l L
	r R
	d discriminator
}

func (v Either[L, R]) MapLeft(f func(l L)) {
	if v.d == left {
		f(v.l)
	}
}

func (v Either[L, R]) MapRight(f func(r R)) {
	if v.d == right {
		f(v.r)
	}
}

func (v Either[T, E]) IsLeft() bool {
	return v.d == left
}

func (v Either[T, E]) IsRight() bool {
	return v.d == right
}

func (v Either[T, E]) Left() T {
	if v.d != left {
		panic("violated usage of either")
	}
	return v.l
}

func (v Either[T, E]) Right() E {
	if v.d != right {
		panic("violated usage of either")
	}
	return v.r
}
