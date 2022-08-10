package maybesql

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/UsingCoding/fpgo/pkg/maybe"
)

// wrapper that combines two sql interfaces
type valuer interface {
	sql.Scanner
	driver.Valuer
}

type Maybe[T any] struct {
	maybe.Maybe[T]
	value valuer
}

func (m *Maybe[T]) Scan(src any) error {
	if src == nil {
		m.Maybe = maybe.NewNone[T]()
		return nil
	}

	if m.value == nil {
		t, ok := src.(T)
		if !ok {
			return errors.New("failed to scan")
		}

		m.Maybe = maybe.NewJust(t)
		return nil
	}

	err := m.value.Scan(src)
	if err != nil {
		return err
	}

	value, err := m.value.Value()
	if err != nil {
		return err
	}

	v, ok := value.(T)
	if !ok {
		return errors.New("failed to type cast")
	}

	m.Maybe = maybe.NewJust(v)
	return nil
}

func (m *Maybe[T]) Value() (driver.Value, error) {
	if !maybe.Valid(m.Maybe) {
		return nil, nil
	}
	if m.value == nil {
		return maybe.Just(m.Maybe), nil
	}
	return m.value.Value()
}

func FromMaybe[T any](m maybe.Maybe[T]) (Maybe[T], error) {
	if !maybe.Valid(m) {
		t := new(T)
		if implementsValue(t) {
			return Maybe[T]{
				Maybe: m,
				value: convertToValue(t),
			}, nil
		}

		v := sqlNullObjectFromInterface(t, false)

		return Maybe[T]{
			Maybe: m,
			value: v,
		}, nil
	}

	if implementsValue(maybe.Just(m)) {
		return Maybe[T]{
			Maybe: m,
			value: convertToValue(maybe.Just(m)),
		}, nil
	}

	v := sqlNullObjectFromInterface(maybe.Just(m), true)

	return Maybe[T]{
		Maybe: m,
		value: v,
	}, nil
}

func implementsValue(i interface{}) (ok bool) {
	_, ok = i.(valuer)
	return
}

func convertToValue(i interface{}) (v valuer) {
	v = i.(valuer)
	return
}

func sqlNullObjectFromInterface(t interface{}, valid bool) valuer {
	switch v := t.(type) {
	case string:
		return &sql.NullString{
			String: v,
			Valid:  valid,
		}
	case bool:
		return &sql.NullBool{
			Bool:  v,
			Valid: valid,
		}
	case byte:
		return &sql.NullByte{
			Byte:  v,
			Valid: valid,
		}
	case float64:
		return &sql.NullFloat64{
			Float64: v,
			Valid:   valid,
		}
	case int16:
		return &sql.NullInt16{
			Int16: v,
			Valid: valid,
		}
	case int32:
		return &sql.NullInt32{
			Int32: v,
			Valid: valid,
		}
	case int64:
		return &sql.NullInt64{
			Int64: v,
			Valid: valid,
		}
	case time.Time:
		return &sql.NullTime{
			Time:  v,
			Valid: valid,
		}
	default:
		return nil
	}
}
