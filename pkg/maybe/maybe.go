package maybe

type Maybe[T any] struct {
	v     T
	valid bool
}

func NewJust[T any](v T) Maybe[T] {
	return Maybe[T]{
		v:     v,
		valid: true,
	}
}

// NewNone used for explicit none value
func NewNone[T any]() Maybe[T] {
	return Maybe[T]{}
}

func Valid[T any](maybe Maybe[T]) bool {
	return maybe.valid
}

func Just[T any](maybe Maybe[T]) T {
	return maybe.v
}

func FromPtr[T any](t *T) Maybe[T] {
	if t == nil {
		return NewNone[T]()
	}
	return NewJust[T](*t)
}

func ToPtr[T any](m Maybe[T]) *T {
	if m.valid {
		return &m.v
	}
	return nil
}
